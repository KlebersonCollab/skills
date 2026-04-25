# Specification: SDD Auditor

## Goal
Implement a compliance checker for the Spec-Driven Development (SDD) framework within the `hb` CLI to ensure all features are properly documented and tracked.

## Context
As the number of features grows, consistency in documentation (spec, plan, contract, tasks) becomes harder to maintain manually. This tool provides a "Health Dashboard" for the project's specifications.

## Acceptance Criteria (AC)

### AC 1: Command `hb sdd audit`
- [ ] Implement the command as a subcommand of `sdd`.
- [ ] Must output a summary table in the terminal.

### AC 2: Compliance Verification
- [ ] Check for mandatory files in each feature folder:
    - `spec.md` (or legacy `PRD.md`)
    - `plan.md`
    - `contract.md`
    - `tasks.md`
- [ ] Identify legacy filenames and suggest renaming.

### AC 3: Progress Calculation
- [ ] Parse `tasks.md` files.
- [ ] Calculate percentage: `(completed_tasks / total_tasks) * 100`.
- [ ] Handle files without any tasks markers gracefully.

### AC 4: Exit Codes
- [ ] Exit `0` if all features are compliant (all mandatory files present).
- [ ] Exit `1` if any feature is missing mandatory files.
