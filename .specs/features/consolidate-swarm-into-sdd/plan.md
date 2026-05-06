# Plan: Consolidate Swarm Facilitator into SDD

## 🏗️ Architecture
We are moving from a **Fragmented Orchestration** model to a **Unified SDD Governance** model.

### Execution Strategy
1.  **Extraction**: Pull infra logic (Terminal, Security, Layout) from `swarm-facilitator/SKILL.md`.
2.  **Integration**: Create `sdd/references/swarm-execution.md` with the extracted logic.
3.  **Cross-Reference**: Update `sdd/SKILL.md` to point to the new reference under "Operational Protocols".
4.  **Deletion**: Remove `swarm-facilitator` directory.
5.  **Registry Sync**: Update global documentation.

## 🛠️ Implementation Steps
1.  **Task 1**: Create `sdd/references/swarm-execution.md`.
2.  **Task 2**: Update `sdd/SKILL.md`.
3.  **Task 3**: Delete `swarm-facilitator`.
4.  **Task 4**: Update root README and project state.
