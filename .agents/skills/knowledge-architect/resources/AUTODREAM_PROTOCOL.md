# Auto-Dream Protocol: Background Consolidation

The "Dream State" is a process of background reflection used to optimize the Hub's institutional memory.

## 1. Objectives
- **De-clutter**: Remove redundant or temporary instructions from `MEMORY.md`.
- **Merge**: Combine similar `LEARNINGS.md` entries into unified technical patterns.
- **Sync**: Ensure the `KNOWLEDGE-MAP.mermaid` accurately reflects the source of truth.

## 2. Dream Workflow
When a "Dream" session is initiated:
1.  **Reflection turn**: The agent performs a "mental walk" through all `.specs/` and `MEMORY.md` files.
2.  **Drafting Changes**: Prepare a list of proposed deletions and merges.
3.  **User Approval**: Present the summary to the user: "I've identified 3 redundant patterns in our learnings. Should I merge them?"
4.  **Application**: Execute the clean-up.

## 3. Best Practices
- **Never delete unique insights**: Only merge patterns that are > 80% identical.
- **Archive vs. Delete**: Prefer moving obsolete data to a `HISTORY.md` or `.specs/archive/` rather than hard deletion.
- **Maintain Traceability**: When merging patterns, keep references to the original conversation IDs if possible.
