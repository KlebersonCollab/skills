# API Versioning Strategies — Referência Completa

Guia de estratégias de versionamento de API e gerenciamento de breaking changes.

---

## 1. Por Que Versionar?

APIs são contratos públicos. Quebrar um contrato sem versionamento adequado:

- Derruba clientes em produção instantaneamente
- Gera retrabalho emergencial do lado do consumidor
- Destrói a confiança na sua API
- Viola o **Princípio da Compatibilidade de Cauda** (clients podem ser muito mais velhos que a API)

**Regra de ouro**: Toda API, desde o primeiro endpoint, deve ter uma estratégia de versionamento definida.

---

## 2. Estratégias de Versionamento

### 2.1 URL Versioning (Mais Comum)

```
/api/v1/users
/api/v2/users
```

**Prós:**
- Visível e explícito
- Fácil de testar no browser
- Fácil de fazer cache por versão
- Logs claros por versão

**Contras:**
- URLs ficam "feias" para puristas REST
- Pode ser confuso com versionamento de recursos individuais

**Quando usar:** APIs públicas, terceiros, mobile apps.

### 2.2 Header Versioning

```http
GET /api/users HTTP/1.1
Accept: application/vnd.empresa.v2+json
```

ou

```http
GET /api/users HTTP/1.1
X-API-Version: 2
```

**Prós:**
- URLs "limpas"
- Segue o princípio REST de content negotiation

**Contras:**
- Menos visível (difícil de testar no browser)
- Cache mais complexo (precisa do header no `Vary`)
- Clientes precisam saber qual header usar

**Quando usar:** APIs internas, quando URLs limpas são requisito.

### 2.3 Query Parameter Versioning

```
/api/users?version=2
/api/users?api-version=2026-04-14
```

**Prós:**
- Simples de implementar
- Fácil de testar

**Contras:**
- Pode conflitar com parâmetros de negócio
- Considerado antipadrão por alguns arquitetos

**Quando usar:** APIs simples, prototipagem, migração gradual.

### 2.4 Data-Based Versioning (Azure style)

```
/api/users?api-version=2026-04-14
```

**Prós:**
- Sem ambiguidade de qual versão é mais nova
- Facilita deprecation com datas específicas

**Quando usar:** APIs enterprise com ciclos de vida longos (estilo Azure API).

---

## 3. O Que É Uma Breaking Change?

### ❌ Breaking Changes (sempre requerem nova versão major)

| Categoria | Exemplos |
|-----------|---------|
| **Remoção** | Remover campo, endpoint ou parâmetro |
| **Renomeação** | Renomear campo `user_id` → `userId` |
| **Tipo** | Mudar tipo de campo `string` → `integer` |
| **Obrigatoriedade** | Tornar campo opcional em obrigatório |
| **Comportamento** | Mudar semântica de endpoint existente |
| **Auth** | Tornar endpoint público em privado |
| **Status code** | Mudar `200` → `201` em criação existente |

### ✅ Non-Breaking Changes (compatíveis retroativamente)

| Categoria | Exemplos |
|-----------|---------|
| **Adição** | Adicionar novo campo opcional na resposta |
| **Adição** | Adicionar novo endpoint |
| **Adição** | Adicionar novo parâmetro opcional |
| **Relaxamento** | Tornar campo obrigatório em opcional |
| **Performance** | Melhorar tempo de resposta |
| **Correção** | Corrigir bug sem mudar contrato |

---

## 4. Estratégia de Deprecation

### Ciclo de Vida de uma Versão

```
v1: CURRENT     → versão ativa, suportada
v1: DEPRECATED  → ainda funciona, mas aconselhado a migrar
v1: RETIRED     → removida (retorna 410 Gone)
```

### Comunicação de Deprecation

1. **Documentação** — Marcar versão/campo como deprecated com data de remoção
2. **Response Headers** — Incluir aviso em toda resposta
3. **Email/Changelog** — Notificar consumidores registrados
4. **Prazo mínimo** — Mínimo 6 meses para APIs públicas, 3 meses para internas

### Headers de Deprecation

```http
HTTP/1.1 200 OK
Deprecation: Sun, 01 Jan 2027 00:00:00 GMT
Sunset: Sun, 01 Jul 2027 00:00:00 GMT
Link: <https://docs.empresa.com/api/v2>; rel="successor-version"
```

### GraphQL Deprecation

```graphql
type User {
  id: ID!
  email: String!
  # Campo deprecated — use email
  username: String @deprecated(reason: "Use 'email' field. Será removido em v3 (2027-01-01).")
}
```

---

## 5. Estratégia de Migração

### Abordagem Gradual (Recomendada)

```
Fase 1: Lançar v2 em /api/v2/* (v1 continua funcionando)
Fase 2: Período de coexistência (6+ meses para público, 3+ meses para interno)
Fase 3: Deprecar v1 (headers + documentação + email)
Fase 4: Sunset v1 — retorna 410 Gone com instrução de migração
```

### Resposta de Endpoint Aposentado

```http
HTTP/1.1 410 Gone
Content-Type: application/json

{
  "error": {
    "code": "API_VERSION_RETIRED",
    "message": "A versão v1 desta API foi removida em 2027-01-01.",
    "migration_guide": "https://docs.empresa.com/api/migration/v1-to-v2"
  }
}
```

---

## 6. Versionamento Interno vs Externo

| Aspecto | API Pública | API Interna |
|---------|-------------|-------------|
| **Breaking change notice** | 6+ meses | 3+ meses |
| **Coexistência de versões** | Obrigatória | Recomendada |
| **Documentação** | OpenAPI spec pública | Interno/Confluence |
| **Consumidores notificados** | Email, status page | Slack, PR review |
| **Versões ativas simultâneas** | 2-3 max | 1-2 max |

---

## 7. Checklist de Versionamento

### Antes de Lançar uma Nova Versão
- [ ] Documentar todas as breaking changes
- [ ] Criar guia de migração (v_old → v_new)
- [ ] Definir data de sunset da versão anterior
- [ ] Comunicar consumidores com antecedência mínima
- [ ] Implementar headers de deprecation na versão anterior
- [ ] Atualizar documentação OpenAPI/GraphQL schema

### Após Lançar Nova Versão
- [ ] Monitorar uso por versão (logs/métricas)
- [ ] Identificar clientes ainda na versão antiga
- [ ] Enviar lembretes de migração conforme aproxima sunset
- [ ] Implementar 410 Gone na data de sunset
- [ ] Remover código da versão aposentada após período de retenção

---

## 8. Ferramentas

| Ferramenta | Propósito |
|------------|-----------|
| **[openapi-diff](https://github.com/OpenAPITools/openapi-diff)** | Detectar breaking changes entre specs OpenAPI |
| **[graphql-inspector](https://the-guild.dev/graphql/inspector)** | Detectar breaking changes em schemas GraphQL |
| **[Bump.sh](https://bump.sh/)** | Documentação versioned e diff automático |
| **[Optic](https://www.useoptic.com/)** | API versioning e governance |
