# Spec: Skill Factory Reconstruction (v1.0.0)

## 1. Vision & Context
The `skill-factory` is the governance and automation engine for the Hub's skills. It must ensure that any new skill added to the ecosystem follows the "Gold Standard" (SDD v2.2.0) and provides clarity for both the agent and the developer. Currently, the skill contains empty files, failing to provide any guidance or consistency.

## 2. Acceptance Criteria (AC)
- [ ] **AC1: Standard Alignment**: The `SKILL.md` must implement the SDD v2.2.0 pattern (Frontmatter, Mandatory Prerequisites, Knowledge Verification Chain).
- [ ] **AC2: Logic-First Scaffolding**: Must provide a conceptual blueprint of a skill. The agent must understand the *why* behind every file, enabling it to build the structure manually with precision.
- [ ] **AC3: Protocol-Driven Validation**: Validation must be a mental and structural check of constraints and definitions, not just a binary CLI check.
- [ ] **AC4: Sub-skill Integration**: Must include `skill-factory-bootstrap` and `skill-factory-validator` sub-skills.

## 3. Functional Requirements
- **FR1**: Define the architectural blueprint for a "Gold Standard" skill.
- **FR2**: Implement the "Knowledge Verification Chain" protocol as the primary reasoning loop.
- **FR3**: Define the "Safety Valve" mechanism for skill execution.

## 4. Constraints
- Must be written in English (Instructional content).
- Must follow the directory structure defined in SDD v2.2.0.
