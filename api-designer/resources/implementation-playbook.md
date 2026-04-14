# Implementation Playbook — API Designer

Padrões de implementação, checklists e exemplos de código para design de APIs REST e GraphQL.

---

## 1. REST — Princípios Fundamentais

### Resource-Oriented Architecture

- Recursos são **substantivos** (`users`, `orders`, `products`), nunca verbos
- Use HTTP methods para as ações (GET, POST, PUT, PATCH, DELETE)
- URLs representam hierarquias de recursos
- Convenções de nomenclatura consistentes (plural para coleções)

### Semântica dos Métodos HTTP

| Método | Semântica | Idempotente | Safe |
|--------|-----------|-------------|------|
| `GET` | Recuperar recursos | ✅ | ✅ |
| `POST` | Criar novo recurso | ❌ | ❌ |
| `PUT` | Substituir recurso inteiro | ✅ | ❌ |
| `PATCH` | Atualizar campos parcialmente | ❌* | ❌ |
| `DELETE` | Remover recurso | ✅ | ❌ |

> *PATCH pode ser idempotente dependendo da implementação.

---

## 2. GraphQL — Princípios Fundamentais

### Schema-First Development

- Tipos definem o modelo de domínio
- Queries para leitura de dados
- Mutations para modificação de dados
- Subscriptions para atualizações em tempo real

### Estrutura de Queries

- Clientes solicitam exatamente o que precisam
- Endpoint único, múltiplas operações
- Schema fortemente tipado
- Introspecção embutida

---

## 3. Estratégias de Versionamento

### URL Versioning (mais comum)
```
/api/v1/users
/api/v2/users
```

### Header Versioning
```
Accept: application/vnd.api+json; version=1
```

### Query Parameter Versioning
```
/api/users?version=1
```

**Recomendação**: URL versioning para APIs públicas (visível, explícito, fácil de testar). Header versioning para APIs internas ou com requisitos mais sofisticados.

---

## 4. Padrões REST

### Padrão 1: Design de Coleções de Recursos

```python
# ✅ Correto: endpoints orientados a recursos
GET    /api/users              # Listar usuários (com paginação)
POST   /api/users              # Criar usuário
GET    /api/users/{id}         # Buscar usuário específico
PUT    /api/users/{id}         # Substituir usuário
PATCH  /api/users/{id}         # Atualizar campos do usuário
DELETE /api/users/{id}         # Remover usuário

# Recursos aninhados
GET    /api/users/{id}/orders  # Pedidos de um usuário
POST   /api/users/{id}/orders  # Criar pedido para um usuário

# ❌ Errado: endpoints orientados a ações (evitar)
POST   /api/createUser
POST   /api/getUserById
POST   /api/deleteUser
```

### Padrão 2: Paginação e Filtros

```python
from typing import List, Optional
from pydantic import BaseModel, Field

class PaginationParams(BaseModel):
    page: int = Field(1, ge=1, description="Número da página")
    page_size: int = Field(20, ge=1, le=100, description="Itens por página")

class FilterParams(BaseModel):
    status: Optional[str] = None
    created_after: Optional[str] = None
    search: Optional[str] = None

class PaginatedResponse(BaseModel):
    items: List[dict]
    total: int
    page: int
    page_size: int
    pages: int

    @property
    def has_next(self) -> bool:
        return self.page < self.pages

    @property
    def has_prev(self) -> bool:
        return self.page > 1

# Exemplo com FastAPI
from fastapi import FastAPI, Query

app = FastAPI()

@app.get("/api/users", response_model=PaginatedResponse)
async def list_users(
    page: int = Query(1, ge=1),
    page_size: int = Query(20, ge=1, le=100),
    status: Optional[str] = Query(None),
    search: Optional[str] = Query(None)
):
    query = build_query(status=status, search=search)
    total = await count_users(query)
    offset = (page - 1) * page_size
    users = await fetch_users(query, limit=page_size, offset=offset)

    return PaginatedResponse(
        items=users,
        total=total,
        page=page,
        page_size=page_size,
        pages=(total + page_size - 1) // page_size
    )
```

### Padrão 3: Tratamento de Erros e Status Codes

```python
from fastapi import HTTPException, status
from pydantic import BaseModel
from typing import Any, List, Optional

class ErrorResponse(BaseModel):
    error: str
    message: str
    details: Optional[dict] = None
    timestamp: str
    path: str

class ValidationErrorDetail(BaseModel):
    field: str
    message: str
    value: Any

# Map de status codes padrão
STATUS_CODES = {
    "success": 200,
    "created": 201,
    "no_content": 204,
    "bad_request": 400,
    "unauthorized": 401,
    "forbidden": 403,
    "not_found": 404,
    "conflict": 409,
    "unprocessable": 422,
    "internal_error": 500
}

def raise_not_found(resource: str, id: str):
    raise HTTPException(
        status_code=status.HTTP_404_NOT_FOUND,
        detail={
            "error": "NotFound",
            "message": f"{resource} não encontrado",
            "details": {"id": id}
        }
    )

def raise_validation_error(errors: List[ValidationErrorDetail]):
    raise HTTPException(
        status_code=status.HTTP_422_UNPROCESSABLE_ENTITY,
        detail={
            "error": "ValidationError",
            "message": "Falha na validação da requisição",
            "details": {"errors": [e.dict() for e in errors]}
        }
    )
```

### Padrão 4: HATEOAS

```python
class UserResponse(BaseModel):
    id: str
    name: str
    email: str
    _links: dict

    @classmethod
    def from_user(cls, user, base_url: str):
        return cls(
            id=user.id,
            name=user.name,
            email=user.email,
            _links={
                "self": {"href": f"{base_url}/api/users/{user.id}"},
                "orders": {"href": f"{base_url}/api/users/{user.id}/orders"},
                "update": {
                    "href": f"{base_url}/api/users/{user.id}",
                    "method": "PATCH"
                },
                "delete": {
                    "href": f"{base_url}/api/users/{user.id}",
                    "method": "DELETE"
                }
            }
        )
```

