# SDD Operational Guidelines

This document establishes the practical execution parameters for AI agents utilizing the SDD framework, covering fast-track workflows (Quick Mode) and strict context/token allocation boundaries to preserve high-fidelity reasoning.

---

## Part 1: Quick Mode (Fast-Track Workflow)

**Triggers**: "Quick fix", "Quick adjustment", "Trivial bugfix"

For minor maintenance, isolated configuration updates, or simple bugs touching no more than 3 files that can be fully explained in a single sentence.

### 1. Fast-Track Execution Flow
1. **Skip Discovery, Specify, and Tasks phases**: Jump directly to implementation and verification. Skip creating feature specs, architectural plans, or standalone task lists.
2. **Atomic TDD Execution**: Surgically apply the fix, ensuring that tests continue to pass and no visual or logical regressions occur.
3. **Audit Verification**: Verify build success, lint compliance, and test suite execution.
4. **Direct Commits**: Commit directly to the branch (or use a short-lived branch) without the ceremony of a fully specified PR loop.

### 2. The Safety Valve
If, during Quick Mode execution, you discover that the change:
- Touches >3 files unexpectedly.
- Requires adding new libraries or structural architectural modifications.
- Affects components labeled as "Critical Risks" in `TECHNICAL-MAP.md`.

**Action**: **STOP** execution immediately. Inform the user:
> *"This task seemed quick, but it involves deep architectural changes. Aborting Quick Mode and elevating to Small/Medium scope."*
Then, initialize conventional SDD phases (generate spec, plan, tasks, etc.).

---

## Part 2: Context Limits & Token Allocation

Context management is vital to maintaining logical reasoning capacity in long-running sessions. Follow these strict allocation guidelines to avoid cognitive dilution.

### 1. File Allocation Limits

| File | Max Tokens | ~Words | Warning Threshold |
| :--- | :--- | :--- | :--- |
| `PROJECT.md` | 2,000 | 1,200 | 1,600 (80%) |
| `ROADMAP.md` | 3,000 | 1,800 | 2,400 (80%) |
| `STATE.md` | 10,000 | 6,000 | 7,000 (70%) |
| `spec.md` | 5,000 | 3,000 | 4,000 (80%) |
| `plan.md` | 8,000 | 4,800 | 6,400 (80%) |
| `tasks.md` | 10,000 | 6,000 | 8,000 (80%) |
| `TECHNICAL-MAP.md` | 15,000 | 9,000 | 12,000 (80%) |
| `MEMORY.md` | 5,000 | 3,000 | 4,000 (80%) |

### 2. Context Health Zones
- 🟢 **Healthy** (<40k total tokens): Optimal operations. Silent.
- 🟡 **Moderate** (40-60k total tokens): Discreet footer warning suggesting compaction.
- 🔴 **Critical** (>60k total tokens): Active warning. Suggests immediate context optimization (e.g., using `token-distiller` or archiving chat history).

### 3. Operational Principles
1. **Target**: Maintain less than 40k active tokens loaded in the current prompt payload.
2. **On-Demand Hydration**: Never load multiple feature specifications or codebase maps simultaneously. Load them only when actively operating in their respective phases.
3. **Anchor Consistency**: Treat `STATE.md` and `TECHNICAL-MAP.md` as the ultimate high-fidelity anchors for session consistency.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "SDD-DE-AMBIGUATION"
phase: "IMPLEMENT"
status: "COMPLETED"
last_update: "2026-05-22T14:33:00Z"
evidence_checksum: "8e52f6a"
```
