---
name: sdd
version: 1.5.0
description: Spec-Driven Development. Modular workflow with PRD/RFC, BDD, and Mermaid Diagrams mandate.
category: development-workflow
---

# SDD: Modular & Adaptive Workflow

> Precision at scale. Rigor when needed, speed when possible.

---

## Goal

O SDD tem como objetivo garantir precisão, rastreabilidade e integridade no ciclo de vida de desenvolvimento de software, transformando requisitos ambíguos em código verificado e documentado através de um workflow modular e adaptativo, agora potencializado pelos princípios de **Harness Engineering**.

---

## The Core Principle: Auto-Sizing

The depth of the workflow is determined by the **Complexity** of the task, not a fixed pipeline.

| Level | Scope | Phases | Required Sub-skills | Directory |
|---|---|---|---|---|
| **Quick** | Bug fixes, config, <3 files | (Implement) + (Verify) | `sdd-implementer` | Root |
| **Small** | Clear feature, <5 tasks | (Spec) + Impl + Verify | `sdd-orchestrator`, `sdd-implementer` | `spec/` |
| **Medium** | Feature + UI, <10 tasks | Explorer + Spec (BDD) + Plan + **Contract** + Impl + Verify | `sdd-explorer`, `sdd-orchestrator` | `spec/` |
| **Large** | Multi-component, new module | Planner + Explorer + RFC + Spec + Plan + **Contract** + Impl + Verify | All modules | `.specs/` |
| **Complex** | Ambiguity, high risk | Same as Large + PRD Audit + Score Review | All modules + `sdd-reviewer` | `.specs/` |

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
1.  **Isolamento**: **OBRIGATÓRIO** criar uma feature branch (`feat/`, `fix/`, `docs/`) a partir da `main` antes de iniciar o código. NUNCA commite diretamente na `main`.
2.  **Automação**: Utilizar o **SDD CLI** (`uv run sdd task <feature> <id>`) para marcar o progresso.
3.  **Ciclo de Código**: Utilizar `sdd-implementer` para escrever código test-driven (TDD) alinhado aos cenários BDD.
4.  **Mandato Proibitivo**: **NUNCA** marque uma tarefa como completa em `tasks.md` sem passar 100% pelos testes e sem realizar o commit git correspondente na branch de feature.
5.  **Integridade**: Garantir que cada tarefa em `tasks.md` seja marcada como completa apenas após passar nos testes e ser commitada individualmente.

### Fase 4: REVIEW — Auditoria e Finalização
1.  **Veredito via Sensores**: Utilizar `sdd-reviewer` para auditar a entrega contra a `spec.md` e o `contract.md`, utilizando feedback de ferramentas (testes, linters).
2.  **Score Report**: Gerar o relatório de validação com pontuação de 0-100.
3.  **UAT**: Validar com o usuário se os critérios de aceitação (BDD) foram atendidos.
4.  **Persistência**: Atualizar os logs de memória e estado no Planner antes do encerramento.

---

## The Modular Engine: When, How & Why

Esta skill delega tarefas para sub-skills especializadas para garantir escalabilidade e eficiência de contexto:

| Sub-skill | **Quando** usar? (Trigger) | **Como** delegar? | **Por que** é vital? |
|---|---|---|---|
| **[Explorer](sdd-explorer.skill.md)** | Ao iniciar em código legado ou desconhecido. | Peça um mapeamento da stack e arquitetura (`STACK.md`, `ARCHITECTURE.md`). | Evita alucinações sobre dependências e padrões existentes. |
| **[Planner](sdd-planner.skill.md)** | Ao final de sessões, após decisões ou mudanças de roadmap. | Instrua a atualizar o `STATE.md`, `MEMORY.md` e capturar `LEARNINGS.md`. | Garante a persistência do estado e a continuidade entre sessões (Harness). |
| **[Orchestrator](sdd-orchestrator.skill.md)** | Sempre que uma feature (**Small+**) for iniciada. | Delegue a criação de `spec.md`, `plan.md` e a quebra em `tasks.md`. | Mantém a separação entre "o que fazer" e "como fazer", garantindo rastreabilidade. |
| **[Implementer](sdd-implementer.skill.md)** | Quando o plano e as tarefas atômicas estão prontos. | Delegue a execução de tarefas específicas do `tasks.md` via TDD e commits. | Foca o agente na escrita de código limpo e testes sem a carga cognitiva do design. |
| **[Reviewer](sdd-reviewer.skill.md)** | Quando a implementação (Fase 3) termina ou em marcos críticos. | Peça uma auditoria baseada em evidências contra os ACs (BDD) e o `contract.md`. | Garante que a entrega atende ao padrão de qualidade e critérios originais. |

> [!TIP]
> Em cenários de **Alta Complexidade**, o Orquestrador deve atuar como o "Cérebro Central", mantendo apenas o plano e delegando a pesquisa (Explorer) e a codificação (Implementer) para manter sua própria janela de contexto limpa e focada.

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

### Delegação para Sub-Agentes (Sub-Agent Delegation)
Para maximizar a performance e não estourar a janela de contexto, delegue tarefas pesadas para sub-agentes (quando suportado):
- **Pesquisa/Brownfield:** Delegue e exija apenas o resumo.
- **Implementação e Testes:** Um sub-agente foca no código, edita os arquivos e retorna o status.
- Nunca delegue o *Planejamento* ou a *Criação de Tarefas*; o Orquestrador sempre deve manter o contexto consolidado.

### Verification Standards
A feature is NOT complete until the **Reviewer** issues an `APPROVED` verdict based on evidence (file paths and line numbers).

---

## Referências & Políticas

Para manter a conformidade rigorosa e lidar com situações específicas, consulte as diretrizes detalhadas (agregadas do padrão TLC Spec-Driven):

- **[Brownfield Mapping](references/brownfield-mapping.md)**: Como auditar e mapear repositórios existentes (STACK, ARCHITECTURE, CONCERNS).
- **[Princípios de Codificação](references/coding-principles.md)**: Simplicidade, testes íntegros e foco objetivo.
- **[Limites de Contexto](references/context-limits.md)**: Tamanhos máximos permitidos para artefatos e zonas de alerta de tokens.
- **[Session Handoff](references/session-handoff.md)**: Protocolo obrigatório para pausar e retomar o trabalho sem perder contexto.
- **[Quick Mode](references/quick-mode.md)**: Via expressa para tarefas triviais (<3 arquivos) que não requerem a sobrecarga de todas as fases.
- **[BDD Guide](references/bdd-guide.md)**: Padrão para escrita de especificações de aceitação.

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
- **Verificação Contínua**: Uma tarefa só é considerada concluída após passar 100% nos testes e ser commitada individualmente seguindo a `git-workflow`.
- **Atomicidade Mandatória**: Cada entrada `[x]` no `tasks.md` deve corresponder a pelo menos um commit no Git com a mensagem vinculada ao ID da task.
- **Branch-First**: O desenvolvimento deve sempre ocorrer em branches de curta duração, protegendo a branch principal (`main`).
- **Diagram-as-Code**: Desenhos técnicos devem usar Mermaid integrados ao Markdown.

## Prohibited

- **NUNCA** iniciar a implementação (Fase 3) sem uma `spec.md` aprovada.
- **NUNCA** ignorar a atualização proativa de `STATE.md` e `MEMORY.md`.
- **NUNCA** pular a Fase 4 (Review) e o encerramento das specs ao finalizar uma feature.
- **NUNCA** usar placeholders como "todo" ou "..." em arquivos de especificação finalizados.