---

## 5. Padrões GraphQL

### Padrão 1: Design de Schema

```graphql
# schema.graphql

type User {
  id: ID!
  email: String!
  name: String!
  createdAt: DateTime!

  # Relacionamentos com paginação
  orders(first: Int = 20, after: String, status: OrderStatus): OrderConnection!
  profile: UserProfile
}

type Order {
  id: ID!
  status: OrderStatus!
  total: Money!
  items: [OrderItem!]!
  createdAt: DateTime!
  user: User!
}

# Padrão de paginação cursor-based (Relay spec)
type OrderConnection {
  edges: [OrderEdge!]!
  pageInfo: PageInfo!
  totalCount: Int!
}

type OrderEdge {
  node: Order!
  cursor: String!
}

type PageInfo {
  hasNextPage: Boolean!
  hasPreviousPage: Boolean!
  startCursor: String
  endCursor: String
}

# Enums para type safety
enum OrderStatus {
  PENDING
  CONFIRMED
  SHIPPED
  DELIVERED
  CANCELLED
}

# Custom scalars
scalar DateTime
scalar Money

# Query root
type Query {
  user(id: ID!): User
  users(first: Int = 20, after: String, search: String): UserConnection!
  order(id: ID!): Order
}

# Mutation root com payload pattern
type Mutation {
  createUser(input: CreateUserInput!): CreateUserPayload!
  updateUser(input: UpdateUserInput!): UpdateUserPayload!
  deleteUser(id: ID!): DeleteUserPayload!
}

input CreateUserInput {
  email: String!
  name: String!
  password: String!
}

type CreateUserPayload {
  user: User
  errors: [Error!]
}

type Error {
  field: String
  message: String!
}
```

### Padrão 2: Resolvers com DataLoader (evitar N+1)

```python
from aiodataloader import DataLoader
from typing import List, Optional

class UserLoader(DataLoader):
    """Carrega usuários em batch por ID."""

    async def batch_load_fn(self, user_ids: List[str]) -> List[Optional[dict]]:
        users = await fetch_users_by_ids(user_ids)
        user_map = {user["id"]: user for user in users}
        return [user_map.get(user_id) for user_id in user_ids]

class OrdersByUserLoader(DataLoader):
    """Carrega pedidos em batch por user_id."""

    async def batch_load_fn(self, user_ids: List[str]) -> List[List[dict]]:
        orders = await fetch_orders_by_user_ids(user_ids)
        orders_by_user = {}
        for order in orders:
            uid = order["user_id"]
            orders_by_user.setdefault(uid, []).append(order)
        return [orders_by_user.get(uid, []) for uid in user_ids]

def create_context():
    return {
        "loaders": {
            "user": UserLoader(),
            "orders_by_user": OrdersByUserLoader()
        }
    }
```

---

## 6. Checklists

### ✅ Checklist REST (pré-implementação)

- [ ] Todos os endpoints usam substantivos no plural
- [ ] Métodos HTTP corretos para cada operação (GET=ler, POST=criar, PUT=substituir, PATCH=atualizar, DELETE=remover)
- [ ] Resposta de erro padronizada em todos os endpoints
- [ ] Paginação implementada para todas as coleções
- [ ] Estratégia de versionamento definida
- [ ] Autenticação especificada (JWT, API Key, OAuth2)
- [ ] Rate limiting planejado
- [ ] 404 retornado para recursos inexistentes (não 200 com body vazio)
- [ ] Documentação OpenAPI 3.x gerada ou planejada

### ✅ Checklist GraphQL (pré-implementação)

- [ ] Schema SDL revisado e aprovado
- [ ] Todos os campos non-null marcados com `!`
- [ ] DataLoaders planejados para relações (N+1 prevention)
- [ ] Paginação cursor-based (Relay spec) para coleções
- [ ] Mutations retornam payload com `errors`
- [ ] Campos deprecated marcados com `@deprecated`
- [ ] Complexidade de query limitada
- [ ] Introspecção desabilitada em produção (se necessário)

---

## 7. Pitfalls Comuns

| Problema | Causa | Solução |
|----------|-------|---------|
| Over-fetching/Under-fetching (REST) | Endpoints genéricos demais ou específicos demais | Projetar endpoints para casos de uso reais; considerar GraphQL |
| N+1 queries (GraphQL) | Resolvers fazem queries individuais por item | Usar DataLoaders para batch |
| Breaking changes sem aviso | Alterar contrato público sem versionamento | Versionar API ou usar deprecation strategy |
| Formatos de erro inconsistentes | Cada endpoint retorna erro diferente | Padronizar ErrorResponse globalmente |
| Falta de rate limiting | API sem proteção contra abuso | Implementar desde o início |
| API espelha schema do banco | Coupling entre camadas | Projetar API para o consumidor, não para o banco |
| HTTP 200 para tudo | Ignorar semântica HTTP | Usar status codes corretos sempre |

---

## 8. Recursos Externos

- **[OpenAPI Specification](https://spec.openapis.org/oas/v3.1.0)** — Padrão para documentação REST
- **[GraphQL Spec](https://spec.graphql.org/)** — Especificação oficial GraphQL
- **[Relay Cursor Connections Spec](https://relay.dev/graphql/connections.htm)** — Paginação cursor-based
- **[JSON:API](https://jsonapi.org/)** — Padrão para estrutura de resposta REST
- **[HTTP Status Codes](https://developer.mozilla.org/pt-BR/docs/Web/HTTP/Status)** — Referência MDN
