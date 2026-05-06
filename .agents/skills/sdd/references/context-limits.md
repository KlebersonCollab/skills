# Context Limits and Artifact Size

Context management is vital to maintaining logical reasoning capacity in long tasks. Strictly follow these limits.

## File Limits

| File               | Max Tokens | ~Words    | Warning at         |
| ------------------ | ---------- | --------- | ------------------ |
| PROJECT.md         | 2,000      | 1,200     | 1,600 (80%)        |
| ROADMAP.md         | 3,000      | 1,800     | 2,400              |
| STATE.md           | 10,000     | 6,000     | 7,000 (70%)        |
| spec.md            | 5,000      | 3,000     | 4,000              |
| plan.md            | 8,000      | 4,800     | 6,400              |
| tasks.md           | 10,000     | 6,000     | 8,000              |
| TECHNICAL-MAP.md   | 15,000     | 9,000     | 12,000             |
| MEMORY.md          | 5,000      | 3,000     | 4,000              |

## Context Zones
- 🟢 **Healthy** (<40k total tokens): Silent.
- 🟡 **Moderate** (40-60k): Discreet footer warning.
- 🔴 **Critical** (>60k): Active warning, suggests context optimization (e.g., using `token-distiller` or archiving logs).

## Operational Principles
1. **Target:** Maintain less than 40k active tokens loaded.
2. **On-Demand Loading:** Never load all feature specs simultaneously.
3. **Hierarchy of Importance:** `STATE.md` and `TECHNICAL-MAP.md` are the primary anchors for consistency.
