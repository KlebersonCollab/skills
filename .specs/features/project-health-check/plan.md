# Plan: Project Health Check

## 1. Technical Design
A Python script that uses regex to find the metadata block.

## 2. Execution Steps
1. Create `bin/check-health.py`.
2. Implement logic to scan `STATE.md`.
3. Add colorized output (Green for OK, Red for Fail).

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HEALTH-CHECK-01"
phase: "SPECIFY"
status: "COMPLETED"
last_update: "2026-05-06T09:36:10Z"
evidence_checksum: "NONE"
```
