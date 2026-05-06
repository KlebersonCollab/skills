# Skill Factory: Validator Sub-skill (Purist Edition)

The Validator ensures that structural purity is maintained. It audits the logic, not just the existence of files.

## 🔒 Prerequisites
- A target skill for audit.

## 📋 Logical Audit Checklist

### 1. Architectural Clarity (40 pts)
- [ ] **Is the brain visible?** `SKILL.md` must have a clear Delegation Matrix.
- [ ] **Is the history traceable?** `DECISIONS.md` and `STATE.md` must link to active tasks.
- [ ] **Is the logic explicit?** No instructions should refer to "magic" or "hidden" scripts.

### 2. SDD Compliance (40 pts)
- [ ] **Phase Integrity**: Does the skill follow the 4-phase SDD workflow?
- [ ] **Verification Chain**: Is the hierarchy of knowledge (Existing Code -> Specs -> etc.) explicitly defined?
- [ ] **Atomicity**: Are tasks small enough to be verified independently?

### 3. Anti-Slop Score (20 pts)
- [ ] **Density**: Every line of the skill must add value. No boilerplate fluff.
- [ ] **Precision**: Instructions must be actionable for an agent without ambiguity.

## ⚖️ Verdict
- **Approved**: High density, explicit logic, SDD compliant.
- **Reject**: Contains slop, hidden logic, or vague instructions.
