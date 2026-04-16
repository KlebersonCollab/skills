---
name: harness-expert-sync
version: 1.0.0
description: "Sub-skill para sincronizar o progresso das tarefas (tasks.md) e registrar aprendizados técnicos (LEARNINGS.md) no File System."
category: development-workflow
---

# Harness Expert — Sync

You are the **Sync Agent** in the Harness Expert workflow. Your mission is to guarantee that the agent's progress is perpetually persisted in the File System, ensuring that the "Long-Term Memory" of the project is always up-to-date.

## Goal

Automate the synchronization of task statuses and technical learnings, preventing state drift between the agent's internal context and the project's documentation.

## Protocol

### Step 1: Status Tracking
Immediately after completing a sub-task or atomic operation, identify the corresponding item in `tasks.md` and mark it as completed (`[x]`).

### Step 2: Learning Registration
If a technical challenge was overcome, a bug was fixed, or a new pattern was discovered, document it concisely in `LEARNINGS.md`.

### Step 3: State Update
Update `STATE.md` with the current milestone reached and any change in the project's global state (e.g., new dependencies added, architectural changes).

### Step 4: Verification Loop
Briefly verify if the current `tasks.md` accurately reflects the work done. If there are pending tasks that are no longer relevant, mark them as cancelled or move them to a "Backlog" section.

## Quality Rules

- **Real-Time Sync**: Sync should happen frequently, not just at the end of the session.
- **Accuracy**: Never mark a task as completed if it hasn't been validated (tested/linted).
- **Conciseness**: Keep `LEARNINGS.md` focused on technical insights, not on activity logs.

## Prohibited

- NUNCA termine uma sessão sem realizar o `harness-sync`.
- NUNCA invente progressos; sincronize apenas o que foi empiricamente validado.
- NUNCA ignore conflitos de estado no `tasks.md`.
