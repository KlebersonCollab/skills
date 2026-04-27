---
name: api-architect
version: 1.3.0
description: "API Architect — guides the agent to design interoperable, secure, and resilient API systems, defining contracts and governance standards."
category: api-design
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the guidelines of the `token-distiller` skill.
1. **Context Check**: Have you rehydrated the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define the architecture, schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---
# API Architect

> High-performance API designer and technical governance guardian. "Design is for developers; architecture is for the system."

---

## Goal

Empower the agent to design robust API ecosystems applying REST, GraphQL, and tRPC design principles, ensuring interoperability through rigorous **Style Guides**, security (OWASP), and resilience (Rate Limiting, Circuit Breaker).

The skill ensures that the API is not just "functional," but **consistent**, **secure**, and **sustainable** over time.

---

## When to Use This Skill

- Designing new APIs from scratch (REST, GraphQL, tRPC).
- Defining the **API Style Guide** to ensure consistency across multiple services.
- Refactoring legacy APIs to modern security and resilience standards.
- Implementing governance strategies (versioning, pagination, rate limit).
- Reviewing security against critical vulnerabilities (BOLA/BFLA).

---

## Workflow (4 Fases)

### Phase 1: STRATEGY — Protocol Definition
1.  **Consumer Analysis**: Identify who consumes (Web, Mobile, B2B).
2.  **Protocol Selection**: Choose between REST (Public), GraphQL (Flexible), or tRPC (TS Monorepo).
3.  **Governance Baseline**: Define initial version, deprecation policy, and rate limiting.

### Phase 2: DESIGN — Modeling and Style Guide
1.  **Resource Modeling**: Define entities and relations (plural, nouns).
2.  **Consistency Layer**: Rigorously apply `references/api-style.md` (standard error format, response envelopes).
3.  **Schema Definition**: Write OpenAPI (REST), SDL (GraphQL), or Zod (tRPC).

### Phase 3: RESILIENCE — Security and Performance
1.  **Security Audit**: Apply OWASP API Top 10 checklist (AuthN/AuthZ).
2.  **Stability Patterns**: Define timeouts, retries, and Circuit Breakers.
3.  **Data Filtering**: Implement pagination (Cursor/Offset) and field projection.

### Phase 4: VALIDATE — Automated Verification
1.  **Contract Testing**: Validate if the generated schema is valid and error-free.
2.  **Security Scan**: Review object permissions (BOLA) on all endpoints.
3.  **Documentation Handoff**: Deliver the approved contract and implementation guide.

---

## Output Structure

The execution of this skill must result in the following mandatory artifacts:

| Artifact | Format | Description |
|----------|---------|-----------|
| **API Spec** | `.yaml` / `.graphql` / `.ts` | Complete technical contract (OpenAPI, SDL, or Zod). |
| **API Style Guide** | `.md` | Definition of response, error, and naming patterns. |
| **Security Checklist** | `.md` | Compliance report with OWASP and Auth standards. |
| **Resilience Policy** | `.md` | Definition of rate limits, timeouts, and cache strategy. |

---

## Operational Mandates
- **Contract First**: Every API design must be documented in `spec.md` with a `## API Contract` section.
- **Automated Validation**: Upon finishing a design or update, the agent **MUST** run **`hb api contract <feature>`** to ensure documentation compliance.
- **Resilience**: Define retry patterns and circuit breakers in the architectural proposal.
- **Security by Default**: Every endpoint must be protected by default (requires explicit opt-out).
- **YAGNI in Data**: Return only what is necessary; avoid over-fetching through projections or optional fields.

## Prohibited

- NEVER expose sequential IDs (use UUIDs or HashIDs).
- NEVER use verbs in resource URLs (e.g., `/getUsers` -> use `GET /users`).
- NEVER ignore versioning; public APIs must have a version in the URL or Headers.
- NEVER return logs or stack traces in production error responses.

---

## References

- [`references/rest-best-practices.md`](references/rest-best-practices.md) — REST design guide.
- [`references/api-security-guide.md`](references/api-security-guide.md) — Security and OWASP API Top 10.
- [`references/api-style.md`](references/api-style.md) — (New) Interface consistency patterns.
- [`references/api-rate-limiting.md`](references/api-rate-limiting.md) — Resilience and rate control.
- [`resources/implementation-playbook.md`](resources/implementation-playbook.md) — Implementation checklists.
