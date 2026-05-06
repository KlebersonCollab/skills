# Specification: Dynamic Health Check & Observable Governance Audit

## 🎯 Objective
Refactor the audit system (`bin/check-health.py`) to dynamically discover skills and artifacts, ensuring 100% visibility of the project's compliance with SDD v2.3.0.

## 📝 User Stories
- **US1**: As a developer, I want the audit script to automatically find new skills so I don't have to update it manually.
- **US2**: As a maintainer, I want the audit to distinguish between "documentation of metadata" (templates) and "actual metadata" (state).

## ✅ Acceptance Criteria (AC)
- **AC1**: `bin/check-health.py` must scan the root directory and all subdirectories for `SKILL.md`, `README.md`, and `CHANGELOG.md`.
- **AC2**: The script must include `youtube-transcript` and any other newly created skills.
- **AC3**: Metadata extraction must prioritize the *real* state block (usually at the very end of the file) and ignore blocks inside code fences (like in documentation).
- **AC4**: `sdd/SKILL.md` must have a real metadata block added.
- **AC5**: The audit result must be 100% compliant after the fix.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "DYNAMIC-HEALTH-CHECK"
phase: "SPECIFY"
status: "IN_PROGRESS"
last_update: "2026-05-06T10:10:00Z"
evidence_checksum: "NONE"
```
