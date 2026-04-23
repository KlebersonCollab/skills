---
name: brainstorming
version: 1.1.0
description: "Facilitador de Brainstorming e Design — guia o agente a explorar problemas complexos, gerar ideias divergentes e convergir para especificações sólidas antes de qualquer implementação."
category: design-thinking
---

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar qualquer execução técnica:
0. **Mode Check**: Verificar o modo operacional atual (`.hub-mode`) e aplicar as diretrizes da skill `token-distiller`.
1. **Context Check**: Você reidratou o contexto lendo `STATE.md`, `MEMORY.md` e `LEARNINGS.md`?
2. **Spec Check**: O arquivo `spec.md` existe com requisitos e Critérios de Aceitação (ACs) claros? (BDD mandatório para Medium+).
3. **Plan Check**: O arquivo `plan.md` define a arquitetura, schemas e inclui diagramas **Mermaid**?
4. **Contract Check**: O arquivo `contract.md` foi estabelecido com os sensores de validação?
5. **Task Check**: A lista de tarefas em `tasks.md` está detalhada e atomizada?

---
# Brainstorming

> Facilitador de ideias e resolução de problemas. Pense como um arquiteto e facilitador, não como um executor. O objetivo aqui é "reduzir a velocidade para acertar".

---

## Goal

Capacitar o agente a facilitar sessões de brainstorming de alta qualidade, utilizando técnicas de divergência e convergência para definir problemas, explorar soluções alternativas, avaliar trade-offs e consolidar especificações claras antes de iniciar o desenvolvimento.

---

## Quando Usar Esta Skill

- Antes de iniciar qualquer projeto ou feature complexa.
- Quando o problema não está bem definido.
- Para explorar múltiplas abordagens técnicas antes de escolher uma.
- Quando houver necessidade de alinhar premissas e requisitos não-funcionais.
- Para resolver gargalos criativos ou técnicos.

## Quando NÃO Usar Esta Skill

- Para tarefas triviais ou bugs óbvios.
- Quando a solução já está totalmente especificada e validada.
- Quando você já está em fase de execução/implementação massiva.

---

## Workflow (4 Fases)

### Fase 1: DISCOVER — Definição do Problema
1.  **Questionamento Ativo**: Use a técnica dos "5 Whys" para chegar à raiz do problema.
2.  **Mapeamento de Stakeholders**: Para quem estamos construindo? Quais são suas dores?
3.  **Understanding Lock**: Crie um resumo de entendimento (Objetivos, Non-goals e Premissas) e peça validação antes de prosseguir.

### Fase 2: DIVERGE — Exploração de Ideias
1.  **Multiplicidade**: Gere pelo menos 3 abordagens distintas (ex: Simples, Escalável, Experimental).
2.  **Técnicas Criativas**: Utilize SCAMPER ou First Principles para desafiar o status quo.
3.  **Proibição de Código**: Nesta fase, foque em conceitos, fluxos e arquitetura. **NÃO ESCREVA CÓDIGO DE PRODUÇÃO**.

### Fase 3: CONVERGE — Avaliação e Escolha
1.  **Matriz de Trade-offs**: Liste Prós e Contras de cada abordagem (Complexidade vs Velocidade vs Manutenibilidade).
2.  **Questionamento Dinâmico**: Faça uma pergunta por vez para ajudar o usuário a decidir entre as opções.
3.  **Recomendação**: O agente deve recomendar uma das abordagens baseada no contexto técnico.

### Fase 4: SPECIFY — Consolidação
1.  **Design Specification**: Documente a solução escolhida em Markdown durável.
2.  **Decision Log**: Registre o que foi decidido e por quê.
3.  **Handoff**: Defina os próximos passos claros para a implementação.

---

## Output Structure

A execução desta skill deve resultar nos seguintes artefatos mandatórios, preferencialmente armazenados em `.specs/design/` ou `docs/brainstorming/`:

| Artefato | Formato | Descrição |
|----------|---------|-----------|
| **Design Specification** | `.md` | Documento consolidado com a arquitetura e fluxos da solução escolhida. |
| **Decision Log** | `.md` | Registro histórico das alternativas exploradas e a justificativa da escolha final. |
| **Trade-off Matrix** | Tabela | Comparativo visual entre as abordagens (Simples, Escalável, Experimental). |

---

## Quality Rules

- **Uma Pergunta por Vez**: Nunca sobrecarregue o usuário com múltiplas perguntas em uma única mensagem.
- **Suposições Explícitas**: Se você assumir algo, declare-o claramente como uma premissa.
- **YAGNI (You Ain't Gonna Need It)**: Evite sugerir complexidade desnecessária ou "over-engineering".
- **Facilitação, não Execução**: Recuse-se a codificar até que o design esteja validado no "Understanding Lock".

## Prohibited

- NUNCA inicie a implementação sem validação explícita do design.
- NUNCA assuma requisitos não-funcionais (escalabilidade, segurança) sem perguntar.
- NUNCA ignore os "Non-goals" (o que não faremos).
- NUNCA apresente designs gigantes sem dividi-los em seções digeríveis.

---

## Referências

- [`references/brainstorming-techniques.md`](references/brainstorming-techniques.md) — SCAMPER, 5 Whys, First Principles.
- [`references/decision-frameworks.md`](references/decision-frameworks.md) — Matrizes de decisão e Trade-offs.
