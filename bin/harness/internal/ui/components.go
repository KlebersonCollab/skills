// Package ui provides rich terminal UI components for the Harness agent.
package ui

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

// Component is the base interface for all TUI components.
type Component interface {
	// Render returns the component as an array of lines, each no wider than width.
	Render(width int) []string

	// Invalidate clears any cached render state.
	Invalidate()
}

// Theme defines colors for TUI components.
type Theme struct {
	Accent   func(string) string
	Success  func(string) string
	Warning  func(string) string
	Error    func(string) string
	Muted    func(string) string
	Dim      func(string) string
	Selected func(string) string
	Border   func(string) string
}

// DefaultTheme returns the default ANSI theme.
func DefaultTheme() Theme {
	return Theme{
		Accent:   func(s string) string { return "\033[38;5;99m" + s + "\033[0m" },
		Success:  func(s string) string { return "\033[1;32m" + s + "\033[0m" },
		Warning:  func(s string) string { return "\033[38;5;208m" + s + "\033[0m" },
		Error:    func(s string) string { return "\033[38;5;196m" + s + "\033[0m" },
		Muted:    func(s string) string { return "\033[38;5;248m" + s + "\033[0m" },
		Dim:      func(s string) string { return "\033[90m" + s + "\033[0m" },
		Selected: func(s string) string { return "\033[48;5;99m\033[38;5;15m" + s + "\033[0m" },
		Border:   func(s string) string { return "\033[38;5;242m" + s + "\033[0m" },
	}
}

// --- Text Component ---

// Text renders multi-line text with optional word wrap.
type Text struct {
	content  string
	paddingX int
	paddingY int
	bgFn     func(string) string
}

// NewText creates a new Text component.
func NewText(content string, paddingX, paddingY int) *Text {
	return &Text{
		content:  content,
		paddingX: paddingX,
		paddingY: paddingY,
	}
}

// SetText updates the text content and invalidates.
func (t *Text) SetText(content string) { t.content = content }

// SetBgFn sets an optional background function (e.g., for highlighting).
func (t *Text) SetBgFn(fn func(string) string) { t.bgFn = fn }

func (t *Text) Render(width int) []string {
	if width <= 0 {
		return nil
	}

	contentWidth := width - 2*t.paddingX
	if contentWidth <= 0 {
		contentWidth = 1
	}

	var lines []string

	// Top padding
	for i := 0; i < t.paddingY; i++ {
		lines = append(lines, strings.Repeat(" ", width))
	}

	// Word wrap content
	wrapped := wordWrap(t.content, contentWidth)
	for _, line := range wrapped {
		padded := strings.Repeat(" ", t.paddingX) + line + strings.Repeat(" ", width-len(line)-2*t.paddingX)
		if t.bgFn != nil {
			padded = t.bgFn(padded)
		}
		lines = append(lines, padded)
	}

	// Bottom padding
	for i := 0; i < t.paddingY; i++ {
		lines = append(lines, strings.Repeat(" ", width))
	}

	return lines
}

func (t *Text) Invalidate() {}

// --- Spacer Component ---

type Spacer struct {
	height int
}

func NewSpacer(height int) *Spacer {
	return &Spacer{height: height}
}

func (s *Spacer) Render(width int) []string {
	lines := make([]string, s.height)
	for i := range lines {
		lines[i] = ""
	}
	return lines
}

func (s *Spacer) Invalidate() {}

// --- Box Component ---

type Box struct {
	paddingX int
	paddingY int
	bgFn     func(string) string
	children []Component
}

func NewBox(paddingX, paddingY int, bgFn func(string) string) *Box {
	return &Box{
		paddingX: paddingX,
		paddingY: paddingY,
		bgFn:    bgFn,
	}
}

func (b *Box) AddChild(child Component) {
	b.children = append(b.children, child)
}

func (b *Box) Clear() {
	b.children = nil
}

func (b *Box) Render(width int) []string {
	var lines []string

	// Top padding
	for i := 0; i < b.paddingY; i++ {
		line := strings.Repeat(" ", width)
		if b.bgFn != nil {
			line = b.bgFn(line)
		}
		lines = append(lines, line)
	}

	// Children
	for _, child := range b.children {
		childLines := child.Render(width)
		for _, line := range childLines {
			if line == "" {
				line = strings.Repeat(" ", width)
			}
			if b.bgFn != nil {
				// Apply background only to padding area
				line = strings.Repeat(" ", b.paddingX) + line
				line = b.bgFn(line)
			}
			lines = append(lines, line)
		}
	}

	// Bottom padding
	for i := 0; i < b.paddingY; i++ {
		line := strings.Repeat(" ", width)
		if b.bgFn != nil {
			line = b.bgFn(line)
		}
		lines = append(lines, line)
	}

	return lines
}

