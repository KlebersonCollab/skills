# Plan: Standardize Token Distiller

## Architecture
This is a documentation-only refactor to align with the new governance standard.

## Proposed Changes

### 1. `token-distiller/SKILL.md` Refactor
- Update frontmatter to v2.2.0.
- Update `## 🔒 Prerequisites (Mandatory)` to point to `.specs/project/`.
- Add `## 🔄 4-Phase Workflow` defining how the skill is used (Discovery of context, Specification of compression, Implementation of modes, Verification of density).
- Add `## 🛠️ Operational Protocols` including the `Knowledge Verification Chain`.

### 2. Cleanup
- Delete `token-distiller/MEMORY.md`.
- Delete `token-distiller/LEARNINGS.md`.

## Verification Plan
1. Run `skill-factory-validator` on the `token-distiller` directory.
2. Verify ACs manually.
