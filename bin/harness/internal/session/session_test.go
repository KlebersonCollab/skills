package session

import (
	"path/filepath"
	"strings"
	"testing"
)

func TestNewTree(t *testing.T) {
	tree := NewTree("test-sess", "test task")
	if tree.SessionID != "test-sess" {
		t.Errorf("expected SessionID 'test-sess', got %q", tree.SessionID)
	}
	if tree.RootTask != "test task" {
		t.Errorf("expected RootTask 'test task', got %q", tree.RootTask)
	}
	if len(tree.Nodes) != 0 {
		t.Errorf("expected empty nodes, got %d", len(tree.Nodes))
	}
}

func TestAddNode(t *testing.T) {
	tree := NewTree("s1", "task")
	n1 := tree.AddNode("system", "init", 10)

	if tree.ActiveNode != n1 {
		t.Errorf("expected ActiveNode %q, got %q", n1, tree.ActiveNode)
	}
	if len(tree.Nodes) != 1 {
		t.Errorf("expected 1 node, got %d", len(tree.Nodes))
	}

	node := tree.Nodes[n1]
	if node.Role != "system" {
		t.Errorf("expected role 'system', got %q", node.Role)
	}
	if node.Content != "init" {
		t.Errorf("expected content 'init', got %q", node.Content)
	}
	if node.Tokens != 10 {
		t.Errorf("expected 10 tokens, got %d", node.Tokens)
	}
	if node.ParentID != "" {
		t.Errorf("expected empty ParentID for root, got %q", node.ParentID)
	}
}

func TestAddNode_Chain(t *testing.T) {
	tree := NewTree("s1", "task")
	n1 := tree.AddNode("system", "sys", 0)
	n2 := tree.AddNode("user", "hello", 5)
	n3 := tree.AddNode("assistant", "hi", 3)

	if tree.Nodes[n2].ParentID != n1 {
		t.Errorf("expected n2 parent to be n1, got %q", tree.Nodes[n2].ParentID)
	}
	if tree.Nodes[n3].ParentID != n2 {
		t.Errorf("expected n3 parent to be n2, got %q", tree.Nodes[n3].ParentID)
	}
}

func TestGetLinearHistory(t *testing.T) {
	tree := NewTree("s1", "task")
	n1 := tree.AddNode("system", "sys", 0)
	tree.AddNode("user", "hello", 5) // n2
	n3 := tree.AddNode("assistant", "hi", 3)

	history, err := tree.GetLinearHistory(n3)
	if err != nil {
		t.Fatalf("GetLinearHistory failed: %v", err)
	}

	if len(history) != 3 {
		t.Errorf("expected 3 nodes in history, got %d", len(history))
	}
	if history[0].NodeID != n1 {
		t.Errorf("expected first node to be %q, got %q", n1, history[0].NodeID)
	}
	if history[2].NodeID != n3 {
		t.Errorf("expected last node to be %q, got %q", n3, history[2].NodeID)
	}
}

func TestGetLinearHistory_Empty(t *testing.T) {
	tree := NewTree("s1", "task")
	_, err := tree.GetLinearHistory("nonexistent")
	if err == nil {
		t.Error("expected error for nonexistent node")
	}
}

func TestGetLinearHistory_FromActive(t *testing.T) {
	tree := NewTree("s1", "task")
	tree.AddNode("system", "sys", 0)
	n2 := tree.AddNode("user", "hello", 5)

	history, err := tree.GetLinearHistory("")
	if err != nil {
		t.Fatalf("GetLinearHistory failed: %v", err)
	}
	if len(history) != 2 {
		t.Errorf("expected 2 nodes, got %d", len(history))
	}
	if history[len(history)-1].NodeID != n2 {
		t.Errorf("expected last node to be active %q, got %q", n2, history[len(history)-1].NodeID)
	}
}

