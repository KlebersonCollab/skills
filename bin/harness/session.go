package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Node struct {
	NodeID    string `json:"node_id"`
	ParentID  string `json:"parent_id"`
	Role      string `json:"role"` // system, user, assistant
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
	Tokens    int    `json:"tokens"`
}

type SessionTree struct {
	SessionID  string          `json:"session_id"`
	RootTask   string          `json:"root_task"`
	ActiveNode string          `json:"active_node"`
	Nodes      map[string]Node `json:"nodes"`
}

func NewSessionTree(sessionID string, rootTask string) *SessionTree {
	return &SessionTree{
		SessionID:  sessionID,
		RootTask:   rootTask,
		ActiveNode: "",
		Nodes:      make(map[string]Node),
	}
}

func LoadSessionTree(path string) (*SessionTree, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tree SessionTree
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tree); err != nil {
		return nil, err
	}
	return &tree, nil
}

func (t *SessionTree) Save(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(t)
}

func (t *SessionTree) AddNode(role string, content string, tokens int) string {
	nodeID := fmt.Sprintf("node-%d", len(t.Nodes))
	parentID := t.ActiveNode

	node := Node{
		NodeID:    nodeID,
		ParentID:  parentID,
		Role:      role,
		Content:   content,
		Timestamp: time.Now().Format(time.RFC3339),
		Tokens:    tokens,
	}

	t.Nodes[nodeID] = node
	t.ActiveNode = nodeID
	return nodeID
}

// GetLinearHistory traverses backwards from the active node to the root and reverses it to get chronological order
func (t *SessionTree) GetLinearHistory(startNodeID string) ([]Node, error) {
	var history []Node
	currentID := startNodeID

	if currentID == "" {
		currentID = t.ActiveNode
	}

	for currentID != "" {
		node, exists := t.Nodes[currentID]
		if !exists {
			return nil, fmt.Errorf("node '%s' not found in session tree", currentID)
		}
		history = append(history, node)
		currentID = node.ParentID
	}

	// Reverse the slice to put in chronological order
	for i, j := 0, len(history)-1; i < j; i, j = i+1, j-1 {
		history[i], history[j] = history[j], history[i]
	}

	return history, nil
}

func (t *SessionTree) FormatLinearHistoryText(startNodeID string) (string, error) {
	history, err := t.GetLinearHistory(startNodeID)
	if err != nil {
		return "", err
	}

	var builder string
	for _, node := range history {
		roleLabel := "User"
		switch node.Role {
		case "system":
			roleLabel = "System"
		case "assistant":
			roleLabel = "Assistant"
		}
		builder += fmt.Sprintf("[%s]: %s\n", roleLabel, node.Content)
	}
	return builder, nil
}

func (t *SessionTree) GetTotalTokens() int {
	totalTokens := 0
	for _, n := range t.Nodes {
		totalTokens += n.Tokens
	}
	return totalTokens
}

func (t *SessionTree) EstimateUSD(provider string) float64 {
	totalTokens := t.GetTotalTokens()
	ratePerToken := 0.0
	switch strings.ToLower(provider) {
	case "deepseek":
		ratePerToken = 0.000002 // Média de $2.00 por 1M de tokens
	case "gemini":
		ratePerToken = 0.0000035 // Média de $3.50 por 1M de tokens
	default:
		ratePerToken = 0.0 // Ollama local ou custom grátis
	}
	return float64(totalTokens) * ratePerToken
}

func GetSkillsMetrics(root string) (available int, active int) {
	skills := DiscoverSkills(root)
	// For now, we consider all discovered skills as available and active
	return len(skills), len(skills)
}

