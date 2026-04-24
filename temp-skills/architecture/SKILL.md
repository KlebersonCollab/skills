---
name: architecture
version: 2.0.1
description: "Systems Architect — guides the agent to design scalable, resilient, and distributed systems (CQRS, Event-Driven) using mandatory ADRs, Fitness Functions, and Mermaid Diagrams for visualization."
category: architecture
uses:
  - sdd
  - token-distiller
references:
  - clean-code-mentor
  - api-architect
---

# Architecture (v2.0.1)

> High-performance systems designer and guardian of evolutionary simplicity. "Architecture is what remains when you take all the code away."

---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the guidelines of the `token-distiller` skill.
1. **Context Check**: Have you rehydrated the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define the architecture, schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---
## Goal

Empower the agent to design high-quality and complex software architectures, spanning distributed systems (CQRS, Event-Driven) and modular monoliths. The skill ensures that every decision is justified by trade-offs, prioritizes simplicity, and is protected by **Fitness Functions** and visually documented via **Mermaid**.

---

## When to Use This Skill

- When designing systems that require high read/write scalability (CQRS).
- When designing asynchronous and decoupled workflows (Event-Driven).
- During major refactorings or architecture migrations (e.g., Monolith to Microservices).
- To define automated governance rules via Fitness Functions.
- To document critical technical decisions that impact the long term (ADRs).

---

## Workflow (4 Fases)

### Phase 1: CONTEXT — Requirements Discovery
1.  **Map Constraints**: Identify time limits, budget, and current stack.
2.  **Identify Load Patterns**: Differentiate read vs. write volumes (indicative of CQRS).
3.  **Need for Decoupling**: Evaluate if synchronous communication is a bottleneck (indicative of Event-Driven).

### Phase 2: ANALYSIS — Trade-off Evaluation
1.  **Explore Alternatives**: Compare approaches (e.g., Simple vs. Scalable).
2.  **Consistency Analysis**: Evaluate if the business tolerates Eventual Consistency or requires Strong Consistency.
3.  **Simplicity Analysis**: Justify the introduction of complex patterns (CQRS/Events) only if the scale requirement demands it.

### Phase 3: DESIGN — Pattern Selection and Contracts
1.  **Command/Query Modeling**: If using CQRS, clearly define Read and Write models.
2.  **Event Design**: Define message schemas, idempotency strategies, and DLQs (Dead Letter Queues).
3.  **Design Components**: Apply SOLID, DRY, and YAGNI to the proposed design.

### Phase 4: DOCUMENT — Recording and Governance (ADR & Fitness)
1.  **Write the ADR**: Document the technical choice using the official template.
2.  **Define Fitness Functions**: Propose scripts or automated tests to validate that the proposed architecture will not suffer regressions (e.g., avoiding circular coupling).
3.  **Handoff**: Deliver the architectural specification ready for implementation.

---

## Output Structure

The execution of this skill results in the following mandatory artifacts in `.specs/architecture/`:

| Artifact | Format | Description |
|----------|---------|-----------|
| **ADR-NNN** | `.md` | Architecture Decision Record with justification and impact. |
| **System Map** | **Mermaid** | Renderable diagram of components and flows. **Mandatory**. |
| **Fitness Specs** | `.py` / `.sh` | Definition of automated tests for architecture governance. |
| **Trade-off Matrix**| Table | Comparison between the analyzed alternatives. |

---

## Quality Rules

- **Simplicity First**: Do not use CQRS or Events if a simple relational database solves the problem.
- **Mandatory Idempotency**: Every event-oriented design must provide for repeated processing without side effects.
- **Fitness-Driven**: Every important architectural constraint must have a way to be automatically validated.
- **Visual-First**: Every complex component or flow must be documented with Mermaid in the `System Map`.
- **Justified Decisions**: Structural changes require a documented "Why."

## Prohibited

- NEVER propose distributed systems without an operational cost and latency analysis.
- NEVER use events for communication that requires an immediate (synchronous) response.
- NEVER ignore the complexity of managing eventual consistency on the front-end.
- NEVER start design without understanding throughput and availability requirements.

---

## References

- [`references/architectural-principles.md`](references/architectural-principles.md) — SOLID, KISS, YAGNI.
- [`references/cqrs-and-events.md`](references/cqrs-and-events.md) — (New) Command and event design.
- [`references/evolutionary-architecture.md`](references/evolutionary-architecture.md) — (New) Fitness Functions and evolution.
- [`references/adr-template.md`](references/adr-template.md) — Official ADR template.
