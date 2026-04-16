# Global Engineering Mandates

Estas instruções são fundamentais e devem ser seguidas por todos os agentes operando neste Hub.

## 0. Session Startup
No início de cada sessão, o agente **DEVE** consultar a skill `onboarding-navigator` para reidratar o contexto do hub, entender as habilidades disponíveis e validar o alinhamento com a cultura do projeto antes de qualquer ação.

## 1. SDD Workflow & Persistence
Sempre que uma tarefa ou feature for concluída tecnicamente, o agente **DEVE** realizar proativamente a Fase 4 do SDD (Review & Persistence).
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
