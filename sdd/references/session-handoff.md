# Session Handoff

**Triggers:** "Pause work", "End session", "Continue work", "Resume"

## Pausing Work
Before ending a session's context, ensure that all current state is accurately saved.
1. **Update STATE.md:** Note the current active task, the last blockers faced, and the immediate next steps.
2. **Atomic Commit:** Ensure that ongoing code is committed to the feature branch (use `[WIP]` if necessary).
3. **Task Check:** Mark in `tasks.md` exactly where the execution stopped.

## Resuming Work
When starting or resuming a session:
1. Rehydrate the context by first reading `STATE.md`.
2. Identify the `Active Task`.
3. Validate the current state of the feature branch in Git.
4. Review any files in `LEARNINGS.md` or recently established new rules.

## Continuity Rules
The main premise of the Handoff is that the next session or the next agent should have *exactly* the same understanding and direction of how the task was left. Short-term memory is lost, so the `STATE.md` file becomes the persistent brain of the operation.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "SDD-CORE"
phase: "MANAGEMENT"
status: "COMPLETED"
last_update: "2026-05-06T09:42:30Z"
evidence_checksum: "NONE"
```
