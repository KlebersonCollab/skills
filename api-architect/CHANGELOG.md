# Changelog — API Architect

Todas as mudanças notáveis nesta skill serão documentadas neste arquivo.

## [1.3.0] - 2026-04-14

### Alterado
- **Workflow & References**: Melhoria no workflow para referenciar explicitamente o uso do `references/api-style.md` durante a fase de design.

## [1.2.0] - 2026-04-14

### Alterado
- **Renomeação**: Skill renomeada de `api-designer` para `api-architect` para refletir um escopo maior de governança e resiliência.
- **Style Guide**: Adição de `references/api-style.md` com padrões de envelopes de resposta, erros e paginação.
- **Resilience**: Inclusão de padrões de Timeouts, Retries e Circuit Breakers.

## [1.1.0] - 2026-04-14

### Adicionado
- **Suporte a tRPC**: Referência completa em `references/trpc-patterns.md` para projetos TypeScript Fullstack.
- **Guia de Segurança (OWASP API Top 10)**: Novo arquivo `references/api-security-guide.md` focando em BOLA, Broken Auth, Mass Assignment, etc.
- **Rate Limiting**: Novo guia em `references/api-rate-limiting.md` cobrindo algoritmos (Token Bucket) e padronização de headers (status 429).
- **Árvore de Decisão**: Integrada ao `SKILL.md` para facilitar a escolha entre REST, GraphQL, tRPC e gRPC.
- **Padrões de Paginação Avançada**: Suporte a Cursor-based pagination (Relay style) e envelopes de resposta.

---

## [1.0.0] - 2026-04-01

### Adicionado
- Versão inicial da skill com suporte a REST e GraphQL.
- Workflow em 4 fases: Discover, Model, Specify e Validate.
