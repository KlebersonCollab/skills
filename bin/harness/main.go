package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var (
	readFileRegex    = regexp.MustCompile(`<tool:read_file\s+path="([^"]+)"\s*/>`)
	writeFileRegex   = regexp.MustCompile(`(?s)<tool:write_file\s+path="([^"]+)">([\s\S]*?)</tool:write_file>`)
	patchFileRegex   = regexp.MustCompile(`(?s)<tool:patch_file\s+path="([^"]+)"><target>([\s\S]*?)</target><replacement>([\s\S]*?)</replacement></tool:patch_file>`)
	execCmdRegex     = regexp.MustCompile(`(?s)<tool:execute_command>([\s\S]*?)</tool:execute_command>`)
	searchFilesRegex = regexp.MustCompile(`<tool:search_files\s+([\s\S]*?)\s*/?>`)
	patternAttrRegex = regexp.MustCompile(`pattern="([^"]*)"`)
	queryAttrRegex   = regexp.MustCompile(`query="([^"]*)"`)
)

func main() {
	if len(os.Args) > 1 {
		cmd := os.Args[1]
		switch cmd {
		case "init":
			runInitWizard()
			return
		case "help":
			printHelp()
			return
		default:
			fmt.Printf("Comando desconhecido: %s. Use 'harness help' para ver os comandos disponíveis.\n", cmd)
			return
		}
	}

	runInteractiveLoop()
}

func printHelp() {
	fmt.Println("🚀 Harness - AI Agent Execution Engine (Go Edition)")
	fmt.Println("\nUso:")
	fmt.Println("  harness        - Inicia o loop de console interativo do agente")
	fmt.Println("  harness init   - Assistente interativo de configuração de LLM Providers")
	fmt.Println("  harness help   - Exibe esta mensagem de ajuda")
	fmt.Println("\nComandos de console (/slash commands):")
	fmt.Println("  /history       - Desenha a árvore DAG cronológica da sessão atual")
	fmt.Println("  /checkout <id> - Altera o ponteiro ativo para o nó selecionado (branching)")
	fmt.Println("  /sessions      - Lista todas as sessões salvas no diretório local")
	fmt.Println("  /skills        - Lista todas as skills disponíveis no workspace")
	fmt.Println("  /exit, /quit   - Encerra a execução do harness")
}

