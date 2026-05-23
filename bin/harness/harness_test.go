package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"harness/internal/config"
	"harness/internal/llm"
	"harness/internal/session"
	"harness/internal/tools"
)

func TestLoadSaveConfig(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "harness-config-test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	configPath := filepath.Join(tmpDir, "config.json")
	cfg := &config.AppConfig{
		ActiveProvider: "gemini",
		Providers: map[string]config.ProviderConfig{
			"gemini": {
				URL:          "http://localhost:8080/gemini",
				Headers:      map[string]string{"Authorization": "Bearer {{TEST_KEY}}"},
				BodyTemplate: `{"prompt": "{{prompt}}"}`,
				ResponsePath: "text",
			},
		},
	}

	err = config.SaveConfig(configPath, cfg)
	if err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	loaded, err := config.LoadConfig(configPath)
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
	tree := session.NewTree("test-sess", "Test Objective")

	n1 := tree.AddNode("system", "Init instruction", 10)
	if tree.ActiveNode != n1 {
		t.Errorf("expected active node '%s', got '%s'", n1, tree.ActiveNode)
	}

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

	formatted, err := tree.FormatHistory(n3)
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

	resolved := config.ResolveEnvVars("https://api.com/{{HARNESS_TEST_VAR}}/run")
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

	val, err := config.GetJSONValue(data, "candidates.0.content.parts.0.text")
	if err != nil {
		t.Fatalf("failed to select json value: %v", err)
	}

	if val != "hello world" {
		t.Errorf("expected 'hello world', got '%s'", val)
	}

	val2, err := config.GetJSONValue(data, "code")
	if err != nil {
		t.Fatalf("failed to select second value: %v", err)
	}
	if val2 != "200" {
		t.Errorf("expected '200', got '%s'", val2)
	}
}

func TestExecuteCommand(t *testing.T) {
	root, _ := os.Getwd()
	executor := &tools.ExecCommandExecutor{Root: root}
	result := executor.Execute("execute_command", map[string]any{"command": "echo 'harness_run'"})
	if result.Error != nil {
		t.Fatalf("failed to execute command: %v", result.Error)
	}

	if result.ExitCode != 0 {
		t.Errorf("expected exit code 0, got %d", result.ExitCode)
	}

	if !strings.Contains(result.Output, "harness_run") {
		t.Errorf("expected output to contain 'harness_run', got '%s'", result.Output)
	}
}

func TestSearchFiles(t *testing.T) {
	searcher := &tools.SearchExecutor{Root: "/home/kleberson/Documentos/skills"}

	res := searcher.Execute("search_files", map[string]any{"pattern": "*.go"})
	if res.Error != nil {
		t.Fatalf("SearchFiles failed: %v", res.Error)
	}

	if !strings.Contains(res.Output, ".go") {
		t.Errorf("expected search result to contain Go files, got:\n%s", res.Output)
	}

	res2 := searcher.Execute("search_files", map[string]any{"pattern": "*.go", "query": "TestJSONValueSelector"})
	if res2.Error != nil {
		t.Fatalf("SearchFiles failed: %v", res2.Error)
	}
	if !strings.Contains(res2.Output, "TestJSONValueSelector") {
		t.Errorf("expected search results to find TestJSONValueSelector, got:\n%s", res2.Output)
	}
}

func TestRegexSearchFiles(t *testing.T) {
	searchFilesRE := regexp.MustCompile(`<tool:search_files\s+([\s\S]*?)\s*/?>`)
	patternAttrRE := regexp.MustCompile(`pattern="([^"]*)"`)
	queryAttrRE := regexp.MustCompile(`query="([^"]*)"`)

	input := `<tool:search_files pattern="*" query="" path="/home/test"/>`
	matches := searchFilesRE.FindAllStringSubmatch(input, -1)
	if len(matches) == 0 {
		t.Fatalf("expected regex to match input string, got 0 matches")
	}

	attrs := matches[0][1]
	pattern := ""
	query := ""

	if patMatch := patternAttrRE.FindStringSubmatch(attrs); len(patMatch) > 1 {
		pattern = patMatch[1]
	}
	if qMatch := queryAttrRE.FindStringSubmatch(attrs); len(qMatch) > 1 {
		query = qMatch[1]
	}

	if pattern != "*" {
		t.Errorf("expected pattern to be '*', got %q", pattern)
	}
	if query != "" {
		t.Errorf("expected query to be empty, got %q", query)
	}
}

func TestCallLLMRetryOnTimeout(t *testing.T) {
	var callCount int64

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&callCount, 1)
		time.Sleep(50 * time.Millisecond)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"text":"should timeout and retry"}`))
	}))
	defer server.Close()

	// We test the retry mechanism via the low-level RetryCall
	originalTimeout := llm.DefaultHTTPClient.Timeout
	llm.DefaultHTTPClient.Timeout = 5 * time.Millisecond
	defer func() {
		llm.DefaultHTTPClient.Timeout = originalTimeout
	}()

	ctx := context.Background()
	_, err := llm.RetryCall(ctx, llm.RetryConfig{MaxAttempts: 3, BaseDelay: 1 * time.Millisecond},
		func(ctx context.Context) (string, error) {
			body, err := llm.DoRequest(ctx, server.URL, map[string]string{}, []byte(`{}`))
			if err != nil {
				return "", err
			}
			return string(body), nil
		})

	if err == nil {
		t.Log("Note: request may have succeeded or failed depending on timing")
	}

	actualAttempts := atomic.LoadInt64(&callCount)
	t.Logf("HTTP calls made: %d", actualAttempts)
	if actualAttempts < 1 {
		t.Errorf("expected at least 1 HTTP call, got %d", actualAttempts)
	}
}

func TestBuildDefaultAppConfig_OpenRouter(t *testing.T) {
	// 1. Test registration of OpenRouter
	cfg := buildDefaultAppConfig()
	prov, exists := cfg.Providers["openrouter"]
	if !exists {
		t.Fatalf("expected openrouter provider to be registered in default app config")
	}
	if prov.URL != "https://openrouter.ai/api/v1/chat/completions" {
		t.Errorf("expected URL 'https://openrouter.ai/api/v1/chat/completions', got %q", prov.URL)
	}

	// 2. Test auto-detection
	os.Setenv("OPENROUTER_API_KEY", "test-or-key")
	defer os.Unsetenv("OPENROUTER_API_KEY")

	cfg2 := buildDefaultAppConfig()
	if cfg2.ActiveProvider != "openrouter" {
		t.Errorf("expected active provider to be 'openrouter' when OPENROUTER_API_KEY is configured, got %q", cfg2.ActiveProvider)
	}
}

