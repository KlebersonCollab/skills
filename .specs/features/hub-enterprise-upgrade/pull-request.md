## Descrição
Release V4.0.0 que finaliza o épico do **Hub Enterprise Upgrade**. Esta release eleva a maturidade do repositório garantindo que todas as 21 skills do Hub estejam 100% aderentes ao Spec-Driven Development (SDD), incluindo proteções automatizadas e verificações adversariais (GAN-style).

## Mudanças e Entregas (Metas 1 a 6)
- **Meta 1 (Test Suite)**: Cobertura base via `pytest` para blindar scripts da raiz.
- **Meta 2 (Knowledge Distiller)**: Automação completa do Grafo de Conhecimento relacional no formato Mermaid, lendo do *frontmatter* das skills.
- **Meta 3 (DevSecOps Skill)**: Inauguração da trilha de auditoria e Shift-Left Security.
- **Meta 4 (Installer CLI)**: Motor de instalação local ou remota (via git sparse-checkout).
- **Meta 5 (Dashboard)**: Lançamento do UI Marketplace construído em Streamlit.
- **Meta 6 (Bunker Release Strategy)**: Implementação do mandato `make release-check` para garantir que novas versões só entrem após linting do *Ruff*, formatações e passagens de testes (`100% verde`).

## Checklist de Qualidade (SDD Exit Gate)
- [x] O `STATE.md` e `LEARNINGS.md` foram atualizados?
- [x] `make release-check` executado com zero falhas?
- [x] O CHANGELOG-HUB.md reflete as alterações mais recentes?
- [x] A Tag de release (v4.0.0) foi gerada no último commit validado?

## Evidência
- O Hub conta agora com 24 arquivos Python reformatados e auditados pelo Ruff, além de 11 testes garantindo 100% de estabilidade nos cenários cruciais do framework.
