# Spec: Architecture Skill v2.0 Upgrade

## 1. Goal
Upgrade the `architecture` skill to version 2.0, incorporating high-performance system design patterns such as **CQRS (Command Query Responsibility Segregation)**, **Event-Driven Architecture**, and **Evolutionary Architecture (Fitness Functions)**. This upgrade will allow the Gemini CLI to design complex, scalable, and resilient systems.

## 2. Target Audience
Software architects and senior developers using Gemini CLI to design distributed systems and high-scale architectures.

## 3. Core Requirements (Functional)
- **CQRS Blueprint**: Guidance for separating read and write models (Command vs. Query).
- **Event-Driven Patterns**: Design of asynchronous message flows, including Idempotency, Event Sourcing, and Eventual Consistency.
- **Fitness Functions**: Mechanism to define automated checks (unit tests or scripts) to validate architectural constraints (e.g., coupling, cycle detection).
- **Architecture Evolution**: Support for designing systems that grow and adapt over time (Evolutionary Architecture).

## 4. Acceptance Criteria (BDD Format)

### Scenario 1: Design a CQRS System
**Given** a requirement for a high-read/low-write application
**When** the `architecture` skill is invoked for design
**Then** it must suggest a CQRS pattern, defining separate models for commands and optimized queries.

### Scenario 2: Design Event-Driven Flows
**Given** a need for decoupling between services
**When** the `architecture` skill analyzes the system
**Then** it must provide a blueprint for message queues, defining event schemas and handling idempotency strategies.

### Scenario 3: Validate with Fitness Functions
**Given** a new architectural constraint (e.g., "Modules A and B must never depend on each other")
**When** the `architecture` skill is in the `Document` phase
**Then** it must suggest an automated "Fitness Function" (script/test) to verify this constraint in CI.

## 5. Technical Requirements (Non-Functional)
- **Backwards Compatibility**: Must not break existing ADR or System Map structures.
- **Reference Depth**: In-depth guides for CQRS, Events, and Evolutionary Architecture.
- **Example Richness**: Provide Mermaid diagrams and code snippets for the new patterns.

## 6. Architecture & Design
- `architecture/SKILL.md`: Main entry point (Update to v2.0.0).
- `architecture/references/`:
    - `cqrs-and-events.md`: Deep dive into CQRS and Event-Driven design.
    - `evolutionary-architecture.md`: Fitness functions and architectural growth.
- `architecture/examples/`:
    - `cqrs-pattern.mermaid`: Diagram for read/write model separation.
    - `event-flow.mermaid`: Diagram for asynchronous message flows.
    - `fitness-function-sample.py`: Example of an automated architectural check.
