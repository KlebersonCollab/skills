# Doctor Diagnostic Checklist

This checklist defines the "hb doctor" protocol for the Skills Hub.

## 1. Environment & Infrastructure
- [ ] **Shell Permissions**: Can the agent execute `run_command` in the workspace?
- [ ] **Config Dirs**: Are `.specs/`, `.agents/`, and `.hub-mode` present?
- [ ] **Git Integrity**: Is the project a git repository? Are submodules (if any) initialized?
- [ ] **Tooling**: Is the `hb` CLI accessible and updated?

## 2. Skill Registry Integrity
- [ ] **SKILL.md Headers**: All skills must have YAML frontmatter (name, version, description).
- [ ] **Broken Links**: Check for `resource:` or `file:///` links that point to non-existent files.
- [ ] **Mandatory Sections**: Do all core skills have a `🔒 Prerequisites` section?
- [ ] **Shadow Skills**: Search for duplicate skills in different directories (e.g., `temp-skills/` vs `.agents/skills/`).

## 3. SDD Compliance
- [ ] **State Sync**: Does `STATE.md` reflect the current active branch?
- [ ] **Task Atomicity**: Do all `tasks.md` files have clear [ ] or [x] states?
- [ ] **Mermaid Syntax**: Verify if Mermaid diagrams in `plan.md` are valid (no broken syntax).

## 4. Knowledge Map
- [ ] **LKG Sync**: Is `KNOWLEDGE-MAP.mermaid` up to date with the latest skill list?
- [ ] **Memory Pruning**: Is `MEMORY.md` becoming too large? (> 50KB suggests need for distillation).
