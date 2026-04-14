# API Designer

> Projetista de APIs REST e GraphQL — guia o agente a criar contratos de API intuitivos, escaláveis e consistentes, desde a modelagem de recursos até a documentação OpenAPI.

[![Versão](https://img.shields.io/badge/Versão-1.0.0-blue)](#changelog)
[![Categoria](https://img.shields.io/badge/Categoria-api--design-purple)](#)
[![Baseado em](https://img.shields.io/badge/Baseado_em-api--design--principles-orange)](https://github.com/sickn33/antigravity-awesome-skills/tree/main/skills/api-design-principles)

---

## 📖 Visão Geral

A skill **API Designer** capacita o agente a projetar APIs de alta qualidade de forma sistemática. Ela aplica os princípios consolidados de design REST e GraphQL para guiar todo o processo — desde a identificação dos consumidores e modelagem de recursos até a especificação completa do contrato e a validação do design.

O foco é garantir que o **contrato de API** seja aprovado e revisado **antes** que qualquer código de implementação seja escrito, eliminando retrabalho custoso e inconsistências.

### Por que usar esta skill?

- APIs mal projetadas geram retrabalho, breaking changes e frustração para os consumidores.
- Esta skill força o processo **Contract-First**, que é a melhor prática da indústria.
- Aplica princípios comprovados de REST e GraphQL de forma consistente e rastreável.

---

## ⚙️ Como Usar

Invoque a skill quando precisar projetar ou revisar uma API:

```
Use a skill api-designer para projetar a API de gerenciamento de pedidos.
Consumidores: app mobile (iOS/Android) e painel admin web.
Operações principais: criar pedido, listar pedidos por usuário, atualizar status, cancelar.
```

A skill executará 4 fases automaticamente:

1. **DISCOVER** — Entende contexto, consumidores e restrições
2. **MODEL** — Modela recursos, endpoints e schema
3. **SPECIFY** — Especifica contratos detalhados com erros, auth e versionamento
4. **VALIDATE** — Revisa consistência e aprova o design

---

## 📚 Recursos Incluídos

| Arquivo | Descrição |
|---------|-----------|
| [`resources/implementation-playbook.md`](resources/implementation-playbook.md) | Padrões de implementação, checklists e exemplos de código (REST + GraphQL) |
| [`references/rest-best-practices.md`](references/rest-best-practices.md) | Guia completo de boas práticas REST |
| [`references/graphql-schema-design.md`](references/graphql-schema-design.md) | Padrões de schema GraphQL e anti-padrões |
| [`references/api-versioning-strategies.md`](references/api-versioning-strategies.md) | Estratégias de versionamento e migração |

---

## 🎯 Casos de Uso

| Cenário | Use a skill? |
|---------|-------------|
| Projetando nova API REST do zero | ✅ Sim |
| Projetando schema GraphQL | ✅ Sim |
| Refatorando API para melhorar usabilidade | ✅ Sim |
| Revisando especificação antes da implementação | ✅ Sim |
| Estabelecendo padrões de API para o time | ✅ Sim |
| Implementando endpoint em framework específico (sem redesign) | ❌ Não |
| Trabalho exclusivo de infraestrutura sem contrato de API | ❌ Não |

---

## 📐 Princípios Aplicados

### REST
- **Resource-Oriented Architecture** — URLs são substantivos, HTTP verbs são ações
- **Stateless** — Cada request é auto-contido
- **HTTP Semântico** — Status codes corretos, idempotência respeitada
- **Paginação e Filtros** — Sempre presentes em coleções
- **Versionamento** — Estratégia definida desde o início

### GraphQL
- **Schema-First Development** — Schema antes dos resolvers
- **Tipagem Forte** — Enums, custom scalars, non-null explícito
- **Paginação Relay** — Cursor-based para coleções
- **DataLoaders** — Prevenção de N+1 consultas

---

## 📝 Changelog

Consulte o [CHANGELOG.md](CHANGELOG.md) para o histórico completo de versões.
