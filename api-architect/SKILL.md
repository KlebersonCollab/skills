---
name: api-architect
version: 1.2.0
description: "Arquiteto de APIs — guia o agente a projetar sistemas de API interoperáveis, seguros e resilientes, definindo contratos e padrões de governança."
category: api-design
---

# API Architect

> Projetista de APIs de alta performance e guardião da governança técnica. "Design is for developers; architecture is for the system."

---

## Goal

Capacitar o agente a projetar ecossistemas de API robustos aplicando princípios de design REST, GraphQL e tRPC, garantindo interoperabilidade através de **Style Guides** rigorosos, segurança (OWASP) e resiliência (Rate Limiting, Circuit Breaker).

A skill assegura que a API não seja apenas "funcional", mas **consistente**, **segura** e **sustentável** ao longo do tempo.

---

## Quando Usar Esta Skill

- Projetando novas APIs do zero (REST, GraphQL, tRPC).
- Definindo o **API Style Guide** para garantir consistência entre múltiplos serviços.
- Refatorando APIs legadas para padrões modernos de segurança e resiliência.
- Implementando estratégias de governança (versionamento, paginação, rate limit).
- Revisando segurança contra vulnerabilidades críticas (BOLA/BFLA).

---

## Workflow (4 Fases)

### Fase 1: STRATEGY — Definição de Protocolo
1.  **Consumer Analysis**: Identificar quem consome (Web, Mobile, B2B).
2.  **Protocol Selection**: Escolher entre REST (Público), GraphQL (Flexível) ou tRPC (Monorepo TS).
3.  **Governance Baseline**: Definir versão inicial, política de deprecation e rate limiting.

### Fase 2: DESIGN — Modelagem e Style Guide
1.  **Resource Modeling**: Definir entidades e relações (plural, substantivos).
2.  **Consistency Layer**: Aplicar o `API Style Guide` (formato de erro padrão, envelopes de resposta).
3.  **Schema Definition**: Escrever OpenAPI (REST), SDL (GraphQL) ou Zod (tRPC).

### Fase 3: RESILIENCE — Segurança e Performance
1.  **Security Audit**: Aplicar checklist de OWASP API Top 10 (AuthN/AuthZ).
2.  **Stability Patterns**: Definir timeouts, retries e Circuit Breakers.
3.  **Data Filtering**: Implementar paginação (Cursor/Offset) e projeção de campos.

### Fase 4: VALIDATE — Verificação Automatizada
1.  **Contract Testing**: Validar se o schema gerado é válido e sem erros.
2.  **Security Scan**: Revisar permissões de objeto (BOLA) em todos os endpoints.
3.  **Documentation Handoff**: Entregar o contrato aprovado e o guia de implementação.

---

## Output Structure

A execução desta skill deve resultar nos seguintes artefatos mandatórios:

| Artefato | Formato | Descrição |
|----------|---------|-----------|
| **API Spec** | `.yaml` / `.graphql` / `.ts` | Contrato técnico completo (OpenAPI, SDL ou Zod). |
| **API Style Guide** | `.md` | Definição de padrões de resposta, erro e nomenclatura. |
| **Security Checklist** | `.md` | Relatório de conformidade com OWASP e padrões de Auth. |
| **Resilience Policy** | `.md` | Definição de limites de taxa, timeouts e estratégia de cache. |

---

## Quality Rules

- **Contract-First**: O código de implementação só deve ser escrito após a aprovação do contrato.
- **Uniform Error Handling**: Erros devem seguir um padrão único de estrutura e status code.
- **Security by Default**: Todo endpoint deve ser protegido por padrão (necessita de opt-out explícito).
- **YAGNI in Data**: Retorne apenas o necessário; evite over-fetching através de projeções ou campos opcionais.

## Prohibited

- NUNCA exponha IDs sequenciais (use UUIDs ou HashIDs).
- NUNCA use verbos em URLs de recursos (ex: `/getUsers` -> use `GET /users`).
- NUNCA ignore o versionamento; APIs públicas devem ter versão na URL ou Headers.
- NUNCA retorne logs ou stack traces em respostas de erro de produção.

---

## Referências

- [`references/rest-best-practices.md`](references/rest-best-practices.md) — Guia de design REST.
- [`references/api-security-guide.md`](references/api-security-guide.md) — Segurança e OWASP API Top 10.
- [`references/api-style.md`](references/api-style.md) — (Novo) Padrões de consistência de interface.
- [`references/api-rate-limiting.md`](references/api-rate-limiting.md) — Resiliência e controle de taxa.
- [`resources/implementation-playbook.md`](resources/implementation-playbook.md) — Checklists de implementação.
