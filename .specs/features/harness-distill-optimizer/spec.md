# Specification: Harness Distill Optimizer

## Goal
Implement an "Agentic Memory Optimizer" in the `hb` CLI to automate the persistence of state and reduce token consumption during long development sessions.

## Context
Long-running agents tend to lose focus or hit token limits. The `harness-expert` skill requires frequent state synchronization. This tool automates the "Context Compression" part of the `token-distiller` skill.

## Acceptance Criteria (AC)

### AC 1: Command `hb harness distill`
- [ ] Implement the command as a subcommand of `harness`.
- [ ] Support flags `--state` and `--code`.

### AC 2: State Distillation (`--state`)
- [ ] Must identify the current active feature (via SDD structure).
- [ ] Must gather "Evidence" of recent work:
    - Recent commits (last 3).
    - Modified files since last sync.
    - Remaining tasks in `tasks.md`.
- [ ] Must generate/update a `STATE.md` and `MEMORY.md` with a "Current Status" snapshot.

### AC 3: Code Distillation (`--code [path]`)
- [ ] Must reuse and expand `internal/distiller` logic.
- [ ] Must allow minifying a file (removing comments/whitespace) specifically for "reading" purposes (not to overwrite the source if not requested).
- [ ] Support `--overwrite` flag to actually minify the source file.

### AC 4: Integration with Rehydrate
- [ ] `hb harness rehydrate` should prioritize reading distilled `STATE.md` files created by this tool.
