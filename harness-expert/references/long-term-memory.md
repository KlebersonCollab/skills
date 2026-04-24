# Long-Term Memory (LTM) in Harness

The `harness-expert` is responsible for ensuring the agent does not suffer from "amnesia" between sessions, keeping the operational state synchronized.

## 🧠 Memory Concepts

### 1. Operational State (`STATE.md`)
It is the current "pointer". It contains task progress, what is pending, and the focus of the current session. It is structured short-term memory.

### 2. Persistent Memory (`MEMORY.md`)
Stores facts, decisions made, and project context that does not change frequently. Ex: "The project uses Postgres 15", "The log pattern is JSON".

### 3. Lessons Learned (`LEARNINGS.md`)
Records mistakes made, corrections, and "eureka moments". It serves to avoid repeating past failures and improve agent efficiency.

## 🔄 Synchronization Flow

1. **Rehydration**: At the start of the session, the agent reads these files via `harness-expert-rehydrate`.
2. **Update**: During execution, the agent updates `STATE.md` as tasks are completed.
3. **Final Persistence**: At the end, `harness-expert-sync` ensures that changes are saved to disk.

## 💡 Why does this matter?
Without LTM, the agent would lose progress with each new interaction, forcing the user to repeat instructions and increasing token cost and time.
