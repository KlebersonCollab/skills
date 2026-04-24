# Architecture — Skill

Systems Architect — guides the agent to design scalable, simple, and well-documented systems, using ADRs and rigorous trade-off analysis.

---

## 🚀 Overview

This skill transforms the agent into a pragmatic software architect. The focus is not just on using cutting-edge technologies, but on ensuring that the solution is as simple and maintainable as possible for the given context. It formalizes the technical decision-making process, ensuring that the "why" behind each architectural choice is recorded and validated.

### Key Capabilities:
- **Contextual Discovery**: Mapping technical, financial, and team constraints.
- **Trade-off Analysis**: Objective comparison between competing approaches.
- **Pattern Selection**: Informed choice of styles (Hexagonal, Microservices, Monolith, etc.).
- **Governance via ADR**: Durable documentation of architectural decisions.

---

## 📂 Skill Structure

```text
architecture/
├── SKILL.md                 # Definition and workflow (v1.0.0)
├── README.md                # This guide
├── CHANGELOG.md             # Version history
├── references/              # Theoretical Knowledge
│   ├── architectural-principles.md
│   ├── pattern-selection.md
│   └── adr-template.md
└── resources/               # Generated ADRs (suggested)
```

---

## 🛠️ How to Use

1. **Activate the skill**: `activate_skill architecture`
2. **Define the Context**: "I want to design the architecture for the new payment processing system."
3. **Analyze Alternatives**: The agent will propose different paths (e.g., Sync vs. Async). Evaluate the listed trade-offs.
4. **Generate the ADR**: After choosing the path, the agent will document the decision in the ADR format for future reference.

---

## 🎯 Example Commands

- *"Analyze the trade-offs between using a relational database and a NoSQL database for this product catalog."*
- *"How to structure this project following hexagonal architecture to facilitate testing?"*
- *"Create an ADR justifying the migration to an event-driven architecture."*
- *"What are the architectural constraints we should consider for a system that needs very low latency?"*
