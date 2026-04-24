# tRPC Patterns — Complete Reference

Guide to patterns for APIs using tRPC in TypeScript monorepos.

---

## 1. What is tRPC?

tRPC (TypeScript Remote Procedure Call) allows you to build APIs with **End-to-End Type Safety** without the need to generate code or maintain manual schemas (like OpenAPI or GraphQL SDL).

**The contract is the TypeScript code itself.**

---

## 2. When to Use tRPC?

### ✅ Use tRPC if:
- The project is **Fullstack TypeScript** (Frontend and Backend in TS).
- You are in a **Monorepo** (or can share types via a package).
- You want maximum development speed and safe refactoring.
- The API consumer is controlled by you (not a public API for third parties).

### ❌ DO NOT use tRPC if:
- The API will be consumed by other languages (Python, Go, Java, etc.).
- You need a market-standard public API (prefer REST/OpenAPI).
- Frontend and Backend are in completely isolated repositories without type sharing.

---

## 3. Procedure Structure

### Queries (Read)
Use for fetching data. Must be pure and without side effects.
```typescript
publicProcedure
  .input(z.object({ id: z.string() }))
  .query(async ({ input, ctx }) => {
    return await ctx.db.user.findUnique({ where: { id: input.id } });
  });
```

### Mutations (Write)
Use for creating, updating, or deleting data.
```typescript
protectedProcedure
  .input(createUserSchema)
  .mutation(async ({ input, ctx }) => {
    return await ctx.db.user.create({ data: input });
  });
```

---

## 4. Design Best Practices

### 4.1 Router Organization
Divide the API into logical sub-routers to maintain scalability.
```typescript
// appRouter.ts
export const appRouter = router({
  user: userRouter,
  post: postRouter,
  auth: authRouter,
});
```

### 4.2 Middlewares and Context
Use middlewares for reusable logic (Auth, Logging, Error Handling).
```typescript
const isAuthed = middleware(({ next, ctx }) => {
  if (!ctx.session) {
    throw new TRPCError({ code: 'UNAUTHORIZED' });
  }
  return next({
    ctx: { session: ctx.session },
  });
});

export const protectedProcedure = publicProcedure.use(isAuthed);
```

### 4.3 Validation with Zod
Always validate `input` using Zod schemas to ensure data arrives correctly at the resolver.

---

## 5. Error Handling

tRPC maps errors to corresponding HTTP status codes.

| tRPC Error Code | HTTP Status |
|-----------------|-------------|
| `BAD_REQUEST` | 400 |
| `UNAUTHORIZED` | 401 |
| `FORBIDDEN` | 403 |
| `NOT_FOUND` | 404 |
| `TIMEOUT` | 408 |
| `CONFLICT` | 409 |
| `INTERNAL_SERVER_ERROR` | 500 |

---

## 6. Versioning in tRPC

Since tRPC shares types, "versioning" usually happens via refactoring in the Monorepo. If you change a field in the Backend, the Frontend will break at compile time, which is the biggest advantage of tRPC.

For gradual migrations:
1. Add the new field as optional.
2. Update the Frontend to use the new field.
3. Remove the old field after migration.
