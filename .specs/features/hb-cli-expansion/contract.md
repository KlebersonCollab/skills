# Validation Contract: HB-CLI Expansion Pack

## 1. Validation Sensors (Automated Tests)

| Sensor ID | Target Command | Validation Logic |
|-----------|----------------|------------------|
| **V-ADR-01** | `hb adr new` | Check if file `ADR-0001-*.md` is created in `.specs/architecture/` with "Proposed" status. |
| **V-ADR-02** | `hb adr list` | Check if the output contains a table with the newly created ADR. |
| **V-LRN-01** | `hb learn` | Verify if a new `## [Date] Subject` entry appears in `.specs/project/LEARNINGS.md`. |
| **V-AUD-01** | `hb audit` | Verify if it detects a missing `CHANGELOG.md` in a dummy skill folder and reduces the score. |
| **V-SKL-01** | `hb skill new` | Check if `SKILL.md`, `CHANGELOG.md`, and `README.md` are created with correct templates. |
| **V-DST-01** | `hb distill` | Compare input vs output size; output must be smaller and functional. |
| **V-DOC-01** | `hb doctor` | Verify if it detects and fixes a missing `.cursorrules` (re-generating it). |

## 2. Quality Gates (Golang Expert)
- **Unit Testing**: All `internal/` packages must have `_test.go` with at least 80% coverage.
- **Linting**: Must pass `go vet` and `staticcheck` without warnings.
- **Build**: Binary must compile for Darwin/AMD64 and Darwin/ARM64.

## 3. Integration Sensors
- **Agent Awareness**: After `hb init`, the `.agent/AGENT.md` must contain the updated mandates including the new CLI tools.
- **Cross-Skill Sync**: Verify if `architecture` skill can successfully call `hb adr list`.
