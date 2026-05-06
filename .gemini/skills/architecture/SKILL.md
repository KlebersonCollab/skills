---
name: architecture
version: 2.2.1
description: "Systems Architect — guides the agent to design scalable, resilient, and distributed systems (CQRS, Event-Driven) using mandatory ADRs, Fitness Functions, and Mermaid Diagrams for visualization."
category: architecture
uses:
  - sdd
  - token-distiller
references:
  - clean-code-mentor
  - api-architect
---

# Architecture (v2.2.1)

> High-performance systems designer and guardian of evolutionary simplicity. "Architecture is what remains when you take all the code away."

---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the guidelines of the `token-distiller` skill.
1. **Context Check**: Rehydrate the context by reading `.specs/project/STATE.md`, `.specs/project/MEMORY.md`, and `.specs/project/LEARNINGS.md`.
2. **Knowledge Check**: Follow the **Knowledge Verification Chain** (Codebase Docs -> Project Docs -> Existing Code -> Web Search).
3. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
4. **Plan Check**: Does the `plan.md` file define the architecture, schemas, and include **Mermaid** diagrams?
5. **Contract Check**: Was the `contract.md` file established with validation sensors?
6. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---

## Goal

Empower the agent to design high-quality and complex software architectures, spanning distributed systems (CQRS, Event-Driven) and modular monoliths. The skill ensures that every decision is justified by trade-offs, prioritizes simplicity, and is protected by **Fitness Functions** and visually documented via **Mermaid**.

---

## Workflow (4 Phases)

### Phase 1: DISCOVERY — Requirements & Constraints
1.  **Map Terrain**: Identify technical, financial, and team constraints.
2.  **Identify Load Patterns**: Differentiate read vs. write volumes (indicative of CQRS).
3.  **Need for Decoupling**: Evaluate if synchronous communication is a bottleneck (indicative of Event-Driven).
4.  **Memory Capture**: Review previous ADRs in `.specs/architecture/` to ensure continuity.

### Phase 2: SPECIFY — Trade-offs & Design
1.  **Explore Alternatives**: Compare approaches (Simple vs. Scalable) in a Trade-off Matrix.
2.  **Consistency Analysis**: Evaluate if the business requires Strong Consistency or tolerates Eventual Consistency.
3.  **Visual Modeling**: Create **Mermaid** diagrams (System Map, Sequence) to validate the flow.
4.  **ADR Creation**: Document the technical choice by creating a new ADR file in `.specs/architecture/` following the official template.

### Phase 3: IMPLEMENT — Pattern Selection & Logic
1.  **Command/Query Modeling**: If using CQRS, clearly define Read and Write models.
2.  **Event Design**: Define message schemas, idempotency strategies, and DLQs.
3.  **Design Components**: Apply SOLID, DRY, and YAGNI. Ensure components are focused and decoupled.
4.  **Fitness Functions**: Define automated tests/scripts to protect architectural integrity.

### Phase 4: REVIEW — Documentation & Governance
1.  **Verdict via Sensors**: Audit the delivery against the original `spec.md` and `plan.md` using evidence from code and tests.
2.  **Finalize ADR**: Ensure the ADR reflects the final implementation and documented impacts.
3.  **Update Knowledge Map**: Update the project's visual architecture in `.specs/architecture/KNOWLEDGE-MAP.mermaid`.
4.  **Persistence**: Capture new architectural patterns in `.specs/project/LEARNINGS.md`.

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
- **Justified Decisions**: Structural changes require a documented "Why" (ADR).

## Prohibited

- NEVER propose distributed systems without an operational cost and latency analysis.
- NEVER use events for communication that requires an immediate (synchronous) response.
- NEVER ignore the complexity of managing eventual consistency on the front-end.
- NEVER start design without understanding throughput and availability requirements.
- NEVER create local memory files (`STATE.md`, etc.); use `.specs/project/`.
- NEVER rely on CLI tools for governance; the Markdown artifacts are the single source of truth.

---

## References

- [`references/architectural-principles.md`](references/architectural-principles.md) — SOLID, KISS, YAGNI.
- [`references/cqrs-and-events.md`](references/cqrs-and-events.md) — Command and event design.
- [`references/evolutionary-architecture.md`](references/evolutionary-architecture.md) — Fitness Functions and evolution.
- [`references/adr-template.md`](references/adr-template.md) — Official ADR template.




---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.372120Z"
evidence_checksum: "NONE"
```
