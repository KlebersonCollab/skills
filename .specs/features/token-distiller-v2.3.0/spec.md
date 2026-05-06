# Specification: Token Distiller v2.3.0 Upgrade

## 1. Context & Objective
The `token-distiller` skill is currently at version 2.2.0. To comply with the latest project standards (SDD v2.3.0 - Observable Governance), it must be refactored to include structured metadata and align with the centralized memory mandates.

## 2. Requirements (Traceable)
- **FR-1**: Update `SKILL.md` frontmatter version to `2.3.0`.
- **FR-2**: Implement Observable Governance metadata block (`<!-- @sdd-state -->`) in all core artifacts.
- **FR-3**: Ensure `SKILL.md` correctly references the centralized Triad of Memory.
- **FR-4**: Update `CHANGELOG.md` documenting the transition to v2.3.0.
- **FR-5**: Remove local `tasks.md` to prevent context fragmentation.

## 3. Acceptance Criteria (AC)
- [ ] `SKILL.md` version is `2.3.0`.
- [ ] All `.md` files in `token-distiller/` (except references/examples) contain the `@sdd-state` block.
- [ ] `token-distiller/tasks.md` is deleted.
- [ ] Project-wide `make audit` passes with 100% compliance.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "TOKEN-DISTILLER-V2.3.0"
phase: "IMPLEMENT"
status: "COMPLETED"
last_update: "2026-05-06T12:55:00Z"
evidence_checksum: "NONE"
```
