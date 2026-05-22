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
	readFileRegex  = regexp.MustCompile(`<tool:read_file\s+path="([^"]+)"\s*/>`)
	writeFileRegex = regexp.MustCompile(`(?s)<tool:write_file\s+path="([^"]+)">([\s\S]*?)</tool:write_file>`)
	patchFileRegex = regexp.MustCompile(`(?s)<tool:patch_file\s+path="([^"]+)"><target>([\s\S]*?)</target><replacement>([\s\S]*?)</replacement></tool:patch_file>`)
	execCmdRegex   = regexp.MustCompile(`(?s)<tool:execute_command>([\s\S]*?)</tool:execute_command>`)
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

	// Dynamic API key verification
	if config.ActiveProvider == "gemini" && os.Getenv("GEMINI_API_KEY") == "" {
		fmt.Println("⚠️  Alerta: A variável de ambiente GEMINI_API_KEY não está definida!")
		fmt.Println("Por favor, defina a chave rodando: export GEMINI_API_KEY=sua_chave")
		fmt.Println()
	}
	if config.ActiveProvider == "deepseek" && os.Getenv("DEEPSEEK_API_KEY") == "" {
		fmt.Println("⚠️  Alerta: A variável de ambiente DEEPSEEK_API_KEY não está definida!")
		fmt.Println("Por favor, defina a chave rodando: export DEEPSEEK_API_KEY=sua_chave")
		fmt.Println()
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("🤖 Bem-vindo ao Harness Terminal AI Agent!")
	fmt.Println("Digite '/help' para obter comandos ou '/exit' para sair.")
	fmt.Println()

	// Initializing Session
	fmt.Print("Digite o objetivo principal da sessão (Root Task): ")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)
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

	fmt.Printf("\n🚀 Sessão criada com ID: \033[35m%s\033[0m\n", sessionID)
	fmt.Printf("Salvando histórico em: \033[36m%s\033[0m\n", sessionFile)
	fmt.Println("Harness pronto para execução!")

	for {
		fmt.Printf("\nharness [\033[33m%s\033[0m]> ", tree.ActiveNode)
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
				tree.PrintTree()
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

func agentExecutionLoop(config *AppConfig, tree *SessionTree, sessionFile string) {
	maxIterations := 8

	for i := 0; i < maxIterations; i++ {
		fmt.Printf("\033[90m[Harness] Enviando prompt para a LLM (%s) (Iteração %d/%d)...\033[0m\n", config.ActiveProvider, i+1, maxIterations)

		historyText, err := tree.FormatLinearHistoryText("")
		if err != nil {
			fmt.Printf("❌ Erro ao formatar histórico: %v\n", err)
			return
		}

		// Inject system rules about tool usage so the agent knows how to use tools
		systemInstructions := `Você é um Harness AI Agent. Você tem acesso às seguintes ferramentas de console em formato XML:\n` +
			`- Ler Arquivo: <tool:read_file path="caminho/relativo"/>\n` +
			`- Escrever/Sobrescrever: <tool:write_file path="caminho/relativo">conteudo</tool:write_file>\n` +
			`- Remendo cirúrgico de bloco único (Patch): <tool:patch_file path="caminho/relativo"><target>conteudo_exato_antigo</target><replacement>novo_conteudo</replacement></tool:patch_file>\n` +
			`- Executar Comando: <tool:execute_command>bash_comando</tool:execute_command>\n\n` +
			`Para usar qualquer ferramenta, emita a tag correspondente. Suas ações de escrita e comandos de bash são automaticamente filtrados pela governança SDD. Você só pode modificar arquivos se o STATE.md estiver em fase de IMPLEMENT ou VERIFY.\n\n`

		promptWithContext := systemInstructions + "Histórico atual da sessão:\n" + historyText + "\nContinue respondendo ao usuário ou chame uma ferramenta se necessário."

		llmResponse, err := CallLLM(config, promptWithContext, "")
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

		// 1. Check read_file
		if matches := readFileRegex.FindAllStringSubmatch(llmResponse, -1); len(matches) > 0 {
			hasTools = true
			for _, m := range matches {
				path := m[1]
				fmt.Printf("🔧 \033[33m[Ferramenta] read_file\033[0m executando para '%s'...\n", path)
				content, err := ReadFile(path)
				if err != nil {
					toolOutputs += fmt.Sprintf("<tool_response:read_file path=\"%s\" status=\"error\">%v</tool_response:read_file>\n", path, err)
					fmt.Printf("❌ Erro: %v\n", err)
				} else {
					toolOutputs += fmt.Sprintf("<tool_response:read_file path=\"%s\" status=\"success\">\n%s\n</tool_response:read_file>\n", path, content)
					fmt.Println("✅ Leitura realizada com sucesso.")
				}
			}
		}

		// 2. Check write_file
		if matches := writeFileRegex.FindAllStringSubmatch(llmResponse, -1); len(matches) > 0 {
			hasTools = true
			for _, m := range matches {
				path := m[1]
				content := m[2]
				fmt.Printf("🔧 \033[33m[Ferramenta] write_file\033[0m executando para '%s'...\n", path)
				err := WriteFile(path, content)
				if err != nil {
					toolOutputs += fmt.Sprintf("<tool_response:write_file path=\"%s\" status=\"error\">%v</tool_response:write_file>\n", path, err)
					fmt.Printf("❌ Erro: %v\n", err)
				} else {
					toolOutputs += fmt.Sprintf("<tool_response:write_file path=\"%s\" status=\"success\"/>\n", path)
					fmt.Println("✅ Escrita realizada com sucesso.")
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
				fmt.Printf("🔧 \033[33m[Ferramenta] patch_file\033[0m executando para '%s'...\n", path)
				err := PatchFile(path, target, replacement)
				if err != nil {
					toolOutputs += fmt.Sprintf("<tool_response:patch_file path=\"%s\" status=\"error\">%v</tool_response:patch_file>\n", path, err)
					fmt.Printf("❌ Erro: %v\n", err)
				} else {
					toolOutputs += fmt.Sprintf("<tool_response:patch_file path=\"%s\" status=\"success\"/>\n", path)
					fmt.Println("✅ Patch aplicado com sucesso.")
				}
			}
		}

		// 4. Check execute_command
		if matches := execCmdRegex.FindAllStringSubmatch(llmResponse, -1); len(matches) > 0 {
			hasTools = true
			for _, m := range matches {
				cmdStr := m[1]
				fmt.Printf("🔧 \033[33m[Ferramenta] execute_command\033[0m: '%s'...\n", cmdStr)
				out, code, err := ExecuteCommand(cmdStr)
				if err != nil {
					toolOutputs += fmt.Sprintf("<tool_response:execute_command status=\"error\" exit_code=\"%d\">%v\n%s</tool_response:execute_command>\n", code, err, out)
					fmt.Printf("❌ Erro (Código %d): %v\n", code, err)
				} else {
					toolOutputs += fmt.Sprintf("<tool_response:execute_command status=\"success\" exit_code=\"%d\">\n%s\n</tool_response:execute_command>\n", code, out)
					fmt.Printf("✅ Comando finalizado com código %d.\n", code)
				}
			}
		}

		if hasTools {
			// Add LLM's intermediate response as assistant node
			tree.AddNode("assistant", llmResponse, len(llmResponse)/4)
			// Add tool execution outputs as system node in DAG
			tree.AddNode("system", toolOutputs, len(toolOutputs)/4)
			tree.Save(sessionFile)
			// Continue execution loop with tool outputs injected
			continue
		}

		// If no tools were invoked, this is the final assistant response. Write to DAG and output
		tree.AddNode("assistant", llmResponse, len(llmResponse)/4)
		tree.Save(sessionFile)

		fmt.Println("\n\033[32m[Assistente]:\033[0m")
		fmt.Println(llmResponse)
		return
	}

	fmt.Println("\n⚠️  Atingido limite de iterações consecutivas de execução de ferramentas.")
}

func filepathWalkDir(root string, fn func(path string, d fs.DirEntry, err error) error) error {
	return filepath.WalkDir(root, fn)
}
