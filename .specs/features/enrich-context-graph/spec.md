# Spec: Enrichment of Context Graph with MemDir Protocol

## 1. Goal
Integrate high-maturity memory management patterns (mined from `src/memdir`) into the `context-graph` skill. This will transform it from a conceptual memory skill into a technically rigorous "Memory Operating System".

## 2. Requirements (FR)
- **FR-1**: Implement the "Index vs Topic" architecture. `MEMORY.md` must only contain pointers.
- **FR-2**: Enforce strict size limits: 200 lines or 25KB for the index file.
- **FR-3**: Standardize memory file frontmatter (name, description, type).
- **FR-4**: Define the taxonomy of memory types: `User`, `Project`, `Local`, `Managed`, `AutoMem`.
- **FR-5**: Provide guidance on the persistence hierarchy (Memory vs Plan vs Tasks).

## 3. Acceptance Criteria (AC - BDD)

### Scenario 1: Index Truncation Logic
**Given** that a project's `MEMORY.md` exceeds 200 lines
**When** the agent audits the memory directory
**Then** it must flag the index as "Truncated" and recommend moving detailed content to topic files.

### Scenario 2: Saving New Memories
**Given** a new insight about the user's preference
**When** saving to memory
**Then** the agent must create a new topic file with YAML frontmatter
**And** add a single-line pointer (markdown link) to `MEMORY.md`.

## 4. Constraints
- Maintain version 1.1.0 for the skill.
- Must use the precise byte limits (25,000 bytes) identified in the mining phase.
