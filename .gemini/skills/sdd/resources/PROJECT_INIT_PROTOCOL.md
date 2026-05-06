# Project Initialization Protocol

Use this protocol when starting a new project or professionalizing an existing one.

## 1. Context Rehydration
Before creating any files, read all available documentation and scan the codebase.

## 2. Core Project Files (.specs/project/)

Generate the following 6 files using `sdd-planner`:

1.  **PROJECT.md**: Define the "North Star". Vision, goals, and target audience.
2.  **ROADMAP.md**: High-level milestones and feature status.
3.  **STATE.md**: Initialize with "Project Setup" as the current task.
4.  **MEMORY.md**: Document the stack and user preferences discovered.
5.  **LEARNINGS.md**: (Empty initially) Ready to capture future insights.
6.  **DECISIONS.md**: Log the initial architectural choices.

## 3. Core Codebase Files (.specs/codebase/)

Generate the following files using `sdd-explorer`:

1.  **STACK.md**: Automated scan of manifests.
2.  **ARCHITECTURE.md**: Structural mapping of the directory.
3.  **CONVENTIONS.md**: Pattern extraction from sample files.
4.  **CONCERNS.md**: Risk audit (Technical Debt, Fragile code).
5.  **TECHNICAL-MAP.md**: Consolidated index.

## 4. Continuity
Update `STATE.md` after this initialization to define the first development feature.
