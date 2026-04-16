# GEMINI Project Instructions

Este arquivo contém mandatos e diretrizes fundamentais para o Gemini CLI neste projeto. Estas instruções devem ser seguidas rigorosamente em cada sessão.

## SDD Workflow Mandates

1.  **Proactive Phase 4 (Persistence)**: Sempre que uma tarefa ou feature for concluída tecnicamente, o agente **DEVE** realizar proativamente a Fase 4 do SDD (Review & Persistence) antes de declarar a finalização ao usuário.
    *   **Atualizar Specs**: Marcar obrigatoriamente **TODAS** as tarefas como concluídas no arquivo `tasks.md` da feature, além de atualizar o `plan.md`.
    *   **Relatórios**: Gerar ou atualizar `validation-report.md` e `verification-report.md`.
    *   **Roadmap**: Atualizar o `ROADMAP.md` e `PROJECT.md` se necessário.
    *   **NÃO** declarar uma tarefa como "finalizada" ou "concluída" sem que os arquivos de rastreio (`tasks.md`) reflitam 100% do progresso.

2.  **State Management**: Manter o `STATE.md`, `MEMORY.md` e `LEARNINGS.md` atualizados para garantir a continuidade do contexto entre sessões.

## Architectural & Technical Mandates

3.  **Architecture-First (ADRs)**: Decisões que alterem a estrutura do sistema, introduzam novos padrões ou tecnologias críticas **DEVEM** ser precedidas por uma análise de trade-offs e registradas em um **ADR** (Architecture Decision Record) na pasta `.specs/architecture/`.

4. **Skill Standardization**: A criação ou atualização de qualquer skill deve ser realizada e validada obrigatoriamente através da **Skill Factory**. Nenhuma skill deve ser registrada sem passar pelo \`skill-factory-validator\`.

5. **Onboarding Synchronization**: Sempre que uma skill for adicionada ou tiver sua versão/propósito alterado, o agente **DEVE** obrigatoriamente atualizar o catálogo de habilidades em \`onboarding-navigator/references/skills-catalog.md\` para manter a integridade do ecossistema.

6. **API Governance**: Projetos de API devem priorizar contratos claros (OpenAPI/SDL), segurança (OWASP API Top 10) e resiliência (Rate Limiting/Circuit Breakers), seguindo o **API Style Guide** do projeto.


6. **Tooling Consistency**: Sempre utilize as ferramentas de gerenciamento de versão e ambiente estabelecidas no projeto (ex: `uv` para Python, `fvm` para Flutter, `npm/pnpm` para Node).

7. **Quality-First (Clean Code)**: Todo código gerado ou revisado **DEVE** seguir rigorosamente os princípios **SOLID**, **YAGNI**, **DRY** e **KISS**. O agente deve atuar como um mentor de qualidade, evitando sobre-engenharia e priorizando a simplicidade e manutenibilidade, conforme detalhado na skill `clean-code-mentor`.

8. **Diagram-as-Code (Mermaid)**: Toda especificação técnica ou desenho de arquitetura **DEVE** ser acompanhado por diagramas **Mermaid** integrados ao Markdown. Isso inclui fluxos de dados, diagramas de sequência para processos assíncronos e mapas de componentes.

9. **Observability-First**: Todo novo serviço ou módulo deve nascer com **Logging Estruturado** (JSON) e **SLIs** (Service Level Indicators) definidos desde a fase de especificação, conforme detalhado na skill `observability-expert`.

10. **Navigation & Culture**: Antes de propor mudanças em processos globais ou iniciar novos módulos de grande escala, o agente **DEVE** consultar as diretrizes da skill `onboarding-navigator` para garantir alinhamento com a cultura e o catálogo de padrões do projeto.

## Global Engineering Standards

0. **Session Startup**: No início de cada sessão, o agente **DEVE** consultar a skill `onboarding-navigator` para reidratar o contexto do hub, entender as habilidades disponíveis e validar o alinhamento com a cultura do projeto antes de qualquer ação.

1.  **Proactive Phase 4 (Persistence)**: Sempre que uma tarefa ou feature for concluída tecnicamente, o agente **DEVE** realizar proativamente a Fase 4 do SDD (Review & Persistence) antes de declarar a finalização ao usuário.

---
*Mandato atualizado em 14 de Abril de 2026.*
