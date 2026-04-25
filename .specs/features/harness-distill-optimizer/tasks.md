# Tasks: Harness Distill Optimizer

## Phase 1: Code Distiller
- [x] Update `internal/distiller` to support more patterns (JS/Go/TS). <!-- id: 0 -->
- [x] Implement `ExecuteDistillCode` in `internal/harness/distill.go`. <!-- id: 1 -->

## Phase 2: State Distiller (The "Brain")
- [x] Implement Git metadata gathering (diffs, logs). <!-- id: 2 -->
- [x] Implement `tasks.md` parser for progress metrics. <!-- id: 3 -->
- [x] Implement `ExecuteDistillState` logic to generate/update `STATE.md`. <!-- id: 4 -->

## Phase 3: CLI Integration
- [x] Add `distill` subcommand to `hb/cmd/harness.go`. <!-- id: 5 -->
- [x] Register flags `--state`, `--code`, `--overwrite`. <!-- id: 6 -->

## Phase 4: Verification
- [x] Test code distillation on a sample file. <!-- id: 7 -->
- [x] Test state distillation on the current `harness-distill-optimizer` feature. <!-- id: 8 -->
