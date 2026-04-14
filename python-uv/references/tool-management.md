# Tool Management

Guia completo sobre `uv tool install` vs `uvx` — quando usar cada um, feature matrix e workflows de manutenção.

---

## Quick Decision Guide

| Cenário | Usar | Por quê |
|---------|------|---------|
| Instalar black/ruff para uso diário | `uv tool install black` | Persistente, sem overhead |
| Testar uma ferramenta nova | `uvx ferramenta` | Temporário, sem instalação |
| Executar MCP server | `uvx mcp-server-sqlite` | Padrão da comunidade |
| Instalar ferramentas de dev | `uv tool install pytest` | Uso frequente |
| Comparar versões | `uvx pacote@1.0.0` | Fácil troca de versão |
| Script local de projeto | `uvx --from . script.py` | Execução one-off |

---

## `uv tool install` — Instalação Persistente

### Comandos Básicos

```bash
# Instalar do PyPI
uv tool install pacote

# Versão específica
uv tool install pacote==1.2.3

# De repositório git
uv tool install git+https://github.com/user/repo.git

# Com extras
uv tool install "pacote[extra1,extra2]"

# Editable install (dev local)
uv tool install -e .
```

### Gerenciamento

```bash
# Listar ferramentas instaladas
uv tool list

# Atualizar uma ferramenta
uv tool upgrade pacote

# Atualizar todas
uv tool upgrade --all

# Desinstalar
uv tool uninstall pacote

# Forçar reinstalação
uv tool install pacote --force
```

### Ferramentas Recomendadas

```bash
# Desenvolvimento
uv tool install ruff       # Linter + formatter
uv tool install mypy       # Type checker
uv tool install pytest     # Testing

# Utilitários
uv tool install httpie     # HTTP client
uv tool install rich-cli   # Terminal formatting
uv tool install pre-commit # Git hooks
```

---

## `uvx` — Execução Temporária

### Comandos Básicos

```bash
# Executar do PyPI
uvx pacote

# Versão específica
uvx pacote@1.2.3

# Com argumentos
uvx pacote arg1 arg2 --flag

# De projeto local
uvx --from /path/to/project script.py

# De diretório atual
uvx --from . script.py
```

### Use Cases

```bash
# MCP Servers
uvx mcp-server-sqlite --db-path /path/to/db
uvx mcp-server-git --repository /path/to/repo

# Teste de versões
uvx black@22.0.0 --check .
uvx black@23.0.0 --check .
uvx black@24.0.0 --check .

# Scripts locais
uvx --from . scripts/setup_db.py
uvx --from . scripts/generate_docs.py
```

---

## Feature Matrix

| Feature | `uv tool install` | `uvx` |
|---------|-------------------|-------|
| **Persistência** | Permanente | Temporário (cache) |
| **Caso de Uso** | Ferramentas diárias | Execução one-off |
| **Velocidade** | Instantâneo (já instalado) | Setup na 1ª execução |
| **Versionamento** | Reinstalar para mudar | Fácil (`@versão`) |
| **Armazenamento** | Disco persistente | Cache gerenciado |
| **Atualização** | `uv tool upgrade` | Automático com `@latest` |
| **Dev local** | `uv tool install -e .` | `uvx --from .` |

---

## Manutenção

### Rotina Semanal

```bash
# Verificar ferramentas instaladas
uv tool list

# Atualizar todas
uv tool upgrade --all

# Limpar cache antigo
uv cache clean

# Remover ferramentas não usadas
uv tool uninstall ferramenta-velha
```

### Documentar Ferramentas do Time

Crie um arquivo `tools.txt` para consistência:

```text
# Ferramentas de desenvolvimento
ruff==0.8.0
mypy==1.13.0
pytest==8.3.0

# Utilitários
httpie==3.2.0
pre-commit==4.0.0
```

Instalar todas:

```bash
cat tools.txt | grep -v '^#' | grep -v '^$' | xargs -n1 uv tool install
```

---

## Anti-Patterns

| ❌ Errado | ✅ Correto | Por quê |
|-----------|-----------|---------|
| `pip install black` global | `uv tool install black` | Evita conflitos de dependências |
| `uv tool install mcp-server` | `uvx mcp-server` | Padrão da comunidade MCP |
| `uvx black .` diariamente | `uv tool install black` | Evita overhead de setup |
| Misturar `pip` e `uv tool` | Escolher um | Evita conflitos |
| Nunca atualizar ferramentas | `uv tool upgrade --all` semanal | Segurança e features |

---

## Troubleshooting

### Ferramenta não encontrada após instalação

```bash
# Verificar PATH
echo $PATH | tr ':' '\n' | grep -i uv

# Adicionar ao PATH
export PATH="$HOME/.local/bin:$PATH"
```

### Conflito de dependências

```bash
uv tool uninstall pacote
uv cache clean
uv tool install pacote
```

### UVX falha com pacote local

```bash
# Verificar que pyproject.toml existe
ls /path/to/project/pyproject.toml

# Usar caminho absoluto com --from
uvx --from /absolute/path/to/project script.py
```
