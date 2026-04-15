# GEMINI Project Instructions

Este arquivo contém mandatos e diretrizes fundamentais para o Gemini CLI neste projeto. Estas instruções devem ser seguidas rigorosamente em cada sessão.

## SDD Workflow Mandates

1.  **Proactive Phase 4 (Persistence)**: Sempre que uma tarefa ou feature for concluída tecnicamente, o agente **DEVE** realizar proativamente a Fase 4 do SDD (Review & Persistence).
    *   **Atualizar Specs**: Marcar tarefas como concluídas em `tasks.md` e `plan.md`.
    *   **Relatórios**: Gerar ou atualizar `validation-report.md` e `verification-report.md`.
    *   **Roadmap**: Atualizar o `ROADMAP.md` e `PROJECT.md` se necessário.
    *   **NÃO** esperar por solicitação explícita do usuário para estas atualizações de encerramento.

2.  **State Management**: Manter o `STATE.md`, `MEMORY.md` e `LEARNINGS.md` atualizados para garantir a continuidade do contexto entre sessões.

## Architectural & Technical Mandates

3.  **Architecture-First (ADRs)**: Decisões que alterem a estrutura do sistema, introduzam novos padrões ou tecnologias críticas **DEVEM** ser precedidas por uma análise de trade-offs e registradas em um **ADR** (Architecture Decision Record) na pasta `.specs/architecture/`.

4.  **Skill Standardization**: A criação de qualquer nova skill deve ser realizada e validada obrigatoriamente através da **Skill Factory**. Nenhuma skill deve ser registrada sem passar pelo `skill-factory-validator`.

5.  **API Governance**: Projetos de API devem priorizar contratos claros (OpenAPI/SDL), segurança (OWASP API Top 10) e resiliência (Rate Limiting/Circuit Breakers), seguindo o **API Style Guide** do projeto.

6. **Tooling Consistency**: Sempre utilize as ferramentas de gerenciamento de versão e ambiente estabelecidas no projeto (ex: `uv` para Python, `fvm` para Flutter, `npm/pnpm` para Node).

7. **Quality-First (Clean Code)**: Todo código gerado ou revisado **DEVE** seguir rigorosamente os princípios **SOLID**, **YAGNI**, **DRY** e **KISS**. O agente deve atuar como um mentor de qualidade, evitando sobre-engenharia e priorizando a simplicidade e manutenibilidade, conforme detalhado na skill `clean-code-mentor`.

## General Engineering Standards


- **Context Efficiency**: Combine buscas e leituras para minimizar turnos. Utilize sub-agentes para tarefas de alto volume.
- **Security**: Nunca exponha, logue ou commite credenciais, segredos ou chaves de API.
- **Git**: Seguir as convenções de commit do projeto ao finalizar tarefas.

---
*Mandato atualizado em 14 de Abril de 2026.*
