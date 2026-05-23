package main

import (
	"bufio"
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"harness/cli"
	"harness/internal/agent"
	"harness/internal/config"
	"harness/internal/session"
	"harness/internal/tools"
	"harness/internal/tui"
)

func runInteractiveLoop(streamMode bool) {
	root := getRoot()
	cfg := loadAppCfg(root)
	if cfg == nil {
		return
	}
	reader := bufio.NewReader(os.Stdin)
	reg := buildProviderReg(cfg)
	if reg == nil {
		return
	}
	toolReg := buildToolRegistry(root, cfg.MCPServers...)
	skillDefs := session.DiscoverSkills(root)
	fmt.Print("Task: ")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)
	if task == "" {
		task = "General exploration"
	}
	sessionHash := md5.Sum([]byte(time.Now().Format(time.RFC3339) + task))
	sessionID := "sess-" + hex.EncodeToString(sessionHash[:4])
	sessionFile := filepath.Join(root, ".harness", "sessions", sessionID+".json")
	tree := session.NewTree(sessionID, task)
	tree.AddNode("system", "Harness Agent session: "+task, 0)
	tree.Save(sessionFile)
	tc := tui.New()
	modelName := extractCurrentModel(cfg)
	tc.WelcomeHeader(cfg.ActiveProvider, modelName, sessionID, task, len(skillDefs))
	tc.Footer(tree, tree.EstimateUSD(cfg.ActiveProvider))
	ac := agent.DefaultConfig()
	ac.StreamMode = streamMode
	ac.WorkspaceRoot = root
	ac.ActiveSkills = skillDefs
	ac.ProviderRegistry = reg
	ac.ActiveProviderName = reg.Names()[0]
	runInteractivePrompt(tree, sessionFile, cfg, ac, toolReg, tc, loadMandates(root), root)
}

