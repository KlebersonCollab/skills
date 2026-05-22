# Specification: LLM API Timeout and Automatic Retry

## Feature Overview
This feature introduces network resilience to the Harness CLI by implementing a strict **2-minute timeout** on LLM API calls and an **automatic retry mechanism** when a call times out or encounters a transient network error.

## Requirements

### 1. 2-Minute Request Timeout
- **Strict Limit**: Every HTTP request sent by `CallLLM` in [client.go](file:///home/kleberson/Documentos/skills/bin/harness/client.go) must be constrained by a strict 2-minute deadline (`2 * time.Minute`).
- **Context Integration**: Utilize Go's `context.WithTimeout` to handle request cancellations gracefully if the remote server fails to respond within the allocated time window.

### 2. Automatic Retry Mechanism
- **Retry Action**: If the request fails due to a network timeout, transient server error, or DNS lookup failure, `CallLLM` must automatically retry the HTTP request.
- **Max Attempts**: The total number of attempts should be bounded to a safe maximum (e.g., up to 3 total attempts: 1 initial call + 2 retries).
- **Graceful Detection**: Verify if the error is a temporary/timeout network error (by asserting `err` to `net.Error` and calling `err.Timeout()`, or checking for `context.DeadlineExceeded`).

### 3. Premium Terminal Log Visibility
- **Clean Interface Feedback**: When a timeout triggers, the Harness CLI must output a clear and premium Unicode warning to the console, informing the user about the failure and indicating the subsequent retry count (e.g., `⚠️  Timeout de 2 minutos atingido para [PROVIDER]. Iniciando nova tentativa ([X]/3)...`).
- **Terminal UI Synchronization**: Ensure this message cleans the current spinner output line appropriately using ANSI escape codes (`\r\033[K`) so that the spinner can continue rendering correctly on subsequent attempts without overlapping text.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HARNESS-TIMEOUT-RETRY"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-22T15:41:00Z"
evidence_checksum: "go-test-pass-retry-001"
```
