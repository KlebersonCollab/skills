# Changelog — API Designer

Todas as mudanças notáveis nesta skill serão documentadas neste arquivo.

## [1.1.0] - 2026-04-14

### Adicionado
- **Suporte a tRPC**: Referência completa em `references/trpc-patterns.md` para projetos TypeScript Fullstack.
- **Guia de Segurança (OWASP API Top 10)**: Novo arquivo `references/api-security-guide.md` focando em BOLA, Broken Auth, Mass Assignment, etc.
- **Rate Limiting**: Novo guia em `references/api-rate-limiting.md` cobrindo algoritmos (Token Bucket) e padronização de headers (status 429).
- **Árvore de Decisão**: Integrada ao `SKILL.md` para facilitar a escolha entre REST, GraphQL, tRPC e gRPC.
- **Padrões de Paginação Avançada**: Suporte a Cursor-based pagination (Relay style) e envelopes de resposta.

### Alterado
- **Workflow**: Atualizado para incluir validações de segurança e rate limit na fase inicial (Discover).
- **README.md**: Expandido para refletir o novo escopo da skill.
- **SKILL.md**: Incrementado para versão 1.1.0 com regras de qualidade "Segurança First".

---

## [1.0.0] - 2026-04-01

### Adicionado
- Versão inicial da skill com suporte a REST e GraphQL.
- Workflow em 4 fases: Discover, Model, Specify e Validate.
- Referências para design REST e Schema GraphQL.
- Estratégias básicas de versionamento.
- Implementation Playbook com exemplos práticos.
