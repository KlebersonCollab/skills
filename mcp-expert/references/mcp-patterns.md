# MCP Production Patterns

This document details the 12 recurring design patterns for production-grade Model Context Protocol (MCP) integrations.

## 1. Remote-First Server
Build the MCP server as a remote HTTP endpoint (using SSE or Fetch) from day one.
- **Why**: Most production agent traffic is moving to the cloud/browser. stdio only works for local child processes.
- **Trade-off**: Requires handling network latency, reliability, and public authentication.

## 2. Intent-Grouped Tools
Group tools around user tasks/intents rather than mapping 1:1 to API endpoints.
- **Why**: Reduces context usage and round-trips. Agents think in tasks, not primitives.
- **Example**: `create_issue_from_thread` instead of `get_messages` + `parse` + `post_issue`.

## 3. Thin Surface
Expose a thin interface with `search` and `execute` tools for massive APIs (e.g., AWS, Cloudflare).
- **Why**: Prevents context window overflow and allows the agent to use the full expressive power of the API via scripts.
- **Security**: Requires a secure, server-side sandbox for execution.

## 4. Inline UI
Return interactive interfaces (charts, forms, dashboards) using **MCP Apps**.
- **Why**: Some data is better inspected visually than described in prose.
- **Implementation**: The client renders components defined by the server.

## 5. Elicited Input (Form Mode)
Pause the tool call to ask the user for missing structured information via a native form.
- **Why**: Keeps the user in the flow instead of asking them to "restart the chat with more info".
- **Use Case**: Destructive confirmations, region selection, or missing IDs.

## 6. External Handoff (URL Mode)
Hand the user to a browser URL for sensitive flows (OAuth, payments).
- **Why**: Prevents sensitive credentials or PII from transiting through the LLM context.
- **Use Case**: Third-party OAuth, Stripe payments.

## 7. Discoverable Auth
Publish standard OAuth metadata so the client can discover how to authenticate.
- **Why**: Reduces friction at "install time" (no more pasting tokens manually).
- **Mechanism**: Use `Client ID Metadata Documents`.

## 8. Vault-Held Credentials
Offload token lifecycle (refresh/rotation) to a platform vault (e.g., Claude Managed Agents).
- **Why**: Simplifies server logic and improves security by keeping secrets out of the tool parameters.

## 9. On-Demand Tool Loading
Lazy-load tool definitions through a tool-search meta-tool.
- **Why**: Saves context tokens when a server has many tools.
- **Requirement**: Excellent tool descriptions for accurate search results.

## 10. Programmatic Tool Calling
Process large tool results (filter, join, aggregate) in a code sandbox before showing them to the model.
- **Why**: Prevents context window overflow and saves tokens on raw data processing.

## 11. Plugin Bundle
Bundle MCP servers with Skills, hooks, and subagents into a single distribution unit.
- **Why**: Prevents configuration drift and ensures all necessary expertise ships with the capability.

## 12. Server-Distributed Skills
The MCP server itself provides the procedural knowledge (Skills) on how to use its tools.
- **Why**: The provider knows the best workflows and safe sequences for their API.
