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
		if node.Role == "system" {
			roleLabel = "System"
		} else if node.Role == "assistant" {
			roleLabel = "Assistant"
		}
		builder += fmt.Sprintf("[%s]: %s\n", roleLabel, node.Content)
	}
	return builder, nil
}

// PrintTree renders the DAG of sessions in a structured console directory format
func (t *SessionTree) PrintTree() {
	// Find roots (nodes with empty parent ID)
	var roots []Node
	for _, n := range t.Nodes {
		if n.ParentID == "" {
			roots = append(roots, n)
		}
	}

	// Sort roots chronologically by timestamp
	sort.Slice(roots, func(i, j int) bool {
		return roots[i].Timestamp < roots[j].Timestamp
	})

	fmt.Printf("🌳 Sessão: %s | Objetivo: \"%s\"\n", t.SessionID, t.RootTask)
	for _, root := range roots {
		t.renderNode(root.NodeID, "", true)
	}
}

func (t *SessionTree) renderNode(nodeID string, indent string, isLast bool) {
	node := t.Nodes[nodeID]

	// Determine bullet character
	marker := "├── "
	nextIndent := indent + "│   "
	if isLast {
		marker = "└── "
		nextIndent = indent + "    "
	}

	activeMarker := ""
	if nodeID == t.ActiveNode {
		activeMarker = " 🌟 (ATIVO)"
	}

	roleColor := "\033[36m" // Cyan for system
	if node.Role == "user" {
		roleColor = "\033[32m" // Green for user
	} else if node.Role == "assistant" {
		roleColor = "\033[33m" // Yellow for assistant
	}
	resetColor := "\033[0m"

	// Short preview of content
	preview := node.Content
	if len(preview) > 50 {
		preview = preview[:47] + "..."
	}
	preview = strings.ReplaceAll(preview, "\n", " ")

	fmt.Printf("%s%s[%s%s%s%s] %s%s%s\n", indent, marker, roleColor, node.NodeID, resetColor, activeMarker, roleColor, preview, resetColor)

	// Find children
	var children []Node
	for _, n := range t.Nodes {
		if n.ParentID == nodeID {
			children = append(children, n)
		}
	}

	// Sort children chronologically
	sort.Slice(children, func(i, j int) bool {
		return children[i].Timestamp < children[j].Timestamp
	})

	for i, child := range children {
		t.renderNode(child.NodeID, nextIndent, i == len(children)-1)
	}
}
