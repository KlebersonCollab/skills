# Estratégias de Ramificação (Branching)

Escolher a estratégia correta depende do tamanho da equipe, cadência de release e maturidade de automação.

## 1. GitHub Flow
**Melhor para**: SaaS, aplicações web, equipes pequenas e entrega contínua.

### Fluxo:
1.  **Branch main**: Representa o estado de produção. Deve estar sempre estável.
2.  **Feature Branches**: Qualquer mudança (nova feature ou correção) deve ser feita em uma branch criada a partir da `main`.
    - Nomeação: `feature/<desc>` ou `fix/<desc>`.
3.  **Pull Request (PR)**: Quando a branch está pronta, abre-se um PR para revisão.
4.  **Discussão e Revisão**: Feedback da equipe e ajustes.
5.  **Merge**: Após aprovação e aprovação do CI, a branch é integrada à `main`.
6.  **Deploy**: O deploy é feito imediatamente após o merge.

---

## 2. Trunk-Based Development
**Melhor para**: Equipes seniores, alta velocidade de desenvolvimento, arquitetura de microsserviços.

### Fluxo:
- Desenvolvedores trabalham em uma única branch principal (`trunk` ou `main`).
- As branches são de vida curtíssima (horas ou 1-2 dias).
- **Feature Flags**: Funcionalidades incompletas são enviadas para produção, mas escondidas atrás de chaves de configuração.
- **Integração Contínua**: O código é integrado várias vezes ao dia.

---

## 3. GitFlow
**Melhor para**: Projetos legados, ambientes regulados, releases agendadas (ex: Apps Mobile em lojas).

### Fluxo:
- **main**: Código em produção.
- **develop**: Branch de integração para a próxima release.
- **feature/**: Criadas a partir da `develop`, integradas de volta na `develop`.
- **release/**: Criadas a partir da `develop` para preparação final de release. Fundem-se na `main` e `develop`.
- **hotfix/**: Criadas a partir da `main` para correções críticas imediatas. Fundem-se na `main` e `develop`.

---

## Tabela Comparativa

| Critério | GitHub Flow | Trunk-Based | GitFlow |
|----------|-------------|-------------|---------|
| **Complexidade** | Baixa | Média | Alta |
| **Velocidade** | Alta | Altíssima | Moderada |
| **Risco de Conflito** | Baixo | Moderado | Alto |
| **Controle de Release** | Contínuo | Contínuo | Agendado |
