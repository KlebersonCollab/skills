# UltraPlan: Remote Exploration Protocol

UltraPlan is used for tasks with high architectural ambiguity that would otherwise saturate the local context.

## 1. Trigger Criteria
Invoke UltraPlan when:
- The task requires understanding > 5 unfamiliar libraries/frameworks.
- The project documentation is sparse or conflicting.
- The initial exploration turn results in > 20 files that need cross-referencing.

## 2. Remote Session Setup
1.  **Draft a Research Query**: Compile all knowns and unknowns into a structured document.
2.  **User Hand-off**: Ask the user to run a "Deep Research" session in a specialized AI (like GPT-4o with Search or Claude with Project capabilities) using the query.
3.  **Synthesis**: The user returns with a consolidated Technical RFC.

## 3. Local Hydration
1.  **Import RFC**: Save the remote findings as `.specs/features/<feature>/ultraplan-results.md`.
2.  **Task Refinement**: Update `tasks.md` based on the new insights.
3.  **Resume Phase 2**: Complete the `spec.md` and `plan.md` locally using the hydrated knowledge.
