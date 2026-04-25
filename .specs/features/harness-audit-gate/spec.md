# Specification: Harness Audit Gate

## Goal
Implement a unified "Quality & Security Gate" in the `hb` CLI to ensure that features meet the minimum standards of the ecosystem before synchronization.

## Context
Currently, audits for Clean Code and Security are fragmented or manual. This feature centralizes these checks under the `harness` command, facilitating the "Exit Gate" protocol of the `harness-expert` skill.

## User Stories
- **As an AI Agent**, I want to run a single command to verify if my implementation is "production-ready".
- **As a Developer**, I want the CLI to block synchronization if critical security or quality issues are found.

## Acceptance Criteria (AC)

### AC 1: Command Availability
- [ ] Command `hb harness audit` must be implemented.
- [ ] Must accept an optional path (defaulting to current directory `.`).

### AC 2: Security Scan (`--sec`)
- [ ] Must check for hardcoded secrets (API Keys, tokens, passwords).
- [ ] Must check for common vulnerabilities (OWASP Top 10) in provided files.
- [ ] Must verify dependency health (if applicable).

### AC 3: Mentor Scan (`--mentor`)
- [ ] Must integrate the `clean-code-mentor` audit logic.
- [ ] Must identify violations of SOLID, KISS, and YAGNI.

### AC 4: Exit Codes and Reports
- [ ] Exit code `0` for success (no critical issues).
- [ ] Exit code `1` for failure (critical issue found).
- [ ] Automatically generate/update `validation-report.md` with the results.

### AC 5: Integration with Sync
- [ ] Update `hb sync` to optionally run this audit or check for a recent valid audit report.
