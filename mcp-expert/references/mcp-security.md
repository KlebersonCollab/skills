# MCP Authentication & Security

Guidelines for securing production MCP integrations.

## 1. Discoverable Auth
Clients should not require manual token configuration.
- **Metadata Documents**: Use HTTPS URLs as `client_id` that point to a JSON describing the client's capabilities and redirect URIs.
- **Discovery**: The client fetches the metadata to initiate the OAuth flow automatically.

## 2. Token Management
- **Vaults**: Prefer platform-managed vaults (e.g., Claude's Vault) to store and inject credentials.
- **Injection**: Tokens should be injected into the transport layer (Headers) rather than passed as tool arguments.
- **Rotation**: Servers should support standard refresh token flows.

## 3. Sandbox Execution (Thin Surface)
When using the "execute script" pattern:
- **Isolation**: Use `gVisor`, `Wasm`, or specialized sandboxes.
- **Resource Limits**: Enforce strict CPU, memory, and timeout limits.
- **Network Gating**: Restrict the sandbox's network access to only the target API.

## 4. Permission Gating
- **Command Risk Classification**: Categorize tools/commands as `low-risk` (auto-approve) or `high-risk` (ask user).
- **Narrow Scopes**: Only request the minimum OAuth scopes necessary for the tool's function.
