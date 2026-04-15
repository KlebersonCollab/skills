# Implementation Playbook — API Designer

Padrões de implementação, checklists e exemplos de código para design de APIs REST, GraphQL e tRPC.

---

## 1. REST — Princípios Fundamentais

### Resource-Oriented Architecture

- Recursos são **substantivos** (`users`, `orders`, `products`), nunca verbos.
- Use HTTP methods para as ações (GET, POST, PUT, PATCH, DELETE).
- URLs representam hierarquias de recursos.
- Convenções de nomenclatura consistentes (plural para coleções).

### Semântica dos Métodos HTTP

| Método | Semântica | Idempotente | Safe |
|--------|-----------|-------------|------|
| `GET` | Recuperar recursos | ✅ | ✅ |
| `POST` | Criar novo recurso | ❌ | ❌ |
| `PUT` | Substituir recurso inteiro | ✅ | ❌ |
| `PATCH` | Atualizar campos parcialmente | ❌* | ❌ |
| `DELETE` | Remover recurso | ✅ | ❌ |

---

## 2. GraphQL — Princípios Fundamentais

### Schema-First Development

- Tipos definem o modelo de domínio.
- Queries para leitura de dados.
- Mutations para modificação de dados.
- Subscriptions para atualizações em tempo real.

---

## 3. tRPC — Princípios Fundamentais (Novo!)

### End-to-End Type Safety

- O contrato é o código TypeScript.
- Sem geração de código necessária.
- Ideal para Monorepos.

#### Exemplo de Router tRPC (v10+)
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

## 4. Padrões de Resiliência: Rate Limiting

### Exemplo de Middleware de Rate Limit (Express/Node)

```javascript
const rateLimit = require('express-rate-limit');

const apiLimiter = rateLimit({
  windowMs: 15 * 60 * 1000, // 15 minutos
  max: 100, // Limite de 100 requests por IP
  standardHeaders: true, // Retorna X-RateLimit-Limit e Remaining
  legacyHeaders: false,
  handler: (req, res) => {
    res.status(429).json({
      error: {
        code: "RATE_LIMIT_EXCEEDED",
        message: "Muitas requisições. Tente novamente em 15 minutos."
      }
    });
  }
});

app.use('/api/', apiLimiter);
```

---

## 5. Checklists de Validação

### ✅ Checklist de Segurança (OWASP Focus)

- [ ] **BOLA Check**: O endpoint valida se `user_id` do token é dono do recurso solicitado?
- [ ] **Mass Assignment**: O DTO de entrada proíbe campos como `is_admin` ou `role`?
- [ ] **Broken Auth**: Tokens JWT usam RS256 e possuem `exp` (expiração)?
- [ ] **Rate Limit**: Existe proteção configurada para endpoints sensíveis?

### ✅ Checklist REST (v1.1)

- [ ] Todos os endpoints usam substantivos no plural.
- [ ] Resposta de erro padronizada (`{ error: { code, message } }`).
- [ ] Paginação cursor-based para feeds ou grandes coleções.
- [ ] Estratégia de versionamento clara na URL ou Header.

### ✅ Checklist tRPC

- [ ] Inputs são validados com Zod?
- [ ] Procedimentos protegidos usam middlewares de autenticação?
- [ ] Erros usam a classe `TRPCError` com códigos semânticos?

---

## 6. Pitfalls Comuns

| Problema | Causa | Solução |
|----------|-------|---------|
| Vazamento de ID (BOLA) | Confiar apenas no ID enviado na URL | Validar propriedade do recurso no backend. |
| Payload muito grande | Falta de limite no body-parser | Configurar limite de tamanho de request. |
| N+1 em GraphQL/tRPC | Múltiplas chamadas ao banco por item | Usar DataLoaders ou Joins eficientes. |
| Falta de Headers de Versão | Clientes quebram após update | Sempre versionar ou usar deprecation strategy. |
