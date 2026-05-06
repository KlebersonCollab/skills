# Specification: Refactor Skill Factory to Purist v2.3.0

## 🎯 Objective
Align the entire `skill-factory` ecosystem with the SDD v2.3.0 Purist standard, ensuring zero fragmentation of memory and elimination of legacy requirements (like local `tasks.md` in skill directories).

## 📝 User Stories
- **US1**: As a developer, I want the Skill Factory to generate skills that are 100% compliant with centralized memory mandates.
- **US2**: As a maintainer, I want the Skill Factory Validator to enforce governance standards without imposing legacy/redundant file requirements.

## ✅ Acceptance Criteria (AC)
- [x] **AC1**: `SKILL.md` updated to v2.3.0, focusing strictly on logic and delegating all state/memory to `.specs/project/`.
- [x] **AC2**: `skill-factory-bootstrap` template updated to remove local `tasks.md` and enforce centralized memory references.
- [x] **AC3**: `skill-factory-validator` checklist updated to remove `tasks.md` requirement and add strict `<!-- @sdd-state -->` verification for all files.
- [x] **AC4**: All reference documentation (`references/`) updated to point to the correct global memory structures.
- [x] **AC5**: `examples/gold-standard` updated to serve as a perfect reference for a v2.3.0 compliant skill.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "REFACTOR-SKILL-FACTORY-V230"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T10:41:00Z"
evidence_checksum: "NONE"
```
