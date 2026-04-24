# API Style Guide: Interface and Consistency Standards

This guide defines the mandatory standards for all APIs designed by the **API Architect** skill, ensuring predictability and ease of consumption.

---

## 1. Response Envelopes (JSON)

### Success
Every success response must return data within a `data` object.

```json
{
  "data": {
    "id": "uuid-123",
    "name": "Kleberson Romero",
    "email": "dev@example.com"
  }
}
```

### Lists and Pagination
Use the metadata standard for collections.

```json
{
  "data": [ ... ],
  "meta": {
    "total": 100,
    "page": 1,
    "limit": 10,
    "has_next": true
  }
}
```

---

## 2. Error Handling

Every error response (4xx, 5xx) must follow the format below to facilitate handling in the frontend/clients.

```json
{
  "error": {
    "code": "ENTITY_NOT_FOUND",
    "message": "The user with ID 123 was not found.",
    "details": {
      "resource": "User",
      "id": "123"
    }
  }
}
```

- **code**: String in `SCREAMING_SNAKE_CASE` (unique for each type of error).
- **message**: Human-friendly message.
- **details**: (Optional) Object with technical details or validation fields.

---

## 3. Naming and Types

- **Fields**: Use `snake_case` for JSON keys (consistency with most persistence libraries).
- **IDs**: Always use strings (UUIDs, ULIDs, or HashIDs). Never expose sequential database IDs.
- **Dates**: Always use ISO 8601 format (`YYYY-MM-DDTHH:mm:ssZ`).
- **Booleans**: Must start with state verbs (e.g., `is_active`, `has_permission`).

---

## 4. HTTP Verbs and Status Codes

| Verb | Action | Success Status | Common Error Status |
|-------|------|----------------|-------------------|
| **GET** | Retrieve resource | 200 OK | 404 Not Found |
| **POST** | Create resource | 201 Created | 400 Bad Request |
| **PUT** | Replace resource | 200 OK | 400 Bad Request |
| **PATCH** | Partially update | 200 OK | 400 Bad Request |
| **DELETE** | Remove resource | 204 No Content | 404 Not Found |

---

## 5. Versioning

- **Via URL**: Priority for public APIs (e.g., `/v1/users`).
- **Via Headers**: Acceptable for internal APIs or minor versions (e.g., `X-API-Version: 1.2`).

---

## 6. Pagination (Cursor vs. Offset)

- **Offset**: Simple, for small collections (e.g., `?page=1&limit=10`).
- **Cursor**: Recommended for feeds or large collections (e.g., `?after=cursor_id`).
