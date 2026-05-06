# Plan: Token Distiller v2.3.0 Upgrade

## 1. Technical Strategy
The upgrade involves surgical edits to Markdown files to inject metadata and update version numbers. We will also clean up redundant files.

## 2. Implementation Sequence

### Phase 1: Core Skill Upgrade
- Modify `token-distiller/SKILL.md`:
  - Update version in YAML frontmatter.
  - Update version in H1 header.
  - Add `@sdd-state` block.
  - Ensure references to Triad of Memory use absolute project paths (or relative to root).

### Phase 2: Documentation Audit
- Update `token-distiller/README.md`:
  - Add `@sdd-state` block.
- Update `token-distiller/CHANGELOG.md`:
  - Add v2.3.0 entry.
  - Add `@sdd-state` block.

### Phase 3: Cleanup
- Delete `token-distiller/tasks.md`.

## 3. Risk & Safety
- **Risk**: Breaking automated skill consumption.
- **Mitigation**: Ensure YAML frontmatter remains valid. Use `make audit` to verify structure.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "TOKEN-DISTILLER-V2.3.0"
phase: "IMPLEMENT"
status: "COMPLETED"
last_update: "2026-05-06T12:56:00Z"
evidence_checksum: "NONE"
```
