# Project Learnings: Skills Hub

## Technical Patterns
- **SDD v2.3.0 Integration**: Observed that modularizing skills into independent folders with their own `SKILL.md` allows for seamless agentic consumption.
- **Hybrid Relationship Mapping**: Combining global mandates (governance) with local content analysis (mentions) creates a high-fidelity visual representation of skill interdependencies.
- **Mandate Synchronization**: Using `bin/sync-skills.sh` to propagate `GLOBAL_MANDATES.md` to root `AGENTS.md` and hidden governance directories (`.gemini`, `.claude`) ensures a single source of truth for engineering mandates.

## Bug Fixes & Gotchas
- **Vite Path Resolution**: When building dashboards inside subdirectories, ensuring correct relative paths to the root `.agents` folder is critical for data extraction scripts.

## Process Insights
- **Manual Bootstrapping**: Discovered that manual creation of `.specs/` structure ensures the agent fully understands the governance layer before implementation.
- **Purist Transition**: Removing CLI dependencies simplifies the environment and reduces "magic" failures, forcing the agent to rely on explicit Markdown contracts which are more resilient and traceable.
- **Governance as Code (GaC)**: Automating metadata and evidence validation via `make audit` ensures 100% compliance and prevents governance drift.
- **Centralized Knowledge Mapping**: Maintaining a root `KNOWLEDGE-MAP.mermaid` provides a consistent dependency view across all project phases.
- **Automated Link Verification**: Using simple scratch scripts to verify the presence of `SKILL.md` and directory structures across a large number of modules is highly effective for maintaining repository integrity at scale.
- **Grilling Alignment Value**: Integrating the glossary (`CONTEXT.md`) and Architectural Decision Records (ADRs) with an interactive "sabatina" (grill-with-docs) drastically reduces semantic noise and design misalignment before the specification and implementation phases begin.
- **De-ambiguation via File Consolidation**: Observed that a large number of reference files (e.g. 24 files in a skill directory) increases token cost and can lead to semantic confusion. Fusing files of the same nature (like phase and session handoffs into a single `handoff-protocol.md`, and `quick-mode`/`context-limits` into `operational-guidelines.md`) clarifies rules and maximizes cognitive efficiency for LLMs.
- **Root Mandate Synchronization**: Elevating local SDD developments (such as context budget check zones, Document Purity, and Phase 0 Grilling) to global mandates ensures that all active sub-agents (Gemini, Claude, and Core) execute with optimized focus and context-awareness.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "GLOBAL-MANDATES-EVOLUTION"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-22T14:30:00Z"
evidence_checksum: "make-audit-success"
```
