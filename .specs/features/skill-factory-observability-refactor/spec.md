# Spec: Skill Factory Observability Refactor (v2.3.0)

## 1. Context
The `skill-factory` is the engine that generates other skills. To maintain repository-wide observability, it must be updated to produce skills that follow the SDD v2.3.0 standard.

## 2. Requirements
- **FR-01**: `skill-factory/SKILL.md` must be updated to v2.3.0 with Observable Governance mandates.
- **FR-02**: `skill-factory-bootstrap` must generate skills with the `<!-- @sdd-state -->` metadata block.
- **FR-03**: `skill-factory-validator` must include checks for metadata blocks and evidence-based task lists.
- **FR-04**: All internal `skill-factory` artifacts must follow the new standard.

## 3. Acceptance Criteria (ACs)
- [ ] `skill-factory/SKILL.md` updated to v2.3.0.
- [ ] `skill-factory-bootstrap` Gold Standard Template updated with metadata placeholders.
- [ ] `skill-factory-validator` checklist includes metadata and evidence checks.
- [ ] Self-validation: The `skill-factory` passes its own v2.3.0 audit.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "FACTORY-REF-01"
phase: "SPECIFY"
status: "COMPLETED"
last_update: "2026-05-06T09:38:00Z"
evidence_checksum: "NONE"
```
