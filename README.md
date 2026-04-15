# 🧠 AI Agent Skills Hub

> Hub centralizado para desenvolvimento, armazenamento e evolução de **Skills** modulares para agentes de IA.

[![Skills](https://img.shields.io/badge/Skills-6-brightgreen)](#-skills-disponíveis)
[![Licença](https://img.shields.io/badge/Licença-MIT-blue)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Em_Desenvolvimento-yellow)](#-roadmap)

---

## 📖 Sobre o Projeto

Este repositório é a **fonte da verdade** para todas as habilidades (skills) utilizadas por agentes de IA. Cada skill é um módulo independente, documentado e versionado, que pode ser integrado a qualquer agente compatível.

Cada skill possui seu próprio diretório com documentação completa — incluindo `SKILL.md`, `README.md` e `CHANGELOG.md`.

---

## 🚀 Skills Disponíveis

| # | Skill | Descrição | Versão |
|---|-------|-----------|--------|
| 1 | **[SDD](sdd/)** | Spec-Driven Development — Workflow modular e adaptativo com **Persistent Memory Protocol** (Triad State). | `1.2.0` |
| 2 | **[Skill Factory](skill-factory/)** | Core Framework para criação padronizada de novas skills com scaffolding, validação e registro automatizados. | `1.1.0` |
| 3 | **[Python com UV](python-uv/)** | Desenvolvimento Python profissional (Django Pro, FastAPI) com UV — gerenciador 10-100x mais rápido. | `2.5.0` |
| 4 | **[API Architect](api-architect/)** | Arquiteto de APIs — design interoperável, seguro e resiliente (OpenAPI, GraphQL, tRPC) com Style Guide. | `1.3.0` |
| 5 | **[Brainstorming](brainstorming/)** | Facilitador de design e resolução de problemas — exploração profunda e validação antes da implementação. | `1.1.0` |
| 6 | **[Architecture](architecture/)** | Arquiteto de Sistemas — design pragmático, análise de trade-offs e registros de decisão (ADR). | `1.0.0` |

> 💡 Clique no nome da skill para acessar sua documentação completa.

---

## 🏗️ Estrutura do Repositório

```
skills/
├── README.md                    # ← Você está aqui
├── LICENSE                      # Licença MIT
├── .specs/                      # Especificações do projeto (SDD)
│   ├── project/
│   │   ├── PROJECT.md           # Visão e objetivos
│   │   └── ROADMAP.md           # Metas e milestones
│   └── codebase/
│       └── STACK.md             # Stack tecnológica
│
└── <skill>/                     # 🧩 Cada skill em seu diretório
    ├── README.md                # Documentação detalhada
    ├── SKILL.md                 # Definição técnica principal
    ├── CHANGELOG.md             # Histórico de versões
    └── *.skill.md               # Sub-skills (se aplicável)
```

---

## 📐 Como Criar uma Nova Skill

Utilize a skill **[Skill Factory](skill-factory/)** para criar novas skills de forma padronizada:

1. **Invoque a Skill Factory** informando `skill_name`, `description`, `category` e opcionalmente `sub_skills`.
2. O **Bootstrap** gera automaticamente todos os arquivos (`SKILL.md`, `README.md`, `CHANGELOG.md` e sub-skills).
3. O **Validator** audita a conformidade da skill gerada.
4. O **Registry** atualiza este `README.md` com a nova skill na tabela.

> 📖 Consulte a [documentação completa do Skill Factory](skill-factory/) para exemplos e detalhes.

### Template Mínimo para `SKILL.md` (referência)

```markdown
---
name: nome-da-skill
version: 1.0.0
description: "Descrição concisa do que a skill faz."
category: categoria
---

# Nome da Skill

> Descrição breve.

## Goal
[Objetivo principal da skill]

## Output Structure
[O que a skill produz]

## Quality Rules
[Regras de qualidade obrigatórias]

## Prohibited
[O que a skill NÃO deve fazer]
```

---

## 🗺️ Roadmap

Consulte o [ROADMAP.md](.specs/project/ROADMAP.md) para a visão completa de evolução do projeto.



---

## 📄 Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).

---

<div align="center">

**Construído com 🧠 por [Kleberson Romero](https://github.com/KlebersonCollab)**

*Precision at scale. Rigor when needed, speed when possible.*

</div>
