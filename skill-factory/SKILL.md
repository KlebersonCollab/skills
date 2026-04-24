---
name: skill-factory
version: 1.1.0
description: "Core Framework for standardized creation of new skills with automated scaffolding, validation, and registration."
category: skill-management
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Have you rehydrated the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define architecture, schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---
# Skill Factory: Core Framework

> Every skill deserves a solid foundation. Consistency breeds quality.

---

## Goal

Automate the standardized creation of new skills with guaranteed quality, eliminating human error in scaffolding and ensuring compliance with Skills Hub standards.

---

## Auto-Sizing

The depth of scaffolding is determined by the skill's **structure**:

| Mode | Scope | Generated Files | Sub-skill Used |
|------|--------|-------------------|---------------------|
| **Quick** | Simple skill, no sub-skills | `SKILL.md`, `README.md`, `CHANGELOG.md` | `skill-factory-bootstrap` |
| **Standard** | Skill with sub-skills | All above + `<skill>-<sub>.skill.md` (N) | `skill-factory-bootstrap` |

After scaffolding, validation is **always** executed (both modes):

| Phase | Responsible |
|-------|-------------|
| **Validate** | `skill-factory-validator` |

---

## Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `skill_name` | string | ✅ | Skill name (slug: lowercase, hyphenated, no spaces). Ex: `deep-research` |
| `description` | string | ✅ | Concise description of what the skill does. |
| `category` | string | ✅ | Functional category (free-form). Ex: `research`, `automation`, `development-workflow` |
| `sub_skills` | list[string] | ❌ | List of sub-skill names to be created. If empty, Quick mode. |

---

## The Modular Engine

This skill delegates tasks to specialized sub-skills:

- **[Bootstrap](skill-factory-bootstrap.skill.md)**: Generates complete scaffolding for the new skill (directory, files, standardized content).
- **[Validator](skill-factory-validator.skill.md)**: Audits the created skill against the compliance checklist and issues a report.

---

## Workflow (3 Phases)

### Phase 1: BOOTSTRAP
Execute `skill-factory-bootstrap` with the received parameters.
- Creates the `<skill_name>/` directory at the repository root.
- Generates all files from standardized templates.

### Phase 2: VALIDATE
Execute `skill-factory-validator` on the newly created skill.
- Verifies structure, frontmatter, content, and naming.
- If **NON-COMPLIANT**, correct before proceeding.

### Phase 3: REGISTER
Update the repository's root `README.md`:
1. Add the new skill to the **Available Skills** table.
2. Increment the skill count badge.

---

## Version Policy

Every new skill **ALWAYS** starts at version `1.0.0`. There is no support for pre-release versions.

---

## Output Structure

Execution of this skill results in the following mandatory artifacts for each new skill created:

| Artifact | Description |
|----------|-----------|
| `SKILL.md` | Technical definition and skill workflow. |
| `README.md` | Presentation and usage documentation. |
| `CHANGELOG.md` | Version history (starting at 1.0.0). |
| `references/` | (Optional) Detailed reference guides. |
| `examples/` | (Optional) Code samples and configurations. |

## Quality Rules

- **Template-First**: Every generated skill must follow the standardized templates defined in the `bootstrap` sub-skill.
- **Validation-Always**: No skill is considered created until the `validator` issues a `COMPLIANT` verdict.
- **Registry-Updated**: The root `README.md` is the source of truth for existing skills and must always be up-to-date.
- **Self-Consistent**: The name in the frontmatter must match the directory name.

## Prohibited

- NEVER create a skill without executing the validation phase.
- NEVER register a NON-COMPLIANT skill in the root README.
- NEVER use special characters, spaces, or uppercase in `skill_name`.
- NEVER create files outside the skill directory (except for the root README registration).
