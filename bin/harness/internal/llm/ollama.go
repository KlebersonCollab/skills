package llm

import (
	"context"
	"fmt"
	"strings"
)

// OllamaProvider implements LLMProvider for local Ollama models.
type OllamaProvider struct {
	host  string
	model string
}

// NewOllamaProvider creates a new Ollama provider.
// If host is empty, "http://localhost:11434" is used.
// If model is empty, "llama3" is used.
func NewOllamaProvider(host string, model string) *OllamaProvider {
	if host == "" {
		host = "http://localhost:11434"
	}
	if model == "" {
		model = "llama3"
	}
	return &OllamaProvider{
		host:  host,
		model: model,
	}
}

func (o *OllamaProvider) Name() string {
	return "ollama"
}

func (o *OllamaProvider) Complete(ctx context.Context, prompt string, opts map[string]any) (string, error) {
	url := o.host + "/api/generate"

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	body := fmt.Sprintf(`{"model":%s,"prompt":%s,"stream":false}`, JSONString(o.model), JSONString(prompt))

	respBody, err := DoRequest(ctx, url, headers, []byte(body))
	if err != nil {
		return "", err
	}

	token := ExtractToken(string(respBody), "response")
	if token == "" {
		return "", fmt.Errorf("ollama: empty response or unexpected format")
	}
	return token, nil
}

func (o *OllamaProvider) Stream(ctx context.Context, prompt string, opts map[string]any) (<-chan string, error) {
	url := o.host + "/api/generate"

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	body := fmt.Sprintf(`{"model":%s,"prompt":%s,"stream":true}`, JSONString(o.model), JSONString(prompt))

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
			// Ollama NDJSON format: {"response":"...","done":false}
			token := ExtractToken(line, "response")
			if token != "" {
				ch <- token
			}
		}
	}()

	return ch, nil
}
