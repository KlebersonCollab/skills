package llm

import (
	"context"
	"fmt"
	"strings"
)

// OpenAIProvider implements LLMProvider for OpenAI (GPT) API.
type OpenAIProvider struct {
	apiKey string
	model  string
}

// NewOpenAIProvider creates a new OpenAI provider.
// If model is empty, "gpt-4o" is used.
func NewOpenAIProvider(apiKey string, model string) *OpenAIProvider {
	if model == "" {
		model = "gpt-4o"
	}
	return &OpenAIProvider{
		apiKey: apiKey,
		model:  model,
	}
}

func (o *OpenAIProvider) Name() string { return "openai" }

func (o *OpenAIProvider) Complete(ctx context.Context, prompt string, opts map[string]any) (string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + o.apiKey,
	}

	body := fmt.Sprintf(`{"model":%s,"messages":[{"role":"user","content":%s}]}`, JSONString(o.model), JSONString(prompt))

	respBody, err := DoRequest(ctx, url, headers, []byte(body))
	if err != nil {
		return "", err
	}

	token := ExtractToken(string(respBody), "choices.0.message.content")
	if token == "" {
		return "", fmt.Errorf("openai: empty response or unexpected format")
	}
	return token, nil
}

func (o *OpenAIProvider) Stream(ctx context.Context, prompt string, opts map[string]any) (<-chan string, error) {
	url := "https://api.openai.com/v1/chat/completions"

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + o.apiKey,
	}

	body := fmt.Sprintf(`{"model":%s,"messages":[{"role":"user","content":%s}],"stream":true}`, JSONString(o.model), JSONString(prompt))

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
