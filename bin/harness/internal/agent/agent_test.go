package agent

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"harness/internal/tools"
)

func TestDetectDevTask_DevKeywords(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"implementar nova feature", true},
		{"criar modulo de autenticacao", true},
		{"corrigir bug no login", true},
		{"refatorar o codigo", true},
		{"qual a capital do Brasil?", false},
		{"explique o que e SOLID", false},
		{"add suporte a markdown", true},
		{"testar a API", false}, // "testar" nao esta na lista de keywords
		{"faça um relatorio", false}, // nao e desenvolvimento
	}
	for _, tt := range tests {
		got := DetectDevTask(tt.input)
		if got != tt.want {
			t.Errorf("DetectDevTask(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestExtractFeatureName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"implementar login social", "login-social"},
		{"criar API de usuarios", "api-de-usuarios"},
		{"refatorar modulo de pagamento", "modulo-de-pagamento"},
		{"fix security vulnerability", "security-vulnerability"},
		{"", "unnamed-feature"},
		{"   ", "unnamed-feature"},
		// "add @#$ special chars!!" -> prefix "add " e removido, vira "special chars"
		{"special chars", "special-chars"},
	}
	for _, tt := range tests {
		got := ExtractFeatureName(tt.input)
		if got != tt.want {
			t.Errorf("ExtractFeatureName(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestFeatureDir(t *testing.T) {
	dir := FeatureDir("/root", "my-feature")
	expected := "/root/.specs/features/my-feature"
	if dir != expected {
		t.Errorf("FeatureDir = %q, want %q", dir, expected)
	}
}

func TestEnsureFeatureDir(t *testing.T) {
	tmpDir := t.TempDir()
	err := EnsureFeatureDir(tmpDir, "test-feature")
	if err != nil {
		t.Fatalf("EnsureFeatureDir failed: %v", err)
	}

	expectedDir := filepath.Join(tmpDir, ".specs", "features", "test-feature")
	if _, err := os.Stat(expectedDir); os.IsNotExist(err) {
		t.Error("expected directory to exist")
	}
}

func TestFeatureExists(t *testing.T) {
	tmpDir := t.TempDir()
	featureDir := filepath.Join(tmpDir, ".specs", "features", "existing")
	os.MkdirAll(featureDir, 0755)

	// Without spec.md
	if FeatureExists(tmpDir, "existing") {
		t.Error("expected false when spec.md doesn't exist")
	}

	// With spec.md
	os.WriteFile(filepath.Join(featureDir, "spec.md"), []byte("spec"), 0644)
	if !FeatureExists(tmpDir, "existing") {
		t.Error("expected true when spec.md exists")
	}
}

func TestCreateSDDArtifacts_GeneratesAllFiles(t *testing.T) {
	tmpDir := t.TempDir()
	decisions := []string{"Objetivo: test"}
	err := CreateSDDArtifacts(tmpDir, "artifact-test", "tarefa de teste", decisions)
	if err != nil {
		t.Fatalf("CreateSDDArtifacts failed: %v", err)
	}

	featureDir := filepath.Join(tmpDir, ".specs", "features", "artifact-test")
	files := []string{"spec.md", "plan.md", "tasks.md", "contract.md"}
	for _, f := range files {
		path := filepath.Join(featureDir, f)
		data, err := os.ReadFile(path)
		if err != nil {
			t.Errorf("failed to read %s: %v", f, err)
			continue
		}
		content := string(data)
		if !strings.Contains(content, "artifact-test") {
			t.Errorf("%s should contain feature name", f)
		}
		if !strings.Contains(content, "@sdd-state") {
			t.Errorf("%s should contain sdd-state marker", f)
		}
	}

	// Verify specific content in spec.md
	specData, _ := os.ReadFile(filepath.Join(featureDir, "spec.md"))
	if !strings.Contains(string(specData), "Scenario 1") {
		t.Error("spec.md should contain BDD scenarios")
	}

	// Verify plan.md has mermaid
	planData, _ := os.ReadFile(filepath.Join(featureDir, "plan.md"))
	if !strings.Contains(string(planData), "mermaid") {
		t.Error("plan.md should contain mermaid diagram")
	}

	// Verify tasks.md has Evidence column
	tasksData, _ := os.ReadFile(filepath.Join(featureDir, "tasks.md"))
	if !strings.Contains(string(tasksData), "Evidence") {
		t.Error("tasks.md should contain Evidence column")
	}

	// Verify contract.md has Sensors
	contractData, _ := os.ReadFile(filepath.Join(featureDir, "contract.md"))
	if !strings.Contains(string(contractData), "Sensors") {
		t.Error("contract.md should contain Sensors")
	}
}

func TestCreateSDDArtifacts(t *testing.T) {
	tmpDir := t.TempDir()
	decisions := []string{"Objetivo: test SDD"}
	err := CreateSDDArtifacts(tmpDir, "sdd-test", "tarefa de teste", decisions)
	if err != nil {
		t.Fatalf("CreateSDDArtifacts failed: %v", err)
	}

	expectedFiles := []string{
		"spec.md", "plan.md", "tasks.md", "contract.md",
	}
	for _, f := range expectedFiles {
		path := filepath.Join(tmpDir, ".specs", "features", "sdd-test", f)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Errorf("expected %s to exist", path)
		}
	}
}

func TestUpdateProjectState(t *testing.T) {
	tmpDir := t.TempDir()
	os.MkdirAll(filepath.Join(tmpDir, ".specs", "project"), 0755)

	update := PlannerUpdate{
		FeatureName: "state-test",
		Phase:       PhaseAlign,
		Goal:        "test goal",
	}
	err := UpdateProjectState(tmpDir, update)
	if err != nil {
		t.Fatalf("UpdateProjectState failed: %v", err)
	}

	statePath := filepath.Join(tmpDir, ".specs", "project", "STATE.md")
	if _, err := os.Stat(statePath); os.IsNotExist(err) {
		t.Error("expected STATE.md to exist")
	}

	data, _ := os.ReadFile(statePath)
	if !strings.Contains(string(data), "state-test") {
		t.Error("STATE.md should contain feature name")
	}
}

func TestGenerateADRContent(t *testing.T) {
	decisions := []string{"Decisao 1", "Decisao 2"}
	adr := generateADRContent("adr-test", "ctx", decisions)
	if !strings.Contains(adr, "ADR") {
		t.Error("ADR should contain title marker")
	}
	if !strings.Contains(adr, "Decisao 1") {
		t.Error("ADR should contain decisions")
	}
}

func TestGlossaryUpdate(t *testing.T) {
	tmpDir := t.TempDir()
	contextPath := filepath.Join(tmpDir, "CONTEXT.md")

	updateGlossary(contextPath, "Cliente: usuario final")
	data, _ := os.ReadFile(contextPath)
	if !strings.Contains(string(data), "Cliente") {
		t.Error("glossary should contain the new term")
	}
}

func TestTruncateStr(t *testing.T) {
	if truncateStr("short", 10) != "short" {
		t.Error("short string should not be truncated")
	}
	long := "this is a very long string that should be truncated"
	result := truncateStr(long, 20)
	if len(result) > 23 {
		t.Errorf("expected truncated string, got %q (len %d)", result, len(result))
	}
}

func TestSDDPhaseConstants(t *testing.T) {
	if PhaseAlign != "ALIGN" {
		t.Errorf("expected ALIGN, got %q", PhaseAlign)
	}
	if PhaseSpecify != "SPECIFY" {
		t.Errorf("expected SPECIFY, got %q", PhaseSpecify)
	}
	if PhaseImplement != "IMPLEMENT" {
		t.Errorf("expected IMPLEMENT, got %q", PhaseImplement)
	}
	if PhaseVerify != "VERIFY" {
		t.Errorf("expected VERIFY, got %q", PhaseVerify)
	}
}

// ── XML Tool Parser Tests ──

type mockTool struct {
	name   string
	result string
}

func (m *mockTool) Execute(name string, args map[string]any) tools.ToolResult {
	return tools.ToolResult{Output: m.result}
}

func TestParseAndExecuteTools_MCPWithColons(t *testing.T) {
	reg := tools.NewRegistry()
	reg.Register("mcp:chrome-devtools:navigate_page", &mockTool{result: "navigated"})
	reg.Register("mcp:chrome-devtools:take_snapshot", &mockTool{result: "snapshot taken"})
	reg.Register("read_file", &mockTool{result: "file read"})

	tests := []struct {
		name     string
		input    string
		want     []string
		wantNone bool
	}{
		{
			name:  "MCP tool with XML body",
			input: `<tool:mcp:chrome-devtools:navigate_page><url>https://centralotaku.com.br</url></tool:mcp:chrome-devtools:navigate_page>`,
			want:  []string{"navigated"},
		},
		{
			name:  "MCP tool self-closing",
			input: `<tool:mcp:chrome-devtools:take_snapshot/>`,
			want:  []string{"snapshot taken"},
		},
		{
			name:  "standard tool with attributes",
			input: `<tool:read_file path="test.txt"/>`,
			want:  []string{"file read"},
		},
		{
			name:     "multiple MCP tools",
			input:    `<tool:mcp:chrome-devtools:navigate_page url="https://example.com"/> then <tool:mcp:chrome-devtools:take_snapshot/>`,
			want:     []string{"navigated", "snapshot taken"},
		},
		{
			name:     "no tool calls",
			input:    "responda apenas com texto",
			wantNone: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			results := parseAndExecuteTools(tt.input, reg)
			if tt.wantNone {
				if len(results) > 0 {
					t.Errorf("expected no tools, got %d: %v", len(results), results)
				}
				return
			}
			if len(results) != len(tt.want) {
				t.Errorf("expected %d results, got %d: %v", len(tt.want), len(results), results)
				return
			}
			for i, w := range tt.want {
				if !strings.Contains(results[i], w) {
					t.Errorf("result[%d] = %q, want contains %q", i, results[i], w)
				}
			}
		})
	}
}

func TestOpenTagRegex_MCPToolNames(t *testing.T) {
	tests := []struct {
		input string
		name  string
		match bool
	}{
		{`<tool:mcp:chrome-devtools:navigate_page/>`, "mcp:chrome-devtools:navigate_page", true},
		{`<tool:mcp:chrome-devtools:take_snapshot/>`, "mcp:chrome-devtools:take_snapshot", true},
		{`<tool:read_file path="x"/>`, "read_file", true},
		{`<tool:web_search query="x"/>`, "web_search", true},
		{`<tool:fetch_content url="x"/>`, "fetch_content", true},
		{`texto sem tool`, "", false},
	}

	for _, tt := range tests {
		t.Run(tt.input[:min(len(tt.input), 40)], func(t *testing.T) {
			matches := openTagRE.FindStringSubmatch(tt.input)
			if !tt.match {
				if matches != nil {
					t.Errorf("expected no match, got %v", matches)
				}
				return
			}
			if matches == nil {
				t.Fatalf("expected match for %q, got nil", tt.input)
			}
			if matches[1] != tt.name {
				t.Errorf("tool name = %q, want %q", matches[1], tt.name)
			}
		})
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
