package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var llmRequestTimeout = 2 * time.Minute

type ProviderConfig struct {
	URL          string            `json:"url"`
	Headers      map[string]string `json:"headers"`
	BodyTemplate string            `json:"body_template"`
	ResponsePath string            `json:"response_path"`
}

type AppConfig struct {
	ActiveProvider string                    `json:"active_provider"`
	Providers      map[string]ProviderConfig `json:"providers"`
}

func LoadConfig(path string) (*AppConfig, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config AppConfig
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

func SaveConfig(path string, config *AppConfig) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(config)
}

func resolveEnvVars(str string) string {
	// Replaces all occurrences of {{VAR_NAME}} with os.Getenv("VAR_NAME")
	resolved := str
	for {
		start := strings.Index(resolved, "{{")
		if start == -1 {
			break
		}
		end := strings.Index(resolved[start:], "}}")
		if end == -1 {
			break
		}
		end = start + end
		varName := resolved[start+2 : end]
		resolved = resolved[:start] + os.Getenv(varName) + resolved[end+2:]
	}
	return resolved
}

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
			idx, err := strconv.Atoi(part)
			if err != nil {
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

func CallLLM(config *AppConfig, prompt string, history string) (string, error) {
	providerName := config.ActiveProvider
	provider, exists := config.Providers[providerName]
	if !exists {
		return "", fmt.Errorf("provider '%s' not found in configuration", providerName)
	}

	url := resolveEnvVars(provider.URL)
	escapedPromptBytes, _ := json.Marshal(prompt)
	escapedPrompt := string(escapedPromptBytes[1 : len(escapedPromptBytes)-1])
	escapedHistoryBytes, _ := json.Marshal(history)
	escapedHistory := string(escapedHistoryBytes[1 : len(escapedHistoryBytes)-1])

	bodyStr := strings.ReplaceAll(provider.BodyTemplate, "{{prompt}}", escapedPrompt)
	bodyStr = strings.ReplaceAll(bodyStr, "{{history}}", escapedHistory)

	maxAttempts := 3
	var lastErr error
	var extractedText string

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		var err error
		extractedText, err = func() (string, error) {
			ctx, cancel := context.WithTimeout(context.Background(), llmRequestTimeout)
			defer cancel()

			req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(bodyStr)))
			if err != nil {
				return "", fmt.Errorf("failed to create request: %w", err)
			}

			for k, v := range provider.Headers {
				resolvedVal := resolveEnvVars(v)
				req.Header.Set(k, resolvedVal)
			}

			req = req.WithContext(ctx)
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return "", err
			}
			defer resp.Body.Close()

			respBody, err := io.ReadAll(resp.Body)
			if err != nil {
				return "", fmt.Errorf("failed to read response body: %w", err)
			}

			if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				return "", fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(respBody))
			}

			var jsonResponse interface{}
			if err := json.Unmarshal(respBody, &jsonResponse); err != nil {
				return "", fmt.Errorf("failed to parse JSON response: %w", err)
			}

			text, err := getJSONValue(jsonResponse, provider.ResponsePath)
			if err != nil {
				return "", fmt.Errorf("failed to extract text from path '%s': %w", provider.ResponsePath, err)
			}

			return text, nil
		}()

		if err == nil {
			return extractedText, nil
		}

		lastErr = err
		// Detect if the error was due to timeout (either context timeout or net timeout)
		isTimeout := false
		if err == context.DeadlineExceeded {
			isTimeout = true
		} else if netErr, ok := err.(interface{ Timeout() bool }); ok && netErr.Timeout() {
			isTimeout = true
		} else if strings.Contains(err.Error(), "context deadline exceeded") || strings.Contains(err.Error(), "Client.Timeout") {
			isTimeout = true
		}

		if attempt < maxAttempts {
			// Clear the current spinner output dynamically before logging the warning
			fmt.Print("\r\033[K")
			if isTimeout {
				fmt.Printf("\033[38;5;208m⚠️  Timeout de 2 minutos atingido para %s. Iniciando nova tentativa (%d/%d)...\033[0m\n", strings.ToUpper(providerName), attempt+1, maxAttempts)
			} else {
				fmt.Printf("\033[38;5;196m⚠️  Erro na chamada para %s: %v. Iniciando nova tentativa (%d/%d)...\033[0m\n", strings.ToUpper(providerName), err, attempt+1, maxAttempts)
			}
			// Small delay before retrying to prevent connection flooding
			time.Sleep(500 * time.Millisecond)
		}
	}

	return "", fmt.Errorf("HTTP request failed after %d attempts: %w", maxAttempts, lastErr)
}
