---
name: sdd
version: 2.0.0
description: Spec-Driven Development. High-performance workflow with Explore-Plan-Act Loop and Plan-Validate-Execute mandates.
category: development-workflow
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Have you rehydrated the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)?
3. **Plan Check**: Does the `plan.md` file define architecture, schemas, and include **Mermaid** diagrams?

---
# SDD: Modular & Adaptive Workflow (v2.0.0)

> Precision at scale. Rigor when needed, speed when possible.

---

## 🏗️ Production Workflow Patterns

### 1. The Explore-Plan-Act Loop
The SDD workflow is now strictly organized around three permission-gated phases:
- **EXPLORE**: Read-only mapping of the terrain. No code changes allowed.
- **PLAN**: Designing the solution and validating it via **Plan-Validate-Execute**.
- **ACT**: Implementation and verification.

### 2. Plan-Validate-Execute (PVE)
For **Medium+** complexity, a verifiable plan (JSON/YAML or strict Markdown) must be validated by a script or the user before any code is modified.

### 3. UltraPlan (Remote Exploration)
Para tarefas de nível **Complex** com alta ambiguidade, o agente deve sugerir um "Teleporte" de exploração.
- **O que:** Desacoplar a exploração inicial para uma sessão remota (web/detached) para evitar a saturação de tokens do Hub local.
- **Gatilho:** Quando o mapeamento inicial (Explore) revelar dependências desconhecidas em mais de 3 sistemas externos ou o plano exceder 20 tasks.
- **Resultado:** O plano aprovado remotamente deve ser trazido de volta como uma especificação técnica consolidada (`spec.md`).

---

## The Core Principle: Auto-Sizing

| Level | Scope | Phases | Required Sub-skills |
|---|---|---|---|
| **Quick** | Bug fixes, config, <3 files | (Act) | `sdd-implementer` |
| **Small** | Clear feature, <5 tasks | (Plan) + (Act) | `sdd-orchestrator` |
| **Medium** | Feature + UI, <10 tasks | Explore + Plan (PVE) + Act | `sdd-explorer`, `sdd-orchestrator` |
| **Large** | Multi-component | Explore + Plan (PVE + RFC) + Act | All modules |
| **Complex** | Ambiguity, high risk | Explore + Plan + Act (Parallel) | All modules + `sdd-reviewer` |

---

## Workflow (4 Phases)

### Phase 1: EXPLORE — Mapping and Context (Read-Only)
1.  **Initialize Feature**: Run **`hb sdd start <feature>`**.
2.  **Terrain Mapping**: Use `sdd-explorer` to understand the stack and architecture.
3.  **Impact Analysis**: Run **`hb map --impact <file>`** to identify regressions.

### Phase 2: PLAN — Contracts and Validation (PVE)
1.  **Write the Spec**: Define requirements in `spec.md` (BDD format mandatory).
2.  **Design the Plan**: Draft `plan.md` with **Mermaid** diagrams.
3.  **Meta-Knowledge Check**: Query the **Context Graph** for precedents/patterns using the `context-graph` skill.
4.  **Log Decisions**: Record critical architectural choices in the project's `DECISIONS.md`.
5.  **Validate the Plan**: Apply **Plan-Validate-Execute**. For destructive tasks, verify the plan artifact before moving to ACT.
6.  **Contract (SDC)**: Define validation sensors in `contract.md`.

### Phase 3: ACT — Atomic Execution
1.  **Isolation**: Create a feature branch.
2.  **Implementation**: Use `sdd-implementer` to write TDD code.
3.  **Task Sync**: Use `hb sdd task <feature> <id>` to track progress.
4.  **Risk Classification**: Approve commands through the harness risk classifier.

### Phase 4: VERIFY — Audit and Handoff
1.  **Verdict**: Use `sdd-reviewer` to audit against `spec.md` and `contract.md`.
2.  **Evaluation**: Run **`hb harness evaluate`** (Score >= 7.0 required).
3.  **Consolidation**: Run `hb sync` to update memory and state.

---

## Quality Rules

- **Explore Before Act**: Never edit a file you haven't first read and mapped.
- **PVE Mandate**: High-stakes operations require a separate validation step before execution.
- **Diagram-as-Code**: All plans must include Mermaid visualizations.
- **Branch-First**: No commits directly to `main`.

## Prohibited

- NEVER mark a task as complete without individual git commits.
- NEVER skip the Explore phase for features in unfamiliar codebases.
- NEVER skip the Plan phase (PVE) for Medium+ tasks, even if the implementation seems trivial.
- NEVER use placeholders in `spec.md` or `plan.md`.
- NEVER ignore the "Visual-First" mandate (Mermaid diagrams are mandatory in all plans).

---

## References

- [12 Agentic Harness Patterns (Explore-Plan-Act)](https://generativeprogrammer.com/p/12-agentic-harness-patterns-from)
- [Skill Authoring Patterns (Plan-Validate-Execute)](https://generativeprogrammer.com/p/skill-authoring-patterns-from-anthropics)
- [BDD Guide](references/bdd-guide.md)
- [UltraPlan Protocol](resource:resources/ULTRAPLAN_PROTOCOL.md)
