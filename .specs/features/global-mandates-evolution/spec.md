# Spec - Global Mandates Evolution

## Goal
Evolve the repository's Global Engineering Mandates in `AGENTS.md` to incorporate learnings and conventions from SDD consolidation: Phase 0 Alignment (Grilling/Glossary), Document Purity (Anti-Fragmentation), and strict Context Management.

## User Review Required
We are modifying `AGENTS.md` at the root, which acts as the central engineering mandates for all agents. The new rules cover:
1. **Phase 0: ALIGN** (Relentless Grilling and Glossaries).
2. **Document Purity Principle** (Anti-fragmentation to optimize context window).
3. **Context Allocation Zones** (Keeping prompts under 40k tokens).

## Acceptance Criteria
### AC-1: Mandate 1 - Phase 0 Alignment
- **Given** the `AGENTS.md` file.
- **When** the new mandates are added.
- **Then** Section 1 (SDD Framework) explicitly requires agents to conduct a "Grilling Session" (Phase 0: ALIGN) and create/update `CONTEXT.md` (Glossary) before specifying requirements.

### AC-2: Mandate 2 - Document Purity and Anti-Fragmentation
- **Given** the `AGENTS.md` file.
- **When** the new mandates are added.
- **Then** Section 3 (Quality Standards) includes a mandate on Document Purity, prohibiting document fragmentation and enforcing unifications like `handoff-protocol.md` and `operational-guidelines.md`.

### AC-3: Mandate 3 - Context Budget Management
- **Given** the `AGENTS.md` file.
- **When** the new mandates are added.
- **Then** Section 9 (Observable Governance) or Section 0 (Bootstrap) includes strict context limits (Healthy <40k, Moderate 40-60k, Critical >60k tokens) to preserve the LLM's logical reasoning.

### AC-4: Verification and Audit Success
- **Given** the modified `AGENTS.md` file.
- **When** `make sync` and `make audit` are executed.
- **Then** the global mandates are successfully synced to all sub-agent and hidden directories (`.agents`, `.gemini`, `.claude`), and the entire workspace is 100% compliant.

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
