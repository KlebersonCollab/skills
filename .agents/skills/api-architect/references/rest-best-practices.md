# REST Best Practices — Complete Reference

Reference guide for REST API design best practices. Use in conjunction with `resources/implementation-playbook.md`.

---

## 1. Endpoint Naming

### ✅ Naming Rules

| Rule | Correct | Incorrect |
|-------|---------|-----------|
| Plural for collections | `/users` | `/user` |
| Nouns, not verbs | `/orders` | `/getOrders` |
| Lowercase with hyphens | `/user-profiles` | `/UserProfiles`, `/user_profiles` |
| Logical hierarchy | `/users/{id}/orders` | `/getUserOrders` |
| No file extensions | `/api/users` | `/api/users.json` |

### URL Patterns

```
# Collection
GET  /api/v1/users

# Individual resource
GET  /api/v1/users/{user_id}

# Sub-resource (relationship)
GET  /api/v1/users/{user_id}/orders
GET  /api/v1/users/{user_id}/orders/{order_id}

# Special action (when really necessary — use sparingly)
POST /api/v1/users/{user_id}/deactivate
POST /api/v1/orders/{order_id}/cancel
```

---

## 2. HTTP Status Codes

### 2xx — Success

| Code | Name | When to Use |
|------|------|-------------|
| `200` | OK | GET, PUT, PATCH with data return |
| `201` | Created | POST that creates resource (include `Location` header) |
| `204` | No Content | DELETE, PUT/PATCH without data return |

### 4xx — Client Error

| Code | Name | When to Use |
|------|------|-------------|
| `400` | Bad Request | Malformed request, invalid parameters |
| `401` | Unauthorized | Not authenticated (no token or invalid token) |
| `403` | Forbidden | Authenticated but no permission |
| `404` | Not Found | Resource does not exist |
| `409` | Conflict | State conflict (e.g., duplicate email) |
| `422` | Unprocessable Entity | Syntactically valid but semantically invalid data |
| `429` | Too Many Requests | Rate limit exceeded |

### 5xx — Server Error

| Code | Name | When to Use |
|------|------|-------------|
| `500` | Internal Server Error | Unexpected server error |
| `502` | Bad Gateway | Upstream service unavailable |
| `503` | Service Unavailable | Server temporarily unavailable |

---

## 3. Standardized Response Structure

### Envelopes vs. Direct Data

- **Direct Data**: Returns the resource directly (simpler, classic REST standard).
- **Envelope Pattern**: Wraps the data in a central object (better for metadata and consistency).

#### Envelope Pattern Example (Recommended for collections)
```json
{
  "success": true,
  "data": {
    "id": "123",
    "name": "João Silva"
  },
  "meta": {
    "request_id": "req_abc123",
    "timestamp": "2026-04-14T20:00:00Z"
  }
}
```

### Pagination Types

1. **Offset Pagination** (`page`, `page_size`): Simple, but inefficient on giant datasets.
2. **Cursor Pagination** (`after`, `first`): More efficient for fast-changing datasets (e.g., feeds) and massive volumes.
3. **Keyset Pagination** (`last_id`, `limit`): Uses the last sorting key.

#### Cursor-Based Example (Relay Style)
```json
{
  "data": [...],
  "meta": {
    "total": 150,
    "has_next_page": true,
    "start_cursor": "YXJyYXljb25uZWN0aW9uOjA=",
    "end_cursor": "YXJyYXljb25uZWN0aW9uOjE5"
  }
}
```

### Error Response
```json
{
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "The sent data is invalid.",
    "details": [
      {
        "field": "email",
        "message": "Invalid email format.",
        "value": "not-an-email"
      }
    ],
    "timestamp": "2026-04-14T20:00:00Z",
    "path": "/api/v1/users",
    "request_id": "req_abc123"
  }
}
```

---

## 4. Important Headers

### Request Headers
```
Authorization: Bearer <jwt_token>
Content-Type: application/json
Accept: application/json
X-Request-ID: <uuid>         # Traceability
X-API-Version: 2             # (if using header versioning)
```

### Response Headers
```
Content-Type: application/json; charset=utf-8
X-Request-ID: <uuid>         # Request ID echo
X-RateLimit-Limit: 1000
X-RateLimit-Remaining: 999
X-RateLimit-Reset: 1713128400
Location: /api/v1/users/123  # After POST 201
```

---

## 5. Filtering, Sorting, and Search

```
# Filtering by fields
GET /api/v1/orders?status=pending&created_after=2026-01-01

# Sorting
GET /api/v1/users?sort=name&order=asc
GET /api/v1/users?sort=-created_at    # - prefix for DESC

# Text search
GET /api/v1/users?search=joao

# Field projection (sparse fieldsets)
GET /api/v1/users?fields=id,name,email

# Combined
GET /api/v1/orders?status=confirmed&sort=-created_at&page=2&page_size=10
```

---

## 6. Security

### Authentication

| Method | When to Use |
|--------|-------------|
| **JWT Bearer** | Stateless APIs, microservices, SPAs |
| **API Key** | B2B integrations, scripts, third-party |
| **OAuth 2.0** | Access delegation, public APIs with scopes |
| **mTLS** | Service-to-service communication (service mesh) |

### Security Best Practices

- **HTTPS mandatory** — Never HTTP in production
- **Rate limiting** — By IP, by API key, by user
- **CORS configured** — Only authorized origins
- **Input validation** — Never trust client data
- **Sanitization** — Prevent XSS and SQL Injection
- **Audit logs** — Record all sensitive operations

---

## 7. Idempotency

| Operation | Idempotent | Explanation |
|----------|-------------|------------|
| `GET /users/123` | ✅ | Multiple calls = same result |
| `DELETE /users/123` | ✅ | Deleting something already deleted = 404 (acceptable) |
| `PUT /users/123` | ✅ | Replace with same payload = same state |
| `POST /users` | ❌ | Creates a new resource on each call |
| `PATCH /users/123` | Depend | May or may not be idempotent |

**Pattern for idempotent POST** (when necessary):
```
POST /api/v1/payments
Idempotency-Key: <client-generated-uuid>
```

---

## 8. OpenAPI Documentation

### Minimum `openapi.yaml` structure

```yaml
openapi: "3.1.0"
info:
  title: "Orders API"
  version: "1.0.0"
  description: "API for managing orders and users."
  contact:
    name: "API Team"
    email: "api@company.com"

servers:
  - url: "https://api.company.com/v1"
    description: "Production"
  - url: "https://api-staging.company.com/v1"
    description: "Staging"

paths:
  /users:
    get:
      summary: "List users"
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
          description: "Users list"
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
      description: "Not authenticated"
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
