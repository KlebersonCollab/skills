# Changelog — SDD (Spec-Driven Development)

All notable changes to this skill will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and versioning follows [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [2.1.0] - 2026-04-27

### Added
- **Modular Orchestration**: Refactored `SKILL.md` to act as a central hub delegating tasks to sub-skills.
- **Safety Valve Protocol**: Mandatory re-planning trigger when implementation complexity drifts or touches high-risk files.
- **Knowledge Verification Chain**: Strict hierarchy for research to prevent hallucinations (Code > Internal Specs > Docs > MCP > Web).
- **Consolidated Mapping**: Integrated 7 fragmented codebase docs into a single, high-fidelity `TECHNICAL-MAP.md`.
- **English-Native Compliance**: All documentation and instructions migrated to English for maximum model performance.

## [1.5.0] - 2026-04-23

### Added
- **Spec-Driven TLC Integration**: Added 5 new advanced operational guidelines in `references/` (`brownfield-mapping.md`, `coding-principles.md`, `context-limits.md`, `session-handoff.md`, `quick-mode.md`).
- **Sub-Agent Delegation**: Documentation of the new delegation protocol for orchestrators to mitigate context window limits.
- **Audit Compliance (Skill Factory)**: Added mandatory `## Goal` and `## Quality Rules` in `sdd-implementer`, `sdd-orchestrator`, and `sdd-reviewer` sub-skills.

### Changed
- **Mass Version Update**: Sub-skills updated to `1.5.0` synchronizing with the base skill.

## [1.4.0] - 2026-04-16

### Added
- **Harness Engineering Integration**: Workflow evolution focusing on contracts, sensors, and deterministic feedback loops.
- **SDD Contracts**: New `contract.md` artifact to define the delivery agreement between Implementer and Reviewer.
- **Sensor-based Reviews**: Verdicts are now based on signals from external tools (linters, tests) with a Global Score (0-100).

### Changed
- **Version Unification**: Synchronization of all sub-skills to version 1.4.0 for ecosystem consistency.
- **Auto-Sizing Refinement**: Updated phase table to include Contract and Review by Score milestones.
- **Naming Correction**: Replaced residual term `ork-orchestrator` with `sdd-orchestrator`.

## [1.3.1] — 2026-04-14

### Added
- **Mermaid Diagrams Mandate**: Mandatory inclusion of Mermaid diagrams in the `Specify` phase for Medium+ levels.

## [1.3.0] — 2026-04-14

### Added
- **PRD/RFC/BDD Integration**: Incorporation of Product Requirements Documents, Request for Comments, and BDD (Given/When/Then) into the Auto-Sizing based workflow.

## [1.2.0] — 2026-04-14

### Added
- **High-Level Workflow**: Introduction of a 4-phase lifecycle (Discovery, Specify, Implement, Review) to guide the full SDD journey before delegation to sub-skills.

## [1.1.0] — 2026-04-14

### Added
- **Persistent Memory Protocol**: Integration of long-term memory with `MEMORY.md` and `LEARNINGS.md`.
- **sdd-planner Evolution**: Added templates and protocols for context rehydration and incremental learning capture.
- **Mandatory Structure Evolution**: Inclusion of new memory files in `.specs/project/`.

## [1.0.0] — 2026-04-14

### Added
- **SKILL.md**: Main definition of the SDD workflow with Auto-Sizing and mandatory directory structure.
- **sdd-explorer.skill.md**: Sub-skill for mapping existing codebases (Brownfield).
- **sdd-planner.skill.md**: Sub-skill for project vision management and persistent memory (`STATE.md`).
- **sdd-orchestrator.skill.md**: Sub-skill for technical specification with Knowledge Verification Chain.
- **sdd-implementer.skill.md**: Sub-skill for atomic and test-driven implementation.
- **sdd-reviewer.skill.md**: Sub-skill for auditing against acceptance criteria with evidence.
