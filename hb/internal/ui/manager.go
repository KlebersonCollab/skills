package ui

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"golang.org/x/term"
	"github.com/klebersonromero/hb/internal/ui/layout"
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

	width, height, _ := TerminalSize()
	if width < 10 {
		width = 80
	}
	if height < 5 {
		height = 5
	}

	m.clear()

	// 1. Define Layout Tree
	root := layout.NewNode(layout.Style{
		Direction: layout.Column,
		Width:     width,
		Height:    4, // Status bar height
	})

	header := layout.NewNode(layout.Style{
		Direction: layout.Row,
		Justify:   layout.JustifySpaceBetween,
		Padding:   layout.EdgeValues{Left: 1, Right: 1},
		FlexGrow:  1,
	})

	taskRow := layout.NewNode(layout.Style{
		Direction: layout.Row,
		Padding:   layout.EdgeValues{Left: 1, Right: 1},
		FlexGrow:  1,
	})

	body := layout.NewNode(layout.Style{
		Direction: layout.Row,
		Padding:   layout.EdgeValues{Left: 1, Right: 1},
		FlexGrow:  1,
	})

	root.AddChild(header)
	root.AddChild(taskRow)
	root.AddChild(body)

	// 2. Calculate Layout
	root.Calculate(width, 4)

	// 3. Render to Virtual Screen
	screen := layout.NewScreen(width, 4)
	
	// Draw Header (Feature | Status)
	feature := state.FeatureName
	if feature == "" {
		feature = "IDLE"
	}
	screen.DrawText(header.Result.X, header.Result.Y, fmt.Sprintf("[%s]", feature))
	screen.DrawText(header.Result.X+header.Result.Width-len(state.Status)-2, header.Result.Y, state.Status)

	// Draw Task
	if state.CurrentTask != "" {
		taskStr := truncate(state.CurrentTask, width-2)
		screen.DrawText(taskRow.Result.X, taskRow.Result.Y, taskStr)
	}

	// Draw Body (Progress Bar + Cost)
	barWidth := 20
	filled := int(state.Progress * float64(barWidth))
	bar := ""
	for i := 0; i < barWidth; i++ {
		if i < filled {
			bar += "■"
		} else {
			bar += "□"
		}
	}
	
	barText := fmt.Sprintf("%s %d%%", bar, int(state.Progress*100))
	screen.DrawText(body.Result.X, body.Result.Y, barText)
	
	costInfo := fmt.Sprintf("$%.2f | %d tkn", state.Cost, state.TokenCount)
	screen.DrawText(body.Result.X+body.Result.Width-len(costInfo)-1, body.Result.Y, costInfo)

	// 4. Output to terminal
	fmt.Fprint(m.output, SaveCursor)
	// We want to draw at the bottom of the screen
	fmt.Fprint(m.output, fmt.Sprintf("\033[%d;1H", height-3)) // Move to line height-3
	fmt.Fprint(m.output, screen.String())
	fmt.Fprint(m.output, RestoreCursor)
	
	m.lastLines = 4
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
