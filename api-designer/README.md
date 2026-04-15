# API Designer — Skill

Projetista de APIs REST, GraphQL e tRPC — guia o agente a criar contratos de API intuitivos, escaláveis e seguros, desde a modelagem de recursos até a segurança (OWASP API Top 10).

---

## 🚀 Visão Geral

Esta skill transforma o agente em um arquiteto de APIs sênior. Ela não apenas ajuda a definir endpoints, mas orienta a escolha da tecnologia certa para cada contexto, garante que a API seja resiliente a abusos (Rate Limiting) e segura contra ataques modernos (OWASP).

### Capacidades Principais:
- **Design Multiestilo**: Suporte nativo a REST, GraphQL e tRPC.
- **Decision Tree**: Guia para escolha de arquitetura (Monorepo vs Polyglot).
- **Segurança Nativa**: Foco em BOLA/IDOR, Broken Auth e Mass Assignment.
- **Resiliência**: Padrões de Rate Limiting e tratamento de erros acionáveis.
- **Documentação**: Geração de OpenAPI 3.x e Schemas SDL/tRPC.

---

## 📂 Estrutura da Skill

```text
api-designer/
├── SKILL.md                 # Definição, Workflow e Regras (v1.1.0)
├── README.md                # Este guia
├── CHANGELOG.md             # Histórico de versões
├── references/              # Conhecimento Teórico
│   ├── rest-best-practices.md
│   ├── graphql-schema-design.md
│   ├── trpc-patterns.md     # NOVO!
│   ├── api-security-guide.md # NOVO!
│   ├── api-rate-limiting.md  # NOVO!
│   └── api-versioning-strategies.md
└── resources/               # Guia Prático
    └── implementation-playbook.md
```

---

## 🛠️ Como Usar

1. **Ative a skill**: `activate_skill api-designer`
2. **Siga o Workflow**: O agente passará pelas fases de DISCOVER, MODEL, SPECIFY e VALIDATE.
3. **Consulte as Referências**: Se tiver dúvidas sobre segurança ou tRPC, peça ao agente para consultar os arquivos na pasta `references/`.

---

## 🎯 Exemplos de Comandos

- *"Projete uma API REST para um sistema de e-commerce seguindo os padrões de segurança da skill."*
- *"Escolha a melhor tecnologia de API para um monorepo TypeScript e crie o contrato inicial."*
- *"Revise minha especificação de API atual contra o OWASP API Top 10."*
