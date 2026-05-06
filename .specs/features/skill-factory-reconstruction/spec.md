# Spec: Skill Factory Reconstruction (v1.0.0)

## 1. Vision & Context
The `skill-factory` is the governance and automation engine for the Hub's skills. It must ensure that any new skill added to the ecosystem follows the "Gold Standard" (SDD v2.2.0) and provides clarity for both the agent and the developer. Currently, the skill contains empty files, failing to provide any guidance or consistency.

## 2. Acceptance Criteria (AC)
- [ ] **AC1: Standard Alignment**: The `SKILL.md` must implement the SDD v2.2.0 pattern (Frontmatter, Mandatory Prerequisites, Knowledge Verification Chain).
- [ ] **AC2: Functional Scaffolding**: Must provide clear instructions on how to bootstrap a new skill directory with all required memory artifacts (`STATE.md`, `MEMORY.md`, `LEARNINGS.md`, `DECISIONS.md`).
- [ ] **AC3: Validation Logic**: Must define how to audit a skill for "Slop-Free" design and structural integrity (Law of SDD).
- [ ] **AC4: Sub-skill Integration**: Must include `skill-factory-bootstrap` and `skill-factory-validator` sub-skills.

## 3. Functional Requirements
- **FR1**: Define the mandatory structure for a new skill (Folders and Files).
- **FR2**: Implement the "Knowledge Verification Chain" protocol as defined in the master SDD skill.
- **FR3**: Provide a `hb skill new <name>` pattern instruction for the agent.

## 4. Constraints
- Must be written in English (Instructional content).
- Must follow the directory structure defined in SDD v2.2.0.
