# Specification: Structured Review (hb review)

## Goal
Implement a semantic audit command `hb review` that validates staged changes against the current SDD task context, ensuring atomic commits and preventing "code pollution".

## Acceptance Criteria (BDD)
### Scenario 1: Semantic Validation of Staged Changes
- **Given** I have staged changes in Git
- **And** the current feature has a `tasks.md` with an active task
- **When** I run `hb review`
- **Then** the command should parse the diff and compare it with the task description
- **And** it should flag any changes in files not related to the feature
- **And** it should detect "debug trash" (e.g., console.log, fmt.Printf)

### Scenario 2: Structured Diff Display
- **Given** I run `hb review`
- **When** the diff is analyzed
- **Then** it should display a clean, structured diff in the terminal
- **And** it should provide a "PASS/FAIL" verdict for the commit

## Constraints
- Must use `internal/ui` for premium output.
- Must be cross-platform (mac/linux).
- Must not block the commit if the user explicitly ignores (warning mode by default).
