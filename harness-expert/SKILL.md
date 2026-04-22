---
name: harness-expert
version: 2.1.0
description: "Motor técnico para Harness Engineering. Implementa loops de feedback adversariais (GAN-style) para garantir qualidade de produção, além de gestão de estado e determinismo."
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
# Harness Expert (v2.1.0)

> "Se você não é o modelo, então você é o harness." — Infraestrutura de suporte para transformar LLMs em agentes operacionais de alta performance.

---

## Goal

O objetivo desta skill é fornecer o "maquinário" (CLI, scripts, automação) para que o **SDD Planner** execute suas diretrizes. Enquanto o Planner define o *quê* e o *porquê*, o Harness fornece o *como* técnico para manter o estado determinístico e fechar loops de feedback adversariais (GAN-style) que garantem que a entrega final seja **production-ready**.

---

## 🔄 GAN-style Feedback Loop (Quality Enforcement)

Para tarefas de alta complexidade (**Large/Complex**), o Harness deve operar em modo **Adversarial**:

| Papel | Responsabilidade | Ferramentas |
|-------|------------------|-------------|
| **Generator** | Implementar a feature conforme a Spec e o Plan. | SDD CLI, UV, Frameworks (React, Flutter). |
| **Evaluator** | Atuar como QA implacável, testando a aplicação viva e rejeitando "AI slop". | Playwright, Screenshotting, Log Analysis. |

### 📊 Evaluation Rubric (Target: Score >= 7.0)

| Critério | Peso | Descrição |
|----------|------|-----------|
| **Design Quality** | 0.3 | Coesão visual, ausência de padrões genéricos de IA, estética premium. |
| **Originality** | 0.2 | Decisões técnicas/visuais criativas e customizadas para o problema. |
| **Craft** | 0.3 | Polimento, micro-interações, tipografia, espaçamento e performance. |
| **Functionality** | 0.2 | Funcionamento robusto de todas as features e tratamento de erros. |

---

## Workflow (4 Fases)

### Fase 1: AUTOMATED REHYDRATE — Carregamento Técnico
1.  **Injeção de Contexto**: Executar scripts de leitura em massa de `.specs/` para reidratar a memória do agente.
2.  **Health Check**: Validar se os arquivos de estado (`STATE.md`, `tasks.md`) estão íntegros.
3.  **Context Compression**: Utilizar `harness-expert-compress` para manter os tokens críticos.

### Fase 2: OPERATE — Automação de Execução (Generator Mode)
1.  **Tool Orchestration**: Invocar MCPs e scripts locais conforme o plano.
2.  **Implementation Loop**: Escrever código focado em passar nos ACs da `spec.md`.
3.  **Self-Correction**: Corrigir bugs de linter/testes imediatamente via `sdd-implementer`.

### Fase 3: ADVERSARIAL REVIEW — Ciclo de Feedback (Evaluator Mode)
1.  **Live Testing**: Iniciar o servidor de dev e utilizar Playwright para testar fluxos reais.
2.  **Strict Scoring**: Atribuir notas de 1 a 10 baseadas na rubrica. **NUNCA elogiar trabalho medíocre**.
3.  **Feedback Iteration**: Gerar o `validation-report.md` com críticas acionáveis para o Generator.

### Fase 4: FINAL SYNC — Determinismo e Encerramento
1.  **Exit Gates**: Garantir que a pontuação GAN atingiu o limiar de aprovação (Default: 7.0).
2.  **Auto-Sync**: Atualizar `STATE.md`, `MEMORY.md` e `LEARNINGS.md`.
3.  **Handoff**: Preparar os artefatos de distribuição em `dist/`.

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
