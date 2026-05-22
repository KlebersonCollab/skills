# Spec - Harness Transformation (Phase 1)

## Goal
Transform the passive Skills Hub into an active terminal-based agent execution harness written in Go (`bin/harness`). The engine will run an interactive loop with slash commands, load domain skills as dynamic prompt contexts, execute local tool subprocesses (`read`, `write`, `patch`, `bash`), and support configurable LLM providers via dynamic HTTP JSON templates.

---

## Acceptance Criteria

### AC-1: Terminal Interactive Console Loop (Slash Commands)
* **Given** the workspace root.
* **When** executing `./bin/harness` or running the compiled harness binary.
* **Then** the terminal boots into an interactive console prompt loop (`harness [node-0]> `) supporting:
  * Typing normal text: Sends a user message to the active node in the session tree.
  * `/history`: Displays the tree history of messages.
  * `/checkout <node_id>`: Switches the active node to the specified ID, enabling non-linear conversation branching.
  * `/sessions`: Lists local active sessions.
  * `/skills`: Lists available skills in the hub.
  * `/exit` or `/quit`: Cleanly exits the harness.

### AC-2: Minimal Tool Set (Subprocesses)
* **Given** an active agent execution loop inside the harness.
* **When** the agent issues a tool call.
* **Then** the harness executes the action using its built-in minimal tools:
  * `read_file(path)`: Safe, bounded file reading.
  * `write_file(path, content)`: Writes or completely overwrites a file.
  * `patch_file(path, target, replacement)`: Surgically replaces a contiguous block of text.
  * `execute_command(command)`: Safely executes a shell script or program via subprocess (`exec.Command` in Go), tracking exit codes and console outputs.

### AC-3: Dynamic Provider Templates (Configurable HTTP REST API)
* **Given** a new LLM provider (Ollama, Gemini, local proxy).
* **When** configured in `.harness/config.json`.
* **Then** the harness parses the JSON structure defining:
  * `url`: The endpoint API endpoint.
  * `headers`: Header key-value pairs (for authorization, content-types, etc.).
  * `body_template`: The request body structure, replacing keywords like `{{prompt}}` or `{{history}}`.
  * `response_path`: The JSON path to extract the assistant response string.

### AC-4: SDD Gated Automation
* **Given** the active project state.
* **When** running a command that changes files.
* **Then** the harness verifies the current phase in `.specs/project/STATE.md` and restricts the agent from jumping straight to implementation without completing align/specify/plan steps.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HARNESS-TRANSFORMATION"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-22T14:48:00Z"
evidence_checksum: "go-test-pass-005"
```
