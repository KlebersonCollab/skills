# 🧠 AI Agent Skills Hub

> Hub centralizado para desenvolvimento, armazenamento e evolução de **Skills** modulares para agentes de IA.

[![Skills](https://img.shields.io/badge/Skills-1-brightgreen)](#-skills-disponíveis)
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
| 1 | **[SDD](sdd/)** | Spec-Driven Development — Workflow modular e adaptativo para desenvolvimento orientado a especificações. | `1.0.0` |

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

1. **Crie um diretório** com o nome da skill na raiz do repositório.
2. **Crie o arquivo `SKILL.md`** — a definição técnica principal da skill.
3. **Crie o `README.md`** — documentação detalhada, sub-skills e exemplos de uso.
4. **Crie o `CHANGELOG.md`** — histórico de versões seguindo [Keep a Changelog](https://keepachangelog.com/pt-BR/).
5. **Atualize este `README.md`** adicionando a nova skill na tabela [Skills Disponíveis](#-skills-disponíveis).

### Template Mínimo para `SKILL.md`

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

### Próximos Passos
- [ ] Core Framework para criação padronizada de novas skills
- [ ] Skill de Pesquisa Avançada
- [ ] Skill de Automação de Infraestrutura
- [ ] Conectores para APIs externas

---

## 📄 Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).

---

<div align="center">

**Construído com 🧠 por [Kleberson Romero](https://github.com/KlebersonCollab)**

*Precision at scale. Rigor when needed, speed when possible.*

</div>
