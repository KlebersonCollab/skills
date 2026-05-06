# Specification: SDD Observability Refactor (v2.3.0)

## 1. Context & Problem
The current SDD implementation is "implicit" and relies on the agent's memory to track progress. This makes it difficult to audit the process or build visual tools to monitor state.

## 2. Requirements (FR - Functional Requirements)
- **FR-01**: All SDD artifacts (`spec.md`, `plan.md`, `STATE.md`, etc.) MUST contain a structured metadata block in YAML format.
- **FR-02**: The `tasks.md` file MUST include a mandatory "Evidence" field for every completed task.
- **FR-03**: The `sdd/SKILL.md` MUST define a "Logical Pipeline" where phases are explicitly linked to metadata states.
- **FR-04**: The system MUST remain human-readable (Markdown-first).

## 3. Acceptance Criteria (ACs)
- [ ] `sdd/SKILL.md` is updated to v2.3.0 with "Observable Governance" section.
- [ ] `resources/tasks-template.md` includes the `Evidence` column.
- [ ] `resources/spec-template.md` and `resources/plan-template.md` include metadata placeholders.
- [ ] A visual script can theoretically parse the new metadata (Validation test).

## 4. Metadata Schema (Proposed)
```yaml
<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "FEAT-XXX"
phase: "DISCOVERY | SPECIFY | IMPLEMENT | VERIFY"
status: "IN_PROGRESS | COMPLETED | BLOCKED"
last_update: "ISO-TIMESTAMP"
evidence_checksum: "GIT-HASH | LOG-PATH | NONE"
```
```
