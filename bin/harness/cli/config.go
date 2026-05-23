// Package cli provides CLI argument parsing and help for the harness agent.
package cli

import (
	"fmt"
	"os"
	"strings"
)

// Config holds all parsed CLI flags and arguments.
type Config struct {
	StreamMode    bool
	PrintMode     bool
	RPCMode       bool
	VoiceMode     bool
	VoiceListen   bool
	VoiceSpeak    bool
	VoiceDuration int
	ThinkingLevel string
	ProviderName  string
	ModelPattern  string
	SessionID     string
	SessionDir    string
	ForkID        string
	PromptArgs    []string
	Command       string // "swarm", "mcp-server", "mcp-client", "init", "help", "voice"
	SubCommandIdx int    // index in PromptArgs where subcommand args begin
}

// Parse parses CLI arguments and populates a Config.
// The boolean return indicates whether the program should continue (true) or exit (false).
// When false, the caller MUST os.Exit or return immediately.
func Parse(args []string) (*Config, bool) {
	cfg := &Config{}

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch {
		case arg == "--stream" || arg == "-s":
			cfg.StreamMode = true
		case arg == "--print" || arg == "-p":
			cfg.PrintMode = true
		case arg == "--mode" && i+1 < len(args):
			i++
			if args[i] == "rpc" {
				cfg.RPCMode = true
			}
		case arg == "--voice" || arg == "-V":
			cfg.VoiceMode = true
		case strings.HasPrefix(arg, "--voice="):
			cfg.VoiceMode = true
			modeOpt := strings.TrimPrefix(arg, "--voice=")
			switch modeOpt {
			case "listen":
				cfg.VoiceListen = true
			case "speak":
				cfg.VoiceSpeak = true
			case "full":
				cfg.VoiceListen = true
				cfg.VoiceSpeak = true
			}
		case arg == "--voice-duration" && i+1 < len(args):
			i++
			fmt.Sscanf(args[i], "%d", &cfg.VoiceDuration)
		case strings.HasPrefix(arg, "--voice-duration="):
			fmt.Sscanf(strings.TrimPrefix(arg, "--voice-duration="), "%d", &cfg.VoiceDuration)
		case arg == "--continue" || arg == "-c":
			cfg.SessionID = "__continue__"
		case arg == "--resume" || arg == "-r":
			cfg.SessionID = "__resume__"
		case strings.HasPrefix(arg, "--session="):
			cfg.SessionID = strings.TrimPrefix(arg, "--session=")
		case arg == "--session" && i+1 < len(args):
			i++
			cfg.SessionID = args[i]
		case strings.HasPrefix(arg, "--session-dir="):
			cfg.SessionDir = strings.TrimPrefix(arg, "--session-dir=")
		case arg == "--session-dir" && i+1 < len(args):
			i++
			cfg.SessionDir = args[i]
		case strings.HasPrefix(arg, "--thinking="):
			cfg.ThinkingLevel = strings.TrimPrefix(arg, "--thinking=")
		case arg == "--thinking" && i+1 < len(args):
			i++
			cfg.ThinkingLevel = args[i]
		case strings.HasPrefix(arg, "--provider="):
			cfg.ProviderName = strings.TrimPrefix(arg, "--provider=")
		case arg == "--provider" && i+1 < len(args):
			i++
			cfg.ProviderName = args[i]
		case strings.HasPrefix(arg, "--model="):
			cfg.ModelPattern = strings.TrimPrefix(arg, "--model=")
		case arg == "--model" && i+1 < len(args):
			i++
			cfg.ModelPattern = args[i]
		case strings.HasPrefix(arg, "--fork="):
			cfg.ForkID = strings.TrimPrefix(arg, "--fork=")
		case arg == "--fork" && i+1 < len(args):
			i++
			cfg.ForkID = args[i]
		case arg == "--offline":
			os.Setenv("HARNESS_OFFLINE", "1")
		case arg == "--verbose":
			os.Setenv("HARNESS_VERBOSE", "1")
		case arg == "--list-models":
			ListModels()
			return cfg, false
		case arg == "--help" || arg == "-h":
			PrintHelp()
			return cfg, false
		case arg == "--version" || arg == "-v":
			fmt.Println("Harness AI Agent v2.3.0")
			return cfg, false
		case strings.HasPrefix(arg, "-"):
			fmt.Printf("Flag desconhecida: %s\n", arg)
			os.Exit(1)
		default:
			cfg.PromptArgs = append(cfg.PromptArgs, arg)
		}
	}

	// Detect top-level commands among prompt args.
	if len(cfg.PromptArgs) > 0 && !cfg.PrintMode && cfg.SessionID == "" && cfg.ForkID == "" {
		cmd := cfg.PromptArgs[0]
		switch cmd {
		case "swarm", "mcp-server", "mcp-client", "init", "help":
			cfg.Command = cmd
			cfg.SubCommandIdx = 1
		case "voice":
			cfg.Command = "voice"
			cfg.VoiceMode = true
			cfg.SubCommandIdx = 1
			// Parse sub-flags: voice full | voice listen | voice speak | voice duration=5
			for _, opt := range cfg.PromptArgs[1:] {
				switch opt {
				case "full":
					cfg.VoiceListen = true
					cfg.VoiceSpeak = true
				case "listen":
					cfg.VoiceListen = true
				case "speak":
					cfg.VoiceSpeak = true
				default:
					if strings.HasPrefix(opt, "duration=") {
						fmt.Sscanf(strings.TrimPrefix(opt, "duration="), "%d", &cfg.VoiceDuration)
					}
				}
			}
		}
	}

	// Default voice mode: full duplex (listen + speak)
	if cfg.VoiceMode {
		if !cfg.VoiceListen && !cfg.VoiceSpeak {
			cfg.VoiceListen = true
			cfg.VoiceSpeak = true
		}
		if cfg.VoiceDuration <= 0 {
			cfg.VoiceDuration = 8
		}
		_ = cfg.VoiceDuration // suppress unused warning (used at runtime)
	}

	// Set env vars from parsed flags.
	if cfg.ThinkingLevel != "" {
		os.Setenv("HARNESS_THINKING_LEVEL", cfg.ThinkingLevel)
	}
	if cfg.ProviderName != "" {
		os.Setenv("HARNESS_PROVIDER", cfg.ProviderName)
	}
	if cfg.ModelPattern != "" {
		os.Setenv("HARNESS_MODEL", cfg.ModelPattern)
	}

	return cfg, true
}