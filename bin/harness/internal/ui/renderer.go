// Package ui provides terminal rendering utilities for the Harness agent.
package ui

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"harness/internal/session"
)

// Renderer handles terminal output formatting.
type Renderer struct {
	Theme  Theme
	mu     sync.Mutex // protects concurrent access to terminal output
}

// NewRenderer creates a new terminal renderer.
func NewRenderer() *Renderer {
	return &Renderer{Theme: DefaultTheme()}
}

// safeRepeat wraps strings.Repeat to prevent panic on negative count.
func safeRepeat(s string, count int) string {
	if count < 0 {
		count = 0
	}
	if count > 200 {
		count = 200 // safety limit
	}
	return strings.Repeat(s, count)
}

// WelcomeBoard displays the session welcome message using simple prints (no components).
func (r *Renderer) WelcomeBoard(providerName, sessionID, task, sessionFile string, availSkills, activeSkills int) {
	dispFile := sessionFile
	if len(dispFile) > 44 {
		dispFile = "..." + dispFile[len(dispFile)-41:]
	}
	dispTask := task
	if len(dispTask) > 38 {
		dispTask = dispTask[:35] + "..."
	}

	w := 76
	line := r.Theme.Accent("┌" + safeRepeat("─", w-2) + "┐")

	fmt.Println()
	fmt.Println(line)
	fmt.Printf("%s   🚀  %-55s  %s\n",
		r.Theme.Accent("│"), r.Theme.Accent("HARNESS AI AGENT (Go Edition)"), r.Theme.Accent("│"))
	fmt.Println(r.Theme.Accent("├" + safeRepeat("─", w-2) + "┤"))
	fmt.Printf("%s   Active Provider  : %-44s  %s\n", r.Theme.Accent("│"), r.Theme.Success(strings.ToUpper(providerName)), r.Theme.Accent("│"))
	fmt.Printf("%s   Session ID       : %-44s  %s\n", r.Theme.Accent("│"), r.Theme.Accent(sessionID), r.Theme.Accent("│"))
	fmt.Printf("%s   Root Task        : \"%-41s\"  %s\n", r.Theme.Accent("│"), dispTask, r.Theme.Accent("│"))
	fmt.Printf("%s   Skills Hub       : %d disponiveis | %d ativas       %s\n", r.Theme.Accent("│"), availSkills, activeSkills, r.Theme.Accent("│"))
	fmt.Printf("%s   History Log      : %-44s  %s\n", r.Theme.Accent("│"), r.Theme.Dim(dispFile), r.Theme.Accent("│"))
	fmt.Println(r.Theme.Accent("├" + safeRepeat("─", w-2) + "┤"))
	fmt.Printf("%s   Digite /help para comandos | /exit para sair       %s\n", r.Theme.Accent("│"), r.Theme.Accent("│"))
	fmt.Println(r.Theme.Accent("└" + safeRepeat("─", w-2) + "┘"))
}

// AssistantCard prints an assistant response in a formatted box.
// This is designed to be safe to call from any goroutine.
func (r *Renderer) AssistantCard(tree *session.Tree, provider string, nodeID string, content string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	lines := strings.Split(content, "\n")
	providerUpper := strings.ToUpper(provider)

	// Header
	header := fmt.Sprintf("🤖 ASSISTENTE (%s)", providerUpper)
	headerLen := len(header) + 6 // emoji + spaces + parens
	headerLine := fmt.Sprintf("%s%s%s",
		r.Theme.Accent("┌── "),
		r.Theme.Success(header),
		r.Theme.Accent(safeRepeat("─", 55-headerLen)))
	fmt.Println()
	fmt.Println(headerLine)

	// Content
	for _, line := range lines {
		fmt.Printf("%s  %s\n", r.Theme.Accent("│"), line)
	}

	// Footer
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	nodeTokens := 0
	if n, ok := tree.Nodes[nodeID]; ok {
		nodeTokens = n.Tokens
	}
	totalTokens := tree.TotalTokens()
	cost := tree.EstimateUSD(provider)
	costStr := fmt.Sprintf("$%.4f USD", cost)
	if cost == 0.0 {
		costStr = "Free/Local"
	}

	footerInfo := fmt.Sprintf("no: %s | 🪙 %dt (total: %dt) | 💰 %s | 🕒 %s",
		nodeID, nodeTokens, totalTokens, costStr, timestamp)
	footerWidth := 75 - len(footerInfo)
	if footerWidth < 1 {
		footerWidth = 1
	}
	fmt.Printf("%s%s%s\n",
		r.Theme.Accent("└── "),
		r.Theme.Dim(fmt.Sprintf("(%s)", footerInfo)),
		r.Theme.Accent(safeRepeat("─", footerWidth)))
}

