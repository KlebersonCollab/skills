# Project Incremental Wisdom (Learnings)

## Automation & Scripting
- **Zipping Hidden Folders**: When creating ZIPs of hidden folders (like `.claude`), it is fundamental to enter the parent directory (`cd dist_staging`) before executing the `zip` command to ensure the ZIP contains the folder as its root, facilitating use by the end user.
- **Dynamic Skill Discovery**: Using `find . -maxdepth 2 -name "SKILL.md"` is a robust strategy for listing skills without needing a manual manifest file, making the project self-expanding.

## CI/CD (GitHub Actions)
- The separation of artifacts by name in `upload-artifact` v4 greatly improves the user experience.
- **GitHub Releases vs Artifacts**: For final software distribution (like skill ZIPs), Releases are preferable because they are permanent and allow static links (`/releases/latest/download/`), which simplifies integration and documentation.
- **Tag-Driven Releases**: Automating release creation from tag pushes (`v*`) ensures that only validated versions are distributed to users.

## Skill Management & Documentation
- **Centralized Version Tracking**: Maintaining a centralized catalog with updated versions of all 16 skills is essential for ecosystem consistency.
- **Mermaid Diagrams for Navigation**: Visual diagrams in the skill catalog significantly improve the understanding of the ecosystem and workflows.
- **Automated Version Collection**: Scripts to automatically extract versions and descriptions from `SKILL.md` files prevent the catalog from becoming outdated.
- **Onboarding Navigator as Living Documentation**: The onboarding skill must evolve with the ecosystem, reflecting changes in real-time.

## 2026-04-16 - Knowledge Graphing Strategy
- **Learning**: Implementing a modular 'knowledge-architect' is superior to bloating the 'sdd' or 'harness' skills.
- **Pattern**: Using Mermaid for relational mapping allows agents to visualize impact before coding.

## 2026-04-18 - Expert Skills & Global Sync Audit
- **Sync Blind Spots**: It was discovered that synchronizing only text files (like global mandates) is not enough for the ecosystem. New skills created in the root remain "invisible" to agents if the sync script does not perform a complete directory mirroring to hidden folders (`.agent/skills`, etc.).
- **Automation of Consistency**: Simple "cross-check" tools between the README and the skill frontmatter (like `verify_versions.py`) are fundamental to prevent the repository from entering a state of documental entropy.
- **Skill Consolidation**: Integrating knowledge from multiple external sources (GitHub) and modularizing into specific references (like `orm_performance.md` and `security.md`) makes the skill much more actionable than a single giant file.
- **Mandatory Skill Governance**: Explicitly defining in global mandates that the use of skills is mandatory reduces ad-hoc execution and ensures that agents operate within the quality tracks established by the Hub.
- **Workflow Automation with Makefile**: Creating a centralized Makefile reduces the developer's cognitive load and ensures that complex processes (like mandate sync and dist generation) are executed identically by any agent or human.
- **Automated Knowledge Mapping**: The `generate_knowledge_map.py` script allows visualizing the project topology and dependencies between skills and features in real-time, facilitating impact analysis in refactorings.

### [2026-04-19] Rigorous Changelog Validation
- **Occurrence**: Pipeline failure due to date format in CHANGELOG.md.
- **Solution**: The local validator strictly requires the format '## [X.Y.Z] - YYYY-MM-DD' (with simple hyphen and spaces). Using a dash (—) or alternative formats results in FAIL.

### [2026-04-20] Deterministic Enforcement & Context Inertia
- **Context Inertia**: LLMs tend to ignore "cleanup" instructions (sync/memory) upon reaching the user's direct goal. The solution is to transform these actions into imperative and visual **Exit Gates** in the context.
- **Enforcement Layering**: Reinforcement works best in layers:
    1. **Global Mandate** (Bootstrap/Exit Gate at the top of the context).
    2. **Local Hook** (Prerequisites within the specific skill).
    3. **Total Visibility** (Automatic compilation of all skills in .cursorrules).