func runInitWizard() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("🔧 --- HARNESS LLM PROVIDER SETUP WIZARD --- 🔧")
	fmt.Println("Este assistente irá gerar o arquivo '.harness/config.json' automaticamente.")
	fmt.Println()

	root, err := FindWorkspaceRoot()
	if err != nil {
		fmt.Printf("❌ Erro: Não foi possível localizar o workspace root. %v\n", err)
		return
	}

	harnessDir := filepath.Join(root, ".harness")
	configPath := filepath.Join(harnessDir, "config.json")

	fmt.Println("Escolha o Provider:")
	fmt.Println("1) Gemini (Google Generative AI)")
	fmt.Println("2) Ollama (Local Llama/Deepseek)")
	fmt.Println("3) DeepSeek (API Oficial)")
	fmt.Println("4) Custom (Qualquer API compatível REST JSON)")
	fmt.Print("Selecione (1-4) [1]: ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	if choice == "" {
		choice = "1"
	}

	var activeProvider string
	providers := make(map[string]ProviderConfig)

	switch choice {
	case "1":
		activeProvider = "gemini"
		fmt.Print("Digite a API Key [ou pressione Enter para usar {{GEMINI_API_KEY}}]: ")
		apiKey, _ := reader.ReadString('\n')
		apiKey = strings.TrimSpace(apiKey)

		// If the user pasted the API key, let's offer to save it in their env, or default to the placeholder
		urlStr := "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-pro:generateContent?key={{GEMINI_API_KEY}}"
		if apiKey != "" {
			os.Setenv("GEMINI_API_KEY", apiKey)
			fmt.Println("💡 API Key temporariamente guardada nesta sessão de console. Defina a variável de ambiente GEMINI_API_KEY permanentemente em seu shell.")
		}

		providers["gemini"] = ProviderConfig{
			URL: urlStr,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			BodyTemplate: `{"contents": [{"parts": [{"text": "{{prompt}}"}]}]}`,
			ResponsePath: "candidates.0.content.parts.0.text",
		}
		fmt.Println("✅ Gemini configurado como default.")

	case "2":
		activeProvider = "ollama"
		fmt.Print("Porta/Host do Ollama [http://localhost:11434]: ")
		host, _ := reader.ReadString('\n')
		host = strings.TrimSpace(host)
		if host == "" {
			host = "http://localhost:11434"
		}

		fmt.Print("Modelo do Ollama [llama3]: ")
		model, _ := reader.ReadString('\n')
		model = strings.TrimSpace(model)
		if model == "" {
			model = "llama3"
		}

		providers["ollama"] = ProviderConfig{
			URL: host + "/api/generate",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			BodyTemplate: fmt.Sprintf(`{"model": "%s", "prompt": "{{prompt}}", "stream": false}`, model),
			ResponsePath: "response",
		}
		fmt.Printf("✅ Ollama (%s) configurado como default.\n", model)

	case "3":
		activeProvider = "deepseek"
		fmt.Print("Digite a API Key do DeepSeek [ou pressione Enter para usar {{DEEPSEEK_API_KEY}}]: ")
		apiKey, _ := reader.ReadString('\n')
		apiKey = strings.TrimSpace(apiKey)
		if apiKey != "" {
			os.Setenv("DEEPSEEK_API_KEY", apiKey)
			fmt.Println("💡 API Key temporariamente guardada nesta sessão de console. Defina a variável de ambiente DEEPSEEK_API_KEY permanentemente em seu shell.")
		}

		fmt.Println("Escolha o Modelo:")
		fmt.Println("1) deepseek-v4-pro")
		fmt.Println("2) deepseek-v4-flash")
		fmt.Print("Selecione (1-2) [1]: ")
		mChoice, _ := reader.ReadString('\n')
		mChoice = strings.TrimSpace(mChoice)

		model := "deepseek-v4-pro"
		if mChoice == "2" {
			model = "deepseek-v4-flash"
		}

		providers["deepseek"] = ProviderConfig{
			URL: "https://api.deepseek.com/chat/completions",
			Headers: map[string]string{
				"Content-Type":  "application/json",
				"Authorization": "Bearer {{DEEPSEEK_API_KEY}}",
			},
			BodyTemplate: fmt.Sprintf(`{"model": "%s", "messages": [{"role": "user", "content": "{{prompt}}"}], "thinking": {"type": "enabled"}, "reasoning_effort": "high", "stream": false}`, model),
			ResponsePath: "choices.0.message.content",
		}
		fmt.Printf("✅ DeepSeek (%s) configurado como default.\n", model)

	case "4":
		fmt.Print("Nome do Provider Custom: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		if name == "" {
			name = "custom"
		}
		activeProvider = name

		fmt.Print("Endpoint URL (ex: http://api.openai.com/v1/chat/completions): ")
		urlStr, _ := reader.ReadString('\n')
		urlStr = strings.TrimSpace(urlStr)

		fmt.Print("Body Template JSON (use {{prompt}} e {{history}}): ")
		bodyTemp, _ := reader.ReadString('\n')
		bodyTemp = strings.TrimSpace(bodyTemp)

		fmt.Print("Response extraction JSONPath (ex: choices.0.message.content): ")
		resPath, _ := reader.ReadString('\n')
		resPath = strings.TrimSpace(resPath)

		providers[name] = ProviderConfig{
			URL:          urlStr,
			Headers:      map[string]string{"Content-Type": "application/json"},
			BodyTemplate: bodyTemp,
			ResponsePath: resPath,
		}
		fmt.Printf("✅ Provider customizado '%s' configurado.\n", name)
	}

	appConfig := &AppConfig{
		ActiveProvider: activeProvider,
		Providers:      providers,
	}

	if err := os.MkdirAll(harnessDir, 0755); err != nil {
		fmt.Printf("❌ Falha ao criar diretório '.harness': %v\n", err)
		return
	}

	if err := SaveConfig(configPath, appConfig); err != nil {
		fmt.Printf("❌ Falha ao salvar arquivo de configuração: %v\n", err)
		return
	}

	fmt.Printf("🎉 Configuração salva com sucesso em %s!\n", configPath)
}

