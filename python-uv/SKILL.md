---
name: python-uv
version: 2.6.0
description: "Skill para desenvolvimento Python profissional com UV. Use para gerenciar dependências e ambientes. Para desenvolvimento de APIs ou Web, utilize em conjunto com as skills 'fastapi-expert' ou 'django-expert'."
category: development-workflow
---

# Python com UV

> Gerenciador de pacotes e projetos Python extremamente rápido, escrito em Rust. Substitui pip, pipx, poetry, pyenv e mais — com performance 10-100x superior.

---

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar qualquer execução técnica:
1. **Context Check**: Você reidratou o contexto lendo `STATE.md`, `MEMORY.md` e `LEARNINGS.md`?
2. **Spec Check**: O arquivo `spec.md` existe com requisitos e Critérios de Aceitação (ACs) claros? (BDD mandatório para Medium+).
3. **Plan Check**: O arquivo `plan.md` define a arquitetura, schemas e inclui diagramas **Mermaid**?
4. **Contract Check**: O arquivo `contract.md` foi estabelecido com os sensores de validação?
5. **Task Check**: A lista de tarefas em `tasks.md` está detalhada e atomizada?

---
## Goal

Capacitar o agente a gerenciar projetos Python com performance extrema utilizando o **UV**. A skill unifica o gerenciamento de pacotes, versões de Python e ambientes virtuais em um workflow único, rápido e determinístico.

---

## Version Awareness

**Versão Recomendada: UV 0.9.7+**

Antes de começar, verifique a versão instalada:

```bash
uv --version
```

**Mudanças por versão:**

| Versão | Mudança |
|--------|---------|
| **0.9.6+** | Python 3.14 é o default (antes era 3.13) |
| **0.9.6+** | Free-threaded Python 3.14+ sem opt-in explícito |
| **0.9.6+** | `uv build --clear` para limpar artefatos de build |
| **0.9.7+** | Updates de segurança para handling de tar/ZIP |

Se a versão for anterior a 0.9.0, atualize:

```bash
# macOS / Linux
curl -LsSf https://astral.sh/uv/install.sh | sh

# macOS via Homebrew
brew install uv

# Windows
powershell -ExecutionPolicy ByPass -c "irm https://astral.sh/uv/install.ps1 | iex"
```

Veja [Recent Changes](references/python-environment.md) para detalhes de migração.

## When to Use This Skill

Use this skill when:
- Criando ou inicializando um projeto Python novo
- Gerenciando dependências com `pyproject.toml`
- Instalando ou gerenciando versões de Python
- Configurando ambientes virtuais
- Instalando ferramentas CLI de desenvolvimento (black, ruff, mypy, pytest)
- Executando scripts Python com dependências inline (PEP 723)
- Configurando CI/CD com UV (GitHub Actions)
- Decidindo entre `uv tool install` vs `uvx`
- Migrando de pip, pipx ou poetry para UV
- Resolvendo erros relacionados ao UV

Skip this skill when:
- O projeto usa Python 2.x (UV não suporta)
- O projeto NÃO usa Python como linguagem
- Apenas precisa de documentação básica de pip (sem UV)

---

## Workflow (4 Fases)

### Fase 1: ENVIRONMENT — Configuração e Alinhamento
1.  **Verificar Versão**: Executar `uv --version` para garantir compatibilidade (Recomendado 0.9.7+).
2.  **Definir Versão Python**: Se não existir `.python-version`, use `uv python pin 3.12` (ou versão desejada).
3.  **Ambiente Isolado**: Garantir que o `.venv/` existe ou será criado via `uv venv`.

### Fase 2: PROJECT — Inicialização e Estrutura
1.  **Scaffolding**: Para novos projetos, use `uv init`. Para migrações, use `uv init --bare`.
2.  **Declarar Dependências**: Adicionar pacotes via `uv add` (produção) e `uv add --dev` (ferramentas).
3.  **Sync Global**: Executar `uv sync` para criar o `uv.lock` e sincronizar o ambiente virtual.

