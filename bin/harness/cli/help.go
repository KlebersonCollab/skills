package cli

import "fmt"

// PrintHelp prints the general usage information.
func PrintHelp() {
	fmt.Println("Harness AI Agent v2.3.0")
	fmt.Println("Usage: harness [flags] [command|prompt]")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  --stream, -s           Stream output tokens in real-time")
	fmt.Println("  --print, -p            Print mode (one-shot prompt)")
	fmt.Println("  --voice, -V            Voice mode (uses arecord + Gemini STT)")
	fmt.Println("  --voice=listen         Voice mode: listen only (no TTS)")
	fmt.Println("  --voice=speak          Voice mode: speak only (no STT)")
	fmt.Println("  --voice=full           Voice mode: full duplex (default)")
	fmt.Println("  --voice-duration=N     Recording duration in seconds (default: 8)")
	fmt.Println("  --continue, -c         Continue last session")
	fmt.Println("  --resume, -r           Resume a session interactively")
	fmt.Println("  --session=<id>         Load a specific session")
	fmt.Println("  --thinking=<level>     Set thinking level (off|minimal|low|medium|high|xhigh)")
	fmt.Println("  --provider=<name>      Select LLM provider")
	fmt.Println("  --model=<pattern>      Select model pattern")
	fmt.Println("  --fork=<id>            Fork a session")
	fmt.Println("  --offline              Disable network tools")
	fmt.Println("  --verbose              Enable verbose logging")
	fmt.Println("  --version, -v          Show version")
	fmt.Println("  --help, -h             Show this help")
	fmt.Println("  --list-models          List available providers")
	fmt.Println("  --mode rpc             Run in RPC server mode")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  init                         Interactive provider setup")
	fmt.Println("  swarm [args]                 Run swarm mode")
	fmt.Println("  mcp-server [args]            Run MCP server")
	fmt.Println("  mcp-client <transport> <addr> Run MCP client")
	fmt.Println()
	fmt.Println("Voice Commands:")
	fmt.Println("  harness --voice              Interactive voice session")
	fmt.Println("  harness -V                   (same as above)")
	fmt.Println("  harness voice full           Full duplex voice")
	fmt.Println("  harness voice listen         Listen only (STT)")
	fmt.Println("  harness voice duration=5     5-second recordings")
	fmt.Println()
	fmt.Println("  # In voice mode:")
	fmt.Println("  Press Enter to start recording")
	fmt.Println("  Speak for up to N seconds (default: 8)")
	fmt.Println("  Press Ctrl+C to quit")
	fmt.Println()
	fmt.Println("Slash Commands:")
	fmt.Println("  /help, /session, /history, /name, /model")
	fmt.Println("  /checkout, /fork, /clone, /export")
	fmt.Println()
	fmt.Println("Inline:")
	fmt.Println("  !cmd        Run command, pipe output as context")
	fmt.Println("  !!cmd       Run command, show output")
	fmt.Println("  @file       Load file contents")
	fmt.Println()
	fmt.Println("Shortcuts: Ctrl+C clear, Escape cancel, Ctrl+L models")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  harness -p \"Explain quantum computing\"")
	fmt.Println("  harness -c")
	fmt.Println("  harness --voice")
	fmt.Println("  harness voice listen")
	fmt.Println("  harness --voice=full --voice-duration=5")
	fmt.Println("  harness --thinking high")
}

// ListModels prints all available provider names.
func ListModels() {
	fmt.Println("Available providers:")
	for _, p := range []string{"gemini", "deepseek", "openai", "anthropic", "groq",
		"openrouter", "together", "fireworks", "mistral", "xai",
		"cerebras", "deepinfra", "huggingface", "ollama"} {
		fmt.Printf("  - %s\n", p)
	}
	fmt.Println("Use --provider=<name> to select, --model=<id> for specific model.")
}