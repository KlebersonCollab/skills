---
name: token-distiller
version: 2.3.0
description: "Token density manager. Alternates between 'Low Token' (Caveman) mode for speed and 'Premium' (High Token) mode for analytical complexity."
category: utility
---

## 🔒 Prerequisites (Mandatory)
This skill operates integrated with the **SDD** framework. Before any technical execution:
0. **Mode Check**: Verify the current operational mode in `.hub-mode`.
1. **Context Check**: Rehydrate state by reading `.specs/project/STATE.md`, `.specs/project/MEMORY.md`, and `.specs/project/LEARNINGS.md`.
2. **Spec Check**: Does the `spec.md` file exist with clear requirements?
3. **Plan Check**: Does the `plan.md` file define the compression architecture?

---

# Token Distiller: Dual-Mode Communication (v2.3.0)

> "Density is the currency of autonomy. Spend wisely."

---

## 🧩 Delegation Matrix
As a utility skill, Token Distiller operates as a unified agent across all phases:

| Phase | Sub-Skill | Primary Artifact | Purpose |
|---|---|---|---|
| **AUDIT** | Self | Context Report | Assess token pressure and context health. |
| **DENSITY** | Self | Distilled Output | Execute linguistic compression or prose expansion. |

---

## 🔄 4-Phase Workflow

### 1. DISCOVERY
*   **Goal**: Audit current context health and token pressure.
*   **Action**: Assess if context exceeds 80% capacity. If so, perform manual context pruning by summarizing previous tool outputs and removing redundant thought paths. Identify task complexity (Quick/Small vs Medium+).
*   **Output**: Context health report and mode selection.

### 2. SPECIFY
*   **Goal**: Define the required fidelity for the next interaction.
*   **Action**: Set `/mode low` (Caveman) for speed or `/mode high` (Premium) for analysis.
*   **Output**: Mode activation in session state.

### 3. IMPLEMENT
*   **Goal**: Execute communication according to the selected mode.
*   **Action**: Apply linguistic filters (Caveman fragments or Premium prose). Perform micro-compaction of tool results.
*   **Output**: Compressed or high-fidelity responses.

### 4. VERIFY
*   **Goal**: Ensure substance preservation and safety.
*   **Action**: Validate that task IDs and technical patterns are intact. Exit Low Token mode for safety warnings.
*   **Output**: Verified, dense communication.

---

## 🏗️ Operating Modes

### 🪨 Low Token Mode (Caveman)
**Activation:** 'Quick', 'Small' tasks, `/mode low`, `/caveman`.
**Rules:**
- Eliminate articles (the, a, an), fillers (really, basically), and greetings.
- Use fragments: `[object] [action] [reason]. [next step].`
- **Example:** "Bug in middleware. Expiry check uses `<` not `<=`. Fix applied."

### 💎 Premium Mode (High Token)
**Activation:** 'Medium', 'Large', 'Complex' tasks, `/mode high`, `/premium`.
**Rules:**
- Analytical, grammatically complete, and professional prose.
- Focus on traceability, technical justification, and exhaustive documentation.

---

## 🛠️ Operational Protocols

### 1. Knowledge Verification Chain
1. **Global Mandates**: `.specs/codebase/GLOBAL_MANDATES.md` for engineering standards.
2. **Core SDD Skill**: Reference for task sizing and workflow.
3. **Project Specs**: `.specs/project/STATE.md` for current context.
4. **Internal Governance**: `skill-factory` standards for documentation.

### 2. Micro-Compaction (Stubbing)
- **Trigger**: Tool results > 10 turns old AND > 5000 characters.
- **Action**: Replace with: `[Old tool result cleared. Tool: {name}. Summary: {summary}].`

### 3. Safety Valve
- **Prohibited**: NEVER compress source code, safety warnings, or confirmations of irreversible actions.
- **Escalation**: Abandon Low Token mode if ambiguity causes execution errors.

---

## 🚫 Prohibited

- NEVER create local `STATE.md`, `MEMORY.md`, or `LEARNINGS.md` within the skill folder.
- NEVER use placeholders like "todo" or "..." in instructions.
- NEVER compress technical error messages.

---

> **Law of Density**: All substance, no fluff. Efficiency is precision.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "TOKEN-DISTILLER-V2.3.0"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:30:00Z"
evidence_checksum: "NONE"
```

