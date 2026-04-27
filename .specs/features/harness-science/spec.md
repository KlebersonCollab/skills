# Spec: New Skill - Harness Science

## 1. Goal
Create a new skill `harness-science` that introduces the discipline of **Ablation Science** and **Operational Benchmarking** for AI agents. It aims to provide a rigorous framework for measuring how specific capabilities (e.g., background tasks, tool access, "thinking" steps) contribute to the agent's overall success rate.

## 2. Requirements (FR)
- **FR-1**: Define "Ablation Testing" as a core methodology.
- **FR-2**: List standard ablation flags (e.g., `DISABLE_BACKGROUND_TASKS`, `DISABLE_THINKING`, `SIMPLE_MODE`).
- **FR-3**: Establish guidelines for creating "Operational Baselines" (Harness L0).
- **FR-4**: Provide a protocol for "Counter-factual Testing" (e.g., "Would the agent have succeeded without Tool X?").
- **FR-5**: Integrate with the `benchmark-expert` skill for quantitative data.

## 3. Acceptance Criteria (AC - BDD)

### Scenario 1: Conducting an Ablation Experiment
**Given** an agent with multiple active features
**When** a "Thinking Ablation" is performed
**Then** the system must be run with the feature disabled
**And** the output must be compared against the baseline to measure "Reasoning Degradation".

### Scenario 2: Establishing a Baseline
**Given** a new version of the Skills Hub
**When** setting the Harness L0 baseline
**Then** all non-essential experimental features must be disabled to establish a deterministic performance floor.

## 4. Constraints
- Must follow the `skill-factory` structure.
- Must use the `src/entrypoints/cli.tsx` ablation logic as the technical foundation.
