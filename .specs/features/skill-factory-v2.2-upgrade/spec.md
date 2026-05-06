# Spec: Skill Factory v2.2 Upgrade (Root-Only)

## 1. Vision & Context
Upgrade the existing `skill-factory` (currently v2.0.0) located at the project root to v2.2.0. The goal is to align its structure and operational rigor with the `sdd` v2.2.0 master skill, while preserving the advanced "14 Anthropic Patterns" already present in the root version. 

**MANDATE**: Do NOT modify files in `.agents/` or other hidden system directories. All work must be performed on the visible root skill.

## 2. Acceptance Criteria (AC)
- [ ] **AC1: Pattern Synthesis**: Merge the 14 Anthropic Patterns (v2.0.0) with the SDD v2.2.0 layout (Prerequisites, Delegation Matrix, 4-Phase Workflow).
- [ ] **AC2: Structural Purity**: Ensure the skill is "Purist" - no magic CLI dependencies, all logic explicit in Markdown.
- [ ] **AC3: Standard Alignment**: Implement the **Knowledge Verification Chain** and **Safety Valve** protocols exactly as they appear in the `sdd` v2.2.0 skill.
- [ ] **AC4: Root Integrity**: All modifications must happen in the `skill-factory/` root directory.

## 3. Functional Requirements
- **FR1**: Update `SKILL.md` to v2.2.0 layout.
- **FR2**: Refactor `skill-factory-bootstrap.skill.md` and `skill-factory-validator.skill.md` for purist logic.
- **FR3**: Preserve existing references and examples in the root folders.

## 4. Constraints
- English for instructions, Brazilian Portuguese for user communication.
- No modifications to dot-folders (except `.specs/` for documentation).
