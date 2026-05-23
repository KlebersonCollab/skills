// Package internal/config provides configuration types and persistence for LLM providers and app settings.
package config

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ModelCost holds token pricing for a model.
type ModelCost struct {
	Input     float64 `json:"input"`
	Output    float64 `json:"output"`
	CacheRead float64 `json:"cacheRead,omitempty"`
}

// ModelConfig holds configuration for a single model within a provider.
type ModelConfig struct {
	ID            string    `json:"id"`
	Name          string    `json:"name,omitempty"`
	ContextWindow int       `json:"contextWindow,omitempty"`
	MaxTokens     int       `json:"maxTokens,omitempty"`
	InputTypes    []string  `json:"input,omitempty"` // "text", "image", "audio"
	Reasoning     bool      `json:"reasoning,omitempty"`
	Cost          ModelCost `json:"cost"`
}

// ProviderConfig holds the configuration for a single LLM provider.
type ProviderConfig struct {
	URL                string            `json:"url"`
	Headers            map[string]string `json:"headers"`
	BodyTemplate       string            `json:"body_template"`
	ResponsePath       string            `json:"response_path"`
	Streaming          bool              `json:"streaming"`
	StreamResponsePath string            `json:"stream_response_path"`
	Models             []ModelConfig     `json:"models,omitempty"`
}

// MCPServerConfig holds the configuration for an MCP server connection.
type MCPServerConfig struct {
	Name      string `json:"name"`
	Transport string `json:"transport"`
	Address   string `json:"address"`
}

// AppConfig is the root configuration for the Harness application.
type AppConfig struct {
	ActiveProvider string                    `json:"active_provider"`
	Providers      map[string]ProviderConfig `json:"providers"`
	MCPServers     []MCPServerConfig         `json:"mcp_servers,omitempty"`
}

// LoadConfig reads and deserializes the config file at the given path.
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

// SaveConfig serializes and writes the config to the given path.
func SaveConfig(path string, config *AppConfig) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(config)
}

// ResolveEnvVars replaces all {{VAR_NAME}} placeholders with os.Getenv("VAR_NAME").
func ResolveEnvVars(str string) string {
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

// GetJSONValue traverses a parsed JSON tree following a dotted path (e.g., "candidates.0.content.parts.0.text").
func GetJSONValue(data interface{}, path string) (string, error) {
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
				return "", fmt.Errorf("array index %d out of bounds (len %d)", idx, len(v))
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

// FindWorkspaceRoot searches upwards from the current directory for the .specs directory.
func FindWorkspaceRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		specsPath := filepath.Join(dir, ".specs")
		if info, err := os.Stat(specsPath); err == nil && info.IsDir() {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	return "", fmt.Errorf("workspace root (.specs directory) not found from %s", dir)
}

// LoadEnvFile loads environment variables from a .env file in the given directory.
// Format: KEY=VALUE (uma por linha, # para comentários, espaços ignorados).
// Nao sobrescreve variaveis ja definidas no ambiente.
func LoadEnvFile(dir string) error {
	envPath := filepath.Join(dir, ".env")
	f, err := os.Open(envPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // .env e opcional
		}
		return err
	}
	defer f.Close()

	loaded := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		value = strings.Trim(value, `"'`)

		if os.Getenv(key) == "" {
			os.Setenv(key, value)
			loaded++
		}
	}

	if loaded > 0 {
		fmt.Printf("\033[38;5;99m│\033[0m  📄 .env carregado (%d variaveis)\n", loaded)
	}

	return scanner.Err()
}
