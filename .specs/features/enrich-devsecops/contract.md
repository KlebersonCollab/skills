# Contract: Enrichment of DevSecOps Expert

## Agreement
The Implementer will enrich the `devsecops-expert` skill with specific guardrails for agentic safety, ensuring that all dangerous patterns identified in the mining phase are documented and actionable.

## Sensors (Validation)
1.  **Pattern Presence**: `SKILL.md` must contain the string `DANGEROUS_BASH_PATTERNS`.
2.  **BDD Compliance**: The skill must explicitly mention that agents should ask for confirmation when a dangerous pattern is detected.
3.  **Hub Audit**: `hb audit` must pass with score > 90.
