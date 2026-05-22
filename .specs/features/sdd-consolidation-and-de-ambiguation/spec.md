# Spec - SDD Consolidation and De-ambiguation

## Goal
Optimize the SDD skill structure by eliminating redundant, overlapping, and ambiguous reference documentation, reducing token usage, and improving cognitive clarity for AI agents without sacrificing rigor.

## User Review Required
We are merging `session-handoff.md` into `handoff-protocol.md` and consolidating `quick-mode.md` and `context-limits.md` into a single file `references/operational-guidelines.md`. This will require updating links in several files.

## Acceptance Criteria
### AC-1: Handoff Protocol Unification
- **Given** that `references/session-handoff.md` and `references/handoff-protocol.md` exist.
- **When** the refactoring is executed.
- **Then** `session-handoff.md` is deleted, and its contents are elegantly merged into `handoff-protocol.md` under a unified handoff paradigm (Phase vs. Session handoffs).

### AC-2: Operational Guidelines Consolidation
- **Given** that `references/quick-mode.md` and `references/context-limits.md` exist.
- **When** the refactoring is executed.
- **Then** both files are deleted, and their rules are unified in a single, high-fidelity `references/operational-guidelines.md` file.

### AC-3: Internal Link and Reference Updates
- **Given** files in `sdd/` that reference deleted or moved documents.
- **When** the files are scanned and modified.
- **Then** all internal markdown links and rule descriptions point correctly to the new files (`references/handoff-protocol.md` and `references/operational-guidelines.md`).

### AC-4: Compliance and Purity Audit
- **Given** the modified `sdd` skill directory.
- **When** `make audit` and `make sync` are executed.
- **Then** the repository remains 100% compliant, with zero broken links or metadata violations, and is fully synchronized.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "SDD-DE-AMBIGUATION"
phase: "SPECIFY"
status: "IN_PROGRESS"
last_update: "2026-05-22T14:25:00Z"
evidence_checksum: "NONE"
```
