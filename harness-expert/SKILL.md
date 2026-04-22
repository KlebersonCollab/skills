---
name: harness-expert
version: 2.0.0
description: "Motor técnico para Harness Engineering. Fornece ferramentas para gestão de estado, memória e automação de loops de feedback operando sob o sdd-planner."
category: agentic-infrastructure
---

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar qualquer execução técnica:
1. **Context Check**: Você reidratou o contexto lendo `STATE.md`, `MEMORY.md` e `LEARNINGS.md`?
2. **Spec Check**: O arquivo `spec.md` existe com requisitos e Critérios de Aceitação (ACs) claros? (BDD mandatório para Medium+).
3. **Plan Check**: O arquivo `plan.md` define a arquitetura, schemas e inclui diagramas **Mermaid**?
4. **Contract Check**: O arquivo `contract.md` foi estabelecido com os sensores de validação?
5. **Task Check**: A lista de tarefas em `tasks.md` está detalhada e atomizada?

---
# Harness Expert

> "Se você não é o modelo, então você é o harness." — Infraestrutura de suporte para transformar LLMs em agentes operacionais.

---

## Goal

O objetivo desta skill é fornecer o "maquinário" (CLI, scripts, automação) para que o **SDD Planner** execute suas diretrizes. Enquanto o Planner define o *quê* e o *porquê*, o Harness fornece o *como* técnico para manter o estado determinístico e loops de feedback fechados.

---

## Workflow (4 Fases)

### Fase 1: AUTOMATED REHYDRATE — Carregamento Técnico
1.  **Injeção de Contexto**: Executar scripts de leitura em massa de `.specs/` para reidratar a memória do agente conforme definido pelo `sdd-planner`.
2.  **Health Check**: Validar se os arquivos de estado (`STATE.md`, `tasks.md`) estão em formato válido.
3.  **Context Compression**: Utilizar `harness-expert-compress` para reduzir o tamanho da sessão mantendo os tokens críticos.

### Fase 2: OPERATE — Automação de Execução
1.  **Tool Orchestration**: Garantir que MCPs e scripts locais sejam invocados com parâmetros de segurança.
2.  **Harnessing Sub-agents**: Configurar o ambiente para que subagentes especializados operem sem degradar o estado global.

### Fase 3: AUTO-SYNC — Sincronização de Máquina
1.  **CLI Sync**: Utilizar o **SDD CLI** para atualizar atomicamente o status no `tasks.md`.
2.  **Feedback Loop**: Capturar logs de execução e transcrevê-los para `LEARNINGS.md` se houver falhas recorrentes.

### Fase 4: VALIDATE — Autocorreção e Determinismo
1.  **Deterministic Checks**: Rodar linters, type-checks e testes unitários de forma automatizada (Exit Gates).
2.  **Machine Feedback**: O harness deve interpretar logs de erro e propor a correção imediata, fechando o ciclo de desenvolvimento sem intervenção humana.

---

## Output Structure

A execução desta skill resulta nos seguintes artefatos mandatórios:

| Artefato | Formato | Descrição |
|----------|---------|-----------|
| **Tasks Sync** | `tasks.md` | Lista de tarefas atualizada com o progresso real. |
| **State Snapshot** | `STATE.md` | Resumo do estado atual do sistema e da feature. |
| **Learnings Log** | `LEARNINGS.md` | Registro de problemas resolvidos e lições técnicas. |
| **Validation Report** | `validation-report.md` | Relatório de testes e conformidade do código gerado. |

---

## Quality Rules

- **Persistence First**: O File System é a única fonte da verdade. Nenhuma decisão deve ser tomada sem consultar as specs locais.
- **SDD Compliance**: Todo desenvolvimento deve seguir rigorosamente o Spec-Driven Development.
- **Context Awareness**: Monitorar constantemente o uso de tokens e realizar a "Compactação de Contexto" quando necessário.
- **Deterministic Check**: Toda saída do agente deve passar por ferramentas de verificação (linters/testes) antes de ser considerada final.

## Prohibited

- NUNCA iniciar uma tarefa sem ler as specs existentes.
- NUNCA considerar uma tarefa "concluída" sem atualizar os arquivos de sincronização (`tasks.md`).
- NUNCA ignorar erros de linter ou testes durante a fase de validação.
- NUNCA manter contexto redundante que possa ser resumido no `STATE.md`.
