# Harness TUI Footer — Contract

## Validation Sensors

### Sensor 1: Build
```bash
cd /home/kleberson/Documentos/skills/bin/harness && go build ./...
```
Expected: Exit code 0, compile success.

### Sensor 2: Tests
```bash
cd /home/kleberson/Documentos/skills/bin/harness && go test ./... -count=1 -timeout 30s
```
Expected: All tests pass.

### Sensor 3: DrawFixedFooter Method
- `tui.go` must implement the `DrawFixedFooter` signature.
- Check with: `grep -c "DrawFixedFooter" internal/tui/tui.go` → must be ≥ 1

### Sensor 4: stty size query helper
- `tui.go` must implement `getTerminalSize` or equivalent helper calling `stty`.
- Check with: `grep -c "stty" internal/tui/tui.go` → must be ≥ 1

### Sensor 5: Reset Scroll Region on Shutdown
- `interactive.go` must clear and reset scroll region using `\033[r` or similar sequence on loop exit.
- Check with: `grep -c "\\033\[r" interactive.go` → must be ≥ 1

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
