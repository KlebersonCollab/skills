# GEMINI Project Instructions

<!-- GLOBAL_MANDATES_START -->

# Global Engineering Mandates

Estas instruções são fundamentais e devem ser seguidas por todos os agentes operando neste Hub.

## 0. Session Startup
No início de cada sessão, o agente **DEVE** consultar a skill `onboarding-navigator` para reidratar o contexto do hub, entender as habilidades disponíveis e validar o alinhamento com a cultura do projeto antes de qualquer ação.

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



<!-- GLOBAL_MANDATES_END -->

--- 
*Mandato atualizado em 16 de Abril de 2026.*