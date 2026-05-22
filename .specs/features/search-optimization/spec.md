
# Spec: SearchFiles Performance Optimization (Go Native Parallel)

## Motivation
Current `SearchFiles` uses sequential `filepath.WalkDir`, which is slow on large workspaces. The goal is to achieve ~100x speedup using Go native concurrency, without external dependencies (agnostic to OS).

## Acceptance Criteria (ACs)

- [x] **AC1**: `SearchFiles` spawns a configurable number of workers (default: `runtime.NumCPU()`) to read and match files concurrently.
- [x] **AC2**: A dispatcher goroutine walks the directory tree and sends file paths to a buffered channel.
- [x] **AC3**: Binary detection is performed on the first 512 bytes of each file; binary files are skipped before full read.
- [x] **AC4**: A `[]byte` pool (`sync.Pool`) is used to reduce allocations on small files (< 64KB).
- [x] **AC5**: Results are aggregated and capped at 100 matches (existing behavior preserved).
- [x] **AC6**: All existing tests (`TestSearchFiles`, `TestRegexSearchFiles`) continue to pass.
- [x] **AC7**: No external binaries (`rg`, `grep`) are required; the function works identically on Linux, macOS, Windows.
- [x] **AC8**: A benchmark test demonstrates at least 10x speedup over sequential walk on a workspace with ≥1000 files.</target>

## Non-Goals
- Do not change the function signature or behavior for consumers.
- Do not add new CLI commands or flags.
- Do not modify any other file beyond `tools.go` and the test file `search_test.go` (or `benchmark_test.go`).

## Glossary
- **Dispatcher**: goroutine that walks the FS tree and sends paths to `jobsChan`.
- **Worker**: goroutine that reads a file from `jobsChan`, performs matching, and sends results to `resultsChan`.
- **Buffer pool**: `sync.Pool` holding `[]byte` slices to avoid repeated allocations.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "search-optimization"
phase: "SPEC"
status: "IN_PROGRESS"
last_update: "2026-05-23T16:00:00Z"
```
