---
name: mcp-expert
version: 1.0.0
description: "Model Context Protocol (MCP) Expert — guides the agent to design and build production-ready MCP servers using industry-standard patterns for integration, interaction, and security."
category: agentic-infrastructure
uses:
  - api-architect
  - architecture
  - sdd
---

# MCP Expert (v1.0.0)

> "The bridge between models and production systems." — Master of the Model Context Protocol, ensuring agents can safely and effectively interact with the real world.

---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the guidelines of the `token-distiller` skill.
1. **Context Check**: Have you rehydrated the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)?
3. **Plan Check**: Does the `plan.md` file define the MCP architecture and include **Mermaid** diagrams?

---

## Goal

Empower the agent to design and implement high-quality MCP servers that follow production-ready patterns. This skill ensures that MCP integrations are not just functional proxies, but robust product surfaces designed for agentic consumption.

---

## 🏗️ The 12 Production MCP Patterns

### Group 1: Architecture & Tooling
1.  **Remote-First Server**: Build as HTTP endpoints (SSE/Fetch) from day one to support cloud-native agents.
2.  **Intent-Grouped Tools**: Design tools around user tasks (e.g., `create_issue_from_thread`) rather than 1:1 API endpoint mapping.
3.  **Thin Surface**: For massive APIs (AWS/K8s), use a `search` + `execute` pattern in a secure sandbox.

### Group 2: Rich Interaction
4.  **Inline UI**: Return interactive views (charts, dashboards, forms) using MCP Apps instead of raw JSON.
5.  **Elicited Input**: Use **Form Mode** to pause and ask the user for missing structured data or confirmation.
6.  **External Handoff**: Use **URL Mode** for sensitive downstream flows (OAuth, payments) that should not transit the agent's context.

### Group 3: Authentication & Security
7.  **Discoverable Auth**: Publish standard OAuth metadata for zero-friction client registration and discovery.
8.  **Vault-Held Credentials**: Offload token management to platform vaults; avoid manual token passing in tool calls.

### Group 4: Context Optimization
9.  **On-Demand Tool Loading**: Use tool-search meta-tools to lazy-load definitions and save context tokens.
10. **Programmatic Tool Calling**: Process large results (filter, join, aggregate) in a client-side sandbox before showing the model.

### Group 5: Distribution & Expertise
11. **Plugin Bundle**: Package MCP servers with Skills, hooks, and subagents into a single distribution unit.
12. **Server-Distributed Skills**: Ship domain expertise (Skills) directly from the MCP server to guide the agent's behavior.

---

## Workflow (4 Phases)

### Phase 1: STRATEGY — Mapping the Interface
1.  **Identify Intent**: Group underlying API capabilities into human-centric "Intents".
2.  **Select Interaction Mode**: Determine if the tool needs **Inline UI**, **Elicited Input**, or **External Handoff**.
3.  **Auth Design**: Define how the agent will authenticate against the downstream system (Discoverable OAuth preferred).

### Phase 2: DESIGN — Schema & Patterns
1.  **Tool Definition**: Write descriptive tool definitions that are "agent-friendly" (clear descriptions, typed parameters).
2.  **Interface Design**: Design the UI components (MCP Apps) or Form schemas (Elicitation).
3.  **Architecture**: Choose between stdio (local/CLI) or HTTP (remote/production) transport.

### Phase 3: IMPLEMENT — Building the Server
1.  **SDK Selection**: Use the official MCP SDKs (TS/Python/Go).
2.  **Sandbox Enforcement**: If using **Thin Surface**, ensure the execution environment is strictly sandboxed.
3.  **Error Handling**: Return clear, actionable error messages that the agent can use for self-correction.

### Phase 4: VALIDATE — Compliance Audit
1.  **Tool Selection Test**: Verify if the agent can correctly identify and use the tools based on descriptions.
2.  **Auth Flow Audit**: Ensure sensitive data never hits the context window unnecessarily.
3.  **Performance Check**: Verify token usage of tool definitions and response payloads.

---

## Output Structure

| Artifact | Format | Description |
|----------|---------|-----------|
| **MCP Spec** | `.md` | Description of the server's tools, resources, and prompts. |
| **Tool Catalog** | Table | Mapping of user intents to specific tools and interaction modes. |
| **Auth Document** | `.json` | Discoverable OAuth metadata or credential injection plan. |
| **Interface Mockups**| `.png`/`.md`| Mockups for Inline UI or Elicited Forms. |

---

## Quality Rules

- **Intent > Endpoints**: Never build a 1:1 proxy; build a task-oriented surface.
- **Context is Gold**: Minimize tool definition sizes; use `On-Demand Loading` for >20 tools.
- **Safety First**: Destructive actions **MUST** use `Elicited Input` (confirmation) or `Plan-Validate-Execute`.
- **Descriptive Excellence**: Tool descriptions must be "pushy" and clear about *when* to use them.

## Prohibited

- NEVER pass raw API keys in tool parameters.
- NEVER return multi-megabyte JSON payloads directly to the model (use `Programmatic Calling`).
- NEVER use stdio for servers intended for web or mobile agent clients.

---

## References

### 📖 Local Guides
- [`references/mcp-patterns.md`](references/mcp-patterns.md) — Detailed breakdown of the 12 patterns.
- [`references/interaction-modes.md`](references/interaction-modes.md) — Guide to Inline UI, Elicitation, and Handoff.
- [`references/mcp-security.md`](references/mcp-security.md) — Auth and Sandboxing best practices.

### 🌐 External Resources
- [12 MCP Patterns Behind Production Agents](https://substack.com/home/post/p-195426893)
- [Anthropic: Building agents that reach production systems with MCP](https://claude.com/blog/building-agents-that-reach-production-systems-with-mcp)
- [MCP Specification](https://modelcontextprotocol.io/specification/2025-11-25/)
