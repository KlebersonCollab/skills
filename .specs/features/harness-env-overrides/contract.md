# Contract: Harness Environment Overrides & Root .env Support

This contract guarantees feature delivery through automated verification sensors.

## Acceptance Criteria

1. Harness builds successfully with `make harness-build`.
2. Existing and new tests pass with `make harness-test`.
3. Running `make audit` passes successfully ("100% COMPLIANT").
4. Root-level `.env.example` exists.

## Verification Sensors

- **Build Sensor**: `go build -o bin/harness-cli bin/harness/*.go` must exit with code 0.
- **Test Sensor**: `cd bin/harness && go test ./...` must exit with code 0.
- **Audit Sensor**: `make audit` must run successfully.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HARNESS-ENV-OVERRIDES"
phase: "SPECIFY"
status: "COMPLETED"
last_update: "2026-05-23T02:49:00Z"
evidence_checksum: "ffb546e"
```
