# Exemplo: GitHub Actions CI/CD Completo

Workflow de produção com matrix de versões, caching, quality checks e publicação automática.

---

## Workflow CI (Validação em PRs e Pushes)

```yaml
name: CI

on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main]

concurrency:
  group: ${{ github.workflow }}-${{ github.ref }}
  cancel-in-progress: true

jobs:
  # ────────────────────────────────────────────
  # Quality Checks
  # ────────────────────────────────────────────
  quality:
    name: Quality Checks
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

      - name: Check formatting
        run: uv run ruff format --check .

      - name: Check linting
        run: uv run ruff check .

      - name: Type check
        run: uv run mypy .

  # ────────────────────────────────────────────
  # Tests com matrix de versões e OS
  # ────────────────────────────────────────────
  test:
    name: Test (Python ${{ matrix.python-version }}, ${{ matrix.os }})
    needs: quality
    runs-on: ${{ matrix.os }}
    strategy:
      fail-fast: false
      matrix:
        os: [ubuntu-latest, macos-latest, windows-latest]
        python-version: ["3.11", "3.12", "3.13"]

    steps:
      - uses: actions/checkout@v4

      - name: Install UV
        uses: astral-sh/setup-uv@v4
        with:
          enable-cache: true

      - name: Set up Python ${{ matrix.python-version }}
        run: uv python install ${{ matrix.python-version }}

      - name: Install dependencies
        run: uv sync --dev --locked

      - name: Run tests
        run: uv run pytest --cov=src/ --cov-report=xml

      - name: Upload coverage (Ubuntu only)
        if: matrix.os == 'ubuntu-latest' && matrix.python-version == '3.12'
        uses: codecov/codecov-action@v4
        with:
          file: coverage.xml
          token: ${{ secrets.CODECOV_TOKEN }}
```

---

## Workflow CD (Publicação ao criar Release)

```yaml
name: Publish

on:
  release:
    types: [published]

jobs:
  publish:
    name: Build & Publish
    runs-on: ubuntu-latest
    permissions:
      id-token: write  # Para trusted publishing do PyPI

    steps:
      - uses: actions/checkout@v4

      - name: Install UV
        uses: astral-sh/setup-uv@v4

      - name: Set up Python
        run: uv python install 3.12

      - name: Install dependencies
        run: uv sync --locked

      - name: Build distribution
        run: uv build --clear

      - name: Publish to PyPI
        run: uv publish
        env:
          UV_PUBLISH_TOKEN: ${{ secrets.PYPI_TOKEN }}

      - name: Upload artifacts
        uses: actions/upload-artifact@v4
        with:
          name: dist
          path: dist/
```

---

## Dicas

### Cache eficiente
O `astral-sh/setup-uv@v4` com `enable-cache: true` cacheia automaticamente pacotes entre runs, reduzindo o tempo de CI significativamente.

### `--locked` é essencial
Sempre use `uv sync --locked` em CI para garantir reprodução exata. Se o lockfile estiver desatualizado, o CI falhará — forçando atualização explícita.

### Matrix otimizada
Quality checks rodam uma vez (não precisam de matrix). Testes pesados rodam em matrix para validação cross-platform.
