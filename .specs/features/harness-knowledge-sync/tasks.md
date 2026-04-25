# Tasks: Harness Knowledge Sync

## Phase 1: Git Integration
- [x] Implement `getModifiedOwners` in `internal/harness/map.go` to find changed skills/features. <!-- id: 0 -->

## Phase 2: Inference Engine
- [x] Implement logic to detect missing links between features and skills. <!-- id: 1 -->
- [x] Parse `KNOWLEDGE-MAP.mermaid` to check for existing relations. <!-- id: 2 -->

## Phase 3: CLI & Persistence
- [x] Add `map` subcommand to `hb/cmd/harness.go`. <!-- id: 3 -->
- [x] Implement `--apply` logic to update the Mermaid file. <!-- id: 4 -->

## Phase 4: Verification
- [x] Test detection of a new feature link. <!-- id: 5 -->
- [x] Verify that redundant links are not suggested. <!-- id: 6 -->
