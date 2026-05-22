
# Learnings & Patterns — Skills Hub

## Recent Learnings (2026-05-23)

### SSE Streaming in Go
- **Pattern**: Using `http.Client` with `io.Copy` and a channel to stream tokens in real-time.
- **Key Implementation** (`client.go`):
  - The SSE stream is initiated via a POST request with `Accept: text/event-stream`.
  - Response body is read line-by-line, parsing `data:` lines for token content.
  - Tokens are pushed into a `chan string` which is consumed by the caller.
  - Retry logic (3 retries, 2-minute timeout) ensures resilience against transient failures.
- **Lesson**: Go channels provide a clean, goroutine-safe way to stream partial responses without blocking.

### Testing Strategy
- **Smoke tests** (`smoke_test.go`) validate the full pipeline end-to-end.
- **Unit tests** (`harness_test.go`, `distiller_test.go`) cover edge cases (security checks, caveman modes, session DAG).
- **Retry tests** (`TestCallLLMRetryOnTimeout`) verify timeout handling without actually calling an external LLM (mock).

### Token Distiller Patterns
- **Auto-Clarity**: Three levels of security checks are applied before compression (security, irreversible, multi-step).
- **Caveman modes**: Lite, Full, Ultra — progressively more aggressive compression.
- **Backtick protection**: Preserves code blocks during compression.

## Patterns to Reuse
1. **Modular Go project structure**: Keep core logic in `client.go`, `session.go`, `tools.go`, each with clear responsibilities (SRP).
2. **SSE with channels**: Use `chan string` for streaming, with a context or timeout for cancellation.
3. **Test isolation**: Use environment variable mocks (`HARNESS_MOCK_LLM`) to avoid external dependencies.

## Problems Solved
- **Timeout handling**: Added retry mechanism with exponential backoff for LLM calls.
- **JSON parsing errors**: Robust selector extraction with fallback and error reporting.

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "harness-streaming-sse"
phase: "LEARNINGS"
status: "COMPLETED"
last_update: "2026-05-23T16:00:00Z"
evidence_checksum: "go-test-pass-1.106s"
```
