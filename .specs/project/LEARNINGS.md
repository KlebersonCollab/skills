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

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-22T14:19:00Z"
evidence_checksum: "make-audit-success"
```
