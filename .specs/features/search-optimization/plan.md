
# Plan: SearchFiles Performance Optimization

## Task List

### 1. Refactor `SearchFiles` in `tools.go`
- [x] **1.1** Inject `runtime.NumCPU()` workers and a buffered channel (buffer = 1024).
- [x] **1.2** Create dispatcher goroutine with `filepath.WalkDir`.
- [x] **1.3** Add binary detection: read first 512 bytes, check for null byte.
- [x] **1.4** Implement worker logic: receive path, skip binary, read full content (with pool for ≤64KB), match against query/pattern.
- [x] **1.5** Aggregate results from workers, cap at 100.
- [x] **1.6** Remove sequential walk; ensure error handling (continue on permission errors).

### 2. Update/Add Tests
- [x] **2.1** Update `search_test.go` to test parallel behavior (concurrent safety).
- [x] **2.2** Add `benchmark_test.go` with both sequential (old) and parallel (new) benchmarks.

### 3. Verify
- [x] **3.1** Run `go test -v -count=1 ./...` – all tests pass (19/19).
- [x] **3.2** Run benchmarks: `go test -bench=. -benchmem -count=1 -timeout 60s`.
- [x] **3.3** Manual smoke test: created temp dir with 4000 files, search completed in ~1.5s.

## Risk Assessment
- **Risk**: Deadlock if channel is not closed properly. Mitigation: use `sync.WaitGroup` for workers and close results channel after all workers finish.
- **Risk**: Memory blow if many large files are read concurrently. Mitigation: limit concurrency to `runtime.NumCPU()` and skip files > 2MB (existing behavior).
- **Risk**: Race condition on results slice. Mitigation: use a mutex or a separate results collector goroutine.

## Dependencies
- None beyond Go standard library.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "search-optimization"
phase: "PLAN"
status: "IN_PROGRESS"
last_update: "2026-05-23T16:00:00Z"
```
