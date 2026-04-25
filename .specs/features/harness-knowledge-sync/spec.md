# Specification: Harness Knowledge Sync

## Goal
Implement a pro-active Knowledge Graph manager in the `hb` CLI that identifies architectural changes via Git analysis and ensures the `KNOWLEDGE-MAP.mermaid` reflects the current reality of the ecosystem.

## Context
The current `hb map` is reactive (it only maps what is already documented in `SKILL.md`). The `harness-knowledge-sync` tool will be proactive, detecting new files or features that haven't been linked to any skill yet.

## Acceptance Criteria (AC)

### AC 1: Command `hb harness map --auto`
- [ ] Implement the command/flag.
- [ ] Must analyze `git diff` to identify modified components (features, skills, internals).

### AC 2: Change Classification
- [ ] If a new feature is detected in `.specs/features/`, identify its dependencies.
- [ ] If a file in `hb/internal/` is changed, identify which CLI command uses it.

### AC 3: Suggestion Engine
- [ ] Compare detected changes with the current `KNOWLEDGE-MAP.mermaid`.
- [ ] If a relation is missing (e.g., a feature uses a skill but isn't linked), suggest the Mermaid code to add it.
- [ ] Provide an `--apply` flag to automatically update the Mermaid file.

### AC 4: Visualization
- [ ] Show a "Knowledge Delta" report:
    - Nodes added/modified.
    - Missing links detected.
