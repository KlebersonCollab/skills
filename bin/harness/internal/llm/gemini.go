package llm

import (
	"context"
	"fmt"
	"strings"
)

// GeminiProvider implements LLMProvider for Google Generative AI (Gemini).
type GeminiProvider struct {
	apiKey string
	model  string
}

// NewGeminiProvider creates a new Gemini provider.
// If model is empty, "gemini-1.5-pro" is used.
func NewGeminiProvider(apiKey string, model string) *GeminiProvider {
	if model == "" {
		model = "gemini-1.5-pro"
	}
	return &GeminiProvider{
		apiKey: apiKey,
		model:  model,
	}
}

func (g *GeminiProvider) Name() string {
	return "gemini"
}

func (g *GeminiProvider) Complete(ctx context.Context, prompt string, opts map[string]any) (string, error) {
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent?key=%s", g.model, g.apiKey)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	body := fmt.Sprintf(`{"contents":[{"parts":[{"text":%s}]}]}`, JSONString(prompt))

	respBody, err := DoRequest(ctx, url, headers, []byte(body))
	if err != nil {
		return "", err
	}

	token := ExtractToken(string(respBody), "candidates.0.content.parts.0.text")
	if token == "" {
		return "", fmt.Errorf("gemini: empty response or unexpected format")
	}
	return token, nil
}

func (g *GeminiProvider) Stream(ctx context.Context, prompt string, opts map[string]any) (<-chan string, error) {
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:streamGenerateContent?key=%s", g.model, g.apiKey)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	body := fmt.Sprintf(`{"contents":[{"parts":[{"text":%s}]}]}`, JSONString(prompt))

	lines, err := DoStreamRequest(ctx, url, headers, []byte(body))
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
			// SSE format: data: {...}
			if strings.HasPrefix(line, "data:") {
				dataStr := strings.TrimSpace(line[5:])
				if dataStr == "" || dataStr == "[DONE]" {
					continue
				}
				token := ExtractToken(dataStr, "candidates.0.content.parts.0.text")
				if token != "" {
					ch <- token
				}
			}
		}
	}()

	return ch, nil
}
