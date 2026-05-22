
# Project Memory — Skills Hub

## Context
This document maintains persistent context about the project's architecture, decisions, and patterns.

## Key Files & Directories
- `bin/harness/` — Go-based Harness Agent (modular structure)
- `bin/harness/main.go` — Entrypoint and core logic (33.7 KB)
- `bin/harness/client.go` — Client communication (incl. SSE streaming)
- `bin/harness/distiller.go` — Token distiller module
- `bin/harness/session.go` — Session management
- `bin/harness/tools.go` — Harness tools (read/write/search/execute commands)
- `bin/harness/harness_test.go` — Main test suite (19 tests)
- `.hub-mode` — Mode control file
- `.specs/project/STATE.md` — State machine artifact

## Architectural Decisions (captured)
- **Go as primary language** for the Harness Agent — chosen for performance and single binary distribution.
- **SSE streaming** implemented with HTTP long-polling and channel-based token delivery in `client.go`.
- **Modular CLI** in `__main__.py` (Python) orchestrates calls to the Go binary.
- **Token distiller** uses three compression levels: Lite, Full, Ultra.

## Preferences
- Portuguese (Brazil) for project communication.
- Conventional commits in English (`feat:`, `fix:`, `chore:`, etc.).
- Feature branches via `git worktree` for Medium+ tasks.

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "harness-streaming-sse"
phase: "MEMORY"
status: "COMPLETED"
last_update: "2026-05-23T16:00:00Z"
evidence_checksum: "go-test-pass-1.106s"
```
