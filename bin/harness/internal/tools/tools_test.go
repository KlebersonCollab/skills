package tools

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRegistry_RegisterAndExecute(t *testing.T) {
	r := NewRegistry()
	r.Register("echo", &mockTool{output: "hello"})

	result := r.Execute("echo", nil)
	if result.Error != nil {
		t.Fatalf("unexpected error: %v", result.Error)
	}
	if result.Output != "hello" {
		t.Errorf("expected 'hello', got %q", result.Output)
	}
}

func TestRegistry_NotFound(t *testing.T) {
	r := NewRegistry()
	result := r.Execute("nonexistent", nil)
	if result.Error == nil {
		t.Error("expected error for unknown tool")
	}
}

func TestRegistry_DuplicatePanics(t *testing.T) {
	r := NewRegistry()
	r.Register("dup", &mockTool{})

	defer func() {
		if rec := recover(); rec == nil {
			t.Error("expected panic for duplicate registration")
		}
	}()
	r.Register("dup", &mockTool{})
}

func TestRegistry_Names(t *testing.T) {
	r := NewRegistry()
	r.Register("a", &mockTool{})
	r.Register("b", &mockTool{})

	names := r.Names()
	if len(names) != 2 {
		t.Errorf("expected 2 names, got %d", len(names))
	}
}

func TestReadFileExecutor(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")
	os.WriteFile(filePath, []byte("hello world"), 0644)

	e := &ReadFileExecutor{Root: tmpDir}
	result := e.Execute("read_file", map[string]any{"path": "test.txt"})
	if result.Error != nil {
		t.Fatalf("unexpected error: %v", result.Error)
	}
	if !strings.Contains(result.Output, "hello world") {
		t.Errorf("expected 'hello world', got %q", result.Output)
	}
}

func TestReadFileExecutor_MissingPath(t *testing.T) {
	e := &ReadFileExecutor{}
	result := e.Execute("read_file", nil)
	if result.Error == nil {
		t.Error("expected error for missing path")
	}
}

func TestWriteFileExecutor(t *testing.T) {
	tmpDir := t.TempDir()
	e := &WriteFileExecutor{Root: tmpDir}

	result := e.Execute("write_file", map[string]any{
		"path":    "newfile.txt",
		"content": "test content",
	})
	if result.Error != nil {
		t.Fatalf("unexpected error: %v", result.Error)
	}

	data, _ := os.ReadFile(filepath.Join(tmpDir, "newfile.txt"))
	if string(data) != "test content" {
		t.Errorf("expected 'test content', got %q", string(data))
	}
}

func TestPatchFileExecutor(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "patch.txt")
	os.WriteFile(filePath, []byte("hello world foo"), 0644)

	e := &PatchFileExecutor{Root: tmpDir}
	result := e.Execute("patch_file", map[string]any{
		"path":        "patch.txt",
		"target":      "world",
		"replacement": "there",
	})
	if result.Error != nil {
		t.Fatalf("unexpected error: %v", result.Error)
	}

	data, _ := os.ReadFile(filePath)
	if string(data) != "hello there foo" {
		t.Errorf("expected 'hello there foo', got %q", string(data))
	}
}

func TestPatchFileExecutor_NotFound(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "nope.txt")
	os.WriteFile(filePath, []byte("hello"), 0644)

	e := &PatchFileExecutor{Root: tmpDir}
	result := e.Execute("patch_file", map[string]any{
		"path":        "nope.txt",
		"target":      "nonexistent",
		"replacement": "x",
	})
	if result.Error == nil {
		t.Error("expected error for target not found")
	}
}

func TestExecCommandExecutor_Echo(t *testing.T) {
	e := &ExecCommandExecutor{Root: "/tmp"}
	result := e.Execute("execute_command", map[string]any{
		"command": "echo 'hello harness'",
	})
	if result.Error != nil {
		t.Fatalf("unexpected error: %v", result.Error)
	}
	if result.ExitCode != 0 {
		t.Errorf("expected exit code 0, got %d", result.ExitCode)
	}
	if !strings.Contains(result.Output, "hello harness") {
		t.Errorf("expected output to contain 'hello harness', got %q", result.Output)
	}
}

func TestExecCommandExecutor_Error(t *testing.T) {
	e := &ExecCommandExecutor{Root: "/tmp"}
	result := e.Execute("execute_command", map[string]any{
		"command": "exit 42",
	})
	if result.ExitCode != 42 {
		t.Errorf("expected exit code 42, got %d", result.ExitCode)
	}
}

func TestCodeSearchExecutor_MissingQuery(t *testing.T) {
	e := &CodeSearchExecutor{}
	result := e.Execute("code_search", nil)
	if result.Error == nil {
		t.Error("expected error for missing query")
	}
}

func TestCodeSearchExecutor_QueryValidation(t *testing.T) {
	e := &CodeSearchExecutor{}
	result := e.Execute("code_search", map[string]any{
		"query": "HelloWorldTestQueryXYZ",
	})
	// This may fail in CI without network, but should not crash
	if result.Error != nil {
		t.Logf("code_search network-dependent: %v", result.Error)
	}
}

func TestWebSearchExecutor_MissingQuery(t *testing.T) {
	e := &WebSearchExecutor{}
	result := e.Execute("web_search", nil)
	if result.Error == nil {
		t.Error("expected error for missing query")
	}
}

func TestFetchContentExecutor_MissingArgs(t *testing.T) {
	e := &FetchContentExecutor{}
	result := e.Execute("fetch_content", nil)
	if result.Error == nil {
		t.Error("expected error for missing args")
	}
}

// --- mocks ---

type mockTool struct {
	output string
	err    error
}

func (m *mockTool) Execute(name string, args map[string]any) ToolResult {
	if m.err != nil {
		return ToolResult{Error: m.err}
	}
	return ToolResult{Output: m.output}
}
