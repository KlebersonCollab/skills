# Spec: Standardize Token Distiller

Standardize the `token-distiller` skill to comply with `skill-factory` v2.2.0 (Purist SDD Gold Standard).

## Functional Requirements
- FR-1: Update `token-distiller/SKILL.md` to follow the v2.2.0 template (Prerequisites, Knowledge Chain, 4-Phase Workflow).
- FR-2: Eliminate local memory files (`STATE.md`, `MEMORY.md`, `LEARNINGS.md`) in `token-distiller/`.
- FR-3: Ensure `token-distiller` passes the `skill-factory-validator`.

## Acceptance Criteria
- AC-1: `SKILL.md` starts with `## 🔒 Prerequisites (Mandatory)` referencing `.specs/project/`.
- AC-2: `SKILL.md` contains a `Knowledge Verification Chain`.
- AC-3: `SKILL.md` explicitly defines the 4-Phase Workflow (Discovery, Specify, Implement, Verify).
- AC-4: Files `token-distiller/MEMORY.md` and `token-distiller/LEARNINGS.md` are deleted.
- AC-5: `skill-factory-validator` score is 100/100.
