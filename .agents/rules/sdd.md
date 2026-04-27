---
name: sdd
version: 1.5.0
description: Spec-Driven Development. Modular workflow with PRD/RFC, BDD, and Mermaid Diagrams mandate.
category: development-workflow
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Have you rehydrated the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define architecture, schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---

# SDD: Modular & Adaptive Workflow

> Precision at scale. Rigor when needed, speed when possible.

---

## Goal

The objective of SDD is to ensure precision, traceability, and integrity throughout the software development lifecycle, transforming ambiguous requirements into verified and documented code through a modular and adaptive workflow, now enhanced by **Harness Engineering** principles.

---

## The Core Principle: Auto-Sizing

The depth of the workflow is determined by the **Complexity** of the task, not a fixed pipeline.

| Level | Scope | Phases | Required Sub-skills | Directory |
|---|---|---|---|---|
| **Quick** | Bug fixes, config, <3 files | (Implement) + (Verify) | `sdd-implementer` | Root |
| **Small** | Clear feature, <5 tasks | (Spec) + Impl + Verify | `sdd-orchestrator`, `sdd-implementer` | `spec/` |
| **Medium** | Feature + UI, <10 tasks | Explorer + Spec (BDD) + Plan + **Contract** + Impl + Verify | `sdd-explorer`, `sdd-orchestrator` | `spec/` |
| **Large** | Multi-component, new module | Planner + Explorer + RFC + Spec + Plan + **Contract** + Impl + Verify | All modules | `.specs/` |
| **Complex** | Ambiguity, high risk | Same as Large + PRD Audit + Score Review | All modules + `sdd-reviewer` | `.specs/` |

---

## Workflow (4 Phases)

The SDD lifecycle follows an iterative and rigorous delivery flow:

