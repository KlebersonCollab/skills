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
	// Parse flags before positional commands
	streamMode := false
	remainingArgs := os.Args[1:]

	for i := 0; i < len(remainingArgs); i++ {
		arg := remainingArgs[i]
		if arg == "--stream" || arg == "-s" {
			streamMode = true
			remainingArgs = append(remainingArgs[:i], remainingArgs[i+1:]...)
			i--
		} else if arg == "--help" || arg == "-h" {
			remainingArgs[i] = "help"
		}
	}

	if len(remainingArgs) > 0 {
		cmd := remainingArgs[0]
		switch cmd {
		case "swarm":
			runSwarm()
			return
		case "mcp-server":
			runMCPServer()
			return
		case "mcp-client":
			if len(remainingArgs) < 3 {
				fmt.Println("Usage: harness mcp-client <transport> <address>")
				fmt.Println("  transport: \x27stdio\x27 or \x27sse\x27")
				fmt.Println("  address: command+args for stdio, URL for sse")
				return
			}
			transportType := remainingArgs[1]
			address := strings.Join(remainingArgs[2:], " ")
			runMCPClient(transportType, address)
			return
		case "init":
			runInitWizard()
			return
		case "help":
			printHelp()
			return
		default:
			fmt.Printf("Comando desconhecido: %s. Use 'harness help' para ver os comandos disponiveis.\n", cmd)
			return
		}
	}

	runInteractiveLoop(streamMode)
}

func printHelp() {
	fmt.Println("🚀 Harness - AI Agent Execution Engine (Go Edition)")
	fmt.Println("\nUso:")
	fmt.Println("  harness          - Inicia o loop de console interativo do agente")
	fmt.Println("  harness init     - Assistente interativo de configuração de LLM Providers")
	fmt.Println("  harness help     - Exibe esta mensagem de ajuda")
	fmt.Println("\nFlags:")
	fmt.Println("  --stream, -s     - Habilita modo SSE streaming (tokens em tempo real)")
	fmt.Println("  --help, -h       - Exibe esta mensagem de ajuda")
	fmt.Println("\nComandos de console (/slash commands):")
	fmt.Println("  /history         - Desenha a árvore DAG cronológica da sessão atual")
	fmt.Println("  /checkout <id>   - Altera o ponteiro ativo para o nó selecionado (branching)")
	fmt.Println("  /sessions        - Lista todas as sessões salvas no diretório local")
	fmt.Println("  /skills          - Lista todas as skills disponíveis no workspace")
	fmt.Println("  /exit, /quit     - Encerra a execução do harness")
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

func runInteractiveLoop(streamMode bool) {
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
