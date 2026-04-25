# Contract: Harness Audit Gate

## Command Interface

| Command | Args | Flags | Description |
| :--- | :--- | :--- | :--- |
| `hb harness audit` | `[path]` | `--sec` | Runs only security audit. |
| | | `--mentor` | Runs only Clean Code audit. |
| | | `--all` | Runs all audits (default). |
| | | `--json` | Output result in JSON format. |

## Output Artifact

**File**: `validation-report.md` (updated/created in the root of the audited path).

### Format
```markdown
# Validation Report: [Feature Name]
- **Date**: 2026-04-25
- **Status**: ✅ PASS / ❌ FAIL

## 🛡️ Security Audit
- [x] Secret Leak Detection: OK
- [ ] Dependency Check: 1 Warning

## 💎 Quality Audit (Mentor)
- [x] SOLID Compliance: OK
- [x] DRY Principle: OK

## 🚨 Critical Issues
- No critical issues found.
```

## Exit Codes
- `0`: Success.
- `1`: Error or Critical Issue found.
