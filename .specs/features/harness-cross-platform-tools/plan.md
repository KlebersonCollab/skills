# Plan: Implementation Tasks for OS-Agnostic Tools and Native File Search

This document outlines the detailed step-by-step programming tasks to achieve OS-agnostic execution and pure Go file searching inside the Harness CLI.

## Tasks

### Phase 1: Core Tools (Go Engine)
- [ ] **Task 1.1**: Refactor `ExecuteCommand` in [tools.go](file:///home/kleberson/Documentos/skills/bin/harness/tools.go).
  - Add OS runtime detection using `runtime.GOOS`.
  - Use `exitError.ExitCode()` instead of type assertions to `syscall.WaitStatus`.
- [ ] **Task 1.2**: Implement `SearchFiles(pattern string, query string)` in `tools.go` using pure Go stdlib (`filepath.WalkDir`, `os.Open`, `strings.Contains`).
  - Add ignore filters for system files.
  - Implement pattern matching and text query matching.

### Phase 2: Interpreter and Loop Integration
- [ ] **Task 2.1**: Update regex matching in [main.go](file:///home/kleberson/Documentos/skills/bin/harness/main.go) to support XML tag parsing.
- [ ] **Task 2.2**: Integrate `SearchFiles` into the loop in `main.go`. Handle tool responses and output logs nicely.
- [ ] **Task 2.3**: Update LLM system instructions with the new tool description.

### Phase 3: Verification
- [ ] **Task 3.1**: Write tests in [harness_test.go](file:///home/kleberson/Documentos/skills/bin/harness/harness_test.go) for `SearchFiles`.
- [ ] **Task 3.2**: Execute tests and run compliance audit (`make audit`).

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HARNESS-CROSS-PLATFORM-TOOLS"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-22T15:23:00Z"
evidence_checksum: "go-test-and-make-audit-pass-004"
```
