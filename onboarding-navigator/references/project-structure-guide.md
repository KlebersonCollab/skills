# Project Structure Guide: The Source of Truth

This guide explains the project's physical and logical organization to ensure efficient navigation and compliance with the SDD workflow.

---

## 1. Global Hierarchy

```text
skills/
├── .specs/                      # The Heart of SDD
│   ├── project/                 # Vision, Roadmap, and Memory Triad
│   ├── codebase/                # Stack, Conventions, and Global Architecture
│   └── features/                # Detailed specifications per functionality
├── .gemini/ | .claude/ | .agent/ # Agent Configurations and Mandates
├── <skill-name>/                # Skill Directories (e.g., python-uv)
│   ├── SKILL.md                 # Technical definition
│   ├── references/              # Support guides
│   └── examples/                # Real samples
├── README.md                    # Central skill registry
└── .gitignore                   # Versioning rules (includes Agent Mandates)
```

## 2. The Role of the `.specs/` folder

Here resides all project strategic intelligence. Never start a change without consulting this directory.
- **`STATE.md`**: What is happening now?
- **`MEMORY.md`**: What have we learned and should not forget?
- **`LEARNINGS.md`**: Solutions for bugs and discovered technical patterns.

## 3. Agent Directories

These directories contain the **`GEMINI.md`**, **`CLAUDE.md`**, and **`AGENT.md`** files.
- They are mandatory and define how the AI agent should behave.
- **Important**: These `.md` files are versioned (not ignored), ensuring all agents follow the same quality rules.

## 4. Where to find everything?

| Need (What do you want to do?) | Location / Recommended Skill |
|-------------------------------------|---------------------------------|
| **Create a new skill in the Hub** | `skill-factory/` |
| **Orchestrate an Agent Swarm** | `swarm-facilitator/` |
| **Consult/Create ADRs and Macro Architecture** | `architecture/` or `.specs/codebase/` |
| **Design Contracts and APIs (REST/GraphQL)** | `api-architect/` |
| **Ensure/Verify Clean Code rules** | `clean-code-mentor/` |
| **Save state and Manage Session Memory** | `harness-expert/` or `.specs/project/STATE.md` |
| **Map and Visualize dependency graphs** | `knowledge-architect/` or `KNOWLEDGE-MAP.mermaid` |
| **Generate base projects / Boilerplates from scratch** | `scaffolding-expert/` |
| **Configure CI/CD, Boards, or Azure Pipelines** | `azure-devops/` |
| **Plan next sprint or Roadmap** | `.specs/project/ROADMAP.md` |

## 5. Navigation Tip
Always start your session by asking the agent to "Rehydrate the context by reading the memory files in the `.specs/` folder". This ensures it knows exactly where we left off.
