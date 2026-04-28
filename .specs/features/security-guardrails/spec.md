# Specification: Security Guardrails

## Goal
Implement automated security gates to prevent sensitive data leakage (Secrets, PII) and ensure compliance with Hub security standards.

## Functional Requirements
1. **Secret Detection**: Scan staged changes for API keys, passwords, and tokens.
2. **PII Masking**: Detect patterns of Personally Identifiable Information (Emails, CPF, Credit Cards).
3. **Audit Logging**: Record all security violations in a dedicated log.

## Acceptance Criteria (BDD)
- **Given** a staged file contains a hardcoded API Key
- **When** the `hb review` or a pre-commit hook runs
- **Then** the commit must be blocked and a security warning displayed.

- **Given** a file contains a Brazilian CPF number
- **When** analyzed by the guardrails
- **Then** it must be flagged as a PII violation.
