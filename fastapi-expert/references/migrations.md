# Database Migrations with Alembic (Async)

O gerenciamento de mudanças no esquema do banco de dados deve ser versionado e automatizado.

## 1. Setup Assíncrono com UV

Sempre utilize o template `async` do Alembic para compatibilidade com FastAPI.

```bash
uv add alembic sqlalchemy asyncpg  # ou aiosqlite
uv run alembic init -t async migrations
```

## 2. Configuração do `env.py`

No arquivo `migrations/env.py`, integre suas configurações do Pydantic e seus modelos SQLAlchemy:

```python
from app.core.config import settings
from app.models.base import Base  # Metadata que contém todos os modelos

# ...
config = context.config
config.set_main_option("sqlalchemy.url", settings.DATABASE_URL)
target_metadata = Base.metadata
# ...
```

## 3. Workflow de Migração

| Comando | Descrição |
|---------|-----------|
| `uv run alembic revision --autogenerate -m "desc"` | Cria uma nova migração detectando mudanças nos modelos. |
| `uv run alembic upgrade head` | Aplica todas as migrações pendentes. |
| `uv run alembic downgrade -1` | Reverte a última migração aplicada. |
| `uv run alembic history` | Lista o histórico de migrações. |

## 4. Best Practices
- **Manual Review**: Sempre revise o script gerado por `--autogenerate` antes de aplicar. Ele pode errar em renomeações de colunas.
- **Atomic Migrations**: Cada migração deve ser pequena e focar em uma única mudança lógica.
- **Production Safety**: NUNCA rode `upgrade head` automaticamente em produção sem um backup recente do banco.