### Fase 3: DEVELOP — Qualidade e Execução
1.  **Configurar Tooling**: Definir regras no `pyproject.toml` para Ruff e Mypy.
2.  **Execução Segura**: Sempre utilize `uv run <comando>` para garantir que o código rode no ambiente correto.
3.  **Qualidade Incremental**: Rodar `uv run ruff check .` e `uv run pytest` periodicamente.

### Fase 4: DEPLOY — Build e Distribuição
1.  **Lock Reproduzível**: Garantir que o `uv.lock` está atualizado e no controle de versão.
2.  **Compilação**: Usar `uv build --clear` para gerar distribuições limpas (sdist/wheel).
3.  **Publicação**: Executar `uv publish` para distribuir no PyPI ou repositório privado.

---

## UV Commands Overview

| Comando | Propósito | Exemplo |
|---------|-----------|---------|
| `uv init` | Criar novo projeto | `uv init meu-projeto` |
| `uv add` | Adicionar dependência | `uv add requests pydantic` |
| `uv remove` | Remover dependência | `uv remove requests` |
| `uv sync` | Sincronizar ambiente | `uv sync --dev` |
| `uv lock` | Gerar/atualizar lockfile | `uv lock` |
| `uv run` | Executar no ambiente virtual | `uv run pytest` |
| `uv pip install` | Instalar pacotes (modo pip) | `uv pip install requests` |
| `uv tool install` | Instalar CLI tool global isolada | `uv tool install black` |
| `uvx` | Executar pacote em ambiente temporário | `uvx ruff check .` |
| `uv venv` | Criar ambiente virtual | `uv venv .venv` |
| `uv python install` | Instalar versão Python | `uv python install 3.12` |
| `uv python pin` | Fixar versão Python no projeto | `uv python pin 3.12` |
| `uv build` | Construir distribuição | `uv build --clear` |
| `uv publish` | Publicar no PyPI | `uv publish` |

---

## Tool vs UVX Decision Tree

```text
Precisa executar um pacote Python?
│
├─ Usa diariamente/frequentemente?
│  └─ SIM → `uv tool install pacote`
│     Exemplos: black, ruff, mypy, pytest
│
├─ Execução temporária / teste?
│  └─ SIM → `uvx pacote`
│     Exemplos: testar nova ferramenta, comparar versões
│
├─ MCP server?
│  └─ SIM → `uvx pacote` ou `uvx --from path script.py`
│     Exemplos: mcp-server-sqlite, custom MCP servers
│
└─ Script local de projeto?
   └─ SIM → `uvx --from . script.py`
      Exemplos: scripts project-specific
```

---

## Project Initialization

```bash
# Criar novo projeto
uv init meu-projeto
cd meu-projeto

# Estrutura gerada:
# meu-projeto/
# ├── .python-version
# ├── pyproject.toml
# ├── README.md
# └── src/
#     └── meu_projeto/
#         └── __init__.py

# Adicionar dependências
uv add requests pydantic

# Adicionar dependências de desenvolvimento
uv add --dev pytest ruff mypy

# Executar
uv run python -m meu_projeto
```

### Gerenciamento de Dependências

```bash
# Adicionar dependência de produção
uv add requests pydantic

# Adicionar dependência de desenvolvimento
uv add --dev pytest ruff mypy

# Adicionar dependência opcional
uv add --optional dev pytest-cov

# Remover dependência
uv remove requests

# Sincronizar ambiente (instala tudo do pyproject.toml)
uv sync --dev

# Gerar lockfile
uv lock

# Sincronizar do lockfile (reprodução exata)
uv sync --locked
```

---

## Virtual Environment Management

