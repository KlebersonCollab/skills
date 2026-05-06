# Technical Plan: Refactor Skill Factory to Purist v2.3.0

## 🛠️ Architecture Overview
The refactor will focus on "cleaning" the instruction set of the Skill Factory sub-agents to remove any "memory leakage" (local files) and hardening the governance audit.

## 📦 Changeset

### 1. `SKILL.md` (Main)
- Update version to 2.3.0.
- Clarify that ALL operational memory (tasks, state, learnings) resides in `.specs/project/`.
- Explicitly prohibit the creation of `tasks.md` within the skill directory.

### 2. `skill-factory-bootstrap.skill.md`
- Refactor the **Gold Standard Template** to remove `tasks.md` from the generated folder structure.
- Update the template's prerequisites to point to global paths.
- Harden the `Observable Governance` section to ensure `feature_id` is derived from the current feature spec.

### 3. `skill-factory-validator.skill.md`
- Remove the `tasks.md` check from the structural integrity list.
- Add a check to verify if the skill's prerequisites point to the global Tríade de Memória.
- Ensure 100% rejection for any skill containing local memory files.

### 4. `examples/gold-standard/`
- Update the example files to match the new v2.3.0 standard perfectly.

## 🧪 Verification Strategy
- Run `skill-factory-validator` against its own refactored directory.
- Verify that `make audit` passes with 100% compliance.
- Test generating a "dummy" skill to ensure the bootstrap template is correct.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "REFACTOR-SKILL-FACTORY-V230"
phase: "SPECIFY"
status: "IN_PROGRESS"
last_update: "2026-05-06T10:24:00Z"
evidence_checksum: "NONE"
```
