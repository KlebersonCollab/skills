---
name: golang-testing-expert
version: 1.1.0
description: "Expert level Go testing patterns including TDD, table-driven tests, benchmarks, fuzzing, and HTTP testing."
category: development-testing
---

# Golang Testing Expert

> Expert level Go testing patterns including TDD, table-driven tests, benchmarks, fuzzing, and HTTP testing.

---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Did you rehydrate the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define the architecture, schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---

## Goal

Act as a QA and Testing expert for Go (Golang). This skill provides rigorous guidelines for implementing the TDD cycle, ensuring high code coverage, validating performance via benchmarks, and increasing resilience via fuzzing, following the language's most modern standards.

---

## 🛠️ Expert Workflows

### 1. The TDD Cycle (Red-Green-Refactor)
When developing new features, strictly follow:
1. **RED**: Write a test that fails first.
2. **GREEN**: Write the minimum code necessary to pass the test.
3. **REFACTOR**: Improve the code while keeping tests green.

### 2. Table-Driven Testing
Use anonymous structs to define multiple test cases in a single function, ensuring coverage of edge cases and errors.

### 3. Integration Testing with Testcontainers
For tests that depend on infrastructure (Database, Redis, etc.), use **testcontainers-go** to spin up ephemeral Docker instances, ensuring total isolation.

### 4. Verification Commands
Use the following commands to ensure quality:

| Scenario | Command |
|---------|---------|
| **Unit** | `go test ./internal/... -short -count=1` |
| **Integration** | `go test ./internal/repository/... -count=1 -timeout 120s` |
| **Race Detector** | `go test ./... -race -count=1` |
| **Coverage** | `go test ./... -coverprofile=coverage.out && go tool cover -func=coverage.out` |

---

## ⚡ Testing Checklist
- [ ] Was the RED-GREEN-REFACTOR cycle followed?
- [ ] Were Table-Driven Tests used for success and error cases?
- [ ] Were subtests (`t.Run`) used to isolate cases?
- [ ] Do helpers call `t.Helper()` for accurate error logs?
- [ ] Are resources cleaned up via `t.Cleanup()` or `defer`?
- [ ] Was code coverage (`go test -cover`) verified?
- [ ] Was the race detector (`go test -race`) executed?

---

## Output Structure

Execution of this skill results in the following artifacts and guidance:

| Artifact | Description |
|----------|-----------|
| **Test Suite** | Robust and idiomatic `_test.go` files. |
| **Mocks & Stubs** | Lightweight implementations for dependency isolation. |
| **Benchmarks** | Performance and memory allocation analysis. |
| **Fuzz Tests** | Mutation tests for complex input validation. |
| **Coverage Report** | Detailed code coverage reports. |

---

## Quality Rules

- **Behavior over Implementation**: Test public behavior, not internal details.
- **Fast Feedback**: Keep unit tests extremely fast.
- **Meaningful Names**: Use subtest names that explain the scenario (e.g., "success_valid_user").
- **No Sleep**: NEVER use `time.Sleep()` in tests; use channels or synchronization signals.
- **TDD First**: Always write the test before the production code.

## Prohibited

- **NEVER** use `time.Sleep()` for synchronization in tests.
- **NEVER** ignore errors returned in cleanup functions.
- **NEVER** fail to use `t.Helper()` in auxiliary test functions.
- **NEVER** commit production code without the corresponding test suite passing.
- **NEVER** use heavy mock frameworks when a simple interface solves the problem.

---

## 📚 References

This skill is organized into 3 fundamental pillars:

1. [Testing Foundations](references/foundations.md) (TDD, Table-driven, Subtests)
2. [Advanced Testing](references/advanced.md) (Mocks, Helpers, Golden Files, HTTP)
3. [Performance & Fuzzing](references/performance.md) (Benchmarks, Fuzzing, Coverage)
