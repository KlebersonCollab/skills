package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"harness/internal/agent"
	"harness/internal/session"
)

func runPrintMode(streamMode bool, promptArgs []string) {
	prompt := strings.Join(promptArgs, " ")
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		stdinData, _ := io.ReadAll(os.Stdin)
		if len(stdinData) > 0 {
			prompt = string(stdinData) + "\n\n" + prompt
		}
	}
	if prompt == "" {
		fmt.Fprintln(os.Stderr, "No prompt. Use: harness -p \"prompt\"")
		os.Exit(1)
	}
	root := getRoot()
	cfg := loadAppCfg(root)
	if cfg == nil {
		return
	}
	reg := buildProviderReg(cfg)
	if reg == nil {
		return
	}
	toolReg := buildToolRegistry(root, cfg.MCPServers...)
	tree := session.NewTree("print-"+time.Now().Format("150405"), prompt)
	tree.AddNode("user", prompt, len(prompt)/4)
	ac := agent.DefaultConfig()
	ac.StreamMode = streamMode
	ac.WorkspaceRoot = root
	ac.ProviderRegistry = reg
	ac.ActiveProviderName = reg.Names()[0]
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	result, err := agent.Run(ctx, ac, &agent.StdLogger{}, toolReg, tree, prompt, loadMandates(root))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(result)
}

