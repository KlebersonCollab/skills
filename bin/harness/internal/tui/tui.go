// Package tui provides a terminal UI similar to pi.dev for the Harness agent.
package tui

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"harness/internal/session"
)

const (
	CReset = "\033[0m"
	CBold  = "\033[1m"
	CDim   = "\033[2m"

	CAccent   = "\033[38;5;99m"
	CGreen    = "\033[1;32m"
	CRed      = "\033[38;5;196m"
	CYellow   = "\033[38;5;208m"
	CBlue     = "\033[38;5;39m"
	CCyan     = "\033[38;5;86m"
	CWhite    = "\033[38;5;15m"
	CGray     = "\033[38;5;248m"
	CDarkGray = "\033[90m"
)

type ThinkingLevel int

const (
	ThinkingOff ThinkingLevel = iota
	ThinkingMinimal
	ThinkingLow
	ThinkingMedium
	ThinkingHigh
	ThinkingXHigh
)

func (t ThinkingLevel) BorderColor() string {
	switch t {
	case ThinkingOff:
		return CDarkGray
	case ThinkingMinimal:
		return CBlue
	case ThinkingLow:
		return CCyan
	case ThinkingMedium:
		return CGreen
	case ThinkingHigh:
		return CYellow
	case ThinkingXHigh:
		return CRed
	default:
		return CAccent
	}
}

func (t ThinkingLevel) String() string {
	switch t {
	case ThinkingOff:
		return "off"
	case ThinkingMinimal:
		return "minimal"
	case ThinkingLow:
		return "low"
	case ThinkingMedium:
		return "medium"
	case ThinkingHigh:
		return "high"
	case ThinkingXHigh:
		return "xhigh"
	default:
		return "off"
	}
}

type TUI struct {
	mu            sync.Mutex
	thinkingLevel ThinkingLevel
	providerName  string
	sessionID     string
}

func New() *TUI {
	return &TUI{thinkingLevel: ThinkingOff}
}

func safeRepeat(s string, count int) string {
	if count < 0 {
		count = 0
	}
	if count > 300 {
		count = 300
	}
	return strings.Repeat(s, count)
}

func visibleLen(s string) int {
	n := 0
	in := false
	for _, r := range s {
		if r == '\033' {
			in = true
		} else if in {
			if r == 'm' {
				in = false
			}
		} else {
			n++
		}
	}
	return n
}

func padRight(s string, w int) string {
	v := visibleLen(s)
	if v >= w {
		return s
	}
	return s + safeRepeat(" ", w-v)
}

// WelcomeHeader prints the startup info (no box borders, token-efficient).
func (t *TUI) WelcomeHeader(providerName, modelName, sessionID, task string, availSkills int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.providerName = providerName
	t.sessionID = sessionID

	fmt.Println()
	fmt.Printf("%s%s %s%s\n", CAccent, CBold, "HARNESS AI AGENT", CReset)
	fmt.Printf("%s  Provider: %s%s\n", CGreen, CReset, strings.ToUpper(providerName))
	if modelName != "" {
		fmt.Printf("%s  Model:    %s%s\n", CYellow, CReset, modelName)
	}
	fmt.Printf("%s  Session:  %s%s\n", CAccent, CReset, sessionID)
	fmt.Printf("%s  Task:     %s\"%s\"%s\n", CGray, CReset, task, CReset)
	fmt.Printf("%s  Skills:   %s%d disponiveis%s\n", CCyan, CReset, availSkills, CReset)
	fmt.Printf("%s  %s/help for commands  |  /exit to quit%s\n", CDarkGray, CDim, CReset)
}

// MessageBubble prints a formatted message without border decorations (token-efficient).
func (t *TUI) MessageBubble(role, content string, tokens int) {
	t.mu.Lock()
	defer t.mu.Unlock()

	fmt.Println()

	for _, line := range strings.Split(content, "\n") {
		switch role {
		case "user":
			fmt.Printf("  %s\n", line)
		case "assistant":
			fmt.Printf("  %s\n", line)
		case "system":
			fmt.Printf("%s  ⚙️ %s%s\n", CDarkGray, CDim, line)
		}
	}

	if role == "assistant" {
		fmt.Printf("%s  ─ %d tokens ─%s\n", CDarkGray, tokens, CReset)
	}
}

