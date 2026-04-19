# ADR 002: Dynamic Scaffolding Engine Selection

## Status
Accepted

## Context
O hub de skills cresceu para incluir diversos ecossistemas (Python, FastAPI, Django, Flutter). A criação de novos projetos exige muito esforço manual dos agentes para escrever arquivos base (boilerplates). Precisamos de uma ferramenta de scaffolding que seja:
1. Fácil de executar via CLI (sem instalação global pesada).
2. Suporte templates parametrizáveis (Jinja2).
3. Permita atualizações de templates em projetos existentes.

## Decision
Adotaremos o **Copier** como engine oficial de scaffolding.

## Rationale
- **Zero Install**: Pode ser executado via `uvx copier` (rápido e isolado).
- **Flexibilidade**: Suporta Git repos, pastas locais e composição de templates.
- **Python-Native**: Alinhado com o core de automação do Hub (scripts em Python).

## Consequences
- Todas as skills que oferecem criação de projetos devem ter uma pasta `template/` compatível com Copier.
- Os agentes devem ser instruídos a usar o comando `uvx copier copy <template_path> <dest>` como primeiro passo em novos apps.
