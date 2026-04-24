# GEMINI Project Instructions

<!-- GLOBAL_MANDATES_START -->

# Global Engineering Mandates

These instructions are fundamental and must be followed by all agents operating in this Hub.

## 🔒 0. SESSION BOOTSTRAP (EXECUTE BEFORE ANY RESPONSE)
Before responding to the user, the agent **MUST** perform this mental and operational checklist:
1. **Rehydrate Context**: Read `.specs/project/STATE.md`, `MEMORY.md`, and `LEARNINGS.md`.
2. **Onboarding Check**: Consult `onboarding-navigator` to align with the project culture.
3. **Task Sizing**: Classify task complexity (Quick, Small, Medium, Large, Complex) according to the SDD table.
4. **Skill Matching**: Consult the **Skill Router** below to select the appropriate tools.
5. **SDD Verification**: If it is a development task, validate if `spec.md` and `plan.md` exist.

## 📍 SKILL ROUTER
Use this guide to identify the mandatory skill for each context:

| If the task involves... | USE this skill |
|-------------------------|----------------|
| Project/Feature Start   | `onboarding-navigator` |
| Specification & Planning | `sdd` (`orchestrator`, `planner`) |
| Python Development (FastAPI/Django) | `python-uv` + `fastapi-expert` / `django-expert` |
| Frontend Development (React/Next) | `frontend-expert` |
| Mobile Development (Flutter) | `flutter-fvm` |
| Backend Development (Go) | `golang-expert` |
| Architecture & ADRs     | `architecture` + `api-architect` |
| Quality & Clean Code    | `clean-code-mentor` |
| Observability & SRE     | `observability-expert` |
| Brainstorming & Ideation | `brainstorming` |
| Automation & Scaffolding | `scaffolding-expert` |
| Knowledge Management    | `knowledge-architect` |
| State Synchronization   | `harness-expert` |
| Git & Commit Management | `git-workflow` |

## 1. SDD Framework (Mandatory for Development)
Any construction, development, or significant refactoring task **MUST** utilize the **SDD (Spec-Driven Development)** framework from the initial planning stage.
- **Workflow & Persistence**: Upon technically completing a task, the agent **MUST** proactively perform SDD Phase 4 (Review & Persistence).
  - **Update Specs**: Mark tasks as completed in `tasks.md` and update `plan.md`.
  - **Reports**: Generate or update `validation-report.md`.
  - **State**: Keep `STATE.md`, `MEMORY.md`, and `LEARNINGS.md` updated via **HB CLI**.

## 2. Architectural Integrity
- **Architecture-First**: Critical decisions must be recorded in an **ADR** in the `.specs/architecture/` folder.
- **Skill Standardization**: New skills must be validated via `skill-factory-validator`.
- **Diagram-as-Code**: Technical drawings must use **Mermaid** integrated into Markdown.

## 3. Quality Standards (Clean Code)
- Strictly follow **SOLID, YAGNI, DRY, and KISS**.
- Prioritize simplicity and maintainability; avoid over-engineering.
- All code must pass automated validation (linters/tests).

## 4. Operational Infrastructure (Harness Expert)
- **State Synchronization**: The agent **MUST** ensure that the operational state (long-term memory) is synchronized via `harness-expert` at the end of each iteration.
- **Context Efficiency**: Proactively use context compression to avoid token saturation, keeping `STATE.md` as the progress compass.

## 5. Knowledge Management (Knowledge Architect)
- **Local Knowledge Graph (LKG)**: Every new feature or significant architectural change **MUST** be mapped in `KNOWLEDGE-MAP.mermaid`.
- **Relationship Mapping**: Before major refactors, the agent **MUST** use `knowledge-architect` to analyze the impact on existing entities and relationships.

## 6. Mandatory Skill Usage
- **Skill-Driven Execution**: The agent **MUST** identify and use the most appropriate skill for each task. Executing any technical activity without the support of a specific skill or the Hub's automation framework is prohibited.

## 7. Dynamic Scaffolding (Scaffolding Expert)
- **Zero-Boilerplate Policy**: The agent **SHOULD** prefer using template CLI tools (like `uvx copier` via `scaffolding-expert`) to initialize the structure of projects and complex files (like `pyproject.toml`, Dockerfiles). Boilerplate code should NEVER be written line by line if a scaffolding template is available.

## 8. Autonomous Skill Discovery & Installation
- **On-Demand Intelligence**: If a task requires a specialized domain where no local skill is installed, the agent **MUST** proactively use `hb listskills` to discover if a relevant skill exists in the remote Hub.
- **Just-in-Time Setup**: If found, the agent **MUST** use `hb install <skill> --remote` to equip itself with the necessary intelligence before starting execution.

## 9. Version Control & Git (Git Workflow)
Every agent **MUST** strictly follow the `git-workflow` skill for any versioning operation:
- **English-Only Commits**: All commit messages must be written in **English**.
- **Conventional Commits**: Use of the Conventional Commits standard is mandatory.
- **SDD Linking**: Commits must be atomic and, whenever possible, reference task IDs from `tasks.md`.
- **History Integrity**: Using `rebase` is mandatory to maintain a linear history before merging into `main`.

## 🔒 10. SESSION EXIT GATE (EXECUTE BEFORE ENDING)
Before ending the session or delivering the task, the agent **MUST** validate:
1. **Tasks Update**: Does `tasks.md` reflect the real state of implementation?
2. **State Sync**: Has `STATE.md` been updated with progress and next steps?
3. **Learnings Capture**: Have new patterns or fixed bugs been added to `LEARNINGS.md`?
4. **Validation**: Has the code passed linters/tests and received a score in `validation-report.md`?
5. **Knowledge Update**: Does `KNOWLEDGE-MAP.mermaid` need an update?


<!-- GLOBAL_MANDATES_END -->

--- 
*Mandate updated via HB-CLI.*