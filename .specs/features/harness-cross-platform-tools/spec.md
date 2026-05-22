# Specification: OS-Agnostic Tools and Native File Search

## Feature Overview
This feature introduces complete OS-agnostic support to the Harness CLI executive tools and implements a native, pure Go **file search tool** (`search_files`).

## Requirements

### 1. OS-Agnostic Execution Tools
- **Subprocess Execution**: The Harness CLI must run shell/bash commands transparently on both Unix-based systems (Linux/macOS) and Windows.
- **Dynamic Shell Resolution**: If running on Windows, command routing must dynamically swap the execution shell (using `cmd.exe /c` or PowerShell as fallback when bash is unavailable).
- **Cross-Platform Exit Status**: Replace Unix-specific `syscall.WaitStatus` extraction with standard Go library `exitError.ExitCode()`.

### 2. Native File Search Tool
- **No System Tool Dependency**: Searching must not rely on `grep` or `find`. It must execute in pure Go.
- **Search Capabilities**:
  - **Pattern Matching (Glob)**: Support wildcard file name matching (e.g., `*.go`, `*main*`).
  - **Text Matching (Query)**: Support case-insensitive text string search inside files, returning file paths and line numbers.
- **Safety Boundaries**: Only search inside the resolved workspace root, ignoring folders like `.git`, `.harness`, `.gemini`, `.claude`, `node_modules`, `bin`, and files larger than 2MB.

### 3. XML Tool Schema Integration
- Implement the XML tag parsing:
  `<tool:search_files pattern="*.go" query="some_text" />`
  (both attributes are optional).

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HARNESS-CROSS-PLATFORM-TOOLS"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-22T15:23:00Z"
evidence_checksum: "go-test-and-make-audit-pass-004"
```
