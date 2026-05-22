# SDD Handoff Protocol

This protocol defines the "Baton Pass" mechanisms to ensure zero-loss context transmission, applicable both during development phase transitions and when pausing/resuming work sessions.

---

## Part 1: Phase Handoffs (The Baton Pass)

This section defines transitions between different development lifecycle phases.

### 0. State Transition Validation (Gated Check)
Before initiating any phase handoff, the agent MUST perform a "Metadata Audit":
- [ ] `STATE.md` phase reflects the *source* phase of the handoff.
- [ ] All source artifacts (`spec.md`, `plan.md`, etc.) are marked `status: COMPLETED`.
- [ ] `tasks.md` has evidence for all tasks in the source phase.
- [ ] **Prohibited**: Never jump to `VERIFY` if `IMPLEMENT` tasks are still in `IN_PROGRESS`.

### 1. Discovery ➡️ Specify (The Context Pass)
**Performed by**: `sdd-explorer` ➡️ `sdd-orchestrator`
- **Pre-requisite**: `TECHNICAL-MAP.md` is updated with the current stack and architecture.
- **Handoff Artifacts**:
    - `.specs/codebase/TECHNICAL-MAP.md`: Reality check for the Architect.
    - `STATE.md`: Updated to `PHASE: SPECIFY`.
- **Validation**:
    - The Orchestrator must acknowledge the "Critical Risks" in `CONCERNS.md`.
    - **Metadata Audit**: Verify `STATE.md` contains the `<!-- @sdd-state -->` block with the correct phase.

### 2. Specify ➡️ Implement (The Contract Pass)
**Performed by**: `sdd-orchestrator` ➡️ `sdd-implementer`
- **Pre-requisite**: Spec and Plan are complete and include BDD and Mermaid diagrams.
- **Handoff Artifacts**:
    - `spec.md`: Functional requirements.
    - `plan.md`: Technical blueprint.
    - `tasks.md`: Atomic checklist.
    - `contract.md`: Validation sensors and score threshold.
- **Validation**:
    - The Implementer must verify that each task has a clear "How to test" description.
    - **Metadata Audit**: Verify `spec.md` and `plan.md` contain the `<!-- @sdd-state -->` block with `status: COMPLETED`.

### 3. Implement ➡️ Verify (The Evidence Pass)
**Performed by**: `sdd-implementer` ➡️ `sdd-reviewer`
- **Pre-requisite**: All tasks in `tasks.md` are marked as complete `[x]` and committed.
- **Handoff Artifacts**:
    - `tasks.md`: Evidence of task completion.
    - `Changed Files`: The actual code implementation.
    - `Local Test Results`: Preliminary proof of success.
- **Validation**:
    - The Reviewer runs the **Sensors** defined in `contract.md`.
    - **Metadata Audit**: Verify `tasks.md` uses the structured table format and includes the `Evidence` column for all `[x]` tasks.

### 4. Verify ➡️ Persistence (The Wisdom Pass)
**Performed by**: `sdd-reviewer` ➡️ `sdd-planner`
- **Pre-requisite**: Final score meets the `contract.md` threshold.
- **Handoff Artifacts**:
    - `validation-report.md`: The final verdict.
    - `STATE.md`: Finalized session status.
- **Action**: The Planner extracts `LEARNINGS.md` (what we discovered) and `MEMORY.md` (what we standardized) from the implementation cycle.

---

## Part 2: Session Handoffs (Pause & Resume)

This section defines transitions between AI work sessions to prevent short-term memory loss when a session ends or pauses.

### 1. Pausing Work
Before ending a session's context or pausing execution:
1. **Update STATE.md**: Document the current active task, the last blockers faced, and immediate next steps.
2. **Atomic Commit**: Ensure ongoing work is committed to the feature branch (use `[WIP]` commits if incomplete).
3. **Task Check**: Mark exactly where execution stopped in the local `tasks.md`.

### 2. Resuming Work
When starting or resuming a session:
1. **Rehydrate Context**: First, read `.specs/project/STATE.md`, `MEMORY.md`, and `LEARNINGS.md`.
2. **Identify Active Task**: Read `STATE.md` to see exactly what needs to be done next.
3. **Git Sync**: Validate the current state of the feature branch in Git.
4. **Wisdom Check**: Review recently added technical insights or standards in `LEARNINGS.md`.

### 3. Continuity Rules
The persistent operational memory in `STATE.md` represents the persistent brain of the project. Every session transition must treat `STATE.md` as the single source of truth for execution state.

---

## 🔄 The Swarm Loop
In multi-agent environments, each handoff is a signal to "Switch Persona" or "Start New Thread" to ensure the context window remains focused on the current phase's specific mission.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "SDD-DE-AMBIGUATION"
phase: "IMPLEMENT"
status: "COMPLETED"
last_update: "2026-05-22T14:32:00Z"
evidence_checksum: "8e52f6a"
```
