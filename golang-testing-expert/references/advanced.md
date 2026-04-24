# Pillar 2: Advanced Testing in Go

This guide explores advanced techniques to ensure the quality of complex systems, focusing on isolating dependencies and validating complex interfaces like HTTP APIs.

---

## 🎭 Mocking & Dependency Injection

In Go, the best way to mock is through interfaces. Create lightweight manual mocks to avoid heavy framework dependencies when possible.

### Manual Mock Example
```go
type MockStore struct {
    GetFunc func(id string) (User, error)
}

func (m *MockStore) Get(id string) (User, error) {
    return m.GetFunc(id)
}
```

### Best Practices
- **Manual Mocks**: Prefer manual mocks based on interfaces when simulation logic is simple.
- **Mockery / GoMock**: Use generators when you need complex assertions (e.g., "method called exactly twice").
- **Interface Segregation**: Keep interfaces small to facilitate mock creation.

---

## 🌐 HTTP Testing (Package `httptest`)

Use `net/http/httptest` to test handlers without starting a real server.

### Essential Components
- **httptest.NewRecorder()**: Captures the handler's response.
- **httptest.NewRequest()**: Creates a mocked request.
- **httptest.NewServer()**: Starts a real server on a random port (ideal for testing HTTP clients).

### Handler Test Example
```go
func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/health", nil)
    w := httptest.NewRecorder()
    HealthHandler(w, req)
    
    if w.Code != http.StatusOK {
        t.Errorf("got status %d", w.Code)
    }
}
```

---

## 🏆 Golden Files

To test large outputs (complex JSON, HTML, generated files), save the "expected result" in a file within the `testdata/` folder.

- Compare the current result with the content of the golden file.
- Use a flag (`-update`) to update golden files when behavior changes intentionally.

---

## 📦 TestData & Fixtures

- Use the `testdata/` folder for storing static test files. Go ignores this folder during compilation.
- Prefer injecting data via code (Factory Pattern) rather than depending on complex states in external files when possible.

---

## 🚀 Integration & Concurrency

### Integration Testing
- **Testcontainers**: Use `testcontainers-go` to run real dependencies (Postgres, Redis) in Docker.
- **Gock**: Excellent for mocking external HTTP calls made by your system.

### Concurrency Testing
- **Race Detector**: Always run `go test -race ./...`.
- **Determinism**: Avoid `time.Sleep`. Use channels or `sync.WaitGroup` to coordinate task completion.

---

## 📚 Recommended Libraries
- `testcontainers/testcontainers-go`: For Docker-based integration tests.
- `h2non/gock`: For mocking outgoing HTTP calls.
- `stretchr/testify/assert`: For more readable assertions.