- **Redundancy is Reliability**: Having the SDD Hook both in the mandate and within the stack skill ensures the agent doesn't skip steps, even if the task is short.
- **The "Brain Shell" Pattern**: Grouping operational mandates under an "Enforcement Shell" subgraph in the Knowledge Map helps the agent visualize that governance is not optional, but the system's skeleton.

### [2026-04-20] Automated Distilling & Memory Resilience
- **Automated Knowledge Distiller**: Extracting metadata directly from the YAML frontmatter of the skills to generate the Mermaid map reduces human error and allows dynamic visualization of versions and categories. Using visual badges (🛡️) for SDD compliance creates a layer of positive social/technical pressure for quality maintenance.
- **Auto-Fix as Resilience**: In complex ecosystems, tools that only detect errors (Sensors) are not enough. Having scripts that "repair" the state (like `auto_fix_memory.py`) allows the system to recover operational integrity without human intervention, which is crucial for autonomous agents.
- **The v2.0 Milestone**: Transitioning to a deterministic ecosystem (Bootstrap/Exit Gates) eliminates the need to "hope" the agent follows the patterns. Governance is now encoded in the session lifecycle.

### [2026-04-20] Code Audit Findings (Full Hub)
- **Duplication Crisis**: Skills are duplicated in ~5 locations (`.agent/skills/`, `.gemini/skills/`, `.claude/skills/`, `dist_staging/`, `.claude/worktrees/`). This violates DRY and creates a risk of versioning inconsistency.
- **Safe File Operations**: The `sync_mandates.py` script uses `shutil.rmtree` before `copytree`, creating a risk of data loss if there is an interruption. Solution: use TemporaryDirectory with safe transaction.
- **Code Sharing Opportunity**: Directory walking logic and exclude lists are replicated in 4+ scripts. Creating a shared `scripts/utils.py` would solve this.

### [2026-04-21] Skill Enrichment & Content Density
- **Learning**: A skill without practical examples and external references is 50% less effective for an AI agent. Context density (`examples/` and `references/` folders) allows the agent to understand not only *what* to do, but *how* to do it following market standards.
- **Pattern**: Using standardized directories (`references/`, `resources/`, `examples/`) in all skills creates a predictable navigation interface for the agent, reducing information search latency.

### [2026-04-22] Consolidation of "Expert" Skills
- **Learning**: Integrating knowledge from multiple external sources (Patterns, Security, TDD, Verification) into a single local "Expert" skill drastically reduces context switching and creates a robust "Gold Standard".
- **Pattern**: The 12-phase workflow (from Setup to Verification) ensures that no technical pillar is neglected during the development lifecycle.

### [2026-04-22] Mobile Modernization and Resilience
- **Learning**: Integrating modern patterns (Dart 3 Sealed Classes) into Flutter skills not only improves the code but forces the agent to think in terms of **State Exhaustiveness**, eliminating entire classes of UI bugs.
- **Pattern**: Including **Mandatory Accessibility** and **Global Error Handling** in the skill's core workflow ensures these "non-functional" pillars are treated as first-class requirements in all tasks.

### [2026-04-22] Cohesion Audit & Validator Evolution
- **Learning (Validator Lag)**: Validation scripts tend to lag behind documental governance (e.g.: `MEMORY.md`). It is vital that governance policy changes are immediately reflected in the auditor's code (`validate_skills.py`).
- **Pattern (Mass Fixing)**: Using mass-fixing scripts (`fix_missing_hooks.py`) is the only scalable way to maintain cohesion in a Hub with 20+ skills, ensuring the "Gold Standard" is applied uniformly without human error.
- **Enforcement Success**: Manual cohesion audit followed by automation raised the Hub from 40% to 100% compliance with the `🔒 Prerequisites (Mandatory)` mandate.