func runInteractiveLoop() {
	root, err := FindWorkspaceRoot()
	if err != nil {
		fmt.Printf("❌ Erro de Workspace: %v\n", err)
		return
	}

	configPath := filepath.Join(root, ".harness", "config.json")
	config, err := LoadConfig(configPath)
	if err != nil {
		fmt.Println("⚠️  Configuração não encontrada ou corrompida.")
		fmt.Println("👉 Execute 'harness init' para configurar seu provider de LLM antes de rodar.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("🤖 Bem-vindo ao Harness Terminal AI Agent!")
	fmt.Println("Digite '/help' para obter comandos ou '/exit' para sair.")
	fmt.Println()

	// Dynamic API key verification and interactive input to prevent 401 Authentication errors
	if config.ActiveProvider == "gemini" && os.Getenv("GEMINI_API_KEY") == "" {
		fmt.Println("⚠️  Alerta: A variável de ambiente GEMINI_API_KEY não está definida!")
		fmt.Print("👉 Por favor, digite a sua chave de API Gemini (ou pressione Enter para ignorar): ")
		apiKey, _ := reader.ReadString('\n')
		apiKey = strings.TrimSpace(apiKey)
		if apiKey != "" {
			os.Setenv("GEMINI_API_KEY", apiKey)
			fmt.Println("✅ GEMINI_API_KEY definida com sucesso para esta execução de console!")
		} else {
			fmt.Println("❌ Nenhuma chave informada. Chamadas ao Gemini podem falhar com erro 401.")
		}
		fmt.Println()
	}
	if config.ActiveProvider == "deepseek" && os.Getenv("DEEPSEEK_API_KEY") == "" {
		fmt.Println("⚠️  Alerta: A variável de ambiente DEEPSEEK_API_KEY não está definida!")
		fmt.Print("👉 Por favor, digite a sua chave de API DeepSeek (sk-...) (ou pressione Enter para ignorar): ")
		apiKey, _ := reader.ReadString('\n')
		apiKey = strings.TrimSpace(apiKey)
		if apiKey != "" {
			os.Setenv("DEEPSEEK_API_KEY", apiKey)
			fmt.Println("✅ DEEPSEEK_API_KEY definida com sucesso para esta execução de console!")
		} else {
			fmt.Println("❌ Nenhuma chave informada. Chamadas ao DeepSeek podem falhar com erro 401.")
		}
		fmt.Println()
	}

	// Initializing Session
	fmt.Print("Digite o objetivo principal da sessão (Root Task): ")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)

	// Intercept accidental export commands pasted into the prompt
	if strings.HasPrefix(task, "export ") || strings.Contains(task, "=") {
		parts := strings.Split(task, "=")
		if len(parts) == 2 {
			varName := strings.TrimSpace(strings.ReplaceAll(parts[0], "export", ""))
			varValue := strings.TrimSpace(parts[1])
			if varName == "DEEPSEEK_API_KEY" || varName == "GEMINI_API_KEY" {
				os.Setenv(varName, varValue)
				fmt.Printf("💡 Detectado comando de export no prompt. Configurada chave %s em tempo de execução!\n", varName)
				fmt.Print("👉 Digite agora o objetivo REAL da sessão (Root Task): ")
				task, _ = reader.ReadString('\n')
				task = strings.TrimSpace(task)
			}
		}
	}

	if task == "" {
		task = "Sessão exploratória geral"
	}

	sessionHash := md5.Sum([]byte(time.Now().Format(time.RFC3339) + task))
	sessionID := "sess-" + hex.EncodeToString(sessionHash[:4])
	sessionFile := filepath.Join(root, ".harness", "sessions", sessionID+".json")

	tree := NewSessionTree(sessionID, task)
	// Create initial node
	tree.AddNode("system", fmt.Sprintf("Instrução do Harness Agent para o objetivo: %s", task), 0)
	tree.Save(sessionFile)

	// Fetch metrics for initial Welcome Board
	availSkills, activeSkills := GetSkillsMetrics(root)

	// Truncate path for beauty
	dispFile := sessionFile
	if len(dispFile) > 44 {
		dispFile = "..." + dispFile[len(dispFile)-41:]
	}

	dispTask := task
	if len(dispTask) > 38 {
		dispTask = dispTask[:35] + "..."
	}

	// Premium Welcome Board with micro-aesthetic boxes and pixel-perfect borders
	fmt.Printf("\n\033[38;5;99m┌────────────────────────────────────────────────────────────────────────────┐\033[0m\n")
	fmt.Printf("\033[38;5;99m│\033[0m   🚀  \033[1;38;5;159mHARNESS AI AGENT EXECUTION ENGINE (Go Edition)\033[0m                         \033[38;5;99m│\033[0m\n")
	fmt.Printf("\033[38;5;99m├────────────────────────────────────────────────────────────────────────────┤\033[0m\n")
	fmt.Print(formatBoxLine("Active Provider  : ", strings.ToUpper(config.ActiveProvider), "\033[1;32m"))
	fmt.Print(formatBoxLine("Session ID       : ", sessionID, "\033[1;35m"))
	fmt.Print(formatBoxLine("Root Task        : ", "\""+dispTask+"\"", "\033[0m"))
	
	skillsVal := fmt.Sprintf("%d disponíveis | %d ativas", availSkills, activeSkills)
	fmt.Print(formatBoxLine("Skills Hub       : ", skillsVal, "\033[1;36m"))
	
	tokensVal := fmt.Sprintf("%d tokens | $0.0000 USD", 0)
	if config.ActiveProvider == "ollama" {
		tokensVal = "0 tokens | $0.00 (Local/Free)"
	}
	fmt.Print(formatBoxLine("Session Cost     : ", tokensVal, "\033[1;32m"))
	fmt.Print(formatBoxLine("History Log      : ", dispFile, "\033[90m"))
	fmt.Printf("\033[38;5;99m├────────────────────────────────────────────────────────────────────────────┤\033[0m\n")
	fmt.Printf("\033[38;5;99m│\033[0m   💡 Digite \033[33m/help\033[0m para ver os comandos de controle slash disponíveis.        \033[38;5;99m│\033[0m\n")
	fmt.Printf("\033[38;5;99m│\033[0m   💡 Digite \033[33m/exit\033[0m para encerrar o Harness com segurança.                    \033[38;5;99m│\033[0m\n")
	fmt.Printf("\033[38;5;99m└────────────────────────────────────────────────────────────────────────────┘\033[0m\n")

	for {
		fmt.Printf("\n\033[1;38;5;99mharness\033[0m \033[38;5;198m❯\033[0m [\033[36m%s\033[0m] \033[38;5;198m❯\033[0m ", tree.ActiveNode)
		userInput, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		userInput = strings.TrimSpace(userInput)
		if userInput == "" {
			continue
		}

		if strings.HasPrefix(userInput, "/") {
			parts := strings.Fields(userInput)
			cmd := parts[0]
			switch cmd {
			case "/exit", "/quit":
				fmt.Println("Saindo do Harness. Até breve!")
				return
			case "/help":
				printHelp()
			case "/history":
				tree.PrintTree(config.ActiveProvider, root)
			case "/checkout":
				if len(parts) < 2 {
					fmt.Println("Erro: especifique o ID do nó. Uso: /checkout node-X")
					continue
				}
				nodeID := parts[1]
				if _, exists := tree.Nodes[nodeID]; !exists {
					fmt.Printf("Erro: Nó '%s' não existe nesta árvore.\n", nodeID)
					continue
				}
				tree.ActiveNode = nodeID
				tree.Save(sessionFile)
				fmt.Printf("Ponteiro de execução alterado para o nó: %s 🌟\n", nodeID)
			case "/sessions":
				listSessions(root)
			case "/skills":
				listSkills(root)
			default:
				fmt.Printf("Comando /slash desconhecido: %s\n", cmd)
			}
			continue
		}

		// Process normal text prompt: Add user node to DAG
		tree.AddNode("user", userInput, len(userInput)/4)
		tree.Save(sessionFile)

		agentExecutionLoop(config, tree, sessionFile)
	}
}

func listSessions(root string) {
	sessionsDir := filepath.Join(root, ".harness", "sessions")
	files, err := os.ReadDir(sessionsDir)
	if err != nil {
		fmt.Printf("Erro ao listar diretório de sessões: %v\n", err)
		return
	}

	fmt.Println("\n🗂️  Sessões ativas no diretório:")
	found := false
	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(sessionsDir, file.Name())
			t, err := LoadSessionTree(filePath)
			if err == nil {
				fmt.Printf("  • ID: \033[35m%s\033[0m | Objetivo: \"%s\" (Nós: %d)\n", t.SessionID, t.RootTask, len(t.Nodes))
				found = true
			}
		}
	}
	if !found {
		fmt.Println("  Nenhuma sessão encontrada.")
	}
}

