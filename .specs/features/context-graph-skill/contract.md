# Contract: Context Graph Skill Validation

## Validation Sensors

The following sensors must pass for the skill to be considered "Done":

1. **Existence Sensor**:
   - [ ] `SKILL.md` must exist and contain the section "Core Concepts (The Memory Triad)".
   - [ ] `examples/gold-standard-decision-ledger.md` must exist.
   - [ ] `resources/context-graph-architecture.mermaid` must exist.

2. **Quality Sensor**:
   - [ ] Every decision in `examples/` must include a "Rationale" field.
   - [ ] The "Prohibited" section in `SKILL.md` must explicitly forbid "Implicit Memory".

3. **Compliance Sensor**:
   - [ ] The skill must be registered in the root `README.md` with the correct version.
   - [ ] Total skill count in root `README.md` must be accurate (26).

## Exit Gate (Handoff)
- [ ] SDD status check via `hb sdd status` (if available).
- [ ] No placeholders in any artifact.
