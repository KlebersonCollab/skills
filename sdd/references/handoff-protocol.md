# SDD Handoff Protocol

This protocol defines the "Baton Pass" between different phases and sub-skills to ensure zero-loss context transmission.

---

## 1. Discovery ➡️ Specify (The Context Pass)
**Performed by**: `sdd-explorer` ➡️ `sdd-orchestrator`

- **Pre-requisite**: `TECHNICAL-MAP.md` is updated with the current stack and architecture.
- **Handoff Artifacts**:
    - `.specs/codebase/TECHNICAL-MAP.md`: Reality check for the Architect.
    - `STATE.md`: Updated to `PHASE: SPECIFY`.
- **Validation**: The Orchestrator must acknowledge the "Critical Risks" in `CONCERNS.md` before starting the spec.

## 2. Specify ➡️ Implement (The Contract Pass)
**Performed by**: `sdd-orchestrator` ➡️ `sdd-implementer`

- **Pre-requisite**: Spec and Plan are complete and include BDD and Mermaid diagrams.
- **Handoff Artifacts**:
    - `spec.md`: Functional requirements.
    - `plan.md`: Technical blueprint.
    - `tasks.md`: Atomic checklist.
    - `contract.md`: Validation sensors and score threshold.
- **Validation**: The Implementer must verify that each task has a clear "How to test" description.

## 3. Implement ➡️ Verify (The Evidence Pass)
**Performed by**: `sdd-implementer` ➡️ `sdd-reviewer`

- **Pre-requisite**: All tasks in `tasks.md` are marked as complete `[x]` and committed.
- **Handoff Artifacts**:
    - `tasks.md`: Evidence of task completion.
    - `Changed Files`: The actual code implementation.
    - `Local Test Results`: Preliminary proof of success.
- **Validation**: The Reviewer runs the **Sensors** defined in `contract.md` to confirm implementation reality.

## 4. Verify ➡️ Persistence (The Wisdom Pass)
**Performed by**: `sdd-reviewer` ➡️ `sdd-planner`

- **Pre-requisite**: Final score meets the `contract.md` threshold.
- **Handoff Artifacts**:
    - `validation-report.md`: The final verdict.
    - `STATE.md`: Finalized session status.
- **Action**: The Planner extracts `LEARNINGS.md` (what we discovered) and `MEMORY.md` (what we standardized) from the implementation cycle.

---

## 🔄 The Swarm Loop
In multi-agent environments, each handoff is a signal to "Switch Persona" or "Start New Thread" to ensure the context window remains focused on the current phase's specific mission.
