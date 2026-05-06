# Skill Factory: Validator Sub-skill

This sub-skill performs a quality audit on existing or newly created skills to ensure they meet the Hub's "Gold Standard".

## 🔒 Prerequisites
- Target skill path must be provided.

## 📋 Audit Checklist

### 1. Structural Integrity (Score: 40 pts)
- [ ] Presence of `SKILL.md` (10 pts)
- [ ] Presence of Memory Triad (`STATE.md`, `MEMORY.md`, `LEARNINGS.md`) (10 pts)
- [ ] Presence of `.specs/` directory (10 pts)
- [ ] No empty files (Files > 0 bytes) (10 pts)

### 2. SDD Compliance (Score: 40 pts)
- [ ] `SKILL.md` contains a Delegation Matrix (10 pts)
- [ ] `SKILL.md` contains a Knowledge Verification Chain (10 pts)
- [ ] `SKILL.md` defines a 4-Phase Workflow (10 pts)
- [ ] Frontmatter version is Semantic (e.g., 1.0.0) (10 pts)

### 3. "Slop-Free" Design (Score: 20 pts)
- [ ] Descriptions are clear and actionable (10 pts)
- [ ] Instructions avoid placeholders like "todo" or "..." (10 pts)

## ⚖️ Scoring Verdict
- **90 - 100**: ✅ APPROVED (Gold Standard)
- **70 - 89**: ⚠️ NEEDS POLISH (Minor fixes required)
- **< 70**: ❌ REJECTED (High Slop Detected)

## 🛠️ Remediation
If a skill fails the audit, the validator must provide a specific list of missing items or patterns to be corrected.
