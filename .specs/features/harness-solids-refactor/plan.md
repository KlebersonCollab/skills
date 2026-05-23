# Harness SOLID Refactor — Architecture Plan

## Target Architecture

```mermaid
graph TB
    subgraph CLI["cmd/harness (main.go)"]
        Root[main()] --> ParseFlags
        Root --> CmdRouter{Cmd Router}
        CmdRouter -->|"swarm"| SwarmCmd
        CmdRouter -->|"mcp-server"| MCPServerCmd
        CmdRouter -->|"init"| InitCmd
        CmdRouter -->|default| InteractiveCmd
    end

    subgraph LLM["internal/llm/"]
        ProviderI[<<interface>> LLMProvider]
        Gemini[GeminiProvider]
        DeepSeek[DeepSeekProvider]
        Ollama[OllamaProvider]
        Registry[ProviderRegistry]
        ProviderI -->|implementa| Gemini
        ProviderI -->|implementa| DeepSeek
        ProviderI -->|implementa| Ollama
        Registry o-- ProviderI
    end

    subgraph SESSION["internal/session/"]
        Tree[SessionTree]
        Node[Node]
    end

    subgraph TOOLS["internal/tools/"]
        ExecutorI[<<interface>> ToolExecutor]
        ReadFile
        WriteFile
        PatchFile
        SearchFiles
        ExecCmd
        ExecutorI --> ReadFile
        ExecutorI --> WriteFile
        ExecutorI --> PatchFile
        ExecutorI --> SearchFiles
        ExecutorI --> ExecCmd
    end

    subgraph MCP["internal/mcp/"]
        Server[MCPServer]
        Client[MCPClient]
        TransportI[<<interface>> MCPTransport]
        StdioT[StdioTransport]
        SSET[SSETransport]
        TransportI --> StdioT
        TransportI --> SSET
    end

    subgraph AGENT["internal/agent/"]
        Loop[AgentLoop]
        Parser[XMLParser]
        Loop --> Parser
    end

    subgraph DIST["internal/distiller/"]
        Caveman[ToCaveman]
        Compactor[CompactHistory]
        Safety[isDestructiveCommand]
    end

    subgraph UI["internal/ui/"]
        Renderer[TerminalRenderer]
        Spinner[Spinner]
        Board[WelcomeBoard]
    end

    InteractiveCmd -->|cria| Loop
    Loop -->|llm call| Registry
    Loop -->|tool exec| ExecutorI
    Loop -->|session| Tree
    Loop -->|compress| Caveman
    Loop -->|compact| Compactor
    Loop -->|safety| Safety
    Loop -->|render| Renderer
    SwarmCmd -->|usa| Registry
    SwarmCmd -->|usa| Loop
    MCPServerCmd -->|inicia| Server
    Server -->|registra| ExecutorI
```

## Package Structure

```
bin/harness/
├── main.go                    # Only flag parsing + routing
├── cmd/
│   ├── swarm.go               # Swarm command entry
│   ├── mcp.go                 # MCP server/client commands
│   └── init.go                # Init wizard command
├── internal/
│   ├── agent/
│   │   ├── loop.go            # AgentExecutionLoop (core)
│   │   └── parser.go          # XML tool parser
│   ├── llm/
│   │   ├── provider.go        # LLMProvider interface
│   │   ├── registry.go        # Provider registry
│   │   ├── gemini.go          # Gemini provider
│   │   ├── deepseek.go        # DeepSeek provider
│   │   ├── ollama.go          # Ollama provider
│   │   └── call.go            # Shared HTTP + retry logic
│   ├── tools/
│   │   ├── executor.go        # ToolExecutor interface
│   │   ├── read.go            # ReadFile
│   │   ├── write.go           # WriteFile
│   │   ├── patch.go           # PatchFile
│   │   ├── search.go          # SearchFiles
│   │   └── exec.go            # ExecuteCommand
│   ├── session/
│   │   ├── tree.go            # SessionTree + Node
│   │   └── persistence.go     # Save/Load
│   ├── mcp/
│   │   ├── server.go          # MCPServer
│   │   ├── client.go          # MCPClient
│   │   └── transport.go       # MCPTransport interface + impls
│   ├── distiller/
│   │   ├── caveman.go         # ToCaveman
│   │   ├── compact.go         # CompactHistory
│   │   └── safety.go          # isDestructiveCommand
│   ├── ui/
│   │   ├── renderer.go        # Terminal rendering
│   │   ├── spinner.go         # Spinner
│   │   └── board.go           # WelcomeBoard
│   └── config/
│       └── config.go          # AppConfig, ProviderConfig, Load/Save
└── go.mod
```

## Interface Contracts

### LLMProvider
```go
type LLMProvider interface {
    Name() string
    Complete(ctx context.Context, prompt string, opts map[string]any) (string, error)
    Stream(ctx context.Context, prompt string, opts map[string]any) (<-chan string, error)
}
```

### ToolExecutor
```go
type ToolExecutor interface {
    Execute(ctx context.Context, toolName string, args map[string]any) ToolResult
}
```

### MCPTransport
```go
type MCPTransport interface {
    Send(ctx context.Context, msg []byte) error
    Receive(ctx context.Context) ([]byte, error)
    Close() error
}
```

## Migration Strategy (In-Place)

To avoid breaking existing compiled binary, the refactoring will be **incremental**:

1. **Phase 1**: Create `internal/` package structure (new files, no deletion)
2. **Phase 2**: Migrate logic from `package main` files to `internal/`
3. **Phase 3**: Refactor `main.go` to use new packages
4. **Phase 4**: Complete `agentExecutionLoop`
5. **Phase 5**: Complete `runInteractiveLoop`
6. **Phase 6**: Clean up old code, fix go.mod

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "harness-solids-refactor"
phase: "SPECIFY"
status: "IN_PROGRESS"
last_update: "2026-05-22T16:05:00Z"
```
