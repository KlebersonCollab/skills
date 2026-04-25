# Specification: Context Graph Skill

## Goal
Implement a Meta-Knowledge Graph (Context Graph) skill to capture agentic decision traces and rationale, solving the "First Project Cost" and ensuring long-term institutional memory.

## User Stories

### Story 1: Decision Traceability
**As an** AI Agent,
**I want to** record the rationale behind every major architectural or logic choice,
**So that** future agents (or humans) can understand why the current system was built this way.

**Acceptance Criteria:**
- [ ] Must provide a template for `DECISIONS.md`.
- [ ] Each trace must include: ID, Date, Agent, Status, Context, Decision, Rationale, and Alternatives.
- [ ] Must allow updating traces with "Outcomes" (Success/Failure).

### Story 2: Pattern Distillation
**As an** AI Agent,
**I want to** identify recurring successful decisions across sessions,
**So that** I can promote them to "Standard Procedures" and improve accuracy.

**Acceptance Criteria:**
- [ ] Must provide a way to link individual decisions to general `PATTERNS.md`.
- [ ] Patterns must include "Applicable Domains" and "Success Rate".

### Story 3: Context Navigation
**As an** AI Agent,
**I want to** query past decisions before starting a new task,
**So that** I can reuse proven strategies and avoid known gotchas.

**Acceptance Criteria:**
- [ ] The skill must include a mandatory "Explore Phase" task to query the Context Graph.
- [ ] Must support visual representation via Mermaid Diagrams.

## BDD Scenarios

### Scenario: Recording a new Decision Trace
**Given** an agent is using the `context-graph` skill
**And** a critical architectural choice has been made
**When** the agent runs the "Capture" workflow
**Then** a new entry must be added to `DECISIONS.md` with full rationale
**And** the `CONTEXT-MAP.mermaid` should be updated to reflect the new node.

### Scenario: Reusing a Pattern
**Given** a project has an existing `PATTERNS.md` with a "Kebab-case" naming rule
**When** a new agent initializes a feature
**And** queries the Context Graph
**Then** the agent must adopt "Kebab-case" automatically based on the retrieved pattern.
