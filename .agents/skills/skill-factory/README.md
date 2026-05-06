# Skill Factory — Capability Architect

> "Structure is the bridge to autonomy. Precision in design, rigor in execution."

[![Version](https://img.shields.io/badge/Version-2.2.0-blue)](#changelog)
[![Sub-skills](https://img.shields.io/badge/Sub--skills-2-brightgreen)](#-sub-skills)

---

## 📖 Overview

The **Skill Factory** is the central engine for standardized skill authoring within the Skills Hub. It automates scaffolding, validation, and registration, ensuring every new capability adheres to the project's premium "Gold Standards" and structural integrity.

Instead of manual file creation which leads to inconsistency, the agent utilizes this skill to:
1. **Design**: Perform a conceptual blueprinting of the skill's logic.
2. **Generate**: Scaffold all necessary files from logic-first templates.
3. **Validate**: Audit the skill against SDD v2.2.0 and structural purity checklists.
4. **Register**: Ensure the skill is correctly indexed in the project's global memory and roadmap.

> **Principle**: If SDD ensures code is specified before being written, the Skill Factory ensures every skill is structured before being defined.

---

## ⚙️ Standard Structure (SDD v2.2.0)

| Mode | Use Case | Generated Artifacts |
|------|-------------|-----------------|
| **Quick** | Simple skill without sub-skills | `SKILL.md`, `README.md`, `CHANGELOG.md`, `tasks.md` |
| **Standard** | Complex skill with sub-skills | All above + `<skill>-<sub>.skill.md` (N) + `DECISIONS.md` |

---

## 🧩 Sub-skills

| Sub-skill | File | Responsibility |
|-----------|---------|------------------|
| **Bootstrap** | [skill-factory-bootstrap.skill.md](skill-factory-bootstrap.skill.md) | Generates the complete scaffolding of the new skill using logic-first templates. |
| **Validator** | [skill-factory-validator.skill.md](skill-factory-validator.skill.md) | Audits the created skill against the compliance checklist and emits an audit report. |

---

## 🚀 How to Use

### Example 1: Creating a Simple Skill (Quick Mode)

**Input:**
```
Create a new skill named "deep-research" for advanced web research.
Category: research.
No sub-skills.
```

**The agent executes:**
1. `skill-factory-bootstrap` with `skill_name=deep-research`, `description="Advanced web research with result synthesis"`, `category=research`.
2. `skill-factory-validator` to audit `deep-research/`.
3. Update project roadmap in `.specs/project/ROADMAP.md`.

**Result:**
```
deep-research/
├── SKILL.md          # Definition with placeholders for completion
├── README.md         # Documentation with badges and structure
├── tasks.md          # Initial task list for execution
└── CHANGELOG.md      # Entry [1.0.0] with today's date
```

---

## ✅ Compliance Checklist

Every skill created by the framework is validated against this checklist:

| # | Check | Description |
|---|-------|-----------|
| 1 | **Structural** | All mandatory files exist (Root Integrity) |
| 2 | **Memory** | Operational memory is aggregated in `.specs/project/` (No local triad) |
| 3 | **Frontmatter** | Valid YAML with `name`, `version`, `description`, `category` |
| 4 | **Content** | Mandatory H2 sections present (Goal, Prerequisites, Quality Rules, Prohibited) |
| 5 | **Naming** | Directory, files, and frontmatter follow standard patterns |

---

## 📐 Versioning Policy

- Every new skill starts at version **`1.0.0`**.
- Changelogs follow the [Keep a Changelog](https://keepachangelog.com/en/1.1.0/) format.
- Versioning adheres to [Semantic Versioning](https://semver.org/).

---

## 📝 Changelog

Consult the [CHANGELOG.md](CHANGELOG.md) for the full version history.
