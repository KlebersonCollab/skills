# Testing Foundations in Go

The native `testing` package is the most powerful tool to ensure the correctness of your code. This guide covers the essential patterns.

---

## 1. Anatomy of a Test

- **Files**: Must end with `_test.go`.
- **Functions**: Must start with `Test` and accept `*testing.T`.
- **Command**: `go test ./...` to run all tests in the project.

```go
func TestHello(t *testing.T) {
    got := Hello("World")
    want := "Hello World"
    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}
```

---

## 2. Table-Driven Tests (TDT)

This is the idiomatic Go pattern for testing multiple scenarios cleanly.

```go
func TestSum(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive numbers", 1, 2, 3},
        {"negative numbers", -1, -1, -2},
        {"zero", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Sum(tt.a, tt.b); got != tt.want {
                t.Errorf("Sum() = %v, want %v", got, tt.want)
            }
        })
    }
}
```

---

## 3. Assertions and Failures

- **t.Error/t.Errorf**: Records the failure but continues test execution.
- **t.Fatal/t.Fatalf**: Records the failure and stops the test immediately (useful if the rest of the test depends on that step).
- **stretchr/testify**: A popular library that adds readable assertions (`assert.Equal`, `assert.NoError`).

---

## 4. Setup and Teardown

- **TestMain**: Use the `TestMain(m *testing.M)` function for global suite configurations (e.g., starting a temporary database).
- **t.Cleanup**: The preferred method for cleaning up specific resources of a test (Go 1.14+).

---

## 5. Code Coverage

- **Command**: `go test -cover ./...`
- **Visual Report**: 
  1. `go test -coverprofile=cover.out ./...`
  2. `go tool cover -html=cover.out`

---

## 📚 Supporting Libraries (Testing Foundations)
- `stretchr/testify/assert`: For simple assertions.
- `stretchr/testify/require`: For fatal failures.
