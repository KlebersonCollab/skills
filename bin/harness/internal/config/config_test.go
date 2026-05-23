package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResolveEnvVars(t *testing.T) {
	os.Setenv("TEST_CONFIG_VAR", "magic-value")
	defer os.Unsetenv("TEST_CONFIG_VAR")

	result := ResolveEnvVars("https://api.com/{{TEST_CONFIG_VAR}}/endpoint")
	expected := "https://api.com/magic-value/endpoint"
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestResolveEnvVars_NoPlaceholder(t *testing.T) {
	result := ResolveEnvVars("https://api.com/static")
	if result != "https://api.com/static" {
		t.Errorf("expected unchanged string, got %q", result)
	}
}

func TestResolveEnvVars_UnsetVar(t *testing.T) {
	result := ResolveEnvVars("https://api.com/{{UNSET_VAR_XYZ}}/endpoint")
	if result != "https://api.com//endpoint" {
		t.Errorf("expected empty replacement, got %q", result)
	}
}

func TestGetJSONValue_Basic(t *testing.T) {
	data := map[string]interface{}{
		"name": "test",
		"count": 42,
	}

	val, err := GetJSONValue(data, "name")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "test" {
		t.Errorf("expected 'test', got %q", val)
	}
}

func TestGetJSONValue_Nested(t *testing.T) {
	data := map[string]interface{}{
		"candidates": []interface{}{
			map[string]interface{}{
				"content": map[string]interface{}{
					"parts": []interface{}{
						map[string]interface{}{
							"text": "hello world",
						},
					},
				},
			},
		},
	}

	val, err := GetJSONValue(data, "candidates.0.content.parts.0.text")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "hello world" {
		t.Errorf("expected 'hello world', got %q", val)
	}
}

func TestGetJSONValue_IntToString(t *testing.T) {
	data := map[string]interface{}{
		"code": 200,
	}

	val, err := GetJSONValue(data, "code")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "200" {
		t.Errorf("expected '200', got %q", val)
	}
}

func TestGetJSONValue_NotFound(t *testing.T) {
	data := map[string]interface{}{"a": 1}
	_, err := GetJSONValue(data, "b")
	if err == nil {
		t.Error("expected error for missing key")
	}
}

func TestGetJSONValue_OutOfBounds(t *testing.T) {
	data := map[string]interface{}{
		"items": []interface{}{1, 2},
	}
	_, err := GetJSONValue(data, "items.5")
	if err == nil {
		t.Error("expected error for out of bounds index")
	}
}

func TestLoadEnvFile(t *testing.T) {
	tmpDir := t.TempDir()
	envPath := filepath.Join(tmpDir, ".env")
	content := []byte("# Comentario\nKEY1=value1\nKEY2=value2 with spaces\nEMPTY=\n")
	if err := os.WriteFile(envPath, content, 0644); err != nil {
		t.Fatalf("failed to write .env: %v", err)
	}

	os.Unsetenv("KEY1")
	os.Unsetenv("KEY2")
	os.Unsetenv("EMPTY")

	if err := LoadEnvFile(tmpDir); err != nil {
		t.Fatalf("LoadEnvFile failed: %v", err)
	}

	if os.Getenv("KEY1") != "value1" {
		t.Errorf("expected KEY1=value1, got %q", os.Getenv("KEY1"))
	}
	if os.Getenv("KEY2") != "value2 with spaces" {
		t.Errorf("expected KEY2='value2 with spaces', got %q", os.Getenv("KEY2"))
	}
}

func TestLoadEnvFile_NotExists(t *testing.T) {
	tmpDir := t.TempDir()
	if err := LoadEnvFile(tmpDir); err != nil {
		t.Errorf("expected nil for missing .env, got %v", err)
	}
}

func TestLoadEnvFile_DoesNotOverride(t *testing.T) {
	tmpDir := t.TempDir()
	envPath := filepath.Join(tmpDir, ".env")
	os.WriteFile(envPath, []byte("EXISTING=from_file"), 0644)

	os.Setenv("EXISTING", "from_env")
	defer os.Unsetenv("EXISTING")

	LoadEnvFile(tmpDir)

	if os.Getenv("EXISTING") != "from_env" {
		t.Errorf("expected env var to keep original value, got %q", os.Getenv("EXISTING"))
	}
}

func TestSaveAndLoadConfig(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "config.json")

	original := &AppConfig{
		ActiveProvider: "test",
		Providers: map[string]ProviderConfig{
			"test": {
				URL:          "http://localhost:9999",
				ResponsePath: "result",
				Headers:      map[string]string{"X-Key": "val"},
			},
		},
	}

	if err := SaveConfig(path, original); err != nil {
		t.Fatalf("SaveConfig failed: %v", err)
	}

	loaded, err := LoadConfig(path)
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	if loaded.ActiveProvider != "test" {
		t.Errorf("expected ActiveProvider 'test', got %q", loaded.ActiveProvider)
	}
	if _, ok := loaded.Providers["test"]; !ok {
		t.Error("expected 'test' provider in loaded config")
	}
}

func TestGetSDDPhase(t *testing.T) {
	root := filepath.Join(t.TempDir(), ".specs", "project")
	os.MkdirAll(root, 0755)
	statePath := filepath.Join(root, "STATE.md")
	os.WriteFile(statePath, []byte("---\n```yaml\nphase: \"IMPLEMENT\"\n```\n---\n"), 0644)

	phase, err := GetSDDPhase(filepath.Dir(filepath.Dir(root)))
	if err != nil {
		t.Fatalf("GetSDDPhase failed: %v", err)
	}
	if phase != "IMPLEMENT" {
		t.Errorf("expected IMPLEMENT, got %q", phase)
	}
}

func TestGetSDDPhase_NotFound(t *testing.T) {
	tmpDir := t.TempDir()
	_, err := GetSDDPhase(tmpDir)
	if err == nil {
		t.Error("expected error for missing .specs/project/STATE.md")
	}
}

func TestVerifySDDGated_Allowed(t *testing.T) {
	root := filepath.Join(t.TempDir(), ".specs", "project")
	os.MkdirAll(root, 0755)
	statePath := filepath.Join(root, "STATE.md")
	os.WriteFile(statePath, []byte("---\n```yaml\nphase: \"IMPLEMENT\"\n```\n---\n"), 0644)

	wsRoot := filepath.Dir(filepath.Dir(root))
	if err := VerifySDDGated(wsRoot); err != nil {
		t.Errorf("expected nil for IMPLEMENT phase, got %v", err)
	}
}

func TestVerifySDDGated_Blocked(t *testing.T) {
	root := filepath.Join(t.TempDir(), ".specs", "project")
	os.MkdirAll(root, 0755)
	statePath := filepath.Join(root, "STATE.md")
	os.WriteFile(statePath, []byte("---\n```yaml\nphase: \"DISCOVERY\"\n```\n---\n"), 0644)

	wsRoot := filepath.Dir(filepath.Dir(root))
	if err := VerifySDDGated(wsRoot); err == nil {
		t.Error("expected error for DISCOVERY phase")
	}
}
