// Package session provides session tree management for tracking agent conversation history.
package session

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Node represents a single message/turn in the session tree.
type Node struct {
	NodeID    string `json:"node_id"`
	ParentID  string `json:"parent_id"`
	Role      string `json:"role"` // system, user, assistant
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
	Tokens    int    `json:"tokens"`
}

// GetContent implements distiller.HasContent.
func (n *Node) GetContent() string { return n.Content }

// SetContent implements distiller.HasContent.
func (n *Node) SetContent(s string) { n.Content = s }

// GetRole implements distiller.HasContent.
func (n *Node) GetRole() string { return n.Role }

// Tree represents a session as a DAG (Directed Acyclic Graph) of nodes.
type Tree struct {
	mu         sync.RWMutex     `json:"-"`                    // Thread safety
	SessionID  string           `json:"session_id"`
	RootTask   string           `json:"root_task"`
	Name       string           `json:"name,omitempty"`       // Display name
	Tags       []string         `json:"tags,omitempty"`
	ActiveNode string           `json:"active_node"`
	Nodes      map[string]Node  `json:"nodes"`
	CreatedAt  string           `json:"created_at,omitempty"`
	UpdatedAt  string           `json:"updated_at,omitempty"`
	Model      string           `json:"model,omitempty"`
	Provider   string           `json:"provider,omitempty"`
	TokenCost  float64          `json:"token_cost,omitempty"`
}

// NewTree creates a new session tree with the given ID and root task.
func NewTree(sessionID string, rootTask string) *Tree {
	now := time.Now().UTC().Format(time.RFC3339)
	return &Tree{
		SessionID:  sessionID,
		RootTask:   rootTask,
		ActiveNode: "",
		Nodes:      make(map[string]Node),
		CreatedAt:  now,
		UpdatedAt:  now,
	}
}

// Lock locks the tree for writing.
func (t *Tree) Lock() {
	t.mu.Lock()
}

// Unlock unlocks the tree.
func (t *Tree) Unlock() {
	t.mu.Unlock()
}

// RLock locks the tree for reading.
func (t *Tree) RLock() {
	t.mu.RLock()
}

// RUnlock unlocks the tree for reading.
func (t *Tree) RUnlock() {
	t.mu.RUnlock()
}

// SetName sets the display name of the session. Thread-safe.
func (t *Tree) SetName(name string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.Name = name
	t.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
}

// SetModel records which model/provider was used.
func (t *Tree) SetModel(model, provider string) {
	t.Model = model
	t.Provider = provider
	t.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
}

// Touch updates the UpdatedAt timestamp.
func (t *Tree) Touch() {
	t.UpdatedAt = time.Now().UTC().Format(time.RFC3339)
}

// Info returns a formatted summary of the session. Thread-safe.
func (t *Tree) Info(provider string) string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	var b strings.Builder
	b.WriteString(fmt.Sprintf("🆔 ID:        \033[35m%s\033[0m\n", t.SessionID))
	if t.Name != "" {
		b.WriteString(fmt.Sprintf("📛 Nome:      %s\n", t.Name))
	}
	b.WriteString(fmt.Sprintf("🎯 Objetivo:  \"%s\"\n", t.RootTask))
	b.WriteString(fmt.Sprintf("📊 Nos:       %d\n", len(t.Nodes)))
	b.WriteString(fmt.Sprintf("🪙 Tokens:    %d\n", t.TotalTokens()))
	b.WriteString(fmt.Sprintf("💰 Custo:     $%.4f USD\n", t.EstimateUSD(provider)))
	b.WriteString(fmt.Sprintf("📍 Ativo:     %s\n", t.ActiveNode))
	b.WriteString(fmt.Sprintf("🕐 Criado:    %s\n", t.CreatedAt[:19]))
	b.WriteString(fmt.Sprintf("🕐 Atualizado: %s\n", t.UpdatedAt[:19]))
	if t.Provider != "" {
		b.WriteString(fmt.Sprintf("🤖 Provider:  %s\n", strings.ToUpper(t.Provider)))
	}
	if len(t.Tags) > 0 {
		b.WriteString(fmt.Sprintf("🏷️ Tags:      %s\n", strings.Join(t.Tags, ", ")))
	}
	return b.String()
}

