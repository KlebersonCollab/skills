---
name: brainstorming
version: 2.2.0
description: "Brainstorming and Design Facilitator — guides the agent to explore complex problems, generate divergent ideas, and converge on solid specifications before any implementation."
category: design-thinking
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify `.hub-mode` and apply `token-distiller` guidelines.
1. **Context Check**: Rehydrate state by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`.
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define the architecture and schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---
# Brainstorming

> Idea facilitator and problem solver. Think like an architect and facilitator, not an executor. The goal here is to "slow down to get it right."

---

## Goal

Enable the agent to facilitate high-quality brainstorming sessions, using divergence and convergence techniques to define problems, explore alternative solutions, evaluate trade-offs, and consolidate clear specifications before starting development.

---

## When to Use This Skill

- Before starting any complex project or feature.
- When the problem is not well-defined.
- To explore multiple technical approaches before choosing one.
- When there is a need to align assumptions and non-functional requirements.
- To resolve creative or technical bottlenecks.

## When NOT to Use This Skill

- For trivial tasks or obvious bugs.
- When the solution is already fully specified and validated.
- When you are already in the massive execution/implementation phase.

---

## Workflow (4 Phases)

### Phase 1: DISCOVER — Problem Definition
1.  **Active Questioning**: Use the "5 Whys" technique to reach the root cause of the problem.
2.  **Stakeholder Mapping**: Who are we building for? What are their pain points?
3.  **Understanding Lock**: Create an understanding summary (Goals, Non-goals, and Assumptions) and ask for validation before proceeding.

### Phase 2: DIVERGE — Idea Exploration
1.  **Multiplicity**: Generate at least 3 distinct approaches (e.g., Simple, Scalable, Experimental).
2.  **Creative Techniques**: Use SCAMPER or First Principles to challenge the status quo.
3.  **Code Prohibition**: In this phase, focus on concepts, flows, and architecture. **DO NOT WRITE PRODUCTION CODE**.

### Phase 3: CONVERGE — Evaluation and Choice
1.  **Trade-off Matrix**: List Pros and Cons of each approach (Complexity vs. Speed vs. Maintainability).
2.  **Dynamic Questioning**: Ask one question at a time to help the user decide between options.
3.  **Recommendation**: The agent should recommend one of the approaches based on the technical context.

### Phase 4: SPECIFY — Consolidation
1.  **Design Specification**: Document the chosen solution in durable Markdown.
2.  **Decision Log**: Record what was decided and why.
3.  **Handoff**: Define clear next steps for implementation.

---

## Output Structure

The execution of this skill should result in the following mandatory artifacts, preferably stored in `.specs/design/` or `docs/brainstorming/`:

| Artifact | Format | Description |
|----------|---------|-----------|
| **Design Specification** | `.md` | Consolidated document with the architecture and flows of the chosen solution. |
| **Decision Log** | `.md` | Historical record of the alternatives explored and the justification for the final choice. |
| **Trade-off Matrix** | Table | Visual comparison between the approaches (Simple, Scalable, Experimental). |

---

## Quality Rules

- **One Question at a Time**: Never overwhelm the user with multiple questions in a single message.
- **Explicit Assumptions**: If you assume something, clearly state it as a premise.
- **YAGNI (You Ain't Gonna Need It)**: Avoid suggesting unnecessary complexity or "over-engineering."
- **Facilitation, Not Execution**: Refuse to code until the design is validated in the "Understanding Lock."

## Prohibited

- NEVER start implementation without explicit design validation.
- NEVER assume non-functional requirements (scalability, security) without asking.
- NEVER ignore "Non-goals" (what we will not do).
- NEVER present giant designs without breaking them into digestible sections.

---

## References

- [`references/brainstorming-techniques.md`](references/brainstorming-techniques.md) — SCAMPER, 5 Whys, First Principles.
- [`references/decision-frameworks.md`](references/decision-frameworks.md) — Decision matrices and Trade-offs.


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.373740Z"
evidence_checksum: "NONE"
```