### [2026-04-22] GAN-style Harness Integration & Quality Enforcement
- **Learning**: AI agents tend to be "pathological optimists" when evaluating their own work. Separating the role of **Generator** (who implements) and **Evaluator** (who criticizes relentlessly) via the `harness-expert` skill creates an adversarial tension that eliminates "AI slop" and ensures production-level polish.
- **Pattern (The Quality Loop)**: Introducing a 4-pillar rubric (**Design Quality, Craft, Originality, Functionality**) with a "Pass Threshold" transforms SDD Phase 4 (Review) into an iterative cycle, ensuring the final result is not only functional but premium.
- **Implementation (Cross-Skill Fusion)**: Integrating this logic directly into the `sdd` framework and `onboarding-navigator` institutionalizes quality as part of the infrastructure, and not just as an optional desire.

### [2026-04-22] Idiomatic Go Patterns & Composition
- **Learning**: In Go, simplicity and composition are the pillars of maintainability. Adopting patterns like "Accept Interfaces, Return Structs" and "Zero Value Useful" drastically reduces the need for complex mocks and verbose initializations, making the system more resilient and easier to understand.
- **Pattern (Interface Decoupling)**: Defining interfaces on the consumer side (Consumer-Defined Interfaces) is the key to real decoupling in Go, allowing different implementations to satisfy requirements without the provider needing to know the contract.
- **Enforcement (Concurrency Safety)**: Adopting `errgroup` simplifies goroutine orchestration and ensures errors in parallel flows are not lost, raising the Hub's concurrency safety standard.

### [2026-04-22] SDD-Git Compliance & Mandatory Rigor
- **Learning**: Technical rigor is not optional. The absence of explicit prohibitive mandates can lead to the relaxation of critical workflows (like atomic commits per task). Governance must be encoded as an insurmountable "guardrail", not just a recommendation.
- **Pattern (Atomic Cycle)**: The `Task -> Test -> Commit` cycle must be the indivisible atom of development. Any progress marking in `tasks.md` without the corresponding validation artifacts (tests) and persistence (git) is an integrity failure.
- **Enforcement (Prohibitive Mandates)**: Inserting prohibitive mandates ("NEVER...") in the core skills (`sdd`) is more effective for Hub consistency than passive guidelines, as it creates a clear Exit Gate for agent intelligence.
- **Branching Strategy (GitHub Flow)**: Committing directly to `main` is an anti-pattern that compromises stability. Using **Feature Branches** is mandatory to isolate development and allow audits before integration. The Hub must operate in "Always Shippable" mode.

### [2026-04-22] Python Expert Enrichment & UV Synergy
- **Learning**: Integrating language patterns (Patterns) and testing methodology (TDD) directly into the environment management skill (`python-uv`) creates a very powerful "Language + Tool" foundation. UV is not just an installer; it is the engine that enables the `Test -> Refactor` cycle with zero latency.
- **Pattern**: Using `Protocol` instead of multiple inheritance or rigid abstract base class inheritance is the modern idiomatic path for Python, facilitating Mocking and maintenance of large codebases.
- **Enforcement**: Coverage goal (80%+) and mandatory TDD cycle must be an integral part of the "Quality Rules" documentation of any technical stack skill.

### [2026-04-22] Benchmark Expert & Git Rigor
- **Learning**: Adapting external skills (ECC) requires not just translation, but re-contextualization for local mandates (SDD, Git Workflow). Failing to start development with a feature branch is a sign of "Context Drift" that must be corrected immediately via Exit Gate audit.
- **Pattern**: The lifecycle of a new skill must be: `Spec -> Branch -> Scaffolding -> Atomic Commits -> Validation -> Registry`.
- **Enforcement**: Using deterministic benchmarks (average of 3 runs) and regression verdict (>15% FAIL) transform performance from a subjective desire into a verifiable technical requirement in SDD Phase 4.

# Project Learnings

## Technical Learnings
- **Go Microservices Patterns (ECC Integration)**: 
    - The Clean Architecture structure with `internal/domain`, `service`, `repository`, and `handler` is the gold standard for microservices in Go.
    - Using `sqlc` to generate type-safe SQL code and `Wire` for explicit dependency injection significantly reduces boilerplate and runtime errors.
    - `testcontainers-go` is the recommended tool for robust and isolated integration tests.
