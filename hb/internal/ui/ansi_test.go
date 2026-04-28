package ui

import (
	"testing"
)

func TestANSIHelpers(t *testing.T) {
	tests := []struct {
		name     string
		got      string
		expected string
	}{
		{"MoveTo", MoveTo(10, 20), "\x1b[10;20H"},
		{"MoveUp", MoveUp(5), "\x1b[5A"},
		{"MoveUp Zero", MoveUp(0), ""},
		{"MoveDown", MoveDown(3), "\x1b[3B"},
		{"MoveDown Zero", MoveDown(0), ""},
		{"Colorize", Colorize("Hello", Green), "\x1b[32mHello\x1b[0m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, tt.got)
			}
		})
	}
}

func TestConstants(t *testing.T) {
	// Just to touch the constants for coverage (though Go doesn't require this for coverage, it's good practice)
	if CSI != "\x1b[" {
		t.Errorf("CSI constant is wrong")
	}
}
