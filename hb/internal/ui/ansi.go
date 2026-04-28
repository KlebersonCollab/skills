package ui

import "fmt"

// ANSI escape sequences
const (
	Escape = "\x1b"
	CSI    = Escape + "["

	// Cursor management
	HideCursor    = CSI + "?25l"
	ShowCursor    = CSI + "?25h"
	SaveCursor    = CSI + "s"
	RestoreCursor = CSI + "u"

	// Clearing
	ClearScreen     = CSI + "2J"
	ClearLine       = CSI + "2K"
	ClearLineRight  = CSI + "K"
	ClearBelow      = CSI + "J"

	// Formatting
	Reset     = CSI + "0m"
	Bold      = CSI + "1m"
	Dim       = CSI + "2m"
	Italic    = CSI + "3m"
	Underline = CSI + "4m"
	Inverse   = CSI + "7m"

	// Colors (Standard)
	Black   = CSI + "30m"
	Red     = CSI + "31m"
	Green   = CSI + "32m"
	Yellow  = CSI + "33m"
	Blue    = CSI + "34m"
	Magenta = CSI + "35m"
	Cyan    = CSI + "36m"
	White   = CSI + "37m"

	// Bright Colors
	BrightBlack   = CSI + "90m"
	BrightRed     = CSI + "91m"
	BrightGreen   = CSI + "92m"
	BrightYellow  = CSI + "93m"
	BrightBlue    = CSI + "94m"
	BrightMagenta = CSI + "95m"
	BrightCyan    = CSI + "96m"
	BrightWhite   = CSI + "97m"
)

// MoveTo moves the cursor to a specific row and column (1-indexed)
func MoveTo(row, col int) string {
	return fmt.Sprintf("%s%d;%dH", CSI, row, col)
}

// MoveUp moves the cursor up N lines
func MoveUp(n int) string {
	if n <= 0 {
		return ""
	}
	return fmt.Sprintf("%s%dA", CSI, n)
}

// MoveDown moves the cursor down N lines
func MoveDown(n int) string {
	if n <= 0 {
		return ""
	}
	return fmt.Sprintf("%s%dB", CSI, n)
}

// Colorize wraps text in ANSI color codes
func Colorize(text, color string) string {
	return color + text + Reset
}
