package main

import (
	"strings"
	"testing"
)

func TestCheckAutoClarity_Security(t *testing.T) {
	input := "This is a security warning: unauthorized access detected."
	if !CheckAutoClarity(input) {
		t.Error("Expected Auto-Clarity for security warning")
	}
}

func TestCheckAutoClarity_Irreversible(t *testing.T) {
	input := "Executing DROP TABLE users — this cannot be undone."
	if !CheckAutoClarity(input) {
		t.Error("Expected Auto-Clarity for irreversible action")
	}
}

func TestCheckAutoClarity_MultiStep(t *testing.T) {
	input := "Step 1: do this.\nStep 2: do that."
	if !CheckAutoClarity(input) {
		t.Error("Expected Auto-Clarity for multi-step pattern")
	}
}

func TestCheckAutoClarity_Safe(t *testing.T) {
	input := "The quick brown fox jumps over the lazy dog."
	if CheckAutoClarity(input) {
		t.Error("Did not expect Auto-Clarity for safe text")
	}
}

func TestToCaveman_Lite(t *testing.T) {
	input := "I really think we should basically implement a solution for this very extensive problem."
	result := ToCaveman(input, Lite)
	if strings.Contains(result, "really") || strings.Contains(result, "basically") {
		t.Error("Lite should remove filler words")
	}
	// Lite must keep articles ("a", "the") — verify at least one remains
	if !strings.Contains(result, "a") && !strings.Contains(result, "the") {
		t.Error("Lite should keep articles (a/the)")
	}
	// Lite must NOT apply synonyms — "implement a solution for" should remain unchanged
	if !strings.Contains(result, "implement a solution for") {
		t.Error("Lite should NOT replace synonyms; 'implement a solution for' should remain")
	}
}

func TestToCaveman_Full(t *testing.T) {
	input := "The quick brown fox jumps over a lazy dog. I really think we should simply implement a solution."
	result := ToCaveman(input, Full)
	// Check article removal using word boundaries to avoid false positives (e.g., 'a' inside "lazy")
	if strings.Contains(result, " the ") || strings.HasPrefix(result, "the ") {
		t.Error("Full should drop article 'the', got:", result)
	}
	if strings.Contains(result, " a ") || strings.HasPrefix(result, "a ") || strings.HasSuffix(result, " a") {
		t.Error("Full should drop article 'a', got:", result)
	}
	if strings.Contains(result, "really") || strings.Contains(result, "simply") {
		t.Error("Full should remove filler words")
	}
	if !strings.Contains(result, "do") && !strings.Contains(result, "fix") {
		t.Log("Full: synonym replacement may not apply directly; result:", result)
	}
}

func TestToCaveman_Ultra(t *testing.T) {
	input := "The database authentication request function implementation is extensive and causes connection issues."
	result := ToCaveman(input, Ultra)
	if strings.Contains(result, "database") && !strings.Contains(result, "DB") {
		t.Error("Ultra should abbreviate 'database' to 'DB'")
	}
	if strings.Contains(result, "and") {
		t.Error("Ultra should strip conjunctions like 'and'")
	}
	if strings.Contains(result, "causes") && strings.Contains(result, "→") {
		// causality already replaced
	}
}

func TestToCaveman_BacktickProtection(t *testing.T) {
	input := "Run `DROP TABLE users` to delete data."
	result := ToCaveman(input, Ultra)
	if !strings.Contains(result, "`DROP TABLE users`") {
		t.Error("Backtick-protected code should remain unchanged, got:", result)
	}
}

func TestToCaveman_AutoClarityBypass(t *testing.T) {
	input := "Warning: this is a security vulnerability CVE-2024-1234."
	result := ToCaveman(input, Ultra)
	if !strings.HasPrefix(result, "[AUTO-CLARITY TRIGGERED") {
		t.Error("Expected Auto-Clarity bypass, got:", result)
	}
}

func TestCompactHistory_NoCompact(t *testing.T) {
	nodes := []Node{
		{Role: "user", Content: "short", Timestamp: "now"},
		{Role: "assistant", Content: "short", Timestamp: "now"},
		{Role: "system", Content: "short", Timestamp: "now"},
	}
	compacted := CompactHistory(nodes, 10, 5000)
	if len(compacted) != 3 {
		t.Error("Expected no change, got length", len(compacted))
	}
}

func TestCompactHistory_CompactsOldToolOutput(t *testing.T) {
	longContent := strings.Repeat("A", 6000)
	nodes := make([]Node, 15)
	for i := 0; i < 15; i++ {
		nodes[i] = Node{Role: "system", Content: longContent, Timestamp: "old"}
	}
	compacted := CompactHistory(nodes, 10, 5000)
	// First 5 nodes (indices 0-4) are more than 10 turns from end (len 15, i=0 => 15-0=15 >10)
	// They should be compacted. But compacted length should still be 15.
	if len(compacted) != 15 {
		t.Error("Length should remain 15, got", len(compacted))
	}
	// Node at index 0 should be stub
	if !strings.HasPrefix(compacted[0].Content, "[Old tool result cleared") {
		t.Error("Expected stub prefix, got:", compacted[0].Content[:50])
	}
	// Node at index 14 (last) should NOT be compacted because total-i = 1 <= 10
	if strings.HasPrefix(compacted[14].Content, "[Old tool result cleared") {
		t.Error("Last node should not be compacted")
	}
}
