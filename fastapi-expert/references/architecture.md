# FastAPI Clean Architecture

Para APIs de nível enterprise, a `fastapi-expert` recomenda a separação em camadas para garantir testabilidade e manutenção.

## 1. Camada de Repositório (Data Access)

Abstrai o acesso ao banco de dados (SQLAlchemy, Tortoise, etc).

```python
from typing import Generic, TypeVar, Type
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import select

T = TypeVar("T")

class BaseRepository(Generic[T]):
    def __init__(self, model: Type[T]):
        self.model = model

    async def get_by_id(self, db: AsyncSession, id: int) -> T | None:
        result = await db.execute(select(self.model).where(self.model.id == id))
        return result.scalars().first()
```

## 2. Camada de Serviço (Business Logic)

Onde reside a inteligência da aplicação. Endpoints devem apenas chamar serviços.

```python
class UserService:
    def __init__(self, repo: UserRepository):
        self.repo = repo

    async def create_user(self, db: AsyncSession, data: UserCreate) -> User:
        # Lógica de validação de negócio, hashing de senha, etc.
        ...
```

## 3. Lifespan Management

Gerenciamento moderno de recursos da aplicação.

```python
from contextlib import asynccontextmanager
from fastapi import FastAPI

@asynccontextmanager
async def lifespan(app: FastAPI):
    # Startup: Conectar ao banco, carregar modelos de IA, etc.
    await database.connect()
    yield
    # Shutdown: Fechar conexões
    await database.disconnect()

app = FastAPI(lifespan=lifespan)
```
