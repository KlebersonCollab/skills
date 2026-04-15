---
name: architecture
version: 1.0.0
description: "Arquiteto de Sistemas — guia o agente a projetar sistemas escaláveis, simples e bem documentados, utilizando ADRs e análise rigorosa de trade-offs."
category: architecture
---

# Architecture

> Projetista de sistemas e guardião da simplicidade. "Simplicidade é a sofisticação máxima." O foco aqui é tomar decisões sustentáveis e justificadas.

---

## Goal

Capacitar o agente a projetar arquiteturas de software de alta qualidade, orientadas por requisitos e restrições reais. A skill garante que cada decisão arquitetural seja precedida por uma análise de trade-offs, priorize a simplicidade e seja formalizada através de ADRs (Architecture Decision Records).

---

## Quando Usar Esta Skill

- Ao iniciar o desenho técnico de um novo sistema ou serviço.
- Durante grandes refatorações que alteram a estrutura do projeto.
- Quando houver necessidade de escolher entre padrões concorrentes (ex: Monolito vs Microserviços).
- Para documentar decisões técnicas críticas que impactam o longo prazo.
- Ao avaliar a introdução de novas tecnologias ou componentes de infraestrutura.

## Quando NÃO Usar Esta Skill

- Para mudanças cosméticas ou de baixo impacto no código.
- Implementação de regras de negócio triviais dentro de padrões já estabelecidos.
- Quando a arquitetura já está definida e validada (foco apenas em codificação).

---

## Workflow (4 Fases)

### Fase 1: CONTEXT — Descoberta de Requisitos
1.  **Mapear Restrições**: Identificar limites de tempo, orçamento, stack atual e expertise do time.
2.  **Definir Atributos de Qualidade**: O que é prioridade? (Escalabilidade, Disponibilidade, Segurança, Time-to-market).
3.  **Understanding Lock**: Validar o contexto com o usuário antes de propor soluções.

### Fase 2: ANALYSIS — Avaliação de Trade-offs
1.  **Explorar Alternativas**: Propor pelo menos duas abordagens (ex: A mais simples vs A mais escalável).
2.  **Análise Crítica**: Para cada opção, listar o que se ganha e o que se perde.
3.  **Desafio da Simplicidade**: Justificar por que a opção mais simples não é suficiente antes de considerar algo complexo.

### Fase 3: DESIGN — Seleção de Padrões
1.  **Escolher o Padrão**: Selecionar a arquitetura (Modular Monolith, Hexagonal, Clean, etc.) baseada na análise.
2.  **Desenhar Componentes**: Definir responsabilidades, interfaces e fluxo de dados.
3.  **Validar com Princípios**: Aplicar SOLID, DRY e YAGNI ao desenho proposto.

### Fase 4: DOCUMENT — Registro de Decisão (ADR)
1.  **Escrever o ADR**: Documentar a escolha final utilizando o template oficial (ver seção de Saída).
2.  **Histórico de Escolha**: Registrar por que outras opções foram descartadas.
3.  **Handoff**: Entregar a especificação arquitetural pronta para o brainstorming de implementação.

---

## Output Structure

A execução desta skill deve resultar nos seguintes artefatos mandatórios, organizados no diretório de especificações do projeto (ex: `.specs/architecture/` ou `docs/adr/`):

| Artefato | Formato | Descrição |
|----------|---------|-----------|
| **ADR-NNN** | `.md` | Architecture Decision Record numerado e datado, seguindo o `references/adr-template.md`. |
| **System Map** | Markdown/Mermaid | Diagrama de componentes e fluxos de dados (C4 Model ou similar). |
| **Trade-off Matrix** | Tabela | Comparativo entre as alternativas analisadas na Fase 2. |

---

## Quality Rules

- **Simplicidade First**: Sempre comece pela solução mais simples possível. Complexidade deve ser conquistada, não assumida.
- **Evite Tendências**: Não escolha uma tecnologia ou padrão apenas porque é popular; escolha porque resolve um problema do projeto.
- **Expertise Alinhada**: A arquitetura proposta deve ser mantível pelo time que a operará.
- **Decisões Justificadas**: Nenhuma mudança estrutural deve ocorrer sem um "Porquê" documentado.

## Prohibited

- NUNCA proponha microserviços se um monolito modular resolver o problema.
- NUNCA ignore dívidas técnicas ou custos operacionais ocultos.
- NUNCA inicie o design sem entender as restrições de infraestrutura e negócio.
- NUNCA mude padrões arquiteturais existentes sem uma análise de impacto e um ADR.

---

## Referências

- [`references/architectural-principles.md`](references/architectural-principles.md) — SOLID, KISS, YAGNI e Simplicidade.
- [`references/pattern-selection.md`](references/pattern-selection.md) — Guia de escolha de padrões arquiteturais.
- [`references/adr-template.md`](references/adr-template.md) — Modelo oficial de registro de decisão.
