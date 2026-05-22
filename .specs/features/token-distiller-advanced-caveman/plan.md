# Plan: Token Distiller Advanced Caveman Mode

## 1. Technical Strategy
This upgrade adds multi-tier linguistic compression levels and Auto-Clarity safety protocols to `token-distiller`. We will modify the skill specification, update the changelog, and write a robust Python implementation in the example script that mimics these compression layers.

## 2. Implementation Sequence

### Phase 1: Specifications Update
- Update `token-distiller/SKILL.md`:
  - Increase version to `2.4.0` in frontmatter and headers.
  - Document `lite`, `full`, and `ultra` modes in detail.
  - Document the strict structural pattern: `[thing] [action] [reason]. [next step].`
  - Define exact Auto-Clarity boundaries and trigger rules.
- Update `token-distiller/README.md`:
  - Update instructions to mention the intensity levels.
- Update `token-distiller/CHANGELOG.md`:
  - Add version `2.4.0` entry highlighting the advanced levels and safety triggers.

### Phase 2: Python Script Evolution
- Modify `token-distiller/examples/distill.py`:
  - Add `--level [lite|full|ultra]` flag using `argparse`.
  - Refactor `to_caveman(text, level)` to apply different rules based on the level.
  - Implement a simple parser/heuristic for Auto-Clarity (e.g. if the input matches patterns for warnings, list of steps, or destructive words, output the raw message as per safety protocols).
  - Structure the output to follow the `[thing] [action] [reason]. [next step].` pattern where applicable.

### Phase 3: Audit & Sync
- Run `make audit` to ensure governance files are perfectly formatted and compliant.
- Run `make sync` to propagate the skill files to the `.agents` and `.gemini` folders.

## 3. Risk & Safety
- **Risk**: Over-compressing critical instructions or warnings leading to user error.
- **Mitigation**: Strictly define and implement Auto-Clarity triggers to bypass compression for all warnings and irreversible actions.

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
