# Plan - Global Mandates Evolution

## Proposed Changes

We will edit the central `/home/kleberson/Documentos/skills/AGENTS.md` and then propagate it to the sub-agent and hidden runtime directories using `make sync`.

### Specific modifications in `/home/kleberson/Documentos/skills/AGENTS.md`:

1. **Section 0 (SESSION BOOTSTRAP)**:
   - Add a requirement to respect Context Budget Zones (<40k tokens target).
2. **Section 1 (SDD Framework)**:
   - Formally add **Phase 0: ALIGN** as a mandatory step. Explain that the agent must relentlessly challenge the proposed plans against the existing domain model (Grilling Session) and map precise domain terminology inside `CONTEXT.md` before initiating specification in `spec.md`.
3. **Section 3 (Quality Standards)**:
   - Add **Document Purity Principle**: Mandate against fragmenting operational guides or files. Force the unification of similar operational assets (e.g. Phase and Session Handoffs, or Quick-Mode and Context limits) into high-fidelity files (`handoff-protocol.md`, `operational-guidelines.md`) to minimize token bloat and cognitive friction.
4. **Metadata block update**:
   - Change `feature_id` to `GLOBAL-MANDATES-EVOLUTION`.
   - Update phase and status.

### Flow diagram of synchronization (Mermaid)

```mermaid
graph TD
    AM[AGENTS.md root] ➡️ MS{make sync}
    MS ➡️ GA[.gemini/rules/agent.md]
    MS ➡️ CA[.claude/CLAUDE.md]
    MS ➡️ AA[.agents/rules/agent.md]
    MS ➡️ AM2[AGENTS.md espelhos]
```

## Verification Plan

### Automated Tests
- Run `make sync` to propagate `AGENTS.md` globally to all ecosystem nodes.
- Run `make audit` to verify metadata blocks, evidence checksums, and compliance of all `.specs` and files.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "GLOBAL-MANDATES-EVOLUTION"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-22T14:30:00Z"
evidence_checksum: "ffb546e"
```
