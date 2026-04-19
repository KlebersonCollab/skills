# ADR 003: Multi-Agent Orchestration Protocol

## Status
Accepted

## Context
Grandes épicos de desenvolvimento causam saturação de contexto e degradação na qualidade dos agentes se mantidos em uma única sessão. Precisamos de um protocolo para dividir o trabalho entre diferentes "personas" de agentes especializados.

## Decision
Adotaremos o protocolo **SDD-Swarm (Handoff via Filesystem)**.

## Rationale
- **Agnóstico de Plataforma**: Funciona em qualquer agente (Gemini, Claude, GPT) que tenha acesso ao filesystem.
- **Estado Persistente**: O framework SDD (`plan.md`, `tasks.md`, `STATE.md`) serve como a única fonte de verdade.
- **Handoff Explícito**: Introduzimos o artefato `handoff.md` para instruções de passagem de turno.

## Consequences
- Épicos complexos serão divididos em tarefas para diferentes personas (Arquiteto -> Dev -> QA).
- O agente atual deve criar o `handoff.md` e encerrar a sessão quando sua sub-tarefa estiver concluída.