// PrintTree renders the session tree.
func (r *Renderer) PrintTree(tree *session.Tree, provider string, rootPath string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var roots []session.Node
	for _, n := range tree.Nodes {
		if n.ParentID == "" {
			roots = append(roots, n)
		}
	}

	totalNodes := len(tree.Nodes)
	totalTokens := tree.TotalTokens()
	cost := tree.EstimateUSD(provider)
	availSkills, _ := session.GetSkillsMetrics(rootPath)

	costStr := fmt.Sprintf("$%.4f USD", cost)
	if cost == 0.0 {
		costStr = "$0.00 (Local/Free)"
	}

	fmt.Printf("\n%s\n", r.Theme.Accent("┌── 🌳 ARVORE DE DECISAO ──"))
	fmt.Printf("%s  🆔 ID: %s\n", r.Theme.Accent("│"), r.Theme.Accent(tree.SessionID))
	fmt.Printf("%s  🎯 Objetivo: \"%s\"\n", r.Theme.Accent("│"), tree.RootTask)
	fmt.Printf("%s  📊 %d nos | 🪙 %d tokens | 💰 %s\n", r.Theme.Accent("│"), totalNodes, totalTokens, costStr)
	fmt.Printf("%s  🎓 %d skills disponiveis\n", r.Theme.Accent("│"), availSkills)
	fmt.Printf("%s\n", r.Theme.Accent("├── 🌲 NOS ──"))

	for _, root := range roots {
		r.renderNode(tree, root.NodeID, "", true)
	}
	fmt.Printf("%s\n", r.Theme.Accent("└── ──"))
}

func (r *Renderer) renderNode(tree *session.Tree, nodeID string, indent string, isLast bool) {
	node := tree.Nodes[nodeID]

	marker := "├── "
	nextIndent := indent + "│   "
	if isLast {
		marker = "└── "
		nextIndent = indent + "    "
	}

	activeMarker := ""
	if nodeID == tree.ActiveNode {
		activeMarker = r.Theme.Warning(" ★ ATIVO")
	}

	roleIcon := "⚙️"
	roleColor := "\033[38;5;86m"
	roleLabel := "System"
	switch node.Role {
	case "user":
		roleIcon = "👤"
		roleColor = "\033[38;5;120m"
		roleLabel = "User"
	case "assistant":
		roleIcon = "🤖"
		roleColor = "\033[38;5;99m"
		roleLabel = "Assistant"
	}

	timeStr := ""
	tParsed, err := time.Parse(time.RFC3339, node.Timestamp)
	if err == nil {
		timeStr = tParsed.Format("15:04:05")
	} else if len(node.Timestamp) >= 19 {
		timeStr = node.Timestamp[11:19]
	} else {
		timeStr = node.Timestamp
	}

	tokenInfo := ""
	if node.Tokens > 0 {
		tokenInfo = fmt.Sprintf(" (🪙 %dt)", node.Tokens)
	}

	preview := node.Content
	preview = strings.ReplaceAll(preview, "\n", " ")
	if len(preview) > 55 {
		preview = preview[:52] + "..."
	}

	fmt.Printf("%s%s[%s%s]%s %s%s%s - %s%s%s\n",
		indent, marker, roleColor, node.NodeID, activeMarker,
		roleColor, roleIcon, roleLabel, timeStr,
		r.Theme.Dim(tokenInfo), "\033[0m")

	var children []session.Node
	for _, n := range tree.Nodes {
		if n.ParentID == nodeID {
			children = append(children, n)
		}
	}

	for i, child := range children {
		r.renderNode(tree, child.NodeID, nextIndent, i == len(children)-1)
	}
}

// ShowError displays a structured error message.
func (r *Renderer) ShowError(err error, providerName string) {
	info := ClassifyError(err, providerName)
	fmt.Println(FormatErrorForDisplay(info))
}

// ShowRecovery displays a recovery message.
func (r *Renderer) ShowRecovery(info ErrorInfo, attempt int, maxAttempts int) {
	msg := fmt.Sprintf("Tentando novamente (%d/%d)...", attempt+1, maxAttempts)
	switch info.Category {
	case CategoryRateLimit:
		fmt.Printf("\r%s %s\n", r.Theme.Warning("⏳ Rate-limit,"), r.Theme.Muted(msg))
	case CategoryNetwork:
		fmt.Printf("\r%s %s\n", r.Theme.Warning("🔄 Rede,"), r.Theme.Muted(msg))
	case CategoryContext:
		fmt.Printf("\r%s %s\n", r.Theme.Warning("⏰ Timeout,"), r.Theme.Muted(msg))
	case CategoryServer:
		fmt.Printf("\r%s %s\n", r.Theme.Warning("🔧 Servidor,"), r.Theme.Muted(msg))
	default:
		fmt.Printf("\r%s %s\n", r.Theme.Warning("⚠️"), r.Theme.Muted(msg))
	}
}

// ShowFallback displays a provider fallback message.
func (r *Renderer) ShowFallback(fromProvider, toProvider string) {
	fmt.Printf("%s %s → %s\n",
		r.Theme.Warning("⚡ Fallback automatico"),
		strings.ToUpper(fromProvider),
		strings.ToUpper(toProvider))
}

// Spinner represents a terminal spinner animation.
type Spinner struct {
	stopChan chan struct{}
}

// StartSpinner starts an animated spinner with a suffix label.
func StartSpinner(suffix string) *Spinner {
	s := &Spinner{stopChan: make(chan struct{})}
	go func() {
		frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		i := 0
		start := time.Now()
		for {
			select {
			case <-s.stopChan:
				fmt.Print("\r\033[K")
				return
			default:
				elapsed := time.Since(start).Seconds()
				fmt.Printf("\r\033[38;5;99m%s\033[0m \033[1m%s\033[0m \033[90m(%.1fs)\033[0m",
					frames[i%len(frames)], suffix, elapsed)
				i++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	return s
}

// Stop stops the spinner.
func (s *Spinner) Stop() {
	close(s.stopChan)
}
