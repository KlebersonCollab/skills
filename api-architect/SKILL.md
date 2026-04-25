---
name: api-architect
version: 1.4.0
description: "API Architect — guides the agent to design interoperable, secure, and agent-friendly API systems, defining contracts and intent-based integration patterns."
category: api-design
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the guidelines of the `token-distiller` skill.
1. **Context Check**: Have you rehydrated the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)?
3. **Plan Check**: Does the `plan.md` file define the architecture, schemas, and include **Mermaid** diagrams?

---
# API Architect (v1.4.0)

> High-performance API designer and technical governance guardian. "Design is for developers; architecture is for the system."

---

## 🏗️ Agent-Facing API Patterns

When designing APIs intended for LLM/Agent consumption (e.g., MCP servers):

1.  **Intent-Grouped Tools**: Do not map 1:1 to REST endpoints. Group tools by user intent (e.g., `create_issue_from_thread` instead of `get_thread` + `parse` + `create_issue`).
2.  **Thin Surface Pattern**: For massive APIs (200+ endpoints), expose a `search` tool and an `execute` tool in a secure sandbox instead of listing every endpoint.
3.  **Agent-Ready Outputs**: Return data optimized for model reasoning (summarized, normalized) rather than raw database dumps.
4.  **Inline UI Primitives**: Design responses that can return structured UI (charts, forms) via MCP Apps.

---

## Goal

Empower the agent to design robust API ecosystems applying REST, GraphQL, and agentic design principles, ensuring interoperability, security (OWASP), and sustainability.

---

## Workflow (4 Phases)

### Phase 1: STRATEGY — Protocol & Persona
1.  **Consumer Analysis**: Identify if the primary consumer is a Human, a Machine, or an **Agent**.
2.  **Protocol Selection**: REST (Public), GraphQL (Flexible), or MCP (Agent-specific).
3.  **Governance Baseline**: Define versioning, deprecation, and rate limiting.

### Phase 2: DESIGN — Modeling & Intent
1.  **Intent Mapping**: If designing for agents, group resources into "Intent-Based Tools".
2.  **Schema Definition**: Write OpenAPI, SDL, or MCP Tool Definitions.
3.  **Consistency Layer**: Apply `references/api-style.md` (standard error formats).

### Phase 3: RESILIENCE — Security & Sandbox
1.  **Security Audit**: Apply OWASP API Top 10.
2.  **Stability Patterns**: Define timeouts, retries, and Circuit Breakers.
3.  **Sandbox Design**: If using **Thin Surface**, define the execution environment boundaries.

### Phase 4: VALIDATE — Contract Testing
1.  **Contract Testing**: Validate schema against representative "User Intents".
2.  **Security Scan**: Review object permissions (BOLA/BFLA).
3.  **Handoff**: Deliver the approved contract and implementation guide.

---

## Output Structure

| Artifact | Format | Description |
|----------|---------|-----------|
| **API Spec** | `.yaml` / `.json` | Complete technical contract (OpenAPI or MCP). |
| **Intent Map** | Table | Mapping of user goals to API operations. |
| **Security Checklist** | `.md` | Compliance report with OWASP standards. |

---

## Quality Rules

- **Contract-First**: Code follows the contract, not the other way around.
- **Uniform Error Handling**: Standardized structure for all error responses.
- **YAGNI in Data**: Return only what is necessary for the task at hand.
- **Agent-Friendly**: Tool descriptions must clearly state *why* and *when* an agent should call them.

## Prohibited

- NEVER expose sequential IDs (use UUIDs).
- NEVER use verbs in REST resource URLs.
- NEVER return raw logs in production error responses.
- NEVER map 1:1 endpoints to tools if they serve a single user intent.

---

## References

- [12 MCP Patterns Behind Production Agents](https://substack.com/home/post/p-195426893)
- [`references/rest-best-practices.md`](references/rest-best-practices.md)
- [`references/api-security-guide.md`](references/api-security-guide.md)
