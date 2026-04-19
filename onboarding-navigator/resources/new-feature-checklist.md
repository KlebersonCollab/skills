# Checklist: Iniciando Nova Funcionalidade

Este checklist garante que toda nova funcionalidade siga rigorosamente a cultura de engenharia e os padrões de qualidade do Hub.

## 🏁 Pré-requisitos
1. [ ] **Validar Alinhamento**: Consultar a skill `onboarding-navigator` para confirmar se a abordagem proposta faz sentido para o ecossistema.
2. [ ] **Identificar Skills**: Listar quais das 18 skills serão necessárias para a implementação (ex: `fastapi-expert`, `architecture`).

## 🏗️ Fase 1: Especificação (SDD)
3. [ ] **Criar PRD/RFC**: Definir o "porquê" e o "quê" no diretório `.specs/features/<feature_name>/`.
4. [ ] **Desenhar Diagramas**: Criar diagramas de sequência ou fluxo usando Mermaid para visualizar a lógica.
5. [ ] **Escrever ADR**: Caso haja uma decisão arquitetural significativa, registrá-la em `.specs/architecture/`.

## 🛠️ Fase 2: Planejamento Técnico
6. [ ] **Definir Tasks**: Criar o arquivo `tasks.md` com a decomposição atômica das atividades.
7. [ ] **Plano de Testes**: Definir como a funcionalidade será validada (Unitários, Integração, E2E).

## 💻 Fase 3: Execução
8. [ ] **Scaffolding**: Usar `scaffolding-expert` se houver necessidade de novos diretórios ou estruturas padronizadas.
9. [ ] **Implementação Iterativa**: Seguir o ciclo Plan -> Act -> Validate para cada tarefa do `tasks.md`.

## 🚀 Fase 4: Persistência e Review
10. [ ] **Atualizar Grafo de Conhecimento**: Mapear as novas entidades no `KNOWLEDGE-MAP.mermaid` via `knowledge-architect`.
