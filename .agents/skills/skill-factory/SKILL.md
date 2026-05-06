---
name: skill-factory
version: 1.0.0
description: Governance and Automation engine for creating and validating Hub skills. Ensures SDD v2.2.0 compliance and "Anti-Slop" design.
category: automation-scaffolding
---

## 🔒 Prerequisites (Mandatory)
This skill acts as the architect for new agent capabilities. Before execution:
0. **Mode Check**: Verify `.hub-mode` and apply `token-distiller` guidelines.
1. **Pattern Check**: Read the latest SDD master spec to align with current "Gold Standards".
2. **Knowledge Check**: Follow the **Knowledge Verification Chain** (see below).

---

# Skill Factory: Capability Architect (v1.0.0)

> Scaling intelligence through rigorous structure and automated governance.

---

## 🧩 Delegation Matrix
The Skill Factory delegates specific lifecycle tasks to its sub-agents:

| Phase | Sub-Skill | Primary Artifact | Purpose |
|---|---|---|---|
| **BOOTSTRAP** | `skill-factory-bootstrap` | New Skill Directory | Scaffolding of folders, base files, and memory artifacts. |
| **VALIDATE** | `skill-factory-validator` | `audit-report.md` | Verification of structural integrity and SDD compliance. |

---

## 🔄 Factory Workflow

### 1. Creation (Bootstrap)
*   **Goal**: Zero-to-hero skill initialization.
*   **Action**: Use `skill-factory-bootstrap` to create the directory structure.
*   **Output**: A fully populated directory with `SKILL.md`, `README.md`, and the `.specs/` hierarchy.

### 2. Validation (Audit)
*   **Goal**: Ensure the new skill isn't "Slop".
*   **Action**: Use `skill-factory-validator` to scan the new skill.
*   **Output**: An audit score and required fixes.

---

## 🛠️ Operational Protocols

### 1. The "Gold Standard" Mandate
Every skill created MUST include:
- **Semantic Frontmatter**: Accurate metadata for the Skill Router.
- **Memory Triad**: `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`.
- **SDD Hierarchy**: Feature specs must live in `.specs/features/`.

### 2. Knowledge Verification Chain
To prevent pattern drift, the Factory follows:
1. **Core SDD Skill**: The ultimate source of truth for workflow patterns.
2. **Existing Skills**: High-performance skills (like `sdd` or `git-workflow`) as references.
3. **Internal Governance**: `.specs/codebase/CONVENTIONS.md`.

---

## 🏗️ Production Patterns

### Atomic Scaffolding
Never create a skill folder without its mandatory files. If the process is interrupted, the validator must flag the skill as `INCOMPLETE`.

### Language Policy
- **Logic & Instructions**: English (Universal agent language).
- **Communication**: Respect USER preferences (e.g., Brazilian Portuguese) for reporting results.

---

## 🚫 Prohibited

- NEVER create empty `SKILL.md` files.
- NEVER skip the memory artifact initialization.
- NEVER approve a skill that doesn't follow the 4-Phase SDD workflow.

---

## 6. Directory Structure for New Skills (Mandatory)

```text
[skill-name]/
├── SKILL.md
├── README.md
├── MEMORY.md
├── LEARNINGS.md
├── STATE.md
├── DECISIONS.md
├── tasks.md
├── examples/
├── references/
├── resources/
└── .specs/
    ├── features/
    ├── project/
    └── codebase/
```

---

> **Law of the Factory**: A skill without structure is just a prompt. Structure is the bridge to autonomy.
