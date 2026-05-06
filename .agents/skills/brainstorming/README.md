# Brainstorming — Skill

Brainstorming and Design Facilitator — guides the agent to explore complex problems, generate divergent ideas, and converge on solid specifications before any implementation.

---

## 🚀 Overview

This skill transforms the agent into a design thinking facilitator. The goal is not to code fast, but to **think slow** to ensure the solution built is the correct one. It uses established techniques to break through creative blocks and validate assumptions before spending resources on implementation.

### Key Capabilities:
- **Disciplined Discovery**: Identifying the root cause via "5 Whys."
- **Divergent Generation**: Exploration of multiple architectures/approaches (SCAMPER).
- **Strategic Convergence**: Evaluation of trade-offs and grounded decision-making.
- **Understanding Lock**: Security gate to ensure full alignment with the user.

---

## 📂 Skill Structure

```text
brainstorming/
├── SKILL.md                 # Definition and workflow (v1.0.0)
├── README.md                # This guide
├── CHANGELOG.md             # Version history
├── references/              # Theoretical Knowledge
│   ├── brainstorming-techniques.md
│   └── decision-frameworks.md
└── resources/               # Templates (optional)
```

---

## 🛠️ How to Use

1. **Activate the skill**: `activate_skill brainstorming`
2. **Start Discovery**: "I'd like to brainstorm the new notification architecture."
3. **Follow the Facilitator**: The agent will ask questions one at a time to map the problem and assumptions.
4. **Validate the Understanding Lock**: Before proceeding to the solution, confirm the understanding summary sent by the agent.

---

## 🎯 Example Commands

- *"Let's brainstorm how to reduce the latency of our search API."*
- *"Use the SCAMPER technique to propose new features for our dashboard."*
- *"What are the trade-offs between using a relational or NoSQL database for this use case?"*
