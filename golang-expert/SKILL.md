---
name: golang-expert
version: 1.2.0
description: "Expert level Go development skill focused on performance, idiomatic concurrency, and clean architecture."
category: development
---

# Golang Expert

> Expert level Go development skill focused on performance, idiomatic concurrency, and clean architecture.

---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the guidelines of the `token-distiller` skill.
1. **Context Check**: Have you rehydrated the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define the architecture, schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---

## ⚡ Slash Commands (Operational)
Use these commands to guide your behavior during the session:

| Command | Description | Expected Action |
|---------|-----------|---------------|
| `/plan` | Feature Planning | Generates/updates `spec.md` and `plan.md` with ADRs. |
| `/go-build` | Compilation Fix | Runs `go build ./...` and fixes reported errors. |
| `/go-test` | TDD Cycle | Starts TDD workflow (Red-Green-Refactor) with `golang-testing-expert`. |
| `/go-review` | Code Review | Reviews Go idioms, error handling, and concurrency. |
| `/go-bench` | Performance Benchmark | Runs **`hb bench <target>`** and compares with baseline. |
| `/security-scan` | Security Audit | Searches for exposed secrets and known vulnerabilities. |

---

## Goal

Act as a senior software architect and engineer specialized in Go (Golang). This skill provides guidelines, patterns, and best practices to build scalable, resilient, and high-performance systems, utilizing the full idiomatic potential of the language and its modern ecosystem.

---

## 🛠️ Expert Workflows

### 1. New Project Scaffolding & Architecture
When starting a new project, follow the **Clean Architecture** structure:
1. `go mod init <module_path>`
2. Folder layout:
   - `cmd/<app>/main.go`: Entry point (DI, Graceful Shutdown).
   - `internal/domain/`: Entities, repository interfaces, and sentinel errors.
   - `internal/service/`: Pure business logic.
   - `internal/repository/`: Data access (e.g., `postgres/` using `sqlc`).
   - `internal/handler/`: Transport (`grpc/`, `rest/` subfolders).
   - `internal/config/`: Loading via env/yaml.
3. Mandatory Tools:
   - **sqlc**: For type-safe SQL code generation.
   - **Wire**: For explicit dependency injection.
   - **golangci-lint**: For rigorous static analysis.

### 2. High-Concurrency Design
For concurrent flows:
1. Design the Data Flow first.
2. Identify contention points.
3. Use `context.Context` for cancellation propagation.
4. Use `golang.org/x/sync/errgroup` for goroutine coordination and error propagation.
5. Implement Graceful Shutdown for all goroutines.
6. Validate with `go test -race`.

---

## ⚡ Performance Checklist
- [ ] Has the code passed `go test -bench`?
- [ ] Was escape analysis (`-gcflags="-m"`) verified to avoid unnecessary Heap allocations?
- [ ] Were Slices and Maps pre-allocated when size is known (`make([]T, 0, cap)`)?
- [ ] Was `strings.Builder` used for string concatenation in loops?
- [ ] Were object pools (`sync.Pool`) used for frequently allocated structures?
- [ ] Is error handling using `wrapping` (%w) efficiently?
- [ ] Is the `slog` package configured for the correct log level in production?

---

## Output Structure

The execution of this skill results in the following artifacts and guidance:

| Artifact | Description |
|----------|-----------|
| **Architecture ADR** | Grounded technical decisions for Go projects. |
| **Idiomatic Code** | Implementations that strictly follow `go-code-style` and `naming`. |
| **Concurrency Flow** | Secure concurrency flow designs (goroutines, channels, context). |
| **Optimization Report** | Performance analysis, benchmarks, and profiling suggestions. |
| **Validation Suite** | Robust unit and integration tests (with `testify`). |

---

## Quality Rules

- **Idiomatic First**: Always prioritize simplicity ("Clear is better than clever") and official conventions (`go fmt`, `go lint`).
- **Interfaces and Structs**: "Accept interfaces, return structs".
- **Zero Value Useful**: Design types so their zero value is usable without explicit initialization (e.g., `sync.Mutex`, `bytes.Buffer`).
- **Functional Options**: Use the Functional Options pattern for APIs with complex or optional configurations.
- **Consumer-Defined Interfaces**: Define interfaces in the package that consumes them, not the one that provides them, to avoid unnecessary coupling.
- **Safety Always**: Avoid race conditions, memory leaks, and improper use of `unsafe`.
- **Modern Go**: Use features from recent versions (Generics, `slog`, `errors.Join`, etc.).
- **Zero-Boilerplate**: Use utility libraries like `samber/lo` and `samber/mo` to reduce repetitive code noise.
- **Documentation**: Exported types and functions **must** have documentation comments.

## Prohibited

- **NEVER** ignore errors (`_ = err`); always handle or propagate with context.
- **NEVER** use `panic` for normal flow control; reserve for catastrophic errors.
- **NEVER** start goroutines without a clear cancellation or termination mechanism (leaks).
- **NEVER** place `context.Context` inside structs; it must always be the first parameter of functions.
- **NEVER** mix business logic with infrastructure details (Clean Architecture).
- **NEVER** use the `init()` function for initialization; prefer explicit injection in `main()`.
- **NEVER** use global mutable state.

---

## 📚 References

This skill is organized into 5 fundamental pillars:

1. [Foundations & Style](references/foundations.md)
2. [Concurrency & Advanced](references/concurrency.md)
3. [Development & Tooling](references/development.md)
4. [Systems & Architecture](references/architecture.md)
5. [Ecosystem & Samber](references/ecosystem.md)
6. [Testing Expert (Specialized Skill)](../golang-testing-expert/SKILL.md)
