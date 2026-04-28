# Tasks: Structured Review (hb review)

## Phase 1: Git & Diff Infrastructure
- [x] T1: Implement Git Diff helper to capture staged changes. `id: REV-001`
- [x] T2: Create a simple diff parser for unified hunks in Go. `id: REV-002`

## Phase 2: SDD Context Integration
- [x] T3: Implement task extractor (finds active task in `tasks.md`). `id: REV-003`
- [ ] T4: Implement file-to-task mapper (checks if touched files are expected). `id: REV-004`

## Phase 3: Semantic Auditor
- [x] T5: Implement basic heuristics (print detector, large hunk warnings). `id: REV-005`
- [x] T6: Create the audit engine that combines context and diff. `id: REV-006`

## Phase 4: CLI & UI
- [x] T7: Create the `hb review` command. `id: REV-007`
- [ ] T8: Integrate `internal/ui` for premium audit reports. `id: REV-008`
