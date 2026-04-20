# Hub Global Changelog
*Última atualização: 2026-04-20 16:20:52*

Este arquivo consolida as atualizações mais recentes de todas as skills do Hub.

## 🧩 api-architect
## [1.3.0] - 2026-04-14

### Alterado
- **Workflow & References**: Melhoria no workflow para referenciar explicitamente o uso do `references/api-style.md` durante a fase de design.

## [1.2.0] - 2026-04-14

### Alterado
- **Renomeação**: Skill renomeada de `api-designer` para `api-architect` para refletir um escopo maior de governança e resiliência.
- **Style Guide**: Adição de `references/api-style.md` com padrões de envelopes de resposta, erros e paginação.
- **Resilience**: Inclusão de padrões de Timeouts, Retries e Circuit Breakers.

## [1.1.0] - 2026-04-14

### Adicionado
- **Suporte a tRPC**: Referência completa em `references/trpc-patterns.md` para projetos TypeScript Fullstack.
- **Guia de Segurança (OWASP API Top 10)**: Novo arquivo `references/api-security-guide.md` focando em BOLA, Broken Auth, Mass Assignment, etc.
- **Rate Limiting**: Novo guia em `references/api-rate-limiting.md` cobrindo algoritmos (Token Bucket) e padronização de headers (status 429).
- **Árvore de Decisão**: Integrada ao `SKILL.md` para facilitar a escolha entre REST, GraphQL, tRPC e gRPC.
- **Padrões de Paginação Avançada**: Suporte a Cursor-based pagination (Relay style) e envelopes de resposta.

---

---

## 🧩 architecture
## [2.0.1] - 2026-04-14
### Adicionado
- Diagramas **Mermaid** tornam-se mandatórios para o `System Map`.
- Foco em visualização de fluxos e componentes distribuídos.

## [2.0.0] - 2026-04-14

### Adicionado
- **CQRS Blueprint**: Orientação para separação de modelos de leitura e escrita para alta escala.
- **Event-Driven Architecture**: Suporte a fluxos de mensagens assíncronas, idempotência e consistência eventual.
- **Evolutionary Architecture**: Introdução de **Fitness Functions** para automação da governança arquitetural.
- Novos guias de referência para CQRS/Eventos e Arquitetura Evolutiva.
- Exemplos de diagramas Mermaid para fluxos CQRS e Eventos.

## [1.0.0] - 2026-04-14

### Adicionado
- Versão inicial da skill de Arquiteto de Sistemas.
- Workflow em 4 fases: Context, Analysis, Design e Document.
- Foco em Simplicidade como prioridade máxima ("Less is more").
- Integração de ADRs (Architecture Decision Records) no fluxo.
- Referências para princípios (SOLID, DRY, YAGNI) e seleção de padrões.

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

## 🧩 brainstorming
## [1.1.0] - 2026-04-14

### Adicionado
- **Output Structure**: Seção H2 obrigatória detalhando os artefatos de saída (Design Spec, Decision Log, Trade-off Matrix).

## [1.0.0] - 2026-04-14

### Adicionado
- Versão inicial da skill de facilitador de brainstorming.
- Workflow em 4 fases: Discover, Diverge, Converge e Specify.
- Conceito de "Understanding Lock" para validação de premissas.
- Referências para técnicas criativas (SCAMPER, 5 Whys, First Principles).
- Regras de qualidade para facilitadores (uma pergunta por vez, suposições explícitas).

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

## 🧩 django-expert
## [1.1.0] - 2026-04-20

### Added
- Seção `🔒 Prerequisites (Mandatory)` vinculando a skill obrigatoriamente ao framework SDD.

## [1.0.0] - 2026-04-18

### Added
- Inicialização da skill `django-expert`.
- Guia de performance ORM (N+1 solutions).
- Integração com HTMX e padrões de reatividade.
- Suporte nativo ao workflow `python-uv`.
- Estrutura recomendada de camadas (Services/Managers).

---

## 🧩 fastapi-expert
## [1.1.0] - 2026-04-20

### Added
- Seção `🔒 Prerequisites (Mandatory)` vinculando a skill obrigatoriamente ao framework SDD.

## [1.0.0] - 2026-04-18

### Added
- Inicialização da skill `fastapi-expert`.
- Integração com o padrão `python-uv`.
- Regras estritas para `Annotated` e `Pydantic V2`.
- Workflow de 4 fases para design e implementação de APIs.
- Exemplos de "Do's and Don'ts" baseados no padrão oficial do FastAPI.

---

## 🧩 flutter-fvm
## [1.2.0] - 2026-04-20

### Added
- Seção `🔒 Prerequisites (Mandatory)` vinculando a skill obrigatoriamente ao framework SDD.

## [1.1.0] - 2026-04-15

### Added
- Advanced testing patterns guide with layer-by-layer testing strategies
- Flutter security guide based on OWASP Mobile Top 10 (2024)
- Riverpod state management testing patterns
- Widget testing best practices with keys and platform testing
- Security scanning workflows for hardcoded secrets and dependencies
- Build protection recommendations with obfuscation

