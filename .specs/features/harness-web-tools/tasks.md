# Harness Web Tools — Task List

## Phase 1: Web Search Provider Layer
- [x] **T1.1**: Create `internal/llm/web.go` — WebSearchProvider interface + WebSearchRegistry with fallback
- [x] **T1.2**: Implement Perplexity provider (`POST https://api.perplexity.ai/chat/completions`)
- [x] **T1.3**: Implement Exa provider (`POST https://api.exa.ai/search`)
- [x] **T1.4**: Implement Gemini Web provider (via Google Search grounding API)
- [x] **T1.5**: Implement fallback chain (Perplexity → Exa → Gemini)
- [x] **T1.6**: Tools for web search (included in tool integration)

## Phase 2: Tool Executors
- [x] **T2.1**: Create `internal/tools/websearch.go` — WebSearch executor (XML → provider → response)
- [x] **T2.2**: Create `internal/tools/codesearch.go` — CodeSearch executor (GitHub API + Stack Exchange API)
- [x] **T2.3**: Create `internal/tools/fetch.go` — FetchContent executor (URL fetch + YouTube + GitHub + video)

## Phase 3: Integration
- [x] **T3.1**: Register new tools in `internal/agent/loop.go` — XML parser (webSearchRE, codeSearchRE, fetchContentRE)
- [x] **T3.2**: Register new tools in `cmd_mcp.go` — MCP server
- [x] **T3.3**: Add `parseGenericAttrs` — generic XML attribute parser for complex JSON types
- [x] **T3.4**: Wire API keys from env vars in `main.go` and `cmd_mcp.go`

## Phase 4: Validation
- [x] **T4.1**: `go build ./...` passing
- [x] **T4.2**: `go test ./... -count=1` passing (27/27)
- [x] **T4.3**: All 3 web tools registered in tool registry + agent parser + MCP server
- [x] **T4.4**: Update STATE.md

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "harness-web-tools"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-22T16:30:00Z"
evidence_checksum: "go-build-pass-go-test-27"
```
