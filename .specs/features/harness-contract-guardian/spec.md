# Specification: Harness Contract Guardian

## Goal
Implement an architectural drift detection tool in the `hb` CLI that ensures the implementation strictly follows the established `contract.md`.

## Context
In the SDD framework, the `contract.md` is the source of truth for interfaces. Currently, verifying if the implementation matches the contract is a manual task. This tool automates this verification as a "Harness Gate".

## Acceptance Criteria (AC)

### AC 1: Command `hb harness contract`
- [ ] Implement the command as a subcommand of `harness`.
- [ ] Must accept a feature name as an argument.

### AC 2: Contract Parsing
- [ ] Must parse `contract.md` to identify:
    - **Artifacts**: Files that must exist.
    - **Commands/Interfaces**: CLI commands or API endpoints defined in tables.

### AC 3: Verification Logic
- [ ] **File Check**: Verify if specified files exist in the project.
- [ ] **Interface Check**: 
    - For Go: Search for command declarations or function names.
    - For Python/JS: Search for route definitions or class methods.
- [ ] **Version Check**: (Optional) Verify if version in contract matches implementation.

### AC 4: Reporting
- [ ] Display a "Contract Compliance Report".
- [ ] Exit code `0` if 100% compliant.
- [ ] Exit code `1` if drift is detected.
- [ ] Update `validation-report.md` with the "Contract Status".
