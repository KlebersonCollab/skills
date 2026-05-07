# Skill Validation Checklist: COMPLIANT Status

A skill is considered 'COMPLIANT' when it meets all the standards defined in this checklist. This ensures consistency and reliability across the Hub.

## 1. Structure & Organization
- [ ] Folder contains `SKILL.md` (Core instructions).
- [ ] Folder contains `README.md` (Usage documentation for humans).
- [ ] Folder contains `CHANGELOG.md` (Version tracking).
- [ ] Sub-directories exist: `examples/`, `references/`, `resources/`.
- [ ] No empty `.gitkeep` files remain (unless the folder is intended to be empty).

## 2. Technical Quality (SKILL.md)
- [ ] Role definition is clear and focused (One skill, one responsibility).
- [ ] Instructions use structured Markdown (Goal, Output, Rules, Prohibited).
- [ ] Quality rules follow the project's engineering mandates (SDD, KISS, etc.).
- [ ] 'Prohibited' section includes security and safety constraints.
- [ ] Mermaid diagrams are used to visualize complex flows or architectures.

## 3. Idiomatic consistency
- [ ] Prompts are technical and professional in tone.
- [ ] Terminology is consistent with the Hub (e.g., 'Skill', 'Sub-skill', 'Agent', 'Hub').
- [ ] Language is appropriate for the target AI model (Clear, unambiguous, and directive).

## 4. Validation & Testing
- [ ] At least one `example/` scenario is provided and functional.
- [ ] A `validation-report.md` has been generated for the current version.
- [ ] All listed resources in `SKILL.md` actually exist in the folder.

## 5. Metadata & Registry
- [ ] Skill is registered in the main `onboarding-navigator/references/skills-catalog.md`.
- [ ] Version number is updated and consistent across all files.
- [ ] Licensing information is present if applicable.