```bash
# UV cria automaticamente ao usar uv sync/run
# Mas se precisar criar manualmente:
uv venv .venv

# Ativar (macOS / Linux)
source .venv/bin/activate

# Ativar (Windows Git Bash)
. .venv/Scripts/activate

# Ativar (Windows CMD)
.venv\Scripts\activate.bat

# Executar SEM ativar (recomendado)
uv run python script.py
uv run pytest
```

---

## Inline Script Metadata (PEP 723)

Scripts Python self-contained com dependências definidas em comentários:

```python
# /// script
# dependencies = [
#   "requests>=2.28.0",
#   "rich>=13.0.0",
# ]
# requires-python = ">=3.11"
# ///

import requests
from rich.console import Console

console = Console()
response = requests.get("https://api.github.com")
console.print(response.json())
```

Executar com instalação automática de dependências:

```bash
uv run script.py
```

**Quando usar**: Scripts únicos, utilitários, automação, exemplos compartilháveis.
**Quando NÃO usar**: Projetos multi-arquivo (use `pyproject.toml`).

Veja [Inline Script Metadata Reference](references/inline-script-metadata.md) para exemplos completos.

---

## Development Tools Setup

```bash
# Instalar ferramentas de desenvolvimento uma vez
uv tool install ruff
uv tool install mypy
uv tool install pytest

# Usar diariamente (disponíveis globalmente)
ruff format .
ruff check .
mypy .
pytest tests/
```

### Configuração recomendada no `pyproject.toml`

```toml
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
```

---

## Best Practices

### Project Management

**DO:**
- Use `uv init` para novos projetos (cria estrutura completa)
- Use `uv add` / `uv remove` para gerenciar dependências (não edite pyproject.toml manualmente)
- Sempre versione `uv.lock` no git para reprodução exata
- Use `uv sync --locked` em CI/CD para builds reproduzíveis
- Use `uv run` para executar comandos (não precisa ativar venv)

**DON'T:**
- Não instale pacotes globalmente sem venv
- Não misture `pip install` e `uv add` no mesmo projeto
- Não ignore o `uv.lock` no `.gitignore`
- Não use `@latest` em produção (instável)

### Tool Management

**DO:**
- Use `uv tool install` para ferramentas de uso diário (ruff, mypy, pytest)
- Use `uvx` para execuções temporárias e teste de ferramentas
- Mantenha ferramentas atualizadas com `uv tool upgrade --all`

**DON'T:**
- Não use `pip install` global para CLI tools (causa conflitos)
- Não use `uvx` para ferramentas diárias (overhead desnecessário)
- Não misture instalações pip e uv tool

### Code Quality

**DO:**
- Sempre use type hints em código Python
- Configure `ruff` para linting e formatting
- Configure `mypy` em modo strict
- Escreva testes com `pytest`

**DON'T:**
- Não promova código sem type hints
- Não pule linting, formatting ou type checking
- Não use práticas deprecated como `pip` sem venv

---

## Troubleshooting

### "UV não encontrado" / "spawn uvx ENOENT"
**Causa**: UV/UVX não está no PATH.
**Solução**:
```bash
# macOS/Linux — adicionar ao ~/.zshrc ou ~/.bashrc
export PATH="$HOME/.local/bin:$PATH"
source ~/.zshrc

# Verificar
uv --version
```

### "Package not found"
**Causa**: Nome do pacote incorreto ou pacote não publicado no PyPI.
**Solução**:
```bash
# Verificar no PyPI
curl https://pypi.org/pypi/nome-do-pacote/json | head -5

# Para projetos locais, verificar pyproject.toml
ls pyproject.toml
```

### "No virtual environment found"
**Causa**: Ambiente virtual não ativado ou inexistente.
**Solução**:
```bash
# Criar e sincronizar
uv sync

# Ou executar diretamente (sem ativar)
uv run python script.py
```

### Conflitos de versão
**Causa**: Versões incompatíveis de dependências.
**Solução**:
```bash
# Verificar árvore de dependências
uv tree

# Verificar conflitos reversos
uv tree --reverse

# Limpar cache e ressincronizar
uv cache clean
uv sync --recreate
```

