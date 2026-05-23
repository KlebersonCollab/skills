package main

import (
	"fmt"
	"os"

	"harness/internal/config"
	"harness/internal/llm"
	"harness/internal/tools"
)
func buildToolRegistry(root string, mcpServers ...config.MCPServerConfig) *tools.Registry {
	toolReg := tools.NewRegistry()

	// ── Built-in tools ──
	toolReg.Register("read_file", &tools.ReadFileExecutor{Root: root})
	toolReg.Register("write_file", &tools.WriteFileExecutor{Root: root})
	toolReg.Register("patch_file", &tools.PatchFileExecutor{Root: root})
	toolReg.Register("execute_command", &tools.ExecCommandExecutor{Root: root})
	toolReg.Register("search_files", &tools.SearchExecutor{Root: root})

	// web_search: only register if at least one API key is available
	hasPerplexity := os.Getenv("PERPLEXITY_API_KEY") != ""
	hasExa := os.Getenv("EXA_API_KEY") != ""
	hasGemini := os.Getenv("GEMINI_API_KEY") != ""
	if hasPerplexity || hasExa || hasGemini {
		webReg := llm.NewWebSearchRegistry(
			os.Getenv("PERPLEXITY_API_KEY"), os.Getenv("EXA_API_KEY"), os.Getenv("GEMINI_API_KEY"),
		)
		toolReg.Register("web_search", &tools.WebSearchExecutor{Registry: webReg})
	}

	toolReg.Register("code_search", &tools.CodeSearchExecutor{GitHubToken: os.Getenv("GITHUB_TOKEN")})
	toolReg.Register("fetch_content", &tools.FetchContentExecutor{})
	toolReg.Register("devtools", &tools.ChromeDevToolsTool{Root: root})
	toolReg.Register("paste_clipboard_image", &tools.PasteClipboardImageExecutor{})
	toolReg.Register("image_info", &tools.ImageInfoExecutor{})

	// ── MCP servers from config ──
	for _, srv := range mcpServers {
		client, err := connectMCPServer(srv)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\033[38;5;208m⚠️  MCP server '%s' connection failed: %v\033[0m\n", srv.Name, err)
			continue
		}

		count, err := registerMCPTools(toolReg, srv.Name, client)
		if err != nil {
			fmt.Fprintf(os.Stderr, "\033[38;5;208m⚠️  MCP server '%s' tool registration: %v\033[0m\n", srv.Name, err)
			continue
		}

		fmt.Printf("\033[38;5;32m✅ MCP server '%s': %d tools registradas\033[0m\n", srv.Name, count)
	}

	return toolReg
}

