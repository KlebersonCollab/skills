# Specification: Teleport Knowledge (hb teleport)

## Goal
Implement `hb teleport` to serialize and export feature-specific knowledge (Learnings, Decisions, Memory) to a global repository, enabling cross-project knowledge inheritance.

## Acceptance Criteria (BDD)
### Scenario 1: Exporting Feature Knowledge
- **Given** I am working on a feature with `LEARNINGS.md` and `MEMORY.md`
- **When** I run `hb teleport export`
- **Then** the command should pack these files into a "Knowledge Bundle"
- **And** it should save/upload it to the central knowledge hub.

### Scenario 2: Global Knowledge Inheritance
- **Given** a new project or feature
- **When** I run `hb teleport import <feature-source>`
- **Then** it should extract the learnings and append them to the local `MEMORY.md` to avoid repeating past mistakes.

## Constraints
- Serialization format: JSON or Markdown Bundles.
- Default central path: `.specs/knowledge/global/`.
