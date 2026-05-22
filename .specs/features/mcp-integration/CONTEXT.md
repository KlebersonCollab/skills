
# Glossary — MCP Integration

## Domain Terms

| Term | Definition |
|------|------------|
| **MCP (Model Context Protocol)** | Open protocol by Anthropic that standardizes communication between LLMs and external tools/data sources. Defines a JSON-RPC based interface for tool discovery, invocation, and context management. |
| **MCP Server** | A process that exposes tools, resources, and prompts via the MCP protocol. Can be local (stdio) or remote (SSE/HTTP). |
| **MCP Client** | A process that connects to MCP servers to discover and invoke tools. The Harness Agent would act as an MCP client. |
| **Tool** | A capability exposed by an MCP server (e.g., read_file, execute_command). Defined by name, description, and input JSON Schema. |
| **Resource** | A data source exposed by an MCP server (e.g., file contents, database queries). Accessed via URI scheme. |
| **Prompt** | A pre-defined message template exposed by an MCP server. |
| **JSON-RPC** | The transport protocol used by MCP — lightweight remote procedure call protocol using JSON. |
| **stdio transport** | MCP transport over standard input/output — used for local subprocess communication. |
| **SSE transport** | MCP transport over Server-Sent Events — used for remote/network communication. |

## Current State

The Harness Agent currently uses direct REST API calls to LLM providers (Gemini, Ollama, DeepSeek) with hardcoded tool definitions in `main.go`. MCP integration would standardize this into a protocol-based approach.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "mcp-integration"
phase: "DISCOVERY"
status: "IN_PROGRESS"
last_update: "2026-05-23T16:25:00Z"
evidence_checksum: "discovery-glossary"
```
