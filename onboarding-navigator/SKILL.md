---
name: onboarding-navigator
version: 1.5.0
description: "Guia mestre do Hub de Skills. Fornece overview de todas as habilidades locais, mentoria cultural e suporte à navegação técnica no ecossistema do projeto, incluindo fluxos avançados de Harness Engineering."
category: project-management
---

# Onboarding Navigator

> "Para navegar com precisão, você precisa conhecer o mapa." — Esta skill é o guia autoritativo para todas as 25 habilidades e padrões deste repositório.

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

Atuar como o ponto de entrada principal e guia contínuo para o Hub de Skills. A missão é orientar o agente e o usuário sobre qual habilidade utilizar para cada problema, garantir o alinhamento com a cultura de engenharia (KISS, SDD) e facilitar a navegação por toda a documentação local.

## Output Structure

A execução desta skill resulta nos seguintes artefatos:

| Artefato | Formato | Descrição |
|----------|---------|-----------|
| **Skill Roadmap** | `.md` | Sugestão de sequência de skills para resolver o problema. |
| **Decision Draft** | `.md` (ADR/RFC) | Esboço fundamentado de decisão técnica. |
| **Checklist** | `.md` | Plano de passos iniciais para onboarding ou feature. |
| **Ecosystem Map** | `.md` com Mermaid | Diagrama visual do ecossistema de skills. |

## Quality Rules

- **Local-First**: Priorizar sempre a documentação contida neste repositório e o catálogo de skills.
- **Decision-Support**: Nunca apenas "escolher", mas guiar o processo de escolha através de perguntas diagnósticas.
- **Pedagogical**: Explicar o motivo de cada padrão ou skill recomendada.
- **Mermaid-Enabled**: Utilizar visualização visual para explicar fluxos de onboarding ou tomada de decisão.
- **Stats-Aware**: Sempre mencionar o número total de skills (25) e suas categorias.
- **Token-Efficient**: Sempre verificar o modo de operação atual (`.hub-mode`) e aplicar a densidade de tokens apropriada via `token-distiller`.

## Prohibited

- **NUNCA** ignorar a existência de uma skill já implementada ao sugerir uma solução.
- **NUNCA** referenciar arquivos externos inexistentes no repositório local sem contexto.
- **NUNCA** iniciar uma tarefa de larga escala sem validar o alinhamento cultural nesta skill.
- **NUNCA** esquecer de consultar `MEMORY.md` e `LEARNINGS.md` para lições anteriores.

---

## 📊 Visão Geral do Ecossistema

```mermaid
flowchart TD
    Start[Início da Sessão] --> Onboarding{onboarding-navigator}
    
    Onboarding --> Analise[Análise do Contexto]
    Analise --> Catálogo[Consulta ao Catálogo de 25 Skills]
    Catálogo --> Recomendação[Recomendação Específica de 25 Skills]
    Recomendação --> Handoff[Handoff para Skill Especializada]
    
    style Start fill:#e1f5fe
    style Onboarding fill:#81c784
```

---

## 🔄 Workflow (4 Fases)

### Fase 1: 🎪 WELCOME — Boas-vindas e Mapeamento
1.  **Reconhecer o Terreno**: Identificar o estado atual do repositório e os objetivos da sessão.
2.  **Apresentar o Hub**: Utilizar o `references/skills-catalog.md` para dar um overview das **25 habilidades disponíveis** com diagramas Mermaid.
3.  **Boot de Sessão (Token Optimization)**: Ler o arquivo `.hub-mode` e ativar o modo correspondente via `token-distiller`. Sugerir o modo **Low Token** para tarefas simples detectadas no bootstrap.
4.  **Checklist de Início**: Sugerir as primeiras ações baseadas na necessidade do usuário.

