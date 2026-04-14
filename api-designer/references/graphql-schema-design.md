# GraphQL Schema Design — Referência Completa

Guia de referência de padrões de schema GraphQL e anti-padrões. Use em conjunto com o `resources/implementation-playbook.md`.

---

## 1. Fundamentos de Tipagem

### Tipos Escalares

```graphql
# Escalares built-in
String
Int
Float
Boolean
ID

# Custom scalars (declarar e implementar resolver)
scalar DateTime   # ISO 8601: "2026-04-14T20:00:00Z"
scalar Date       # ISO 8601: "2026-04-14"
scalar Money      # Centavos como Int ou string decimal
scalar UUID       # "550e8400-e29b-41d4-a716-446655440000"
scalar JSON       # Dados arbitrários (usar com cuidado)
scalar Upload     # Para file upload (multipart)
```

### Non-Null e Listas

```graphql
# Non-null: campo obrigatório (nunca retorna null)
name: String!

# Lista nullable de itens non-null
items: [OrderItem!]

# Lista non-null de itens non-null (o mais comum em coleções)
items: [OrderItem!]!

# Lista nullable de itens nullable (evitar)
items: [OrderItem]
```

### Enums

```graphql
enum OrderStatus {
  PENDING
  CONFIRMED
  PROCESSING
  SHIPPED
  DELIVERED
  CANCELLED
  REFUNDED
}

enum SortOrder {
  ASC
  DESC
}
```

---

## 2. Patterns de Schema

### Padrão: Relay Connection (paginação cursor-based)

```graphql
# Padrão Relay para paginação — use sempre para coleções grandes
type UserConnection {
  edges: [UserEdge!]!
  pageInfo: PageInfo!
  totalCount: Int!
}

type UserEdge {
  node: User!
  cursor: String!
}

type PageInfo {
  hasNextPage: Boolean!
  hasPreviousPage: Boolean!
  startCursor: String
  endCursor: String
}

# Uso na query
type Query {
  # Paginação forward (first/after)
  users(
    first: Int = 20
    after: String
    search: String
    sortBy: UserSortField
    sortOrder: SortOrder
  ): UserConnection!

  # Paginação backward (last/before) — opcional
  # last: Int
  # before: String
}
```

### Padrão: Mutation Payloads

```graphql
# ✅ Correto: mutations retornam payload com errors
type Mutation {
  createUser(input: CreateUserInput!): CreateUserPayload!
  updateUser(input: UpdateUserInput!): UpdateUserPayload!
  deleteUser(id: ID!): DeleteUserPayload!
}

input CreateUserInput {
  email: String!
  name: String!
  password: String!
  role: UserRole = USER
}

type CreateUserPayload {
  user: User          # null se falhou
  errors: [UserError!]!  # vazio se sucesso
}

type UserError {
  field: String       # null para erros globais
  message: String!
  code: String!       # Ex: "EMAIL_TAKEN", "WEAK_PASSWORD"
}

# ❌ Errado: mutation retorna tipo diretamente (sem tratamento de erro)
# createUser(input: CreateUserInput!): User!
```

### Padrão: Input Types

```graphql
# Inputs específicos por operação (não reutilizar entre create/update)
input CreateUserInput {
  email: String!
  name: String!
  password: String!
}

input UpdateUserInput {
  id: ID!
  name: String        # nullable = campo opcional na atualização
  email: String
}

# Para updates parciais (patch), use campos nullable
# Para updates completos (put), use campos non-null
```

### Padrão: Interfaces e Unions

```graphql
# Interface — tipos com campos em comum
interface Node {
  id: ID!
}

interface Timestamped {
  createdAt: DateTime!
  updatedAt: DateTime!
}

type User implements Node & Timestamped {
  id: ID!
  email: String!
  name: String!
  createdAt: DateTime!
  updatedAt: DateTime!
}

# Union — retorno que pode ser de tipos diferentes
union SearchResult = User | Product | Order

type Query {
  search(query: String!): [SearchResult!]!
}
```

---

## 3. Queries Avançadas

### Fragments e Inline Fragments

