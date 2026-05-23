package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// AnthropicProvider implements LLMProvider for Anthropic (Claude) API.
// Uses Anthropic Messages API format.
type AnthropicProvider struct {
	apiKey string
	model  string
}

// NewAnthropicProvider creates a new Anthropic provider.
// If model is empty, "claude-sonnet-4-20250514" is used.
func NewAnthropicProvider(apiKey string, model string) *AnthropicProvider {
	if model == "" {
		model = "claude-sonnet-4-20250514"
	}
	return &AnthropicProvider{
		apiKey: apiKey,
		model:  model,
	}
}

func (a *AnthropicProvider) Name() string { return "anthropic" }

func (a *AnthropicProvider) Complete(ctx context.Context, prompt string, opts map[string]any) (string, error) {
	url := "https://api.anthropic.com/v1/messages"

	headers := map[string]string{
		"Content-Type":      "application/json",
		"x-api-key":         a.apiKey,
		"anthropic-version": "2023-06-01",
	}

	bodyMap := map[string]interface{}{
		"model":    a.model,
		"max_tokens": 4096,
		"messages": []map[string]interface{}{
			{"role": "user", "content": prompt},
		},
	}

	bodyBytes, _ := json.Marshal(bodyMap)

	respBody, err := DoRequest(ctx, url, headers, bodyBytes)
	if err != nil {
		return "", err
	}

	// Anthropic returns: content[0].text
	token := ExtractToken(string(respBody), "content.0.text")
	if token == "" {
		return "", fmt.Errorf("anthropic: empty response or unexpected format")
	}
	return token, nil
}

func (a *AnthropicProvider) Stream(ctx context.Context, prompt string, opts map[string]any) (<-chan string, error) {
	url := "https://api.anthropic.com/v1/messages"

	headers := map[string]string{
		"Content-Type":      "application/json",
		"x-api-key":         a.apiKey,
		"anthropic-version": "2023-06-01",
	}

	bodyMap := map[string]interface{}{
		"model":       a.model,
		"max_tokens":  4096,
		"stream":      true,
		"messages": []map[string]interface{}{
			{"role": "user", "content": prompt},
		},
	}

	bodyBytes, _ := json.Marshal(bodyMap)

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
			// Anthropic SSE format: event: content_block_delta\n data: {...}
			// Each data line has "type": "content_block_delta" with delta.text
			if strings.HasPrefix(line, "data:") {
				dataStr := strings.TrimSpace(line[5:])
				if dataStr == "" || dataStr == "[DONE]" {
					continue
				}
				// Extract delta text
				token := ExtractToken(dataStr, "delta.text")
				if token != "" {
					ch <- token
				}
			}
		}
	}()

	return ch, nil
}
