# Spec: New Skill - CLI Expert

## 1. Goal
Create a new skill `cli-expert` that guides agents in designing and implementing AI-native Command Line Interfaces (CLIs). It should bridge the gap between traditional CLI design and tools optimized for LLM consumption.

## 2. Requirements (FR)
- **FR-1**: Define the "Slash Command" architecture (Prompt vs Local vs Bundled).
- **FR-2**: Provide guidelines for building interactive TUIs using React/Ink.
- **FR-3**: Establish metadata standards for tool registration (name, description, parameters, availability).
- **FR-4**: Include patterns for "Command Groups" (e.g., help, internal, user-facing).
- **FR-5**: Define context-aware command visibility (e.g., only show `/thinkback` if memory exists).

## 3. Acceptance Criteria (AC - BDD)

### Scenario 1: Designing a new Slash Command
**Given** that the agent is creating a new tool
**When** defining the command structure
**Then** it must include a human-readable description AND an LLM-optimized prompt hint.

### Scenario 2: TUI Design
**Given** a need for user feedback in the terminal
**When** designing the UI component
**Then** the skill must recommend using "Stateful UI" (like Ink) to avoid messy log streams.

## 4. Constraints
- Must follow the `skill-factory` structure.
- Must use the `src/commands.ts` patterns as the technical Gold Standard.
