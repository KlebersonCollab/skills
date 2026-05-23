// Package agent provides the core agent execution loop and XML tool parser.
package agent

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"harness/internal/distiller"
	"harness/internal/llm"
	"harness/internal/session"
	"harness/internal/tools"
)

// Regex for XML attributes
var (
	patternAttrRE = regexp.MustCompile(`pattern="([^"]*)"`)
	queryAttrRE   = regexp.MustCompile(`query="([^"]*)"`)
	// Matches opening tag: <tool:name ...> or <tool:name .../>
	openTagRE     = regexp.MustCompile(`<tool:([a-zA-Z_][a-zA-Z0-9_:.-]*)((?:\s+[^>]*?)?)\s*/?>`) 
)

// Budget thresholds
const (
	BudgetWarning  = 40000 // tokens: comecar a monitorar
	BudgetCritical = 60000 // tokens: comprimir agressivamente
	BudgetMax      = 80000 // tokens: compactar nos antigos
)

// Config holds configuration for the agent loop.
type Config struct {
	MaxIterations            int
	MaxConsecutiveFailures   int
	StreamMode               bool
	DistillLevel             distiller.Level
	EnableDistiller          bool
	EnableCompactHistory     bool
	CompactMaxAgeTurns       int
	CompactMaxCharLen        int
	EnableBudgetManagement   bool                  // Ativar gerenciamento automatico de tokens
	BudgetWarning            int                   // Threshold amarelo (default: 40000)
	BudgetCritical           int                   // Threshold vermelho (default: 60000)
	WorkspaceRoot            string                // Raiz do workspace para descoberta de skills
	ActiveSkills             []session.SkillDef     // Skills disponiveis no workspace
	ProviderRegistry         *llm.ProviderRegistry  // Registry com fallback entre providers
	ActiveProviderName       string                // Nome do provider primario
	StatusFn                 func(string)          // Callback for status updates (called from agent loop)
}

// DefaultConfig returns sensible defaults for the agent loop.
func DefaultConfig() Config {
	return Config{
		MaxIterations:            25,
		MaxConsecutiveFailures:   8,
		StreamMode:               false,
		DistillLevel:             distiller.Full,
		
		EnableCompactHistory:     true,
		CompactMaxAgeTurns:       10,
		CompactMaxCharLen:        5000,
		EnableBudgetManagement:   true,
		BudgetWarning:            BudgetWarning,
		BudgetCritical:           BudgetCritical,
	}
}

// Logger abstracts output for the agent loop (allows mocking/testing).
type Logger interface {
	Log(format string, args ...interface{})
	LogRaw(text string)
}

// StdLogger implements Logger using fmt.Printf.
type StdLogger struct{}

