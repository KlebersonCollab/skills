# Architecture — Skill

Systems Architect — guides the agent to design scalable, simple, and well-documented systems, using ADRs and rigorous trade-off analysis.

---

## 🚀 Overview

This skill transforms the agent into a pragmatic software architect. The focus is not just on using cutting-edge technologies, but on ensuring that the solution is as simple and maintainable as possible for the given context. It formalizes the technical decision-making process, ensuring that the "why" behind each architectural choice is recorded and validated through Markdown artifacts.

### Key Capabilities:
- **Contextual Discovery**: Mapping technical, financial, and team constraints.
- **Trade-off Analysis**: Objective comparison between competing approaches.
- **Pattern Selection**: Informed choice of styles (Hexagonal, Microservices, Monolith, etc.).
- **Governance via ADR**: Durable documentation of architectural decisions in `.specs/architecture/`.

---

## 📂 Skill Structure

```text
architecture/
├── SKILL.md                 # Definition and workflow (v2.2.1)
├── README.md                # This guide
├── CHANGELOG.md             # Version history
├── references/              # Theoretical Knowledge
│   ├── architectural-principles.md
│   ├── cqrs-and-events.md
│   ├── evolutionary-architecture.md
│   ├── pattern-selection.md
│   └── adr-template.md
└── resources/               # Assets and templates
```

> [!IMPORTANT]
> **Centralized Memory**: This skill does NOT store its own `STATE.md`, `MEMORY.md`, or `LEARNINGS.md`. All operational state is maintained in `.specs/project/`.

---

## 🛠️ How to Use

1. **Prerequisites**: Ensure `.specs/project/STATE.md` is initialized.
2. **Define the Context**: "I want to design the architecture for the new payment processing system."
3. **Analyze Alternatives**: The agent will propose different paths (e.g., Sync vs. Async) in a Trade-off Matrix.
4. **Document via ADR**: Create a new ADR in `.specs/architecture/` using the template.
5. **Verify**: Perform a logic-based audit against the `spec.md` and `plan.md`.

---

## 🎯 Example Commands

- *"Analyze the trade-offs between using a relational database and a NoSQL database for this product catalog."*
- *"How to structure this project following hexagonal architecture to facilitate testing?"*
- *"Create an ADR justifying the migration to an event-driven architecture."*
- *"What are the architectural constraints we should consider for a system that needs very low latency?"*


