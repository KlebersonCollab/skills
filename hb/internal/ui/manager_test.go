package ui

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestUIManager_Basic(t *testing.T) {
	var buf bytes.Buffer
	m := NewUIManager(&buf)
	
	// Force isTerminal for testing
	m.isTerminal = true

	m.SetState(UIState{
		FeatureName: "TestFeature",
		CurrentTask: "Coding the UI",
		Progress:    0.5,
		TokenCount:  1000,
		Cost:        0.02,
	})

	m.Start()
	// Wait for at least one tick
	time.Sleep(150 * time.Millisecond)
	m.Stop()

	output := buf.String()
	
	if !strings.Contains(output, "TestFeature") {
		t.Errorf("Output should contain FeatureName, got: %s", output)
	}
	if !strings.Contains(output, "Coding the UI") {
		t.Errorf("Output should contain Task description")
	}
	if !strings.Contains(output, "50%") {
		t.Errorf("Output should contain progress percentage")
	}
	if !strings.Contains(output, HideCursor) {
		t.Errorf("Output should hide cursor")
	}
	if !strings.Contains(output, ShowCursor) {
		t.Errorf("Output should restore cursor")
	}
}

func TestTruncate(t *testing.T) {
	tests := []struct {
		s      string
		max    int
		expect string
	}{
		{"Hello World", 5, "He..."},
		{"Hi", 5, "Hi"},
		{"Hello World", 2, "..."},
	}

	for _, tt := range tests {
		got := truncate(tt.s, tt.max)
		if got != tt.expect {
			t.Errorf("truncate(%q, %d) = %q, expect %q", tt.s, tt.max, got, tt.expect)
		}
	}
}

func TestUIManager_NonTerminal(t *testing.T) {
	var buf bytes.Buffer
	m := NewUIManager(&buf)
	m.isTerminal = false

	m.Start()
	m.SetState(UIState{FeatureName: "ShouldNotSeeThis"})
	time.Sleep(50 * time.Millisecond)
	m.Stop()

	if buf.Len() > 0 {
		t.Errorf("Non-terminal mode should not write to output, got: %s", buf.String())
	}
}

func TestUIManager_EmptyFeature(t *testing.T) {
	var buf bytes.Buffer
	m := NewUIManager(&buf)
	m.isTerminal = true
	m.SetState(UIState{FeatureName: ""})
	m.render()
	
	if !strings.Contains(buf.String(), "[IDLE]") {
		t.Errorf("Empty feature should show [IDLE], got: %s", buf.String())
	}
}

func TestTerminalSize_Fallback(t *testing.T) {
	// Since we can't easily force an error on os.Stdout in tests, 
	// we just verify the function exists and runs.
	w, h, _ := TerminalSize()
	if w == 0 || h == 0 {
		t.Errorf("TerminalSize returned zeros")
	}
}

func TestIsTerminal_NonFile(t *testing.T) {
	var buf bytes.Buffer
	if isTerminal(&buf) {
		t.Errorf("isTerminal should be false for bytes.Buffer")
	}
}

func TestTerminalSize_Error(t *testing.T) {
	oldFunc := TerminalSizeFunc
	defer func() { TerminalSizeFunc = oldFunc }()
	
	TerminalSizeFunc = func() (int, int, error) {
		return 80, 24, fmt.Errorf("forced error")
	}
	
	w, h, err := TerminalSize()
	if err == nil {
		t.Errorf("Expected error")
	}
	if w != 80 || h != 24 {
		t.Errorf("Expected fallback values")
	}
}

func TestUIManager_SmallWidth(t *testing.T) {
	oldFunc := TerminalSizeFunc
	defer func() { TerminalSizeFunc = oldFunc }()
	
	TerminalSizeFunc = func() (int, int, error) {
		return 5, 5, nil
	}
	
	var buf bytes.Buffer
	m := NewUIManager(&buf)
	m.isTerminal = true
	m.render()
	
	if !strings.Contains(buf.String(), "[IDLE]") {
		t.Errorf("Should handle small width gracefully")
	}
}

func TestTruncate_EdgeCases(t *testing.T) {
	if truncate("test", 1) != "..." {
		t.Errorf("Expected ... for tiny maxLen")
	}
}
