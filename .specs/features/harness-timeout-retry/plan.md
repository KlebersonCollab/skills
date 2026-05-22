# Plan: Implementation Tasks for LLM API Timeout and Automatic Retry

This document outlines the detailed step-by-step programming tasks to achieve network resilience with a 2-minute request timeout and automatic retries in the Harness CLI.

## Tasks

### Phase 1: CallLLM Refactoring (Go Engine)
- [ ] **Task 1.1**: Update `CallLLM` in [client.go](file:///home/kleberson/Documentos/skills/bin/harness/client.go) to support a 3-attempt execution loop.
- [ ] **Task 1.2**: Implement `context.WithTimeout` (set to `2 * time.Minute`) inside `CallLLM` for each request attempt.
- [ ] **Task 1.3**: Detect network timeouts or connection reset errors safely (checking for `context.DeadlineExceeded` or `net.Error`'s `Timeout()`).

### Phase 2: Terminal UI feedback & Spinner Synchronization
- [ ] **Task 2.1**: Implement a premium ANSI warning log that cleans the terminal line (`\r\033[K`) before printing the retry notification.
- [ ] **Task 2.2**: Integrate visual styling following the Harness CLI yellow/orange premium Unicode design guidelines.

### Phase 3: Verification & Auditing
- [ ] **Task 3.1**: Create a unit test mock in [harness_test.go](file:///home/kleberson/Documentos/skills/bin/harness/harness_test.go) that simulates a slow endpoint (deliberately delaying response for >2 minutes or 2 seconds for faster test runs) to verify that `CallLLM` triggers a timeout and initiates retries.
- [ ] **Task 3.2**: Execute tests using `go test -v ./bin/harness/...` and verify correctness.
- [ ] **Task 3.3**: Run `make audit` to ensure 100% compliant status with SDD v2.3.0.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HARNESS-TIMEOUT-RETRY"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-22T15:41:00Z"
evidence_checksum: "go-test-pass-retry-001"
```
