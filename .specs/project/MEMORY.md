# Project Persistent Memory

## Skill Conventions
- **Marker**: A folder is considered a skill if it contains a `SKILL.md` file in its root.
- **Standard Structure**: A skill must contain `CHANGELOG.md`, `README.md`, `SKILL.md` and optionally `references/` and `resources/` folders.
- **Mandatory Hook**: Execution skills MUST contain the `🔒 Prerequisites (Mandatory)` section linking them to SDD.

## Agent Mapping
- **Claude**: `.claude/` folder, `CLAUDE.md` master file.
- **Gemini**: `.gemini/` folder, `GEMINI.md` master file.
- **Agent**: `.agent/` folder, `AGENT.md` master file.

## Distribution Pipeline
- Artifacts are generated dynamically by collecting skills from the root.
- The final structure in the ZIPs preserves the agent's folder name (e.g., `.claude/`) to facilitate direct installation.

## Operational Governance
- **Mandatory Use of Skills**: As per `GLOBAL_MANDATES.md`, any technical execution MUST be mediated by a specific Hub skill.
- **Deterministic Lifecycle**: Every agent operating in the Hub must mandatorily execute the **Session Bootstrap** (start) and the **Session Exit Gate** (end).

## Git Workflow Conventions
- **Conventional Commits**: Mandatory for all commits in the repository. Messages must be written in **English**.
- **SDD Alignment**: Commits must be linked to task IDs in `tasks.md` when applicable.
- **PR Quality**: Every Pull Request must follow the official template and include the SDD quality checklist.