// ShowThinking prints a processing indicator and returns a stop function.
func (t *TUI) ShowThinking(providerName string) func() {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	start := time.Now()
	done := make(chan struct{})
	statuses := make(chan string, 10)

	go func() {
		i := 0
		lastStatus := ""
		for {
			select {
			case <-done:
				fmt.Print("\r\033[K")
				return
			case s := <-statuses:
				lastStatus = s
			default:
				elapsed := time.Since(start).Seconds()
				p := strings.ToUpper(providerName)
				if lastStatus != "" {
					fmt.Printf("\r%s%s%s %s[%s] %s %.1fs%s",
						CAccent, frames[i%len(frames)], CReset,
						CDim, p, lastStatus, elapsed, CReset)
				} else {
					fmt.Printf("\r%s%s%s %s[%s] thinking... %.1fs%s",
						CAccent, frames[i%len(frames)], CReset,
						CDim, p, elapsed, CReset)
				}
				i++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

	return func() { close(done); fmt.Print("\r") }
}

// UpdateStatus sends a status update to the thinking spinner.
// Should only be called while ShowThinking's stop function hasn't been called.
func (t *TUI) UpdateStatus(status string) {
	fmt.Printf("\r%s\033[K", CDim+status+CReset) // overwrite spinner line
}

// PrintTree renders the session tree.
func (t *TUI) PrintTree(tree *session.Tree, rootPath string) {
	t.mu.Lock()
	defer t.mu.Unlock()

	fmt.Printf("\n%s🌳 Session Tree%s\n", CAccent, CReset)

	var roots []session.Node
	for _, n := range tree.Nodes {
		if n.ParentID == "" {
			roots = append(roots, n)
		}
	}

	fmt.Printf("%s  %d nodes | 🪙 %d tokens%s\n", CDim, len(tree.Nodes), tree.TotalTokens(), CReset)

	for _, root := range roots {
		t.printNode(tree, root.NodeID, "", true)
	}
}

func (t *TUI) printNode(tree *session.Tree, nodeID, indent string, isLast bool) {
	node := tree.Nodes[nodeID]

	marker := "├── "
	nextIndent := indent + "│   "
	if isLast {
		marker = "└── "
		nextIndent = indent + "    "
	}

	active := ""
	if nodeID == tree.ActiveNode {
		active = " " + CGreen + "⬤ active" + CReset
	}

	roleColor := CCyan
	roleIcon := "⚙"
	switch node.Role {
	case "user":
		roleColor = CGreen
		roleIcon = "👤"
	case "assistant":
		roleColor = CAccent
		roleIcon = "🤖"
	}

	timeStr := ""
	if len(node.Timestamp) >= 19 {
		timeStr = " " + node.Timestamp[11:19]
	}

	info := ""
	if node.Tokens > 0 {
		info = fmt.Sprintf(" %s(%dt)%s", CDarkGray, node.Tokens, CReset)
	}

	preview := strings.ReplaceAll(node.Content, "\n", " ")
	if len(preview) > 50 {
		preview = preview[:47] + "..."
	}

	fmt.Printf("%s%s%s[%s%s]%s %s%s%s%s%s%s\n",
		CDarkGray, indent, marker, roleColor, node.NodeID, active,
		roleColor, roleIcon, timeStr, info, CDim, preview)

	var children []session.Node
	for _, n := range tree.Nodes {
		if n.ParentID == nodeID {
			children = append(children, n)
		}
	}

	for i, child := range children {
		t.printNode(tree, child.NodeID, nextIndent, i == len(children)-1)
	}
}

// Footer prints the status bar.
func (t *TUI) Footer(tree *session.Tree, tokCost float64) {
	t.mu.Lock()
	defer t.mu.Unlock()

	w := 78
	totalTokens := 0
	if tree != nil {
		totalTokens = tree.TotalTokens()
	}

	left := "Ready"
	right := fmt.Sprintf("%s | 🪙 %dt | 💰 $%.4f", t.sessionID, totalTokens, tokCost)

	line := fmt.Sprintf("%s▐%s %s %s▌%s", CDarkGray, CReset, CDim+left+CReset, CDarkGray+right+CReset, CReset)
	fmt.Printf("\n%s\n", padRight(line, w))
}
