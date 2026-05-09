# Feature Specification: Governance Hardening & Global Knowledge Hub

## 🎯 Goal
Implement "Governance as Code" by automating the validation of engineering mandates and centralizing the project's knowledge map to ensure strict compliance with SDD v2.3.0.

## 📋 Requirements

### FR-1: Centralized Knowledge Map
- **FR-1.1**: Create a `KNOWLEDGE-MAP.mermaid` file in the project root.
- **FR-1.2**: The map must visualize the relationships between Core Skills (`sdd`, `brainstorming`, `architecture`) and Domain Skills.
- **FR-1.3**: The map must be linked in `README.md` and `GLOBAL_MANDATES.md`.

### FR-2: Automated Governance Audit
- **FR-2.1**: Implement a Python script `bin/audit-governance.py`.
- **FR-2.2**: The script must validate that all `.md` files in `.specs/` and skill folders contain a valid `<!-- @sdd-state -->` block.
- **FR-2.3**: The script must check if `evidence_checksum` is provided for tasks marked as completed.
- **FR-2.4**: The script must verify the presence of mandatory files (`STATE.md`, `MEMORY.md`, `LEARNINGS.md`) in `.specs/project/`.

### FR-3: Mandate Update
- **FR-3.1**: Update `GLOBAL_MANDATES.md` to include the execution of `make audit` as a mandatory step in the "Session Exit Gate".
- **FR-3.2**: Add the "Chain Loading" concept to the Skill Router.

## ✅ Acceptance Criteria (AC)
- [x] `KNOWLEDGE-MAP.mermaid` exists and is renderable.
- [x] `bin/audit-governance.py` runs and correctly identifies missing metadata blocks.
- [x] `Makefile` includes an `audit` command.
- [x] `GLOBAL_MANDATES.md` contains the new automation mandates.
- [x] All skills are mapped in the central Knowledge Hub.

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
