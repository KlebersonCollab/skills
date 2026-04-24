# Quick Mode

**Triggers:** "Quick fix", "Quick adjustment", "Trivial bugfix"

For tasks that fall into the **Quick** category (simple bug fixes, configuration, changes in up to 3 files that can be described in one sentence).

## Flow
1. **Skip the Specify, Design, and Tasks phases.** The bureaucracy is unnecessary.
2. **Execute (Implement & Verify):** Go straight to fixing the code.
3. Ensure that tests continue to pass.
4. Confirm that there are no visual or logical regressions.
5. **Commit directly** (or use a short-lived branch without creating an entire feature specification).

## Safety Valve
If, upon starting the task in *Quick Mode*, you realize that the implementation will require structural changes, the addition of new libraries, or substantially modifying more than 3 files, **STOP**.

In this case, declare to the user:
> *"This task seemed quick, but it involves deep architectural changes. Aborting Quick Mode and elevating to Small/Medium scope."*
Then, start the conventional steps of the SDD framework (generate spec, tasks, etc).
