# Example: Decision Ledger (Context Graph)

This example shows how an agent should log a decision trace following the `context-graph` skill.

---

## [DT-2026-001] Choice of Storage Pattern for Knowledge Map

**Date**: 2026-04-25
**Agent**: `knowledge-architect`
**Status**: `Approved`

### Context
The project needs a way to store the relationship between 500+ files. We need to decide between a centralized JSON file or decentralized Mermaid snippets in each directory.

### Decision
Use **Mermaid in a central `KNOWLEDGE-MAP.mermaid`** supplemented by **bidirectional Markdown links** in the files.

### Rationale
- **Visualization**: Mermaid allows instant rendering in GitHub/VSCode.
- **Traceability**: Markdown links allow `grep` and LSP traversal without custom tools.
- **Scalability**: JSON would be easier for machines, but harder for humans to audit during the "Human-in-the-loop" phase.

### Alternatives Considered
1. **Centralized JSON**: Rejected because it's not "Human-Readable" without a viewer.
2. **Neo4j Instance**: Rejected as it violates the "Text-First" rule of the repository (Prohibited in `knowledge-architect`).

### Tribal Knowledge Applied
> "In this repository, we prioritize transparency and version control over complex binary database state."

### Expected Outcome
Improved developer experience during architectural reviews.

---

## [DT-2026-002] Naming Convention for Feature Modules

**Date**: 2026-04-25
**Agent**: `context-graph`
**Status**: `Pending Outcome`

### Decision
Adopt `kebab-case` for all folder names.

### Rationale
Consistency with existing skills (e.g., `token-distiller`, `git-workflow`).

### Outcome (Updated 2026-05-01)
**Result**: `Success`. No conflicts with OS-level case sensitivity.
