# Tasks: HB-CLI Premium UI

## Phase 1: Infrastructure (ANSI & Engine)
- [x] T1: Setup `hb/internal/ui` package and basic ANSI constants. `id: UI-001`
- [x] T2: Implement `ANSIHelper` with 100% test coverage (cursor management). `id: UI-002`
- [x] T3: Implement `TerminalSize` detection for Go (cross-platform). `id: UI-003`

## Phase 2: Core UI Logic
- [x] T4: Create `UIState` struct and `UIManager` initialization. `id: UI-004`
- [x] T5: Implement the rendering loop (Ticker) with "zero-flicker" logic. `id: UI-005`
- [x] T6: Implement the `StatusBar` component (Feature, Progress, Tokens). `id: UI-006`

## Phase 3: Integration & Testing
- [x] T7: Integrate `UIManager` into the `hb` root command (via a global flag `--ui`). `id: UI-007`
- [x] T8: Connect `internal/sdd` progress to the `UIState`. `id: UI-008`
- [x] T9: Final 100% coverage audit and validation report. `id: UI-009`

## Phase 4: Polish
- [x] T10: Add colors (using `chalk`-like library in Go) and icons. `id: UI-010`
