---
name: harness-expert
version: 2.1.0
description: "Technical engine for Harness Engineering. Implements adversarial feedback loops (GAN-style) to ensure production quality, in addition to state management and determinism."
category: agentic-infrastructure
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Did you rehydrate the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
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
| **Generator** | Implement the feature according to the Spec and Plan. | SDD CLI, UV, Frameworks (React, Flutter). |
| **Evaluator** | Act as a relentless QA, testing the live application and rejecting "AI slop". | Playwright, Screenshotting, Log Analysis. |

### 📊 Evaluation Rubric (Target: Score >= 7.0)

| Criterion | Weight | Description |
|----------|------|-----------|
| **Design Quality** | 0.3 | Visual cohesion, absence of generic AI patterns, premium aesthetics. |
| **Originality** | 0.2 | Creative technical/visual decisions customized for the problem. |
| **Craft** | 0.3 | Polish, micro-interactions, typography, spacing, and performance. |
| **Functionality** | 0.2 | Robust functioning of all features and error handling. |

---

## Workflow (4 Phases)

### Phase 1: AUTOMATED REHYDRATE — Technical Loading
1.  **Context Injection**: Execute mass reading scripts of `.specs/` to rehydrate agent memory.
2.  **Health Check**: Validate if state files (`STATE.md`, `tasks.md`) are intact.
3.  **Context Compression**: Use `harness-expert-compress` to maintain critical tokens.

### Phase 2: OPERATE — Execution Automation (Generator Mode)
1.  **Tool Orchestration**: Invoke MCPs and local scripts according to the plan.
2.  **Implementation Loop**: Write code focused on passing the ACs in `spec.md`.
3.  **Self-Correction**: Immediately fix linter/test bugs via `hb audit` and `sdd-implementer`.

### Phase 3: ADVERSARIAL REVIEW — Feedback Cycle (Evaluator Mode)
1.  **Live Testing**: Start dev server and use Playwright to test real flows.
2.  **Strict Scoring**: Assign grades from 1 to 10 based on the rubric. **NEVER praise mediocre work**.
3.  **Feedback Iteration**: Generate `validation-report.md` with actionable criticism for the Generator.

### Phase 4: FINAL SYNC — Determinism and Closing
1.  **Exit Gates**: Ensure the GAN score reached the approval threshold (Default: 7.0).
2.  **Auto-Sync**: Update `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`.
3.  **Handoff**: Prepare distribution artifacts in `dist/`.

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
