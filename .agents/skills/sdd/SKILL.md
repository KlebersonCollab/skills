---
name: sdd
version: 2.2.0
description: Spec-Driven Development. High-performance modular workflow with Explore-Plan-Act Loop and Safety-Valve protocols. Orchestrates specialized sub-skills. [TLC Spec Driven](https://github.com/tech-leads-club/agent-skills/tree/main/packages/skills-catalog/skills/(development)/tlc-spec-driven)
category: development-workflow
---

## 🔒 Prerequisites (Mandatory)
This skill operates as the central brain for the **SDD** framework. Before any technical execution:
0. **Mode Check**: Verify `.hub-mode` and apply `token-distiller` guidelines.
1. **Context Check**: Rehydrate state by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`.
2. **Knowledge Check**: Follow the **Knowledge Verification Chain** (see below).

---

# SDD: Modular & Adaptive Workflow (v2.2.0)

> Precision at scale. Rigor when needed, speed when possible.

---

## 🧩 Delegation Matrix
The SDD workflow delegates heavy lifting to specialized sub-skills to maintain a clean context window:

| Phase | Sub-Skill | Primary Artifact | Purpose |
|---|---|---|---|
| **DISCOVERY** | `sdd-explorer` | `TECHNICAL-MAP.md` | Read-only mapping of stack, architecture, and risks. |
| **SPECIFY** | `sdd-orchestrator` | `spec.md`, `plan.md` | Requirement analysis (BDD) and technical design. |
| **IMPLEMENT** | `sdd-implementer` | Tested Code | Surgical implementation with TDD and atomic commits. |
| **VERIFY** | `sdd-reviewer` | `validation-report.md` | Evidence-based audit against Acceptance Criteria. |

---

## 🔄 4-Phase Workflow

The SDD follows a rigorous cycle to ensure integrity and traceability:

### 1. DISCOVERY
*   **Goal**: Understand the terrain and align with the project vision.
*   **Action**: Use `sdd-planner` to read `PROJECT.md` and `ROADMAP.md`. Use `sdd-explorer` to map stack and architecture.
*   **Output**: Updated `TECHNICAL-MAP.md` and "rehydrated" context via `STATE.md`.
*   **Handoff**: Follow [Handoff Protocol](references/handoff-protocol.md) Section 1.
*   **Trigger**: Session start or new complex feature.

### 2. SPECIFY
*   **Goal**: Define "what", "how" and record technical decisions.
*   **Action**: Use `sdd-orchestrator` for Specs/Plan and `sdd-planner` to update Roadmap and record decisions.
*   **Outputs**: `spec.md`, `plan.md`, `tasks.md`, `contract.md` + updates to `DECISIONS.md` and `ROADMAP.md`.
*   **Handoff**: Follow [Handoff Protocol](references/handoff-protocol.md) Section 2.
*   **Trigger**: Before any implementation (**Small+**).

### 3. IMPLEMENT
*   **Goal**: Technical execution with continuous progress tracking.
*   **Action**: Use `sdd-implementer` for code/tests. Update `STATE.md` continuously with task progress.
*   **Output**: Verified code and updated `tasks.md`.
*   **Handoff**: Follow [Handoff Protocol](references/handoff-protocol.md) Section 3.
*   **Trigger**: After Phase 2 is approved.

### 4. VERIFY
*   **Goal**: Validate delivery and capture learnings.
*   **Action**: Use `sdd-reviewer` for audit and `sdd-planner` to capture discovered patterns and update memory.
*   **Output**: Updated `validation-report.md`, `LEARNINGS.md`, and `MEMORY.md`. State finalization in `STATE.md`.
*   **Handoff**: Follow [Handoff Protocol](references/handoff-protocol.md) Section 4.
*   **Trigger**: Technical completion of Phase 3 tasks.

---

## 🛠️ Operational Protocols

### 1. The Safety Valve (Automatic Escalation)
Every implementation task must be continuously monitored for complexity drift.
- **Trigger**: If a task reveals unexpected structural changes, touches >3 files, or modifies components listed as "Fragile" in `TECHNICAL-MAP.md`.
- **Action**: **STOP** execution. Notify the user: *"Task complexity has escalated beyond initial sizing. Re-evaluating the plan."*
- **Outcome**: Re-invoke `sdd-orchestrator` to update `plan.md` or `tasks.md`.

### 2. Knowledge Verification Chain
To prevent hallucinations and pattern drift, follow this strict hierarchy:
1. **Existing Code**: Scan for established patterns, conventions, and similar logic.
2. **Internal Specs**: Consult `TECHNICAL-MAP.md`, `CONVENTIONS.md`, and `STATE.md`.
3. **Project Documentation**: Check READMEs and official internal docs.
4. **MCP/External Tools**: Query `context-graph` or perform targeted web searches.
5. **Flag Uncertainty**: If Steps 1-4 yield no proof, flag as "Uncertain" to the user. **Never assume.**

### 3. Context Budget Management
- **Base Context**: Always load `STATE.md` (current progress) and `MEMORY.md` (preferences).
- **On-Demand**: Only load sub-skill definitions and specific feature specs when active in that phase.
- **Limit**: Target <30k tokens for the base environment to leave maximum room for reasoning.

### 4. Real-Time Progress Monitoring
- **Mandatory View**: Use **`hb sdd status --ui`** at the beginning and end of each session.
- **Watch Mode**: During implementation (IMPLEMENT phase), use **`hb sdd status --ui --watch`** to ensure task atomicity and progress visibility.

### 5. Swarm & Multi-Agent Execution
When working with multiple parallel agents or personas:
- **Infrastructure**: Utilize Tmux or iTerm2 for session persistence and visual tracking.
- **Security**: Propagate `DANGEROUS_BASH_PATTERNS` to all swarm members.
- **Protocol**: Follow the [Swarm Execution Guide](references/swarm-execution.md).

---

## 🏗️ Production Patterns

### Explore-Plan-Act Loop
Strict permission-gated flow:
- **EXPLORE**: No code modifications allowed.
- **PLAN**: Design validation via **Plan-Validate-Execute (PVE)**.
- **ACT**: Atomic execution on feature branches.

### Visual-First Mandate
All technical plans (`plan.md`) MUST include **Mermaid** diagrams to visualize data flow and component relationships.

---

## 🚫 Prohibited

- NEVER skip the **DISCOVERY** phase for unknown codebases.
- NEVER mark a task as complete without an atomic Git commit.
- NEVER use placeholders in `spec.md` or `plan.md`.
- NEVER ignore "Fragile" warnings in `TECHNICAL-MAP.md`.

---

## 5. State & Memory Management

### Persistence Protocol (via `sdd-planner`)
Always update `STATE.md` at the end of every session or after major decisions.
1.  **Operational Memory (`STATE.md`)**: Tasks, session status, and blockers.
2.  **Persistent Knowledge (`MEMORY.md`)**: Enduring facts, style guides, and preferences.
3.  **Incremental Wisdom (`LEARNINGS.md`)**: Bug solutions and discovered patterns.

---

## 6. Directory Structure (Mandatory)

### Feature-Specific (`spec/` or `.specs/features/[name]/`)
- `spec.md`: Traceable requirements (FR-X) and Acceptance Criteria (ACs).
- `plan.md`: Technical design and schemas.
- `tasks.md`: Atomic task list with status.
- `contract.md`: Delivery agreement and validation sensors.

### Project-Wide (`.specs/project/`)
- `PROJECT.md`: Core vision and "North Star".
- `ROADMAP.md`: Milestones and feature status.
- `STATE.md`: Operational memory (current status, blockers).
- `MEMORY.md`: Persistent knowledge (patterns, preferences).
- `LEARNINGS.md`: Incremental wisdom (solved bugs, tricks).
- `DECISIONS.md`: Log of architectural and design decisions.

### Codebase Mapping (`.specs/codebase/`)
- `TECHNICAL-MAP.md`: Mapping of stack and dependencies.
- `CONVENTIONS.md`: Coding patterns and linting.
- `ARCHITECTURE.md`: High-level structural view.

---

> **Law of SDD**: If it's not in the spec, it doesn't exist. If it's not verified, it's not done.

---

## References
- [Brownfield Mapping Guide](references/brownfield-mapping.md)
- [Coding Principles](references/coding-principles.md)
- [Context Management](references/context-limits.md)
- [BDD Standard](references/bdd-guide.md)
- [Handoff Protocol](references/handoff-protocol.md)
- [Swarm Execution Guide](references/swarm-execution.md)
