# Spec: Hub Purity Alignment (Core & Quality)

## 1. Goal
Finalize the transition to the "Purist" SDD model by removing legacy CLI dependencies from the core framework and critical skills.

## 2. Requirements

### [FR-1] Core SDD Purity
- **Description**: Remove all references to the `hb` CLI tool (e.g., `hb sdd status`, `hb learn`) from the central `sdd/SKILL.md`.
- **Reason**: The Purist standard mandates logic-first, manual governance via Markdown artifacts to prevent hidden dependencies.

### [FR-2] Quality Skill Alignment
- **Description**: Refactor `clean-code-mentor/SKILL.md` to remove the "Mandatory Tooling" section that requires HB CLI.
- **Reason**: Standardize the quality review process to follow the logic-first approach.

## 3. Acceptance Criteria
- [AC-1] `sdd/SKILL.md` must describe progress monitoring via `tasks.md` and `STATE.md` instead of CLI commands.
- [AC-2] `clean-code-mentor/SKILL.md` must producing `Review Report` and `Refactoring Diff` artifacts without relying on automated audits.
- [AC-3] All modified files must be at version 2.2.x or higher to indicate SDD compliance.
