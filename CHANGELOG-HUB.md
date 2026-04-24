# Hub Global Changelog

*Last updated: 2026-04-23 23:03:43*

This file consolidates the most recent updates from all skills in the Hub.

## [4.5.0] - 2026-04-23
### 🚀 Hub Super-Tools (Governance & Productivity)
- **hb adr [new|list]**: Automação total de Architecture Decision Records.
- **hb audit**: Auditoria de conformidade (Python/Go/Hub) com sistema de pontuação.
- **hb learn**: Registro estruturado de aprendizados e padrões em `LEARNINGS.md`.
- **hb skill new**: Scaffolding instantâneo de novas skills seguindo o padrão Hub.
- **hb doctor**: Diagnóstico e reparo automático de integridade do ambiente.
- **hb distill**: Otimização de tokens via compressão de arquivos e sessões.
- **hb sdd review**: Validador automatizado de artefatos SDD para features.

## [4.4.0] - 2026-04-23
### 🚀 Governance & Autonomy Hardening
- **hb init**: Novo comando root para inicialização "Zero-to-Hero" de projetos (Bootstrap + Mandatos + Core Skills).
- **hb sdd task**: Gestão de progresso atômica para tarefas SDD via CLI.
- **hb sync --remote**: Capacidade de baixar mandatos globais diretamente do repositório central.
- **Autonomous Skill Discovery**: Mandato global #8 implementado, obrigando agentes a buscarem e instalarem skills remotas sob demanda.
- **Skill Awareness**: Atualização de todas as 25+ skills para integração total com o ferramental HB-CLI v4.4.0.

## [4.3.0] - 2026-04-23
### 🚀 Hub Infrastructure Revolution (HB-CLI)
- **HB-CLI (Harness Binary)**: Substituição total de scripts Python de infraestrutura por binário nativo Go de alta performance.
- **Plug & Play Intelligence**: Motor de instalação recursivo que resolve dependências automaticamente (`hb install`).
- **Multi-Agent Governance**: Distribuição automática de regras para Gemini, Claude, Cursor e outros agentes (`hb rules`).
- **Industrial SDD**: Workflow Spec-Driven Development integrado nativamente ao CLI (`hb sdd`).
- **Zero-Link-Debt**: Auditoria automática de integridade de links e documentação (`hb check`).
- **Multi-OS Release**: Automação via GitHub Actions para distribuição de binários Win/Mac/Linux.
- **HB listskills**: Novo comando para listar skills disponíveis remotamente no GitHub diretamente via CLI.

---

## 🧩 api-architect
## [1.3.0] - 2026-04-14

### Changed
- **Workflow & References**: Improved workflow to explicitly reference the use of `references/api-style.md` during the design phase.

## [1.2.0] - 2026-04-14

### Changed
- **Renaming**: Skill renamed from `api-designer` to `api-architect` to reflect a larger scope of governance and resilience.
- **Style Guide**: Added `references/api-style.md` with standards for response envelopes, errors, and pagination.
- **Resilience**: Included Timeout, Retry, and Circuit Breaker patterns.

## [1.1.0] - 2026-04-14

### Added
- **tRPC Support**: Full reference in `references/trpc-patterns.md` for fullstack TypeScript projects.
- **Security Guide (OWASP API Top 10)**: New `references/api-security-guide.md` file focusing on BOLA, Broken Auth, Mass Assignment, etc.
- **Rate Limiting**: New guide in `references/api-rate-limiting.md` covering algorithms (Token Bucket) and header standardization (429 status).
- **Decision Tree**: Integrated into `SKILL.md` to facilitate choosing between REST, GraphQL, tRPC, and gRPC.
- **Advanced Pagination Patterns**: Support for Cursor-based pagination (Relay style) and response envelopes.

---

---

## 🧩 architecture
## [2.0.1] - 2026-04-14
### Added
- **Mermaid Diagrams** are now mandatory for the `System Map`.
- Focus on visualizing distributed components and flows.

## [2.0.0] - 2026-04-14

