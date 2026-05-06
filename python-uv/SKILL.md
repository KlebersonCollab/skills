---
name: python-uv
description: Expert em ecossistema Python moderno com UV, focado em performance (Django/Async), segurança e governança SDD.
version: 4.0.0
tags: [python, uv, django, async, performance, devops]
---

# Python UV Expert

Esta skill governa o desenvolvimento de aplicações Python de alta performance utilizando o `uv` como motor de execução e gestão. Ela integra padrões avançados de Django, desenvolvimento assíncrono e conformidade rigorosa com o framework **SDD (Spec-Driven Development)**.

## 🔒 Prerequisites (Mandatory)

Este skill opera DENTRO do framework **SDD**. Antes de iniciar qualquer execução técnica:
0. **Mode Check**: Verifique o modo operacional atual (`.hub-mode`) e aplique as diretrizes do `token-distiller`.
1. **Context Check**: Você reidratou o contexto lendo `STATE.md`, `MEMORY.md` e `LEARNINGS.md` em `.specs/project/`?
2. **Spec Check**: O arquivo `spec.md` existe com requisitos claros e Critérios de Aceite (ACs)? (BDD obrigatório para Medium+).
3. **Plan Check**: O arquivo `plan.md` define arquitetura, esquemas e inclui diagramas **Mermaid**?
4. **Contract Check**: O arquivo `contract.md` foi estabelecido com sensores de validação?
5. **Task Check**: A lista de tarefas em `tasks.md` está detalhada e atomizada?

---

## 🚀 Core Knowledge: UV & Environment Management

O `uv` é a "Source of Truth" para o ambiente. Substitui pip, poetry, pyenv e pipx com performance Rust.

### 1. Gestão de Versões (Python 3.14+)
- **Novo Padrão**: A partir do UV 0.9.6, o padrão é o **Python 3.14**.
- **Free-Threading (PEP 703)**: Suporte nativo ao Python 3.14t (no-GIL) para paralelismo real em multi-core.
- **Comandos**:
    - `uv python install 3.14` (Instalação)
    - `uv python pin 3.14` (Fixa versão no projeto via `.python-version`)
    - `uv run --python 3.14t script.py` (Execução em modo free-threaded)

### 2. Gestão de Projetos e Dependências
- **Inicialização**: `uv init project-name` (Estrutura `src/` layout recomendada).
- **Sincronização**: `uv sync --dev --locked` (Garante paridade total com `uv.lock`).
- **Adição de Pacotes**: 
    - `uv add requests` (Produção)
    - `uv add --dev ruff mypy pytest` (Desenvolvimento)
    - `uv add --optional group-name package` (Dependências opcionais)

### 3. PEP 723: Scripts Inline
Para automações e ferramentas de arquivo único, use metadados inline:
```python
# /// script
# requires-python = ">=3.11"
# dependencies = ["httpx", "rich"]
# ///
import httpx
# ... lógica aqui
```
Execute com: `uv run script.py`.

---

## 🏗️ Expert Domain: Django Professional

Focado em escalabilidade corporativa e eliminação de gargalos comuns.

### 1. Performance ORM (N+1 Resolution)
- **Regra de Ouro**: NUNCA realize queries em loops.
- **select_related**: Use para Foreign Keys e One-to-One (SQL JOIN).
- **prefetch_related**: Use para Many-to-Many e relações reversas (Queries separadas com IN).
- **Bulk Ops**: Use `bulk_create` e `bulk_update` para operações em massa.

### 2. Service Layer & Architecture
- **Fat Models/Views são Proibidos**: Mova a lógica de negócio complexa para `services.py`.
- **Secure Querysets**: Sempre filtre por `owner` ou `tenant` no `get_queryset` para evitar **BOLA/IDOR**.

### 3. Async Django (ASGI)
- Use views assíncronas (`async def`) para operações de I/O bloqueante (chamadas de API externas).
- Motor recomendado: `uvicorn` via `uv run`.

---

## ⚡ Expert Domain: Async Mastery

### 1. Padrões de Concorrência
- **gather**: Execute múltiplas tarefas simultâneas.
- **Semaphores**: Limite a concorrência para proteger recursos externos.
```python
sem = asyncio.Semaphore(10)
async with sem:
    await call_external_api()
```
- **uvloop**: Utilize para ganhos de performance de 2-4x em Linux/macOS.

### 2. Testing Async
- Configure `pytest-asyncio` no `pyproject.toml` com `asyncio_mode = "auto"`.

---

## 🛠️ Operational Workflows

### 1. Pipeline de Qualidade (TDD)
Siga o ciclo **Red-Green-Refactor**:
1. `uv add --dev pytest pytest-cov`
2. Escreva o teste em `tests/`.
3. Execute: `uv run pytest`.
4. Implemente o código mínimo.
5. Refatore.

### 2. CI/CD & Docker
- **GitHub Actions**: Use `astral-sh/setup-uv@v4` com `enable-cache: true`.
- **Docker**: Utilize builds multi-stage para manter imagens leves, copiando apenas o `.venv`.

---

## 🔒 Prohibitive Mandates & Anti-Patterns

- **NUNCA** use `pip install` globalmente. Use `uv tool install` para ferramentas ou `uv venv` para projetos.
- **NUNCA** ignore o arquivo `uv.lock`. Ele deve ser versionado no Git.
- **NUNCA** execute comandos de gerenciamento sem o prefixo `uv run` (ex: `uv run python manage.py migrate`).
- **NUNCA** use argumentos padrão mutáveis (`list`, `dict`) em funções.
- **EVITE** o uso de `type()` para checagem; prefira `isinstance()`.

---

## 📚 References & Resources

- [Project Management](file:///Users/klebersonromero/Projetos/Kleberson/skills/python-uv/references/project-management.md)
- [Django Workflow](file:///Users/klebersonromero/Projetos/Kleberson/skills/python-uv/references/django-workflow.md)
- [Async Development](file:///Users/klebersonromero/Projetos/Kleberson/skills/python-uv/references/async-development.md)
- [Python Environment (3.14)](file:///Users/klebersonromero/Projetos/Kleberson/skills/python-uv/references/python-environment.md)
- [Testing Standards](file:///Users/klebersonromero/Projetos/Kleberson/skills/python-uv/references/testing.md)
- [CI/CD & Docker](file:///Users/klebersonromero/Projetos/Kleberson/skills/python-uv/references/ci-cd-workflows.md)
- [PEP 723 Metadata](file:///Users/klebersonromero/Projetos/Kleberson/skills/python-uv/references/inline-script-metadata.md)