### Enhanced
- Updated SKILL.md with security and advanced testing integration
- Expanded "When to Use This Skill" section with security considerations
- Enhanced DEVELOP phase with testing by architectural layers
- Enhanced DEPLOY phase with security analysis and secure builds
- Updated Best Practices with security do's and don'ts
- Added references to new advanced guides in testing-and-quality.md

## [1.0.0] - 2026-04-14

### Added
- Initial version of the Flutter with FVM skill.
- Comprehensive workflow for Flutter development using FVM.
- Reference guides for Environment, Project, Dependencies, Testing, and Deployment.
- Examples for `pubspec.yaml`, `analysis_options.yaml`, and GitHub Actions.

---

## 🧩 frontend-expert
## [1.1.0] - 2026-04-20
### Added
- Seção `🔒 Prerequisites (Mandatory)` vinculando a skill obrigatoriamente ao framework SDD.

## [1.0.0] - 2026-04-19
### Added
- Versão inicial da skill `frontend-expert`.
- Suporte a React, Next.js, Tailwind v4 e Shadcn/UI.
- Workflow focado em Server Components e Acessibilidade.

---

## 🧩 golang-expert
## [1.1.0] - 2026-04-20

### Adicionado
- Seção `🔒 Prerequisites (Mandatory)` vinculando a skill obrigatoriamente ao framework SDD.

## [1.0.0] - 2026-04-19

### Adicionado
- **SKILL.md**: Definição principal da skill de nível Expert em Go.
- **README.md**: Documentação de apresentação e uso.
- **Estrutura de Referências**: Consolidação de 35 fontes de conhecimento externas em 5 pilares locais.
- **foundations.md**: Pilar 1 focado em fundamentos e estilo idiomático.
- **concurrency.md**: Pilar 2 focado em concorrência e segurança.
- **development.md**: Pilar 3 focado em ferramentas e processos.
- **architecture.md**: Pilar 4 focado em sistemas e infraestrutura.
- **ecosystem.md**: Pilar 5 focado em bibliotecas modernas e Samber.

---

## 🧩 harness-expert
## [2.0.0] - 2026-04-20

### Alterado
- **Refatoração de Papel**: A skill agora foca exclusivamente no **motor técnico** (infraestrutura agêntica).
- **Desacoplamento**: Removida a gestão estratégica de memória, delegando o protocolo de visão operacional para o `sdd-planner`.
- **Workflow**: Atualizado para incluir as fases de `AUTOMATED REHYDRATE`, `OPERATE`, `AUTO-SYNC` e `VALIDATE`.

## [1.1.0] - 2026-04-18
 — 2026-04-18

### Alterado
- **Renomeação Global**: Transição oficial de "Hardness" para **Harness Engineering** para alinhar com padrões da indústria.
- **Enhanced Rehydration**: Adicionado passo de **Bootstrap** para validação do ambiente operacional e instalação automática de dependências.

### Adicionado
- **Harness Principles**: Documentação de Feed Forward, Feedback e Sensores no README.

## [1.0.0] — 2026-04-15

### Adicionado
- **SKILL.md**: Definição principal da skill de Harness Engineering.
- **README.md**: Documentação detalhada e visão geral.
- **harness-expert-sync.skill.md**: Sub-skill para sincronização automática de estado.
- **harness-expert-rehydrate.skill.md**: Sub-skill para injeção de contexto operacional.
- **harness-expert-compress.skill.md**: Sub-skill para compactação de contexto (Context Compressor).
- **Resources**: Implementado `compressor.py` para análise e compactação de tarefas.
- **Exemplos**: Adicionado script de demonstração de sincronização.

---

## 🧩 knowledge-architect
## [1.0.0] - 2026-04-16
 — 2026-04-16

### Adicionado
- **SKILL.md**: Definição inicial do protocolo de Local Knowledge Graphing (LKG).
- **README.md**: Visão geral e cenários de aplicação para Grafos de Conhecimento Locais.
- **Workflow**: Implementado ciclo de 4 fases (Extract, Link, Traverse, Prune).
- **Artefatos**: Definidos `KNOWLEDGE-MAP.mermaid` e `RELATIONS.md` como saídas mandatórias.

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
## [1.2.1] - 2026-04-20
### Fixed
- Auditoria de sincronização: Corrigida a contagem total de skills de 18 para 20 em toda a documentação.
- Atualização das estatísticas de categorias (5 Metodologia, 3 Arquitetura, 8 DevOps, 2 Automação, 2 Navegação).

## [1.2.0] - 2026-04-19
### Alterado
- Auditoria de completude: Atualizado de 11 para 18 skills no catálogo visual e de decisão.
- Adição dos fluxos de `Scaffolding`, `Orchestrator`, `FastAPI` e `Django` na Matriz de Decisão.
- Restauração completa de nós faltantes na hierarquia do Mermaid.

## [1.1.0] - 2026-04-15

