# CI/CD Workflows

Guia completo para integração do UV em pipelines de CI/CD, Docker e publicação de pacotes.

---

## GitHub Actions

### Workflow Básico

```yaml
name: CI

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Install UV
        uses: astral-sh/setup-uv@v4

      - name: Set up Python
        run: uv python install 3.12

      - name: Install dependencies
        run: uv sync --dev --locked

      - name: Check formatting
        run: uv run ruff format --check .

      - name: Check linting
        run: uv run ruff check .

      - name: Type check
        run: uv run mypy .

      - name: Run tests
        run: uv run pytest --cov=src/
```

### Workflow com Matrix

```yaml
name: CI Matrix

on: [push, pull_request]

jobs:
  test:
    runs-on: ${{ matrix.os }}
    strategy:
      matrix:
        os: [ubuntu-latest, macos-latest, windows-latest]
        python-version: ["3.11", "3.12", "3.13"]

    steps:
      - uses: actions/checkout@v4

      - name: Install UV
        uses: astral-sh/setup-uv@v4

      - name: Set up Python ${{ matrix.python-version }}
        run: uv python install ${{ matrix.python-version }}

      - name: Install dependencies
        run: uv sync --dev --locked

      - name: Run tests
        run: uv run pytest
```

### Workflow com Cache

```yaml
name: CI with Cache

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Install UV
        uses: astral-sh/setup-uv@v4
        with:
          enable-cache: true

      - name: Set up Python
        run: uv python install 3.12

      - name: Install dependencies
        run: uv sync --dev --locked

      - name: Run tests
        run: uv run pytest
```

Veja [github-actions-example.md](../examples/github-actions-example.md) para um workflow completo de produção.

---

## Docker

### Dockerfile Básico

```dockerfile
FROM python:3.12-slim

# Instalar UV
COPY --from=ghcr.io/astral-sh/uv:latest /uv /usr/local/bin/uv

WORKDIR /app

# Copiar dependências primeiro (cache de layers)
COPY pyproject.toml uv.lock ./

# Instalar dependências
RUN uv sync --locked --no-dev

# Copiar código fonte
COPY src/ src/

# Executar
CMD ["uv", "run", "python", "-m", "meu_projeto"]
```

### Dockerfile Multi-stage

```dockerfile
# Stage 1: Build
FROM python:3.12-slim AS builder

COPY --from=ghcr.io/astral-sh/uv:latest /uv /usr/local/bin/uv

WORKDIR /app
COPY pyproject.toml uv.lock ./
RUN uv sync --locked --no-dev

# Stage 2: Runtime
FROM python:3.12-slim

COPY --from=builder /app/.venv /app/.venv
COPY src/ /app/src/

ENV PATH="/app/.venv/bin:$PATH"
WORKDIR /app

CMD ["python", "-m", "meu_projeto"]
```

---

## Packaging e Publicação

### Build

```bash
# Build simples
uv build

# Build com limpeza de artefatos antigos (0.9.6+)
uv build --clear

# Verificar artefatos gerados
ls dist/
# meu-projeto-0.1.0.tar.gz
# meu-projeto-0.1.0-py3-none-any.whl
```

### Publicar no PyPI

```bash
# Publicar no PyPI (requer token)
uv publish

# Publicar no TestPyPI
uv publish --publish-url https://test.pypi.org/legacy/
```

### GitHub Actions para Publicação

```yaml
name: Publish

on:
  release:
    types: [published]

jobs:
  publish:
    runs-on: ubuntu-latest
    permissions:
      id-token: write  # Para trusted publishing
    steps:
      - uses: actions/checkout@v4

      - name: Install UV
        uses: astral-sh/setup-uv@v4

      - name: Build
        run: uv build --clear

      - name: Publish to PyPI
        run: uv publish
        env:
          UV_PUBLISH_TOKEN: ${{ secrets.PYPI_TOKEN }}
```

---

## Desenvolvimento Local

### Workflow Diário

```bash
# 1. Clone / pull
git clone <repo>
cd <project>

# 2. Sincronizar ambiente
uv sync --dev

# 3. Pipeline de qualidade
uv run ruff format --check .
uv run ruff check .
uv run mypy .
uv run pytest

# 4. Desenvolver com hot reload (web)
uv run uvicorn app.main:app --reload

# 5. Antes do commit
uv run ruff format .
uv run ruff check --fix .
uv run pytest
```

### Pre-commit com UV

```yaml
# .pre-commit-config.yaml
repos:
  - repo: local
    hooks:
      - id: ruff-format
        name: ruff format
        entry: uv run ruff format
        language: system
        types: [python]

      - id: ruff-check
        name: ruff check
        entry: uv run ruff check --fix
        language: system
        types: [python]

      - id: mypy
        name: mypy
        entry: uv run mypy
        language: system
        types: [python]
```

```bash
uv tool install pre-commit
pre-commit install
```
