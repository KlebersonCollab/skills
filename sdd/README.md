# SDD — Spec-Driven Development

> Precision at scale. Rigor when needed, speed when possible.

[![Versão](https://img.shields.io/badge/Versão-1.4.0-blue)](#changelog)
[![Sub-skills](https://img.shields.io/badge/Sub--skills-5-brightgreen)](#-sub-skills)

---

## 📖 Visão Geral

Skill de **desenvolvimento orientado a especificações** com workflow modular e adaptativo. O SDD ajusta automaticamente a profundidade do processo à complexidade da tarefa (Auto-Sizing), garantindo rigor quando necessário e velocidade quando possível.

A versão `1.4.0` integra os princípios de **Harness Engineering**, introduzindo **Contratos de Desenvolvimento (SDC)** e validação determinística via **Sensores**.

> **Lei do SDD**: Se não está na spec, não existe. Se não foi verificado, não está pronto.

---

## ⚙️ Auto-Sizing

A profundidade do workflow é determinada pela **complexidade** da tarefa:

| Nível | Escopo | Fases | Sub-skills Utilizadas |
|-------|--------|-------|-----------------------|
| **Quick** | Bug fixes, configs, <3 arquivos | Implementar → Verificar | `implementer` |
| **Small** | Feature clara, <5 tarefas | Spec → Impl → Verificar | `orchestrator`, `implementer` |
| **Medium** | Feature + UI, <10 tarefas | Explorar → Spec → Impl → Verificar | `explorer`, `orchestrator` |
| **Large** | Multi-componente, novo módulo | Planejar → Explorar → Spec → Design → Impl → Verificar | Todas |
| **Complex** | Ambiguidade, alto risco | Todas as fases + UAT Interativo | Todas + `reviewer` |

---

## 🧩 Sub-skills

| Sub-skill | Arquivo | Responsabilidade |
|-----------|---------|------------------|
| **Explorer** | [sdd-explorer.skill.md](sdd-explorer.skill.md) | Mapeia codebases existentes gerando `STACK.md`, `ARCHITECTURE.md`, `CONVENTIONS.md` e outros artefatos de contexto. |
| **Planner** | [sdd-planner.skill.md](sdd-planner.skill.md) | Gerencia a visão do projeto e a **Memória Triádica** (`STATE.md`, `MEMORY.md`, `LEARNINGS.md`). |
| **Orchestrator** | [sdd-orchestrator.skill.md](sdd-orchestrator.skill.md) | Traduz requisitos em especificações técnicas (`spec.md`, `plan.md`, `tasks.md`) com rastreabilidade completa. |
| **Implementer** | [sdd-implementer.skill.md](sdd-implementer.skill.md) | Executa código atômico e test-driven, seguindo rigorosamente as specs e convenções do projeto. |
| **Reviewer** | [sdd-reviewer.skill.md](sdd-reviewer.skill.md) | Audita implementações contra critérios de aceitação com relatório de evidências e UAT. |

---

## 🔗 Knowledge Verification Chain

Ao pesquisar ou decidir, siga esta hierarquia estrita:

1. **Codebase Docs** → Consulte `.specs/codebase/` (gerado pelo Explorer).
2. **Project Docs** → Verifique READMEs, `PROJECT.md`, `ROADMAP.md`.
3. **Existing Code** → Leia o código-fonte diretamente.
4. **Web Search** → Apenas documentação oficial.
5. **Flag Uncertainty** → Se os passos 1-4 falharem, comunique a incerteza ao usuário.

---

## 📁 Estrutura de Diretórios (Obrigatória)

### Por Feature (`spec/` ou `.specs/features/[nome]/`)
- `spec.md` — Requisitos rastreáveis (FR-X) e critérios de aceitação.
- `plan.md` — Design técnico e schemas.
- `tasks.md` — Lista atômica de tarefas com status.

### Por Projeto (`.specs/`)
- `project/` — `PROJECT.md`, `ROADMAP.md`, `STATE.md`, `MEMORY.md`, `LEARNINGS.md`.
- `codebase/` — `STACK.md`, `ARCHITECTURE.md`, `CONVENTIONS.md`, `CONCERNS.md`.

---

## 📋 Protocolos Operacionais

### Safety Valve
Se uma tarefa identificada como **Quick** ou **Small** revelar complexidade oculta (>5 passos ou dependências profundas), **PARE**. Formalize a tarefa como **Large** com spec e plan completos.

### Persistent Memory (Triad State)
Sempre atualize o `STATE.md` (operacional), `MEMORY.md` (fatos) e `LEARNINGS.md` (sabedoria) via `sdd-planner` ao final de cada sessão ou após decisões importantes.

### Verification Standards
Uma feature **NÃO** está completa até que o **Reviewer** emita um veredito `APPROVED` baseado em evidências (paths e line numbers).

---

## 📝 Changelog

Consulte o [CHANGELOG.md](CHANGELOG.md) para o histórico completo de versões.
