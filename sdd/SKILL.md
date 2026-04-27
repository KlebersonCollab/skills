---
name: sdd
version: 2.1.0
description: Spec-Driven Development. High-performance modular workflow with Explore-Plan-Act Loop and Safety-Valve protocols. Orchestrates specialized sub-skills.
category: development-workflow
---

## 🔒 Prerequisites (Mandatory)
This skill operates as the central brain for the **SDD** framework. Before any technical execution:
0. **Mode Check**: Verify `.hub-mode` and apply `token-distiller` guidelines.
1. **Context Check**: Rehydrate state by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`.
2. **Knowledge Check**: Follow the **Knowledge Verification Chain** (see below).

---

# SDD: Modular & Adaptive Workflow (v2.1.0)

> Precision at scale. Rigor when needed, speed when possible.

---

## 🧩 Delegation Matrix
The SDD workflow delegates heavy lifting to specialized sub-skills to maintain a clean context window:

| Phase | Sub-Skill | Primary Artifact | Purpose |
|---|---|---|---|
| **EXPLORE** | `sdd-explorer` | `TECHNICAL-MAP.md` | Read-only mapping of stack, architecture, and risks. |
| **PLAN** | `sdd-orchestrator` | `spec.md`, `plan.md` | Requirement analysis (BDD) and technical design. |
| **ACT** | `sdd-implementer` | Tested Code | Surgical implementation with TDD and atomic commits. |
| **VERIFY** | `sdd-reviewer` | `validation-report.md` | Evidence-based audit against Acceptance Criteria. |

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
- NEVER skip the **EXPLORE** phase for unknown codebases.
- NEVER mark a task as complete without an atomic Git commit.
- NEVER use placeholders in `spec.md` or `plan.md`.
- NEVER ignore "Fragile" warnings in `TECHNICAL-MAP.md`.

---

## References
- [Brownfield Mapping Guide](references/brownfield-mapping.md)
- [Coding Principles](references/coding-principles.md)
- [Context Management](references/context-limits.md)
- [BDD Standard](references/bdd-guide.md)
