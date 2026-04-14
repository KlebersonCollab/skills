---
name: api-designer
version: 1.0.0
description: "Projetista de APIs REST e GraphQL — guia o agente a criar contratos de API intuitivos, escaláveis e consistentes, desde a modelagem de recursos até a documentação OpenAPI."
category: api-design
---

# API Designer

> Projetista de APIs REST e GraphQL — guia o agente a criar contratos de API intuitivos, escaláveis e consistentes, desde a modelagem de recursos até a documentação OpenAPI.

---

## Goal

Capacitar o agente a projetar APIs de alta qualidade aplicando os princípios de design REST e GraphQL de forma sistemática: modelagem de recursos, escolha de estilo, convenções de nomenclatura, tratamento de erros, versionamento, paginação, autenticação e documentação.

A skill garante que o contrato de API seja **consistente**, **developer-friendly** e **preparado para evolução** antes de qualquer linha de código de implementação ser escrita.

---

## Quando Usar Esta Skill

- Projetando novas APIs REST ou GraphQL do zero
- Refatorando APIs existentes para melhorar usabilidade
- Estabelecendo padrões de design de API para o time
- Revisando especificações de API antes da implementação
- Migrando entre paradigmas (REST → GraphQL ou vice-versa)
- Gerando documentação OpenAPI/Swagger a partir de contratos
- Otimizando APIs para casos de uso específicos (mobile, integrações terceiras)

## Quando NÃO Usar Esta Skill

- Você precisa apenas de orientação de implementação em um framework específico
- O trabalho é exclusivamente de infraestrutura, sem contratos de API
- Você não tem permissão para alterar ou versionar interfaces públicas existentes

---

## Output Structure

A skill produz os seguintes artefatos:

| Artefato | Descrição |
|----------|-----------|
| **Contrato de API** | Especificação OpenAPI 3.x (YAML/JSON) ou schema GraphQL (SDL) |
| **Mapa de Recursos** | Diagrama de recursos, relacionamentos e endpoints |
| **Guia de Erros** | Catálogo padronizado de códigos e mensagens de erro |
| **Checklist de Revisão** | Lista de verificação pré-implementação |

---

## Workflow (4 Fases)

### Fase 1: DISCOVER — Entender o Contexto
1. Identificar os **consumidores** da API (frontend, mobile, third-party, B2B).
2. Mapear os **casos de uso principais** (user stories ou operações CRUD críticas).
3. Definir **restrições** (SLA, limites de payload, requisitos de segurança).
4. Escolher o **estilo de API** adequado: REST, GraphQL ou híbrido.

### Fase 2: MODEL — Modelar Recursos e Schema
**Para REST:**
- Identificar os recursos (substantivos, não verbos).
- Mapear relacionamentos e hierarquias de URL.
- Definir verbos HTTP corretos por operação.
- Planejar paginação, filtros e ordenação.

**Para GraphQL:**
- Projetar o schema SDL (tipos, queries, mutations, subscriptions).
- Definir tipos de input e payload de mutations.
- Planejar DataLoaders para evitar N+1.
- Desenhar estratégia de paginação (cursor-based / Relay spec).

### Fase 3: SPECIFY — Especificar Contratos
- Detalhar request/response para cada endpoint ou operação.
- Definir estratégia de **versionamento** (URL, header, query param).
- Especificar **autenticação** (Bearer JWT, API Key, OAuth2).
- Padronizar **respostas de erro** (código, mensagem, detalhes).
- Documentar rate limits e headers obrigatórios.

### Fase 4: VALIDATE — Revisar e Aprovar
- Aplicar o `resources/implementation-playbook.md` como checklist.
- Verificar consistência de nomenclatura (plural, snake_case vs camelCase).
- Confirmar idempotência dos métodos HTTP.
- Revisar com stakeholders antes de iniciar implementação.

---

## Princípios de Design Aplicados

### REST
- **Resource-Oriented**: URLs são substantivos (`/users`, `/orders`), não verbos.
- **HTTP Semântico**: GET (safe+idempotent), PUT/DELETE (idempotent), POST (not idempotent).
- **Stateless**: Cada request contém toda a informação necessária.
- **HATEOAS** (quando aplicável): Respostas incluem links para ações disponíveis.

### GraphQL
- **Schema-First**: Schema define o contrato antes dos resolvers.
- **Tipagem Forte**: Enums, custom scalars, tipos non-null explícitos.
- **Paginação Relay**: Cursor-based para coleções grandes.
- **Mutation Payloads**: Retornam dados e erros estruturados.

---

## Quality Rules

- **Consistência Acima de Tudo**: Nomenclatura, estrutura de erro e paginação devem ser uniformes em toda a API.
- **Contract-First**: O contrato deve estar aprovado antes da implementação iniciar.
- **Versioning-Ready**: Toda API deve ter uma estratégia de versionamento definida desde o início.
- **Erros Acionáveis**: Mensagens de erro devem ser claras e incluir detalhes que permitam ao consumidor corrigir o problema.
- **Documentação como Código**: OpenAPI spec ou SDL GraphQL devem ser gerados e mantidos junto ao código.

## Prohibited

- NUNCA usar verbos em endpoints REST (`/createUser`, `/getOrders`).
- NUNCA retornar sempre HTTP 200 com erros embarcados no body (exceto GraphQL com convenção explícita).
- NUNCA projetar sem considerar os consumidores reais da API.
- NUNCA ignorar breaking changes — toda alteração de contrato público exige versionamento.
- NUNCA iniciar implementação sem contrato revisado.
- NUNCA espelhar diretamente o schema do banco de dados nos endpoints de API.

---

## Referências

Consulte os recursos detalhados para padrões, checklists e templates:

- [`resources/implementation-playbook.md`](resources/implementation-playbook.md) — Padrões de implementação, checklists e exemplos de código
- [`references/rest-best-practices.md`](references/rest-best-practices.md) — Guia completo de design REST
- [`references/graphql-schema-design.md`](references/graphql-schema-design.md) — Padrões de schema GraphQL e anti-padrões
- [`references/api-versioning-strategies.md`](references/api-versioning-strategies.md) — Estratégias de versionamento e migração
