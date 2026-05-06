# Feature: Skill Synchronization Script

## Context
O usuário necessita de um script para automatizar a cópia/sincronização das pastas de "skills" localizadas na raiz do repositório para os diretórios de governança `.agents/skills` e `.gemini/skills`.

## Goal
Criar um script Bash (`bin/sync-skills.sh`) que:
1. Identifique as skills mandatórias.
2. Sincronize-as com os destinos especificados.
3. Garanta a integridade dos arquivos.

## Acceptance Criteria
- [x] Script `bin/sync-skills.sh` criado e executável.
- [x] As skills listadas são copiadas para `.agents/skills/`.
- [x] As skills listadas são copiadas para `.gemini/skills/`.
- [x] O script utiliza caminhos relativos à raiz do projeto.
- [x] O script reporta o progresso de forma clara.
- [x] Integração com `Makefile` via comando `make sync`.
- [x] Sincronização do `GLOBAL_MANDATES.md` para arquivos de instrução dos agentes (.gemini, .claude, .agents).

## Skills List
- architecture
- benchmark-expert
- brainstorming
- clean-code-mentor
- django-expert
- fastapi-expert
- flutter-fvm
- git-workflow
- observability-expert
- python-uv
- sdd
- skill-factory
- token-distiller
- youtube-transcript

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "SKILL-SYNC"
phase: "VERIFY"
status: COMPLETED
last_update: "2026-05-06T13:47:30Z"
evidence_checksum: "NONE"
```
