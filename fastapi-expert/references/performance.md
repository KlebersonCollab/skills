# FastAPI Performance Reference

## 1. Async vs Sync Decision Matrix

O FastAPI executa funções `async def` diretamente no event loop e funções `def` em um threadpool.

| Tipo de Função | Quando Usar | Por que? |
|----------------|-------------|----------|
| `async def` | I/O não bloqueante (HTTPX, Motor, Asyncpg) | Evita bloquear o event loop. |
| `def` | I/O bloqueante (Requests, SQLAlchemy, File I/O) | Roda em threads separadas para não travar a API. |
| `def` | CPU-bound (Cálculos pesados, Processamento de imagem) | Threadpool lida melhor sem travar outras requisições. |

### ⚠️ O Perigo do "Fake Async"
NUNCA use `time.sleep()` ou bibliotecas bloqueantes (como `requests`) dentro de um `async def`. Isso trava TODAS as requisições da sua API simultaneamente.

## 2. Rust-Powered Serialization (Pydantic V2)

O FastAPI usa Pydantic V2, que é escrito em Rust. Para máxima performance:

- **Type Hints Estritos:** Quanto mais preciso o tipo, mais rápido o Rust valida.
- **Annotated:** Use `Annotated` para evitar que o Pydantic precise inferir tipos complexos múltiplas vezes.
- **Exclude Unset:** Ao retornar modelos, use `response_model_exclude_unset=True` para reduzir o tamanho do payload.

## 3. Tooling
- **Uvicorn vs Gunicorn:** Use Uvicorn para dev e Gunicorn com workers Uvicorn para produção.
- **uv run:** Sempre execute via `uv run` para garantir que as otimizações do ambiente virtual estejam ativas.
