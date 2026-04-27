# Pillar 3: Development & Tooling

> Development in Go is empowered by integrated tools that ensure the quality and agility of the software lifecycle.

---

## 💻 CLI Applications (Command Line)

Go is the language of choice for modern CLI tools:
- **Flag Package**: Use the native `flag` package for simple applications.
- **Cobra & Viper**: Use libraries like `spf13/cobra` and `spf13/viper` for complex applications with subcommands and flexible configurations.
- **Interactivity**: Provide clear visual feedback and support for input/output pipes.

---

## 📦 Dependency Management (Go Modules)

The module system is the foundation for reproducibility:
- **go mod tidy**: Keep the `go.mod` and `go.sum` files always clean and synchronized.
- **Vendoring**: Consider `go mod vendor` in corporate environments where direct internet access is restricted in CI/CD.
- **Semantic Versioning**: Strictly follow semantic versioning when publishing public packages.

---

## 🔍 Linting & Static Analysis (Code Analysis)

Static analysis is mandatory for Expert code:
- **golangci-lint**: The standard orchestrator of linters in Go. Configure it with a `.golangci.yml` file, activating `errcheck`, `staticcheck`, `revive`, `goimports`, `misspell`, and `govet`.
- **Gofmt / Goimports**: Automatic formatting and intelligent management of imported packages (`goimports -w .`).
- **Custom Linters**: Use specific linters for security (`gosec`) and logical bugs (`errcheck`, `staticcheck`).
- **Essential Commands**:
  - `go mod tidy`: Synchronize dependencies.
  - `go test -race -cover ./...`: Run tests with race condition detection and coverage.
  - `go build ./...`: Validate compilation of the entire project.
  - `golangci-lint run`: Execute full static analysis.

---

## 🧪 Testing Strategies (Expert Testing)

In Go, test code is first-class code:
- **Native Testing**: Use the `testing` package and the `_test.go` structure.
- **Table-Driven Tests**: Standardize your test cases into tables for greater readability and coverage.
  ```go
  func TestSum(t *testing.T) {
      cases := []struct { a, b, want int }{ {1, 2, 3}, {4, 5, 9} }
      for _, c := range cases {
          if got := Sum(c.a, c.b); got != c.want { ... }
      }
  }
  ```
- **Testify**: Use the `stretchr/testify` library for rich assertions and robust mocks.
- **Sub-tests**: Use `t.Run` to isolate test cases within the same function.
- **Specialization**: For advanced TDD, Mocks, Benchmarks, and Fuzzing patterns, use the specialized skill [golang-testing-expert](../../golang-testing-expert/SKILL.md).

---

## 🛠️ Troubleshooting & Debugging (Problem Solving)

- **Delve (dlv)**: The standard debugger for Go. Use it for deep inspection of memory and execution.
- **Stack Traces**: Understand how to read and interpret panic traces and goroutine dumps.
- **Profiling**: Use `pprof` to identify CPU, memory, and lock contention bottlenecks.

---

## 📚 Supporting Libraries (Development)
- `stretchr/testify`: For assertions and mocks.
- `samber/lo`: For slice and map manipulation utilities in tests.
