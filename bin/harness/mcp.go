package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"harness/internal/config"
	"harness/internal/mcp"
	"harness/internal/tools"
)
func runMCPServer() {
	root := getRoot()
	server := mcp.NewServer(buildToolRegistry(root))
	fmt.Fprintf(os.Stderr, "MCP Server started\n")
	server.Start()
}

func runMCPClient(transportType, address string) {
	ctx := context.Background()
	var transport mcp.Transport
	var err error
	switch transportType {
	case "stdio":
		parts := strings.Split(address, " ")
		transport, err = mcp.NewStdioTransport(parts[0], parts[1:]...)
	case "sse":
		transport = mcp.NewSSETransport(address)
	default:
		fmt.Fprintf(os.Stderr, "Unknown transport: %s\n", transportType)
		return
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return
	}
	client := mcp.NewClient(transport)
	defer client.Close()
	client.Initialize(ctx)
	tools2, _ := client.ListTools(ctx)
	fmt.Printf("%d tools available\n", len(tools2))
	for _, t := range tools2 {
		fmt.Printf("  %s: %s\n", t.Name, t.Description)
	}
}

// connectMCPServer connects to an MCP server from config and returns a client.
func connectMCPServer(srv config.MCPServerConfig) (*mcp.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var transport mcp.Transport
	var err error
	switch srv.Transport {
	case "stdio":
		parts := strings.Fields(srv.Address)
		if len(parts) == 0 {
			return nil, fmt.Errorf("empty address for stdio transport")
		}
		transport, err = mcp.NewStdioTransport(parts[0], parts[1:]...)
	case "sse":
		transport = mcp.NewSSETransport(srv.Address)
	default:
		return nil, fmt.Errorf("unknown MCP transport: %s", srv.Transport)
	}
	if err != nil {
		return nil, fmt.Errorf("%s transport: %w", srv.Transport, err)
	}

	client := mcp.NewClient(transport)
	if err := client.Initialize(ctx); err != nil {
		transport.Close()
		return nil, fmt.Errorf("initialize %s: %w", srv.Name, err)
	}

	return client, nil
}

// registerMCPTools connects to an MCP server and registers its tools in the registry.
func registerMCPTools(toolReg *tools.Registry, name string, client *mcp.Client) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	toolDefs, err := client.ListTools(ctx)
	if err != nil {
		return 0, fmt.Errorf("list tools from %s: %w", name, err)
	}

	count := 0
	for _, t := range toolDefs {
		regName := fmt.Sprintf("mcp:%s:%s", name, t.Name)
		// Capture loop variable
		toolName := t.Name
		toolReg.Register(regName, &tools.MCPToolExecutor{
			CallFn: func(ctx context.Context, _ string, args map[string]interface{}) (string, error) {
				result, err := client.CallTool(ctx, toolName, args)
				if err != nil {
					return "", err
				}
				if result.IsError {
					var errText string
					for _, item := range result.Content {
						errText += item.Text
					}
					return "", fmt.Errorf("%s", errText)
				}
				var texts []string
				for _, item := range result.Content {
					if item.Type == "text" {
						texts = append(texts, item.Text)
					}
				}
				return strings.Join(texts, "\n"), nil
			},
			ToolName:    t.Name,
			Description: t.Description,
		})
		count++
	}

	return count, nil
}

// ── Swarm ─────────────────────────────────────────────────────────────────

