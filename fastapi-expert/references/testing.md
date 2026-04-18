# FastAPI Testing with UV & Pytest

Testar APIs assíncronas exige configuração correta do loop de eventos e do cliente HTTP.

## 1. Configuração de Fixtures (`tests/conftest.py`)

```python
import pytest
import asyncio
from httpx import AsyncClient
from app.main import app
from app.core.database import get_db

@pytest.fixture(scope="session")
def event_loop():
    loop = asyncio.get_event_loop_policy().new_event_loop()
    yield loop
    loop.close()

@pytest.fixture
async def client():
    # Exemplo de override de dependência (ex: banco em memória)
    async with AsyncClient(app=app, base_url="http://test") as ac:
        yield ac
```

## 2. Exemplo de Teste Assíncrono

Sempre utilize o decorator `@pytest.mark.asyncio`.

```python
import pytest
from httpx import AsyncClient

@pytest.mark.asyncio
async def test_read_main(client: AsyncClient):
    response = await client.get("/")
    assert response.status_code == 200
    assert response.json() == {"msg": "Hello World"}
```

## 3. Integration Tests (Database & Cleanups)

Use `dependency_overrides` para injetar um banco de dados de teste (ex: SQLite in-memory).

```python
from app.core.database import get_db

async def override_get_db():
    async with TestingSessionLocal() as session:
        yield session

app.dependency_overrides[get_db] = override_get_db
```

### Resetando o Banco entre Testes
Use fixtures para criar/limpar as tabelas em cada teste:

```python
@pytest.fixture(autouse=True)
async def setup_db():
    async with engine.begin() as conn:
        await conn.run_sync(Base.metadata.create_all)
    yield
    async with engine.begin() as conn:
        await conn.run_sync(Base.metadata.drop_all)
```

## 4. Execução com UV
```bash
uv add --dev pytest pytest-asyncio httpx aiosqlite
uv run pytest
```
