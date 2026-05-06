# Plan: Skill Factory Observability Refactor

## 1. Architecture & Design
We are updating the "Genetic Template" of the Skill Hub. This is a recursive refactor: updating the tool that updates the tools.

## 2. Execution Steps

### Step 1: Update `skill-factory/SKILL.md`
- Bump version to `2.3.0`.
- Add "Observable Governance" section.

### Step 2: Refactor `skill-factory-bootstrap.skill.md`
- Update the "Gold Standard Template" to include:
    - Metadata block at the end of `SKILL.md`.
    - Mandatory table-based `tasks.md` with Evidence column.
- Bump version to `2.3.0`.

### Step 3: Refactor `skill-factory-validator.skill.md`
- Add structural checks for `<!-- @sdd-state -->`.
- Add audit rule for the `Evidence` column in `tasks.md`.
- Bump version to `2.3.0`.

### Step 4: Verification
- Use the new `skill-factory` to bootstrap a dummy skill and verify it has the correct metadata.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "FACTORY-REF-01"
phase: "SPECIFY"
status: "COMPLETED"
last_update: "2026-05-06T09:38:20Z"
evidence_checksum: "NONE"
```
