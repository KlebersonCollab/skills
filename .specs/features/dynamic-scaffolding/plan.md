# Plan: Dynamic Scaffolding

## Goal
Implementar um sistema de geração de templates dinâmicos (scaffolding) para acelerar a criação de projetos e reduzir o boilerplate manual, integrando o `uvx copier` como engine padrão no Hub de Skills.

## Architectural Decision
- **Engine**: [ADR-002] Copier (via `uvx`).
- **Storage**: Pasta `templates/` dentro de cada skill que suportar scaffolding.
- **Trigger**: Instrução explícita no `SKILL.md` orientando o agente a usar comandos CLI específicos.

## Phases
1. **Infra**: Definir estrutura de pastas globais e templates base.
2. **Integration**: Atualizar as skills `python-uv`, `fastapi-expert` e `django-expert` com comandos de scaffolding.
3. **Skill Creation**: Criar ou atualizar a skill `scaffolding-expert` para gerenciar a descoberta de templates.
4. **Validation**: Testar a geração de um projeto FastAPI completo via template.

## Verification Strategy
- Executar `uvx copier copy` localmente e validar se a estrutura gerada segue os padrões de Clean Code do Hub.
- Validar se o `KNOWLEDGE-MAP.mermaid` reflete a nova dependência de scaffolding.
