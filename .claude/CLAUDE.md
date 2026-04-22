# CLAUDE Project Instructions

<!-- GLOBAL_MANDATES_START -->

# Global Engineering Mandates

Estas instruções são fundamentais e devem ser seguidas por todos os agentes operando neste Hub.

## 🔒 0. SESSION BOOTSTRAP (EXECUTE BEFORE ANY RESPONSE)
Antes de responder ao usuário, o agente **DEVE** realizar este checklist mental e operacional:
1. **Rehydrate Context**: Ler `.specs/project/STATE.md`, `MEMORY.md` e `LEARNINGS.md`.
2. **Onboarding Check**: Consultar `onboarding-navigator` para alinhar com a cultura do projeto.
3. **Task Sizing**: Classificar a complexidade da tarefa (Quick, Small, Medium, Large, Complex) conforme a tabela SDD.
4. **Skill Matching**: Consultar o **Skill Router** abaixo para selecionar as ferramentas adequadas.
5. **SDD Verification**: Se a tarefa for de desenvolvimento, validar se `spec.md` e `plan.md` existem.

## 📍 SKILL ROUTER
Utilize este guia para identificar a skill mandatória para cada contexto:

| Se a tarefa envolve... | USE esta skill |
|------------------------|----------------|
| Início de Projeto/Feature | `onboarding-navigator` |
| Especificação e Planejamento | `sdd` (`orchestrator`, `planner`) |
| Desenvolvimento Python (FastAPI/Django) | `python-uv` + `fastapi-expert` / `django-expert` |
| Desenvolvimento Frontend (React/Next) | `frontend-expert` |
| Desenvolvimento Mobile (Flutter) | `flutter-fvm` |
| Desenvolvimento Backend (Go) | `golang-expert` |
| Arquitetura e ADRs | `architecture` + `api-architect` |
| Qualidade e Clean Code | `clean-code-mentor` |
| Observabilidade e SRE | `observability-expert` |
| Brainstorming e Ideação | `brainstorming` |
| Automação e Scaffolding | `scaffolding-expert` |
| Gestão de Conhecimento | `knowledge-architect` |
| Sincronização de Estado | `harness-expert` |
| Gestão de Git e Commits | `git-workflow` |

## 1. SDD Framework (Mandatory for Development)
Qualquer tarefa de construção, desenvolvimento ou refatoração significativa **DEVE** obrigatoriamente utilizar o framework **SDD (Spec-Driven Development)** desde o planejamento inicial.
- **Workflow & Persistence**: Ao concluir tecnicamente uma tarefa, o agente **DEVE** realizar proativamente a Fase 4 do SDD (Review & Persistence).
  - **Atualizar Specs**: Marcar tarefas como concluídas no `tasks.md` e atualizar o `plan.md`.
  - **Relatórios**: Gerar ou atualizar `validation-report.md`.
  - **Estado**: Manter `STATE.md`, `MEMORY.md` e `LEARNINGS.md` atualizados via **SDD CLI**.

## 2. Architectural Integrity
- **Architecture-First**: Decisões críticas devem ser registradas em um **ADR** na pasta `.specs/architecture/`.
- **Skill Standardization**: Novas skills devem ser validadas obrigatoriamente via `skill-factory-validator`.
- **Diagram-as-Code**: Desenhos técnicos devem usar **Mermaid** integrados ao Markdown.

## 3. Quality Standards (Clean Code)
- Seguir rigorosamente **SOLID, YAGNI, DRY e KISS**.
- Priorizar a simplicidade e a manutenibilidade; evitar sobre-engenharia.
- Todo código deve passar por validação automatizada (linters/testes).

## 4. Operational Infrastructure (Harness Expert)
- **State Synchronization**: O agente **DEVE** garantir que o estado operacional (memória de longo prazo) seja sincronizado via `harness-expert` ao final de cada iteração.
- **Context Efficiency**: Utilizar a compactação de contexto proativamente para evitar saturação de tokens, mantendo o `STATE.md` como a bússola do progresso.

## 5. Knowledge Management (Knowledge Architect)
- **Local Knowledge Graph (LKG)**: Toda nova feature ou mudança arquitetural significativa **DEVE** ser mapeada no `KNOWLEDGE-MAP.mermaid`.
- **Relationship Mapping**: Antes de grandes refatorações, o agente **DEVE** utilizar o `knowledge-architect` para analisar o impacto nas entidades e relações existentes.

## 6. Mandatory Skill Usage
- **Skill-Driven Execution**: O agente **DEVE** obrigatoriamente identificar e utilizar a skill mais adequada para cada tarefa. A execução de qualquer atividade técnica sem o suporte de uma skill específica ou do framework de automação do Hub é proibida.

## 7. Dynamic Scaffolding (Scaffolding Expert)
- **Zero-Boilerplate Policy**: O agente **DEVE** preferir utilizar as ferramentas de CLI de templates (como `uvx copier` via `scaffolding-expert`) para inicializar a estrutura de projetos e arquivos complexos (como `pyproject.toml`, Dockerfiles). O código boilerplate NUNCA deve ser escrito linha por linha se existir um template de scaffolding disponível.

## 9. Version Control & Git (Git Workflow)
Todo agente **DEVE** seguir rigorosamente a skill `git-workflow` para qualquer operação de versionamento:
- **English-Only Commits**: Todas as mensagens de commit devem ser escritas em **Inglês**.
- **Conventional Commits**: O uso do padrão Conventional Commits é obrigatório.
- **SDD Linking**: Commits devem ser atômicos e, sempre que possível, referenciar IDs de tarefas do `tasks.md`.
- **History Integrity**: O uso de `rebase` é mandatório para manter um histórico linear antes de merges na `main`.

## 🔒 10. SESSION EXIT GATE (EXECUTE BEFORE ENDING)
Antes de encerrar a sessão ou entregar a tarefa, o agente **DEVE** validar:
1. **Tasks Update**: O `tasks.md` reflete o estado real da implementação?
2. **State Sync**: O `STATE.md` foi atualizado com o progresso e próximos passos?
3. **Learnings Capture**: Novos padrões ou bugs corrigidos foram para o `LEARNINGS.md`?
4. **Validation**: O código passou por linters/testes e recebeu um Score no `validation-report.md`?
5. **Knowledge Update**: O `KNOWLEDGE-MAP.mermaid` precisa de atualização?


<!-- GLOBAL_MANDATES_END -->

--- 
*Mandato atualizado em 16 de Abril de 2026.*