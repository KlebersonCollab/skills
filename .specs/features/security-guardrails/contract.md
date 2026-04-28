# Validation Contract: Security Guardrails

## Validation Sensors
- [ ] **Sensor: Secret Detection**: 100% detection rate for standard API Keys in tests.
- [ ] **Sensor: PII Accuracy**: Must detect CPF patterns with < 5% false positives.
- [ ] **Sensor: Mandatory Block**: Commit must fail if a security violation is found.

## Evidence
- Test vectors in `internal/security/scanner_test.go`.
- Audit logs showing blocked attempts.
