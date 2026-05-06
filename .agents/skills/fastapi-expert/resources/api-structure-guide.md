# Guia de Estrutura de API FastAPI Escalável

Este guia define a estrutura recomendada para projetos FastAPI que precisam escalar em complexidade e volume de dados.

## Estrutura de Pastas Recomendada

```text
app/
├── main.py              # Ponto de entrada da aplicação
├── api/                 # Camada de roteamento (v1, v2, etc.)
│   ├── api_v1/
│   │   ├── api.py       # Agregador de routers
│   │   └── endpoints/   # Implementação das rotas
│   │       ├── users.py
│   │       └── items.py
├── core/                # Configurações globais e segurança
│   ├── config.py
│   └── security.py
├── crud/                # Operações de banco de dados (Create, Read, Update, Delete)
├── db/                  # Configuração da Database e Session
│   ├── base.py          # Importa todos os modelos para o Alembic
│   └── session.py
├── dependencies/        # Injeção de dependência (Auth, DB, etc.)
├── models/              # Modelos de banco de dados (SQLAlchemy/SQLModel)
├── schemas/             # Schemas de validação Pydantic (Request/Response)
└── services/            # Lógica de negócio complexa
```

## Princípios Chave

### 1. Routers Modulares
Use `APIRouter` para separar as rotas por domínio. Evite colocar muita lógica diretamente nos endpoints; delegue para `crud` ou `services`.

### 2. Schemas Pydantic
Mantenha seus modelos de banco de dados (`models/`) separados dos seus modelos de API (`schemas/`). Isso permite que a API evolua sem quebrar o contrato com o cliente.

### 3. Injeção de Dependências
Use o sistema de `Depends()` do FastAPI para gerenciar sessões de banco de dados e autenticação. Isso facilita o teste unitário através de overrides.

### 4. Camada de CRUD vs Service
- **CRUD**: Operações puras de banco de dados.
- **Service**: Lógica de negócio que pode envolver múltiplos CRUDs, chamadas externas ou processamento complexo.

## Exemplo de Endpoint Limpo

```python
@router.post("/", response_model=schemas.User)
def create_user(
    *,
    db: Session = Depends(deps.get_db),
    user_in: schemas.UserCreate,
    current_user: models.User = Depends(deps.get_current_active_superuser),
) -> Any:
    """
    Cria um novo usuário.
    """
    user = crud.user.get_by_email(db, email=user_in.email)
    if user:
        raise HTTPException(status_code=400, detail="User already exists")
    return crud.user.create(db, obj_in=user_in)
```
