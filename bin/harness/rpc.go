package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"harness/internal/agent"
	"harness/internal/session"
)

func runRPCMode() {
	root := getRoot()
	cfg := loadAppCfg(root)
	if cfg == nil {
		os.Exit(1)
	}
	reg := buildProviderReg(cfg)
	if reg == nil {
		os.Exit(1)
	}
	toolReg := buildToolRegistry(root, cfg.MCPServers...)
	tree := session.NewTree("rpc-"+time.Now().Format("150405"), "RPC")
	ac := agent.DefaultConfig()
	ac.WorkspaceRoot = root
	ac.ProviderRegistry = reg
	ac.ActiveProviderName = reg.Names()[0]
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var cmd struct {
			Type    string `json:"type"`
			ID      string `json:"id,omitempty"`
			Message string `json:"message,omitempty"`
		}
		if err := json.Unmarshal([]byte(line), &cmd); err != nil {
			resp, _ := json.Marshal(map[string]interface{}{
				"type": "response", "command": "parse", "success": false, "error": err.Error(),
			})
			fmt.Println(string(resp))
			continue
		}
		switch cmd.Type {
		case "prompt":
			ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
			tree.AddNode("user", cmd.Message, len(cmd.Message)/4)
			result, err := agent.Run(ctx, ac, &agent.StdLogger{}, toolReg, tree, cmd.Message, "")
			cancel()
			emitEvent(map[string]interface{}{"type": "agent_end", "messages": []map[string]string{{"role": "assistant", "content": result}}})
			resp, _ := json.Marshal(map[string]interface{}{"type": "response", "command": "prompt", "success": true, "id": cmd.ID})
			fmt.Println(string(resp))
			_ = err
		case "get_state":
			resp, _ := json.Marshal(map[string]interface{}{
				"type": "response", "command": "get_state", "success": true, "id": cmd.ID,
				"data": map[string]interface{}{"sessionId": tree.SessionID, "messageCount": len(tree.Nodes)},
			})
			fmt.Println(string(resp))
		case "abort":
			resp, _ := json.Marshal(map[string]interface{}{"type": "response", "command": "abort", "success": true, "id": cmd.ID})
			fmt.Println(string(resp))
		default:
			resp, _ := json.Marshal(map[string]interface{}{
				"type": "response", "command": cmd.Type, "success": false, "id": cmd.ID,
				"error": fmt.Sprintf("unknown: %s", cmd.Type),
			})
			fmt.Println(string(resp))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "RPC scanner error: %v\n", err)
	}
}

func emitEvent(v interface{}) {
	data, _ := json.Marshal(v)
	fmt.Println(string(data))
}

