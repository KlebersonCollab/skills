---
name: sdd
version: 1.5.0
description: Spec-Driven Development. Modular workflow with PRD/RFC, BDD, and GAN-style Adversarial Review.
category: development-workflow
---

# SDD: Modular & Adaptive Workflow

> Precision at scale. Rigor when needed, speed when possible.

---

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar qualquer execução técnica:
1. **Context Check**: Você reidratou o contexto lendo `STATE.md`, `MEMORY.md` e `LEARNINGS.md`?
2. **Spec Check**: O arquivo `spec.md` existe com requisitos e Critérios de Aceitação (ACs) claros? (BDD mandatório para Medium+).
3. **Plan Check**: O arquivo `plan.md` define a arquitetura, schemas e inclui diagramas **Mermaid**?
4. **Contract Check**: O arquivo `contract.md` foi estabelecido com os sensores de validação?
5. **Task Check**: A lista de tarefas em `tasks.md` está detalhada e atomizada?

---
## Goal

O SDD tem como objetivo garantir precisão, rastreabilidade e integridade no ciclo de vida de desenvolvimento de software, transformando requisitos ambíguos em código verificado e documentado através de um workflow modular e adaptativo, agora potencializado pelos princípios de **Harness Engineering** e loops adversariais (**GAN-style**).

---

## The Core Principle: Auto-Sizing

The depth of the workflow is determined by the **Complexity** of the task, not a fixed pipeline.

| Level | Scope | Phases | Required Sub-skills | Directory |
|---|---|---|---|---|
| **Quick** | Bug fixes, config, <3 files | (Implement) + (Verify) | `sdd-implementer` | Root |
| **Small** | Clear feature, <5 tasks | (Spec) + Impl + Verify | `sdd-orchestrator`, `sdd-implementer` | `spec/` |
| **Medium** | Feature + UI, <10 tasks | Explorer + Spec (BDD) + Plan + **Contract** + Impl + Verify | `sdd-explorer`, `sdd-orchestrator` | `spec/` |
| **Large** | Multi-component, new module | Planner + Explorer + RFC + Spec + Plan + **Contract** + Impl + Verify | All modules | `.specs/` |
| **Complex** | Ambiguity, high risk | Same as Large + GAN Adversarial Loop + Score Review | All modules + `harness-expert` | `.specs/` |

---

## Workflow (4 Fases)

O ciclo de vida do SDD segue um fluxo iterativo e rigoroso de entrega:

### Fase 1: DISCOVERY — Mapeamento e Contexto
1.  **Explorar o Terreno**: Utilizar `sdd-explorer` para entender a stack e arquitetura atual.
2.  **Definir Visão**: Utilizar `sdd-planner` para alinhar os objetivos no `PROJECT.md` e `ROADMAP.md`.
3.  **Captura de Memória**: Reidratar o contexto lendo `STATE.md`, `MEMORY.md` e `LEARNINGS.md`.
4.  **Requisitos de Produto**: Para tarefas **Complex**, auditar o PRD (Product Requirements Document) existente.

### Fase 2: SPECIFY — Contratos e Planejamento
1.  **Escrever a Spec**: Utilizar `sdd-orchestrator` para definir requisitos técnicos (`spec.md`). Para níveis **Medium+**, utilizar formato **BDD** (Given/When/Then) nos Critérios de Aceitação.
2.  **Proposta Técnica (RFC)**: Para níveis **Large+**, elaborar uma **RFC** (Request for Comments) detalhando a arquitetura e trade-offs antes do `plan.md`.
3.  **Desenhar a Solução**: Elaborar o `plan.md` com arquitetura e schemas. **Mandatório**: Incluir diagramas **Mermaid** (Flowcharts, Sequence, Class) para visualizar fluxos e estruturas em níveis **Medium+**.
4.  **Contrato de Desenvolvimento (SDC)**: Criar o `contract.md` definindo o acordo entre o Implementador e o Revisor sobre o que será entregue e como será medido (Sensores).
5.  **Atomic Tasks**: Gerar a lista de tarefas (`tasks.md`) para execução.

### Fase 3: IMPLEMENT — Execução Atômica
1.  **Automação**: Utilizar o **SDD CLI** (`uv run sdd task <feature> <id>`) para marcar o progresso.
2.  **Ciclo de Código**: Utilizar `sdd-implementer` para escrever código test-driven (TDD) alinhado aos cenários BDD.
2.  **Integridade**: Garantir que cada tarefa em `tasks.md` seja marcada como completa apenas após passar nos testes.

### Fase 4: REVIEW — Auditoria e Finalização
1.  **Adversarial Review (GAN)**: Para tarefas **Large/Complex**, utilizar a skill `harness-expert` para executar o loop Gerador-Avaliador até atingir a nota de corte (Default: 7.0).
2.  **Veredito via Sensores**: Utilizar `sdd-reviewer` para auditar a entrega contra a `spec.md` e o `contract.md`.
3.  **Score Report**: Gerar o relatório de validação (`validation-report.md`) com pontuação de 0-100 e rubrica GAN.
4.  **UAT**: Validar com o usuário se os critérios de aceitação (BDD) foram atendidos.
5.  **Persistência**: Atualizar os logs de memória e estado no Planner antes do encerramento.