### Added
- Diagramas Mermaid para visualização do ecossistema de skills
- Matriz de decisão visual com fluxos de trabalho recomendados
- Estatísticas do hub (11 skills total, por categoria)
- Seção de "Fluxo de Atualização" para manutenção da skill
- Referências a STATE.md, MEMORY.md e LEARNINGS.md no workflow

### Updated
- Catálogo de skills com versões atualizadas de todas as 11 habilidades
- Descrições precisas extraídas dos arquivos SKILL.md originais
- Correção de versões: flutter-fvm (1.0.0 → 1.1.0), azure-devops (1.0.0 → 1.1.0)
- Estrutura do SKILL.md com melhor organização visual
- Regras de qualidade com foco em "Stats-Aware" (consciência do ecossistema)

### Fixed
- Referências a versões desatualizadas no catálogo
- Falta de diagramação visual no processo de onboarding
- Documentação incompleta sobre a contagem total de skills

---

## 🧩 python-uv
## [2.6.0] - 2026-04-20

### Adicionado
- Seção `🔒 Prerequisites (Mandatory)` vinculando a skill obrigatoriamente ao framework SDD.

## [2.5.0] - 2026-04-14

### Adicionado
- **Workflow Operacional**: Introdução de um workflow estruturado em 4 fases (Environment, Project, Develop, Deploy) para orientar agentes de IA passo a passo.

## [2.4.0] - 2026-04-14

### Adicionado
- **Output Structure**: Seção H2 obrigatória detalhando os artefatos gerados pelo UV (pyproject.toml, uv.lock, .python-version).

---

## 🧩 scaffolding-expert
## [1.0.0] - 2026-04-19
### Adicionado
- Estrutura base para a skill.
- Especificação de workflow para uso com `uvx copier`.

---

## 🧩 sdd
## [1.4.0] - 2026-04-16
 — 2026-04-18

### Adicionado
- **Harness Engineering Integration**: Evolução do workflow com foco em contratos, sensores e feedback loops determinísticos.
- **SDD Contracts**: Novo artefato `contract.md` para definir o acordo de entrega entre Implementador e Revisor.
- **Sensor-based Reviews**: Vereditos agora são baseados em sinais de ferramentas externas (linters, testes) com Score Global (0-100).

### Alterado
- **Unificação de Versões**: Sincronização de todas as sub-skills para a versão 1.4.0 para consistência do ecossistema.
- **Refinamento do Auto-Sizing**: Atualização da tabela de fases para incluir os marcos de Contrato e Review por Score.
- **Correção de Nomenclatura**: Substituído o termo residual `ork-orchestrator` por `sdd-orchestrator`.

## [1.3.1] — 2026-04-14

### Adicionado
- **Mermaid Diagrams Mandate**: Tornada obrigatória a inclusão de diagramas Mermaid na fase de `Specify` para níveis Medium+.

## [1.3.0] — 2026-04-14

### Adicionado
- **PRD/RFC/BDD Integration**: Incorporação de Product Requirements Documents, Request for Comments e BDD (Given/When/Then) ao workflow baseado no Auto-Sizing.

---

## 🧩 skill-factory
## [1.1.0] - 2026-04-16
 — 2026-04-16

### Adicionado
- **skill-factory-validator**: Adicionado Check 6 (Registry Status) e Check 7 (Onboarding Sync).
- **GitHub Actions**: Integração com CI/CD para validação automática de conformidade.

## [1.0.0] — 2026-04-14

### Adicionado
- **SKILL.md**: Definição principal do Core Framework com Auto-Sizing (Quick/Standard), workflow em 3 fases e Version Policy.
- **skill-factory-bootstrap.skill.md**: Sub-skill para scaffolding automatizado com templates padronizados para `SKILL.md`, `README.md`, `CHANGELOG.md` e sub-skills.
- **skill-factory-validator.skill.md**: Sub-skill para validação de conformidade com 5 checks (Structural, Frontmatter, Content, Naming, Registry) e relatório de evidências.
- **README.md**: Documentação detalhada com exemplos de uso para ambos os modos (Quick e Standard).

---

## 🧩 swarm-facilitator
## [1.1.0] - 2026-04-20
### Alterado
- **Renomeação**: Skill renomeada de `multi-agent-orchestrator` para `swarm-facilitator` para evitar colisão semântica com o `sdd-orchestrator`.
- **Identidade**: Atualizada para focar explicitamente na coreografia e protocolos de handoff de personas.

## [1.0.0] - 2026-04-19
### Adicionado
- Estrutura inicial e definição das personas do Swarm.
- Protocolo formal de Handoff.

---

## 🧩 youtube-transcript
## [1.0.0] - 2026-04-15
 — 2026-04-15

### Adicionado
- **SKILL.md**: Definição principal da skill com fluxo de decisão e fallback.
- **README.md**: Documentação de visão geral e uso.
- **Estrutura Base**: Scaffolding inicial seguindo os padrões do Skills Hub.

---
