# Changelog — Skill Factory

All notable changes to this skill will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/), and versioning follows [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

---
 
 ## [2.2.0] - 2026-05-06
 
 ### Added
 - **Purist Mandate**: Enforced logic-first instructions (Markdown protocols) over hidden CLI dependencies.
 - **Memory Centralization**: Integrated pointers to `.specs/project/` global memory, prohibiting redundant local files.
 - **Gold Standard Alignment**: Upgraded to SDD v2.2.0 operational rigor, including Prerequisites Locks and Knowledge Verification Chains.
 - **Root Integrity**: Formalized development at the project root, prohibiting modifications to system-level `.agents/` directories.
 
 ### Changed
 - **SKILL.md**: Refactored for v2.2.0 with the 14 Anthropic pattern synthesis.
 - **skill-factory-bootstrap**: Refactored as a conceptual blueprint agent.
 - **skill-factory-validator**: Refactored to audit memory centralization and structural purity.
 

## [1.1.0] - 2026-04-16

### Added
- **skill-factory-validator**: Added Check 6 (Registry Status) and Check 7 (Onboarding Sync).
- **GitHub Actions**: CI/CD integration for automatic compliance validation.

## [1.0.0] — 2026-04-14

### Added
- **SKILL.md**: Main definition of the Core Framework with Auto-Sizing (Quick/Standard), 3-phase workflow, and Version Policy.
- **skill-factory-bootstrap.skill.md**: Sub-skill for automated scaffolding with standardized templates for `SKILL.md`, `README.md`, `CHANGELOG.md`, and sub-skills.
- **skill-factory-validator.skill.md**: Sub-skill for compliance validation with 5 checks (Structural, Frontmatter, Content, Naming, Registry) and evidence report.
- **README.md**: Detailed documentation with usage examples for both modes (Quick and Standard).
