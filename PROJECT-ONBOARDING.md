# AI Agent Skills Hub - Project Governance (Purist SDD)

This document explains how to adopt the governance and intelligence of this Hub in your projects using the **Purist SDD v2.2.0** standard.

## 1. Governance Architecture

We follow a **Logic-First** approach. Governance is not enforced by a closed CLI, but by transparent Markdown artifacts and strict operational mandates.

### The Triad of Memory
Every project MUST maintain the following structure in `.specs/project/`:
- `STATE.md`: Captures the current active feature, progress, and next steps.
- `MEMORY.md`: Stores long-term facts, user preferences, and architectural decisions.
- `LEARNINGS.md`: Documents technical insights, bug fixes, and discovered patterns.

## 2. Project Bootstrapping (Manual)

To initialize a project with this governance:

1.  **Create the Specs Directory**:
    ```bash
    mkdir -p .specs/project .specs/architecture .specs/features
    ```
2.  **Initialize Memory**:
    Copy the templates or create empty files for `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`.
3.  **Define Global Mandates**:
    Link to or copy the `GLOBAL_MANDATES.md` from this Hub to your project's `.specs/codebase/`.

## 3. Adopting Skills

Since we no longer rely on a specialized installer CLI, you can "equip" your agent by:

1.  **Root Reference**: Directing your agent to read the `SKILL.md` and `references/` folders from this repository.
2.  **Local Mirroring**: Copying the desired skill folder (e.g., `python-uv/`) to your project root.
3.  **System Instruction**: Adding the core mandates to your agent's system prompt or `.cursorrules`.

## 4. The SDD Workflow

All development MUST follow the 4-phase SDD cycle:

1.  **DISCOVERY**: Map the codebase and rehydrate context from the Triad of Memory.
2.  **SPECIFY**: Create `spec.md`, `plan.md`, `contract.md`, and `tasks.md` in `.specs/features/[feature]`.
3.  **IMPLEMENT**: Execute tasks atomically with TDD and individual commits.
4.  **REVIEW**: Audit against Acceptance Criteria and update the Triad of Memory.

## 5. Quality Standards

- **Conventional Commits**: Mandatory for clear history.
- **Mermaid Diagrams**: Mandatory for visualizing complex flows in `plan.md`.
- **Atomic Progress**: Never mark a task as done without a corresponding commit and validation.
