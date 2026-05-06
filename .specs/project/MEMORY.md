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
- **Hardened SDD Governance (2026-04-25)**: Established Zero-Tolerance for skipping Plan/Explore phases in Medium+ tasks. Every new skill must be linked to a `.specs/features/` folder.
- **ADK Alignment (2026-05-06)**: Integrated the 5-layer Agent Development Kit model (Memory, Knowledge, Guardrails, Delegation, Distribution) as the core architectural standard for the Hub.

## Git Workflow Conventions
- **Conventional Commits**: Mandatory for all commits in the repository. Messages must be written in **English**.
- **SDD Alignment**: Commits must be linked to task IDs in `tasks.md` when applicable.
- **PR Quality**: Every Pull Request must follow the official template and include the SDD quality checklist.
+
+## Technical UI Governance
+- **Premium Terminal UI (2026-04-27)**: Decidido pelo uso de ANSI Escape Codes puros (Save/Restore Cursor) em vez de frameworks de alto nível para permitir a "auto-injeção" de status bars no CLI `hb` sem quebrar o fluxo de logs e manter alta performance cross-platform.
