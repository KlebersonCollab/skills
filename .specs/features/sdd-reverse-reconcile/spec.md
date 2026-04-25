# Specification: SDD Reverse Reconciler

## Goal
Implement a "Reverse Documentation Filling" tool in the `hb` CLI that reconciles the SDD artifacts (tasks, contracts) with the current reality of the codebase.

## Context
Often implementation happens faster than documentation updates. This tool closes the gap by detecting implemented features and marking them as done in the SDD files.

## Acceptance Criteria (AC)

### AC 1: Command `hb sdd reconcile`
- [ ] Implement the command as a subcommand of `sdd`.
- [ ] Support `--feature <name>` or all by default.

### AC 2: Evidence Discovery
- [ ] Scan for files matching feature patterns.
- [ ] Scan for CLI commands (`Use: "command"`) registered in `hb/cmd/`.
- [ ] Scan for Go functions/structs named after feature components.

### AC 3: Task Reconciliation
- [ ] If evidence of a feature component exists (file or symbol), update `tasks.md` to mark corresponding tasks as completed `[x]`.

### AC 4: Contract Generation
- [ ] If `contract.md` is missing but the implementation exists, generate a skeleton `contract.md` listing the identified interfaces and artifacts.
