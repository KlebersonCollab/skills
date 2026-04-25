# Tasks: Harness Audit Gate

## Phase 1: Infrastructure
- [x] Create `internal/security/scanner.go` with secret detection logic. <!-- id: 0 -->
- [x] Create `internal/harness/audit.go` to orchestrate audits. <!-- id: 1 -->

## Phase 2: CLI Implementation
- [x] Add `audit` subcommand to `hb/cmd/harness.go`. <!-- id: 2 -->
- [x] Implement flag parsing (`--sec`, `--mentor`, `--all`). <!-- id: 3 -->

## Phase 3: Reporting
- [x] Implement markdown generator for `validation-report.md`. <!-- id: 4 -->
- [x] Integrate report generation into the audit command. <!-- id: 5 -->

## Phase 4: Integration & UX
- [x] Update `hb sync` to verify the audit status before proceeding. <!-- id: 6 -->
- [x] Add colorized output for terminal visibility. <!-- id: 7 -->

## Phase 5: Verification
- [x] Run `hb harness audit` on a test file with a fake API key to verify detection. <!-- id: 8 -->
- [x] Run `hb audit` (full suite) to ensure project integrity. <!-- id: 9 -->
