# Spec: Project Health Check

## 1. Context
We need a way to quickly verify if the project is following the SDD v2.3.0 standards (Observable Governance).

## 2. Requirements
- **FR-1**: Check if `.specs/project/STATE.md` exists.
- **FR-2**: Check if `STATE.md` contains the `<!-- @sdd-state -->` block.
- **FR-3**: Output a "Success" or "Failure" message with details.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HEALTH-CHECK-01"
phase: "SPECIFY"
status: "COMPLETED"
last_update: "2026-05-06T09:36:00Z"
evidence_checksum: "NONE"
```
