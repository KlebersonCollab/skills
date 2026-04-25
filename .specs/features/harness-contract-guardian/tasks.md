# Tasks: Harness Contract Guardian

## Phase 1: Parser & Infrastructure
- [x] Implement `internal/harness/contract.go` basic structure. <!-- id: 0 -->
- [x] Implement `ParseContractMD` to extract interfaces from tables. <!-- id: 1 -->

## Phase 2: Validation Engine
- [x] Implement file existence check logic. <!-- id: 2 -->
- [x] Implement symbol search logic (Grep-based). <!-- id: 3 -->
- [x] Create `ExecuteContractCheck` orchestrator. <!-- id: 4 -->

## Phase 3: CLI Integration
- [x] Add `contract` subcommand to `hb/cmd/harness.go`. <!-- id: 5 -->

## Phase 4: Verification
- [x] Test on `harness-audit-gate` feature (should pass). <!-- id: 6 -->
- [x] Test on a new feature with a fake contract (should fail). <!-- id: 7 -->
