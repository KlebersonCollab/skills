# Pillar 2: Concurrency & Advanced

> Concurrency in Go is not about parallelism, but about the composition of independent tasks that communicate via channels and cancellation signals.

---

## 🏎️ Concurrency Fundamentals (Goroutines & Channels)

- **Share Memory by Communicating**: Follow the Go mantra — do not communicate by sharing memory; share memory by communicating.
- **Channels**: Use channels for synchronization and data transfer between goroutines.
  - Buffered channels for controlled temporal decoupling.
  - Unbuffered channels for rigid synchronization.
- **Select Statement**: Use `select` to orchestrate multiple channels and timeouts.

---

## 🚦 Context Management (`context` Package)

The `context` package is the backbone of any resilient system in Go:
- **Propagation**: Always pass `context.Context` as the first argument to long-running functions (IO, DB, HTTP).
- **Cancellation**: Use `WithCancel`, `WithTimeout`, or `WithDeadline` to avoid "zombie" goroutines that run indefinitely.
- **Values**: Use `WithValue` only for transient and global data (e.g., Request ID, Logs), never for business logic.

---

## 🛡️ Concurrency Safety (Security)

- **Race Conditions**: Use `go test -race` religiously to detect race conditions.
- **WaitGroups**: Use `sync.WaitGroup` to wait for a set of goroutines to finish.
- **errgroup**: Use `golang.org/x/sync/errgroup` to coordinate groups of goroutines that can return errors, allowing the entire group to be canceled if one fails.
  ```go
  g, ctx := errgroup.WithContext(ctx)
  for _, url := range urls {
      url := url
      g.Go(func() error {
          return fetch(ctx, url)
      })
  }
  if err := g.Wait(); err != nil { return err }
  ```
- **Mutexes**: Use `sync.Mutex` or `sync.RWMutex` to protect shared state when channels are not the best option.
- **Atomic Operations**: For simple counters or state flags, use `sync/atomic` for maximum performance without locks.

---

## 🧩 Advanced Design Patterns

- **Fan-out/Fan-in**: Distribute heavy work across multiple goroutines and consolidate the results into a single channel.
- **Worker Pools**: Limit the number of concurrent goroutines to avoid resource overload.
- **Graceful Shutdown**: Capture OS signals (SIGINT, SIGTERM) to close resources and finalize goroutines before terminating the process.
- **Circuit Breaker**: Implement resilience patterns for cascading failures in external services.
- **Singleton**: Use `sync.Once` for safe one-time initializations for multiple goroutines.

---

## 💉 Dependency Injection (DI)

- **Interface-based**: Inject dependencies as interfaces, not as concrete types, to facilitate mocks in testing.
- **Constructor Pattern**: Use `New...` functions to initialize structs with their dependencies.
- **DI Containers**: For large-scale projects, use dependency injection containers to avoid "Constructor Hell".

---

## 📚 Supporting Libraries (Concurrency & Advanced)
- `samber/do`: A simple and powerful type-based dependency injection library.
- `samber/hot`: For dynamic reloading of configurations or code without interrupting the service.