func (l *StdLogger) Log(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
func (l *StdLogger) LogRaw(text string) {
	fmt.Print(text)
}

// Run executes the agent loop: system prompt → LLM → parse XML → execute tools → store → iterate.
func Run(
	ctx context.Context,
	cfg Config,
	log Logger,
	toolReg *tools.Registry,
	tree *session.Tree,
	rootTask string,
	mandatesContext string,
) (string, error) {
	if cfg.MaxIterations <= 0 {
		cfg = DefaultConfig()
	}

	consecutiveFailures := 0
	var lastResponse string

	// Build skills context at start
	skillsContext := buildSkillsContext(cfg.ActiveSkills)

	for i := 0; i < cfg.MaxIterations; i++ {
		select {
		case <-ctx.Done():
			return lastResponse, ctx.Err()
		default:
		}

		// ── 1. CONTEXT BUDGET MANAGEMENT ──
		totalTokens := tree.TotalTokens()
		if cfg.EnableBudgetManagement && totalTokens > 0 {
			budgetLog := ""

			if totalTokens > cfg.BudgetCritical {
				budgetLog = fmt.Sprintf("\033[38;5;196m⚠️  Orçamento CRÍTICO: %d tokens — comprimindo histórico...\033[0m", totalTokens)
				// Modo ultra: comprimir tudo
				if cfg.EnableDistiller {
					// Aplica Caveman Ultra nos nós antigos via compactação
					nodes := nodesToSlice(tree)
					compacted := distiller.Compact(nodes, cfg.CompactMaxAgeTurns, cfg.CompactMaxCharLen)
					// Write back compacted nodes (those that are > maxAgeTurns away from end)
					for _, cn := range compacted {
						if node, ok := cn.(*session.Node); ok {
							tree.Nodes[node.NodeID] = *node
						}
					}
				}
			} else if totalTokens > cfg.BudgetWarning {
				budgetLog = fmt.Sprintf("\033[38;5;208m⚠️  Orçamento elevado: %d tokens\033[0m", totalTokens)
			}

			if budgetLog != "" && log != nil {
				log.Log("%s\n", budgetLog)
			}
		}

		// ── 2. BUILD HISTORY ──
		historyText, err := tree.FormatHistory("")
		if err != nil {
			return lastResponse, fmt.Errorf("error formatting history: %w", err)
		}

		// ── 3. APPLY DISTILLER TO HISTORY (not system prompt) ──
		if cfg.EnableDistiller && totalTokens > cfg.BudgetWarning {
			historyText = distiller.ToCaveman(historyText, cfg.DistillLevel)
		}

		// ── 4. DETECT SKILL MENTIONS ──
		activeSkillContent := detectAndLoadSkills(cfg.WorkspaceRoot, cfg.ActiveSkills, rootTask, lastResponse)

		// ── 5. BUILD SYSTEM PROMPT ──
		systemPrompt := buildSystemPrompt(rootTask, skillsContext, mandatesContext, historyText, toolReg, activeSkillContent)

		// ── 6. ESTIMATE PROMPT TOKENS ──
		promptTokens := estimateTokens(systemPrompt)
		if cfg.EnableBudgetManagement && promptTokens+totalTokens > BudgetMax {
			log.Log("\033[38;5;196m⚠️  Prompt muito grande (%d tok). Removendo histórico mais antigo...\033[0m\n", promptTokens)
			// Fallback: usa só o ultimo turno
			lastNodes := getLastNodes(tree, 3)
			historyText = lastNodes
			_ = historyText
		}

		// ── 7. CALL LLM (COM FALLBACK) ──

		if cfg.StatusFn != nil {
			cfg.StatusFn("🤔 contacting LLM...")
		}

		// Check for images in the user input and prepare opts
		llmOpts := make(map[string]any)
		if images := extractImagesFromInput(rootTask, cfg.WorkspaceRoot); len(images) > 0 {
			llmOpts["images"] = images
			if cfg.StatusFn != nil {
				cfg.StatusFn(fmt.Sprintf("📷 %d imagem(ns) detectada(s), enviando modo multimodal...", len(images)))
			}
		}

		var response string
		if cfg.ProviderRegistry != nil {
			// Usa registry com fallback entre providers
			if cfg.StreamMode {
				var ch <-chan string
				ch, err = cfg.ProviderRegistry.StreamWithFallback(ctx, systemPrompt, llmOpts, func(format string, args ...any) {
					log.Log(format, args...)
				})
				if err == nil {
					var sb strings.Builder
					for token := range ch {
						log.LogRaw(token)
						sb.WriteString(token)
					}
					response = sb.String()
					log.Log("\n")
				}
			} else {
				response, err = cfg.ProviderRegistry.CompleteWithFallback(ctx, systemPrompt, llmOpts, func(format string, args ...any) {
					log.Log(format, args...)
				})
			}
		} else {
			// Fallback seguro: sem registry configurado
			err = fmt.Errorf("no provider registry configured")
		}

		if err != nil {
			consecutiveFailures++
			log.Log("❌ LLM call failed (%d/%d): %v\n", consecutiveFailures, cfg.MaxConsecutiveFailures, err)
			if consecutiveFailures >= cfg.MaxConsecutiveFailures {
				return lastResponse, fmt.Errorf("too many consecutive failures: %w", err)
			}
			time.Sleep(500 * time.Millisecond)
			continue
		}
		consecutiveFailures = 0

		// ── 7.5 DISPLAY USAGE ──
		if usage, ok := llmOpts["_usage"].(interface{ Format() string }); ok {
			if formatted := usage.Format(); formatted != "" {
				log.Log("\033[38;5;244m%s\033[0m\n", formatted)
			}
		}

		// ── 8. STORE RESPONSE ──
		tokens := estimateTokens(response)
		tree.AddNode("assistant", response, tokens)
		lastResponse = response

		// ── 9. SAVE SESSION ──
		if cfg.WorkspaceRoot != "" {
			sessionFile := filepath.Join(cfg.WorkspaceRoot, ".harness", "sessions", tree.SessionID+".json")
			tree.Save(sessionFile)
		}

		// ── 10. PARSE AND EXECUTE XML TOOLS ──
		if cfg.StatusFn != nil {
			cfg.StatusFn("🔍 parsing tools...")
		}
		toolResults := parseAndExecuteTools(response, toolReg)
		if len(toolResults) > 0 {
			if cfg.StatusFn != nil {
				cfg.StatusFn(fmt.Sprintf("🔧 executing %d tool(s)...", len(toolResults)))
			}
			for _, tr := range toolResults {
				tree.AddNode("system", tr, 0)
			}
			continue // Go to next iteration to process tool results
		}

		// ── 11. NO TOOL CALLS — FINAL ANSWER ──
		return response, nil
	}

	return lastResponse, fmt.Errorf("agent loop reached max iterations (%d) without final answer", cfg.MaxIterations)
}

// callStream handles streaming LLM calls.
func callStream(ctx context.Context, log Logger, provider llm.LLMProvider, prompt string) (string, error) {
	ch, err := provider.Stream(ctx, prompt, nil)
	if err != nil {
		return "", err
	}

	var fullText strings.Builder
	log.Log("\033[38;5;99m║\033[0m  ")
	for token := range ch {
		log.LogRaw(token)
		fullText.WriteString(token)
	}
	log.Log("\n")

	return fullText.String(), nil
}

// buildSystemPrompt constructs the full system prompt with skills, mandates, history, and tools.
func buildSystemPrompt(task string, skillsContext string, mandatesContext string, history string, toolReg *tools.Registry, activeSkillContent string) string {
	var b strings.Builder

	if mandatesContext != "" {
		b.WriteString(mandatesContext)
		b.WriteString("\n\n")
	}

	b.WriteString("Você é o Harness AI Agent. Você tem acesso às seguintes ferramentas em formato XML:\n\n")

	for _, name := range toolReg.Names() {
		b.WriteString(fmt.Sprintf("- %s: <tool:%s .../>\n", getToolDescription(name), name))
	}

	b.WriteString("\n")

	if skillsContext != "" {
		b.WriteString(skillsContext)
		b.WriteString("\n")
	}

	// Inject active skill content if a skill was detected
	if activeSkillContent != "" {
		b.WriteString(activeSkillContent)
		b.WriteString("\n")
	}

	b.WriteString(fmt.Sprintf("Objetivo da sessão: %s\n\n", task))

	if history != "" {
		b.WriteString("Histórico da sessão:\n")
		b.WriteString(history)
		b.WriteString("\n")
	}

	b.WriteString("\nResponda com texto ou use as ferramentas XML para interagir com o sistema.")

	return b.String()
}

// buildSkillsContext creates the standard skills listing prompt from discovered skills.
func buildSkillsContext(skills []session.SkillDef) string {
	if len(skills) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString("As seguintes Skills de Engenharia estão ativas e disponíveis no Hub (cada uma possui regras rígidas e diretrizes de desenvolvimento no arquivo SKILL.md de seu respectivo diretório):\n")
	for _, sk := range skills {
		b.WriteString(fmt.Sprintf("  - %s: %s (caminho: %s)\n", sk.Name, sk.Description, filepath.Join(sk.Path, "SKILL.md")))
	}
	b.WriteString("\n")
	return b.String()
}

// detectAndLoadSkills checks if the user input or last assistant response mentions any known skill.
// If a match is found, it loads the SKILL.md content and returns it as an injected system instruction.
func detectAndLoadSkills(root string, skills []session.SkillDef, userInput string, lastResponse string) string {
	combined := strings.ToLower(userInput + " " + lastResponse)

	for _, sk := range skills {
		nameLower := strings.ToLower(sk.Name)
		if strings.Contains(combined, nameLower) {
			content, err := loadSkillFile(filepath.Join(sk.Path, "SKILL.md"))
			if err == nil && content != "" {
				return fmt.Sprintf("\n## 🎓 SKILL ATIVADA: %s\nA skill '%s' foi detectada no contexto da tarefa. Siga rigorosamente as diretrizes abaixo:\n\n%s\n",
					sk.Name, sk.Name, content)
			}
		}
	}

	return ""
}

// loadSkillFile reads a SKILL.md file and returns its content for injection.
func loadSkillFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// parseAndExecuteTools scans text for XML tool calls and executes them.
// Supports 3 formats for all tools:
//   1. Attributes: <tool:name key="val"/>
//   2. JSON body:  <tool:name>{"key":"val"}</tool:name>
//   3. XML body:   <tool:name><key>val</key></tool:name>
func parseAndExecuteTools(text string, toolReg *tools.Registry) []string {
	var results []string

	// Find all tool opening tags and their positions
	matches := openTagRE.FindAllStringSubmatchIndex(text, -1)
	if matches == nil {
		return nil
	}

	for _, loc := range matches {
		fullStart := loc[0]
		fullEnd := loc[1]
		toolStart := loc[2]
		toolEnd := loc[3]
		attrStart := loc[4]
		attrEnd := loc[5]

		toolName := text[toolStart:toolEnd]
		attrs := ""
		if attrStart >= 0 && attrEnd >= 0 {
			attrs = strings.TrimSpace(text[attrStart:attrEnd])
		}

		// Check if self-closing (ends with />)
		openTag := text[fullStart:fullEnd]
		isSelfClosing := strings.HasSuffix(openTag, "/>")

		body := ""
		if !isSelfClosing {
			// Find closing tag: </tool:name>
			closeTag := fmt.Sprintf("</tool:%s>", toolName)
			closeIdx := strings.Index(text[fullEnd:], closeTag)
			if closeIdx >= 0 {
				body = strings.TrimSpace(text[fullEnd : fullEnd+closeIdx])
			}
		}

		// Parse args from attributes (format 1)
		args := parseGenericAttrs(attrs)

		// Parse body if present (formats 2 and 3)
		if body != "" {
			// Format 2: JSON body
			if strings.HasPrefix(body, "{") || strings.HasPrefix(body, "[") {
				var jsonBody map[string]interface{}
				if err := json.Unmarshal([]byte(body), &jsonBody); err == nil {
					for k, v := range jsonBody {
						if _, exists := args[k]; !exists {
							args[k] = v
						}
					}
				}
			}

			// Format 3: XML elements — extract <key>value</key>
			// Using non-backreference regex + manual validation
			elemRE := regexp.MustCompile(`<([a-zA-Z0-9_-]+)>([\s\S]*?)</[a-zA-Z0-9_-]+>`)
			for _, elem := range elemRE.FindAllStringSubmatch(body, -1) {
				key := elem[1]
				val := strings.TrimSpace(elem[2])
				// Verify closing tag matches opening tag
				closeTag := fmt.Sprintf("</%s>", key)
				if strings.Contains(body, closeTag) {
					if _, exists := args[key]; !exists {
						args[key] = val
					}
				}
			}
		}

		// Execute the tool
		res := toolReg.Execute(toolName, args)
		if res.Error != nil {
			results = append(results, fmt.Sprintf("%s error: %v", toolName, res.Error))
		} else if res.Output != "" {
			outputMsg := fmt.Sprintf("%s:\n%s", toolName, res.Output)
			if toolName == "execute_command" && res.ExitCode != 0 {
				outputMsg += fmt.Sprintf("\n(exit code: %d)", res.ExitCode)
			}
			results = append(results, outputMsg)
		} else {
			if toolName == "execute_command" && res.ExitCode != 0 {
				results = append(results, fmt.Sprintf("%s: (completed with exit code: %d, no output)", toolName, res.ExitCode))
			} else {
				results = append(results, fmt.Sprintf("%s: (executed successfully, empty output)", toolName))
			}
		}
	}

	return results
}

// parseGenericAttrs parses key="value" pairs from a string of XML attributes.
func parseGenericAttrs(attrs string) map[string]any {
	result := make(map[string]any)
	
	// Match key="value" patterns
	attrRE := regexp.MustCompile(`(\w+)="([^"]*)"`)
	for _, match := range attrRE.FindAllStringSubmatch(attrs, -1) {
		key := match[1]
		val := match[2]
		
		// Try to parse special types
		switch {
		case val == "true" || val == "false":
			result[key] = val == "true"
		case isNumber(val):
			result[key] = parseNumber(val)
		case strings.HasPrefix(val, "[") && strings.HasSuffix(val, "]"):
			// JSON array
			result[key] = parseJSONArray(val)
		default:
			result[key] = val
		}
	}
	
	return result
}

func isNumber(s string) bool {
	for _, c := range s {
		if c < '0' || c > '9' {
			if c != '.' && c != '-' {
				return false
			}
		}
	}
	return len(s) > 0
}

func parseNumber(s string) float64 {
	var v float64
	fmt.Sscanf(s, "%f", &v)
	return v
}

func parseJSONArray(s string) []interface{} {
	var v []interface{}
	if err := json.Unmarshal([]byte(s), &v); err == nil {
		return v
	}
	return nil
}

func getToolDescription(name string) string {
	// MCP tools: use the name as a readable description
	if strings.HasPrefix(name, "mcp:") {
		parts := strings.SplitN(name, ":", 3)
		if len(parts) == 3 {
			return fmt.Sprintf("[MCP:%s] %s — use <tool:%s .../>", parts[1], parts[2], name)
		}
		return fmt.Sprintf("MCP tool: %s", name)
	}

	descriptions := map[string]string{
		"read_file":       "Read file contents by path",
		"write_file":      "Write or overwrite a file",
		"patch_file":      "Apply surgical text replacement in a file",
		"search_files":    "Search files by name pattern and/or text content",
		"execute_command": "Execute a bash command",
		"web_search":      "Search the web with multiple queries. Supports: query (single), queries (JSON array), numResults, recencyFilter (day/week/month/year), domainFilter, provider (auto/perplexity/exa/gemini)",
		"code_search":     "Search code on GitHub and Stack Overflow. Params: query (string), maxTokens (number)",
		"fetch_content":   "Fetch content from URLs, YouTube, GitHub repos, or local videos. Params: url (single), urls (JSON array), path (local video), prompt (question about video), timestamp, frames (number)",

		"devtools":        "Browser/page interaction. Actions: fetch(url) - always works via HTTP, html(url) - tries Chrome + fallback HTTP, screenshot(url) - needs Chrome, list-tabs - needs Chrome DevTools",
		"paste_clipboard_image": "Read an image from the system clipboard (PrtSc/Ctrl+C de imagem). Retorna metadados e base64. Sem suporte a visão no LLM atual.",
		"image_info":     "Get metadata about an image file: path (string) - caminho absoluto. Retorna tipo, dimensões e tamanho.",
	}
	if desc, ok := descriptions[name]; ok {
		return desc
	}
	// Clean up the name for display: replace underscores and colons
	display := strings.ReplaceAll(name, "_", " ")
	display = strings.ReplaceAll(display, ":", ".")
	return display
}

// estimateTokens estima o numero de tokens em um texto.
// Usa heuristica: 1 token ~ 4 caracteres para prosa,
// mas codigo (backticks) conta como 1 token ~ 2 caracteres.
func estimateTokens(text string) int {
	if text == "" {
		return 0
	}

	total := 0
	inCode := false
	buf := ""

	for _, ch := range text {
		if ch == '`' {
			if inCode {
				// Fechou bloco de codigo
				total += len(buf) / 2
				if total < 1 {
					total = 1
				}
				buf = ""
				inCode = false
			} else {
				// Abriu bloco de codigo
				total += len(buf) / 4
				if total < 1 {
					total = 1
				}
				buf = ""
				inCode = true
			}
			continue
		}
		buf += string(ch)
	}

	// Restante
	if inCode {
		total += len(buf) / 2
	} else {
		total += len(buf) / 4
	}

	if total < 1 {
		total = 1
	}
	return total
}

// nodesToSlice converte o mapa de nodes do Tree em um slice ordenado por timestamp.
func nodesToSlice(tree *session.Tree) []distiller.HasContent {
	// Pega todos os nodes
	allNodes := make([]*session.Node, 0, len(tree.Nodes))
	for _, n := range tree.Nodes {
		copyN := n // copia para evitar pointer aliasing
		allNodes = append(allNodes, &copyN)
	}

	// Ordena por timestamp (mais antigos primeiro)
	// Usa bubble sort simples - map size e tipicamente < 50
	for i := 0; i < len(allNodes); i++ {
		for j := i + 1; j < len(allNodes); j++ {
			if allNodes[i].Timestamp > allNodes[j].Timestamp {
				allNodes[i], allNodes[j] = allNodes[j], allNodes[i]
			}
		}
	}

	result := make([]distiller.HasContent, len(allNodes))
	for i, n := range allNodes {
		result[i] = n
	}
	return result
}

// getLastNodes retorna os ultimos N turnos como string formatada.
func getLastNodes(tree *session.Tree, n int) string {
	history, err := tree.GetLinearHistory("")
	if err != nil {
		return ""
	}

	if len(history) > n {
		history = history[len(history)-n:]
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
	return b.String()
}

// imageExtensions is the set of recognized image file extensions.
var imageExtensions = map[string]bool{
	".png": true, ".jpg": true, ".jpeg": true, ".gif": true, ".webp": true,
}

// extractImagesFromInput scans user input for image file references (file:// URIs, absolute paths)
// and returns them as base64 data URIs ready for multimodal API calls.
func extractImagesFromInput(input string, workspaceRoot string) []string {
	words := strings.Fields(input)
	var images []string

	for _, w := range words {
		path := ""

		// Handle file:// URIs
		if strings.HasPrefix(w, "file://") {
			ref := strings.TrimPrefix(w, "file://")
			ref = strings.ReplaceAll(ref, "%20", " ")
			path = ref
		}

		// Handle @file references
		if strings.HasPrefix(w, "@") {
			ref := w[1:]
			if filepath.IsAbs(ref) {
				path = ref
			} else if workspaceRoot != "" {
				path = filepath.Join(workspaceRoot, ref)
			}
		}

		// Handle absolute paths without prefix (if they look like image paths)
		if path == "" && strings.HasPrefix(w, "/") {
			path = w
		}

		if path == "" {
			continue
		}

		ext := strings.ToLower(filepath.Ext(path))
		if !imageExtensions[ext] {
			continue
		}

		// Read and encode the image
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		mime := "image/png"
		switch ext {
		case ".jpg", ".jpeg":
			mime = "image/jpeg"
		case ".gif":
			mime = "image/gif"
		case ".webp":
			mime = "image/webp"
		}

		// Size limit: skip images over 10MB
		if len(data) > 10*1024*1024 {
			continue
		}

		dataURI := fmt.Sprintf("data:%s;base64,%s", mime, base64.StdEncoding.EncodeToString(data))
		images = append(images, dataURI)
	}

	return images
}