func (b *Box) Invalidate() {
	for _, child := range b.children {
		child.Invalidate()
	}
}

// --- Container Component ---

type Container struct {
	children []Component
}

func NewContainer() *Container {
	return &Container{}
}

func (c *Container) AddChild(child Component) {
	c.children = append(c.children, child)
}

func (c *Container) Clear() {
	c.children = nil
}

func (c *Container) Render(width int) []string {
	var lines []string
	for _, child := range c.children {
		childLines := child.Render(width)
		lines = append(lines, childLines...)
	}
	return lines
}

func (c *Container) Invalidate() {
	for _, child := range c.children {
		child.Invalidate()
	}
}

// --- Border Component ---

type Border struct {
	title   string
	theme   Theme
	style   string // "single", "double", "none"
}

func NewBorder(title string, theme Theme) *Border {
	return &Border{
		title: title,
		theme: theme,
		style: "single",
	}
}

func (b *Border) Render(width int) []string {
	if width <= 0 {
		return nil
	}

	if width < 4 {
		width = 4
	}

	var top, bottom string

	switch b.style {
	case "double":
		top = "╔" + "═"
		bottom = "╚" + "═"
	default:
		top = "┌" + "─"
		bottom = "└" + "─"
	}

	if b.title != "" {
		titleStr := " " + b.title + " "
		repeatCount := width - utf8.RuneCountInString(top) - utf8.RuneCountInString(titleStr) - 1
		if repeatCount < 0 {
			repeatCount = 0
		}
		top += titleStr + strings.Repeat("─", repeatCount)
	} else {
		top += strings.Repeat("─", width-2)
	}
	top += "┐"

	bottom += strings.Repeat("─", width-2) + "┘"

	return []string{
		b.theme.Border(top),
		b.theme.Border(bottom),
	}
}

func (b *Border) Invalidate() {}

// --- Progress Bar Component ---

type ProgressBar struct {
	percent  float64
	width    int
	label    string
	theme    Theme
}

func NewProgressBar(theme Theme) *ProgressBar {
	return &ProgressBar{
		theme: theme,
	}
}

func (p *ProgressBar) SetProgress(percent float64, label string) {
	p.percent = percent
	p.label = label
}

func (p *ProgressBar) Render(width int) []string {
	if width <= 0 {
		return nil
	}

	barWidth := width - 2 // brackets
	if barWidth < 10 {
		barWidth = 10
	}

	filled := int(float64(barWidth) * p.percent / 100.0)
	if filled > barWidth {
		filled = barWidth
	}

	bar := "["
	bar += strings.Repeat("█", filled)
	bar += strings.Repeat("░", barWidth-filled)
	bar += "]"

	label := fmt.Sprintf(" %3.0f%% ", p.percent)
	if p.label != "" {
		label = " " + p.label + " " + label
	}
	bar += label

	return []string{bar}
}

func (p *ProgressBar) Invalidate() {}

// --- Selector Component ---

type Selector struct {
	items    []string
	selected int
	visible  int
	offset   int
	theme    Theme
	OnSelect func(index int, item string)
	OnCancel func()
}

func NewSelector(items []string, theme Theme) *Selector {
	return &Selector{
		items:   items,
		theme:   theme,
		visible: 10,
	}
}

func (s *Selector) SetVisible(n int) { s.visible = n }

func (s *Selector) MoveUp() {
	if s.selected > 0 {
		s.selected--
		if s.selected < s.offset {
			s.offset = s.selected
		}
	}
}

func (s *Selector) MoveDown() {
	if s.selected < len(s.items)-1 {
		s.selected++
		if s.selected >= s.offset+s.visible {
			s.offset = s.selected - s.visible + 1
		}
	}
}

func (s *Selector) Select() {
	if s.OnSelect != nil {
		s.OnSelect(s.selected, s.items[s.selected])
	}
}

func (s *Selector) Cancel() {
	if s.OnCancel != nil {
		s.OnCancel()
	}
}

func (s *Selector) Render(width int) []string {
	if width <= 0 {
		return nil
	}

	var lines []string
	end := s.offset + s.visible
	if end > len(s.items) {
		end = len(s.items)
	}

	for i := s.offset; i < end; i++ {
		prefix := "  "
		item := s.items[i]

		if i == s.selected {
			prefix = "▸ "
			item = s.theme.Selected(item)
		}

		line := prefix + truncateToWidth(item, width-2)
		lines = append(lines, line)
	}

	if end < len(s.items) {
		lines = append(lines, s.theme.Dim(fmt.Sprintf("  ... e mais %d itens", len(s.items)-end)))
	}

	return lines
}

