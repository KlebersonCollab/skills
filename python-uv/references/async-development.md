# Async Development with UV — Referência Completa

Guia de padrões e ferramentas para desenvolvimento de aplicações assíncronas em Python usando o ecossistema `uv`.

---

## 1. Por que usar UV para Async?

O `uv` facilita a gestão de projetos assíncronos de alta performance ao:
- **Resolução de conflitos rápida**: Bibliotecas async costumam ter árvores de dependência complexas.
- **Ambientes Isolados**: Garante que versões específicas de loop (como `uvloop`) sejam instaladas corretamente.
- **Execução Direta**: `uv run` permite rodar scripts async com dependências inline (PEP 723) sem setup manual.

---

## 2. Setup de Projeto Async

### Inicialização
```bash
uv init async-app
cd async-app

# Adicionar dependências essenciais
uv add httpx pydantic anyio
uv add --dev pytest-asyncio trio
```

### Bibliotecas Recomendadas

| Categoria | Biblioteca | Comando UV |
|-----------|------------|------------|
| **HTTP Client** | `httpx` | `uv add httpx` |
| **Event Loop (Fast)** | `uvloop` | `uv add uvloop` |
| **Web Framework** | `FastAPI` | `uv add fastapi` |
| **Banco de Dados** | `SQLAlchemy` (async) | `uv add sqlalchemy[asyncio]` |
| **Testes** | `pytest-asyncio` | `uv add --dev pytest-asyncio` |

---

## 3. Padrões de Código no Ecossistema UV

### Execução Concorrente (Gather)
Padrão para buscar múltiplos recursos simultaneamente.

```python
import asyncio
import httpx

async def fetch_urls(urls: list[str]):
    async with httpx.AsyncClient() as client:
        tasks = [client.get(url) for url in urls]
        responses = await asyncio.gather(*tasks)
        return [r.status_code for r in responses]

# Executar com uv
# uv run python script.py
```

### Proteção de Recursos (Semaphores)
Evita sobrecarga em APIs externas ao limitar a concorrência.

```python
semaphore = asyncio.Semaphore(10) # No máximo 10 requests simultâneos

async def safe_fetch(url):
    async with semaphore:
        # Lógica de fetch aqui
        pass
```

---

## 4. Testando Código Async com UV

Configure o `pytest-asyncio` no seu `pyproject.toml` para evitar decorators repetitivos.

```toml
[tool.pytest.ini_options]
asyncio_mode = "auto"
asyncio_default_fixture_loop_scope = "function"
```

### Exemplo de Teste
```python
import pytest
from app import fetch_data

@pytest.mark.asyncio
async def test_fetch():
    data = await fetch_data("https://api.example.com")
    assert data["status"] == "ok"

# Rodar testes
# uv run pytest
```

---

## 5. Scripts Inline (PEP 723) para Async

Ideal para ferramentas de scraping ou automação rápida.

```python
# /// script
# dependencies = ["httpx", "rich"]
# ///
import asyncio
import httpx
from rich import print

async def main():
    async with httpx.AsyncClient() as client:
        r = await client.get("https://httpbin.org/get")
        print(r.json())

if __name__ == "__main__":
    asyncio.run(main())
```

---

## 6. Performance: Usando `uvloop`

Em sistemas Linux/macOS, o `uvloop` pode tornar o loop do asyncio 2-4x mais rápido.

```python
import asyncio
import sys

if sys.platform != "win32":
    import uvloop
    asyncio.set_event_loop_policy(uvloop.EventLoopPolicy())

async def main():
    # Sua aplicação aqui
    pass

asyncio.run(main())
```

---

## 7. Checklist de Desenvolvimento Async

- [ ] Usou `asyncio.run()` como ponto de entrada?
- [ ] Evitou funções bloqueantes como `time.sleep()`? (Use `asyncio.sleep()`)
- [ ] Configurou timeouts em todas as chamadas de rede?
- [ ] Usou `uv run` para garantir que as dependências async estão no ambiente correto?
- [ ] Implementou tratamento de erro robusto no `gather` com `return_exceptions=True`?
