package llm

import (
	"context"
	"fmt"
	"strings"
)

// GroqProvider implements LLMProvider for Groq (fast inference) API.
// Uses OpenAI-compatible chat completions format.
type GroqProvider struct {
	apiKey string
	model  string
}

// NewGroqProvider creates a new Groq provider.
// If model is empty, "llama-3.3-70b-versatile" is used.
func NewGroqProvider(apiKey string, model string) *GroqProvider {
	if model == "" {
		model = "llama-3.3-70b-versatile"
	}
	return &GroqProvider{
		apiKey: apiKey,
		model:  model,
	}
}

func (g *GroqProvider) Name() string { return "groq" }

func (g *GroqProvider) Complete(ctx context.Context, prompt string, opts map[string]any) (string, error) {
	url := "https://api.groq.com/openai/v1/chat/completions"

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + g.apiKey,
	}

	body := fmt.Sprintf(`{"model":%s,"messages":[{"role":"user","content":%s}]}`, JSONString(g.model), JSONString(prompt))

	respBody, err := DoRequest(ctx, url, headers, []byte(body))
	if err != nil {
		return "", err
	}

	token := ExtractToken(string(respBody), "choices.0.message.content")
	if token == "" {
		return "", fmt.Errorf("groq: empty response or unexpected format")
	}
	return token, nil
}

func (g *GroqProvider) Stream(ctx context.Context, prompt string, opts map[string]any) (<-chan string, error) {
	url := "https://api.groq.com/openai/v1/chat/completions"

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + g.apiKey,
	}

	body := fmt.Sprintf(`{"model":%s,"messages":[{"role":"user","content":%s}],"stream":true}`, JSONString(g.model), JSONString(prompt))

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
