# Project Management

Guia completo para inicialização, gerenciamento de dependências, lockfiles, estrutura de projeto e ferramentas de qualidade com UV.

---

## Inicialização de Projeto

### Criar novo projeto

```bash
uv init meu-projeto
cd meu-projeto
```

Estrutura gerada:

```
meu-projeto/
├── .python-version      # Versão Python fixada
├── pyproject.toml       # Configuração do projeto
├── README.md
└── src/
    └── meu_projeto/
        └── __init__.py
```

### Inicializar projeto existente

```bash
cd projeto-existente
uv init --pyproject       # Gera pyproject.toml
```

---

## pyproject.toml

O `pyproject.toml` é a **fonte única da verdade** para o projeto:

```toml
[project]
name = "meu-projeto"
version = "0.1.0"
description = "Meu projeto Python"
requires-python = ">=3.11"
dependencies = [
    "requests>=2.31.0",
    "pydantic>=2.0.0",
    "click>=8.1.0",
]

[project.optional-dependencies]
dev = [
    "pytest>=7.4.0",
    "ruff>=0.1.0",
    "mypy>=1.5.0",
    "pytest-cov>=4.0.0",
]

[project.scripts]
meu-cli = "meu_projeto.cli:main"

[build-system]
requires = ["hatchling"]
build-backend = "hatchling.build"

[tool.ruff]
line-length = 88
target-version = "py311"

[tool.ruff.format]
quote-style = "double"

[tool.mypy]
python_version = "3.11"
strict = true
warn_return_any = true
warn_unused_configs = true

[tool.pytest.ini_options]
testpaths = ["tests"]
python_files = ["test_*.py"]
addopts = "-xvs"
```

Veja [pyproject-toml-example.md](../examples/pyproject-toml-example.md) para um template completo anotado.

---

## Gerenciamento de Dependências

### Adicionar

```bash
# Dependência de produção
uv add requests pydantic

# Dependência de desenvolvimento
uv add --dev pytest ruff mypy

# Dependência com versão específica
uv add "requests>=2.31.0,<3.0.0"

# Dependência opcional (grupo)
uv add --optional test pytest-cov
```

### Remover

```bash
uv remove requests
uv remove --dev pytest-cov
```

### Sincronizar

```bash
# Instalar tudo do pyproject.toml
uv sync

# Incluindo deps de desenvolvimento
uv sync --dev

# A partir do lockfile (reprodução exata)
uv sync --locked

# Recriar ambiente do zero
uv sync --recreate
```

### Inspecionar

```bash
# Árvore de dependências
uv tree

# Dependências reversas (quem depende de quê)
uv tree --reverse

# Listar pacotes instalados
uv pip list
```

---

## Lockfile (`uv.lock`)

O lockfile garante **reprodução exata** do ambiente:

```bash
# Gerar/atualizar lockfile
uv lock

# Sincronizar a partir do lockfile
uv sync --locked
```

**Regras:**
- ✅ **Sempre versione** `uv.lock` no git
- ✅ **Use `--locked`** em CI/CD para builds reproduzíveis
- ❌ **Nunca edite** `uv.lock` manualmente

---

## Estrutura de Projeto Recomendada

### Src Layout (Recomendado)

```
meu-projeto/
├── pyproject.toml
├── uv.lock
├── README.md
├── .gitignore
├── .python-version
├── src/
│   └── meu_projeto/
│       ├── __init__.py
│       ├── core.py
│       ├── models.py
│       └── utils.py
├── tests/
│   ├── __init__.py
│   ├── conftest.py
│   ├── test_core.py
│   └── test_models.py
└── scripts/
    └── deploy.py
```

### .gitignore para projetos UV

```gitignore
# Python
__pycache__/
*.py[cod]
*.egg-info/
dist/
build/

# Virtual Environment
.venv/

# IDE
.idea/
.vscode/
*.swp

# NÃO ignore uv.lock — sempre versione!
# uv.lock  ← NÃO faça isso
```

---

## Workspaces (Monorepo)

```toml
# pyproject.toml na raiz do monorepo
[tool.uv.workspace]
members = ["packages/*"]
```

```
monorepo/
├── pyproject.toml          # Workspace root
├── uv.lock                 # Lockfile compartilhado
└── packages/
    ├── core/
    │   ├── pyproject.toml
    │   └── src/
    ├── api/
    │   ├── pyproject.toml
    │   └── src/
    └── cli/
        ├── pyproject.toml
        └── src/
```

---

## Ferramentas de Qualidade

### Ruff (Linter + Formatter)

```bash
# Instalar como dev dep
uv add --dev ruff

# Formatar código
uv run ruff format .

# Verificar linting
uv run ruff check .

# Corrigir automaticamente
uv run ruff check --fix .
```

### mypy (Type Checking)

```bash
uv add --dev mypy

# Executar type checking
uv run mypy .
```

### pytest (Testes)

```bash
uv add --dev pytest pytest-cov

# Executar testes
uv run pytest

# Com cobertura
uv run pytest --cov=src/

# Testes específicos
uv run pytest tests/test_core.py -xvs
```

### Workflow de qualidade completo

```bash
# Executar toda pipeline de qualidade
uv run ruff format --check .
uv run ruff check .
uv run mypy .
uv run pytest --cov=src/
```
