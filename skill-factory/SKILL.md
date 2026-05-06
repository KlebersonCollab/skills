---
name: skill-factory
version: 2.3.0
description: "Governance and Automation engine for high-performance Skill Authoring. Implements Anthropics's 14 patterns synthesized with SDD v2.2.0 operational rigor."
category: skill-management
---

## 🔒 Prerequisites (Mandatory)
This skill acts as the architect for the Hub's capabilities. Before execution:
0. **Mode Check**: Verify `.hub-mode` and apply `token-distiller` guidelines.
1. **Context Check**: Rehydrate state by reading `.specs/project/STATE.md`, `.specs/project/MEMORY.md`, and `.specs/project/LEARNINGS.md`.
2. **Knowledge Check**: Follow the **Knowledge Verification Chain** (see below).

---

# Skill Factory: Capability Architect (v2.2.0)

> "Structure is the bridge to autonomy. Precision in design, rigor in execution."

---

## 🧩 Delegation Matrix
The Skill Factory delegates specific lifecycle tasks to its sub-agents:

| Phase | Sub-Skill | Primary Artifact | Purpose |
|---|---|---|---|
| **DESIGN** | `skill-factory-bootstrap` | Skill Blueprint | Conceptual and physical mapping of the skill's logic. |
| **AUDIT** | `skill-factory-validator` | `audit-report.md` | Verification of structural and logical integrity. |

---

## 🏗️ Skill Authoring Patterns (The Core)

### Group 1: Discovery & Selection
1.  **Activation Metadata**: Precise descriptions that trigger the agent correctly.
2.  **Exclusion Clause**: Explicit "Do NOT use for..." lines to prevent hijacking.

### Group 2: Context Economy
3.  **Context Budget**: Every sentence must justify its token cost.
4.  **Progressive Disclosure**: Move details to sub-files (`REFERENCE.md`, `FORMS.md`).

### Group 3: Instruction Calibration
5.  **Control Tuning**: Match instruction freedom to task fragility.
6.  **Explain-the-Why**: State rules + rationale for better generalization.
7.  **Template Scaffold**: Provide explicit placeholders for reports/logs.
8.  **In-Skill Examples**: Embed 2-3 Few-shot pairs for style calibration.
9.  **Known Gotchas**: Document failure modes to prevent hallucinations.

### Group 4: Workflow Control
10. **Execution Checklist**: Copyable checklists to track progress.
11. **Self-Correcting Loop**: "Produce-validate-fix" loops for quality.
12. **Plan-Validate-Execute**: Verified plan before performing side effects.

### Group 5: Executable Logic
13. **Utility Bundle**: Purpose-built scripts for deterministic logic.
14. **Autonomy Calibration**: Declare `allowed-tools` to minimize friction.

---

## 🔄 4-Phase Workflow

### 1. DISCOVERY
*   **Goal**: Understand the functional domain and intent.
*   **Action**: Use `skill-factory-bootstrap` to perform an **Intent Audit**.
*   **Output**: Defined triggers and core task scope.

### 2. SPECIFY
*   **Goal**: Design the skill's soul and instructions.
*   **Action**: Apply **Control Tuning** and **Explain-the-Why**. Define the sub-file structure.
*   **Output**: Initial draft of `SKILL.md` and required sub-files.

### 3. IMPLEMENT
*   **Goal**: Calibrate and finalize instructions.
*   **Action**: Add **Few-Shot Examples** and **Known Gotchas**.
*   **Output**: Verified skill directory.

### 4. VERIFY
*   **Goal**: Audit against SDD and Slop-Free standards.
*   **Action**: Use `skill-factory-validator` to produce an audit score.
*   **Output**: Final `audit-report.md` and registration.

---

## 🛠️ Operational Protocols

### 1. The "Structural Purity" Mandate
Every skill created MUST be self-documenting:
- **Explicit Logic**: No hidden dependencies or "magic" CLI commands.
- **Centralized Memory**: Skills do NOT store their own `STATE.md`, `MEMORY.md`, or `LEARNINGS.md`. Operational memory MUST be aggregated in `.specs/project/`.
- **Atomic Intent**: Every task and spec must have a clear ID.

### 2. Knowledge Verification Chain
1. **Core SDD Skill**: Ultimate truth for workflow patterns.
2. **Existing Skills**: High-performance references (e.g., `sdd`, `git-workflow`).
3. **Internal Governance**: `.specs/codebase/CONVENTIONS.md`.

### 3. Safety Valve
If a skill design exceeds 5 sub-skills or complex external dependencies, **STOP** and re-evaluate for orchestration.

### 4. Observable Governance (v2.3.0)
Every skill created MUST support the automated auditing protocol.
- **Metadata**: Append `<!-- @sdd-state -->` to all `.md` files.
- **Evidence**: All tasks MUST be tracked via an evidence-based table in `tasks.md`.

---


## 🚫 Prohibited

- NEVER modify system directories (e.g., `.agents/`, `.gemini/`) unless explicitly authorized.
- NEVER create local `STATE.md`, `MEMORY.md`, or `LEARNINGS.md` within a skill; use `.specs/project/` instead.
- NEVER approve a skill that doesn't follow the 4-Phase SDD workflow.
- NEVER use placeholders like "todo" or "..." in instructions.

---

> **Law of the Factory**: A skill without structure is just a prompt. Structure is the bridge to autonomy.

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
