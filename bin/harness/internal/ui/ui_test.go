package ui

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestNewText(t *testing.T) {
	text := NewText("hello", 1, 1)
	if text == nil {
		t.Fatal("expected non-nil Text")
	}
	text.SetText("world")
	text.SetBgFn(func(s string) string { return s })
}

func TestTextRender(t *testing.T) {
	text := NewText("hello world", 0, 0)
	lines := text.Render(80)
	if len(lines) == 0 {
		t.Fatal("expected at least one line")
	}
	if !containsAny(lines, "hello") {
		t.Errorf("expected 'hello' in rendered output, got %v", lines)
	}
}

func TestTextRender_Width(t *testing.T) {
	text := NewText("hello world this is a test", 0, 0)
	lines := text.Render(10)
	if len(lines) == 0 {
		t.Fatal("expected lines for narrow width")
	}
}

func TestSpacer(t *testing.T) {
	s := NewSpacer(3)
	lines := s.Render(80)
	if len(lines) != 3 {
		t.Errorf("expected 3 lines, got %d", len(lines))
	}
}

func TestBox(t *testing.T) {
	b := NewBox(1, 1, nil)
	b.AddChild(NewText("inside", 0, 0))
	lines := b.Render(80)
	if len(lines) == 0 {
		t.Fatal("expected lines from box")
	}
}

func TestBox_Clear(t *testing.T) {
	b := NewBox(1, 1, nil)
	b.AddChild(NewText("a", 0, 0))
	b.Clear()
	lines := b.Render(80)
	_ = lines // should not panic
}

func TestContainer(t *testing.T) {
	c := NewContainer()
	c.AddChild(NewText("line1", 0, 0))
	c.AddChild(NewText("line2", 0, 0))
	lines := c.Render(80)
	if len(lines) != 2 {
		t.Errorf("expected 2 lines, got %d", len(lines))
	}
}

func TestContainer_Clear(t *testing.T) {
	c := NewContainer()
	c.AddChild(NewText("a", 0, 0))
	c.Clear()
	lines := c.Render(80)
	if len(lines) != 0 {
		t.Errorf("expected 0 lines after clear, got %d", len(lines))
	}
}

func TestBorder(t *testing.T) {
	b := NewBorder("Test Title", DefaultTheme())
	lines := b.Render(80)
	if len(lines) != 2 {
		t.Errorf("expected 2 lines (top+bottom), got %d", len(lines))
	}
	if !strings.Contains(lines[0], "Test Title") {
		t.Errorf("expected title in border, got %q", lines[0])
	}
}

func TestBorder_Narrow(t *testing.T) {
	b := NewBorder("X", DefaultTheme())
	lines := b.Render(4)
	_ = lines // just ensure no panic
}

func TestProgressBar(t *testing.T) {
	p := NewProgressBar(DefaultTheme())
	p.SetProgress(50.0, "loading")
	lines := p.Render(50)
	if len(lines) == 0 {
		t.Fatal("expected progress bar line")
	}
	if !strings.Contains(lines[0], "50%") {
		t.Errorf("expected 50%%, got %q", lines[0])
	}
}

func TestProgressBar_Zero(t *testing.T) {
	p := NewProgressBar(DefaultTheme())
	p.SetProgress(0, "")
	lines := p.Render(30)
	if len(lines) == 0 {
		t.Fatal("expected line even for 0%")
	}
}

func TestProgressBar_OneHundred(t *testing.T) {
	p := NewProgressBar(DefaultTheme())
	p.SetProgress(100, "done")
	lines := p.Render(30)
	if !strings.Contains(lines[0], "100%") {
		t.Errorf("expected 100%%, got %q", lines[0])
	}
}

func TestSelector(t *testing.T) {
	items := []string{"opcao A", "opcao B", "opcao C"}
	s := NewSelector(items, DefaultTheme())
	lines := s.Render(80)
	if len(lines) == 0 {
		t.Fatal("expected selector lines")
	}
}

