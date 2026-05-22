# Specification: Token Distiller Advanced Caveman Mode

## 1. Context & Objective
The `token-distiller` skill needs an upgrade to support advanced compression levels and safety mechanisms. This feature will introduce three levels of "Caveman Mode" intensity (`lite`, `full`, `ultra`), strict pattern formatting, and Auto-Clarity safety triggers, while explicitly leaving out any Classical Chinese (Wenyan) modes.

## 2. Requirements (Traceable)
- **FR-1**: Define three intensity levels in `SKILL.md`: `lite`, `full` (default), and `ultra`.
- **FR-2**: Define exact compression rules for each level:
  - `lite`: No filler, keep articles + full sentences, professional but tight.
  - `full`: Drop articles, use fragments, short synonyms.
  - `ultra`: Abbreviate prose words, strip conjunctions, use arrows for causality (`X → Y`).
- **FR-3**: Implement **Auto-Clarity triggers** in the specification and example script (suspend compression for security warnings, destructive confirmations, complex multi-step instructions, or extreme ambiguity).
- **FR-4**: Restructure the response pattern requirement: `[thing] [action] [reason]. [next step].`
- **FR-5**: Implement command options for `/caveman [lite|full|ultra]` and `/mode low [lite|full|ultra]`.
- **FR-6**: Refactor `examples/distill.py` to support all three compression levels and demonstrate the Auto-Clarity rules.

## 3. Acceptance Criteria (AC)
- [ ] `SKILL.md` documents the `lite`, `full`, and `ultra` intensity levels, Auto-Clarity rules, and the strict pattern.
- [ ] `examples/distill.py` supports `--level [lite|full|ultra]` and demonstrates correct output transformations and Auto-Clarity bypasses.
- [ ] `CHANGELOG.md` reflects version `2.4.0` with these additions.
- [ ] All `.md` files contain the correct `<!-- @sdd-state -->` governance block.
- [ ] Project-wide `make audit` passes successfully with 100% compliance.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "TOKEN-DISTILLER-ADVANCED-CAVEMAN"
phase: "IMPLEMENT"
status: "IN_PROGRESS"
last_update: "2026-05-22T14:02:46Z"
evidence_checksum: "NONE"
```
