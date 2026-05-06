# Specification: Consolidate Swarm Facilitator into SDD

## 🎯 Objective
Eliminate governance redundancy by merging the infrastructure and orchestration logic from `swarm-facilitator` into the core `sdd` skill, establishing SDD as the absolute "Source of Truth" for both workflow and execution.

## 📋 Requirements
- **FR-1: Logic Migration**: Move Terminal Management, Permission Bridge, and Layout Management from `swarm-facilitator` to `sdd/references/swarm-execution.md`.
- **FR-2: Persona Mapping**: Explicitly map SDD sub-agents to the personas (Architect, Developer, Reviewer) and their recommended skills.
- **FR-3: Deprecation**: Delete the `swarm-facilitator` skill directory to remove redundancy.
- **FR-4: SDD Update**: Reference the new `swarm-execution.md` in `sdd/SKILL.md`.

## ✅ Acceptance Criteria (BDD)
- **AC-1: No Redundancy**
  - **Given** the skill catalog
  - **When** I search for "handoff" or "swarm"
  - **Then** only the `sdd` skill should contain the definitive protocol.

- **AC-2: Feature Parity**
  - **Given** an agent needing Tmux or Security Sync
  - **When** it consults `sdd/references/swarm-execution.md`
  - **Then** it must find the exact instructions previously in `swarm-facilitator`.

- **AC-3: Clean Registry**
  - **Given** the root `README.md` and catalog
  - **When** I check the list of skills
  - **Then** `swarm-facilitator` must be removed or marked as deprecated/merged.