func TestSelector_Navigation(t *testing.T) {
	items := []string{"a", "b", "c"}
	s := NewSelector(items, DefaultTheme())

	if s.selected != 0 {
		t.Errorf("expected selected=0, got %d", s.selected)
	}

	s.MoveDown()
	if s.selected != 1 {
		t.Errorf("expected selected=1 after MoveDown, got %d", s.selected)
	}

	s.MoveDown()
	if s.selected != 2 {
		t.Errorf("expected selected=2, got %d", s.selected)
	}

	s.MoveDown() // should stay at max
	if s.selected != 2 {
		t.Errorf("expected selected=2 at boundary, got %d", s.selected)
	}

	s.MoveUp()
	if s.selected != 1 {
		t.Errorf("expected selected=1 after MoveUp, got %d", s.selected)
	}
}

func TestSelector_Callbacks(t *testing.T) {
	items := []string{"a", "b"}
	s := NewSelector(items, DefaultTheme())

	selected := ""
	s.OnSelect = func(idx int, item string) {
		selected = item
	}
	s.Select()
	if selected != "a" {
		t.Errorf("expected 'a', got %q", selected)
	}

	cancelled := false
	s.OnCancel = func() { cancelled = true }
	s.Cancel()
	if !cancelled {
		t.Error("expected cancel callback")
	}
}

func TestStatusBar(t *testing.T) {
	s := NewStatusBar(DefaultTheme())
	s.SetLeft("left text")
	s.SetRight("right text")
	lines := s.Render(80)
	if len(lines) == 0 {
		t.Fatal("expected status bar line")
	}
}

func TestAnimatedSpinner(t *testing.T) {
	s := NewAnimatedSpinner("loading")
	lines := s.Render(80)
	if len(lines) == 0 {
		t.Fatal("expected spinner line")
	}
}

func TestErrorCard(t *testing.T) {
	suggestions := []string{"tente de novo", "verifique a config"}
	e := NewErrorCard("Erro 500", "servidor falhou", suggestions, DefaultTheme())
	lines := e.Render(60)
	if len(lines) == 0 {
		t.Fatal("expected error card lines")
	}
}

func TestDefaultTheme(t *testing.T) {
	theme := DefaultTheme()
	if theme.Accent("x") == "x" {
		t.Error("expected Accent to add ANSI codes")
	}
	if theme.Success("x") == "x" {
		t.Error("expected Success to add ANSI codes")
	}
	if theme.Error("x") == "x" {
		t.Error("expected Error to add ANSI codes")
	}
}

func TestWordWrap(t *testing.T) {
	lines := wordWrap("hello world this is a test", 10)
	if len(lines) == 0 {
		t.Fatal("expected wrapped lines")
	}
}

func TestWordWrap_Empty(t *testing.T) {
	lines := wordWrap("", 10)
	if lines != nil {
		t.Errorf("expected nil for empty text, got %v", lines)
	}
}

func TestWordWrap_Narrow(t *testing.T) {
	lines := wordWrap("hello world", 3)
	if len(lines) == 0 {
		t.Fatal("expected lines for narrow width")
	}
}

func TestVisibleWidth(t *testing.T) {
	noANSI := visibleWidth("hello")
	if noANSI != 5 {
		t.Errorf("expected 5, got %d", noANSI)
	}

	withANSI := visibleWidth("\033[31mhello\033[0m")
	if withANSI != 5 {
		t.Errorf("expected 5 for ANSI-wrapped text, got %d", withANSI)
	}
}

func TestTruncateToWidth(t *testing.T) {
	if truncateToWidth("hello", 10) != "hello" {
		t.Error("short string should not truncate")
	}
	result := truncateToWidth("hello world this is long", 10)
	if len([]rune(result)) > 12 {
		t.Errorf("expected truncated string, got %q (len %d)", result, len([]rune(result)))
	}
}

func TestTruncateToWidth_Zero(t *testing.T) {
	if truncateToWidth("hello", 0) != "" {
		t.Error("expected empty for zero width")
	}
}

