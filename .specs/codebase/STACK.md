# Technology Stack

## Language & Format
- **Markdown** (`.md`): Exclusive format for skill definitions, specs, and documentation.

## Methodology
- **SDD v1.3.1** (Spec-Driven Development): Modular and adaptive workflow with **PRD/RFC/BDD** integration, **Diagram-as-Code** support, and persistent state management.

## Visualization Tools
- **Mermaid**: Mandatory standard for architectural diagrams and process flows integrated into Markdown.

## Observability & Telemetry
- **OpenTelemetry (OTel)**: Standard for distributed tracing and metrics design.
- **Structured Logging (JSON)**: Standard for machine-readable logging.

## Automation Ecosystems
- **UV**: Python project and package management.
- **FVM**: Flutter SDK version management.
- **Azure DevOps CLI**: Infrastructure automation and platform governance.

## Distribution & CI/CD
- **GitHub Actions**: Pipeline automation for validation, artifact generation, and distribution.
- **Artifacts (.zip)**: Pre-configured packages for multiple agents (.claude, .gemini, .agent) automatically generated.

## Version Control
- **Git**: Repository versioning.
- **GitHub**: Hosting and collaboration.

## Development Tools
- **Antigravity (AI Agent)**: AI agent used for pair programming and skill execution.

## Scaffolding
- **Skill Factory**: Core Framework (`skill-factory/`) for standardized creation of new skills.
- **Embedded Templates**: Scaffolding templates are embedded within sub-skill instructions (pure Markdown, no external tools).
- **Mandatory Validation**: Every created skill undergoes automated auditing before registration.

## Conventions
- Each skill resides in its own directory (e.g., `sdd/`).
- The main file of a skill is named `SKILL.md`.
- Sub-skills follow the pattern `<skill-name>-<sub-skill>.skill.md`.
- The root `README.md` must be updated with each new skill added.
- Every new skill must start at version `1.0.0`.
