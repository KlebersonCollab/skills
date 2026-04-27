---
name: "devsecops-expert"
version: "1.0.0"
description: "Skill focused on DevSecOps, Information Security, SAST, DAST, and hardening practices in infrastructure and code."
uses:
  - "architecture"
  - "git-workflow"
references:
  - "sdd"
---

# 🛡️ DevSecOps Expert Skill

## 1. Description and Purpose
The `devsecops-expert` skill elevates software security quality within the ecosystem, guiding the agent to audit dependencies, validate configuration, and adopt "Secure-by-Design" practices.

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Did you rehydrate the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)?
3. **Plan Check**: Does the `plan.md` file define the architecture and schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

Before implementing or validating solutions involving infrastructure, keys, public routes, or pipelines:
- **Check for Secret Leaks**: Never commit hardcoded keys, passwords, or tokens.
- **Validate Attack Surface**: Every exposed port or endpoint must use RBAC/Authentication.
- **Initial SAST**: Ensure dependencies do not have open CVEs before running scripts.

## Goal
Ensure and shield the continuous development cycle (DevSecOps), providing static (SAST) and dynamic (DAST) auditing in code and infra.

## Output Structure
The execution of this skill produces the following artifacts:
| Artifact | Location | Description |
|----------|-------------|-----------|
| **Validation Report** | `validation-report.md` | Security report. |

## Quality Rules
- Zero tolerance for credential leaks.
- Every structural modification must pass SAST.

## Prohibited
- NEVER upload keys to version control.
- NEVER ignore injection failures or CVEs exposed in logs.

## 3. Security Standards (Mandatory)
1. **Zero Trust Architecture**: No microservice or component trusts another natively without tokenization/mutual TLS.
2. **Dependency Audit**: Require the use of scanning tools (e.g., `bandit`, `trivy`, or `npm audit`).
3. **OWASP Top 10**: Active prevention against SQL/NoSQL injection, XSS, CSRF, and Broken Object Level Authorization (BOLA).

## 4. Adversarial Review Guide (Security Gate)
- The `harness-expert` may call this skill at the end of an architectural refactoring to "attack" the created logical decisions (Threat Modeling / STRIDE).
