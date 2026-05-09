# Tasks: Governance Hardening

## Phase 1: Knowledge Mapping
- [x] Create root `KNOWLEDGE-MAP.mermaid` [ID: GH-1.1] [Evidence: b134813]
- [x] Link `KNOWLEDGE-MAP.mermaid` in `README.md` [ID: GH-1.2] [Evidence: b134813]
- [x] Link `KNOWLEDGE-MAP.mermaid` in `GLOBAL_MANDATES.md` [ID: GH-1.3] [Evidence: b134813]

## Phase 2: Audit Automation
- [x] Create `bin/audit-governance.py` script [ID: GH-2.1] [Evidence: b134813]
- [x] Add `audit` target to `Makefile` [ID: GH-2.2] [Evidence: b134813]
- [x] Execute `make audit` and identify violations [ID: GH-2.3] [Evidence: b134813]

## Phase 3: Mandate Hardening
- [x] Update `GLOBAL_MANDATES.md` Section 7 (Exit Gate) to include audit [ID: GH-3.1] [Evidence: b134813]
- [x] Update `GLOBAL_MANDATES.md` Section 📍 (Skill Router) with Chain Loading notes [ID: GH-3.2] [Evidence: b134813]
- [x] Run `make sync` to propagate mandates [ID: GH-3.3] [Evidence: b134813]

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "GOVERNANCE-HARDENING"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-09T12:09:00Z"
evidence_checksum: "b134813"
```
