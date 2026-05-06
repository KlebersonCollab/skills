# Pillar 3: Performance & Fuzzing in Go

Validate not only logical correctness but also efficiency and resilience against unexpected inputs.

---

## ⏱️ Benchmarking

Go has native support for benchmarks integrated into the `go test` command.

- Functions must start with `Benchmark`.
- Use the `b.N` loop for repeated executions controlled by the runtime.
- **b.ResetTimer()**: Use after heavy setups so as not to distort execution time.
- **b.ReportAllocs()**: Activates memory allocation reporting per operation.

```go
func BenchmarkProcess(b *testing.B) {
    data := setup()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Process(data)
    }
}
```

---

## 🧪 Fuzz Testing (Go 1.18+)

Fuzzing generates random inputs to find security flaws or crashes that manual tests would not catch.

- Use `f.Add()` to provide initial examples (seed corpus).
- Go will attempt to mutate these inputs to break your code.

```go
func FuzzParse(f *testing.F) {
    f.Add("initial input")
    f.Fuzz(func(t *testing.T, str string) {
        Parse(str) // Must be resilient to any string
    })
}
```

---

## 📊 Test Coverage

- **go test -cover**: Displays the coverage percentage in the console.
- **go test -coverprofile=c.out**: Generates a profile file.
- **go tool cover -html=c.out**: Opens a browser viewer highlighting untested lines.

---

## 🏁 Race Detection

Always run your tests with the `-race` flag in CI environments and during local development of concurrent code.

```bash
go test -race ./...
```

---

## 🔍 Profiling (pprof)

`pprof` is the definitive tool for identifying bottlenecks:

- **CPU Profile**: Where the code spends the most time.
- **Memory/Heap Profile**: Where excessive allocations occur.
- **Usage**:
  1. Generate profile: `go test -cpuprofile=cpu.out -bench=.`
  2. Analyze: `go tool pprof -http=:8080 cpu.out`
