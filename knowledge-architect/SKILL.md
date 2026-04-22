---
name: knowledge-architect
version: 1.0.0
description: "Arquitetura de conhecimento local via grafos relacionais (Local GraphRAG). Transforma documentação estática em um mapa de conexões para navegação inteligente de contexto."
category: knowledge-management
---

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar qualquer execução técnica:
1. **Context Check**: Você reidratou o contexto lendo `STATE.md`, `MEMORY.md` e `LEARNINGS.md`?
2. **Spec Check**: O arquivo `spec.md` existe com requisitos e Critérios de Aceitação (ACs) claros? (BDD mandatório para Medium+).
3. **Plan Check**: O arquivo `plan.md` define a arquitetura, schemas e inclui diagramas **Mermaid**?
4. **Contract Check**: O arquivo `contract.md` foi estabelecido com os sensores de validação?
5. **Task Check**: A lista de tarefas em `tasks.md` está detalhada e atomizada?

---
# Knowledge Architect

> "O conhecimento não está nos arquivos, mas nas conexões entre eles." — Protocolo para transformar documentação linear em um grafo relacional local.

---

## Goal

O objetivo desta skill é capacitar o agente a construir, manter e navegar em um **Grafo de Conhecimento Local (LKG)**. Ela permite que o agente entenda as relações complexas entre requisitos, decisões arquiteturais (ADRs), componentes de código e estados de sessão, superando as limitações da busca puramente linear ou vetorial.

---

## Workflow (4 Fases)

### Fase 1: EXTRACT — Destilação de Entidades
1.  **Identificação de Nós**: Ler os arquivos do projeto e identificar "entidades" (ex: Features, Componentes, Classes, ADRs, Requisitos).
2.  **Mapeamento de Relações**: Identificar como essas entidades se conectam (ex: `Feature A` **implementada por** `Classe B`, `Decisão C` **impacta** `Módulo D`).
3.  **Metadados Semânticos**: Extrair tags, versões e dependências críticas de cada nó.

### Fase 2: LINK — Construção do Grafo
1.  **Documentação em Mermaid**: Criar ou atualizar o `KNOWLEDGE-MAP.mermaid` (Mapa da Mente do projeto).
2.  **Cross-Referencing**: Inserir links bidirecionais (`[[link]]`) em arquivos Markdown para facilitar o caminhamento manual e automático.
3.  **Indexação de Contexto**: Manter um índice de "vizinhos próximos" para entidades críticas.

### Fase 3: TRAVERSE — Navegação Inteligente
1.  **Caminhamento de Impacto**: Antes de uma alteração, seguir as conexões no grafo para identificar efeitos colaterais em módulos distantes.
2.  **Descoberta de Contexto**: Usar o grafo para decidir quais arquivos *adicionais* devem ser lidos (além do arquivo alvo) para obter a visão completa.
3.  **Poda de Irrelevância**: Ignorar ramos do grafo que não possuem conexão semântica com a tarefa atual, economizando tokens.

### Fase 4: PRUNE — Evolução e Limpeza
1.  **Refatoração de Conhecimento**: Remover conexões obsoletas ou entidades deletadas do código.
2.  **Consolidação de Clusters**: Agrupar nós muito conectados em "Módulos de Conhecimento" para simplificar a visão macro.
3.  **Sync com Harness**: Exportar o mapa de relações para que o `harness-expert` possa realizar a reidratação relacional.

---

## Output Structure

A execução desta skill resulta nos seguintes artefatos mandatórios (geralmente em `.specs/knowledge/` ou na raiz da skill):

| Artefato | Formato | Descrição |
|----------|---------|-----------|
| **Knowledge Map** | `KNOWLEDGE-MAP.mermaid` | Diagrama visual de todas as entidades e relações do projeto. |
| **Relations Registry** | `RELATIONS.md` | Lista textual de conexões e suas justificativas técnicas. |
| **Context Clusters** | `CLUSTERS.json` | (Opcional) Agrupamento de arquivos que devem sempre ser lidos juntos. |

---

## Quality Rules

- **Gold Standard Benchmark**: Utilize os exemplos em `examples/gold-standard/` como referência para mapeamentos complexos.
- **Relation over Content**: Foque em *como* as coisas se conectam, não apenas no que elas fazem.
- **Visual-First**: Todo grafo complexo deve ser visualizável via Mermaid.
- **Bi-directional**: Sempre que possível, as relações devem ser mapeadas nos dois sentidos (quem usa e quem é usado).
- **KISS Graphing**: Evite criar nós para detalhes triviais; foque em componentes arquiteturais e requisitos.

## Prohibited

- NUNCA crie grafos que não possam ser renderizados (erros de sintaxe Mermaid).
- NUNCA mantenha relações órfãs (nós sem conexões relevantes).
- NUNCA use ferramentas externas de banco de dados de grafos; tudo deve ser persistido em arquivos de texto no repositório.
- NUNCA ignore a atualização do mapa após grandes refatorações de código.
