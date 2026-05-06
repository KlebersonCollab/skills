# Exemplo: `pyproject.toml` Completo

Template anotado de `pyproject.toml` para projetos Python modernos com UV.

---

## Template Completo

```toml
# ============================================================
# PROJECT METADATA
# ============================================================
[project]
name = "meu-projeto"
version = "0.1.0"
description = "Descrição concisa do projeto"
readme = "README.md"
license = { text = "MIT" }
requires-python = ">=3.11"
authors = [
    { name = "Seu Nome", email = "seu@email.com" },
]
keywords = ["python", "uv", "exemplo"]
classifiers = [
    "Development Status :: 3 - Alpha",
    "Intended Audience :: Developers",
    "License :: OSI Approved :: MIT License",
    "Programming Language :: Python :: 3",
    "Programming Language :: Python :: 3.11",
    "Programming Language :: Python :: 3.12",
    "Programming Language :: Python :: 3.13",
]

# ============================================================
# DEPENDENCIES
# ============================================================
dependencies = [
    "requests>=2.31.0",
    "pydantic>=2.0.0",
    "click>=8.1.0",
]

[project.optional-dependencies]
dev = [
    "pytest>=7.4.0",
    "pytest-cov>=4.0.0",
    "ruff>=0.8.0",
    "mypy>=1.13.0",
    "pre-commit>=4.0.0",
]
docs = [
    "mkdocs>=1.5.0",
    "mkdocs-material>=9.0.0",
]

# ============================================================
# ENTRY POINTS (CLI commands)
# ============================================================
[project.scripts]
meu-cli = "meu_projeto.cli:main"

# ============================================================
# BUILD SYSTEM
# ============================================================
[build-system]
requires = ["hatchling"]
build-backend = "hatchling.build"

# ============================================================
# TOOL: RUFF (Linter + Formatter)
# ============================================================
[tool.ruff]
line-length = 88
target-version = "py311"
src = ["src"]

[tool.ruff.lint]
select = [
    "E",    # pycodestyle errors
    "W",    # pycodestyle warnings
    "F",    # pyflakes
    "I",    # isort
    "N",    # pep8-naming
    "UP",   # pyupgrade
    "B",    # flake8-bugbear
    "SIM",  # flake8-simplify
    "TCH",  # flake8-type-checking
]
ignore = [
    "E501",  # line too long (handled by formatter)
]

[tool.ruff.lint.isort]
known-first-party = ["meu_projeto"]

[tool.ruff.format]
quote-style = "double"
indent-style = "space"

# ============================================================
# TOOL: MYPY (Type Checker)
# ============================================================
[tool.mypy]
python_version = "3.11"
strict = true
warn_return_any = true
warn_unused_configs = true
disallow_untyped_defs = true

[[tool.mypy.overrides]]
module = ["tests.*"]
disallow_untyped_defs = false

# ============================================================
# TOOL: PYTEST (Testing)
# ============================================================
[tool.pytest.ini_options]
testpaths = ["tests"]
python_files = ["test_*.py"]
python_classes = ["Test*"]
python_functions = ["test_*"]
addopts = "-xvs --cov=src/ --cov-report=term-missing"

[tool.coverage.run]
source = ["src"]
omit = ["tests/*"]

[tool.coverage.report]
exclude_lines = [
    "pragma: no cover",
    "if TYPE_CHECKING:",
    "if __name__ == .__main__.",
]

# ============================================================
# UV SPECIFIC
# ============================================================
[tool.uv]
dev-dependencies = [
    "pytest>=7.4.0",
    "pytest-cov>=4.0.0",
    "ruff>=0.8.0",
    "mypy>=1.13.0",
]
```

---

## Uso

```bash
# Instalar todas as dependências
uv sync --dev

# Executar testes
uv run pytest

# Verificar qualidade
uv run ruff format --check .
uv run ruff check .
uv run mypy .

# Build
uv build --clear
```
