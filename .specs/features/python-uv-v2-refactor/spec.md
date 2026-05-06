# Spec: Python UV Expert (v2.2.0 Purist Refactor)

## Context
O objetivo é criar uma skill "Expert" para Python utilizando o `uv`, consolidando conhecimentos avançados de Django Professional, Desenvolvimento Assíncrono, Performance e Gestão de Ambiente Moderno (Python 3.14). A skill deve seguir rigorosamente os padrões SDD v2.2.0 (Purista).

## Requirements (Functional)

- **FR-1: Metadata & Identity**: Definir frontmatter YAML completo seguindo o padrão `skill-factory`.
- **FR-2: SDD Hook**: Incluir a seção obrigatória `🔒 Prerequisites (Mandatory)` vinculada ao workflow SDD.
- **FR-3: Core Knowledge Integration**:
    - **UV Mastery**: Instalação, setup, gestão de projetos (`uv init`, `uv sync`), e lockfiles.
    - **Django Pro**: Padrões de Service Layer, resolução de N+1 (select_related/prefetch_related), Bulk operations e Secure Querysets (anti-BOLA/IDOR).
    - **Async Expert**: Padrões para `httpx`, `uvloop`, `gather` e semáforos.
    - **Python 3.14**: Suporte a Free-threading (no-GIL) e mudança de versão padrão.
    - **PEP 723**: Execução de scripts inline com dependências.
- **FR-4: Advanced Tooling**: Gestão de ferramentas globais via `uv tool` vs `uvx`.
- **FR-5: Quality & CI/CD**: Padrões de TDD (Red-Green-Refactor), Ruff, Mypy e Workflows de GitHub Actions/Docker.
- **FR-6: Prohibitive Mandates**: Listar anti-padrões e comportamentos proibidos (ex: `pip install` global, ignorar `uv.lock`).

## Acceptance Criteria (BDD)

### Scenario 1: SDD Compliance
**Given** que sou um agente operando no Hub
**When** eu leio `python-uv/SKILL.md`
**Then** eu devo encontrar a seção `🔒 Prerequisites (Mandatory)` com os checks de Spec, Plan, Contract e Task.

### Scenario 2: Django Pro Knowledge
**Given** um projeto Django legado com problemas de performance
**When** eu consulto a skill `python-uv`
**Then** ela deve me instruir a usar `select_related` para FKs e `prefetch_related` para M2M, além de sugerir auditoria de N+1.

### Scenario 3: Async & Modern Python
**Given** a necessidade de paralelismo real em Python 3.14
**When** eu utilizo a skill `python-uv`
**Then** ela deve recomendar o uso de Python 3.14t (free-threaded) e explicar como configurar o script via PEP 723.

## Constraints
- **Language**: Português Brasileiro (PT-BR).
- **Structure**: Markdown purista, focado em lógica e protocolos, sem scripts "mágicos" ocultos.
- **Memory**: Não deve conter arquivos locais de memória (`STATE.md`, `MEMORY.md`, etc).
