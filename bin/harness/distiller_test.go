package main

import (
	"strings"
	"testing"

	"harness/internal/distiller"
)

func TestCheckAutoClarity_Security(t *testing.T) {
	input := "This is a security warning: unauthorized access detected."
	if !distiller.CheckAutoClarity(input) {
		t.Error("Expected Auto-Clarity for security warning")
	}
}

func TestCheckAutoClarity_Irreversible(t *testing.T) {
	input := "Executing DROP TABLE users — this cannot be undone."
	if !distiller.CheckAutoClarity(input) {
		t.Error("Expected Auto-Clarity for irreversible action")
	}
}

func TestCheckAutoClarity_MultiStep(t *testing.T) {
	input := "Step 1: do this.\nStep 2: do that."
	if !distiller.CheckAutoClarity(input) {
		t.Error("Expected Auto-Clarity for multi-step pattern")
	}
}

func TestCheckAutoClarity_Safe(t *testing.T) {
	input := "The quick brown fox jumps over the lazy dog."
	if distiller.CheckAutoClarity(input) {
		t.Error("Did not expect Auto-Clarity for safe text")
	}
}

func TestToCaveman_Lite(t *testing.T) {
	input := "I really think we should basically implement a solution for this very extensive problem."
	result := distiller.ToCaveman(input, distiller.Lite)
	if strings.Contains(result, "really") || strings.Contains(result, "basically") {
		t.Error("Lite should remove filler words")
	}
	if !strings.Contains(result, "a") && !strings.Contains(result, "the") {
		t.Error("Lite should keep articles (a/the)")
	}
	if !strings.Contains(result, "implement a solution for") {
		t.Error("Lite should NOT replace synonyms; 'implement a solution for' should remain")
	}
}

func TestToCaveman_Full(t *testing.T) {
	input := "The quick brown fox jumps over a lazy dog. I really think we should simply implement a solution."
	result := distiller.ToCaveman(input, distiller.Full)
	if strings.Contains(result, " the ") || strings.HasPrefix(result, "the ") {
		t.Error("Full should drop article 'the', got:", result)
	}
	if strings.Contains(result, " a ") || strings.HasPrefix(result, "a ") || strings.HasSuffix(result, " a") {
		t.Error("Full should drop article 'a', got:", result)
	}
	if strings.Contains(result, "really") || strings.Contains(result, "simply") {
		t.Error("Full should remove filler words")
	}
}

func TestToCaveman_Ultra(t *testing.T) {
	input := "The database authentication request function implementation is extensive and causes connection issues."
	result := distiller.ToCaveman(input, distiller.Ultra)
	if strings.Contains(result, "database") && !strings.Contains(result, "DB") {
		t.Error("Ultra should abbreviate 'database' to 'DB'")
	}
	if strings.Contains(result, "and") {
		t.Error("Ultra should strip conjunctions like 'and'")
	}
}

func TestToCaveman_BacktickProtection(t *testing.T) {
	input := "Run `DROP TABLE users` to delete data."
	result := distiller.ToCaveman(input, distiller.Ultra)
	if !strings.Contains(result, "`DROP TABLE users`") {
		t.Error("Backtick-protected code should remain unchanged, got:", result)
	}
}

func TestToCaveman_AutoClarityBypass(t *testing.T) {
	input := "Warning: this is a security vulnerability CVE-2024-1234."
	result := distiller.ToCaveman(input, distiller.Ultra)
	if !strings.HasPrefix(result, "[AUTO-CLARITY TRIGGERED") {
		t.Error("Expected Auto-Clarity bypass, got:", result)
	}
}

type testNode struct {
	role    string
	content string
}

func (n *testNode) GetContent() string   { return n.content }
func (n *testNode) SetContent(s string)  { n.content = s }
func (n *testNode) GetRole() string      { return n.role }

func TestCompactHistory_NoCompact(t *testing.T) {
	nodes := []distiller.HasContent{
		&testNode{role: "user", content: "short"},
		&testNode{role: "assistant", content: "short"},
		&testNode{role: "system", content: "short"},
	}
	compacted := distiller.Compact(nodes, 10, 5000)
	if len(compacted) != 3 {
		t.Error("Expected no change, got length", len(compacted))
	}
}

func TestCompactHistory_CompactsOldToolOutput(t *testing.T) {
	longContent := strings.Repeat("A", 6000)
	nodes := make([]distiller.HasContent, 15)
	for i := 0; i < 15; i++ {
		nodes[i] = &testNode{role: "system", content: longContent}
	}
	compacted := distiller.Compact(nodes, 10, 5000)
	if len(compacted) != 15 {
		t.Error("Length should remain 15, got", len(compacted))
	}
	if !strings.HasPrefix(compacted[0].GetContent(), "[Old tool result cleared") {
		t.Error("Expected stub prefix, got:", compacted[0].GetContent()[:50])
	}
	if strings.HasPrefix(compacted[14].GetContent(), "[Old tool result cleared") {
		t.Error("Last node should not be compacted")
	}
}
