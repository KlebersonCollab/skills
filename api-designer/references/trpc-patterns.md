# tRPC Patterns — Referência Completa

Guia de padrões para APIs usando tRPC em monorepos TypeScript.

---

## 1. O que é tRPC?

tRPC (TypeScript Remote Procedure Call) permite construir APIs com **segurança de tipos ponta a ponta** (End-to-End Type Safety) sem a necessidade de gerar código ou manter esquemas manuais (como OpenAPI ou GraphQL SDL).

**O contrato é o próprio código TypeScript.**

---

## 2. Quando Usar tRPC?

### ✅ Use tRPC se:
- O projeto é **Fullstack TypeScript** (Frontend e Backend em TS).
- Você está em um **Monorepo** (ou pode compartilhar tipos via package).
- Você quer velocidade de desenvolvimento máxima e refatoração segura.
- O consumidor da API é controlado por você (não é uma API pública para terceiros).

### ❌ NÃO use tRPC se:
- A API será consumida por outras linguagens (Python, Go, Java, etc.).
- Você precisa de uma API pública padrão de mercado (prefira REST/OpenAPI).
- O Frontend e Backend estão em repositórios completamente isolados sem compartilhamento de tipos.

---

## 3. Estrutura de Procedimentos

### Queries (Leitura)
Use para buscar dados. Devem ser puras e sem efeitos colaterais.
```typescript
publicProcedure
  .input(z.object({ id: z.string() }))
  .query(async ({ input, ctx }) => {
    return await ctx.db.user.findUnique({ where: { id: input.id } });
  });
```

### Mutations (Escrita)
Use para criar, atualizar ou deletar dados.
```typescript
protectedProcedure
  .input(createUserSchema)
  .mutation(async ({ input, ctx }) => {
    return await ctx.db.user.create({ data: input });
  });
```

---

## 4. Melhores Práticas de Design

### 4.1 Organização de Routers
Divida a API em sub-routers lógicos para manter a escalabilidade.
```typescript
// appRouter.ts
export const appRouter = router({
  user: userRouter,
  post: postRouter,
  auth: authRouter,
});
```

### 4.2 Middlewares e Contexto
Use middlewares para lógica reutilizável (Auth, Logging, Error Handling).
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

### 4.3 Validação com Zod
Sempre valide o `input` usando schemas Zod para garantir que os dados cheguem corretos ao resolver.

---

## 5. Tratamento de Erros

O tRPC mapeia erros para status codes HTTP correspondentes.

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

## 6. Versionamento no tRPC

Como o tRPC compartilha tipos, o "versionamento" geralmente acontece via refatoração no Monorepo. Se você mudar um campo no Backend, o Frontend quebrará em tempo de compilação, o que é a maior vantagem do tRPC.

Para migrações graduais:
1. Adicione o novo campo como opcional.
2. Atualize o Frontend para usar o novo campo.
3. Remova o campo antigo após a migração.
