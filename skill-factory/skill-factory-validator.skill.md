---
name: skill-factory-validator
version: 2.3.0
description: "Validator agent for Skill Factory. Audits a skill directory against SDD v2.2.0 compliance and emits a pass/fail report."
category: skill-management
---

## 🔒 Prerequisites (Mandatory)
Before execution:
0. **Context Check**: Read `.specs/project/STATE.md`, `.specs/project/MEMORY.md`, and `.specs/project/LEARNINGS.md`.
1. **Target Access**: Ensure the target skill directory is accessible.

---

# Skill Factory: Validator Agent (v2.2.0)

> "Verification is the final gate for autonomy."

---

## 📋 Validation Checklist

### 1. Structural Integrity (Root Files)
- [ ] `SKILL.md` (Main entry point)
- [ ] `README.md` (General documentation)
- [ ] `tasks.md` (Task tracker with Evidence column)
- [ ] **Metadata Audit**: Every `.md` file must contain the `<!-- @sdd-state -->` block.
- [ ] **Absence of Local Memory**: No `STATE.md`, `MEMORY.md`, or `LEARNINGS.md` in the skill folder.

### 2. SDD Compliance (Protocol Check)
- [ ] **Prerequisites Lock**: Does `SKILL.md` start with `## 🔒 Prerequisites (Mandatory)`?
- [ ] **Delegation Matrix**: Is there a visual matrix mapping phases to sub-skills?
- [ ] **4-Phase Workflow**: Are the 4 phases (Discovery, Specify, Implement, Verify) explicitly defined?
- [ ] **Knowledge Chain**: Is the `Knowledge Verification Chain` present and correctly ordered?

### 3. "Anti-Slop" Quality Audit
- [ ] **Density**: Instructions are concise and free of filler text.
- [ ] **Explicit Logic**: No references to "magic" CLI commands; protocols are Markdown-based.
- [ ] **Placeholder Free**: No "todo", "...", or empty sections.

## ⚖️ Scoring Verdict
- **Approved (Gold Standard)**: 100/100 (All checks pass).
- **Needs Polish**: 70-99 (Minor structural or documentation fixes).
- **Rejected**: < 70 (Missing core protocols, metadata, or evidence assets).
- **Critical Failure**: Any missing `<!-- @sdd-state -->` block results in immediate rejection (Score: 0).

## 🛠️ Remediation Protocol
If a skill is rejected, the validator must provide a specific list of missing protocols and a path to align with the **SDD v2.2.0 Gold Standard**.

---

## 🚫 Prohibited
- NEVER approve a skill with local `STATE.md`, `MEMORY.md`, or `LEARNINGS.md` files.
- NEVER skip the **Prerequisites Lock** check.
- NEVER modify the skill files during validation.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T09:48:00Z"
evidence_checksum: "NONE"
```
