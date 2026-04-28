# Validation Contract: Structured Review

## Validation Sensors
- [x] **Sensor: Diff Parsing Integrity**: Must correctly identify hunks in unified diffs.
- [x] **Sensor: Junk Detection Precision**: Must flag `fmt.Printf` and `console.log` without false positives.
- [x] **Sensor: Plan Alignment**: Must verify file presence in `plan.md` of the active feature.
- [x] **Sensor: Score Logic**: Score must decrease proportionally to the number of warnings.

## Evidence
- Unit tests in `internal/git/parser_test.go`
- Manual verification of `hb review` output with debug code injection.
