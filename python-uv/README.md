# Python com UV

> Skill para uso e boas práticas de Python utilizando UV como gerenciador de pacotes e versionamento

[![Versão](https://img.shields.io/badge/Versão-1.0.0-blue)](#changelog)

---

## 📖 Visão Geral

**UV** é um gerenciador de pacotes Python extremamente rápido desenvolvido pela [Astral](https://astral.sh/) (criadores do Ruff). Ele combina as funcionalidades de múltiplas ferramentas em uma única CLI:

- 📦 **Gerenciador de pacotes** (substitui pip)
- 🌐 **Gerenciador de ambientes** (substitui venv/virtualenv)
- 🐍 **Gerenciador de versões Python** (substitui pyenv)
- 🔒 **Gerador de lockfiles** (semelhante a poetry/pipenv)
- 🏗️ **Builder de projetos** (para packaging e publicação)

Esta skill ensina como usar o UV efetivamente seguindo **boas práticas de Python moderno**.

---

## 🚀 Instalação do UV

### macOS / Linux
```bash
# Instalar via curl
curl -LsSf https://astral.sh/uv/install.sh | sh

# Ou via Homebrew (macOS)
brew install uv

# Ou via pipx (se já tiver Python)
pipx install uv
```

### Windows
```powershell
# Via PowerShell
powershell -c "irm https://astral.sh/uv/install.ps1 | iex"

# Ou via pip (se já tiver Python)
pip install uv
```

### Verificar instalação
```bash
uv --version
```

---

## ⚙️ Configuração Inicial

### Inicializar um novo projeto
```bash
# Criar novo projeto Python
mkdir meu-projeto
cd meu-projeto
uv init

# Ou inicializar projeto existente
uv init --pyproject
```

O comando `uv init` cria:
- `pyproject.toml` (configuração do projeto)
- `README.md` (template)
- `.gitignore` (para Python)
- Estrutura básica de diretórios

### Configurar versão de Python
```bash
# Usar Python 3.12 (uv baixa automaticamente se necessário)
uv python pin 3.12

# Listar versões Python disponíveis
uv python list

# Instalar versão específica
uv python install 3.11
```

---

## 📦 Gerenciamento de Dependências

### pyproject.toml
O `pyproject.toml` é a fonte única da verdade para seu projeto:

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
    "black>=23.0.0",
]
test = ["pytest>=7.4.0"]
```

### Adicionar dependências
```bash
# Dependência de produção
uv add requests pydantic

# Dependência de desenvolvimento
uv add --dev pytest ruff mypy

# Dependência opcional
uv add --optional dev pytest-cov
```

### Remover dependências
```bash
uv remove requests
```

### Sincronizar ambiente
```bash
# Instalar todas as dependências do pyproject.toml
uv sync

# Instalar com dependências de desenvolvimento
uv sync --dev

# Forçar recriação do ambiente virtual
uv sync --recreate
```

### Lockfile (uv.lock)
```bash
# Gerar/atualizar lockfile
uv lock

# Sincronizar a partir do lockfile (reprodução exata)
uv sync --locked
```

---

## 🧪 Boas Práticas Python Moderno

### 1. Type Hints (Sempre!)
```python
from typing import Optional, List, Dict
from pydantic import BaseModel

class User(BaseModel):
    id: int
    name: str
    email: Optional[str] = None
    tags: List[str] = []

def process_users(users: List[User]) -> Dict[int, str]:
    return {user.id: user.name for user in users}
```

### 2. Formatação com Ruff
```bash
# Instalar ruff
uv add --dev ruff

# Formatar código
ruff format .

# Verificar linting
ruff check .

# Corrigir automaticamente
ruff check --fix .
```

Configurar no `pyproject.toml`:
```toml
[tool.ruff]
line-length = 88
target-version = "py311"

[tool.ruff.format]
quote-style = "double"
```

### 3. Type Checking com mypy
```bash
# Instalar mypy
uv add --dev mypy

# Executar type checking
uv run mypy .
```

Configurar no `pyproject.toml`:
```toml
[tool.mypy]
python_version = "3.11"
strict = true
warn_return_any = true
warn_unused_configs = true
```

### 4. Testes com pytest
```bash
# Instalar pytest
uv add --dev pytest pytest-cov

# Executar testes
uv run pytest

# Com cobertura
uv run pytest --cov=.

# Testes específicos
uv run pytest tests/test_user.py -xvs
```

### 5. Estrutura de Projeto Recomendada
```
meu-projeto/
├── pyproject.toml
├── uv.lock
├── README.md
├── .gitignore
├── src/
│   └── meu_projeto/
│       ├── __init__.py
│       ├── core.py
│       └── utils.py
├── tests/
│   ├── __init__.py
│   ├── test_core.py
│   └── conftest.py
└── scripts/
    └── deploy.py
```

---

## 🔄 Workflows

### Desenvolvimento Local
```bash
# 1. Clonar repositório
git clone <repo>
cd <project>

# 2. Configurar ambiente
uv sync --dev

# 3. Executar verificação de qualidade
uv run ruff check .
uv run mypy .
uv run pytest

# 4. Desenvolver com hot reload (se aplicável)
uv run uvicorn app.main:app --reload
```

### CI/CD (GitHub Actions)
```yaml
name: CI

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Install UV
        run: |
          curl -LsSf https://astral.sh/uv/install.sh | sh
          echo "$HOME/.cargo/bin" >> $GITHUB_PATH
          
      - name: Set up Python
        run: uv python install 3.11
          
      - name: Install dependencies
        run: uv sync --dev
        
      - name: Check formatting
        run: uv run ruff format --check .
        
      - name: Check linting
        run: uv run ruff check .
        
      - name: Type check
        run: uv run mypy .
        
      - name: Run tests
        run: uv run pytest --cov=.
```

### Packaging e Publicação
```bash
# Construir distribuição
uv build

# Publicar no PyPI
uv publish

# Instalar localmente em modo desenvolvimento
uv pip install -e .
```

### Monorepo com Workspaces
```toml
# pyproject.toml na raiz do monorepo
[tool.uv]
workspace = true

# Em cada subprojeto
[tool.uv]
workspace = true
members = ["packages/*"]
```

---

## 🛠️ Comandos Essenciais

### Ambiente Virtual
```bash
# Ativar ambiente
source .venv/bin/activate  # Linux/macOS
.venv\Scripts\activate     # Windows

# Executar comando no ambiente virtual (sem ativar)
uv run python script.py
uv run pytest
```

### Inspeção
```bash
# Verificar dependências
uv tree

# Verificar dependências conflitantes
uv tree --reverse

# Listar pacotes instalados
uv pip list
```

### Cache
```bash
# Limpar cache
uv cache clean

# Mostrar informações do cache
uv cache info
```

---

## 🔗 Referências

- **Documentação Oficial do UV**: https://docs.astral.sh/uv/
- **Site da Astral**: https://astral.sh/
- **Ruff (linter/formatter)**: https://docs.astral.sh/ruff/
- **Python Packaging Guide**: https://packaging.python.org/
- **Pydantic (validação de dados)**: https://docs.pydantic.dev/
- **FastAPI (se for usar web)**: https://fastapi.tiangolo.com/

---

## 📝 Changelog

Consulte o [CHANGELOG.md](CHANGELOG.md) para o histórico completo de versões.

---

<div align="center">

**✨ Dica Pro:** Use `uv` para tudo relacionado a Python - é mais rápido, mais simples e mais confiável que a pilha tradicional.

*"One tool to rule them all" - Pythonistas modernos*

</div>