```graphql
# Fragment reutilizável
fragment UserBasic on User {
  id
  name
  email
}

query GetUsers {
  users(first: 10) {
    edges {
      node {
        ...UserBasic
        createdAt
      }
    }
  }
}

# Inline fragment para unions
query Search($query: String!) {
  search(query: $query) {
    ... on User {
      id
      name
      email
    }
    ... on Product {
      id
      name
      price
    }
  }
}
```

### Aliases

```graphql
query GetMultipleUsers {
  adminUsers: users(role: ADMIN, first: 5) {
    edges { node { id name } }
  }
  regularUsers: users(role: USER, first: 5) {
    edges { node { id name } }
  }
}
```

---

## 4. Directives

```graphql
# @deprecated — marcar campos/enums que serão removidos
type User {
  id: ID!
  email: String!
  username: String @deprecated(reason: "Use email instead. Will be removed in v3.")
}

# @skip e @include — lógica condicional no cliente
query GetUser($id: ID!, $includeOrders: Boolean!) {
  user(id: $id) {
    id
    name
    orders @include(if: $includeOrders) {
      edges { node { id status } }
    }
  }
}

# Custom directives (ex: autorização)
directive @auth(requires: Role = USER) on FIELD_DEFINITION

type Mutation {
  deleteUser(id: ID!): DeleteUserPayload! @auth(requires: ADMIN)
}
```

---

## 5. Subscriptions

```graphql
# Real-time updates via WebSocket
type Subscription {
  orderStatusChanged(orderId: ID!): Order!
  newMessageInChat(chatId: ID!): Message!
  systemNotification: Notification!
}
```

---

## 6. Anti-Padrões a Evitar

| Anti-Padrão | Problema | Solução |
|-------------|----------|---------|
| **Resolver fazendo query por item** | N+1 queries massacram o banco | Usar DataLoader para batch |
| **Schema espelhando banco** | Coupling — qualquer mudança no banco quebra a API | Projetar schema para o consumidor |
| **Campos genéricos `data: JSON`** | Perde tipo safety e introspecção | Criar tipos específicos |
| **Mutations sem payload** | Sem tratamento de erro estruturado | Sempre retornar payload com `errors` |
| **Campos non-null `!` em tudo** | Retorna erro quando campo é realmente opcional | Usar non-null apenas quando garantido |
| **Query infinitamente aninhada** | Consumidor pode fazer query de custo O(n^k) | Limitar profundidade e complexidade |
| **Introspecção em produção** | Expõe schema completo para atacantes | Desabilitar ou proteger em prod |

---

## 7. Segurança e Performance

### Limitar Complexidade de Query

```python
# Exemplo com Strawberry (Python)
from strawberry.extensions import QueryDepthLimiter, MaxAliasesLimiter

schema = strawberry.Schema(
    query=Query,
    extensions=[
        QueryDepthLimiter(max_depth=10),
        MaxAliasesLimiter(max_alias_count=15),
    ]
)
```

### Persisted Queries

```
# Em vez de enviar query completa:
POST /graphql
{ "query": "{ users { edges { node { id name } } } }" }

# Enviar hash da query (registrada previamente):
POST /graphql
{ "extensions": { "persistedQuery": { "sha256Hash": "abc123" } } }
```

### Depth Limiting

```
# Máximo recomendado: 10 níveis de profundidade
# Exemplo de query profunda a ser bloqueada:
query {
  user {
    orders {
      items {
        product {
          category {
            parent {
              parent { ... }
            }
          }
        }
      }
    }
  }
}
```

---

## 8. Ferramentas Recomendadas

| Ferramenta | Propósito |
|------------|-----------|
| **[GraphiQL](https://github.com/graphql/graphiql)** | IDE interativa para explorar schema |
| **[Apollo Studio](https://studio.apollographql.com/)** | Gestão de schema, monitoramento |
| **[Strawberry](https://strawberry.rocks/)** | Code-first GraphQL para Python |
| **[Ariadne](https://ariadnegraphql.org/)** | Schema-first GraphQL para Python |
| **[graphql-js](https://graphql.org/graphql-js/)** | Referência JavaScript |
| **[Pothos](https://pothos-graphql.dev/)** | Type-safe schema builder para TypeScript |
| **[graphql-inspector](https://the-guild.dev/graphql/inspector)** | Detectar breaking changes no schema |