func (s *Selector) Invalidate() {}

// --- Status Bar Component ---

type StatusBar struct {
	left   string
	right  string
	theme  Theme
}

func NewStatusBar(theme Theme) *StatusBar {
	return &StatusBar{theme: theme}
}

func (s *StatusBar) SetLeft(text string)  { s.left = text }
func (s *StatusBar) SetRight(text string) { s.right = text }

func (s *StatusBar) Render(width int) []string {
	if width <= 0 {
		return nil
	}

	left := truncateToWidth(s.left, width/2)
	right := truncateToWidth(s.right, width/2)

	padding := width - visibleWidth(left) - visibleWidth(right)
	if padding < 1 {
		padding = 1
	}

	line := s.theme.Dim(left) + strings.Repeat(" ", padding) + s.theme.Dim(right)
	return []string{line}
}

func (s *StatusBar) Invalidate() {}

// --- Spinner Component ---

type AnimatedSpinner struct {
	frames   []string
	label    string
	start    time.Time
	interval time.Duration
	tick     int
}

func NewAnimatedSpinner(label string) *AnimatedSpinner {
	return &AnimatedSpinner{
		frames:   []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
		label:    label,
		start:    time.Now(),
		interval: 100 * time.Millisecond,
	}
}

func (s *AnimatedSpinner) Render(width int) []string {
	elapsed := time.Since(s.start)
	frameIdx := int(elapsed/s.interval) % len(s.frames)
	frame := s.frames[frameIdx]

	line := fmt.Sprintf("\033[38;5;99m%s\033[0m \033[1m%s\033[0m \033[90m(%.1fs)\033[0m",
		frame, s.label, elapsed.Seconds())

	return []string{truncateToWidth(line, width)}
}

func (s *AnimatedSpinner) Invalidate() {}

// --- Error Card Component ---

type ErrorCard struct {
	title    string
	message  string
	suggestions []string
	theme    Theme
}

func NewErrorCard(title, message string, suggestions []string, theme Theme) *ErrorCard {
	return &ErrorCard{
		title:       title,
		message:     message,
		suggestions: suggestions,
		theme:       theme,
	}
}

func (e *ErrorCard) Render(width int) []string {
	var lines []string

	// Top border
	lines = append(lines, e.theme.Error(strings.Repeat("─", width)))

	// Title
	title := fmt.Sprintf(" ❌ %s ", e.title)
	lines = append(lines, e.theme.Error(title))

	// Separator
	lines = append(lines, e.theme.Border("│"))

	// Message (word-wrapped)
	for _, line := range wordWrap(e.message, width-4) {
		lines = append(lines, "  "+line)
	}

	// Suggestions
	if len(e.suggestions) > 0 {
		lines = append(lines, e.theme.Border("│"))
		lines = append(lines, e.theme.Accent(" 💡 Sugestões:"))
		for _, s := range e.suggestions {
			lines = append(lines, "  • "+s)
		}
	}

	// Bottom border
	lines = append(lines, e.theme.Error(strings.Repeat("─", width)))

	return lines
}

func (e *ErrorCard) Invalidate() {}

// --- Helpers ---

func wordWrap(text string, maxWidth int) []string {
	if maxWidth <= 0 {
		return []string{text}
	}

	var lines []string
	words := strings.Fields(text)
	if len(words) == 0 {
		return nil
	}

	current := words[0]
	for _, word := range words[1:] {
		if utf8.RuneCountInString(current)+1+utf8.RuneCountInString(word) <= maxWidth {
			current += " " + word
		} else {
			lines = append(lines, current)
			current = word
		}
	}
	lines = append(lines, current)

	return lines
}

func visibleWidth(s string) int {
	// Strip ANSI codes and count visible runes
	var stripped strings.Builder
	inEscape := false
	for _, r := range s {
		if r == '\033' {
			inEscape = true
		} else if inEscape {
			if r == 'm' {
				inEscape = false
			}
		} else {
			stripped.WriteRune(r)
		}
	}
	return utf8.RuneCountInString(stripped.String())
}

func truncateToWidth(s string, maxWidth int) string {
	if maxWidth <= 0 {
		return ""
	}
	// Simple truncation by rune count (doesn't handle ANSI perfectly but works for basic cases)
	if utf8.RuneCountInString(s) <= maxWidth {
		return s
	}
	return string([]rune(s)[:maxWidth]) + "…"
}
