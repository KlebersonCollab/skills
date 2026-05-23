package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"harness/cli"
	"harness/internal/agent"
	"harness/internal/session"
	"harness/internal/voice"
)

func main() {
	cfg, cont := cli.Parse(os.Args[1:])
	if !cont {
		return
	}

	if cfg.RPCMode {
		runRPCMode()
		return
	}

	if cfg.Command != "" {
		switch cfg.Command {
		case "swarm":
			runSwarm()
			return
		case "mcp-server":
			runMCPServer()
			return
		case "mcp-client":
			args := cfg.PromptArgs[cfg.SubCommandIdx:]
			if len(args) < 2 {
				fmt.Println("Usage: harness mcp-client <transport> <address>")
				return
			}
			runMCPClient(args[0], strings.Join(args[1:], " "))
			return
		case "init":
			runInitWizard()
			return
		case "voice":
			runVoiceMode(cfg)
			return
		case "help":
			cli.PrintHelp()
			return
		}
	}

	// Voice mode (from --voice flag)
	if cfg.VoiceMode {
		runVoiceMode(cfg)
		return
	}

	if cfg.PrintMode || len(cfg.PromptArgs) > 0 {
		runPrintMode(cfg.StreamMode, cfg.PromptArgs)
		return
	}

	if cfg.ForkID != "" {
		runForkMode(cfg.ForkID, cfg.StreamMode, cfg.SessionDir)
		return
	}

	if cfg.SessionID != "" {
		runWithSession(cfg.SessionID, cfg.StreamMode, cfg.PromptArgs)
		return
	}

	runInteractiveLoop(cfg.StreamMode)
}

