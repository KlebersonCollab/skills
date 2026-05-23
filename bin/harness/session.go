package main

import (
	"bufio"
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"harness/internal/agent"
	"harness/internal/session"
	"harness/internal/tui"
)

func runWithSession(sessionFlag string, streamMode bool, promptArgs []string) {
	root := getRoot()
	cfg := loadAppCfg(root)
	if cfg == nil {
		return
	}
	var sessionFile string
	switch sessionFlag {
	case "__continue__":
		recent := session.FindMostRecentSession(root)
		if recent == nil {
			fmt.Println("No previous session found.")
			return
		}
		sessionFile = recent.File
		fmt.Printf("Continuing: %s\n", recent.ID)
	case "__resume__":
		sessions := session.FindSessions(root)
		if len(sessions) == 0 {
			fmt.Println("No sessions found.")
			return
		}
		fmt.Println("Select session:")
		for i, s := range sessions {
			fmt.Printf("  %d) [%s] %s (%d nodes)\n", i+1, s.ID, s.Task, s.Nodes)
		}
		fmt.Print("Number: ")
		reader := bufio.NewReader(os.Stdin)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			return
		}
		idx := 0
		fmt.Sscanf(line, "%d", &idx)
		if idx < 1 || idx > len(sessions) {
			fmt.Println("Invalid.")
			return
		}
		sessionFile = sessions[idx-1].File
	default:
		if filepath.IsAbs(sessionFlag) {
			sessionFile = sessionFlag
		} else {
			sessionFile = filepath.Join(session.SessionDir(root), sessionFlag+".json")
		}
	}
	tree, err := session.LoadTree(sessionFile)
	if err != nil {
		fmt.Printf("Error loading session: %v\n", err)
		return
	}
	fmt.Printf("Loaded: %s (%d nodes)\n", tree.SessionID, len(tree.Nodes))
	reg := buildProviderReg(cfg)
	if reg == nil {
		return
	}
	toolReg := buildToolRegistry(root, cfg.MCPServers...)
	skillDefs := session.DiscoverSkills(root)
	tc := tui.New()
	ac := agent.DefaultConfig()
	ac.StreamMode = streamMode
	ac.WorkspaceRoot = root
	ac.ActiveSkills = skillDefs
	ac.ProviderRegistry = reg
	ac.ActiveProviderName = reg.Names()[0]
	if len(promptArgs) > 0 {
		userInput := strings.Join(promptArgs, " ")
		tree.AddNode("user", userInput, len(userInput)/4)
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
		defer cancel()
		result, err := agent.Run(ctx, ac, &agent.StdLogger{}, toolReg, tree, userInput, loadMandates(root))
		if err != nil {
			fmt.Printf("\nError: %v\n", err)
		}
		if result != "" {
			tc.MessageBubble("assistant", result, len(result)/4)
		}
		tree.Save(sessionFile)
		return
	}
	modelName := extractCurrentModel(cfg)
	tc.WelcomeHeader(cfg.ActiveProvider, modelName, tree.SessionID, tree.RootTask, len(skillDefs))
	runInteractivePrompt(tree, sessionFile, cfg, ac, toolReg, tc, loadMandates(root), root)
}


func runForkMode(forkID string, streamMode bool, sessionDir string) {
	root := getRoot()
	if sessionDir == "" {
		sessionDir = session.SessionDir(root)
	}
	sourceFile := filepath.Join(sessionDir, forkID+".json")
	source, err := session.LoadTree(sourceFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	newID := "fork-" + hex.EncodeToString([]byte(time.Now().String()))[:8]
	forked := session.ForkSession(source, source.ActiveNode, newID)
	if forked == nil {
		fmt.Println("Fork failed.")
		return
	}
	forkFile := filepath.Join(sessionDir, newID+".json")
	forked.Save(forkFile)
	fmt.Printf("Forked: %s -> %s\n", source.SessionID, forkFile)
}

// ── Helpers ──────────────────────────────────────────────────────────────

