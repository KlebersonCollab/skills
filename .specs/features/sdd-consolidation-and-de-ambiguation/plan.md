# Plan - SDD Consolidation and De-ambiguation

## Proposed Changes

We will refactor the internal files of the `sdd` skill to clean up duplication.

### Proposed File Merges:
1. **Merge** `sdd/references/session-handoff.md` ➡️ `sdd/references/handoff-protocol.md`.
2. **Merge** `sdd/references/quick-mode.md` and `sdd/references/context-limits.md` ➡️ `sdd/references/operational-guidelines.md`.

### Deletions:
- `sdd/references/session-handoff.md` [DELETE]
- `sdd/references/quick-mode.md` [DELETE]
- `sdd/references/context-limits.md` [DELETE]

### New/Modified files:
- `sdd/references/handoff-protocol.md` [MODIFY] - Unified phase and session handoff guide.
- `sdd/references/operational-guidelines.md` [NEW] - Consolidated rules for quick tasks and context management.
- `sdd/SKILL.md` [MODIFY] - Update links and rules.
- `sdd/README.md` [MODIFY] - Update links.
- `sdd-implementer.skill.md` [MODIFY] - Update links.
- `sdd-orchestrator.skill.md` [MODIFY] - Update links.
- `sdd-planner.skill.md` [MODIFY] - Update links.
- `sdd-reviewer.skill.md` [MODIFY] - Update links.

### System Architecture Flow (Mermaid)

```mermaid
graph TD
    subgraph Original ["Original Duplicated Flow"]
        HP1[handoff-protocol.md - Phase Transitions]
        SH1[session-handoff.md - Session Pauses]
        QM1[quick-mode.md - Small tasks]
        CL1[context-limits.md - Budgets]
    end

    subgraph Consolidated ["New Clean Flow"]
        HP2["handoff-protocol.md (Unified)"]
        OG2["operational-guidelines.md (Unified)"]
    end

    HP1 ➡️ HP2
    SH1 ➡️ HP2
    QM1 ➡️ OG2
    CL1 ➡️ OG2
```

## Verification Plan

### Automated Tests
- Run `make audit` in the workspace to verify Markdown metadata compliance and link consistency.
- Run `make sync` to propagate any global changes if necessary.

### Manual Verification
- Verify that no remaining references to the deleted files exist in the `sdd/` directory.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "SDD-DE-AMBIGUATION"
phase: "SPECIFY"
status: "IN_PROGRESS"
last_update: "2026-05-22T14:26:00Z"
evidence_checksum: "NONE"
```
