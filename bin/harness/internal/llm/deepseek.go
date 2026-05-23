package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"harness/internal/config"
)

// Usage contains token counts from an API response.
type Usage struct {
	PromptTokens     int     `json:"prompt_tokens"`
	CompletionTokens int     `json:"completion_tokens"`
	TotalTokens      int     `json:"total_tokens"`
	CacheHitTokens   int     `json:"prompt_cache_hit_tokens,omitempty"`
	InputCost        float64 `json:"-"`
	OutputCost       float64 `json:"-"`
	TotalCost        float64 `json:"-"`
}

// formatCost returns a human-readable cost string.
func (u *Usage) Format() string {
	if u == nil || u.TotalTokens == 0 {
		return ""
	}
	costStr := ""
	if u.TotalCost > 0 {
		costStr = fmt.Sprintf(" | $%.6f", u.TotalCost)
	}
	cacheStr := ""
	if u.CacheHitTokens > 0 {
		cacheStr = fmt.Sprintf(" | cache +%d", u.CacheHitTokens)
	}
	return fmt.Sprintf("⎿  ⚡ %d → %d = %d tok%s%s",
		u.PromptTokens, u.CompletionTokens, u.TotalTokens, cacheStr, costStr)
}

// visionModels lists DeepSeek models that support multimodal (image) input.
var visionModels = map[string]bool{
	"deepseek-vl2": true, "deepseek-vl3": true,
	"deepseek-vision": true,
	"deepseek-vis":    true,
}

// DeepSeekProvider implements LLMProvider for the DeepSeek API.
type DeepSeekProvider struct {
	apiKey    string
	model     string
	lastUsage *Usage               // stores usage from streaming responses
	costs     map[string]config.ModelCost // per-model pricing, loaded from config
}

// NewDeepSeekProvider creates a new DeepSeek provider.
// If model is empty, "deepseek-chat" is used.
// modelCosts can be nil (cost tracking disabled) or from config.json.
func NewDeepSeekProvider(apiKey string, model string, modelCosts []config.ModelConfig) *DeepSeekProvider {
	if model == "" {
		model = "deepseek-chat"
	}
	d := &DeepSeekProvider{
		apiKey: apiKey,
		model:  model,
		costs:  make(map[string]config.ModelCost),
	}
	// Build cost map from config
	for _, mc := range modelCosts {
		if mc.Cost.Input > 0 || mc.Cost.Output > 0 {
			d.costs[mc.ID] = config.ModelCost{
				Input:     mc.Cost.Input,
				Output:    mc.Cost.Output,
				CacheRead: mc.Cost.CacheRead,
			}
		}
	}
	return d
}

// supportsVision returns true if the model is known to support image input.
func (d *DeepSeekProvider) supportsVision() bool {
	if visionModels[d.model] {
		return true
	}
	// Heuristic: models with "vl" or "vision" in the name support vision
	lower := strings.ToLower(d.model)
	return strings.Contains(lower, "vl") || strings.Contains(lower, "vision")
}

func (d *DeepSeekProvider) Name() string {
	return "deepseek"
}

func (d *DeepSeekProvider) Complete(ctx context.Context, prompt string, opts map[string]any) (string, error) {
	// Check for multimodal images in opts
	if images, ok := opts["images"].([]string); ok && len(images) > 0 {
		if d.supportsVision() {
			return d.completeMultimodal(ctx, prompt, images, opts)
		}
		// Model doesn't support vision: embed image as text data URI in the prompt
		prompt = prompt + "\n\n[📷 Imagem anexada (modelo " + d.model + " não suporta visão nativa):\n"
		for _, img := range images {
			// Truncate to first 200 chars of base64 to show it's there without blowing tokens
			preview := img
			if len(preview) > 200 {
				preview = preview[:200] + "...]"
			}
			prompt += preview + "\n"
		}
		prompt += "O modelo atual não processa imagens visualmente. Use <tool:image_info path=\"...\"/> para metadados ou troque para um modelo com visão (deepseek-vl2, deepseek-vl3)."
	}

	url := "https://api.deepseek.com/chat/completions"

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + d.apiKey,
	}

	body := fmt.Sprintf(`{"model":%s,"messages":[{"role":"user","content":%s}]}`, JSONString(d.model), JSONString(prompt))

	respBody, err := DoRequest(ctx, url, headers, []byte(body))
	if err != nil {
		return "", err
	}

	token := ExtractToken(string(respBody), "choices.0.message.content")
	if token == "" {
		return "", fmt.Errorf("deepseek: empty response or unexpected format")
	}

	// Parse usage and store in opts
	if usage := d.parseUsage(respBody); usage != nil {
		opts["_usage"] = usage
	}

	return token, nil
}

// parseUsage extracts token usage from the API response and calculates costs.
func (d *DeepSeekProvider) parseUsage(respBody []byte) *Usage {
	var rawResp struct {
		Usage *Usage `json:"usage"`
	}
	if err := json.Unmarshal(respBody, &rawResp); err != nil || rawResp.Usage == nil {
		return nil
	}
	u := rawResp.Usage

	// Calculate costs based on model pricing
	if cost, ok := d.costs[d.model]; ok {
		u.InputCost = float64(u.PromptTokens) * cost.Input / 1_000_000
		u.OutputCost = float64(u.CompletionTokens) * cost.Output / 1_000_000
		u.TotalCost = u.InputCost + u.OutputCost

		// Apply cache discount
		if u.CacheHitTokens > 0 && cost.CacheRead > 0 {
			cacheCost := float64(u.CacheHitTokens) * cost.CacheRead / 1_000_000
			u.TotalCost = float64(u.PromptTokens-u.CacheHitTokens)*cost.Input/1_000_000 +
				cacheCost + u.OutputCost
		}
	}

	return u
}

