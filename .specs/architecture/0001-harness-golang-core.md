# ADR-0001: Golang-Core CLI and Modular HTTP Harness

We have decided to build the new terminal-based active execution harness (`bin/harness`) in **Go (Golang)** rather than Python or Node.js. It will interface with LLM APIs using direct HTTP/REST calls and execute dynamic skill-specific tools as local subprocesses.

## Context

The current ecosystem is a "passive" repository of skills (Markdown files). Transforming it into an active execution engine (like `pi.dev`) requires a runtime that is:
1. Fast, lightweight, and zero-dependency for final developers (compiled native binary).
2. Highly portable across Linux, MacOS, and Windows without runtime environments.
3. Extensible, capable of launching existing python scripts or bash files.
4. Model-agnostic and resilient to API deprecations.

## Decision

1. **Golang Implementation:** The harness will be a standalone compiled binary written in Go.
2. **REST Client Purism:** Instead of heavy SDK dependencies, Go's standard library (`net/http` and `encoding/json`) will be used to invoke LLM APIs.
3. **Dynamic Provider Templates:** Provider URLs, headers, and request body formats will be declared in a local `.harness/config.json` configuration file, allowing seamless hot-swapping and integration of new LLM endpoints without compiling.
4. **Subprocess Tool Execution (`os/exec`):** Domain-specific scripts (e.g. Python lints, shell scripts) are invoked as system-level subprocesses, allowing maximum backward compatibility with existing Skills Hub tooling.

## Status

Accepted

## Consequences

* **Portability:** Developers only need to run the compiled `harness` binary.
* **API Flexibility:** Supports new providers (Ollama, Gemini, local pipelines) instantly by writing JSON templates.
* **Separation of Concerns:** The execution engine (Go) remains decoupled from the tool scripts (Python/Bash).

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HARNESS-TRANSFORMATION"
phase: "ALIGN"
status: "COMPLETED"
last_update: "2026-05-22T14:41:00Z"
evidence_checksum: "ffb546e"
```