// CloneBranch creates a new Tree with the active path copied into a new session. Thread-safe.
func (t *Tree) CloneBranch(newID string) *Tree {
	t.mu.RLock()
	defer t.mu.RUnlock()
	history, err := t.GetLinearHistory("")
	if err != nil {
		return nil
	}

	newTree := NewTree(newID, t.RootTask+" (clone)")
	newTree.Name = t.Name
	newTree.Model = t.Model
	newTree.Provider = t.Provider

	for _, n := range history {
		newTree.AddNode(n.Role, n.Content, n.Tokens)
	}

	return newTree
}

// ExportHTML exports the session as a standalone HTML file. Thread-safe.
func (t *Tree) ExportHTML() string {
	t.mu.RLock()
	defer t.mu.RUnlock()
	history, err := t.GetLinearHistory("")
	if err != nil {
		return fmt.Sprintf("Error: %v", err)
	}

	var messages strings.Builder
	for _, node := range history {
		roleClass := "user"
		roleLabel := "Usuario"
		icon := "👤"
		switch node.Role {
		case "system":
			roleClass = "system"
			roleLabel = "Sistema"
			icon = "⚙️"
		case "assistant":
			roleClass = "assistant"
			roleLabel = "Assistente"
			icon = "🤖"
		}
		messages.WriteString(fmt.Sprintf(`
		<div class="message %s">
			<div class="role">%s %s</div>
			<div class="content">%s</div>
			<div class="meta">%d tokens</div>
		</div>`, roleClass, icon, roleLabel, htmlEscape(node.Content), node.Tokens))
	}

	name := t.Name
	if name == "" {
		name = t.RootTask
	}

	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="pt-BR">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Sessao: %s</title>
	<style>
		* { margin: 0; padding: 0; box-sizing: border-box; }
		body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
		       background: #1a1a2e; color: #e0e0e0; max-width: 800px; margin: 0 auto; padding: 20px; }
		h1 { color: #e94560; border-bottom: 2px solid #e94560; padding-bottom: 10px; margin-bottom: 20px; }
		.meta-info { color: #888; font-size: 0.9em; margin-bottom: 30px; }
		.meta-info span { margin-right: 20px; }
		.message { margin-bottom: 20px; padding: 15px; border-radius: 8px; }
		.message.user { background: #16213e; border-left: 4px solid #0f3460; }
		.message.assistant { background: #1a1a3e; border-left: 4px solid #e94560; }
		.message.system { background: #111; border-left: 4px solid #888; }
		.role { font-weight: bold; margin-bottom: 8px; color: #e94560; }
		.content { line-height: 1.6; white-space: pre-wrap; }
		.content code { background: #2a2a4e; padding: 2px 6px; border-radius: 3px; font-size: 0.9em; }
		.content pre { background: #2a2a4e; padding: 10px; border-radius: 5px; overflow-x: auto; margin: 10px 0; }
		.meta { color: #666; font-size: 0.8em; margin-top: 8px; }
		.footer { margin-top: 40px; padding-top: 20px; border-top: 1px solid #333; color: #666; font-size: 0.8em; }
	</style>
</head>
<body>
	<h1>📝 %s</h1>
	<div class="meta-info">
		<span>🆔 %s</span>
		<span>📊 %s mensagens</span>
		<span>🪙 %s tokens</span>
	</div>
	%s
	<div class="footer">
		Exportado do Harness AI Agent em %s
	</div>
</body>
</html>`, htmlEscape(name), htmlEscape(name), htmlEscape(t.SessionID),
		strconv.Itoa(len(t.Nodes)), strconv.Itoa(t.TotalTokens()),
		messages.String(), time.Now().Format("2006-01-02 15:04:05"))
}

// ExportText exports the session as formatted text (markdown-like).
func (t *Tree) ExportText() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("# Sessao: %s\n\n", t.RootTask))
	if t.Name != "" {
		b.WriteString(fmt.Sprintf("> Nome: %s\n\n", t.Name))
	}

	history, err := t.GetLinearHistory("")
	if err != nil {
		return fmt.Sprintf("Erro ao exportar: %v", err)
	}

	for _, node := range history {
		roleIcon := "👤"
		roleLabel := "Usuario"
		switch node.Role {
		case "system":
			roleIcon = "⚙️"
			roleLabel = "Sistema"
		case "assistant":
			roleIcon = "🤖"
			roleLabel = "Assistente"
		}
		b.WriteString(fmt.Sprintf("## %s %s\n\n", roleIcon, roleLabel))
		b.WriteString(node.Content)
		b.WriteString("\n\n---\n\n")
	}

	b.WriteString(fmt.Sprintf("\n*Exportado de %s | %d nos | %d tokens*\n", t.SessionID, len(t.Nodes), t.TotalTokens()))
	return b.String()
}

// LoadTree loads a session tree from a JSON file.
func LoadTree(path string) (*Tree, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tree Tree
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tree); err != nil {
		return nil, err
	}
	return &tree, nil
}

// Save serializes the session tree to a JSON file. Thread-safe.
func (t *Tree) Save(path string) error {
	t.mu.RLock()
	defer t.mu.RUnlock()

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

// AddNode adds a new node to the tree, linked to the current active node.
// Thread-safe.
func (t *Tree) AddNode(role string, content string, tokens int) string {
	t.mu.Lock()
	defer t.mu.Unlock()

	nodeID := fmt.Sprintf("node-%d", len(t.Nodes))
	parentID := t.ActiveNode

	node := Node{
		NodeID:    nodeID,
		ParentID:  parentID,
		Role:      role,
		Content:   content,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Tokens:    tokens,
	}

	t.Nodes[nodeID] = node
	t.ActiveNode = nodeID
	t.Touch()
	return nodeID
}

// GetLinearHistory traverses backwards from startNodeID to the root.
// Thread-safe (read lock).
func (t *Tree) GetLinearHistory(startNodeID string) ([]Node, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

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

	// Reverse to chronological order
	for i, j := 0, len(history)-1; i < j; i, j = i+1, j-1 {
		history[i], history[j] = history[j], history[i]
	}

	return history, nil
}

// FormatHistory returns the linear history as a formatted string.
func (t *Tree) FormatHistory(startNodeID string) (string, error) {
	history, err := t.GetLinearHistory(startNodeID)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	for _, node := range history {
		roleLabel := "User"
		switch node.Role {
		case "system":
			roleLabel = "System"
		case "assistant":
			roleLabel = "Assistant"
		}
		b.WriteString(fmt.Sprintf("[%s]: %s\n", roleLabel, node.Content))
	}
	return b.String(), nil
}

// TotalTokens returns the sum of tokens across all nodes. Thread-safe.
func (t *Tree) TotalTokens() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	total := 0
	for _, n := range t.Nodes {
		total += n.Tokens
	}
	return total
}

// EstimateUSD estimates the cost in USD based on the provider type.
func (t *Tree) EstimateUSD(provider string) float64 {
	totalTokens := t.TotalTokens()
	var ratePerToken float64
	switch strings.ToLower(provider) {
	case "deepseek":
		ratePerToken = 0.000002
	case "gemini":
		ratePerToken = 0.0000035
	default:
		ratePerToken = 0.0
	}
	return float64(totalTokens) * ratePerToken
}

// DiscoverSkills scans the workspace for skill directories containing SKILL.md files.
func DiscoverSkills(root string) []SkillDef {
	var skills []SkillDef

	searchDirs := []string{
		filepath.Join(root, ".agents", "skills"),
		filepath.Join(root, ".gemini", "skills"),
		filepath.Join(root, ".claude", "skills"),
		root,
	}

	seen := make(map[string]bool)

	for _, dir := range searchDirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, entry := range entries {
			if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") && entry.Name() != "bin" && entry.Name() != "architecture" && entry.Name() != "docs" {
				skillName := entry.Name()
				if seen[skillName] {
					continue
				}

				skillPath := filepath.Join(dir, entry.Name())
				skillMDPath := filepath.Join(skillPath, "SKILL.md")

				if _, err := os.Stat(skillMDPath); err == nil {
					description := extractSkillDescription(skillMDPath)
					skills = append(skills, SkillDef{
						Name:        skillName,
						Description: description,
						Path:        skillPath,
					})
					seen[skillName] = true
				}
			}
		}
	}

	return skills
}

// SkillDef represents a discovered skill in the workspace.
type SkillDef struct {
	Name        string
	Description string
	Path        string
}

func extractSkillDescription(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return "Sem descrição (falha ao ler o arquivo)"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inFrontmatter := false
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "---" {
			if !inFrontmatter {
				inFrontmatter = true
				continue
			}
			break
		}
		if inFrontmatter && strings.HasPrefix(line, "description:") {
			desc := strings.TrimSpace(strings.TrimPrefix(line, "description:"))
			desc = strings.Trim(desc, `"'`)
			return desc
		}
	}

	return "Sem descrição"
}

// GetSkillsMetrics returns the count of available and active skills.
func GetSkillsMetrics(root string) (available int, active int) {
	skills := DiscoverSkills(root)
	return len(skills), len(skills)
}

// ListSessions displays all saved sessions in the .harness/sessions directory.
// SessionInfo holds metadata about a saved session.
type SessionInfo struct {
	ID       string `json:"id"`
	Task     string `json:"task"`
	Nodes    int    `json:"nodes"`
	File     string `json:"file"`
	Modified string `json:"modified"`
}

// FindSessions returns all saved sessions sorted by most recent first.
func FindSessions(root string) []SessionInfo {
	sessionsDir := filepath.Join(root, ".harness", "sessions")
	entries, err := os.ReadDir(sessionsDir)
	if err != nil {
		return nil
	}

	var sessions []SessionInfo
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		filePath := filepath.Join(sessionsDir, entry.Name())
		t, err := LoadTree(filePath)
		if err != nil {
			continue
		}
		info, _ := entry.Info()
		modTime := ""
		if info != nil {
			modTime = info.ModTime().Format("2006-01-02 15:04")
		}
		sessions = append(sessions, SessionInfo{
			ID:       t.SessionID,
			Task:     t.RootTask,
			Nodes:    len(t.Nodes),
			File:     filePath,
			Modified: modTime,
		})
	}

	// Sort by modified time descending (most recent first)
	for i := 0; i < len(sessions); i++ {
		for j := i + 1; j < len(sessions); j++ {
			if sessions[i].Modified < sessions[j].Modified {
				sessions[i], sessions[j] = sessions[j], sessions[i]
			}
		}
	}

	return sessions
}

// ListSessions displays all saved sessions in a formatted list.
func ListSessions(root string) {
	sessions := FindSessions(root)
	if len(sessions) == 0 {
		fmt.Println("\n🗂️  Nenhuma sessão encontrada.")
		return
	}

	fmt.Println("\n🗂️  Sessões ativas:")
	for _, s := range sessions {
		fmt.Printf("  • \033[35m%s\033[0m | \"%s\" (%d nós) \033[90m%s\033[0m\n", s.ID, s.Task, s.Nodes, s.Modified)
	}
}

// ListSkills displays all discovered skills.
// ForkSession creates a new session from a specific node of an existing session.
// Returns the new session tree with history up to the forked node.
func ForkSession(original *Tree, fromNodeID string, newSessionID string) *Tree {
	node, exists := original.Nodes[fromNodeID]
	if !exists {
		return nil
	}

	// Collect history from root to this node
	history, _ := original.GetLinearHistory(fromNodeID)

	newTree := NewTree(newSessionID, original.RootTask+" (fork)")
	for _, n := range history {
		newTree.AddNode(n.Role, n.Content, n.Tokens)
	}
	_ = node // forked node reference

	return newTree
}

// FindMostRecentSession returns the most recently modified session file.
func FindMostRecentSession(root string) *SessionInfo {
	sessions := FindSessions(root)
	if len(sessions) == 0 {
		return nil
	}
	return &sessions[0]
}

// SessionDir returns the session storage directory.
func SessionDir(root string) string {
	return filepath.Join(root, ".harness", "sessions")
}

// htmlEscape escapes special HTML characters.
func htmlEscape(s string) string {
	s = strings.ReplaceAll(s, "&", "&amp;")
	s = strings.ReplaceAll(s, "<", "&lt;")
	s = strings.ReplaceAll(s, ">", "&gt;")
	s = strings.ReplaceAll(s, "\"", "&quot;")
	return s
}

func ListSkills(root string) {
	fmt.Println("\n🎓 Skills disponíveis no hub:")
	files, err := os.ReadDir(root)
	if err != nil {
		fmt.Printf("Erro ao ler diretório raiz: %v\n", err)
		return
	}

	for _, file := range files {
		if file.IsDir() && !strings.HasPrefix(file.Name(), ".") && file.Name() != "bin" && file.Name() != "architecture" && file.Name() != "docs" {
			skillMD := filepath.Join(root, file.Name(), "SKILL.md")
			if _, err := os.Stat(skillMD); err == nil {
				fmt.Printf("  • \033[32m%s\033[0m - %s/SKILL.md\n", file.Name(), file.Name())
			} else {
				subSkillMD := filepath.Join(root, file.Name(), ".agents", "skills", file.Name(), "SKILL.md")
				if _, err := os.Stat(subSkillMD); err == nil {
					fmt.Printf("  • \033[32m%s\033[0m (Complex) - %s\n", file.Name(), subSkillMD)
				} else {
					fmt.Printf("  • \033[32m%s\033[0m (Diretório de Skill)\n", file.Name())
				}
			}
		}
	}
}
