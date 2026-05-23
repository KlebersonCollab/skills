# Harness SOLID Refactor — Task List

## Phase 1: Foundation (internal packages)
- [x] **T1.1**: Create `internal/config/config.go` — AppConfig, ProviderConfig, Load/Save
- [x] **T1.2**: Create `internal/llm/provider.go` — LLMProvider interface + ProviderRegistry
- [x] **T1.3**: Create `internal/llm/call.go` — Shared HTTP helper (retry, template, headers)
- [x] **T1.4**: Create `internal/llm/gemini.go` — Gemini provider
- [x] **T1.5**: Create `internal/llm/deepseek.go` — DeepSeek provider
- [x] **T1.6**: Create `internal/llm/ollama.go` — Ollama provider
- [x] **T1.7**: Create `internal/tools/executor.go` — ToolExecutor interface + Registry
- [x] **T1.8**: Create `internal/session/tree.go` — SessionTree + Node + Skills + Discovery
- [x] **T1.9**: Create `internal/session/persistence.go` — (merged into tree.go)
- [x] **T1.10**: Create `internal/distiller/caveman.go` — ToCaveman + CheckAutoClarity
- [x] **T1.11**: Create `internal/distiller/compact.go` — CompactHistory (Compact + NodeCompactor)
- [x] **T1.12**: Create `internal/distiller/safety.go` — isDestructiveCommand (merged into caveman.go)
- [x] **T1.13**: Create `internal/mcp/transport.go` — MCPTransport (Sender/Receiver/Closer segregadas)
- [x] **T1.14**: Create `internal/mcp/server.go` — MCPServer (OCP: registry pattern)
- [x] **T1.15**: Create `internal/mcp/client.go` — MCPClient
- [x] **T1.16**: Create `internal/ui/renderer.go` — Terminal renderer (desacoplado)
- [x] **T1.17**: Create `internal/agent/parser.go` — XML tool parser (merged into loop.go)
- [x] **T1.18**: Create `internal/agent/loop.go` — AgentExecutionLoop (COMPLETO)

## Phase 2: Migration (package main → internal)
- [x] **T2.1**: Refactor `main.go` — remove runInitWizard → cmd_init.go
- [x] **T2.2**: Refactor `main.go` — interactive loop completo com chamada a agent.Run()
- [x] **T2.3**: Refactor `main.go` — command registry via switch (simplificado)
- [x] **T2.4**: Refactor `session.go` → deleted, replaced by internal/session/
- [x] **T2.5**: Refactor `client.go` → deleted, replaced by internal/llm/ + internal/config/
- [x] **T2.6**: Refactor `tools.go` → deleted, replaced by internal/tools/
- [x] **T2.7**: Refactor `mcp.go` → deleted, replaced by internal/mcp/ + cmd_mcp.go
- [x] **T2.8**: Refactor `distiller.go` → deleted, replaced by internal/distiller/
- [x] **T2.9**: Refactor `ui.go` → deleted, replaced by internal/ui/
- [x] **T2.10**: Refactor `skills.go` → deleted, merged into internal/session/tree.go
- [x] **T2.11**: Refactor `sdd.go` → deleted, merged into internal/config/sdd.go
- [x] **T2.12**: Refactor `swarm.go` → deleted, replaced by cmd_swarm.go (closure bug FIXED)

## Phase 3: Completion (core loop)
- [x] **T3.1**: Complete `agentExecutionLoop` — system prompt → LLM → parse → execute → store → iterate
- [x] **T3.2**: Complete `runInteractiveLoop` — non-slash prompts trigger agent.Run()
- [x] **T3.3**: Integrate `ToCaveman` compression in agent loop
- [x] **T3.4**: Integrate `CompactHistory` support (interface-based)
- [x] **T3.5**: Integrate `isDestructiveCommand` safety gate in agent parser

## Phase 4: Cleanup
- [x] **T4.1**: Delete old redundant files (10 files deleted)
- [x] **T4.2**: Fix `go.mod` version (1.26.1 → 1.23)
- [x] **T4.3**: Update tests (3 test files refactored)
- [x] **T4.4**: `go build ./...` → passing
- [x] **T4.5**: `go test ./... -count=1` → 27/27 passing

## Verification
- [x] `go vet ./...` → clean
- [x] Binary compiles and runs (`harness help`)
- [x] All acceptance criteria validated

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "harness-solids-refactor"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-22T16:00:00Z"
evidence_checksum: "go-test-pass-27/27"
```
