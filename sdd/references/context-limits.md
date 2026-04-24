# Context Limits and Artifact Size

Context management is vital to maintaining logical reasoning capacity in long tasks. Strictly follow these limits.

## File Limits

| File            | Max Tokens | ~Words    | Warning at         |
| --------------- | ---------- | --------- | ------------------ |
| PROJECT.md      | 2,000      | 1,200     | 1,600 (80%)        |
| ROADMAP.md      | 3.000      | 1,800     | 2,400              |
| STATE.md        | 10.000     | 6.000     | 7,000 (70%)        |
| spec.md         | 5.000      | 3.000     | 4,000              |
| design.md       | 8.000      | 4,800     | 6,400              |
| tasks.md        | 10.000     | 6.000     | 8,000              |
| STACK.md        | 2.000      | 1,200     | 1,600              |
| ARCHITECTURE.md | 4.000      | 2.400     | 3.200              |
| CONVENTIONS.md  | 3.000      | 1.800     | 2,400              |
| STRUCTURE.md    | 2.000      | 1,200     | 1,600              |
| TESTING.md      | 4.000      | 2.400     | 3.200              |
| INTEGRATIONS.md | 5.000      | 3.000     | 4,000              |

## Context Zones
- 🟢 **Healthy** (<40k total tokens): Silent.
- 🟡 **Moderate** (40-60k): Discreet footer warning.
- 🔴 **Critical** (>60k): Active warning, suggests context optimization (e.g., archiving logs in `STATE.md`).

## Operational Principles
1. **Target:** Always maintain less than 40k active tokens loaded in context memory.
2. **Operational Reserve:** Ensure the model always has more than 160k available free tokens for reasoning and output.
3. **On-Demand Loading:** Avoid loading multiple feature specs at the same time. Load artifacts only when strictly necessary for the current task.
