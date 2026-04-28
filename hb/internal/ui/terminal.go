package ui

import (
	"os"

	"golang.org/x/term"
)

// TerminalSizeFunc allows mocking the terminal size in tests
var TerminalSizeFunc = func() (int, int, error) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 80, 24, err // Default fallback
	}
	return width, height, nil
}

// TerminalSize returns the width and height of the current terminal
func TerminalSize() (int, int, error) {
	return TerminalSizeFunc()
}
