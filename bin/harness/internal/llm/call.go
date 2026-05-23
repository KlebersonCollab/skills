package llm

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// DefaultHTTPClient is the shared HTTP client used by all providers.
// Exposed as a variable to allow test overrides.
var DefaultHTTPClient = &http.Client{
	Timeout: 2 * time.Minute,
}

// DoRequest performs an HTTP POST with the given body, returning the raw response body.
func DoRequest(ctx context.Context, url string, headers map[string]string, body []byte) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := DefaultHTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	return respBody, nil
}

// DoStreamRequest performs an HTTP POST and returns a channel of raw text lines.
// The channel is closed when the stream ends or an error occurs.
func DoStreamRequest(ctx context.Context, url string, headers map[string]string, body []byte) (<-chan string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create stream request: %w", err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := DefaultHTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		respBody, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	ch := make(chan string, 32)
	go func() {
		defer resp.Body.Close()
		defer close(ch)

		reader := bufio.NewReader(resp.Body)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					ch <- fmt.Sprintf("__error__:%s", err.Error())
				}
				return
			}
			line = strings.TrimSpace(line)
			if line != "" {
				ch <- line
			}
		}
	}()

	return ch, nil
}

// ExtractToken parses a JSON line and extracts the string value at the given dotted path.
// Returns empty string if path is not found.
func ExtractToken(dataStr string, path string) string {
	var data interface{}
	if err := json.Unmarshal([]byte(dataStr), &data); err != nil {
		return ""
	}
	token, err := getJSONValue(data, path)
	if err != nil {
		return ""
	}
	return token
}

// getJSONValue traverses a parsed JSON tree following a dotted path.
func getJSONValue(data interface{}, path string) (string, error) {
	parts := strings.Split(path, ".")
	var current interface{} = data
	for _, part := range parts {
		if part == "" {
			continue
		}
		switch v := current.(type) {
		case map[string]interface{}:
			var ok bool
			current, ok = v[part]
			if !ok {
				return "", fmt.Errorf("key '%s' not found", part)
			}
		case []interface{}:
			idx := 0
			if _, err := fmt.Sscanf(part, "%d", &idx); err != nil {
				return "", fmt.Errorf("expected array index for '%s'", part)
			}
			if idx < 0 || idx >= len(v) {
				return "", fmt.Errorf("array index %d out of bounds", idx)
			}
			current = v[idx]
		default:
			return "", fmt.Errorf("cannot traverse key '%s' on leaf value", part)
		}
	}
	switch val := current.(type) {
	case string:
		return val, nil
	default:
		return fmt.Sprintf("%v", val), nil
	}
}

// RetryConfig defines the retry behavior for LLM calls.
type RetryConfig struct {
	MaxAttempts int
	BaseDelay   time.Duration
}

// DefaultRetry is the default retry configuration.
var DefaultRetry = RetryConfig{
	MaxAttempts: 3,
	BaseDelay:   500 * time.Millisecond,
}

// RetryCall executes fn up to MaxAttempts times, retrying on errors.
func RetryCall(ctx context.Context, cfg RetryConfig, fn func(ctx context.Context) (string, error)) (string, error) {
	var lastErr error

	for attempt := 1; attempt <= cfg.MaxAttempts; attempt++ {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}

		result, err := fn(ctx)
		if err == nil {
			return result, nil
		}

		lastErr = err

		if attempt < cfg.MaxAttempts {
			time.Sleep(cfg.BaseDelay)
		}
	}

	return "", fmt.Errorf("request failed after %d attempts: %w", cfg.MaxAttempts, lastErr)
}

// ResolveTemplate replaces {{prompt}} and {{history}} placeholders in the body template.
func ResolveTemplate(tmpl string, prompt string, history string) string {
	s := strings.ReplaceAll(tmpl, "{{prompt}}", escapeJSONString(prompt))
	s = strings.ReplaceAll(s, "{{history}}", escapeJSONString(history))
	return s
}

// escapeJSONString wraps a string in JSON-safe encoding (no surrounding quotes).
func escapeJSONString(s string) string {
	b, _ := json.Marshal(s)
	if len(b) >= 2 {
		return string(b[1 : len(b)-1])
	}
	return s
}

// JSONString returns a properly JSON-escaped quoted string using json.Marshal.
// Use this instead of fmt.Sprintf("%q", s) which can produce invalid JSON.
func JSONString(s string) string {
	b, _ := json.Marshal(s)
	return string(b)
}