func TestFormatHistory(t *testing.T) {
	tree := NewTree("s1", "task")
	tree.AddNode("system", "sys msg", 0)
	tree.AddNode("user", "user msg", 5)
	tree.AddNode("assistant", "assistant msg", 3)

	formatted, err := tree.FormatHistory("")
	if err != nil {
		t.Fatalf("FormatHistory failed: %v", err)
	}

	if !strings.Contains(formatted, "[System]") {
		t.Errorf("expected [System] in history, got %q", formatted)
	}
	if !strings.Contains(formatted, "[User]") {
		t.Errorf("expected [User] in history, got %q", formatted)
	}
	if !strings.Contains(formatted, "[Assistant]") {
		t.Errorf("expected [Assistant] in history, got %q", formatted)
	}
}

func TestTotalTokens(t *testing.T) {
	tree := NewTree("s1", "task")
	tree.AddNode("system", "a", 10)
	tree.AddNode("user", "b", 20)
	tree.AddNode("assistant", "c", 30)

	total := tree.TotalTokens()
	if total != 60 {
		t.Errorf("expected 60 tokens, got %d", total)
	}
}

func TestTotalTokens_Empty(t *testing.T) {
	tree := NewTree("s1", "task")
	total := tree.TotalTokens()
	if total != 0 {
		t.Errorf("expected 0 tokens for empty tree, got %d", total)
	}
}

func TestSaveAndLoadTree(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "session.json")

	original := NewTree("s1", "task")
	original.AddNode("system", "sys", 5)
	original.AddNode("user", "hello", 3)

	if err := original.Save(path); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	loaded, err := LoadTree(path)
	if err != nil {
		t.Fatalf("LoadTree failed: %v", err)
	}

	if loaded.SessionID != "s1" {
		t.Errorf("expected SessionID 's1', got %q", loaded.SessionID)
	}
	if len(loaded.Nodes) != 2 {
		t.Errorf("expected 2 nodes, got %d", len(loaded.Nodes))
	}
}

func TestLoadTree_NotFound(t *testing.T) {
	_, err := LoadTree("/nonexistent/path.json")
	if err == nil {
		t.Error("expected error for nonexistent file")
	}
}

func TestEstimateUSD_Gemini(t *testing.T) {
	tree := NewTree("s1", "task")
	tree.AddNode("user", "hello", 1000)

	cost := tree.EstimateUSD("gemini")
	if cost != 0.0035 {
		t.Errorf("expected $0.0035 for 1000 gemini tokens, got %f", cost)
	}
}

func TestEstimateUSD_DeepSeek(t *testing.T) {
	tree := NewTree("s1", "task")
	tree.AddNode("user", "hello", 1000)

	cost := tree.EstimateUSD("deepseek")
	if cost != 0.002 {
		t.Errorf("expected $0.002 for 1000 deepseek tokens, got %f", cost)
	}
}

func TestEstimateUSD_Ollama(t *testing.T) {
	tree := NewTree("s1", "task")
	tree.AddNode("user", "hello", 1000)

	cost := tree.EstimateUSD("ollama")
	if cost != 0.0 {
		t.Errorf("expected $0.0 for ollama, got %f", cost)
	}
}

func TestNodeImplementsHasContent(t *testing.T) {
	n := &Node{
		NodeID:  "n1",
		Role:    "user",
		Content: "test content",
	}

	if n.GetContent() != "test content" {
		t.Errorf("expected 'test content', got %q", n.GetContent())
	}
	if n.GetRole() != "user" {
		t.Errorf("expected role 'user', got %q", n.GetRole())
	}

	n.SetContent("updated")
	if n.Content != "updated" {
		t.Errorf("expected Content 'updated', got %q", n.Content)
	}
}

func TestDiscoverSkills_NoDir(t *testing.T) {
	tmpDir := t.TempDir()
	skills := DiscoverSkills(tmpDir)
	if len(skills) != 0 {
		t.Errorf("expected 0 skills in empty dir, got %d", len(skills))
	}
}

func TestGetSkillsMetrics(t *testing.T) {
	avail, active := GetSkillsMetrics("/tmp")
	// Should not panic
	_ = avail
	_ = active
}