// completeMultimodal sends a multimodal message with text + images.
// DeepSeek supports the OpenAI-compatible vision format.
func (d *DeepSeekProvider) completeMultimodal(ctx context.Context, prompt string, images []string, opts map[string]any) (string, error) {
	url := "https://api.deepseek.com/chat/completions"

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + d.apiKey,
	}

	// Build content array: text + images
	var contentParts []map[string]interface{}

	// Add text part
	contentParts = append(contentParts, map[string]interface{}{
		"type": "text",
		"text": prompt,
	})

	// Add image parts
	for _, img := range images {
		// img should be: "data:image/png;base64,..." or "data:image/jpeg;base64,..."
		contentParts = append(contentParts, map[string]interface{}{
			"type": "image_url",
			"image_url": map[string]string{
				"url": img,
			},
		})
	}

	bodyMap := map[string]interface{}{
		"model": d.model,
		"messages": []map[string]interface{}{
			{
				"role":    "user",
				"content": contentParts,
			},
		},
	}

	bodyBytes, err := json.Marshal(bodyMap)
	if err != nil {
		return "", fmt.Errorf("deepseek multimodal marshal: %w", err)
	}

	respBody, err := DoRequest(ctx, url, headers, bodyBytes)
	if err != nil {
		return "", err
	}

	token := ExtractToken(string(respBody), "choices.0.message.content")
	if token == "" {
		return "", fmt.Errorf("deepseek: empty multimodal response")
	}

	// Parse usage
	if usage := d.parseUsage(respBody); usage != nil {
		opts["_usage"] = usage
	}

	return token, nil
}

func (d *DeepSeekProvider) Stream(ctx context.Context, prompt string, opts map[string]any) (<-chan string, error) {
	// Check for multimodal images in opts
	if images, ok := opts["images"].([]string); ok && len(images) > 0 {
		return d.streamMultimodal(ctx, prompt, images, opts)
	}

	url := "https://api.deepseek.com/chat/completions"

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + d.apiKey,
	}

	body := fmt.Sprintf(`{"model":%s,"messages":[{"role":"user","content":%s}],"stream":true}`, JSONString(d.model), JSONString(prompt))

	lines, err := DoStreamRequest(ctx, url, headers, []byte(body))
	if err != nil {
		return nil, err
	}

	ch := make(chan string, 32)
	go func() {
		defer close(ch)
		var lastData string
		for line := range lines {
			if strings.HasPrefix(line, "__error__:") {
				return
			}
			if strings.HasPrefix(line, "data:") {
				dataStr := strings.TrimSpace(line[5:])
				if dataStr == "" || dataStr == "[DONE]" {
					continue
				}
				lastData = dataStr
				token := ExtractToken(dataStr, "choices.0.delta.content")
				if token == "" {
					token = ExtractToken(dataStr, "choices.0.message.content")
				}
				if token != "" {
					ch <- token
				}
			}
		}
		// Parse usage from the last event (may contain usage instead of delta)
		if lastData != "" {
			d.lastUsage = d.parseUsage([]byte(lastData))
			// If lastData doesn't have usage, it might be in the previous data chunk
			// For now, usage in streaming is best-effort
		}
	}()

	return ch, nil
}

// streamMultimodal sends a streaming multimodal message with text + images.
func (d *DeepSeekProvider) streamMultimodal(ctx context.Context, prompt string, images []string, opts map[string]any) (<-chan string, error) {
	url := "https://api.deepseek.com/chat/completions"

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + d.apiKey,
	}

	// Build content array: text + images
	var contentParts []map[string]interface{}
	contentParts = append(contentParts, map[string]interface{}{"type": "text", "text": prompt})
	for _, img := range images {
		contentParts = append(contentParts, map[string]interface{}{
			"type": "image_url",
			"image_url": map[string]string{"url": img},
		})
	}

	bodyMap := map[string]interface{}{
		"model":    d.model,
		"stream":   true,
		"messages": []map[string]interface{}{
			{"role": "user", "content": contentParts},
		},
	}

	bodyBytes, err := json.Marshal(bodyMap)
	if err != nil {
		errCh := make(chan string, 1)
		close(errCh)
		return errCh, fmt.Errorf("deepseek multimodal marshal: %w", err)
	}

	lines, err := DoStreamRequest(ctx, url, headers, bodyBytes)
	if err != nil {
		return nil, err
	}

	ch := make(chan string, 32)
	go func() {
		defer close(ch)
		for line := range lines {
			if strings.HasPrefix(line, "__error__:") {
				return
			}
			if strings.HasPrefix(line, "data:") {
				dataStr := strings.TrimSpace(line[5:])
				if dataStr == "" || dataStr == "[DONE]" {
					continue
				}
				token := ExtractToken(dataStr, "choices.0.delta.content")
				if token == "" {
					token = ExtractToken(dataStr, "choices.0.message.content")
				}
				if token != "" {
					ch <- token
				}
			}
		}
	}()

	return ch, nil
}
