# Specification: Stitch Skills Integration

## 1. Overview
The goal of this feature is to integrate a newly acquired set of "Stitch Skills" (currently residing in `stitch-skills-main/`) into the main Skills Hub architecture (`/Users/klebersonromero/Projetos/Kleberson/skills`). The integration must follow the SDD (Spec-Driven Development) standards, ensuring that skills are either aggregated into existing logical boundaries or separated into new, single-responsibility skills.

## 2. Source Skills Analysis
The downloaded repository contains the following skills:
- `design-md`, `taste-design`, `enhance-prompt`, `stitch-design` (Stitch UI/Prompting)
- `react-components` (React/Vite Generation)
- `shadcn-ui` (Shadcn/UI integration)
- `remotion` (Video generation)
- `stitch-loop` (Autonomous site building loop)

## 3. Integration Strategy Requirements
- **FR-1: Consolidation of Stitch UI Skills.** `design-md`, `taste-design`, `enhance-prompt`, and `stitch-design` must be aggregated into a single `stitch-expert` skill to avoid fragmentation.
- **FR-2: Frontend Aggregation.** `react-components` and `shadcn-ui` must be integrated into the existing `frontend-expert` skill, as they share the same domain (React, styling, component architecture).
- **FR-3: Standalone Extraction.** `remotion` and `stitch-loop` must be created as new, standalone skills (`remotion-expert` and `stitch-loop-expert`), given their distinct operational boundaries (video generation and autonomous looping, respectively).

## 4. Acceptance Criteria (BDD)

### Scenario 1: Aggregation of Stitch UI Skills
**Given** the source skills `design-md`, `taste-design`, `enhance-prompt`, and `stitch-design`
**When** the integration is executed
**Then** a new skill `stitch-expert` is created in the Hub
**And** it consolidates the prompt enhancement, design system synthesis, and UI guidelines from the source skills into a single cohesive SKILL.md.

### Scenario 2: Aggregation of Frontend Skills
**Given** the existing `frontend-expert` skill and the source skills `react-components` and `shadcn-ui`
**When** the integration is executed
**Then** the `frontend-expert` skill is updated to include shadcn/ui and Stitch-to-React component workflows.

### Scenario 3: Standalone Skills Creation
**Given** the source skills `remotion` and `stitch-loop`
**When** the integration is executed
**Then** new skills `remotion-expert` and `stitch-loop` are created in the Hub
**And** their respective workflows and resources are ported over to the standard Hub structure.
