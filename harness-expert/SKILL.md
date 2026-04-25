---
name: harness-expert
version: 3.0.0
description: "Technical engine for Harness Engineering. Implements agentic infrastructure patterns (Memory Tiering, Explore-Plan-Act, Context Compression) and adversarial feedback loops (GAN-style) for production-grade autonomy."
category: agentic-infrastructure
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Run **`hb harness rehydrate [feature]`** to instantly load state, memory, and learnings.
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)?
3. **Plan Check**: Does the `plan.md` file define the architecture, schemas, and include **Mermaid** diagrams?

---
# Harness Expert (v3.0.0)

> "If you are not the model, then you are the harness." — Master of the agentic substrate, ensuring deterministic state, tiered memory, and isolated execution.

---

## 🏗️ Agentic Harness Patterns

### Group 1: Memory & Context
1.  **Persistent Instruction File**: Use `.cursorrules` or `.agent` files for project-level conventions.
2.  **Scoped Context Assembly**: Dynamically load instructions based on the current directory (Monorepo support).
3.  **Tiered Memory**: Organize memory into **Index** (context), **Topic** (on-demand), and **Archive** (disk).
4.  **Dream Consolidation**: Periodic background pruning and deduplication of agent state.
5.  **Context Compression**: multi-stage summarization (Snip, Microcompact, Collapse) for long sessions.

### Group 2: Workflow & Orchestration
6.  **Explore-Plan-Act Loop**: Enforce a strict separation between understanding, designing, and editing.
7.  **Context-Isolated Subagents**: Use specialized agents for research vs. implementation to avoid context pollution.
8.  **Fork-Join Parallelism**: Spawn parallel agents in isolated git worktrees for independent tasks.

### Group 3: Tools & Permissions
9.  **Progressive Tool Expansion**: Start with a minimal toolset and activate specialized tools on demand.
10. **Command Risk Classification**: Deterministic pre-parsing and permission gating for shell commands.
11. **Single-Purpose Tool Design**: Use specialized tools (FileRead, FileEdit) instead of generic bash commands.

### Group 4: Automation
12. **Lifecycle Hooks**: Run automatic actions at `PreToolUse`, `PostToolUse`, `SessionStart`, etc.

---

## 🔄 GAN-style Feedback Loop (Quality Enforcement)

For high-complexity tasks (**Large/Complex**), the Harness must operate in **Adversarial** mode:

| Role | Responsibility | Tools |
|-------|------------------|-------------|
| **Generator** | Implement the feature according to the Spec and Plan. | HB CLI, SDD CLI, Frameworks. |
| **Evaluator** | Act as a relentless QA, testing and scoring. | **`hb harness evaluate`**, Playwright. |

### 📊 Evaluation Rubric (Target: Score >= 7.0)

| Criterion | Weight | Description |
|----------|------|-----------|
| **Design Quality** | 0.3 | Visual cohesion, absence of generic AI patterns, premium aesthetics. |
| **Originality** | 0.2 | Creative technical/visual decisions customized for the problem. |
| **Craft** | 0.3 | Polish, micro-interactions, typography, spacing, and performance. |
| **Functionality** | 0.2 | Robust functioning of all features and error handling. |

---

## Workflow (4 Phases)

### Phase 1: REHYDRATE & MAP — Context Loading
1.  **Context Injection**: Run **`hb harness rehydrate [feature]`**.
2.  **Tiered Loading**: Determine which "Memory Tier" needs to be in context.
3.  **Impact Mapping**: Run **`hb map --impact <file>`** to visualize technical dependencies.

### Phase 2: EXPLORE & PLAN — The Loop
1.  **Isolation**: Ensure sub-agents are spawned for heavy research.
2.  **Strategy**: Define the plan using **Lifecycle Hooks** for automated validations.
3.  **Tool Check**: Activate necessary tools via **Progressive Expansion**.

### Phase 3: ACT — Execution Automation
1.  **Atomic Implementation**: Write code focused on passing the ACs in `spec.md`.
2.  **Risk Mitigation**: Pass all commands through **Risk Classification**.
3.  **Self-Correction**: Immediately fix linter/test bugs via `hb audit`.

### Phase 4: CONSOLIDATE & SYNC — Closing
1.  **Dreaming**: Run background consolidation of `STATE.md` and `MEMORY.md`.
2.  **Strict Scoring**: Execute **`hb harness evaluate`**.
3.  **Final Sync**: Run **`hb sync`** to persist the new state.

---

## Quality Rules

- **Persistence First**: The File System is the only source of truth.
- **Context Awareness**: Use `Context Compression` to prevent window overflow.
- **Deterministic Check**: Every agent output must pass through hooks (linters/tests).
- **Zero Pollution**: Keep research logs out of implementation context.

## Prohibited

- NEVER commit directly to `main` without passing through the Act phase.
- NEVER maintain redundant context that can be summarized in `STATE.md`.
- NEVER bypass Risk Classification for destructive commands.

---

## References

- [12 Agentic Harness Patterns from Claude Code](https://generativeprogrammer.com/p/12-agentic-harness-patterns-from)
- [Claude Code Documentation](https://code.claude.com/)
- [Harness Engineering Principles](https://github.com/anthropics/harness)
