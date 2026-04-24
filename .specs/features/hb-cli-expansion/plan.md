# Execution Plan: HB-CLI Expansion Pack

## 1. Architecture
We will use a modular approach, creating internal packages for each domain logic.

### 1.1. Internal Directory Structure
- `internal/adr/`: Logic for ADR generation and listing.
- `internal/audit/`: Integration with linters and compliance checks.
- `internal/knowledge/`: Logic for `hb learn` and learnings persistence.
- `internal/scaffold/`: Logic for skill scaffolding.
- `internal/distiller/`: Logic for token optimization and session distillation.
- `internal/doctor/`: Diagnostic and repair logic.

## 2. Command Hierarchy (Cobra)
- `hb adr [new|list]`
- `hb audit`
- `hb learn`
- `hb skill [new|validate]`
- `hb distill [file|session]`
- `hb doctor`
- `hb sdd review`

## 3. Implementation Phases

### Phase 1: ADR & Learning (Knowledge Domain)
- [ ] Implement `hb adr new` and `hb adr list`.
- [ ] Implement `hb learn`.
- [ ] **Validation**: Create unit tests for template generation and file parsing.

### Phase 2: Quality & Scaffolding (Engineering Domain)
- [ ] Implement `hb audit` (Python/Go/Hub).
- [ ] Implement `hb skill new`.
- [ ] **Validation**: Verify scaffolding output and audit scores.

### Phase 3: Optimization & Health (Maintenance Domain)
- [ ] Implement `hb distill`.
- [ ] Implement `hb doctor`.
- [ ] **Validation**: Test distillation compression ratio and doctor fix capability.

### Phase 4: SDD Integration
- [ ] Implement `hb sdd review`.
- [ ] Final Hub-wide integration and version bump.

## 4. TDD Strategy
- Each internal package must have a `_test.go` file.
- Use `testify` or standard `testing` package.
- Mocking: Use interfaces for filesystem interaction to facilitate testing.
