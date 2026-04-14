# Python Environment

Guia completo de gerenciamento de versões Python, Python 3.14 default, free-threaded Python e paths por plataforma.

---

## Gerenciamento de Versões Python

### Comandos Essenciais

```bash
# Listar versões disponíveis
uv python list

# Listar versões instaladas
uv python list --installed

# Instalar versão específica
uv python install 3.12
uv python install 3.13
uv python install 3.14

# Fixar versão para o projeto
uv python pin 3.12

# Verificar versão fixada
cat .python-version
```

### Executar com versão específica

```bash
# Executar script com Python específico
uv run --python 3.12 script.py

# Criar venv com versão específica
uv venv --python 3.12 .venv
```

---

## Python 3.14 — Novo Default (UV 0.9.6+)

A partir do UV 0.9.6, `uv python install` sem versão instala **Python 3.14** (antes era 3.13).

### Impacto

```bash
# Antes (UV < 0.9.6)
uv python install    # → Python 3.13

# Agora (UV 0.9.6+)
uv python install    # → Python 3.14
```

### Migração

Se o projeto requer Python 3.13 ou anterior:

```bash
# Fixar versão no projeto
uv python pin 3.13

# Instalar versão específica
uv python install 3.13

# Verificar
cat .python-version
# Output: 3.13
```

---

## Free-Threaded Python (PEP 703)

Python 3.14+ suporta **free-threading** — execução paralela real sem o GIL (Global Interpreter Lock).

### O que muda

- **Antes**: GIL impedia threads Python de executar código em paralelo
- **Agora**: Free-threaded Python permite true multi-core parallelism

### Como usar

```bash
# UV 0.9.6+ — free-threaded é automático para 3.14+
uv python install 3.14

# Ou explicitamente
uv python install 3.14t
```

### Exemplo

```python
# /// script
# requires-python = ">=3.14"
# dependencies = []
# ///

import threading
import time

def cpu_task(n: int) -> int:
    """CPU-intensive task."""
    return sum(i ** 2 for i in range(n))

# Com free-threaded Python, estas threads rodam em paralelo
threads = []
for i in range(4):
    t = threading.Thread(target=cpu_task, args=(10_000_000,))
    threads.append(t)
    t.start()

for t in threads:
    t.join()

print("Threads executaram em paralelo!")
```

### Quando usar

- ✅ Código CPU-intensive com multi-threading
- ✅ Processamento paralelo crítico para performance
- ⚠️ Verificar compatibilidade de bibliotecas antes de usar em produção

---

## Paths por Plataforma

### Cache do UV

| Plataforma | Caminho |
|------------|---------|
| macOS | `~/Library/Caches/uv/` |
| Linux | `~/.cache/uv/` |
| Windows | `%LOCALAPPDATA%\uv\cache\` |

### Ferramentas Instaladas

| Plataforma | Caminho |
|------------|---------|
| macOS / Linux | `~/.local/share/uv/tools/` |
| Windows | `%APPDATA%\uv\tools\` |

### Binários Python

| Plataforma | Caminho |
|------------|---------|
| macOS / Linux | `~/.local/share/uv/python/` |
| Windows | `%APPDATA%\uv\python\` |

### Gerenciamento de Cache

```bash
# Ver tamanho do cache
du -sh ~/.cache/uv/          # macOS/Linux
dir /s %LOCALAPPDATA%\uv\cache  # Windows

# Limpar cache
uv cache clean

# Informações do cache
uv cache info
```

---

## Compatibilidade

| UV Version | Python 3.10 | Python 3.11 | Python 3.12 | Python 3.13 | Python 3.14 | Free-Threading |
|------------|-------------|-------------|-------------|-------------|-------------|----------------|
| 0.9.0-0.9.5 | ✅ | ✅ | ✅ | ✅ (default) | Beta | Experimental |
| 0.9.6+ | ✅ | ✅ | ✅ | ✅ | ✅ (default) | Stable |
| 0.9.7+ | ✅ | ✅ | ✅ | ✅ | ✅ (default) | Stable |
