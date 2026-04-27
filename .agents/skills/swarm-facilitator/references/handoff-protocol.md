# Inter-Agent Handoff Protocol

The `handoff.md` file is the transition contract between personas (e.g., Architect to Engineer). It must ensure that the next agent has full context without redundancy.

## Mandatory Structure

### 1. State Summary
- **What was done**: Succinct list of completed tasks.
- **Blockers**: Impediments encountered and how they were resolved or bypassed.

### 2. Generated Artifacts
- Links to new files, specs (PRD/RFC), ADRs, and Mermaid diagrams.

### 3. Mission for the Next Persona
- **Immediate Objective**: What should be done next.
- **Acceptance Criteria**: How the next agent will know they have completed their part.

### 4. Technical Context (Deep Dive)
- Necessary environment variables.
- Critical design decisions not in ADRs ("tactical" decisions).

## Example Template

```markdown
# Handoff: [Source Persona] -> [Destination Persona]

## 📝 Current Status
- [X] Spec finalized in `.specs/features/auth-v2/prd.md`
- [X] ADR-005 approved for JWT usage.

## 🔗 Artifacts
- Spec: `prd.md`
- Mockup: `ui-draft.mermaid`

## 🎯 Next Steps (For [Destination Persona])
1. Implement the user repository following the `IUserRepository` interface.
2. Validate against tests in `tests/auth_test.py`.

## ⚠️ Notes
- The dev database should be running on port 5432.
```