func runInteractivePrompt(tree *session.Tree, sessionFile string, appCfg *config.AppConfig,
	ac agent.Config, toolReg *tools.Registry, tc *tui.TUI, mandates string, root string) {

	reader := bufio.NewReader(os.Stdin)
	msgQueue := make(chan string, 10)
	agentBusy := false
	var cancel context.CancelFunc
	var cancelCtx context.Context
	cancelCtx, cancel = context.WithTimeout(context.Background(), 3*time.Minute)

	ac.StatusFn = func(msg string) {
		fmt.Printf("\r\033[90m▸ %s\033[0m\033[K\n", msg)
	}

	for {
		if agentBusy {
			fmt.Print("\n\033[38;5;208m⏳ busy — Enter=steer | Esc=cancel\033[0m ")
		} else {
			fmt.Printf("\n\033[1;38;5;99mharness\033[0m \033[38;5;198m❯\033[0m [\033[36m%s\033[0m] \033[38;5;198m❯\033[0m ", tree.ActiveNode)
		}

		userInput, err := readInput(reader)
		if err != nil {
			break
		}

		if userInput == "" && agentBusy {
			cancel()
			fmt.Println("\033[38;5;208mCancelled.\033[0m")
			cancelCtx, cancel = context.WithTimeout(context.Background(), 3*time.Minute)
			agentBusy = false
			continue
		}

		userInput = strings.TrimSpace(userInput)
		if userInput == "" {
			continue
		}

		// Bash inline !!
		if strings.HasPrefix(userInput, "!!") {
			cmd := strings.TrimSpace(strings.TrimPrefix(userInput, "!!"))
			if cmd != "" {
				out, code, execErr := execCmd(cmd, root)
				if execErr != nil {
					fmt.Printf("\033[38;5;196mError: %v\033[0m\n", execErr)
				} else {
					fmt.Printf("\033[90m%s\033[0m\n", out)
					if code != 0 {
						fmt.Printf("\033[38;5;208mExit: %d\033[0m\n", code)
					}
				}
			}
			continue
		}

		// Bash inline !
		if strings.HasPrefix(userInput, "!") && !strings.HasPrefix(userInput, "!!") {
			cmd := strings.TrimSpace(strings.TrimPrefix(userInput, "!"))
			if cmd != "" {
				out, code, execErr := execCmd(cmd, root)
				if execErr != nil {
					fmt.Printf("\033[38;5;196mError: %v\033[0m\n", execErr)
				} else {
					note := ""
					if code != 0 {
						note = fmt.Sprintf(" (exit: %d)", code)
					}
					userInput = fmt.Sprintf("Ran: %s%s\nOutput:\n%s\n\nBased on this:", cmd, note, out)
				}
			}
		}

		// @file references
		userInput = expandFileRefs(userInput, root)

		// Slash commands
		if strings.HasPrefix(userInput, "/") {
			parts := strings.Fields(userInput)
			cmd := parts[0]
			switch cmd {
			case "/exit", "/quit":
				fmt.Println("Bye!")
				return
			case "/help":
				cli.PrintHelp()
			case "/session":
				fmt.Println(tree.Info(appCfg.ActiveProvider))
			case "/history":
				tc.PrintTree(tree, root)
			case "/sessions":
				session.ListSessions(root)
			case "/skills":
				session.ListSkills(root)
			case "/name":
				if len(parts) >= 2 {
					tree.SetName(strings.Join(parts[1:], " "))
					tree.Save(sessionFile)
					fmt.Printf("Renamed: %s\n", strings.Join(parts[1:], " "))
				}
			case "/clone":
				newID := "clone-" + hex.EncodeToString([]byte(time.Now().String()))[:8]
				if cloned := tree.CloneBranch(newID); cloned != nil {
					cloned.Save(filepath.Join(session.SessionDir(root), newID+".json"))
					fmt.Printf("Cloned: %s\n", newID)
				}
			case "/export":
				exportDir := filepath.Join(root, ".harness", "exports")
				os.MkdirAll(exportDir, 0755)
				content := tree.ExportHTML()
				ep := filepath.Join(exportDir, tree.SessionID+".html")
				os.WriteFile(ep, []byte(content), 0644)
				fmt.Printf("Exported: %s\n", ep)
			case "/checkout":
				if len(parts) >= 2 && tree.Nodes[parts[1]].NodeID != "" {
					tree.ActiveNode = parts[1]
					fmt.Printf("Checked out: %s\n", parts[1])
				}
			case "/fork":
				newID := "fork-" + hex.EncodeToString([]byte(time.Now().String()))[:8]
				if f := session.ForkSession(tree, tree.ActiveNode, newID); f != nil {
					f.Save(filepath.Join(session.SessionDir(root), newID+".json"))
					fmt.Printf("Forked: %s\n", newID)
				}
			case "/model", "/models":
				fmt.Println("\033[1m📡 Providers disponíveis:\033[0m")
				for _, name := range ac.ProviderRegistry.Names() {
					marker := " "
					if name == appCfg.ActiveProvider {
						marker = "▶"
					}
					fmt.Printf("  %s \033[36m%s\033[0m\n", marker, name)
					// Show models for this provider
					if pc, ok := appCfg.Providers[name]; ok && len(pc.Models) > 0 {
						for _, m := range pc.Models {
							currentModel := extractModelFromTemplate(pc.BodyTemplate)
							cur := " "
							if m.ID == currentModel {
								cur = "▶"
							}
							costs := ""
							if m.Cost.Input > 0 {
								costs = fmt.Sprintf(" $%.2f/M in, $%.2f/M out", m.Cost.Input, m.Cost.Output)
							}
							fmt.Printf("    %s \033[33m%s\033[0m%s\n", cur, m.ID, costs)
						}
					}
				}
				fmt.Println("\nDica: use \033[35m--provider\033[0m e \033[35m--model\033[0m na linha de comando para selecionar.")
			default:
				fmt.Printf("Unknown: %s\n", cmd)
			}
			continue
		}

		// Message queue
		if agentBusy {
			select {
			case msgQueue <- userInput:
				fmt.Printf("\033[90mQueued (queue: %d)\033[0m\n", len(msgQueue))
			default:
				fmt.Println("\033[38;5;208mQueue full.\033[0m")
			}
			continue
		}

		// Run agent
		tree.AddNode("user", userInput, len(userInput)/4)
		tree.Save(sessionFile)

		stopThinking := tc.ShowThinking(appCfg.ActiveProvider)
		if ac.StatusFn != nil {
			ac.StatusFn("starting...")
		}
		cancelCtx, cancel = context.WithTimeout(context.Background(), 3*time.Minute)

		agentBusy = true
		go func(input string) {
			defer stopThinking()
			defer cancel()
			result, err := agent.Run(cancelCtx, ac, &agent.StdLogger{}, toolReg, tree, input, mandates)

			if err != nil {
				fmt.Printf("\n\033[38;5;196mError: %v\033[0m\n", err)
			}
			if result != "" {
				tc.MessageBubble("assistant", result, len(result)/4)
				tc.Footer(tree, tree.EstimateUSD(appCfg.ActiveProvider))
			}
			tree.Save(sessionFile)

			agentBusy = false
		msgLoop:
			for len(msgQueue) > 0 {
				select {
				case queued := <-msgQueue:
					fmt.Printf("\033[90mDelivering queued...\033[0m\n")
					tree.AddNode("user", queued, len(queued)/4)
					tree.Save(sessionFile)
					r2, e2 := agent.Run(context.Background(), ac, &agent.StdLogger{}, toolReg, tree, queued, mandates)
					if e2 != nil {
						fmt.Printf("\n\033[38;5;196mError: %v\033[0m\n", e2)
					}
					if r2 != "" {
						tc.MessageBubble("assistant", r2, len(r2)/4)
						tc.Footer(tree, tree.EstimateUSD(appCfg.ActiveProvider))
					}
					tree.Save(sessionFile)
				default:
					break msgLoop
				}
			}
		}(userInput)
	}
}

