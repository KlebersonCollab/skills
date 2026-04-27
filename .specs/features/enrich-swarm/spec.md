# Spec: Enrichment of Swarm Facilitator with Governance Patterns

## 1. Goal
Integrate advanced swarm governance and synchronization patterns (mined from `src/utils/swarm`) into the `swarm-facilitator` skill. This will enable secure, resilient, and platform-native multi-agent orchestration.

## 2. Requirements (FR)
- **FR-1**: Implement the "Permission Sync" protocol. The Leader must propagate security guardrails to all Teammates.
- **FR-2**: Define standards for "Terminal Backends" (iTerm2, Tmux, In-Process).
- **FR-3**: Establish "Teammate Lifecycle" guidelines (Init, Heartbeat, Reconnection).
- **FR-4**: Add "Layout Management" patterns for multi-terminal agent visualization.
- **FR-5**: Integrate with `devsecops-expert` for shared dangerous pattern detection.

## 3. Acceptance Criteria (AC - BDD)

### Scenario 1: Teammate Initialization
**Given** a Leader agent starting a Swarm session
**When** a Teammate is spawned
**Then** the Leader must inject a "Teammate Addendum" prompt ensuring the sub-agent understands its role and security constraints.

### Scenario 2: Permission Propagation
**Given** a Leader with a blocked command list
**When** a Teammate tries to run a command
**Then** the Teammate must use the SAME blocked list provided by the Leader (Permission Bridge).

## 4. Constraints
- Maintain version 1.2.0 for the skill.
- Must focus on the technical synchronization between agents, not just communication.
