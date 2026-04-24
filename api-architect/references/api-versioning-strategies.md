# API Versioning Strategies — Complete Reference

Guide to API versioning strategies and management of breaking changes.

---

## 1. Why Version?

APIs are public contracts. Breaking a contract without proper versioning:

- Instantly drops clients in production
- Generates emergency rework on the consumer side
- Destroys trust in your API
- Violates the **Tail Compatibility Principle** (clients can be much older than the API)

**Golden rule**: Every API, from the first endpoint, must have a defined versioning strategy.

---

## 2. Versioning Strategies

### 2.1 URL Versioning (Most Common)

```
/api/v1/users
/api/v2/users
```

**Pros:**
- Visible and explicit
- Easy to test in the browser
- Easy to cache by version
- Clear logs by version

**Cons:**
- URLs look "ugly" to REST purists
- Can be confusing with individual resource versioning

**When to use:** Public APIs, third parties, mobile apps.

### 2.2 Header Versioning

```http
GET /api/users HTTP/1.1
Accept: application/vnd.company.v2+json
```

or

```http
GET /api/users HTTP/1.1
X-API-Version: 2
```

**Pros:**
- "Clean" URLs
- Follows the REST principle of content negotiation

**Cons:**
- Less visible (hard to test in the browser)
- More complex cache (needs the header in `Vary`)
- Clients need to know which header to use

**When to use:** Internal APIs, when clean URLs are a requirement.

### 2.3 Query Parameter Versioning

```
/api/users?version=2
/api/users?api-version=2026-04-14
```

**Pros:**
- Simple to implement
- Easy to test

**Cons:**
- Can conflict with business parameters
- Considered an anti-pattern by some architects

**When to use:** Simple APIs, prototyping, gradual migration.

### 2.4 Date-Based Versioning (Azure style)

```
/api/users?api-version=2026-04-14
```

**Pros:**
- No ambiguity about which version is newer
- Facilitates deprecation with specific dates

**When to use:** Enterprise APIs with long life cycles (Azure API style).

---

## 3. What is a Breaking Change?

### ❌ Breaking Changes (always require a new major version)

| Category | Examples |
|-----------|---------|
| **Removal** | Remove field, endpoint, or parameter |
| **Renaming** | Rename field `user_id` → `userId` |
| **Type** | Change field type `string` → `integer` |
| **Obligatoriness** | Make optional field mandatory |
| **Behavior** | Change semantics of an existing endpoint |
| **Auth** | Make public endpoint private |
| **Status code** | Change `200` → `201` in existing creation |

### ✅ Non-Breaking Changes (backwards compatible)

| Category | Examples |
|-----------|---------|
| **Addition** | Add new optional field in the response |
| **Addition** | Add new endpoint |
| **Addition** | Add new optional parameter |
| **Relaxation** | Make mandatory field optional |
| **Performance** | Improve response time |
| **Correction** | Fix bug without changing contract |

---

## 4. Deprecation Strategy

### Life Cycle of a Version

```
v1: CURRENT     → active, supported version
v1: DEPRECATED  → still works, but advised to migrate
v1: RETIRED     → removed (returns 410 Gone)
```

### Deprecation Communication

1. **Documentation** — Mark version/field as deprecated with removal date
2. **Response Headers** — Include notice in every response
3. **Email/Changelog** — Notify registered consumers
4. **Minimum period** — Minimum 6 months for public APIs, 3 months for internal

### Deprecation Headers

```http
HTTP/1.1 200 OK
Deprecation: Sun, 01 Jan 2027 00:00:00 GMT
Sunset: Sun, 01 Jul 2027 00:00:00 GMT
Link: <https://docs.company.com/api/v2>; rel="successor-version"
```

### GraphQL Deprecation

```graphql
type User {
  id: ID!
  email: String!
  # Deprecated field — use email
  username: String @deprecated(reason: "Use 'email' field. Will be removed in v3 (2027-01-01).")
}
```

---

## 5. Migration Strategy

### Gradual Approach (Recommended)

```
Phase 1: Launch v2 at /api/v2/* (v1 continues to work)
Phase 2: Coexistence period (6+ months for public, 3+ months for internal)
Phase 3: Deprecate v1 (headers + documentation + email)
Phase 4: Sunset v1 — returns 410 Gone with migration instruction
```

### Retired Endpoint Response

```http
HTTP/1.1 410 Gone
Content-Type: application/json

{
  "error": {
    "code": "API_VERSION_RETIRED",
    "message": "Version v1 of this API was removed on 2027-01-01.",
    "migration_guide": "https://docs.company.com/api/migration/v1-to-v2"
  }
}
```

---

## 6. Internal vs. External Versioning

| Aspect | Public API | Internal API |
|---------|-------------|-------------|
| **Breaking change notice** | 6+ months | 3+ months |
| **Version coexistence** | Mandatory | Recommended |
| **Documentation** | Public OpenAPI spec | Internal/Confluence |
| **Consumers notified** | Email, status page | Slack, PR review |
| **Simultaneous active versions** | 2-3 max | 1-2 max |

---

## 7. Versioning Checklist

### Before Launching a New Version
- [ ] Document all breaking changes
- [ ] Create migration guide (v_old → v_new)
- [ ] Define sunset date for the previous version
- [ ] Communicate with consumers with minimum notice
- [ ] Implement deprecation headers in the previous version
- [ ] Update OpenAPI documentation/GraphQL schema

### After Launching a New Version
- [ ] Monitor usage by version (logs/metrics)
- [ ] Identify clients still on the old version
- [ ] Send migration reminders as sunset approaches
- [ ] Implement 410 Gone on the sunset date
- [ ] Remove code for the retired version after the retention period

---

## 8. Tools

| Tool | Purpose |
|------------|-----------|
| **[openapi-diff](https://github.com/OpenAPITools/openapi-diff)** | Detect breaking changes between OpenAPI specs |
| **[graphql-inspector](https://the-guild.dev/graphql/inspector)** | Detect breaking changes in GraphQL schemas |
| **[Bump.sh](https://bump.sh/)** | Versioned documentation and automatic diff |
| **[Optic](https://www.useoptic.com/)** | API versioning and governance |