- **AI Operational Patterns**: 
    - Implementing Slash Commands (`/plan`, `/go-test`, `/go-review`) within the skill documentation itself standardizes AI responses and facilitates the user workflow.
- **Content Loss in Normalization [2026-04-23]**: During large-scale translations or "normalization" tasks, LLMs may tend to summarize content instead of performing a faithful 1:1 translation. This can result in significant loss of technical detail (up to 60% in some cases). 
    - *Mitigation*: Always perform byte-count or line-count audits after bulk translations to identify anomalies. Use tools like `git diff --stat` to find files with disproportionate size changes.

## Architectural Decisions
- **`init()` Prohibition**: Decision to ban the use of `init()` functions to ensure all dependencies are explicitly initialized in `main()`, facilitating testing and debugging.
- **Global Immutability**: Banning global mutable state to avoid unpredictable side effects in concurrent systems.

## Operational Learnings
- **SD-Consistency**: When enriching skills, it is crucial to validate if the command structure and quality rules are aligned with sub-skills (like `golang-testing-expert`).

### [2026-04-23] Validator Rigor & SDD Compliance Hooks
- **Learning**: The Hub's validation script (`validate_skills.py`) imposes absolute lexical rigor for the `## 🔒 Prerequisites (Mandatory)` block. Even if the content is semantically correct, the script looks for specific keywords like `Spec Check`, `Plan Check`, `Contract Check`, and `Task Check`.
- **Solution**: The structure of the Prerequisites block must be kept identical across all skills to ensure the automated audit passes without false negatives. SDD is the "glue" that holds the entire ecosystem together, and its verification is the first security gate of any agentic execution.
- **Impact**: The enrichment of the SDD skill for v1.5.0 now serves as the "gold template" for this section, ensuring new guidelines (like Brownfield Mapping) are integrated without breaking repository governance.
## [2026-04-23] Automated ADR Generation
- **Category**: Pattern
- **Description**: Implementing a CLI tool to automate ADR creation ensures consistency and traceability in architectural decisions.

## [2026-04-24] Teste final
- **Category**: Pattern
- **Description**: Ok


### [BENCHMARK] ./internal/benchmark: 1164579.00 ns/op (2026-04-24)

### [BENCHMARK] ./internal/benchmark: 1140608.00 ns/op (2026-04-24)

### [2026-05-06] Débito de Governança em Integrações Rápidas
- **Learning**: Durante fases de integração acelerada (ex: Stitch Skills), há uma tendência natural de negligenciar os hooks de governança (SDD Hooks) e arquivos mandatórios (`CHANGELOG.md`) nas novas habilidades criadas.
- **Pattern**: Auditorias SDD periódicas são essenciais para identificar esse "drift" documental antes que ele comprometa a capacidade de futuros agentes de operar com segurança no repositório.
- **Mitigation**: Implementar verificação de conformidade de habilidades no pipeline de CI/CD para garantir que nenhuma nova habilidade seja integrada sem os requisitos mínimos de governança.
- **ADK-Based Architecture [2026-05-06]**: A adoção do modelo de 5 camadas do Agent Development Kit permite uma separação clara entre as regras de longo prazo (CLAUDE.md) e as capacidades sob demanda (Skills), reduzindo drasticamente o consumo de tokens e a fadiga do contexto.
- **Centralização de Memória Operacional (SDD v2.2.0) [2026-05-06]**: Skills individuais não devem carregar seus próprios arquivos `STATE.md`, `MEMORY.md` ou `LEARNINGS.md`. Toda memória operacional deve ser agregada em `.specs/project/` para manter uma "Fonte da Verdade" única para o projeto. A criação de memórias locais gera fragmentação e desvio de padrão.
- **Governança de Root Integrity [2026-05-06]**: Alterações em diretórios ocultos de sistema (`.agents/`) são proibidas. O desenvolvimento e a governança de habilidades devem ocorrer de forma explícita na raiz do projeto para garantir transparência e conformidade com o princípio de "Explicit Logic" do SDD.
