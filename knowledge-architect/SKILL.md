---
name: knowledge-architect
version: 1.0.0
description: "Local knowledge architecture via relational graphs (Local GraphRAG). Transforms static documentation into a connection map for intelligent context navigation."
category: knowledge-management
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Did you rehydrate the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Meta-Knowledge Check**: Consult the **Context Graph** (`DECISIONS.md`) to understand the rationale behind existing relations.
3. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define the architecture, schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---

## 🔒 Mandatory Tooling
The use of **HB CLI** is **MANDATORY** for this skill:
- **Focus**: Use `hb project focus "Task"` to track mapping and relational analysis progress.
- **Knowledge Management**: Use `hb learn` to document new architectural patterns or bugs.
- **Visualization**: Use `hb map` to generate and update the `KNOWLEDGE-MAP.mermaid`.

---
# Knowledge Architect

> "Knowledge is not in the files, but in the connections between them." — Protocol for transforming linear documentation into a local relational graph.

---

## Goal

The goal of this skill is to empower the agent to build, maintain, and navigate a **Local Knowledge Graph (LKG)**. It allows the agent to understand complex relationships between requirements, architectural decisions (ADRs), code components, and session states, overcoming the limitations of purely linear or vector search.

---

## Workflow (4 Phases)

### Phase 1: EXTRACT — Entity Distillation
1.  **Node Identification**: Read project files and identify "entities" (e.g., Features, Components, Classes, ADRs, Requirements).
2.  **Relation Mapping**: Identify how these entities connect (e.g., `Feature A` **implemented by** `Class B`, `Decision C` **impacts** `Module D`).
3.  **Semantic Metadata**: Extract tags, versions, and critical dependencies for each node.

### Phase 2: LINK — Graph Construction
1.  **Mermaid Documentation**: Create or update the `KNOWLEDGE-MAP.mermaid` (the project's Mind Map).
2.  **Cross-Referencing**: Insert bidirectional links (`[[link]]`) in Markdown files to facilitate manual and automatic traversal.
3.  **Context Indexing**: Maintain an index of "close neighbors" for critical entities.

### Phase 3: TRAVERSE — Intelligent Navigation
1.  **Impact Traversal**: Before a change, follow graph connections to identify side effects in distant modules.
2.  **Context Discovery**: Use the graph to decide which *additional* files should be read (beyond the target file) to obtain the full view.
3.  **Irrelevance Pruning**: Ignore graph branches that have no semantic connection to the current task, saving tokens.

### Phase 4: PRUNE — Evolution and Cleanup
1.  **Knowledge Refactoring**: Remove obsolete connections or deleted entities from the code.
2.  **Cluster Consolidation**: Group highly connected nodes into "Knowledge Modules" to simplify the macro view.
3.  **Harness Synchronization**: 
    - Use `hb map` to generate the visual knowledge map.
    - Use `hb learn "Subject" --desc "..." --cat "..."` to record structured learnings and patterns.
    - Export the relation map so that `harness-expert` can perform relational rehydration.

### Phase 5: DREAM — Knowledge Consolidation (Background)
1.  **Trigger**: Every 5 sessions or after a `Phase 4: REVIEW` of a Large/Complex feature.
2.  **Self-Reflection**: The agent should suggest a "Dream State" to the user: "Would you like me to consolidate our recent learnings?"
3.  **Action**:
    - Review `LEARNINGS.md` and merge duplicated patterns.
    - Prune obsolete context from `MEMORY.md`.
    - Update the `KNOWLEDGE-MAP.mermaid` to reflect the final architectural state.
    - Generate a "State of the Hub" report.

---

## Output Structure

Execution of this skill results in the following mandatory artifacts (usually in `.specs/knowledge/` or the skill root):

| Artifact | Format | Description |
|----------|---------|-----------|
| **Knowledge Map** | `KNOWLEDGE-MAP.mermaid` | Visual diagram of all project entities and relationships. |
| **Relations Registry** | `RELATIONS.md` | Textual list of connections and their technical justifications. |
| **Context Clusters** | `CLUSTERS.json` | (Optional) Grouping of files that should always be read together. |

---

## Quality Rules

- **Gold Standard Benchmark**: Use the examples in `examples/gold-standard/` as a reference for complex mappings.
- **Relation over Content**: Focus on *how* things connect, not just what they do.
- **Visual-First**: Every complex graph must be viewable via Mermaid.
- **Bi-directional**: Whenever possible, relationships should be mapped both ways (who uses and who is used).
- **KISS Graphing**: Avoid creating nodes for trivial details; focus on architectural components and requirements.

## Prohibited

- NEVER create graphs that cannot be rendered (Mermaid syntax errors).
- NEVER maintain orphan relationships (nodes without relevant connections).
- NEVER use external graph database tools; everything must be persisted in text files in the repository.
- NEVER ignore map updates after major code refactorings.
- NEVER perform a "Dream" consolidation without user confirmation if it involves deleting information.

---

## 5. References
- [Local GraphRAG Patterns](resource:resources/GRAPHRAG_PATTERNS.md)
- [Auto-Dream Protocol](resource:resources/AUTODREAM_PROTOCOL.md)