---

## The Modular Engine

This skill delegates tasks to specialized sub-skills for maximum scalability:

- **[Explorer](sdd-explorer.skill.md)**: Maps existing codebases (`STACK.md`, `ARCHITECTURE.md`, `CONVENTIONS.md`).
- **[Planner](sdd-planner.skill.md)**: Manages project vision (`ROADMAP.md`) and session memory (`STATE.md`).
- **[Orchestrator](sdd-orchestrator.skill.md)**: Translates requirements (PRD) into technical specifications (`spec.md` com BDD, `plan.md`/`RFC`).
- **[Implementer](sdd-implementer.skill.md)**: Writes atomic, test-driven code and manages git commits.
- **[Reviewer](sdd-reviewer.skill.md)**: Audits implementation against BDD scenarios with evidence.
- **[Harness](harness-expert.skill.md)**: Drives the adversarial quality loop and state management.

---

## The Knowledge Verification Chain

When researching or deciding, follow this strict hierarchy:

1. **Codebase Docs**: Consult `.specs/codebase/` (generated by Explorer).
2. **Project Docs**: Check READMEs, `PROJECT.md`, `ROADMAP.md`.
3. **Existing Code**: Read the source directly for patterns.
4. **Web Search**: Official documentation only.
5. **Flag Uncertainty**: If Steps 1-4 fail, communicate uncertainty to the user.

---

## Operational Protocols

### The Safety Valve
If a task identified as **Quick** or **Small** reveals hidden complexity (>5 steps or deep dependencies), **STOP**. Formalize the task as **Large** by creating a full spec and plan.

### State Management
Utilize the **Planner** (`sdd-planner`) to manage three distinct types of memory:
1.  **Operational Memory (`STATE.md`)**: Tasks, session status, and blockers.
2.  **Persistent Knowledge (`MEMORY.md`)**: Enduring facts, style guides, and user preferences.
3.  **Incremental Wisdom (`LEARNINGS.md`)**: Solutions to bugs and code patterns discovered.

This triad ensures that context persists even if the agent is restarted or the context window shifts.

### Verification Standards
A feature is NOT complete until the **Reviewer** issues an `APPROVED` verdict based on evidence (file paths and line numbers).

---

## Directory Structure (Mandatory)

### Feature-Specific (`spec/` or `.specs/features/[name]/`)
- `spec.md`: Traceable requirements (FR-X) and ACs.
- `plan.md`: Technical design, schemas, and Mermaid diagrams.
- `tasks.md`: Atomic task list with status.

### Project-Wide (`.specs/`)
- `project/`: `PROJECT.md`, `ROADMAP.md`, `STATE.md`, `MEMORY.md`, `LEARNINGS.md`.
- `codebase/`: `STACK.md`, `ARCHITECTURE.md`, `CONVENTIONS.md`, `CONCERNS.md`.

## Output Structure

A execução deste workflow deve resultar nos seguintes artefatos mandatórios, organizados em `.specs/features/[name]/`:

| Artefato | Arquivo | Descrição |
|----------|---------|-----------|
| **Specification** | `spec.md` | Requisitos funcionais, ACs (BDD) e restrições. |
| **Technical Plan** | `plan.md` | Arquitetura, schemas e diagramas Mermaid. |
| **Contract** | `contract.md` | Acordo de entrega e sensores de validação. |
| **Atomic Tasks** | `tasks.md` | Lista detalhada de tarefas com status. |
| **Audit Reports** | `validation-report.md` | Relatórios de conformidade com Score e evidências. |

## Quality Rules

- **Gold Standard Benchmark**: Sempre consulte `examples/gold-standard/` para alinhar a qualidade da entrega ao nível de rigor esperado.
- **Rigor Adaptativo**: O nível de detalhamento deve seguir estritamente a tabela de Auto-Sizing.
- **BDD-First**: Critérios de aceitação para níveis Medium+ devem obrigatoriamente usar Given/When/Then.
- **Verificação Contínua**: Uma tarefa só é considerada concluída após passar nos testes e ser marcada em `tasks.md`.
- **Diagram-as-Code**: Desenhos técnicos devem usar Mermaid integrados ao Markdown.

## Prohibited

- **NUNCA** iniciar a implementação (Fase 3) sem uma `spec.md` aprovada.
- **NUNCA** ignorar a atualização proativa de `STATE.md` e `MEMORY.md`.
- **NUNCA** pular a Fase 4 (Review) e o encerramento das specs ao finalizar uma feature.
- **NUNCA** usar placeholders como "todo" ou "..." em arquivos de especificação finalizados.
