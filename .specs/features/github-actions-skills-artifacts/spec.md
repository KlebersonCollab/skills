# Specification: GitHub Actions Skills Artifacts

Este documento define os requisitos para a automação da geração de artefatos (.zip) contendo as skills do projeto formatadas para diferentes agentes (Claude, Gemini e Agent).

## Problem Statement

Atualmente, as skills são mantidas na raiz do projeto e replicadas manualmente para as pastas dos agentes (`.claude/skills/`, `.gemini/skills/`, `.agent/skills/`). Isso dificulta a manutenção e a distribuição das skills para os usuários finais, que precisam de um pacote pronto para uso.

## Goals

- Centralizar a "fonte da verdade" das skills nas pastas raiz.
- Automatizar a distribuição das skills para as pastas de cada agente no pipeline.
- Gerar 3 artefatos .zip individuais para Claude, Gemini e Agent.
- Facilitar o download e uso imediato pelos usuários.

## Acceptance Criteria (BDD)

### Scenario: Geração de artefatos no Push para a Main
**Given** que novas alterações foram integradas na branch `main`
**When** o pipeline do GitHub Actions for disparado
**Then** ele deve coletar todas as pastas na raiz que possuem um arquivo `SKILL.md`
**And** ele deve criar uma estrutura temporária para cada agente:
    - `.claude/` contendo `CLAUDE.md` e `skills/` (com todas as skills coletadas)
    - `.gemini/` contendo `GEMINI.md` e `skills/` (com todas as skills coletadas)
    - `.agent/` contendo `AGENT.md` e `skills/` (com todas as skills coletadas)
**And** ele deve gerar 3 arquivos ZIP: `claude-skills.zip`, `gemini-skills.zip`, `agent-skills.zip`
**And** os arquivos ZIP devem estar disponíveis como artefatos do workflow.

### Scenario: Integridade das Skills nos Artefatos
**Given** que uma skill como `api-architect` existe na raiz
**When** o artefato `claude-skills.zip` for gerado
**Then** o conteúdo de `skills/api-architect/` no ZIP deve ser idêntico ao da raiz
**And** o arquivo `CLAUDE.md` deve estar na raiz do ZIP.

## Constraints

- O pipeline não deve modificar o repositório original (não deve fazer push de volta).
- O processo deve ser idempotente e rápido.
- Deve suportar a adição de novas skills na raiz automaticamente.
