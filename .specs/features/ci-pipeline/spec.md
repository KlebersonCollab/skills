# Specification: CI Pipeline (Skill Hub Guard)

**Feature ID:** FR-CI-PIPELINE
**Status:** DRAFT
**Complexity:** Small
**Category:** DevOps / Quality

## Goal
Garantir que toda contribuição ou alteração no Hub de Skills mantenha a integridade estrutural e de versionamento, executando validações automáticas em cada Push ou Pull Request.

## 🎯 Requisitos Funcionais

| ID | Título | Descrição |
|----|--------|-----------|
| FR-1 | **Trigger Automático** | O workflow deve rodar em `push` na branch `main` e em qualquer `pull_request`. |
| FR-2 | **Validação de Skills** | Executar `make validate` para garantir que o `SKILL.md` e outros arquivos mandatórios existam. |
| FR-3 | **Consistência de Versão** | Executar `make verify-vers` para garantir que o README bate com os SKILL.md. |
| FR-4 | **Relatório de Falha** | Caso algum check falhe, o build do GitHub deve falhar, bloqueando o merge. |

## ✅ Critérios de Aceitação (BDD)

### Cenário 1: Build bem-sucedido
**Given** que eu envio um código que passa em todos os validadores locais
**When** o GitHub Action é disparado
**Then** todos os steps (Validate, Verify Versions) devem ficar verdes.

### Cenário 2: Bloqueio por erro de validação
**Given** que eu envio uma nova skill sem o arquivo `CHANGELOG.md`
**When** o GitHub Action roda
**Then** o step `make validate` deve falhar
**And** o status do commit deve ficar vermelho.

## 🛠️ Restrições Técnicas
- Plataforma: GitHub Actions.
- Runtime: Python 3.12 (via `astral-sh/setup-uv`).
- Deve usar o `Makefile` existente como entrypoint.