### Added
- **CQRS Blueprint**: Guidance for separating read and write models for high scale.
- **Event-Driven Architecture**: Support for asynchronous message flows, idempotency, and eventual consistency.
- **Evolutionary Architecture**: Introduction of **Fitness Functions** for automating architectural governance.
- New reference guides for CQRS/Events and Evolutionary Architecture.
- Mermaid diagram examples for CQRS and Event flows.

## [1.0.0] - 2026-04-14

### Added
- Initial version of the Systems Architect skill.
- 4-phase workflow: Context, Analysis, Design, and Document.
- Focus on Simplicity as top priority ("Less is more").
- Integration of ADRs (Architecture Decision Records) into the flow.
- References for principles (SOLID, DRY, YAGNI) and pattern selection.

---

## 🧩 azure-devops
## [1.1.0] - 2026-04-14

### Added
- **Governance & IaC Support**: Full management of Variable Groups and Service Connections.
- **Administration & Security**: Commands for Users, Teams, Permissions, and Extension management.
- **Advanced CLI Integration**: Support for `az devops configure` and raw API `invoke`.
- **Wiki Management**: Support for project and code-based wikis.
- New reference guides for Infrastructure, Administration, and Advanced CLI.
- Examples for Service Connection templates and Raw API calls.

## [1.0.0] - 2026-04-14

### Added
- Initial version of the Azure DevOps skill.
- Support for Azure Boards (Work Items, WIQL).
- Support for Azure Repos (PRs, Branches).
- Support for Azure Pipelines (Builds, Releases).
- Support for Azure Artifacts and Test Plans.
- Detailed reference guides for authentication and core modules.
- Examples for JSON Patch and WIQL queries.

---

## 🧩 benchmark-expert
## [1.0.0] - 2026-04-22

### Added
- Initialization of the `benchmark-expert` skill.
- Support for 4 performance modes: Browser, API, Build, and Comparison.
- Mandatory integration with the SDD framework.
- Reference guides for performance targets and regression analysis.
- Comparison report examples.

---

## 🧩 brainstorming
## [1.2.1] - 2026-04-23

### Changed
- Translated all skill documentation and references to **English** as per the new repository standard.

## [1.2.0] - 2026-04-23

### Added
- Integrated **Vanderbilt Prompt Patterns** (Flipped Interaction, Cognitive Verifier, etc).
- Added **Brainstorming Patterns Catalog** (Perspective Multiplication, Constraint Variation, Inversion, Analogical Transfer).
- New reference files in `references/` and templates in `resources/`.
- Support for structured **SCAMPER** and **Time Travel Mapping**.

## [1.1.0] - 2026-04-14

### Added
- Created the structured brainstorming framework.
- Integrated "5 Whys" and Stakeholder Mapping.
- Established the "Understanding Lock" as a safety gate.

---

## 🧩 clean-code-mentor
## [1.0.0] - 2026-04-14