// --- Error Recovery Tests ---

func TestClassifyError_Auth(t *testing.T) {
	info := ClassifyError(fmt.Errorf("401 Unauthorized"), "gemini")
	if info.Category != CategoryAuth {
		t.Errorf("expected Auth, got %v", info.Category)
	}
	if info.Recoverable {
		t.Error("auth errors should not be recoverable")
	}
}

func TestClassifyError_RateLimit(t *testing.T) {
	info := ClassifyError(fmt.Errorf("429 Too Many Requests"), "deepseek")
	if info.Category != CategoryRateLimit {
		t.Errorf("expected RateLimit, got %v", info.Category)
	}
	if !info.Recoverable {
		t.Error("rate limit errors should be recoverable")
	}
}

func TestClassifyError_Network(t *testing.T) {
	info := ClassifyError(fmt.Errorf("connection refused"), "ollama")
	if info.Category != CategoryNetwork {
		t.Errorf("expected Network, got %v", info.Category)
	}
}

func TestClassifyError_Timeout(t *testing.T) {
	info := ClassifyError(fmt.Errorf("context deadline exceeded"), "gemini")
	if info.Category != CategoryContext {
		t.Errorf("expected Context, got %v", info.Category)
	}
}

func TestClassifyError_APIKey(t *testing.T) {
	info := ClassifyError(fmt.Errorf("API key not set"), "gemini")
	if info.Category != CategoryAuth {
		t.Errorf("expected Auth for missing key, got %v", info.Category)
	}
}

func TestClassifyError_Nil(t *testing.T) {
	info := ClassifyError(nil, "gemini")
	if !info.Recoverable {
		t.Error("nil error should be recoverable")
	}
}

func TestDecideRecovery_Abort(t *testing.T) {
	info := ErrorInfo{Recoverable: false, Category: CategoryAuth}
	action, _ := DecideRecovery(info, 0, 3)
	if action != ActionAbort {
		t.Errorf("expected Abort for non-recoverable, got %v", action)
	}
}

func TestDecideRecovery_Retry(t *testing.T) {
	info := ErrorInfo{Recoverable: true, Category: CategoryNetwork, RetryAfter: time.Second}
	action, delay := DecideRecovery(info, 0, 3)
	if action != ActionRetry {
		t.Errorf("expected Retry, got %v", action)
	}
	if delay <= 0 {
		t.Errorf("expected positive delay, got %v", delay)
	}
}

func TestDecideRecovery_Fallback(t *testing.T) {
	info := ErrorInfo{Recoverable: true, Category: CategoryUnknown}
	action, _ := DecideRecovery(info, 3, 3)
	if action != ActionFallback {
		t.Errorf("expected Fallback after max attempts, got %v", action)
	}
}

func TestFormatErrorForDisplay(t *testing.T) {
	info := ErrorInfo{
		Message:     "test error",
		Category:    CategoryNetwork,
		Recoverable: true,
		Suggestions: []string{"check connection"},
		ProviderName: "gemini",
	}
	display := FormatErrorForDisplay(info)
	if !strings.Contains(display, "test error") {
		t.Errorf("expected error message in display, got %q", display)
	}
	if !strings.Contains(display, "GEMINI") {
		t.Errorf("expected provider name in display, got %q", display)
	}
	if !strings.Contains(display, "check connection") {
		t.Errorf("expected suggestion in display, got %q", display)
	}
}

func TestCategoryString(t *testing.T) {
	if CategoryNetwork.String() != "network" {
		t.Errorf("expected 'network', got %q", CategoryNetwork.String())
	}
	if CategoryAuth.String() != "auth" {
		t.Errorf("expected 'auth', got %q", CategoryAuth.String())
	}
	if Category(-1).String() != "unknown" {
		t.Errorf("expected 'unknown', got %q", Category(-1).String())
	}
}

// Helper

func containsAny(lines []string, substr string) bool {
	for _, l := range lines {
		if strings.Contains(l, substr) {
			return true
		}
	}
	return false
}
