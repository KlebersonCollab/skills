# Project Memory: Skills Hub

## Project Identity
- **Name**: AI Agent Skills Hub
- **Purpose**: A centralized repository of modular, specialized skills for AI agents, governed by SDD v2.3.0.
- **Core Philosophy**: Logic-First, Markdown-based governance, and Purist SDD (Spec-Driven Development).

## Context Facts
- The project is transitioning to a "Purist" SDD model (SDD v2.3.0).
- Governance as Code (GaC) implemented via `make audit`.
- Centralized Knowledge Hub established in the project root.
- Multiple specialized skills exist (e.g., `python-uv`, `sdd`, `architecture`).
- Language standard for external documentation: English.
- Language standard for communication with user: Brazilian Portuguese (per user global rule).

## Architectural Decisions
- [ADR-001] Transition to Purist SDD: Removing CLI dependencies in favor of Markdown-based governance artifacts.
- [ADR-002] Centralized Memory: All project state and context must reside in `.specs/project/`.
- [ADR-003] Skills Visualization: Implementation of a React-based Dashboard and Graph to visualize the ecosystem (hub-ui-skills).
- [ADR-004] Governance as Code: Mandatory automated auditing of SDD metadata and evidence via `make audit`.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-08T19:32:00Z"
evidence_checksum: "b134813"
```

