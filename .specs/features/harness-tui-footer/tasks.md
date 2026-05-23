# SDD Tasks: tasks.md (`HARNESS-TUI-FOOTER`)

Atomic task list for tracking implementation progress of the fixed TUI status footer.

---

## 📊 Phase Progress Monitoring

| Phase | Task | Status | Evidence (Commit/Log) |
| :--- | :--- | :---: | :--- |
| **1. PREP** | Bootstrap Spec/Plan/Tasks/Contract | [x] | spec.md, plan.md, tasks.md, contract.md |
| **2. CORE** | Implement `getTerminalSize` and `DrawFixedFooter` in `tui.go` | [x] | tui.go |
| **2. CORE** | Integrate scroll region setup and footer rendering in `interactive.go` | [x] | interactive.go |
| **2. CORE** | Add exit and interrupt handlers to reset scroll region on CLI shutdown | [x] | interactive.go (defer region reset) |
| **3. FINAL** | Verify compilation and run test suites (`go test ./...`) | [x] | make harness-build && make harness-test |
| **3. FINAL** | Audit project governance compliance (`make audit`) | [x] | make audit (100% COMPLIANT) |

---
<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HARNESS-TUI-FOOTER"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-23T04:02:00Z"
evidence_checksum: "go-build-and-test-pass"
```
