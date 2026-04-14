# REST Best Practices — Referência Completa

Guia de referência de boas práticas para design de APIs REST. Use em conjunto com o `resources/implementation-playbook.md`.

---

## 1. Nomenclatura de Endpoints

### ✅ Regras de Nomenclatura

| Regra | Correto | Incorreto |
|-------|---------|-----------|
| Plural para coleções | `/users` | `/user` |
| Substantivos, não verbos | `/orders` | `/getOrders` |
| Lowercase com hifens | `/user-profiles` | `/UserProfiles`, `/user_profiles` |
| Hierarquia lógica | `/users/{id}/orders` | `/getUserOrders` |
| Sem extensões de arquivo | `/api/users` | `/api/users.json` |

### Padrão de URLs

```
# Coleção
GET  /api/v1/users

# Recurso individual
GET  /api/v1/users/{user_id}

# Sub-recurso (relação)
GET  /api/v1/users/{user_id}/orders
GET  /api/v1/users/{user_id}/orders/{order_id}

# Ação especial (quando realmente necessário — use com parcimônia)
POST /api/v1/users/{user_id}/deactivate
POST /api/v1/orders/{order_id}/cancel
```

---

## 2. HTTP Status Codes

### 2xx — Sucesso

| Code | Nome | Quando Usar |
|------|------|-------------|
| `200` | OK | GET, PUT, PATCH com retorno de dados |
| `201` | Created | POST que cria recurso (inclua `Location` header) |
| `204` | No Content | DELETE, PUT/PATCH sem retorno de dados |

### 4xx — Erro do Cliente

| Code | Nome | Quando Usar |
|------|------|-------------|
| `400` | Bad Request | Request malformada, parâmetros inválidos |
| `401` | Unauthorized | Não autenticado (sem token ou token inválido) |
| `403` | Forbidden | Autenticado mas sem permissão |
| `404` | Not Found | Recurso não existe |
| `409` | Conflict | Conflito de estado (ex: email duplicado) |
| `422` | Unprocessable Entity | Dados válidos sintaticamente mas inválidos semanticamente |
| `429` | Too Many Requests | Rate limit excedido |

### 5xx — Erro do Servidor

| Code | Nome | Quando Usar |
|------|------|-------------|
| `500` | Internal Server Error | Erro inesperado no servidor |
| `502` | Bad Gateway | Serviço upstream indisponível |
| `503` | Service Unavailable | Servidor temporariamente indisponível |

---

## 3. Estrutura de Resposta Padronizada

### Resposta de Sucesso (recurso)
```json
{
  "data": {
    "id": "123",
    "name": "João Silva",
    "email": "joao@exemplo.com",
    "created_at": "2026-04-14T20:00:00Z"
  }
}
```

### Resposta de Sucesso (coleção)
```json
{
  "data": [...],
  "meta": {
    "total": 150,
    "page": 1,
    "page_size": 20,
    "pages": 8
  },
  "links": {
    "self": "/api/v1/users?page=1",
    "next": "/api/v1/users?page=2",
    "prev": null,
    "first": "/api/v1/users?page=1",
    "last": "/api/v1/users?page=8"
  }
}
```

### Resposta de Erro
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Os dados enviados são inválidos.",
    "details": [
      {
        "field": "email",
        "message": "Formato de e-mail inválido.",
        "value": "nao-e-um-email"
      }
    ],
    "timestamp": "2026-04-14T20:00:00Z",
    "path": "/api/v1/users",
    "request_id": "req_abc123"
  }
}
```

---

## 4. Headers Importantes

### Request Headers
```
Authorization: Bearer <jwt_token>
Content-Type: application/json
Accept: application/json
X-Request-ID: <uuid>         # Rastreabilidade
X-API-Version: 2             # (se usar header versioning)
```

### Response Headers
```
Content-Type: application/json; charset=utf-8
X-Request-ID: <uuid>         # Echo do request ID
X-RateLimit-Limit: 1000
X-RateLimit-Remaining: 999
X-RateLimit-Reset: 1713128400
Location: /api/v1/users/123  # Após POST 201
```

---

## 5. Filtragem, Ordenação e Busca

```
# Filtragem por campos
GET /api/v1/orders?status=pending&created_after=2026-01-01

# Ordenação
GET /api/v1/users?sort=name&order=asc
GET /api/v1/users?sort=-created_at    # Prefixo - para DESC

# Busca textual
GET /api/v1/users?search=joao

# Projeção de campos (sparse fieldsets)
GET /api/v1/users?fields=id,name,email

# Combinado
GET /api/v1/orders?status=confirmed&sort=-created_at&page=2&page_size=10
```

---

## 6. Segurança

### Autenticação

| Método | Quando Usar |
|--------|-------------|
| **JWT Bearer** | APIs stateless, microservices, SPAs |
| **API Key** | Integrações B2B, scripts, third-party |
| **OAuth 2.0** | Delegação de acesso, APIs públicas com escopos |
| **mTLS** | Comunicação serviço-a-serviço (service mesh) |

### Boas Práticas de Segurança

- **HTTPS obrigatório** — Nunca HTTP em produção
- **Rate limiting** — Por IP, por API key, por usuário
- **CORS configurado** — Apenas origens autorizadas
- **Validação de input** — Nunca confiar em dados do cliente
- **Sanitização** — Prevenir XSS e SQL Injection
- **Logs de auditoria** — Registrar todas as operações sensíveis

---

## 7. Idempotência

| Operação | Idempotente | Explicação |
|----------|-------------|------------|
| `GET /users/123` | ✅ | Múltiplas chamadas = mesmo resultado |
| `DELETE /users/123` | ✅ | Deletar algo que já foi deletado = 404 (aceitável) |
| `PUT /users/123` | ✅ | Substituir com mesmo payload = mesmo estado |
| `POST /users` | ❌ | Cria novo recurso a cada chamada |
| `PATCH /users/123` | Depende | Pode ou não ser idempotente |

**Padrão para POST idempotente** (quando necessário):
```
POST /api/v1/payments
Idempotency-Key: <uuid-gerado-pelo-cliente>
```

---

## 8. Documentação OpenAPI

### Estrutura mínima do `openapi.yaml`

```yaml
openapi: "3.1.0"
info:
  title: "API de Pedidos"
  version: "1.0.0"
  description: "API para gerenciamento de pedidos e usuários."
  contact:
    name: "Equipe de API"
    email: "api@empresa.com"

servers:
  - url: "https://api.empresa.com/v1"
    description: "Produção"
  - url: "https://api-staging.empresa.com/v1"
    description: "Staging"

paths:
  /users:
    get:
      summary: "Listar usuários"
      operationId: "listUsers"
      tags: ["users"]
      parameters:
        - name: page
          in: query
          schema:
            type: integer
            default: 1
        - name: page_size
          in: query
          schema:
            type: integer
            default: 20
            maximum: 100
      responses:
        "200":
          description: "Lista de usuários"
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/UserListResponse"
        "401":
          $ref: "#/components/responses/Unauthorized"

components:
  schemas:
    User:
      type: object
      required: [id, email, name]
      properties:
        id:
          type: string
          format: uuid
        email:
          type: string
          format: email
        name:
          type: string
        created_at:
          type: string
          format: date-time

  responses:
    Unauthorized:
      description: "Não autenticado"
      content:
        application/json:
          schema:
            $ref: "#/components/schemas/ErrorResponse"

  securitySchemes:
    BearerAuth:
      type: http
      scheme: bearer
      bearerFormat: JWT

security:
  - BearerAuth: []
```
