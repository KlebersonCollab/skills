# Specification: Python Patterns Integration

## Context
The project lacks a high-level architectural decision skill for Python development. A new "python-patterns" skill has been proposed to fill this gap, providing guidance on framework selection, async vs sync strategies, and project structure.

## Objective
Integrate the "python-patterns" skill into the Agent Skills Hub as a mandatory architectural gateway.

## Acceptance Criteria
- [ ] New skill `python-patterns` created in `.agents/skills/python-patterns/SKILL.md`.
- [ ] `python-patterns` skill correctly parsed and appearing in the Visual Dashboard (via metadata).
- [ ] `AGENTS.md` Skill Router updated to include `python-patterns` as the first step for Python development.
- [ ] `README.md` updated to include `python-patterns` in the "Languages & Frameworks" section.
- [ ] `.specs/codebase/GLOBAL_MANDATES.md` updated to include `python-patterns` principles as mandatory.
- [ ] `STATE.md` and `MEMORY.md` updated to reflect the new skill.

## Constraints
- Follow SDD v2.3.0.
- All technical mandates must be visible in Markdown.
- Commits must follow `git-workflow` (English, Conventional).

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "python-patterns-integration"
phase: "SPECIFY"
status: "IN_PROGRESS"
last_update: "2026-05-09T00:12:00Z"
evidence_checksum: "NONE"
```
