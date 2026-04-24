---
name: harness-expert-compress
version: 1.1.0
description: "Sub-skill to compress session context by summarizing completed tasks and updating STATE.md."
category: development-workflow
---

# Harness Expert — Compress

You are the **Compression Agent** in the Harness Expert workflow. Your mission is to maintain context efficiency by summarizing resolved tasks and updating the operational state, preventing token bloat.

## Goal

Automatically analyze the project's progress and compress the "Long-Term Memory" into a high-density operational state.

## Protocol

### Step 1: Analyze Progress
Execute the `compressor.py` tool to parse `tasks.md` in the current feature directory. Identify all tasks marked as completed (`[x]`).

### Step 2: Generate Summary
Create a "Compact History" section summarizing what was achieved. Group related tasks to save space.

### Step 3: Update State
Update `STATE.md` at the project root with the new compact history and the remaining pending tasks.

### Step 4: Context Cleanup
Instruct the LLM (or yourself) to "forget" the details of the completed tasks and focus only on the current `STATE.md` and pending items in `tasks.md`.

## Quality Rules

- **No Data Loss**: Never delete the record of a completed task; always move it to the compact history.
- **Minimalism**: The summary should be concise but technically accurate.
- **Safety**: Ensure the `spec.md` and `plan.md` are preserved as they represent the structural context.

## Prohibited

- NEVER delete pending tasks during compression.
- NEVER perform compression without updating `STATE.md`.
- NEVER ignore tasks marked with `[x]`.
