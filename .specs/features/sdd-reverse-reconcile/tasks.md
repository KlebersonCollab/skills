# Tasks: SDD Reverse Reconciler

## Phase 1: Evidence Logic
- [x] Implement `findEvidence(root, keyword)` helper. <!-- id: 0 -->
- [x] Implement `ReconcileTasks(featureDir)` logic. <!-- id: 1 -->

## Phase 2: Contract Generation
- [x] Implement `GenerateMissingContract(featureDir)` based on found symbols. <!-- id: 2 -->

## Phase 3: CLI Integration
- [x] Add `reconcile` subcommand to `hb/cmd/sdd.go`. <!-- id: 3 -->

## Phase 4: Verification
- [x] Test reconciliation on a feature with 0% progress but working code. <!-- id: 4 -->
- [x] Verify that `contract.md` is generated correctly. <!-- id: 5 -->