func listSkills(root string) {
	fmt.Println("\n🎓 Skills disponíveis no hub:")
	files, err := os.ReadDir(root)
	if err != nil {
		fmt.Printf("Erro ao ler diretório raiz: %v\n", err)
		return
	}

	for _, file := range files {
		if file.IsDir() && !strings.HasPrefix(file.Name(), ".") && file.Name() != "bin" && file.Name() != "architecture" && file.Name() != "docs" {
			// Check if contains a SKILL.md
			skillMD := filepath.Join(root, file.Name(), "SKILL.md")
			if _, err := os.Stat(skillMD); err == nil {
				fmt.Printf("  • \033[32m%s\033[0m - %s/SKILL.md\n", file.Name(), file.Name())
			} else {
				// Search inside .agents/skills if applicable, or check sub-skill folder
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

func printAssistantCard(tree *SessionTree, provider string, nodeID string, content string) {
	lines := strings.Split(content, "\n")
	repeatCount := 50 - len(provider)
	if repeatCount < 0 {
		repeatCount = 0
	}
	fmt.Printf("\n\033[38;5;99m┌── \033[1;32m🤖 ASSISTENTE (%s)\033[0m \033[38;5;99m%s\033[0m\n", strings.ToUpper(provider), strings.Repeat("─", repeatCount))
	
	for _, line := range lines {
		fmt.Printf("\033[38;5;99m│\033[0m  %s\n", line)
	}
	
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	nodeTokens := 0
	if n, ok := tree.Nodes[nodeID]; ok {
		nodeTokens = n.Tokens
	}
	totalTokens := tree.GetTotalTokens()
	cost := tree.EstimateUSD(provider)
	costStr := fmt.Sprintf("$%.4f USD", cost)
	if cost == 0.0 {
		costStr = "Free/Local"
	}
	
	footerInfo := fmt.Sprintf("nó: %s | 🪙  %dt (total: %dt) | 💰 %s | 🕒 %s", nodeID, nodeTokens, totalTokens, costStr, timestamp)
	repeatCountFooter := 75 - len(footerInfo)
	if repeatCountFooter < 0 {
		repeatCountFooter = 0
	}
	fmt.Printf("\033[38;5;99m└── \033[90m(%s)\033[0m \033[38;5;99m%s\033[0m\n", footerInfo, strings.Repeat("─", repeatCountFooter))
}

func formatBoxLine(label string, val string, valColor string) string {
	totalWidth := 70
	visibleLen := len(label) + len(val)
	padding := totalWidth - visibleLen
	if padding < 0 {
		padding = 0
	}
	return fmt.Sprintf("\033[38;5;99m│\033[0m   %s%s%s%s   \033[38;5;99m│\033[0m\n", label, valColor, val, strings.Repeat(" ", padding))
}

func startSpinner(suffix string) chan struct{} {
	stopChan := make(chan struct{})
	go func() {
		frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
		i := 0
		start := time.Now()
		for {
			select {
			case <-stopChan:
				fmt.Print("\r\033[K")
				return
			default:
				elapsed := time.Since(start).Seconds()
				fmt.Printf("\r\033[38;5;99m%s\033[0m \033[1m%s\033[0m \033[90m(%.1fs)\033[0m", frames[i%len(frames)], suffix, elapsed)
				i++
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	return stopChan
}

func agentExecutionLoop(config *AppConfig, tree *SessionTree, sessionFile string) {
	maxIterations := 25
	consecutiveFailures := 0
	maxConsecutiveFailures := 8

	for i := 0; i < maxIterations; i++ {
		historyText, err := tree.FormatLinearHistoryText("")
		if err != nil {
			fmt.Printf("❌ Erro ao formatar histórico: %v\n", err)
			return
		}

		root, err := FindWorkspaceRoot()
		skillsText := "Nenhuma skill ativa encontrada no Hub."
		if err == nil {
			activeSkills := GetActiveSkillsList(root)
			if len(activeSkills) > 0 {
				skillsText = "As seguintes Skills de Engenharia estão ativas e disponíveis no Hub (cada uma possui regras rígidas e diretrizes de desenvolvimento no arquivo SKILL.md de seu respectivo diretório):\n"
				for _, sk := range activeSkills {
					skillsText += fmt.Sprintf("  - %s (caminho: %s/SKILL.md)\n", sk, sk)
				}
				skillsText += "\n💡 REQUISITO DE CONFORMIDADE OBRIGATÓRIA: Se o usuário pedir para usar uma skill ou se a tarefa envolver o domínio de uma delas (por exemplo, usar 'sdd', 'python-uv', 'clean-code-mentor', 'git-workflow', etc.), você DEVE obrigatoriamente ler o arquivo SKILL.md correspondente utilizando a ferramenta de leitura <tool:read_file path=\"nome_da_skill/SKILL.md\"/> para entender e aplicar todas as regras de qualidade, governança e engenharia nela especificadas. Não adivinhe as diretrizes; leia a Skill correspondente antes de continuar a execução.\n"
			}
		}

		// Inject system rules about tool usage so the agent knows how to use tools
		systemInstructions := `Você é um Harness AI Agent. Você tem acesso às seguintes ferramentas de console em formato XML:\n` +
			`- Ler Arquivo: <tool:read_file path="caminho/relativo"/>\n` +
			`- Escrever/Sobrescrever: <tool:write_file path="caminho/relativo">conteudo</tool:write_file>\n` +
			`- Remendo cirúrgico de bloco único (Patch): <tool:patch_file path="caminho/relativo"><target>conteudo_exato_antigo</target><replacement>novo_conteudo</replacement></tool:patch_file>\n` +
			`- Buscar Arquivos: <tool:search_files pattern="*.go" query="texto_busca"/> (Ambos os atributos são opcionais. Filtra por nome e/ou busca ocorrências de texto recursivamente no workspace de forma extremamente rápida e cross-platform)\n` +
			`- Executar Comando: <tool:execute_command>bash_comando</tool:execute_command>\n\n` +
			`Para usar qualquer ferramenta, emita a tag correspondente. Suas ações de escrita e comandos de bash são automaticamente filtrados pela governança SDD. Você só pode modificar arquivos se o STATE.md estiver em fase de IMPLEMENT ou VERIFY.\n\n` +
			skillsText + "\n"

		promptWithContext := systemInstructions + "Histórico atual da sessão:\n" + historyText + "\nContinue respondendo ao usuário ou chame uma ferramenta se necessário."

		stopSpinner := startSpinner(fmt.Sprintf("Chamando LLM (%s) [Iteração %d/%d, Falhas Consecutivas: %d/%d]...", config.ActiveProvider, i+1, maxIterations, consecutiveFailures, maxConsecutiveFailures))
		llmResponse, err := CallLLM(config, promptWithContext, "")
		close(stopSpinner)
		time.Sleep(50 * time.Millisecond)

		if err != nil {
			fmt.Printf("❌ Erro na chamada do Provider: %v\n", err)
			return
		}

		llmResponse = strings.TrimSpace(llmResponse)
		if llmResponse == "" {
			fmt.Println("⚠️  LLM respondeu com conteúdo vazio.")
			return
		}

		// Check for tool calls
		hasTools := false
		toolOutputs := ""
		hasAnyFailure := false

		// 1. Check read_file
		if matches := readFileRegex.FindAllStringSubmatch(llmResponse, -1); len(matches) > 0 {
			hasTools = true
			for _, m := range matches {
				path := m[1]
				fmt.Printf("\n\033[36m┌── ⚙️  [FERRAMENTA] LER ARQUIVO ────────────────────────────────────────\033[0m\n")
				fmt.Printf("\033[36m│\033[0m  📂 Arquivo: %s\n", path)
				fmt.Printf("\033[36m│\033[0m  🕒 Executando leitura segura...\n")
				content, err := ReadFile(path)
				if err != nil {
					toolOutputs += fmt.Sprintf("<tool_response:read_file path=\"%s\" status=\"error\">%v</tool_response:read_file>\n", path, err)
					fmt.Printf("\033[36m└── \033[31m❌ Erro: %v\033[0m ───────────────────────────────────────────────────\n", err)
					hasAnyFailure = true
				} else {
					toolOutputs += fmt.Sprintf("<tool_response:read_file path=\"%s\" status=\"success\">\n%s\n</tool_response:read_file>\n", path, content)
					fmt.Printf("\033[36m└── \033[32m✅ Sucesso! Lidos %d bytes.\033[0m ─────────────────────────────────────────\n", len(content))
				}
			}
		}

		// 1b. Check search_files
		if matches := searchFilesRegex.FindAllStringSubmatch(llmResponse, -1); len(matches) > 0 {
			hasTools = true
			for _, m := range matches {
				attrs := m[1]
				pattern := ""
				query := ""

				if patMatch := patternAttrRegex.FindStringSubmatch(attrs); len(patMatch) > 1 {
					pattern = patMatch[1]
				}
				if qMatch := queryAttrRegex.FindStringSubmatch(attrs); len(qMatch) > 1 {
					query = qMatch[1]
				}

				fmt.Printf("\n\033[36m┌── ⚙️  [FERRAMENTA] BUSCAR ARQUIVOS ───────────────────────────────────\033[0m\n")
				fmt.Printf("\033[36m│\033[0m  🔍 Filtro Nome: \"%s\" | Busca Texto: \"%s\"\n", pattern, query)
				fmt.Printf("\033[36m│\033[0m  🕒 Varrendo workspace de forma nativa e segura...\n")

				results, err := SearchFiles(pattern, query)
				if err != nil {
					toolOutputs += fmt.Sprintf("<tool_response:search_files pattern=\"%s\" query=\"%s\" status=\"error\">%v</tool_response:search_files>\n", pattern, query, err)
					fmt.Printf("\033[36m└── \033[31m❌ Erro: %v\033[0m ───────────────────────────────────────────────────\n", err)
					hasAnyFailure = true
				} else {
					toolOutputs += fmt.Sprintf("<tool_response:search_files pattern=\"%s\" query=\"%s\" status=\"success\">\n%s\n</tool_response:search_files>\n", pattern, query, results)
					linesCount := len(strings.Split(results, "\n"))
					if results == "Nenhum arquivo correspondente foi encontrado." {
						linesCount = 0
					}
					fmt.Printf("\033[36m└── \033[32m✅ Sucesso! Encontradas %d correspondências.\033[0m ─────────────────────────\n", linesCount)
				}
			}
		}

		// 2. Check write_file
		if matches := writeFileRegex.FindAllStringSubmatch(llmResponse, -1); len(matches) > 0 {
			hasTools = true
			for _, m := range matches {
				path := m[1]
				content := m[2]
				fmt.Printf("\n\033[33m┌── ⚙️  [FERRAMENTA] ESCREVER ARQUIVO ───────────────────────────────────\033[0m\n")
				fmt.Printf("\033[33m│\033[0m  📂 Arquivo: %s\n", path)
				fmt.Printf("\033[33m│\033[0m  📝 Gravando conteúdo...\n")
				err := WriteFile(path, content)
				if err != nil {
					toolOutputs += fmt.Sprintf("<tool_response:write_file path=\"%s\" status=\"error\">%v</tool_response:write_file>\n", path, err)
					fmt.Printf("\033[33m└── \033[31m❌ Erro: %v\033[0m ───────────────────────────────────────────────────\n", err)
					hasAnyFailure = true
				} else {
					toolOutputs += fmt.Sprintf("<tool_response:write_file path=\"%s\" status=\"success\"/>\n", path)
					fmt.Printf("\033[33m└── \033[32m✅ Sucesso! Escrita concluída.\033[0m ─────────────────────────────────────\n")
				}
			}
		}

		// 3. Check patch_file
		if matches := patchFileRegex.FindAllStringSubmatch(llmResponse, -1); len(matches) > 0 {
			hasTools = true
			for _, m := range matches {
				path := m[1]
				target := m[2]
				replacement := m[3]
				fmt.Printf("\n\033[33m┌── ⚙️  [FERRAMENTA] MODIFICAR ARQUIVO (PATCH) ───────────────────────────\033[0m\n")
				fmt.Printf("\033[33m│\033[0m  📂 Arquivo: %s\n", path)
				fmt.Printf("\033[33m│\033[0m  ✂️  Aplicando patch cirúrgico...\n")
				err := PatchFile(path, target, replacement)
				if err != nil {
					toolOutputs += fmt.Sprintf("<tool_response:patch_file path=\"%s\" status=\"error\">%v</tool_response:patch_file>\n", path, err)
					fmt.Printf("\033[33m└── \033[31m❌ Erro: %v\033[0m ───────────────────────────────────────────────────\n", err)
					hasAnyFailure = true
				} else {
					toolOutputs += fmt.Sprintf("<tool_response:patch_file path=\"%s\" status=\"success\"/>\n", path)
					fmt.Printf("\033[33m└── \033[32m✅ Sucesso! Patch aplicado com sucesso.\033[0m ───────────────────────────────────\n")
				}
			}
		}

		// 4. Check execute_command
		if matches := execCmdRegex.FindAllStringSubmatch(llmResponse, -1); len(matches) > 0 {
			hasTools = true
			for _, m := range matches {
				cmdStr := m[1]
				fmt.Printf("\n\033[35m┌── 💻  [FERRAMENTA] EXECUTAR COMANDO BASH ──────────────────────────────\033[0m\n")
				fmt.Printf("\033[35m│\033[0m  🐚 Comando: %s\n", cmdStr)
				fmt.Printf("\033[35m│\033[0m  🕒 Executando processo local...\n")
				
				stopCmdSpinner := startSpinner(fmt.Sprintf("Executando bash: '%s'...", cmdStr))
				out, code, err := ExecuteCommand(cmdStr)
				close(stopCmdSpinner)
				time.Sleep(50 * time.Millisecond)

				if len(strings.TrimSpace(out)) > 0 {
					fmt.Printf("\033[35m├── 📄  SAÍDA DO COMANDO ──────────────────────────────────────────────\033[0m\n")
					lines := strings.Split(strings.TrimSuffix(out, "\n"), "\n")
					for _, l := range lines {
						fmt.Printf("\033[35m│\033[0m  \033[90m%s\033[0m\n", l)
					}
				}

				if err != nil || code != 0 {
					hasAnyFailure = true
					errMsg := "Erro desconhecido"
					if err != nil {
						errMsg = err.Error()
					} else {
						errMsg = fmt.Sprintf("Comando falhou com código de saída %d", code)
					}
					toolOutputs += fmt.Sprintf("<tool_response:execute_command status=\"error\" exit_code=\"%d\">%s\n%s</tool_response:execute_command>\n", code, errMsg, out)
					fmt.Printf("\033[35m└── \033[31m❌ Erro (Código de saída %d): %s\033[0m ───────────────────────\n", code, errMsg)
				} else {
					toolOutputs += fmt.Sprintf("<tool_response:execute_command status=\"success\" exit_code=\"%d\">\n%s\n</tool_response:execute_command>\n", code, out)
					fmt.Printf("\033[35m└── \033[32m✅ Sucesso! Processo concluído (Código: 0).\033[0m ───────────────────\n")
				}
			}
		}

		if hasTools {
			if hasAnyFailure {
				consecutiveFailures++
			} else {
				consecutiveFailures = 0
			}

			// Add LLM's intermediate response as assistant node
			tree.AddNode("assistant", llmResponse, len(llmResponse)/4)
			// Add tool execution outputs as system node in DAG
			tree.AddNode("system", toolOutputs, len(toolOutputs)/4)
			tree.Save(sessionFile)

			if consecutiveFailures >= maxConsecutiveFailures {
				fmt.Printf("\n\033[31m┌── ⚠️  VÁLVULA DE SEGURANÇA DISPARADA ──────────────────────────────────\033[0m\n")
				fmt.Printf("\033[31m│\033[0m  🚨 Abortando loop de execução do agente.\n")
				fmt.Printf("\033[31m│\033[0m  💥 Motivo: %d falhas consecutivas de ferramentas detectadas.\n", consecutiveFailures)
				fmt.Printf("\033[31m└── ────────────────────────────────────────────────────────────────────\033[0m\n\n")
				return
			}

			// Continue execution loop with tool outputs injected
			continue
		}

		// If no tools were invoked, this is the final assistant response. Write to DAG and output
		tree.AddNode("assistant", llmResponse, len(llmResponse)/4)
		tree.Save(sessionFile)

		printAssistantCard(tree, config.ActiveProvider, tree.ActiveNode, llmResponse)
		return
	}

	fmt.Println("\n⚠️  Atingido limite de iterações consecutivas de execução de ferramentas.")
}

func filepathWalkDir(root string, fn func(path string, d fs.DirEntry, err error) error) error {
	return filepath.WalkDir(root, fn)
}

func GetActiveSkillsList(root string) []string {
	files, err := os.ReadDir(root)
	if err != nil {
		return nil
	}
	var list []string
	for _, file := range files {
		if file.IsDir() && !strings.HasPrefix(file.Name(), ".") && file.Name() != "bin" && file.Name() != "architecture" && file.Name() != "docs" {
			skillMD := filepath.Join(root, file.Name(), "SKILL.md")
			if _, err := os.Stat(skillMD); err == nil {
				list = append(list, file.Name())
			} else {
				subSkillMD := filepath.Join(root, file.Name(), ".agents", "skills", file.Name(), "SKILL.md")
				if _, err := os.Stat(subSkillMD); err == nil {
					list = append(list, file.Name())
				}
			}
		}
	}
	return list
}
