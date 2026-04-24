# GraphQL Schema Design — Complete Reference

Reference guide for GraphQL schema patterns and anti-patterns. Use in conjunction with `resources/implementation-playbook.md`.

---

## 1. Typing Fundamentals

### Scalar Types

```graphql
# Built-in scalars
String
Int
Float
Boolean
ID

# Custom scalars (declare and implement resolver)
scalar DateTime   # ISO 8601: "2026-04-14T20:00:00Z"
scalar Date       # ISO 8601: "2026-04-14"
scalar Money      # Cents as Int or decimal string
scalar UUID       # "550e8400-e29b-41d4-a716-446655440000"
scalar JSON       # Arbitrary data (use with caution)
scalar Upload     # For file upload (multipart)
```

### Non-Null and Lists

```graphql
# Non-null: mandatory field (never returns null)
name: String!

# Nullable list of non-null items
items: [OrderItem!]

# Non-null list of non-null items (the most common in collections)
items: [OrderItem!]!

# Nullable list of nullable items (avoid)
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

## 2. Schema Patterns

### Pattern: Relay Connection (cursor-based pagination)

```graphql
# Relay pattern for pagination — always use for large collections
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

# Usage in query
type Query {
  # Forward pagination (first/after)
  users(
    first: Int = 20
    after: String
    search: String
    sortBy: UserSortField
    sortOrder: SortOrder
  ): UserConnection!

  # Backward pagination (last/before) — optional
  # last: Int
  # before: String
}
```

### Pattern: Mutation Payloads

```graphql
# ✅ Correct: mutations return payload with errors
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
  user: User          # null if failed
  errors: [UserError!]!  # empty if success
}

type UserError {
  field: String       # null for global errors
  message: String!
  code: String!       # e.g., "EMAIL_TAKEN", "WEAK_PASSWORD"
}

# ❌ Wrong: mutation returns type directly (no error handling)
# createUser(input: CreateUserInput!): User!
```

### Pattern: Input Types

```graphql
# Specific inputs per operation (do not reuse between create/update)
input CreateUserInput {
  email: String!
  name: String!
  password: String!
}

input UpdateUserInput {
  id: ID!
  name: String        # nullable = optional field on update
  email: String
}

# For partial updates (patch), use nullable fields
# For full updates (put), use non-null fields
```

### Pattern: Interfaces and Unions

```graphql
# Interface — types with fields in common
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

# Union — return that can be of different types
union SearchResult = User | Product | Order

type Query {
  search(query: String!): [SearchResult!]!
}
```

---

## 3. Advanced Queries

### Fragments and Inline Fragments

```graphql
# Reusable fragment
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

# Inline fragment for unions
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
# @deprecated — mark fields/enums to be removed
type User {
  id: ID!
  email: String!
  username: String @deprecated(reason: "Use email instead. Will be removed in v3.")
}

# @skip and @include — conditional logic on client
query GetUser($id: ID!, $includeOrders: Boolean!) {
  user(id: $id) {
    id
    name
    orders @include(if: $includeOrders) {
      edges { node { id status } }
    }
  }
}

# Custom directives (e.g., authorization)
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

## 6. Anti-Patterns to Avoid

| Anti-Pattern | Problem | Solution |
|-------------|----------|---------|
| **Resolver doing query per item** | N+1 queries slaughter the database | Use DataLoader for batching |
| **Schema mirroring database** | Coupling — any DB change breaks the API | Design schema for the consumer |
| **Generic fields `data: JSON`** | Loses type safety and introspection | Create specific types |
| **Mutations without payload** | No structured error handling | Always return payload with `errors` |
| **Non-null fields `!` everywhere** | Returns error when field is actually optional | Use non-null only when guaranteed |
| **Infinitely nested query** | Consumer can make O(n^k) cost query | Limit depth and complexity |
| **Introspection in production** | Exposes full schema to attackers | Disable or protect in prod |

---

## 7. Security and Performance

### Limit Query Complexity

```python
# Example with Strawberry (Python)
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
# Instead of sending the full query:
POST /graphql
{ "query": "{ users { edges { node { id name } } } }" }

# Send the hash of the query (previously registered):
POST /graphql
{ "extensions": { "persistedQuery": { "sha256Hash": "abc123" } } }
```

### Depth Limiting

```
# Recommended maximum: 10 levels of depth
# Example of deep query to be blocked:
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

## 8. Recommended Tools

| Tool | Purpose |
|------------|-----------|
| **[GraphiQL](https://github.com/graphql/graphiql)** | Interactive IDE to explore schema |
| **[Apollo Studio](https://studio.apollographql.com/)** | Schema management, monitoring |
| **[Strawberry](https://strawberry.rocks/)** | Code-first GraphQL for Python |
| **[Ariadne](https://ariadnegraphql.org/)** | Schema-first GraphQL for Python |
| **[graphql-js](https://graphql.org/graphql-js/)** | JavaScript reference |
| **[Pothos](https://pothos-graphql.dev/)** | Type-safe schema builder for TypeScript |
| **[graphql-inspector](https://the-guild.dev/graphql/inspector)** | Detect breaking changes in schema |
