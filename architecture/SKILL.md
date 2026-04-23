---
name: architecture
version: 2.0.1
description: "Arquiteto de Sistemas — guia o agente a projetar sistemas escaláveis, resilientes e distribuídos (CQRS, Event-Driven) utilizando ADRs, Fitness Functions e Diagramas Mermaid mandatórios para visualização."
category: architecture
---

# Architecture (v2.0.1)

> Projetista de sistemas de alta performance e guardião da simplicidade evolutiva. "Arquitetura é o que resta quando você tira todo o código."

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
## Goal

Capacitar o agente a projetar arquiteturas de software de alta qualidade e complexidade, abrangendo sistemas distribuídos (CQRS, Event-Driven) e monolitos modulares. A skill garante que cada decisão seja justificada por trade-offs, priorize a simplicidade e seja protegida por **Fitness Functions** e documentada visualmente via **Mermaid**.

---

## Quando Usar Esta Skill

- Ao projetar sistemas que exigem alta escalabilidade de leitura/escrita (CQRS).
- Ao desenhar fluxos de trabalho assíncronos e desacoplados (Event-Driven).
- Durante grandes refatorações ou migrações de arquitetura (ex: Monolito para Microsserviços).
- Para definir regras de governança automatizada via Fitness Functions.
- Ao documentar decisões técnicas críticas que impactam o longo prazo (ADRs).

---

## Workflow (4 Fases)

### Fase 1: CONTEXT — Descoberta de Requisitos
1.  **Mapear Restrições**: Identificar limites de tempo, orçamento e stack atual.
2.  **Identificar Padrões de Carga**: Diferenciar volumes de leitura vs. escrita (indicativo de CQRS).
3.  **Necessidade de Desacoplamento**: Avaliar se a comunicação síncrona é um gargalo (indicativo de Event-Driven).

### Fase 2: ANALYSIS — Avaliação de Trade-offs
1.  **Explorar Alternativas**: Comparar abordagens (ex: Simples vs. Escalável).
2.  **Análise de Consistência**: Avaliar se o negócio tolera Consistência Eventual ou exige Consistência Forte.
3.  **Análise de Simplicidade**: Justificar a introdução de padrões complexos (CQRS/Events) apenas se o requisito de escala exigir.

### Fase 3: DESIGN — Seleção de Padrões e Contratos
1.  **Modelagem de Comandos/Consultas**: Se usar CQRS, definir claramente os modelos de Read e Write.
2.  **Design de Eventos**: Definir schemas de mensagens, estratégias de idempotência e DLQs (Dead Letter Queues).
3.  **Desenhar Componentes**: Aplicar SOLID, DRY e YAGNI ao desenho proposto.

### Fase 4: DOCUMENT — Registro e Governança (ADR & Fitness)
1.  **Escrever o ADR**: Documentar a escolha técnica utilizando o template oficial.
2.  **Definir Fitness Functions**: Propor scripts ou testes automatizados para validar se a arquitetura proposta não sofrerá regressões (ex: evitar acoplamento circular).
3.  **Handoff**: Entregar a especificação arquitetural pronta para implementação.

---

## Output Structure

A execução desta skill resulta nos seguintes artefatos mandatórios em `.specs/architecture/`:

| Artefato | Formato | Descrição |
|----------|---------|-----------|
| **ADR-NNN** | `.md` | Architecture Decision Record com justificativa e impacto. |
| **System Map** | **Mermaid** | Diagrama renderizável de componentes e fluxos. **Mandatório**. |
| **Fitness Specs** | `.py` / `.sh` | Definição de testes automatizados para governança da arquitetura. |
| **Trade-off Matrix**| Tabela | Comparativo entre as alternativas analisadas. |

---

## Quality Rules

- **Simplicidade First**: Não use CQRS ou Eventos se um banco de dados relacional simples resolver o problema.
- **Idempotência Obrigatória**: Todo design orientado a eventos deve prever processamento repetido sem efeitos colaterais.
- **Fitness-Driven**: Cada restrição arquitetural importante deve ter uma forma de ser validada automaticamente.
- **Visual-First**: Todo componente ou fluxo complexo deve ser documentado com Mermaid no `System Map`.
- **Decisões Justificadas**: Mudanças estruturais exigem um "Porquê" documentado.

## Prohibited

- NUNCA proponha sistemas distribuídos sem uma análise de custo operacional e latência.
- NUNCA use eventos para comunicação que exige resposta imediata (síncrona).
- NUNCA ignore a complexidade de gerenciar consistência eventual no front-end.
- NUNCA inicie o design sem entender os requisitos de throughput e disponibilidade.

---

## Referências

- [`references/architectural-principles.md`](references/architectural-principles.md) — SOLID, KISS, YAGNI.
- [`references/cqrs-and-events.md`](references/cqrs-and-events.md) — (Novo) Design de comandos e eventos.
- [`references/evolutionary-architecture.md`](references/evolutionary-architecture.md) — (Novo) Fitness Functions e evolução.
- [`references/adr-template.md`](references/adr-template.md) — Modelo oficial de ADR.