// PrintTree renders the DAG of sessions in a structured premium console box format
func (t *SessionTree) PrintTree(provider string, rootPath string) {
	var roots []Node
	for _, n := range t.Nodes {
		if n.ParentID == "" {
			roots = append(roots, n)
		}
	}

	sort.Slice(roots, func(i, j int) bool {
		return roots[i].Timestamp < roots[j].Timestamp
	})

	totalNodes := len(t.Nodes)
	totalTokens := t.GetTotalTokens()
	cost := t.EstimateUSD(provider)
	availSkills, activeSkills := GetSkillsMetrics(rootPath)

	costStr := fmt.Sprintf("$%.4f USD", cost)
	if cost == 0.0 {
		costStr = "$0.00 (Local/Free)"
	}

	fmt.Printf("\n\033[38;5;99m┌── 🌳 HISTÓRICO E DAG DE DECISÃO DA SESSÃO ──────────────────────────────────────────\033[0m\n")
	fmt.Printf("\033[38;5;99m│\033[0m  🆔 ID da Sessão : \033[35m%s\033[0m\n", t.SessionID)
	fmt.Printf("\033[38;5;99m│\033[0m  🎯 Objetivo     : \"%s\"\n", t.RootTask)
	fmt.Printf("\033[38;5;99m│\033[0m  📊 Estatísticas : \033[1m%d\033[0m nós | 🪙  \033[32m%d\033[0m tokens gastos | 💰 Est. Custo: \033[38;5;220m%s\033[0m\n", totalNodes, totalTokens, costStr)
	fmt.Printf("\033[38;5;99m│\033[0m  🎓 Skills Hub   : \033[1;36m%d\033[0m skills registradas | \033[1;32m%d\033[0m skills ativas\n", availSkills, activeSkills)
	fmt.Printf("\033[38;5;99m├── 🌲 ÁRVORE DE DECISÃO ─────────────────────────────────────────────────────────────\033[0m\n")

	for _, r := range roots {
		t.renderNode(r.NodeID, "│  ", true)
	}

	fmt.Printf("\033[38;5;99m└── ──────────────────────────────────────────────────────────────────────────────────\033[0m\n")
}

func (t *SessionTree) renderNode(nodeID string, indent string, isLast bool) {
	node := t.Nodes[nodeID]

	connectorColor := "\033[38;5;242m"
	marker := connectorColor + "├── "
	nextIndent := indent + "\033[38;5;242m│\033[0m   "
	if isLast {
		marker = connectorColor + "└── "
		nextIndent = indent + "    "
	}

	activeMarker := ""
	if nodeID == t.ActiveNode {
		activeMarker = " \033[1;33m★ ATIVO\033[0m"
	}

	roleIcon := "⚙️"
	roleColor := "\033[38;5;86m" // Ciano claro para system
	roleLabel := "System"
	switch node.Role {
	case "user":
		roleIcon = "👤"
		roleColor = "\033[38;5;120m" // Verde para user
		roleLabel = "User"
	case "assistant":
		roleIcon = "🤖"
		roleColor = "\033[38;5;99m" // Roxo para assistant
		roleLabel = "Assistant"
	}
	resetColor := "\033[0m"

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
		tokenInfo = fmt.Sprintf(" \033[90m(🪙 %dt)\033[0m", node.Tokens)
	}

	preview := node.Content
	preview = strings.ReplaceAll(preview, "\n", " ")
	preview = strings.ReplaceAll(preview, "\r", "")
	if len(preview) > 55 {
		preview = preview[:52] + "..."
	}

	fmt.Printf("%s%s[%s%s%s%s] %s%s %s%s - \033[38;5;248m%s\033[0m%s%s\n",
		indent, marker, roleColor, node.NodeID, resetColor, activeMarker, roleColor, roleIcon, roleLabel, resetColor, timeStr, tokenInfo, resetColor)

	var children []Node
	for _, n := range t.Nodes {
		if n.ParentID == nodeID {
			children = append(children, n)
		}
	}

	sort.Slice(children, func(i, j int) bool {
		return children[i].Timestamp < children[j].Timestamp
	})

	for i, child := range children {
		t.renderNode(child.NodeID, nextIndent, i == len(children)-1)
	}
}
