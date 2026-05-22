
# Specification — MCP (Model Context Protocol) Integration

## ID: `mcp-integration`

## Objective
Integrate the Model Context Protocol (MCP) into the Harness Go Agent, enabling:
1. **MCP Client Mode**: Connect to external MCP servers to discover and invoke tools/resources.
2. **MCP Server Mode**: Expose Harness Agent's native tools (read_file, write_file, etc.) as MCP tools for other agents/clients.
3. **Unified Context**: Replace ad-hoc tool definitions with standardized MCP tool schemas.

## Functional Requirements

### FR-1: MCP Client Connection
- The Harness Agent MUST be able to connect to MCP servers via:
  - `stdio` transport (local subprocess)
  - `SSE` transport (remote HTTP endpoint)
- Connection configuration via `.harness/config.json`

### FR-2: Tool Discovery & Invocation
- On connection, the agent MUST discover available tools from the MCP server.
- Tools MUST be invocable via a unified interface (e.g., `/mcp:tool_name args...` or automatically routed by the LLM).
- Tool results MUST be returned in a structured format.

### FR-3: MCP Server Mode
- The Harness Agent MUST expose its native tools as MCP tools:
  - `read_file(path)` → Read file contents
  - `write_file(path, content)` → Write/create file
  - `patch_file(path, target, replacement)` → Surgical file patch
  - `search_files(pattern, query)` → Recursive file search
  - `execute_command(command)` → Execute bash command
- Server mode via `stdio` transport (started with `harness mcp-server`).

### FR-4: Context Bridging
- The MCP client MUST integrate with the existing session DAG (`session.go`):
  - Tool calls via MCP are logged as nodes in the session tree.
  - Results are stored and traceable.

### FR-5: Configuration
- New section in `.harness/config.json` for MCP:
  ```json
  {
    "mcp": {
      "servers": [
        { "name": "local-fs", "transport": "stdio", "command": "npx", "args": ["@modelcontextprotocol/server-filesystem", "/path"] },
        { "name": "github", "transport": "sse", "url": "http://localhost:3001/sse" }
      ],
      "server_mode": { "enabled": true, "port": 9090 }
    }
  }
  ```

## Non-Functional Requirements
- **Zero external dependencies** for core MCP implementation (Go stdlib only).
- **Performance**: Tool invocation via MCP should add <50ms overhead.
- **Graceful degradation**: If MCP server is unavailable, agent continues with native tool execution.

## Acceptance Criteria

- [ ] **AC-1**: `harness mcp-server` starts an MCP server on stdio exposing all 5 native tools.
- [ ] **AC-2**: `harness mcp-client` connects to a remote MCP server and lists available tools.
- [ ] **AC-3**: LLM can invoke MCP tools through the Harness agent (automatic routing).
- [ ] **AC-4**: Session DAG logs MCP tool calls with full input/output traceability.
- [ ] **AC-5**: Configuration via `.harness/config.json` `mcp` section.
- [ ] **AC-6**: `go build` succeeds with zero new external dependencies.
- [ ] **AC-7**: `go test ./...` passes with existing test suite intact.
- [ ] **AC-8**: `make audit` reports 100% compliance.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "mcp-integration"
phase: "SPECIFY"
status: "IN_PROGRESS"
last_update: "2026-05-23T16:30:00Z"
evidence_checksum: "spec-draft-v1"
```
