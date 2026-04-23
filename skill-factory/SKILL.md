---
name: skill-factory
version: 1.1.0
description: "Core Framework para criação padronizada de novas skills com scaffolding, validação e registro automatizados."
category: skill-management
---

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar qualquer execução técnica:
0. **Mode Check**: Verificar o modo operacional atual (`.hub-mode`) e aplicar as diretrizes da skill `token-distiller`.
1. **Context Check**: Você reidratou o contexto lendo `STATE.md`, `MEMORY.md` e `LEARNINGS.md`?
2. **Spec Check**: O arquivo `spec.md` existe com requisitos e Critérios de Aceitação (ACs) claros? (BDD mandatório para Medium+).
3. **Plan Check**: O arquivo `plan.md` define a arquitetura, schemas e inclui diagramas **Mermaid**?
4. **Contract Check**: O arquivo `contract.md` foi estabelecido com os sensores de validação?
5. **Task Check**: A lista de tarefas em `tasks.md` está detalhada e atomizada?

---
# Skill Factory: Core Framework

> Every skill deserves a solid foundation. Consistency breeds quality.

---

## Goal

Automatizar a criação padronizada de novas skills com qualidade garantida, eliminando erros humanos no scaffolding e assegurando conformidade com os padrões do Skills Hub.

---

## Auto-Sizing

A profundidade do scaffolding é determinada pela **estrutura** da skill:

| Mode | Escopo | Arquivos Gerados | Sub-skill Utilizada |
|------|--------|-------------------|---------------------|
| **Quick** | Skill simples, sem sub-skills | `SKILL.md`, `README.md`, `CHANGELOG.md` | `skill-factory-bootstrap` |
| **Standard** | Skill com sub-skills | Todos acima + `<skill>-<sub>.skill.md` (N) | `skill-factory-bootstrap` |

Após o scaffolding, a validação é **sempre** executada (ambos os modos):

| Phase | Responsável |
|-------|-------------|
| **Validate** | `skill-factory-validator` |

---

## Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `skill_name` | string | ✅ | Nome da skill (slug: lowercase, hifenizado, sem espaços). Ex: `deep-research` |
| `description` | string | ✅ | Descrição concisa do que a skill faz. |
| `category` | string | ✅ | Categoria funcional (livre). Ex: `research`, `automation`, `development-workflow` |
| `sub_skills` | list[string] | ❌ | Lista de nomes de sub-skills a serem criadas. Se vazio, modo Quick. |

---

## The Modular Engine

Esta skill delega tarefas para sub-skills especializadas:

- **[Bootstrap](skill-factory-bootstrap.skill.md)**: Gera o scaffolding completo da nova skill (diretório, arquivos, conteúdo padronizado).
- **[Validator](skill-factory-validator.skill.md)**: Audita a skill criada contra o checklist de conformidade e emite relatório.

---

## Workflow (3 Phases)

### Phase 1: BOOTSTRAP
Executar `skill-factory-bootstrap` com os parâmetros recebidos.
- Cria o diretório `<skill_name>/` na raiz do repositório.
- Gera todos os arquivos a partir dos templates padronizados.

### Phase 2: VALIDATE
Executar `skill-factory-validator` na skill recém-criada.
- Verifica estrutura, frontmatter, conteúdo e naming.
- Se **NON-COMPLIANT**, corrigir antes de prosseguir.

### Phase 3: REGISTER
Atualizar o `README.md` raiz do repositório:
1. Adicionar a nova skill na tabela **Skills Disponíveis**.
2. Incrementar o badge de contagem de skills.

---

## Version Policy

Toda nova skill **SEMPRE** inicia na versão `1.0.0`. Não existe suporte para versões pré-release.

---

## Output Structure

A execução desta skill resulta nos seguintes artefatos mandatórios para cada nova skill criada:

| Artefato | Descrição |
|----------|-----------|
| `SKILL.md` | Definição técnica e workflow da skill. |
| `README.md` | Documentação de apresentação e uso. |
| `CHANGELOG.md` | Histórico de versões (inicial 1.0.0). |
| `references/` | (Opcional) Guias de referência detalhados. |
| `examples/` | (Opcional) Amostras de código e configurações. |

## Quality Rules

- **Template-First**: Toda skill gerada deve seguir os templates padronizados definidos na sub-skill `bootstrap`.
- **Validation-Always**: Nenhuma skill é considerada criada até que o `validator` emita `COMPLIANT`.
- **Registry-Updated**: O `README.md` raiz é a fonte da verdade para skills existentes e deve estar sempre atualizado.
- **Self-Consistent**: O nome no frontmatter deve corresponder ao nome do diretório.

## Prohibited

- NUNCA criar uma skill sem executar a fase de validação.
- NUNCA registrar uma skill NON-COMPLIANT no README raiz.
- NUNCA usar caracteres especiais, espaços ou uppercase no `skill_name`.
- NUNCA criar arquivos fora do diretório da skill (exceto o registro no README raiz).
