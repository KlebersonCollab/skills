# Contract: Harness Distill Optimizer

## CLI Interface

| Command | Flags | Description |
| :--- | :--- | :--- |
| `hb harness distill --state` | `--feature <name>` | Compresses current feature state into STATE.md. |
| `hb harness distill --code` | `[path] [--overwrite]` | Minifies code for token reduction. |

## Expected Artifacts

### 1. `STATE.md` (Update)
```markdown
# Current State: [Feature]
- **Snapshot Date**: 2026-04-25
- **Task Progress**: 80% (8/10 completed)
- **Recent Activity**:
  - feat: implemented security scanner
  - refactor: moved logic to internal
- **Modified Files**:
  - hb/internal/security/scanner.go
  - hb/cmd/harness.go
```

### 2. `MEMORY.md` (Update)
- Records technical decisions and blockers encountered since last distill.
