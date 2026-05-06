# Spec: ADK Expert Skill

## Introduction
Create a new skill named `adk-expert` based on "The Agent Development Kit" (ADK) architecture. This skill will serve as a guide for building, managing, and orchestrating advanced agentic systems using the 5-layer model (CLAUDE.md, Skills, Hooks, Subagents, Plugins).

## Requirements (FR)
- **FR-1**: Define the 5 layers of the Agent Development Kit as core knowledge.
- **FR-2**: Provide instructions on how to implement each layer (Memory, Knowledge, Guardrail, Delegation, Distribution).
- **FR-3**: Include patterns for MCP Servers integration and Agent Teams orchestration.
- **FR-4**: Use Brazil Portuguese for explanations as requested by user.
- **FR-5**: Follow the standard SKILL.md format (frontmatter + sections).

## Acceptance Criteria (AC)
### AC-1: Skill Structure
- **Given** the repository root.
- **When** the implementation is finished.
- **Then** a directory `adk-expert/` must exist with a `SKILL.md` file.

### AC-2: Frontmatter Content
- **Given** the `adk-expert/SKILL.md` file.
- **When** inspected.
- **Then** it must contain `name`, `version`, `description`, and `category`.

### AC-3: Layer Content
- **Given** the visual reference from the image.
- **When** writing the skill.
- **Then** all 5 layers must be detailed with their respective responsibilities and examples.

### AC-4: Tone and Language
- **Given** the user's global mandate.
- **When** generating the content.
- **Then** the instructions and descriptions must be in Brazilian Portuguese.

## Constraints
- Must follow SDD standards.
- Must not violate existing architectural rules.
