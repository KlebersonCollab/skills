package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadSaveConfig(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "harness-config-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	configPath := filepath.Join(tmpDir, "config.json")
	cfg := &AppConfig{
		ActiveProvider: "gemini",
		Providers: map[string]ProviderConfig{
			"gemini": {
				URL:          "http://localhost:8080/gemini",
				Headers:      map[string]string{"Authorization": "Bearer {{TEST_KEY}}"},
				BodyTemplate: `{"prompt": "{{prompt}}"}`,
				ResponsePath: "text",
			},
		},
	}

	err = SaveConfig(configPath, cfg)
	if err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	loaded, err := LoadConfig(configPath)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if loaded.ActiveProvider != "gemini" {
		t.Errorf("expected active provider 'gemini', got '%s'", loaded.ActiveProvider)
	}

	prov, ok := loaded.Providers["gemini"]
	if !ok {
		t.Fatalf("gemini provider not found in loaded config")
	}

	if prov.ResponsePath != "text" {
		t.Errorf("expected response path 'text', got '%s'", prov.ResponsePath)
	}
}

func TestSessionTreeDAG(t *testing.T) {
	tree := NewSessionTree("test-sess", "Test Objective")

	// Verify root-like insertion
	n1 := tree.AddNode("system", "Init instruction", 10)
	if tree.ActiveNode != n1 {
		t.Errorf("expected active node '%s', got '%s'", n1, tree.ActiveNode)
	}

	// Verify children insertion
	n2 := tree.AddNode("user", "Hello agent", 5)
	n3 := tree.AddNode("assistant", "Hello user", 5)

	history, err := tree.GetLinearHistory(n3)
	if err != nil {
		t.Fatalf("failed to get history: %v", err)
	}

	if len(history) != 3 {
		t.Errorf("expected history length 3, got %d", len(history))
	}

	if history[0].NodeID != n1 || history[1].NodeID != n2 || history[2].NodeID != n3 {
		t.Errorf("linear history out of order")
	}

	// Test formatting history
	formatted, err := tree.FormatLinearHistoryText(n3)
	if err != nil {
		t.Fatalf("failed to format history: %v", err)
	}

	if !strings.Contains(formatted, "[System]: Init instruction") {
		t.Errorf("formatted text missing system instruction")
	}
	if !strings.Contains(formatted, "[User]: Hello agent") {
		t.Errorf("formatted text missing user prompt")
	}
}

func TestResolveEnvVars(t *testing.T) {
	os.Setenv("HARNESS_TEST_VAR", "antigravity")
	defer os.Unsetenv("HARNESS_TEST_VAR")

	resolved := resolveEnvVars("https://api.com/{{HARNESS_TEST_VAR}}/run")
	expected := "https://api.com/antigravity/run"
	if resolved != expected {
		t.Errorf("expected '%s', got '%s'", expected, resolved)
	}
}

func TestJSONValueSelector(t *testing.T) {
	jsonStr := `{"candidates":[{"content":{"parts":[{"text":"hello world"}]}}],"code":200}`
	var data interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}

	val, err := getJSONValue(data, "candidates.0.content.parts.0.text")
	if err != nil {
		t.Fatalf("failed to select json value: %v", err)
	}

	if val != "hello world" {
		t.Errorf("expected 'hello world', got '%s'", val)
	}

	val2, err := getJSONValue(data, "code")
	if err != nil {
		t.Fatalf("failed to select second value: %v", err)
	}
	if val2 != "200" {
		t.Errorf("expected '200', got '%s'", val2)
	}
}

func TestExecuteCommand(t *testing.T) {
	// Set mock state to allow execute_command
	// We need to create a mock .specs/project/STATE.md inside the current workspace if it's not in IMPLEMENT
	// But our workspace is ALREADY in IMPLEMENT phase! So VerifySDDGated() will pass naturally.
	out, exitCode, err := ExecuteCommand("echo 'harness_run'")
	if err != nil {
		t.Fatalf("failed to execute command: %v", err)
	}

	if exitCode != 0 {
		t.Errorf("expected exit code 0, got %d", exitCode)
	}

	if !strings.Contains(out, "harness_run") {
		t.Errorf("expected output to contain 'harness_run', got '%s'", out)
	}
}
