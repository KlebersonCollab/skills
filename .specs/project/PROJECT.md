# AI Agent Skills Hub - Project Specification

## 1. Vision & Purpose
The **AI Agent Skills Hub** is a centralized ecosystem designed to serve the demands of modern developers by providing a standardized, modular, and high-performance repository of "Skills" for AI agents. 

Our core mission is to maintain a state-of-the-art **Agentic Development Standard**, ensuring that every capability integrated into an agent is deterministic, verifiable, and follows the **Purist SDD (Spec-Driven Development)** methodology.

## 2. Core Governance: The Law of SDD
This project is governed by the **Law of SDD**:
> "If it's not in the spec, it doesn't exist. If it's not verified, it's not done."

### Operational Mandates
- **Logic-First Implementation**: No hidden CLI dependencies; all governance is transparently documented in Markdown.
- **Deterministic Workflows**: Strict adherence to the 4-Phase Loop (Discovery, Specify, Implement, Review).
- **Centralized Memory**: All project state, long-term context, and technical learnings are persisted in the `.specs/project/` memory triad.

## 3. Skill Catalog (The Hub)
The Hub hosts specialized modules that can be consumed by AI agents to perform complex tasks.

### Foundation & Governance
- **sdd**: The core framework driving the entire ecosystem.
- **skill-factory**: The engine for creating and standardizing new skills.
- **git-workflow**: Standards for atomic commits and version integrity.

### Engineering & Architecture
- **architecture**: System design patterns, ADR management, and Mermaid visualization.
- **clean-code-mentor**: SOLID, YAGNI, and DRY enforcement.
- **benchmark-expert**: Performance measurement and regression detection.
- **observability-expert**: SRE, logs, and telemetry standards.

### Languages & Frameworks
- **python-uv**: Modern Python ecosystem management with UV (Django, Async).
- **django-expert**: Specialized Django production architecture.
- **fastapi-expert**: High-performance FastAPI implementation.
- **flutter-fvm**: Professional Flutter development with version management.

### Specialized Tools
- **brainstorming**: Design facilitation and complex problem exploration.
- **token-distiller**: Token density management (Caveman vs. Premium modes).
- **youtube-transcript**: High-performance transcript extraction and processing.

## 4. Architectural Integrity
- **Structure**: All skills must reside in their own directory with a mandatory `SKILL.md`.
- **Interoperability**: Skills are designed to be "equipped" by agents via direct reference or local mirroring.
- **Quality**: Continuous validation against the SDD mandates.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T10:38:00Z"
evidence_checksum: "NONE"
```
