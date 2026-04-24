# Changelog — Harness Expert

All notable changes to this skill will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and versioning follows [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---

## [2.1.0] - 2026-04-22

### Added
- **GAN-style Feedback Loop**: Implementation of the adversarial cycle between Generator and Evaluator for Large/Complex tasks.
- **Evaluation Rubric**: New scoring matrix (Design, Craft, Originality, Functionality) with a target Score >= 7.0.

## [2.0.0] - 2026-04-20

### Changed
- **Role Refactoring**: The skill now focuses exclusively on the **technical engine** (agentic infrastructure).
- **Decoupling**: Strategic memory management removed, delegating the operational vision protocol to the `sdd-planner`.
- **Workflow**: Updated to include `AUTOMATED REHYDRATE`, `OPERATE`, `AUTO-SYNC`, and `VALIDATE` phases.

## [1.1.0] - 2026-04-18

### Changed
- **Global Renaming**: Official transition from "Hardness" to **Harness Engineering** to align with industry standards.
- **Enhanced Rehydration**: Added **Bootstrap** step for operational environment validation and automatic dependency installation.

### Added
- **Harness Principles**: Documentation of Feed Forward, Feedback, and Sensors in the README.

## [1.0.0] — 2026-04-15

### Added
- **SKILL.md**: Main definition of the Harness Engineering skill.
- **README.md**: Detailed documentation and overview.
- **harness-expert-sync.skill.md**: Sub-skill for automated state synchronization.
- **harness-expert-rehydrate.skill.md**: Sub-skill for operational context injection.
- **harness-expert-compress.skill.md**: Sub-skill for context compression (Context Compressor).
- **Resources**: Implemented `compressor.py` for task analysis and compression.
- **Examples**: Added synchronization demonstration script.
