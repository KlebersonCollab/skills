# Specification: Automated Knowledge Distiller

**Feature ID:** FR-KNOWLEDGE-DISTILLER
**Status:** DRAFT
**Complexity:** Medium
**Category:** Documentation / Analysis

## Goal
Transformar o `generate_knowledge_map.py` em um sistema de análise estática profunda que gera diagramas Mermaid dinâmicos, agrupando skills por subgrafos (categorias) e detectando relações de dependência explícitas e implícitas.

## 🎯 Requisitos Funcionais

| ID | Título | Descrição |
|----|--------|-----------|
| FR-1 | **Subgrafos por Categoria** | O diagrama deve agrupar as skills em subgrafos Mermaid baseados no campo `category` do metadata. |
| FR-2 | **Dependências Explícitas** | Suporte a um novo campo `uses: [skill-a, skill-b]` no frontmatter YAML para relações direcionais precisas. |
| FR-3 | **Análise Semântica** | O script deve procurar menções a outras skills não apenas por nome, mas por contexto dentro das seções "Goal" e "Quality Rules". |
| FR-4 | **Badges de Maturidade** | Adicionar ícones/badges baseados no status da skill (ex: 🛡️ para SDD-compliant, 📦 para stable, 🧪 para experimental). |
| FR-5 | **Legenda Dinâmica** | Gerar uma legenda no final do Mermaid explicando as cores e badges. |

## ✅ Critérios de Aceitação (BDD)

### Cenário 1: Agrupamento por Categoria
**Given** que eu tenho skills de "architecture" e "development"
**When** eu executo o Knowledge Distiller
**Then** o arquivo `KNOWLEDGE-MAP.mermaid` deve conter blocos `subgraph Architecture` e `subgraph Development`
**And** as skills correspondentes devem estar dentro de seus respectivos blocos.

### Cenário 2: Detecção de Dependência via Metadata
**Given** que a skill `sdd` tem `uses: [skill-factory]` no frontmatter
**When** o script é executado
**Then** deve haver uma flecha direcionada: `sdd --(uses)--> skill_factory`.

## 🛠️ Restrições Técnicas
- Deve ser retrocompatível com o formato atual.
- Usar `scripts/utils.py` para leitura de metadata.
- Saída: `KNOWLEDGE-MAP.mermaid` na raiz.

## 📊 Matriz de Rastreabilidade
- **FR-1** -> Mermaid `subgraph` syntax.
- **FR-2** -> YAML metadata parsing upgrade.
