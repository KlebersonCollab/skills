package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

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

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(bodyStr)))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	for k, v := range provider.Headers {
		resolvedVal := resolveEnvVars(v)
		req.Header.Set(k, resolvedVal)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
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

	extractedText, err := getJSONValue(jsonResponse, provider.ResponsePath)
	if err != nil {
		return "", fmt.Errorf("failed to extract text from path '%s': %w", provider.ResponsePath, err)
	}

	return extractedText, nil
}