### Added
- Initial version of the Clean Code Mentor skill.
- Support for SOLID principles audit.
- YAGNI and KISS complexity guard.
- DRY (Don't Repeat Yourself) cross-file detection.
- Actionable refactoring suggestions.
- Reference guides for SOLID, YAGNI, KISS, and DRY.
- Examples of "Bad vs. Good" code design.

---

## 🧩 devsecops-expert
## [1.0.0] - 2026-04-22

### Added
- Skill initialization.
- Inclusion of SAST/DAST and Zero Trust mandates.

---

## 🧩 django-expert
## [1.5.0] - 2026-04-22

### Added
- **Deployment Readiness & Verification** section (Phase 12).
- Mandate to use **Ruff** for linting and **Mypy** for Type Checking.
- Pre-deployment checklists and security audit.
- **GitHub Actions** template for automated verification.
- Consolidated redundancies in `SKILL.md`.

## [1.4.0] - 2026-04-22

### Added
- **Testing Excellence** section (Phase 11).
- Mandates for using **Factories** (Factory Boy) instead of static fixtures.
- Ultra-fast test infrastructure configurations (SQLite in-memory, `--nomigrations`).
- **Mocking** patterns for external services and APIs.
- Explicit code coverage goals per layer.

## [1.3.0] - 2026-04-22

### Added
- **Security Hardening** section (Phase 10).
- Hashing patterns with Argon2 and 12+ character password validators.
- Mandates for file type and size validation.
- RBAC (Role-Based Access Control) and Throttling patterns for APIs.
- Prohibitive rules against SQL Injection and unsafe use of `mark_safe`.

---

## 🧩 fastapi-expert
## [1.1.0] - 2026-04-20

### Added
- `🔒 Prerequisites (Mandatory)` section linking the skill mandatorily to the SDD framework.

## [1.0.0] - 2026-04-18

### Added
- Initialization of the `fastapi-expert` skill.
- Integration with the `python-uv` standard.
- Strict rules for `Annotated` and `Pydantic V2`.
- 4-phase workflow for API design and implementation.
- "Do's and Don'ts" examples based on official FastAPI standards.

---

## 🧩 flutter-fvm
## [1.3.0] - 2026-04-22

### Added
- **Modern Dart Patterns**: Full support for Dart 3+ (Sealed classes, Pattern Matching, Records).
- **Mandatory Accessibility**: Guidelines for Semantics, touch targets, and inclusive UX.
- **Advanced Performance**: Selective rebuilding with `MediaQuery.sizeOf` and widget decomposition.
- **Secret Management**: Standardized keys via `--dart-define` and `.env`.
- **Resilience**: Global error capture protocols (`FlutterError.onError`).

### Changed
- Updated `SKILL.md` to reflect ECC technical excellence standards.
- Enriched `references/testing-and-quality.md` with accessibility and golden tests.
- Enriched `references/project-structure.md` with MVVM + Dart 3 integration.

## [1.2.0] - 2026-04-20

### Added
- `🔒 Prerequisites (Mandatory)` section linking the skill mandatorily to the SDD framework.

## [1.1.0] - 2026-04-15

### Added
- Advanced testing patterns guide with layer-by-layer testing strategies.
- Flutter security guide based on OWASP Mobile Top 10 (2024).
- Riverpod state management testing patterns.
- Widget testing best practices with keys and platform testing.
- Security scanning workflows for hardcoded secrets and dependencies.
- Build protection recommendations with obfuscation.

### Enhanced
- Updated SKILL.md with security and advanced testing integration.
- Expanded "When to Use This Skill" section with security considerations.
- Enhanced DEVELOP phase with testing by architectural layers.
- Enhanced DEPLOY phase with security analysis and secure builds.
- Updated Best Practices with security do's and don'ts.
- Added references to new advanced guides in testing-and-quality.md.

---

## 🧩 frontend-expert
## [1.1.0] - 2026-04-20
### Added
- `🔒 Prerequisites (Mandatory)` section linking the skill mandatorily to the SDD framework.

## [1.0.0] - 2026-04-19
### Added
- Initial version of the `frontend-expert` skill.
- Support for React, Next.js, Tailwind v4, and Shadcn/UI.
- Workflow focused on Server Components and Accessibility.

---

## 🧩 git-workflow
## [1.0.0] - 2026-04-22
### Added
- Initial version of the skill imported from ECC and localized.
- Integration with the SDD (Spec-Driven Development) framework.
- Conventional Commits and Branching Strategies guides.
- Mandates for using Merge vs. Rebase.
- Pull Request templates and Git configuration.

---

## 🧩 golang-expert
## [1.2.0] - 2026-04-22

### Added
- Integrated Slash Commands for workflow automation.
- Established mandatory Clean Architecture layout and rigorous testing.

## [1.1.0] - 2026-04-20

### Added
- `🔒 Prerequisites (Mandatory)` section linking the skill mandatorily to the SDD framework.

## [1.0.0] - 2026-04-19

### Added
- **SKILL.md**: Main definition of the Expert level Go skill.
- **README.md**: Presentation and usage documentation.
- **Reference Structure**: Consolidation of 35 external knowledge sources into 5 local pillars.
- **foundations.md**: Pillar 1 focused on foundations and idiomatic style.
- **concurrency.md**: Pillar 2 focused on concurrency and safety.
- **development.md**: Pillar 3 focused on tools and processes.
- **architecture.md**: Pillar 4 focused on systems and infrastructure.
- **ecosystem.md**: Pillar 5 focused on modern libraries and Samber.

---

## 🧩 golang-testing-expert
## [1.1.0] - 2026-04-22
### Added
- Detailed verification commands (Race, Coverage, Integration).
- Integrated TDD cycle with slash commands.

## [1.0.0] - 2026-04-22
### Added
- Skill creation based on ECC.
- References for TDD, Table-Driven Tests, and Subtests.
- References for Mocks, httptest, and Golden Files.
- References for Benchmarks, Fuzzing, and Coverage.

---

## 🧩 harness-expert
## [2.1.0] - 2026-04-22

### Added
- **GAN-style Feedback Loop**: Implementation of the adversarial cycle between Generator and Evaluator for Large/Complex tasks.
- **Evaluation Rubric**: New scoring matrix (Design, Craft, Originality, Functionality) with a target Score >= 7.0.

## [2.2.0] - 2026-04-23
### Added
- **HB listskills Integration**: Support for listing remote skills on GitHub, enabling the agent to "self-expand" its capabilities by discovering new modular intelligence.

## [2.0.0] - 2026-04-20

### Changed
- **Role Refactoring**: The skill now focuses exclusively on the **technical engine** (agentic infrastructure).
- **Decoupling**: Strategic memory management removed, delegating the operational vision protocol to the `sdd-planner`.
- **Workflow**: Updated to include `AUTOMATED REHYDRATE`, `OPERATE`, `AUTO-SYNC`, and `VALIDATE` phases.

## [1.1.0] - 2026-04-18

### Changed
- **Global Renaming**: Official transition from "Hardness" to **Harness Engineering** to align with industry standards.
- **Enhanced Rehydration**: Added **Bootstrap** step for operational environment validation and automatic dependency installation.

### Added
- **Harness Principles**: Documentation of Feed Forward, Feedback, and Sensors in the README.

---

## 🧩 knowledge-architect
## [1.0.0] - 2026-04-16

### Added
- **SKILL.md**: Initial definition of the Local Knowledge Graphing (LKG) protocol.
- **README.md**: Overview and application scenarios for Local Knowledge Graphs.
- **Workflow**: Implemented 4-phase cycle (Extract, Link, Traverse, Prune).
- **Artifacts**: Defined `KNOWLEDGE-MAP.mermaid` and `RELATIONS.md` as mandatory outputs.

---

## 🧩 observability-expert
## [1.0.0] - 2026-04-14

### Added
- Initial version of the Observability Expert skill.
- 4-phase SRE Workflow: Instrument, Collect, Visualize, React.
- Guidelines for Structured Logging (JSON).
- Design patterns for OpenTelemetry (Metrics & Tracing).
- Methodology for SLI/SLO and Alerting policies.
- Reference guides and practical examples with Mermaid diagrams.

---

## 🧩 onboarding-navigator
## [1.5.0] - 2026-04-23
### Changed
- **Integrity Audit**: Strict compliance audit performed.
- **Sync**: Synchronized to reflect the new total of 25 skills.
- **Token Distiller**: Deep integration with Caveman mode for session bootstrap.

## [1.4.0] - 2026-04-22
### Changed
- Update to Hub Exit Gate and SDD audits.

## [1.3.1] - 2026-04-22
### Changed
- **Governance Hardening**: Inclusion of commit and branch audit in the Session Exit Gate.
- **Prohibitive Mandates**: Added rule against ending a session without a Git compliance audit.

---

## 🧩 python-uv
## [3.0.0] - 2026-04-22

### Added
- **Expert Enrichment**: Integration of high-level knowledge about Python patterns and testing.
- **Python Patterns**: New reference in `references/patterns.md` covering EAFP, Protocol, Type Hints, and Performance.
- **Advanced Testing**: New reference in `references/testing.md` covering TDD cycle, advanced Fixtures, and Mocking.
- **Modern Python**: Explicit support for Python 3.9+ idioms and `__slots__` for memory efficiency.

## [2.6.0] - 2026-04-20

### Added
- `🔒 Prerequisites (Mandatory)` section linking the skill mandatorily to the SDD framework.

## [2.5.0] - 2026-04-14

### Added
- **Operational Workflow**: Introduction of a structured 4-phase workflow (Environment, Project, Develop, Deploy) to guide AI agents step-by-step.

---

## 🧩 scaffolding-expert
## [1.0.0] - 2026-04-19
### Added
- Base structure for the skill.
- Workflow specification for use with `uvx copier`.

---

## 🧩 sdd
## [1.5.0] - 2026-04-23

### Added
- **Spec-Driven TLC Integration**: Added 5 new advanced operational guidelines in `references/` (`brownfield-mapping.md`, `coding-principles.md`, `context-limits.md`, `session-handoff.md`, `quick-mode.md`).
- **Sub-Agent Delegation**: Documentation of the new delegation protocol for orchestrators to mitigate context window limits.
- **Audit Compliance (Skill Factory)**: Added mandatory `## Goal` and `## Quality Rules` in `sdd-implementer`, `sdd-orchestrator`, and `sdd-reviewer` sub-skills.

### Changed
- **Mass Version Update**: Sub-skills updated to `1.5.0` synchronizing with the base skill.

## [1.4.0] - 2026-04-16

### Added
- **Harness Engineering Integration**: Workflow evolution focusing on contracts, sensors, and deterministic feedback loops.
- **SDD Contracts**: New `contract.md` artifact to define the delivery agreement between Implementer and Reviewer.
- **Sensor-based Reviews**: Verdicts are now based on signals from external tools (linters, tests) with a Global Score (0-100).

### Changed
- **Version Unification**: Synchronization of all sub-skills to version 1.4.0 for ecosystem consistency.
- **Auto-Sizing Refinement**: Updated phase table to include Contract and Review by Score milestones.
- **Naming Correction**: Replaced residual term `ork-orchestrator` with `sdd-orchestrator`.

## [1.3.1] — 2026-04-14

### Added
- **Mermaid Diagrams Mandate**: Mandatory inclusion of Mermaid diagrams in the `Specify` phase for Medium+ levels.

---

## 🧩 skill-factory
## [1.1.0] - 2026-04-16

### Added
- **skill-factory-validator**: Added Check 6 (Registry Status) and Check 7 (Onboarding Sync).
- **GitHub Actions**: CI/CD integration for automatic compliance validation.

## [1.0.0] — 2026-04-14

### Added
- **SKILL.md**: Main definition of the Core Framework with Auto-Sizing (Quick/Standard), 3-phase workflow, and Version Policy.
- **skill-factory-bootstrap.skill.md**: Sub-skill for automated scaffolding with standardized templates for `SKILL.md`, `README.md`, `CHANGELOG.md`, and sub-skills.
- **skill-factory-validator.skill.md**: Sub-skill for compliance validation with 5 checks (Structural, Frontmatter, Content, Naming, Registry) and evidence report.
- **README.md**: Detailed documentation with usage examples for both modes (Quick and Standard).

---

## 🧩 swarm-facilitator
## [1.1.0] - 2026-04-20
### Changed
- **Renaming**: Skill renamed from `multi-agent-orchestrator` to `swarm-facilitator` to avoid semantic collision with `sdd-orchestrator`.
- **Identity**: Updated to focus explicitly on choreography and persona handoff protocols.

## [1.0.0] - 2026-04-19
### Added
- Initial structure and definition of Swarm personas.
- Formal Handoff protocol.

---

## 🧩 token-distiller
## [1.1.0] - 2026-04-23
### Changed
- **Hardening**: Added mandatory sections (Goal, Output, Rules, Prohibited).
- **Compliance**: Added SDD governance checks to the mandatory hook.
- **Structure**: Created `references/`, `resources/`, and `examples/` folders.

## [1.0.0] - 2026-04-22
### Added
- Initial implementation of the Token Optimization epic.
- Caveman and Premium modes.
- Integration with `hub` CLI.
- Python compression engine.

---

## 🧩 youtube-transcript
## [1.0.0] - 2026-04-15

### Added
- **SKILL.md**: Main skill definition with decision flow and fallback.
- **README.md**: Overview and usage documentation.
- **Base Structure**: Initial scaffolding following Skills Hub standards.

---

