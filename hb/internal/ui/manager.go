package ui

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"golang.org/x/term"
)

// UIState represents the current data to be displayed in the UI
type UIState struct {
	FeatureName string
	CurrentTask string
	Progress    float64 // 0.0 to 1.0
	TokenCount  int
	Cost        float64
	Status      string
}

// UIManager handles the terminal UI lifecycle and rendering
type UIManager struct {
	state      UIState
	mu         sync.RWMutex
	output     io.Writer
	ticker     *time.Ticker
	stopChan   chan struct{}
	lastLines  int
	isTerminal bool
}

// NewUIManager creates a new UI manager
func NewUIManager(out io.Writer) *UIManager {
	return &UIManager{
		output:     out,
		stopChan:   make(chan struct{}),
		isTerminal: isTerminal(out),
	}
}

// SetState updates the UI state
func (m *UIManager) SetState(state UIState) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.state = state
}

// Start begins the rendering loop
func (m *UIManager) Start() {
	if !m.isTerminal {
		return
	}

	fmt.Fprint(m.output, HideCursor)
	m.ticker = time.NewTicker(100 * time.Millisecond)
	go m.renderLoop()
}

// Stop ends the rendering loop and cleans up the terminal
func (m *UIManager) Stop() {
	if !m.isTerminal {
		return
	}

	if m.ticker != nil {
		m.ticker.Stop()
	}
	close(m.stopChan)
	m.clear()
	fmt.Fprint(m.output, ShowCursor)
}

func (m *UIManager) renderLoop() {
	for {
		select {
		case <-m.ticker.C:
			m.render()
		case <-m.stopChan:
			return
		}
	}
}

func (m *UIManager) render() {
	m.mu.RLock()
	state := m.state
	m.mu.RUnlock()

	width, _, _ := TerminalSize()
	if width < 10 {
		width = 80
	}

	m.clear()

	// Build the status bar
	// Pattern: [FEATURE] Task Description... [■■■■□] 80% | $0.05
	barWidth := 20
	filled := int(state.Progress * float64(barWidth))
	progressLine := ""
	for i := 0; i < barWidth; i++ {
		if i < filled {
			progressLine += "■"
		} else {
			progressLine += "□"
		}
	}

	feature := Colorize(fmt.Sprintf("[%s]", state.FeatureName), Cyan)
	if state.FeatureName == "" {
		feature = Colorize("[IDLE]", White)
	}

	statusLine := fmt.Sprintf("%s %s %s %d%% | Tokens: %d | Cost: $%.2f",
		feature,
		truncate(state.CurrentTask, width-50),
		Colorize(progressLine, Green),
		int(state.Progress*100),
		state.TokenCount,
		state.Cost,
	)

	// Inject at bottom
	fmt.Fprint(m.output, SaveCursor)
	fmt.Fprint(m.output, statusLine)
	fmt.Fprint(m.output, RestoreCursor)
	
	m.lastLines = 1 // Track how many lines we occupied
}

func (m *UIManager) clear() {
	if m.lastLines > 0 {
		// In this simple ANSI injection, we just clear the line if we stay on the same row
		// or we can use more advanced cursor movement.
		fmt.Fprint(m.output, "\r"+ClearLine)
	}
}

func isTerminal(w io.Writer) bool {
	if f, ok := w.(*os.File); ok {
		return term.IsTerminal(int(f.Fd()))
	}
	return false
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen < 3 {
		return "..."
	}
	return s[:maxLen-3] + "..."
}
