---
name: context-graph
version: 1.0.0
description: "Meta-Knowledge and Decision Intelligence manager. Implements a Context Graph layer to capture the 'how' and 'why' of agentic decisions, enabling institutional memory and self-improving workflows."
category: intelligence-governance
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Did you rehydrate the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Architecture Link**: If making architectural choices, ensure compliance with existing **ADRs** (via `architecture` skill).
3. **Knowledge Architect Link**: When recording relations, cross-reference with the **Knowledge Map** (`knowledge-architect` skill).
4. **Spec Check**: Does the `spec.md` file exist?
4. **Plan Check**: Does the `plan.md` file include meta-knowledge requirements?

---

# Context Graph (Meta-Knowledge Manager)

> "The domain graph captures what exists; the context graph captures how you built it and why." — Protocol for agentic memory and decision traceability.

---

## Goal

The goal of this skill is to maintain a **Context Graph** (Meta-Graph) that acts as the agent's long-term memory. It captures the rationale behind architectural decisions, extraction patterns, and tribal knowledge that would otherwise be lost between sessions.

---

## Core Concepts (The Memory Triad)

1.  **Episodic Memory (Decision Traces)**: Specific events. "What was decided, why, and what was the outcome?"
2.  **Semantic Memory (Patterns & Tribal Knowledge)**: Generalized facts. "In this domain, 'X' always means 'Y'."
3.  **Procedural Memory (Proven Strategies)**: Learned behaviors. "When encountering document type 'A', use schema 'B' and resolution strategy 'C'."

---

## Workflow (4 Phases)

### Phase 1: CAPTURE — Recording Decisions
1.  **Decision Logging**: Every time a critical choice is made (schema design, logic branch, tool selection), record a **Decision Trace**.
2.  **Rationale Extraction**: Explicitly state the "Why". What alternatives were considered? What constraints (technical/business) forced this choice?
3.  **Provenance**: Attach the 'Who' (user/agent) and 'When' (timestamp/session ID).

### Phase 2: DISTILL — Pattern Recognition
1.  **Pattern Identification**: Analyze multiple decisions to find recurring successes or failures.
2.  **Tribal Knowledge Mapping**: Capture "soft" requirements or industry-specific nuances that aren't in formal docs.
3.  **Schema Evolution**: Track how decisions change over time as the project matures.

### Phase 3: RETRIEVE — Precedent Navigation
1.  **Similarity Search**: Before making a new decision, query the Context Graph for similar past situations.
2.  **Precedent Application**: "Last time we had this conflict, we chose 'X' because of 'Y'. Does 'Y' still apply?"
3.  **Guided Onboarding**: Use traces to explain the current architecture to new agents or humans.

### Phase 4: EVOLVE — Self-Improvement
1.  **Outcome Feedback**: Link decisions to their eventual outcomes (Success/Failure).
2.  **Pattern Optimization**: Promote successful decisions to "Standard Procedures" and flag failed ones as "Known Gotchas".
3.  **Graph Refactoring**: Clean up stale decisions that no longer apply to the current schema.

---

## Output Structure

Execution of this skill results in the following artifacts (usually in `.specs/context/` or the skill root):

| Artifact | Format | Description |
|----------|---------|-----------|
| **Decision Ledger** | `DECISIONS.md` | Sequential log of all architectural and logic choices with rationale. |
| **Context Map** | `CONTEXT-MAP.mermaid` | Visual graph connecting Decisions, Patterns, and Tribal Knowledge. |
| **Meta-Patterns** | `PATTERNS.md` | Reusable procedural strategies for the agent. |

---

## Quality Rules

- **Rationale over Choice**: A decision without a "Why" is a bug. Always document the reasoning.
- **Outcome Linking**: Decisions must be eventually tagged with an outcome (did it work?).
- **Layered Specificity**: Start from Project-specific rules, fall back to Org-rules, then Industry-standards.
- **Traceable Provenance**: Every pattern must point back to the evidence (decision traces) that generated it.

## Prohibited

- NEVER ignore user corrections; they are the highest-priority "Learning Signals" for the Context Graph.
- NEVER store sensitive PII in decision traces.
- NEVER use implicit memory (hoping the agent "just remembers"); everything must be externalized in the graph.
- NEVER allow "Fresh Start" syndrome; every task must begin by querying the Context Graph for precedents.

---

## References

- [Meta Knowledge Graphs (Context Graphs) - Firat Tekiner](https://firattekiner.substack.com/p/meta-knowledge-graphscontext-graphs)
- [Generative Agents: Interactive Simulacra of Human Behavior](https://arxiv.org/abs/2304.03442)
