package llm

import (
	"context"
	"fmt"
	"strings"
)

// OpenAICompatibleProvider implements LLMProvider for any OpenAI-compatible API.
// Covers: OpenRouter, Together AI, Fireworks, Mistral, xAI (Grok),
// Cerebras, Cloudflare Workers AI, Hugging Face, Azure OpenAI, and more.
type OpenAICompatibleProvider struct {
	name   string // provider identifier (e.g., "openrouter", "mistral")
	apiKey string
	model  string
	url    string // API endpoint URL
}

// NewOpenAICompatibleProvider creates a provider for any OpenAI-compatible API.
//
// Parameters:
//   - name: unique identifier (e.g., "openrouter", "mistral", "together")
//   - apiKey: API key for the service
//   - url: full API endpoint URL (e.g., "https://api.openrouter.ai/v1/chat/completions")
//   - model: model name (e.g., "mistral-large", "grok-2")
func NewOpenAICompatibleProvider(name, apiKey, url, model string) *OpenAICompatibleProvider {
	if model == "" {
		model = "default"
	}
	if url == "" {
		url = "https://api.openai.com/v1/chat/completions"
	}
	return &OpenAICompatibleProvider{
		name:   name,
		apiKey: apiKey,
		model:  model,
		url:    url,
	}
}

func (o *OpenAICompatibleProvider) Name() string { return o.name }

func (o *OpenAICompatibleProvider) Complete(ctx context.Context, prompt string, opts map[string]any) (string, error) {
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + o.apiKey,
	}

	body := fmt.Sprintf(`{"model":%s,"messages":[{"role":"user","content":%s}]}`, JSONString(o.model), JSONString(prompt))

	respBody, err := DoRequest(ctx, o.url, headers, []byte(body))
	if err != nil {
		return "", err
	}

	token := ExtractToken(string(respBody), "choices.0.message.content")
	if token == "" {
		return "", fmt.Errorf("%s: empty response or unexpected format", o.name)
	}
	return token, nil
}

func (o *OpenAICompatibleProvider) Stream(ctx context.Context, prompt string, opts map[string]any) (<-chan string, error) {
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + o.apiKey,
	}

	body := fmt.Sprintf(`{"model":%s,"messages":[{"role":"user","content":%s}],"stream":true}`, JSONString(o.model), JSONString(prompt))

	lines, err := DoStreamRequest(ctx, o.url, headers, []byte(body))
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
