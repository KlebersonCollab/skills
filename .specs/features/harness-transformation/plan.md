# Plan - Harness Transformation (Phase 1)

This plan details the surgical implementation of our active Go-based terminal agent harness.

---

## 🛠️ File Structure

We will implement the Go core inside the `bin/harness` directory as a Go module:
* `bin/harness/main.go` - Entrypoint and CLI parser.
* `bin/harness/session.go` - Directed Acyclic Graph (DAG) state, branching logic, and serialization.
* `bin/harness/client.go` - Purist HTTP REST client with dynamic provider configurations.
* `bin/harness/tools.go` - Safe execution of minimal tools (read, write, patch, bash).
* `bin/harness/sdd.go` - Compliance gates and `.specs/project/STATE.md` validation.

---

## ⚙️ Configuration File Schema (`.harness/config.json`)
```json
{
  "active_provider": "gemini",
  "providers": {
    "gemini": {
      "url": "https://generativelanguage.googleapis.com/v1beta/models/gemini-pro:generateContent?key={{GEMINI_API_KEY}}",
      "headers": {
        "Content-Type": "application/json"
      },
      "body_template": "{\"contents\": [{\"parts\": [{\"text\": \"{{prompt}}\"}]}]}",
      "response_path": "candidates.0.content.parts.0.text"
    },
    "ollama": {
      "url": "http://localhost:11434/api/generate",
      "headers": {
        "Content-Type": "application/json"
      },
      "body_template": "{\"model\": \"llama3\", \"prompt\": \"{{prompt}}\", \"stream\": false}",
      "response_path": "response"
    }
  }
}
```

---

## 📋 Steps for Phase 1 MVP
1. **Initialize Go Module:** Setup Go structure inside `bin/harness`.
2. **Implement Session DAG (`session.go`):** Struct and JSON marshaling for session tree nodes and active-pointer branching.
3. **Implement Dynamic REST Client (`client.go`):** Read config file, replace env variables in headers/URL, construct HTTP request payload, parse JSON response dynamically.
4. **Implement CLI Prompt Loop (`main.go`):** Core terminal prompt loop (`harness [node-0]> `) handling slash commands.
5. **Implement Safe Subprocess Tools (`tools.go`):** Basic read/write and subprocess wrapper for safe bash execution.
6. **SDD Integrity Gating (`sdd.go`):** Validate state phases before running file mutations.

---

## 🧪 Verification Plan

### Automated Tests
* Create a lightweight test harness runner in Go (`bin/harness/harness_test.go`) validating:
  * Dynamic provider configuration parsing.
  * Node branching in tree structure.
  * Subprocess command execution tracking.
* Execute `make audit` to verify compliant metadata.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HARNESS-TRANSFORMATION"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-22T14:48:00Z"
evidence_checksum: "go-test-pass-005"
```
