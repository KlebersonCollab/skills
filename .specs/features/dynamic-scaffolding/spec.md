# PRD: Dynamic Scaffolding

## 1. Contexto e Visão
Atualmente as skills dependem fortemente da capacidade do agente de escrever código boilerplate do zero, baseado em regras descritas em texto. A feature *Dynamic Scaffolding* introduz a capacidade de gerar estruturas de projeto (templates) instantaneamente usando ferramentas modernas e rápidas, economizando tokens e tempo.

## 2. Escopo (In Scope)
- Adoção de uma ferramenta de scaffolding baseada em CLI com zero-install via UV (ex: `uvx copier` ou `uvx cookiecutter`).
- Criação de uma pasta `templates/` (ou `scaffolding/`) dentro das skills que mais se beneficiam disso (ex: `python-uv`, `django-expert`).
- Atualização da skill `skill-factory` para incluir o comando de geração (ou criação de uma `scaffolding-expert`).

## 3. Requisitos Técnicos
1. **Engine**: Utilizar `uvx copier` por ser moderno, Python-first e suportar atualizações de templates.
2. **Integração na Skill**: O `SKILL.md` (por exemplo, do FastAPI) irá instruir o agente: *"Para iniciar um novo projeto, execute: `uvx copier copy gh:KlebersonCollab/skills/fastapi-expert/template .`"*.
3. **Template**: Criar pelo menos 1 template funcional como Prova de Conceito (PoC).

## 4. Próximos Passos (Implementação)
- [ ] Definir a engine oficial (Copier vs Cookiecutter).
- [ ] Criar a estrutura base de templates.
- [ ] Ajustar as instruções do agente para rodar bash commands de scaffolding antes de codar.
