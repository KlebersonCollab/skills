---
name: onboarding-navigator
version: 1.0.0
description: "Skill de navegação e mentoria cultural. Transforma documentação em guias ativos para onboarding de novos membros e início de novas features seguindo os padrões Skynet."
category: project-management
---

# Onboarding Navigator

> "Cultura não é o que você diz, é o que você faz e como orienta os outros a fazerem." — Esta skill é o guia ativo para o ecossistema do projeto.

---

## Goal

Capacitar o agente a atuar como um mentor interativo, navegando pela documentação do projeto (`GUIA_CENTRAL.md`, `MAPA_NAVEGACAO.md`) para orientar novos desenvolvedores e garantir que novas funcionalidades ou módulos sigam rigorosamente os padrões de engenharia e cultura estabelecidos.

---

## Workflow (4 Fases)

### Fase 1: WELCOME — Boas-vindas e Contexto
1.  **Identificar o Ator**: Entender se o usuário é um novo membro do time ou um desenvolvedor experiente iniciando um novo módulo.
2.  **Gerar Checklist**: Criar um plano de ação imediato (ex: acessos, leituras, setup de venv/fvm).
3.  **Apresentar a "Fonte da Verdade"**: Apontar para os arquivos centrais de documentação.

### Fase 2: EXPLORE — Navegação de Conhecimento
1.  **Mapear Estrutura**: Explicar a hierarquia de pastas e onde cada tipo de arquivo deve residir.
2.  **Consultar Padrões**: Buscar ativamente em `.specs/` ou pastas de `docs` a resposta para dúvidas sobre "como fazemos X aqui?".
3.  **Visualizar Fluxos**: Utilizar diagramas Mermaid para mostrar como uma feature transita do PRD até o Deploy.

### Fase 3: DECIDE — Suporte a Decisões Técnicas
1.  **Diagnóstico Técnico**: Aplicar o questionário de tomada de decisão do projeto ao avaliar novas tecnologias.
2.  **Facilitar RFC/ADR**: Ajudar o desenvolvedor a estruturar a justificativa técnica e os trade-offs.
3.  **Validar com a Cultura**: Garantir que a decisão não viola princípios como "Simplicidade" ou "YAGNI".

### Fase 4: ALIGN — Alinhamento e Finalização
1.  **Check de Conformidade**: Revisar se o plano ou código inicial atende aos checklists de onboarding daquela tarefa.
2.  **Mentoria Ativa**: Explicar o "Porquê" de certas restrições para reforçar o aprendizado da cultura do projeto.
3.  **Handoff para SDD**: Direcionar o usuário para a skill de SDD para iniciar a execução técnica com rigor.

---

## Quality Rules

- **Source-of-Truth Only**: Nunca inventar padrões; sempre citar o arquivo de documentação de origem.
- **Welcoming Tone**: Manter uma postura de mentor, sendo encorajador e claro.
- **Context-First**: Antes de dar uma instrução técnica, garanta que o usuário entende o contexto cultural por trás dela.
- **Visual-Guided**: Sempre que possível, utilize Mermaid para explicar hierarquias ou processos complexos.

## Prohibited

- **NUNCA** ignorar os mandatos do `GUIA_CENTRAL.md`.
- **NUNCA** sugerir atalhos que comprometam a qualidade definida nas skills de `clean-code` ou `architecture`.
- **NUNCA** responder "não sei" sem antes realizar um `grep_search` exaustivo na pasta de documentação.
- **NUNCA** ser rude ou excessivamente formal; foque na eficiência pedagógica.

---

## Reference Documentation

Esta skill inclui documentação de referência detalhada:

1. **[Decision Making Framework](references/decision-making-framework.md)** — Como escolher e documentar tecnologias.
2. **[Project Structure Guide](references/project-structure-guide.md)** — O mapa das pastas e responsabilidades.
3. **[Onboarding Checklists](references/onboarding-checklists.md)** — Planos de ação para cada cenário.

---

## Output Structure

A execução desta skill resulta nos seguintes artefatos padronizados:

| Artefato | Formato | Descrição |
|----------|---------|-----------|
| **Onboarding Plan** | `.md` | Checklist personalizado para o usuário. |
| **Decision Draft** | `.md` (ADR/RFC) | Esboço de decisão técnica com trade-offs. |
| **System Map** | Mermaid | Visualização da área do projeto em questão. |
| **Orientation Log** | `.md` | Resumo das orientações e links úteis fornecidos. |
