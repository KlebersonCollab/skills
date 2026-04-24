---
name: harness-expert
version: 2.1.0
description: "Technical engine for Harness Engineering. Implements adversarial feedback loops (GAN-style) to ensure production quality, in addition to state management and determinism."
category: agentic-infrastructure
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Run **`hb harness rehydrate [feature]`** to instantly load state, memory, and learnings.
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define the architecture, schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---
# Harness Expert (v2.1.0)

> "If you are not the model, then you are the harness." — Support infrastructure to transform LLMs into high-performance operational agents.

---

## Goal

The goal of this skill is to provide the "machinery" (CLI, scripts, automation) for the **SDD Planner** to execute its directives. While the Planner defines the *what* and *why*, the Harness provides the technical *how* to maintain deterministic state and close adversarial feedback loops (GAN-style) that ensure the final delivery is **production-ready**.

---

## 🔄 GAN-style Feedback Loop (Quality Enforcement)

For high-complexity tasks (**Large/Complex**), the Harness must operate in **Adversarial** mode:

| Role | Responsibility | Tools |
|-------|------------------|-------------|
| **Generator** | Implement the feature according to the Spec and Plan. | HB CLI, SDD CLI, Frameworks. |
| **Evaluator** | Act as a relentless QA, testing and scoring. | **`hb harness evaluate`**, Playwright. |

### 📊 Evaluation Rubric (Target: Score >= 7.0)

> [!IMPORTANT]
> The score is calculated and recorded via **`hb harness evaluate`**. If the score is **< 7.0**, the `hb sync` command will be **BLOCKED**.

| Criterion | Weight | Description |
|----------|------|-----------|
| **Design Quality** | 0.3 | Visual cohesion, absence of generic AI patterns, premium aesthetics. |
| **Originality** | 0.2 | Creative technical/visual decisions customized for the problem. |
| **Craft** | 0.3 | Polish, micro-interactions, typography, spacing, and performance. |
| **Functionality** | 0.2 | Robust functioning of all features and error handling. |

---

## Workflow (4 Phases)

### Phase 1: AUTOMATED REHYDRATE — Technical Loading
1.  **Context Injection**: Run **`hb harness rehydrate [feature]`** to instantly align agent memory with the project state.
2.  **Impact Mapping**: Run **`hb map --impact <file>`** to visualize technical dependencies.
3.  **Health Check**: Run `hb doctor` to ensure the environment is ready for operation.

### Phase 2: OPERATE — Execution Automation (Generator Mode)
1.  **Tool Orchestration**: Invoke MCPs and local scripts according to the plan.
2.  **Implementation Loop**: Write code focused on passing the ACs in `spec.md`.
3.  **Self-Correction**: Immediately fix linter/test bugs via `hb audit` and `sdd-implementer`.

### Phase 3: ADVERSARIAL REVIEW — Feedback Cycle (Evaluator Mode)
1.  **Strict Scoring**: Execute **`hb harness evaluate`** providing grades based on the rubric.
2.  **Deep Audit**: Execute **`hb doctor --deep`** to ensure ADR/Spec compliance.
3.  **Feedback Iteration**: Review the generated `validation-report.md`. If rejected, iterate.

### Phase 4: FINAL SYNC — Determinism and Closing
1.  **Exit Gates**: Ensure the score is >= 7.0 (The CLI will enforce this).
2.  **Capture Insights**: Run **`hb learn`** to record new patterns or bug fixes in `LEARNINGS.md`.
3.  **Auto-Sync**: Execute **`hb sync`** to update the ecosystem and agents.
4.  **Handoff**: Prepare distribution artifacts in `dist/`.

---

## Output Structure

Execution of this skill results in the following mandatory artifacts:

| Artifact | Format | Description |
|----------|---------|-----------|
| **Tasks Sync** | `tasks.md` | Task list updated with real progress. |
| **State Snapshot** | `STATE.md` | Summary of the current system and feature state. |
| **Learnings Log** | `LEARNINGS.md` | Record of solved problems and technical lessons. |
| **Validation Report** | `validation-report.md` | Test and compliance report for the generated code. |

---

## Quality Rules

- **Persistence First**: The File System is the only source of truth. No decisions should be made without consulting local specs.
- **SDD Compliance**: All development must strictly follow Spec-Driven Development.
- **Context Awareness**: Constantly monitor token usage and perform "Context Compression" when necessary.
- **Deterministic Check**: Every agent output must pass through verification tools (linters/tests) before being considered final.

## Prohibited

- NEVER start a task without reading existing specs.
- NEVER consider a task "completed" without updating synchronization files (`tasks.md`).
- NEVER ignore linter or test errors during the validation phase.
- NEVER maintain redundant context that can be summarized in `STATE.md`.
