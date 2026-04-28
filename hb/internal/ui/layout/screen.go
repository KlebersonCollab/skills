package layout

import (
	"strings"
)

// Screen represents a virtual terminal screen buffer.
type Screen struct {
	Width, Height int
	Buffer        [][]rune
}

// NewScreen creates a new screen buffer with the given dimensions.
func NewScreen(w, h int) *Screen {
	buf := make([][]rune, h)
	for i := range buf {
		buf[i] = make([]rune, w)
		for j := range buf[i] {
			buf[i][j] = ' '
		}
	}
	return &Screen{Width: w, Height: h, Buffer: buf}
}

// DrawText writes text at the given coordinates.
func (s *Screen) DrawText(x, y int, text string) {
	runes := []rune(text)
	for i, r := range runes {
		currX := x + i
		if currX >= 0 && currX < s.Width && y >= 0 && y < s.Height {
			s.Buffer[y][currX] = r
		}
	}
}

// DrawBox draws a border box at the given coordinates.
func (s *Screen) DrawBox(x, y, w, h int) {
	if w <= 0 || h <= 0 {
		return
	}

	// Corners
	s.setSafe(x, y, '┌')
	s.setSafe(x+w-1, y, '┐')
	s.setSafe(x, y+h-1, '└')
	s.setSafe(x+w-1, y+h-1, '┘')

	// Horizontal lines
	for i := 1; i < w-1; i++ {
		s.setSafe(x+i, y, '─')
		s.setSafe(x+i, y+h-1, '─')
	}

	// Vertical lines
	for j := 1; j < h-1; j++ {
		s.setSafe(x, y+j, '│')
		s.setSafe(x+w-1, y+j, '│')
	}
}

func (s *Screen) setSafe(x, y int, r rune) {
	if x >= 0 && x < s.Width && y >= 0 && y < s.Height {
		s.Buffer[y][x] = r
	}
}

// String returns the string representation of the screen.
func (s *Screen) String() string {
	var sb strings.Builder
	for i, row := range s.Buffer {
		sb.WriteString(string(row))
		if i < len(s.Buffer)-1 {
			sb.WriteByte('\n')
		}
	}
	return sb.String()
}
