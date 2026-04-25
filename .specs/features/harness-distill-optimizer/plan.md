# Implementation Plan: Harness Distill Optimizer

## Architecture

### 1. CLI Layer
- Subcommand `distill` in `hb/cmd/harness.go`.
- Flags: `--state`, `--code`, `--feature [name]`.

### 2. Domain Layer (`hb/internal/harness/distill.go`)
- `ExecuteDistillState(feature string) error`
- `ExecuteDistillCode(path string, overwrite bool) error`

### 3. Integration Logic
- **State Collection**:
    - Use `git` commands (exec) to get diff summaries.
    - Parse `tasks.md` to count [x] vs [ ].
- **Persistence**:
    - Write to `.specs/features/<feature>/STATE.md`.

## Data Gathering (State)
- **Diff Summary**: `git diff --stat`
- **Recent Logs**: `git log -n 3 --oneline`
- **Task Progress**: Count completed/total tasks.

## Mermaid Diagram
```mermaid
graph TD
    CLI[hb harness distill] --> Router{Flag?}
    Router -- --state --> StateEngine[State Distiller]
    Router -- --code --> CodeEngine[Code Distiller]
    
    StateEngine --> Git[Git Diffs/Logs]
    StateEngine --> SDD[Tasks.md Parser]
    StateEngine --> Write[Write STATE.md / MEMORY.md]
    
    CodeEngine --> Minifier[Internal Distiller]
    Minifier --> File[Optimized File Content]
```
