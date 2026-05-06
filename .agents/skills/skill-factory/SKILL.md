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
| **DESIGN** | `skill-factory-bootstrap` | Skill Blueprint | Conceptual and physical mapping of the skill's logic. |
| **AUDIT** | `skill-factory-validator` | `audit-report.md` | Verification of structural and logical integrity. |

---

## 🔄 Factory Workflow

### 1. Design (Bootstrap)
*   **Goal**: Define the skill's ontological structure.
*   **Action**: Use `skill-factory-bootstrap` to design the directory and file logic.
*   **Output**: A clean, structured skill repository following the SDD v2.2.0 spec.

### 2. Validation (Audit)
*   **Goal**: Ensure the new skill isn't "Slop".
*   **Action**: Use `skill-factory-validator` to scan the new skill.
*   **Output**: An audit score and required fixes.

---

## 🛠️ Operational Protocols

### 1. The "Structural Purity" Mandate
Every skill created MUST be self-documenting:
- **Explicit Logic**: No hidden dependencies or "magic" CLI commands.
- **Memory Triad**: Mandatory `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`.
- **Atomic Intent**: Every task and spec must have a clear ID and verification method.

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