func readInput(reader *bufio.Reader) (string, error) {
	var line strings.Builder
	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			return "", err
		}
		switch r {
		case 0x03: // Ctrl+C
			fmt.Print("\r\033[K")
			return "", nil
		case 0x0c: // Ctrl+L
			fmt.Print("\r\033[K")
			return "/model", nil
		case 0x1b: // Escape
			fmt.Print("\r\033[K")
			return "", nil
		case '\n':
			return line.String(), nil
		case '\r':
			continue
		default:
			line.WriteRune(r)
		}
	}
}

func expandFileRefs(input, root string) string {
	var result strings.Builder
	for _, w := range strings.Fields(input) {
		if result.Len() > 0 {
			result.WriteString(" ")
		}

		// Handle file:// URIs — keep as-is for multimodal processing
		if strings.HasPrefix(w, "file://") {
			result.WriteString(w)
			continue
		}

		// Handle @file references
		if strings.HasPrefix(w, "@") {
			ref := w[1:]
			paths := []string{filepath.Join(root, ref)}
			if filepath.IsAbs(ref) {
				paths = []string{ref}
			}
			found := false
			for _, p := range paths {
				if data, err := os.ReadFile(p); err == nil {
					result.WriteString(string(data))
					found = true
					break
				}
			}
			if !found {
				result.WriteString(w)
			}
		} else {
			result.WriteString(w)
		}
	}
	return result.String()
}




// execCmd runs a command inline (!cmd / !!cmd).
func execCmd(cmd string, root string) (string, int, error) {
	r := (&tools.ExecCommandExecutor{Root: root}).Execute("execute_command", map[string]any{"command": cmd})
	return r.Output, r.ExitCode, r.Error
}