---

## Migration Guides

### From pip
```bash
# Antes                          # Depois
pip install requests      →      uv add requests
pip install -r req.txt    →      uv pip install -r req.txt
pip freeze > req.txt      →      uv lock  (usa uv.lock)
```

### From poetry
```bash
# Antes                          # Depois
poetry add requests       →      uv add requests
poetry install            →      uv sync
poetry run pytest         →      uv run pytest
poetry build              →      uv build
```

### From pipx
```bash
# Antes                          # Depois
pipx install black        →      uv tool install black
pipx run black .          →      uvx black .
```

---

## Reference Documentation

Esta skill inclui documentação de referência detalhada:

1. **[Installation & Setup](references/installation-and-setup.md)** — Instalação cross-platform, virtual environments, troubleshooting
2. **[Project Management](references/project-management.md)** — Init, dependências, lockfile, workspaces, estrutura de projeto
3. **[Tool Management](references/tool-management.md)** — `uv tool install` vs `uvx`, decision matrix, manutenção
4. **[Inline Script Metadata](references/inline-script-metadata.md)** — PEP 723, scripts self-contained, exemplos práticos
5. **[Python Environment](references/python-environment.md)** — Versões, Python 3.14, free-threaded Python, paths
6. **[CI/CD Workflows](references/ci-cd-workflows.md)** — GitHub Actions, Docker, packaging, publicação

---

## Output Structure

A execução desta skill em projetos Python deve resultar na criação ou atualização dos seguintes artefatos padronizados pelo UV:

| Artefato | Arquivo | Descrição |
|----------|---------|-----------|
| **Project Spec** | `pyproject.toml` | Definição de metadados, dependências e configurações de ferramentas (Ruff, Mypy). |
| **Lockfile** | `uv.lock` | Registro determinístico de todas as dependências e sub-dependências (deve ser versionado). |
| **Python Version** | `.python-version` | Arquivo que fixa a versão do Python utilizada no projeto para o UV. |
| **Virtual Env** | `.venv/` | Ambiente virtual isolado contendo as dependências instaladas (não deve ser versionado). |
| **Source Code** | `src/` ou root | Estrutura de pacotes Python seguindo as boas práticas de layout. |

---

## External Resources

- **UV Official Docs**: <https://docs.astral.sh/uv/>
- **UV GitHub**: <https://github.com/astral-sh/uv>
- **Ruff (linter/formatter)**: <https://docs.astral.sh/ruff/>
- **PEP 723 (Inline Metadata)**: <https://peps.python.org/pep-0723/>
- **Python Packaging Guide**: <https://packaging.python.org/>

---

## Quality Rules

- **Referência Oficial**: Sempre referenciar a documentação oficial do UV (<https://docs.astral.sh/uv/>)
- **Exemplos Práticos**: Todos os exemplos devem ser testáveis e reproduzíveis
- **Cross-Platform**: Comandos e configurações devem funcionar em macOS, Linux e Windows
- **Boas Práticas Modernas**: Promover type hints, ruff, pytest, mypy
- **Versionamento Semântico**: Seguir SemVer para configurações de projeto
- **Lockfiles**: Sempre usar e versionar `uv.lock`
- **UV-First**: Preferir comandos UV nativos (`uv add`, `uv run`) sobre equivalentes legados (`pip install`)

## Prohibited

- **Não assumir SO**: Não assumir configurações específicas de sistema operacional sem oferecer alternativas
- **Não duplicar documentação**: Não copiar conteúdo da documentação oficial, apenas referenciar
- **Não promover práticas obsoletas**: Não ensinar `pip install` sem venv, `setup.py`, ou `requirements.txt` como gestão principal
- **Não ignorar type safety**: Não promover código Python sem type hints
- **Não pular validação**: Não pular linting, formatting ou type checking
- **Não misturar gerenciadores**: Não misturar pip e uv no mesmo workflow de projeto