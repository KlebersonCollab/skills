---
name: api-designer
version: 1.1.0
description: "Projetista de APIs REST, GraphQL e tRPC — guia o agente a criar contratos de API intuitivos, escaláveis e seguros, desde a modelagem até a segurança (OWASP)."
category: api-design
---

# API Designer

> Projetista de APIs REST, GraphQL e tRPC — guia o agente a criar contratos de API intuitivos, escaláveis e seguros, desde a modelagem até a segurança (OWASP).

---

## Goal

Capacitar o agente a projetar APIs de alta qualidade aplicando os princípios de design REST, GraphQL e tRPC de forma sistemática: modelagem de recursos, escolha de estilo (Decision Tree), convenções de nomenclatura, tratamento de erros, versionamento, paginação, segurança (OWASP API Top 10) e rate limiting.

A skill garante que o contrato de API seja **consistente**, **seguro** e **preparado para evolução** antes de qualquer linha de código de implementação ser escrita.

---

## Quando Usar Esta Skill

- Projetando novas APIs (REST, GraphQL ou tRPC) do zero.
- Escolhendo a tecnologia ideal para o contexto (Monorepo vs Polyglot).
- Refatorando APIs existentes para melhorar segurança e resiliência.
- Estabelecendo padrões de design e rate limiting para o time.
- Revisando especificações de API contra vulnerabilidades comuns (BOLA/IDOR).
- Gerando documentação OpenAPI/Swagger ou schemas GraphQL/tRPC.

---

## Árvore de Decisão: Qual Estilo Escolher?

| Contexto | Estilo Recomendado | Por que? |
|----------|--------------------|----------|
| **Monorepo TS / Fullstack TS** | **tRPC** | Type-safety ponta a ponta sem geração de código. |
| **API Pública / Interoperabilidade** | **REST + OpenAPI** | Padrão de mercado, universal, cache HTTP eficiente. |
| **Dados Complexos / Múltiplos Frontends** | **GraphQL** | Flexibilidade para o cliente, evita over-fetching. |
| **Microserviços Internos (Alta Perf)** | **gRPC** | Baixa latência, serialização binária eficiente. |

---

## Workflow (4 Fases)

### Fase 1: DISCOVER — Entender o Contexto
1. Identificar os **consumidores** da API (frontend, mobile, third-party, B2B).
2. Aplicar a **Árvore de Decisão** para escolher o estilo (REST, GraphQL, tRPC).
3. Mapear os **casos de uso principais** (user stories ou operações CRUD críticas).
4. Definir **limites de segurança** e **rate limiting** necessários.

### Fase 2: MODEL — Modelar Recursos e Schema
- **REST**: Substantivos no plural, verbos HTTP, hierarquias de URL, paginação (Offset/Cursor).
- **GraphQL**: Schema SDL, tipos fortes, mutations com payloads, paginação Relay.
- **tRPC**: Definição de Routers e Procedimentos tipados com Zod.

### Fase 3: SPECIFY — Especificar Contratos e Segurança
- Detalhar request/response, envelopes e mensagens de erro.
- Aplicar o checklist de **Segurança (OWASP API Top 10)**.
- Definir estratégia de **Rate Limiting** (headers e status 429).
- Especificar autenticação (JWT, API Key, OAuth2) e versionamento.

### Fase 4: VALIDATE — Revisar e Aprovar
- Aplicar o `resources/implementation-playbook.md` como checklist.
- Verificar consistência e proteção contra Mass Assignment.
- Confirmar idempotência e documentação OpenAPI/SDL.

---

## Princípios de Design Aplicados

### REST & tRPC
- **Resource-Oriented**: URLs/Procedimentos focados em entidades.
- **Stateless**: Cada request contém toda a informação necessária.
- **Type-Safe**: Validação rigorosa de inputs (Zod/OpenAPI).

### GraphQL
- **Schema-First**: Contrato define o modelo antes da implementação.
- **Paginação Cursor-based**: Padrão para coleções grandes.
- **Mutation Payloads**: Retorno estruturado de dados e erros.

---

## Quality Rules

- **Segurança First**: Validar permissão de objeto (BOLA) em todos os endpoints/procedimentos.
- **Type Safety**: Priorizar contratos tipados (Zod, SDL, OpenAPI).
- **Consistência**: Envelopes, paginação e erros devem ser uniformes.
- **Rate Limit**: Toda API deve ter uma estratégia de proteção contra abuso.
- **Contract-First**: O contrato deve estar aprovado antes da implementação iniciar.

---

## Referências

Consulte os recursos detalhados para padrões, checklists e templates:

- [`references/rest-best-practices.md`](references/rest-best-practices.md) — Guia completo de design REST e paginação.
- [`references/graphql-schema-design.md`](references/graphql-schema-design.md) — Padrões de schema GraphQL e anti-padrões.
- [`references/trpc-patterns.md`](references/trpc-patterns.md) — APIs tipadas para ecossistemas TypeScript.
- [`references/api-security-guide.md`](references/api-security-guide.md) — Segurança e OWASP API Top 10.
- [`references/api-rate-limiting.md`](references/api-rate-limiting.md) — Resiliência e controle de taxa.
- [`references/api-versioning-strategies.md`](references/api-versioning-strategies.md) — Estratégias de versionamento e migração.
- [`resources/implementation-playbook.md`](resources/implementation-playbook.md) — Padrões de implementação e checklists.
