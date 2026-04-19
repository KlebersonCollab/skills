# Plan: Multi-Agent Orchestrator

## Goal
Criar um protocolo de orquestração multi-agente (Swarm) que utilize os artefatos do framework SDD como "estado compartilhado" para permitir handoffs eficientes entre personas especializadas (Arquiteto, Dev, QA).

## Architectural Decision
- **Protocol**: [ADR-003] Swarm/Handoff via `handoff.md` e `.specs/`.
- **Persistence**: Memória de longo prazo persistida via `harness-expert` e sincronizada no `STATE.md`.
- **Handoff**: O agente atual gera um arquivo de instrução para o próximo e encerra a sessão.

## Phases
1. **Design**: Definir os papéis (Personas) e suas responsabilidades no Hub.
2. **Skill Implementation**: Criar a skill `multi-agent-orchestrator` com prompts de orquestração.
3. **Handoff Artifacts**: Padronizar o template de `handoff.md` dentro de `.specs/features/<feature>/`.
4. **Validation**: Simular um ciclo completo (Arquiteto -> Dev -> QA) em uma feature pequena.

## Verification Strategy
- Validar se o próximo agente consegue retomar o trabalho apenas lendo o `handoff.md` e os arquivos SDD.
- Verificar se as métricas de sucesso (tasks concluídas) são mantidas entre sessões.