### Fase 2: 🔍 EXPLORE — Catálogo de Habilidades
1.  **Match de Necessidade**: Recomendar a skill correta baseada no contexto técnico, usando a **Matriz de Decisão**.
2.  **Guia de Navegação**: Explicar a estrutura da pasta `.specs/` e onde reside a memória do projeto (STATE.md, MEMORY.md, LEARNINGS.md).
3.  **Visualização**: Gerar diagramas Mermaid para explicar a hierarquia de módulos e fluxos de trabalho.

### Fase 3: 🤔 DECIDE — Suporte a Decisões
1.  **Diagnóstico Skynet**: Aplicar o framework de decisão para novas tecnologias ou arquiteturas.
2.  **Alinhamento de Valor**: Garantir que as propostas respeitam a simplicidade e o rigor metodológico.
3.  **Estruturação de ADR/RFC**: Auxiliar na criação de registros de decisão seguindo a skill de `architecture`.

### Fase 4: ✅ ALIGN — Consistência Global
1.  **Check de Padrões**: Validar se os planos seguem as normas de Clean Code e a obrigatoriedade do **SDD** para desenvolvimento.
2.  **Mentoria Cultural**: Relembrar os princípios de "Simplicity at Scale" e "Rigor when needed".
3.  **Handoff de Skill**: Delegar a execução para a skill especializada identificada na Fase 2, garantindo que o ciclo comece pelo **SDD**.
4.  **Session Exit Gate**: Antes de encerrar, validar se cada tarefa marcada como concluída possui seu commit correspondente e se os testes passaram (100%).

---

## 📏 Quality Rules

- **Local-First**: Priorizar sempre a documentação contida neste repositório e o catálogo de skills.
- **Decision-Support**: Nunca apenas "escolher", mas guiar o processo de escolha através de perguntas diagnósticas.
- **Pedagogical**: Explicar o motivo de cada padrão ou skill recomendada.
- **Mermaid-Enabled**: Utilizar visualização visual para explicar fluxos de onboarding ou tomada de decisão.
- **Stats-Aware**: Sempre mencionar o número total de skills (25) e suas categorias.

## 🚫 Prohibited

- **NUNCA** ignorar a existência de uma skill já implementada ao sugerir uma solution.
- **NUNCA** referenciar arquivos externos inexistentes no repositório local sem contexto.
- **NUNCA** iniciar uma tarefa de larga escala sem validar o alinhamento cultural nesta skill.
- **NUNCA** esquecer de consultar `MEMORY.md` e `LEARNINGS.md` para lições anteriores.
- **NUNCA** permitir o início de uma construção ou desenvolvimento sem o framework **SDD**.
- **NUNCA** encerrar uma sessão sem auditar se todos os commits de tarefas (`tasks.md`) foram realizados seguindo a `git-workflow`.

---

## 📚 Reference Documentation

1. **[Skills Catalog](references/skills-catalog.md)** — O mapa autoritativo das 25 habilidades do hub com diagramas Mermaid.
2. **[Decision Making Framework](references/decision-making-framework.md)** — Como escolher e documentar tecnologias.
3. **[Project Structure Guide](references/project-structure-guide.md)** — O mapa das pastas locais.
4. **[Onboarding Checklists](references/onboarding-checklists.md)** — Planos de ação práticos.

---

## 📋 Output Structure

A execução desta skill resulta nos seguintes artefatos:

| Artefato | Formato | Descrição |
|----------|---------|-----------|
| **Skill Roadmap** | `.md` | Sugestão de sequência de skills para resolver o problema. |
| **Decision Draft** | `.md` (ADR/RFC) | Esboço fundamentado de decisão técnica. |
| **Checklist** | `.md` | Plano de passos iniciais para onboarding ou feature. |
| **Ecosystem Map** | `.md` com Mermaid | Diagrama visual do ecossistema de skills. |

---

## 🔄 Fluxo de Atualização

Esta skill **DEVE** ser atualizada sempre que:
1. Nova skill for adicionada ao hub
2. Versões de skills existentes forem atualizadas  
3. Novos padrões de arquitetura forem estabelecidos
4. Lições aprendidas (LEARNINGS.md) forem documentadas
