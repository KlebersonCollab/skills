# Project Decision Ledger (Context Graph)

> Centralized repository of architectural rationale, meta-knowledge, and strategic decisions for the AI Agent Skills Hub.

---

## [DT-2026-003] Reinforcement of SDD Mandates (Mandatory PVE)

**Date**: 2026-04-25
**Agent**: Antigravity
**Status**: `Active`

### Context
During the implementation of the `context-graph` skill, the agent (me) skipped the formal SDD planning phases (Explore/Plan) and went straight to implementation. This resulted in missing validation artifacts and empty resource folders.

### Decision
Harden the SDD mandates to prohibit "Act-Only" workflows for any task involving new intelligence (Skills) or architectural changes. Establish a **Zero-Tolerance Policy** for missing `spec.md` and `plan.md` in Medium+ tasks.

### Rationale
- **Traceability**: Without the Plan phase, the "Why" behind decisions is lost.
- **Visual-First**: Skipping planning often leads to neglected visual assets (`resources/`).
- **Consistency**: The agent must be the first to follow the rules it enforces on others.

### Alternatives Considered
1. **Allow Quick-Mode for Skills**: Rejected. Skills are "Core Intelligence" and require maximum rigor.
2. **Automated Linter**: Ideal, but requires immediate manual enforcement first.

### Tribal Knowledge Applied
> "SDD is not a suggestion; it is the skeleton that supports the Hub's intelligence."

### Expected Outcome
100% compliance in future skill enrichments and new feature developments.

---

## [DT-2026-004] Implementation of the Context Graph Skill

**Date**: 2026-04-25
**Agent**: Antigravity
**Status**: `Completed`

### Context
A gap was identified: we map the domain (Knowledge Architect) but not the *decision process* (Meta-Knowledge).

### Decision
Implement a `context-graph` skill focused on Episodic, Semantic, and Procedural memory layers.

### Rationale
Aligns the repository with state-of-the-art agentic memory research (Tekiner, 2026).

### Expected Outcome
Agents that learn from their own successes and failures across sessions.
