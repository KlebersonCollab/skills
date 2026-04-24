# New Feature Checklist

Mandatory steps for implementing a new feature in the ecosystem.

## 1. Preparation
- [ ] Feature branch created.
- [ ] `STATE.md` updated with "Starting feature [name]".
- [ ] Task ID assigned in `tasks.md`.

## 2. Specification & Planning
- [ ] `spec.md` created with clear requirements.
- [ ] BDD scenarios defined in Acceptance Criteria.
- [ ] `plan.md` created with architecture and schemas.
- [ ] Mermaid diagrams included (Flowchart/Sequence).
- [ ] `contract.md` established with validation sensors.

## 3. Implementation
- [ ] TDD cycle followed (Red-Green-Refactor).
- [ ] Logic implemented in atomic commits per task.
- [ ] English-only compliance for comments and documentation.
- [ ] Linting and type checking passing locally.

## 4. Verification
- [ ] All tests passing 100%.
- [ ] `validation-report.md` generated with evidence.
- [ ] Score review completed.

## 5. Finalization
- [ ] Pull Request opened with link in `validation-report.md`.
- [ ] `tasks.md` updated as 100% completed.
- [ ] `STATE.md` and `MEMORY.md` synchronized.
- [ ] Learnings captured in `LEARNINGS.md`.
