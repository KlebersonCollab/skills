---
name: harness-expert
version: 1.0.0
description: "Skill para Harness Engineering. Foca na infraestrutura de estado, memória de longo prazo e orquestração de agentes via SDD (Spec-Driven Development)."
category: development-workflow
---

# Harness Expert

> "Se você não é o modelo, então você é o harness." — Infraestrutura de suporte para transformar LLMs em agentes operacionais.

---

## Goal

O objetivo principal desta skill é fornecer o "maquinário" necessário para que os agentes operem de forma determinística, segura e com memória de longo prazo. Ela padroniza a gestão de estado (STATE.md, tasks.md), a orquestração de subagentes e o fechamento de loops de feedback (testes/linters).

---

## Workflow (4 Fases)

### Fase 1: REHYDRATE — Injeção de Contexto
1.  **Leitura de Specs**: O agente deve começar cada tarefa lendo o `tasks.md` e o `plan.md` da feature atual.
2.  **Identificação de Estado**: Consultar `STATE.md` e `LEARNINGS.md` para entender o progresso e lições aprendidas em sessões anteriores.
3.  **Compactação**: Se o contexto estiver saturado, resumir o progresso atual no `STATE.md` e limpar o histórico irrelevante.

### Fase 2: OPERATE — Execução Orquestrada
1.  **Delegação**: Utilizar subagentes especializados (`codebase_investigator`, `generalist`) para tarefas de alta densidade ou repetitivas.
2.  **Uso de Tools**: Garantir que as ferramentas corretas (Terminal, File System, MCPs) sejam invocadas com parâmetros precisos.

### Fase 3: SYNC — Sincronização de Estado
1.  **Atualização de Tasks**: Utilizar o **SDD CLI** para atualizar o status no `tasks.md` de forma automatizada.
2.  **Registro de Aprendizado**: Erros encontrados e corrigidos devem ser registrados em `LEARNINGS.md`.
3.  **Atualização do Plano**: Se a estratégia mudar, o `plan.md` deve refletir a nova rota.

### Fase 4: VALIDATE — Determinismo e Verificação
1.  **Loops de Feedback**: Rodar linters, type-checks e testes unitários proativamente.
2.  **Autocorreção**: O harness deve usar os logs de erro para guiar o agente na correção imediata, sem intervenção do usuário.

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
