package distiller

import (
	"strings"
	"testing"
)

func TestCheckAutoClarity_Security(t *testing.T) {
	if !CheckAutoClarity("security warning: unauthorized access") {
		t.Error("expected Auto-Clarity for security warning")
	}
}

func TestCheckAutoClarity_Irreversible(t *testing.T) {
	if !CheckAutoClarity("DROP TABLE users — cannot be undone") {
		t.Error("expected Auto-Clarity for irreversible action")
	}
}

func TestCheckAutoClarity_MultiStep(t *testing.T) {
	if !CheckAutoClarity("Step 1: do this.\nStep 2: do that.") {
		t.Error("expected Auto-Clarity for multi-step")
	}
}

func TestCheckAutoClarity_Safe(t *testing.T) {
	if CheckAutoClarity("the quick brown fox") {
		t.Error("did not expect Auto-Clarity for safe text")
	}
}

func TestToCaveman_Lite_RemovesFillers(t *testing.T) {
	result := ToCaveman("I really think we should basically do this.", Lite)
	if strings.Contains(result, "really") || strings.Contains(result, "basically") {
		t.Error("Lite should remove filler words")
	}
}

func TestToCaveman_Lite_KeepsArticles(t *testing.T) {
	input := "this is a test of the system"
	result := ToCaveman(input, Lite)
	if !strings.Contains(result, " a ") || !strings.Contains(result, " the ") {
		t.Errorf("Lite should keep articles, got: %q", result)
	}
}

func TestToCaveman_Full_RemovesArticles(t *testing.T) {
	result := ToCaveman("The quick brown fox jumps over a lazy dog.", Full)
	if strings.Contains(result, " the ") || strings.HasPrefix(result, "the ") {
		t.Error("Full should drop article 'the'")
	}
}

func TestToCaveman_Ultra_Abbreviates(t *testing.T) {
	result := ToCaveman("The database authentication function is extensive.", Ultra)
	if strings.Contains(result, "database") && !strings.Contains(result, "DB") {
		t.Error("Ultra should abbreviate 'database' to 'DB'")
	}
}

func TestToCaveman_Ultra_StripsConjunctions(t *testing.T) {
	result := ToCaveman("A and B or C but D", Ultra)
	if strings.Contains(result, " and ") || strings.Contains(result, " or ") {
		t.Error("Ultra should strip conjunctions")
	}
}

func TestToCaveman_BacktickProtection(t *testing.T) {
	input := "Run `DROP TABLE users` to delete."
	result := ToCaveman(input, Ultra)
	if !strings.Contains(result, "`DROP TABLE users`") {
		t.Error("backtick-protected code should remain unchanged")
	}
}

func TestToCaveman_AutoClarityBypass(t *testing.T) {
	result := ToCaveman("Warning: CVE-2024-1234 vulnerability", Ultra)
	if !strings.HasPrefix(result, "[AUTO-CLARITY TRIGGERED") {
		t.Error("expected Auto-Clarity bypass")
	}
}

func TestIsDestructiveCommand_RmRf(t *testing.T) {
	if !IsDestructiveCommand("rm -rf /") {
		t.Error("expected 'rm -rf /' to be destructive")
	}
}

func TestIsDestructiveCommand_DropTable(t *testing.T) {
	if !IsDestructiveCommand("DROP TABLE users;") {
		t.Error("expected 'DROP TABLE' to be destructive")
	}
}

func TestIsDestructiveCommand_SafeCommand(t *testing.T) {
	if IsDestructiveCommand("ls -la") {
		t.Error("did not expect 'ls -la' to be destructive")
	}
}

func TestIsDestructiveCommand_Echo(t *testing.T) {
	if IsDestructiveCommand("echo 'hello'") {
		t.Error("did not expect 'echo' to be destructive")
	}
}

func TestCompact_NoCompact(t *testing.T) {
	nodes := []HasContent{
		&testNode{role: "user", content: "short"},
		&testNode{role: "system", content: "short"},
	}
	compacted := Compact(nodes, 10, 5000)
	if len(compacted) != 2 {
		t.Errorf("expected 2 nodes, got %d", len(compacted))
	}
}

func TestCompact_CompactsOldSystemNode(t *testing.T) {
	longContent := strings.Repeat("A", 6000)
	nodes := make([]HasContent, 15)
	for i := 0; i < 15; i++ {
		nodes[i] = &testNode{role: "system", content: longContent}
	}

	compacted := Compact(nodes, 10, 5000)

	// First nodes (far from end) should be compacted
	if !strings.HasPrefix(compacted[0].GetContent(), "[Old tool result") {
		t.Error("expected old node to be compacted")
	}

	// Last node should NOT be compacted
	if strings.HasPrefix(compacted[14].GetContent(), "[Old tool result") {
		t.Error("last node should not be compacted")
	}
}

func TestCompact_KeepsNearNodes(t *testing.T) {
	longContent := strings.Repeat("A", 6000)
	nodes := make([]HasContent, 3)
	for i := 0; i < 3; i++ {
		nodes[i] = &testNode{role: "system", content: longContent}
	}

	compacted := Compact(nodes, 10, 5000)
	// All 3 nodes are within 10 turns of end, none should be compacted
	if strings.HasPrefix(compacted[0].GetContent(), "[Old tool result") {
		t.Error("near nodes should not be compacted")
	}
}

func TestCompact_NonSystemNodes(t *testing.T) {
	longContent := strings.Repeat("A", 6000)
	nodes := make([]HasContent, 15)
	for i := 0; i < 15; i++ {
		role := "user"
		if i%2 == 0 {
			role = "assistant"
		}
		nodes[i] = &testNode{role: role, content: longContent}
	}

	compacted := Compact(nodes, 10, 5000)
	// User/assistant nodes should NOT be compacted even if old and long
	if strings.HasPrefix(compacted[0].GetContent(), "[Old tool result") {
		t.Error("non-system nodes should not be compacted")
	}
}

// --- test helper ---

type testNode struct {
	role    string
	content string
}

func (n *testNode) GetContent() string   { return n.content }
func (n *testNode) SetContent(s string)  { n.content = s }
func (n *testNode) GetRole() string      { return n.role }
