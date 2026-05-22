---
name: token-distiller
version: 2.4.0
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

# Token Distiller: Dual-Mode Communication (v2.4.0)

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
**Activation:** 'Quick', 'Small' tasks, `/mode low [lite|full|ultra]`, `/caveman [lite|full|ultra]`.
**Persistence:** Active every response. Reverts only on "stop caveman" or "normal mode".
**Strict Response Pattern:** `[thing] [action] [reason]. [next step].`

#### 📊 Intensity Levels:
- **lite**:
  - No filler or hedging.
  - Keep articles (a/an/the) and complete grammatical sentences.
  - Professional but tight, crisp communication.
- **full** (default):
  - Drop articles (a/an/the) and filler words (really, basically, just, actually).
  - Sentence fragments are welcome. Use short, direct synonyms (e.g., "big" instead of "extensive", "fix" instead of "implement a solution for").
  - Technical terms, code blocks, errors must remain completely unchanged.
- **ultra**:
  - Max compression: abbreviate prose words (e.g., DB, auth, config, req, res, fn, impl).
  - Strip all conjunctions.
  - Use arrows for causality: `[cause] → [effect]`.
  - **Never** abbreviate code symbols, function names, API names, error strings.

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

### 3. Safety Valve & Auto-Clarity
- **Auto-Clarity Trigger (Skip Caveman)**: Automatically suspend Caveman mode and write full prose when:
  1. **Security Warnings**: Explaining security vulnerabilities or configurations.
  2. **Irreversible Action Confirmations**: Prior to destructive actions (e.g., `DROP TABLE`, deletion, overwriting).
  3. **Multi-step Sequences**: Where fragment order or missing conjunctions risk misinterpretation.
  4. **Linguistic Ambiguity**: If compressing information creates confusion (e.g., "migrate table drop column backup first" is ambiguous about execution order).
  5. **Direct Clarification**: User asks to clarify or repeats a question.
- **Resumption**: Resume selected Caveman intensity level immediately after the clear safety segment is completed.

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
feature_id: "TOKEN-DISTILLER-ADVANCED-CAVEMAN"
phase: "IMPLEMENT"
status: "IN_PROGRESS"
last_update: "2026-05-22T14:02:46Z"
evidence_checksum: "NONE"
```

