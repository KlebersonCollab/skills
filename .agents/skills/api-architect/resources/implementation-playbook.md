# Implementation Playbook — API Designer

Implementation patterns, checklists, and code examples for REST, GraphQL, and tRPC API design.

---

## 1. REST — Fundamental Principles

### Resource-Oriented Architecture

- Resources are **nouns** (`users`, `orders`, `products`), never verbs.
- Use HTTP methods for actions (GET, POST, PUT, PATCH, DELETE).
- URLs represent resource hierarchies.
- Consistent naming conventions (plural for collections).

### HTTP Method Semantics

| Method | Semantics | Idempotent | Safe |
|--------|-----------|-------------|------|
| `GET` | Retrieve resources | ✅ | ✅ |
| `POST` | Create new resource | ❌ | ❌ |
| `PUT` | Replace entire resource | ✅ | ❌ |
| `PATCH` | Partially update fields | ❌* | ❌ |
| `DELETE` | Remove resource | ✅ | ❌ |

---

## 2. GraphQL — Fundamental Principles

### Schema-First Development

- Types define the domain model.
- Queries for data reading.
- Mutations for data modification.
- Subscriptions for real-time updates.

---

## 3. tRPC — Fundamental Principles (New!)

### End-to-End Type Safety

- The contract is the TypeScript code.
- No code generation necessary.
- Ideal for Monorepos.

#### tRPC Router Example (v10+)
```typescript
import { initTRPC, TRPCError } from '@trpc/server';
import { z } from 'zod';

const t = initTRPC.create();

export const appRouter = t.router({
  getUser: t.procedure
    .input(z.string())
    .query(async (opts) => {
      const { input } = opts;
      const user = await db.user.findById(input);
      if (!user) throw new TRPCError({ code: 'NOT_FOUND' });
      return user;
    }),
  createUser: t.procedure
    .input(z.object({ name: z.string(), email: z.string().email() }))
    .mutation(async (opts) => {
      const { input } = opts;
      return await db.user.create(input);
    }),
});
```

---

## 4. Resilience Patterns: Rate Limiting

### Rate Limit Middleware Example (Express/Node)

```javascript
const rateLimit = require('express-rate-limit');

const apiLimiter = rateLimit({
  windowMs: 15 * 60 * 1000, // 15 minutes
  max: 100, // Limit to 100 requests per IP
  standardHeaders: true, // Returns X-RateLimit-Limit and Remaining
  legacyHeaders: false,
  handler: (req, res) => {
    res.status(429).json({
      error: {
        code: "RATE_LIMIT_EXCEEDED",
        message: "Too many requests. Try again in 15 minutes."
      }
    });
  }
});

app.use('/api/', apiLimiter);
```

---

## 5. Validation Checklists

### ✅ Security Checklist (OWASP Focus)

- [ ] **BOLA Check**: Does the endpoint validate if the token's `user_id` owns the requested resource?
- [ ] **Mass Assignment**: Does the input DTO prohibit fields like `is_admin` or `role`?
- [ ] **Broken Auth**: Do JWT tokens use RS256 and have `exp` (expiration)?
- [ ] **Rate Limit**: Is there protection configured for sensitive endpoints?

### ✅ REST Checklist (v1.1)

- [ ] All endpoints use plural nouns.
- [ ] Standardized error response (`{ error: { code, message } }`).
- [ ] Cursor-based pagination for feeds or large collections.
- [ ] Clear versioning strategy in URL or Header.

### ✅ tRPC Checklist

- [ ] Are inputs validated with Zod?
- [ ] Do protected procedures use authentication middlewares?
- [ ] Do errors use the `TRPCError` class with semantic codes?

---

## 6. Common Pitfalls

| Problem | Cause | Solution |
|----------|-------|---------|
| ID Leak (BOLA) | Trusting only the ID sent in the URL | Validate resource ownership in the backend. |
| Payload too large | Lack of limit in body-parser | Configure request size limit. |
| N+1 in GraphQL/tRPC | Multiple DB calls per item | Use DataLoaders or efficient Joins. |
| Missing Version Headers | Clients break after update | Always version or use deprecation strategy. |