### Phase 1: DISCOVERY — Mapping and Context
1.  **Explore the Terrain**: Use `sdd-explorer` to understand the current stack and architecture.
2.  **Define Vision**: Use `sdd-planner` to align objectives in `PROJECT.md` and `ROADMAP.md`.
3.  **Memory Capture**: Rehydrate context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`.
4.  **Product Requirements**: For **Complex** tasks, audit the existing PRD (Product Requirements Document).

### Phase 2: SPECIFY — Contracts and Planning
1.  **Write the Spec**: Use `sdd-orchestrator` to define technical requirements (`spec.md`). For **Medium+** levels, use **BDD** format (Given/When/Then) in Acceptance Criteria.
2.  **Technical Proposal (RFC)**: For **Large+** levels, draft an **RFC** (Request for Comments) detailing architecture and trade-offs before `plan.md`.
3.  **Design the Solution**: Draft `plan.md` with architecture and schemas. **Mandatory**: Include **Mermaid** diagrams (Flowcharts, Sequence, Class) to visualize flows and structures in **Medium+** levels.
4.  **Development Contract (SDC)**: Create `contract.md` defining the agreement between Implementer and Reviewer on what will be delivered and how it will be measured (Sensors).
5.  **Atomic Tasks**: Generate the task list (`tasks.md`) for execution.

### Phase 3: IMPLEMENT — Atomic Execution
1.  **Isolation**: **MANDATORY** create a feature branch (`feat/`, `fix/`, `docs/`) from `main` before starting code. NEVER commit directly to `main`.
2.  **Automation**: Use **SDD CLI** (`uv run sdd task <feature> <id>`) to track progress.
3.  **Code Cycle**: Use `sdd-implementer` to write test-driven code (TDD) aligned with BDD scenarios.
4.  **Prohibitive Mandate**: **NEVER** mark a task as complete in `tasks.md` without passing 100% of tests and performing the corresponding git commit on the feature branch.
5.  **Integrity**: Ensure each task in `tasks.md` is marked as complete only after passing tests and being committed individually.

### Phase 4: REVIEW — Audit and Finalization
1.  **Verdict via Sensors**: Use `sdd-reviewer` to audit delivery against `spec.md` and `contract.md`, using tool feedback (tests, linters).
2.  **Score Report**: Generate a validation report with a score from 0-100.
3.  **UAT**: Validate with the user if Acceptance Criteria (BDD) were met.
4.  **Persistence**: Update memory logs and state in the Planner before closing.

---

## The Modular Engine: When, How & Why

This skill delegates tasks to specialized sub-skills to ensure scalability and context efficiency:

| Sub-skill | **When** to use? (Trigger) | **How** to delegate? | **Why** is it vital? |
|---|---|---|---|
| **[Explorer](sdd-explorer.skill.md)** | When starting in legacy or unknown code. | Request a stack and architecture mapping (`STACK.md`, `ARCHITECTURE.md`). | Prevents hallucinations about existing dependencies and patterns. |
| **[Planner](sdd-planner.skill.md)** | At the end of sessions, after decisions or roadmap changes. | Instruct to update `STATE.md`, `MEMORY.md`, and capture `LEARNINGS.md`. | Ensures state persistence and continuity between sessions (Harness). |
| **[Orchestrator](sdd-orchestrator.skill.md)** | Whenever a feature (**Small+**) is started. | Delegate the creation of `spec.md`, `plan.md`, and the breakdown into `tasks.md`. | Maintains separation between "what to do" and "how to do it," ensuring traceability. |
| **[Implementer](sdd-implementer.skill.md)** | When the plan and atomic tasks are ready. | Delegate execution of specific tasks from `tasks.md` via TDD and commits. | Focuses the agent on writing clean code and tests without the cognitive load of design. |
| **[Reviewer](sdd-reviewer.skill.md)** | When Phase 3 implementation ends or at critical milestones. | Request an evidence-based audit against ACs (BDD) and `contract.md`. | Ensures delivery meets quality standards and original criteria. |

> [!TIP]
> In **High Complexity** scenarios, the Orchestrator should act as the "Central Brain," maintaining only the plan and delegating research (Explorer) and coding (Implementer) to keep its own context window clean and focused.

---

## The Knowledge Verification Chain

When researching or deciding, follow this strict hierarchy:

1. **Codebase Docs**: Consult `.specs/codebase/` (generated by Explorer).
2. **Project Docs**: Check READMEs, `PROJECT.md`, `ROADMAP.md`.
3. **Existing Code**: Read the source directly for patterns.
4. **Web Search**: Official documentation only.
5. **Flag Uncertainty**: If Steps 1-4 fail, communicate uncertainty to the user.

---

## Operational Protocols

### The Safety Valve
If a task identified as **Quick** or **Small** reveals hidden complexity (>5 steps or deep dependencies), **STOP**. Formalize the task as **Large** by creating a full spec and plan.

### State Management
Utilize the **Planner** (`sdd-planner`) to manage three distinct types of memory:
1.  **Operational Memory (`STATE.md`)**: Tasks, session status, and blockers.
2.  **Persistent Knowledge (`MEMORY.md`)**: Enduring facts, style guides, and user preferences.
3.  **Incremental Wisdom (`LEARNINGS.md`)**: Solutions to bugs and code patterns discovered.

This triad ensures that context persists even if the agent is restarted or the context window shifts.

### Sub-Agent Delegation
To maximize performance and avoid exceeding the context window, delegate heavy tasks to sub-agents (when supported):
- **Research/Brownfield:** Delegate and require only a summary.
- **Implementation and Tests:** A sub-agent focuses on code, edits files, and returns status.
- Never delegate *Planning* or *Task Creation*; the Orchestrator must always maintain consolidated context.

### Verification Standards
A feature is NOT complete until the **Reviewer** issues an `APPROVED` verdict based on evidence (file paths and line numbers).

---

## References & Policies

To maintain rigorous compliance and handle specific situations, consult the detailed guidelines (aggregated from the TLC Spec-Driven standard):

- **[Brownfield Mapping](references/brownfield-mapping.md)**: How to audit and map existing repositories (STACK, ARCHITECTURE, CONCERNS).
- **[Coding Principles](references/coding-principles.md)**: Simplicity, integrity in testing, and objective focus.
- **[Context Limits](references/context-limits.md)**: Maximum artifact sizes and token warning zones.
- **[Session Handoff](references/session-handoff.md)**: Mandatory protocol for pausing and resuming work without losing context.
- **[Quick Mode](references/quick-mode.md)**: Express route for trivial tasks (<3 files) not requiring the overhead of all phases.
- **[BDD Guide](references/bdd-guide.md)**: Standard for writing acceptance specifications.

---

## Directory Structure (Mandatory)

### Feature-Specific (`spec/` or `.specs/features/[name]/`)
- `spec.md`: Traceable requirements (FR-X) and ACs.
- `plan.md`: Technical design, schemas, and Mermaid diagrams.
- `tasks.md`: Atomic task list with status.

### Project-Wide (`.specs/`)
- `project/`: `PROJECT.md`, `ROADMAP.md`, `STATE.md`, `MEMORY.md`, `LEARNINGS.md`.
- `codebase/`: `STACK.md`, `ARCHITECTURE.md`, `CONVENTIONS.md`, `CONCERNS.md`.

## Output Structure

Execution of this workflow must result in the following mandatory artifacts, organized in `.specs/features/[name]/`:

| Artifact | File | Description |
|----------|---------|-----------|
| **Specification** | `spec.md` | Functional requirements, ACs (BDD), and constraints. |
| **Technical Plan** | `plan.md` | Architecture, schemas, and Mermaid diagrams. |
| **Contract** | `contract.md` | Delivery agreement and validation sensors. |
| **Atomic Tasks** | `tasks.md` | Detailed task list with status. |
| **Audit Reports** | `validation-report.md` | Compliance reports with score and evidence. |

## Quality Rules

- **Gold Standard Benchmark**: Always consult `examples/gold-standard/` to align delivery quality with the expected level of rigor.
- **Adaptive Rigor**: The level of detail must strictly follow the Auto-Sizing table.
- **BDD-First**: Acceptance criteria for Medium+ levels must use Given/When/Then.
- **Continuous Verification**: A task is only considered complete after passing 100% of tests and being committed individually following `git-workflow`.
- **Mandatory Atomicity**: Each `[x]` entry in `tasks.md` must correspond to at least one Git commit with a message linked to the task ID.
- **Branch-First**: Development must always occur in short-lived branches, protecting the main branch (`main`).
- **Diagram-as-Code**: Technical drawings must use Mermaid integrated into Markdown.

## Prohibited

- **NEVER** start implementation (Phase 3) without an approved `spec.md`.
- **NEVER** ignore proactive updates to `STATE.md` and `MEMORY.md`.
- **NEVER** skip Phase 4 (Review) and spec closure when finishing a feature.
- **NEVER** use placeholders like "todo" or "..." in finalized specification files.