// runVoiceMode starts the interactive voice session using the harness agent.
func runVoiceMode(cfg *cli.Config) {
	root := getRoot()

	// Verify audio dependencies
	missingTTS := voice.CheckDeps()
	if len(missingTTS) > 0 {
		voice.PrintDepsWarning(missingTTS)
	}

	// Check API key for Gemini STT
	hasGeminiKey := true
	if os.Getenv("GEMINI_API_KEY") == "" {
		fmt.Fprintln(os.Stderr, "\033[1;30m⚠️  GEMINI_API_KEY not set. Voice STT will fall back to manual text input.\033[0m")
		hasGeminiKey = false
	}

	// Build voice config
	vc := voice.DefaultConfig()
	vc.Duration = cfg.VoiceDuration
	if vc.Duration <= 0 {
		vc.Duration = 8
	}
	vc.Continuous = true

	if cfg.VoiceListen && cfg.VoiceSpeak {
		vc.Mode = voice.ModeFull
	} else if cfg.VoiceListen {
		vc.Mode = voice.ModeListen
	} else if cfg.VoiceSpeak {
		vc.Mode = voice.ModeSpeak
	}

	// ── Welcome screen ──────────────────────────────────────────────
	fmt.Printf("\n%s%s🎤 HARNESS VOICE MODE%s\n", "\033[1;36m", "\033[1m", "\033[0m")
	fmt.Printf("%s  Record:   %ds%s\n", "\033[32m", vc.Duration, "\033[0m")
	fmt.Printf("%s  Mode:     %s%s\n", "\033[33m", vc.Mode.String(), "\033[0m")
	fmt.Printf("%s  STT:      %sGemini%s\n", "\033[34m", "\033[0m", "\033[0m")
	fmt.Printf("%s  Temp:     %s%s\n", "\033[90m", filepath.Join(root, ".harness", "voice"), "\033[0m")
	fmt.Println()

	vs := voice.NewSession(root, vc)

	var text string
	var err error

	if hasGeminiKey {
		// ── First recording ─────────────────────────────────────────────
		fmt.Print("\033[1;35m🎤 Press Enter to record (type '/t' for text fallback, Ctrl+C to quit):\033[0m ")
		var confirm string
		fmt.Scanln(&confirm)
		confirm = strings.TrimSpace(confirm)

		if confirm == "/t" || confirm == "/text" {
			fmt.Print("\033[1;33m⌨️  Digite sua mensagem:\033[0m ")
			reader := bufio.NewReader(os.Stdin)
			text, _ = reader.ReadString('\n')
			text = strings.TrimSpace(text)
		} else {
			text, err = vs.OneShot()
			if err != nil {
				fmt.Fprintf(os.Stderr, "\033[38;5;208m⚠️  Voice STT error: %v\033[0m\n", err)
				fmt.Print("\033[1;33m👉 Fallback (Keyboard): Digite sua mensagem:\033[0m ")
				reader := bufio.NewReader(os.Stdin)
				text, _ = reader.ReadString('\n')
				text = strings.TrimSpace(text)
			}
		}
	} else {
		// No API key - manual text input
		fmt.Print("\033[1;33m⌨️  Digite sua mensagem:\033[0m ")
		reader := bufio.NewReader(os.Stdin)
		text, _ = reader.ReadString('\n')
		text = strings.TrimSpace(text)
	}

	if text == "" {
		text = "Olá"
	}
	fmt.Printf("\033[1;32m🎤 You:\033[0m %s\n\n", text)

	// ── Process through agent ───────────────────────────────────────
	fmt.Println("\033[1;34m🤖 Processing...\033[0m")
	start := time.Now()

	cfgApp := loadAppCfg(root)
	if cfgApp == nil {
		os.Exit(1)
	}
	reg := buildProviderReg(cfgApp)
	if reg == nil {
		os.Exit(1)
	}
	toolReg := buildToolRegistry(root, cfgApp.MCPServers...)

	ac := agent.DefaultConfig()
	ac.StreamMode = false
	ac.WorkspaceRoot = root
	ac.ActiveSkills = session.DiscoverSkills(root)
	ac.ProviderRegistry = reg
	ac.ActiveProviderName = reg.Names()[0]

	tree := session.NewTree("voice-"+time.Now().Format("150405"), text)
	tree.AddNode("user", text, len(text)/4)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	result, err := agent.Run(ctx, ac, &agent.StdLogger{}, toolReg, tree, text, loadMandates(root))
	cancel()
	elapsed := time.Since(start)

	if err != nil {
		fmt.Fprintf(os.Stderr, "\033[38;5;196mError: %v\033[0m\n", err)
		return
	}

	if result != "" {
		fmt.Printf("\033[1;34m🤖 AI (%s):\033[0m\n\n%s\n\n", formatDuration(elapsed), result)

		// Speak response
		if vc.Mode == voice.ModeFull || vc.Mode == voice.ModeSpeak {
			if err := vs.Speak(result); err != nil {
				fmt.Fprintf(os.Stderr, "\033[38;5;208mTTS: %v\033[0m\n", err)
			}
		}

		// ── Continuous loop ─────────────────────────────────────────
		fmt.Println("\033[1;35mContinuous mode. Press Enter to record, Ctrl+C to quit.\033[0m")
		vs.InteractiveLoop(func(input string) (string, error) {
			tree.AddNode("user", input, len(input)/4)
			ctx2, cancel2 := context.WithTimeout(context.Background(), 3*time.Minute)
			defer cancel2()
			resp, err := agent.Run(ctx2, ac, &agent.StdLogger{}, toolReg, tree, input, loadMandates(root))
			if err != nil {
				return "", err
			}
			tree.AddNode("assistant", resp, len(resp)/4)
			return resp, nil
		})
	}

	fmt.Println("\n\033[1;36mVoice session ended.\033[0m")
}

func formatDuration(d time.Duration) string {
	if d < time.Second {
		return fmt.Sprintf("%dms", d.Milliseconds())
	}
	if d < time.Minute {
		return fmt.Sprintf("%.1fs", d.Seconds())
	}
	return fmt.Sprintf("%dm%ds", int(d.Minutes()), int(d.Seconds())%60)
}