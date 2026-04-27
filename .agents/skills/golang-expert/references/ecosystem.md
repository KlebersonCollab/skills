# Pillar 5: Ecosystem & Samber

> The Go ecosystem has evolved rapidly with the introduction of Generics and native structured logging. Knowing the right tools differentiates a developer from an Expert.

---

## 🚀 Modernizing Go (Modernization)

Go 1.18 (Generics) and 1.21 (slog) changed the rules of the game:
- **Generics**: Use generic types to create reusable data structures and functions without sacrificing type safety (e.g., utility libraries).
- **Native Structured Logging**: Replace `logrus` or `zap` with the native `slog` package whenever possible to reduce external dependencies.
- **errors.Join**: Use the new multi-error aggregation functionality from Go 1.20+.

---

## 🧩 Samber Ecosystem (Expert Productivity)

The `samber` suite of libraries is a reference in modern Go for using Generics to solve common problems elegantly:

### 1. samber/lo (Lodash for Go)
The Swiss Army knife for collection manipulation:
- **Map / Filter / Reduce**: Transform slices and maps without verbose loops.
- **Find / Uniq / GroupBy**: Complex search and grouping operations in a single line.

### 2. samber/mo (Monads & Functional)
Bring functional programming patterns to Go idiomatically:
- **Option[T]**: Handle null/missing values without constant `nil` checks.
- **Result[T]**: Encapsulate results and errors in a single structure.

### 3. samber/do (Dependency Injection)
Simple and type-based dependency injection:
- **Service Discovery**: Register and retrieve dependencies without heavy reflection or complex global configurations.
- **Lifecycle Management**: Manage the lifecycle (initialization/shutdown) of your services.

### 4. samber/oops (Rich Errors)
Industrial-grade error handling:
- **Stack Traces**: Get the exact trace of where the error occurred.
- **Contextual Data**: Automatically attach metadata (User ID, Request ID) to errors.

---

## 📚 Other Popular Libraries

- **Echo / Chi / Gin**: HTTP frameworks and routers for all tastes and needs.
- **sqlc**: Compiles pure SQL into typed Go code — the best way to interact with SQL databases today.
- **GoMock / Testify**: Essential for ensuring the testability of complex systems.

---

## 📈 Staying Updated (Continuous Evolution)

The Go ecosystem never stops. An Expert must follow:
- **Go Blog**: The official channel for announcements of new versions and change proposals (Gopher proposals).
- **Go Weekly**: Weekly newsletter with the best articles and libraries from the community.
- **Github samber/cc-skills-golang**: The base repository that inspired this skill, always updated with new patterns.

---

## 💡 Expert Tip
Don't try to reinvent the wheel. If you need a utility function for slices, check `samber/lo` first. If you need robust error handling, use `samber/oops`. Productivity comes from knowing what to use and when to use it.
