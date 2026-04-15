# 🏛️ API Architect Skill

> Skill para o desenho, governança e arquitetura de ecossistemas de API interoperáveis, seguros e resilientes.

---

## 📖 Visão Geral

A skill **API Architect** (anteriormente `api-designer`) eleva o processo de criação de APIs de um simples "desenho de contrato" para uma arquitetura completa. Ela integra padrões de design modernos (REST, GraphQL, tRPC) com camadas críticas de governança, segurança e resiliência.

Inspirada por padrões de mercado e reforçada por práticas de governança de API, esta skill garante que cada endpoint seja consistente em todo o ecossistema.

---

## 🚀 Principais Recursos

- **API Style Guide**: Definição de padrões de resposta (envelopes), erros e nomenclatura para garantir que a API seja previsível.
- **Multi-Protocolo**: Suporte nativo para decidir entre REST, GraphQL e tRPC com base no contexto técnico.
- **Segurança (OWASP)**: Auditoria sistemática contra os principais riscos de segurança de API (BOLA, IDOR, Mass Assignment).
- **Resiliência Pragmática**: Padrões de Rate Limiting, Timeouts e Circuit Breakers integrados ao design.
- **Contract-First**: Foco total na especificação (OpenAPI, SDL, Zod) antes da implementação.

---

## 📂 Estrutura da Skill

```
api-architect/
├── SKILL.md                 # Definição técnica e workflow
├── README.md                # Esta documentação
├── CHANGELOG.md             # Histórico de versões
├── references/              # Padrões e guias de referência
│   ├── api-style.md         # Style Guide (Envelopes, Erros, Paginação)
│   ├── rest-best-practices.md
│   ├── api-security-guide.md
│   └── api-rate-limiting.md
└── resources/
    └── implementation-playbook.md
```

---

## 🛠️ Como Usar

1. **Invoque a skill** ao planejar uma nova funcionalidade que exija exposição de dados via rede.
2. Siga as **4 Fases** (Strategy, Design, Resilience, Validate) para garantir que nenhum detalhe crítico seja esquecido.
3. Utilize os **referenciais** para definir o formato das respostas e erros logo no início da modelagem.

---

## 📄 Licença

Este módulo faz parte do [Skills Hub](https://github.com/KlebersonCollab/skills) e está sob licença MIT.
