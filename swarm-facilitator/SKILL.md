---
name: swarm-facilitator
version: 1.0.0
description: "Orquestrador de workflows Multi-Agente (Swarm). Define protocolos de handoff, personas e passagem de contexto usando a estrutura SDD."
category: orchestration
---

# Swarm Facilitator

> "Reger múltiplos agentes requer coreografia, não apenas programação. O repositório é o tabuleiro e o Markdown é a linguagem de comunicação inter-agentes."

---

## Goal

Capacitar agentes de IA a atuarem em um fluxo coreografado (Swarm/Crew), assumindo personas especializadas (ex: Arquiteto, Developer, Revisor) e utilizando protocolos de Handoff (passagem de bastão) baseados no ecossistema SDD. Esta skill transforma a interação em um pipeline de linha de produção invisível.

---

## Personas Suportadas

O orquestrador pode instruir o LLM atual ou instruir o usuário a iniciar uma nova thread com uma destas personas:

1. **Solution Architect (Arquiteto)**: Foco em Fases 1 e 2 do SDD. Usa as skills `architecture`, `api-architect` e `brainstorming`. Produz PRDs e ADRs.
2. **Lead Developer (Engenheiro)**: Foco na Fase 3 do SDD. Consome as especificações do Arquiteto. Usa skills de stack (`fastapi-expert`, `flutter-fvm`). Escreve o código produtivo.
3. **Quality Assurance / SRE (Revisor)**: Foco na Fase 4 do SDD. Usa `clean-code-mentor` e `observability-expert`. Faz code review e audita a telemetria.

---

## Workflow (Protocolo de Handoff)

Para orquestrar múltiplos agentes no mesmo projeto local sem perda de contexto, execute este protocolo:

### 1. State Definition (O Estado é Rei)
O Arquiteto deve atualizar o `STATE.md` com o plano de ação macro e salvar as tarefas em `tasks.md`.

### 2. Context Dump (Salvar o Jogo)
Antes de finalizar a sessão do Agente 1 (Arquiteto), ele **deve** compilar as decisões recentes em um arquivo de handover na pasta `.specs/features/<feature>/handoff.md`.
*Exemplo de formato de handoff:*
- O que foi decidido.
- Onde está a especificação (caminho absoluto).
- Qual skill o próximo agente DEVE invocar ao acordar.

### 3. Handoff Message (A Passagem de Bastão)
A última mensagem da sessão do Agente 1 deve ser um prompt formatado para o usuário copiar e colar na nova sessão do Agente 2.
**Template**: 
> "Inicie o desenvolvimento da Feature X. Sou o Arquiteto e deixei as specs em `.specs/...`. Invoque a skill `fastapi-expert` e siga os passos 1 e 2."

---

## Output Structure

- `handoff.md`: Documento contendo as diretrizes para a próxima persona.
- `STATE.md`: Atualizado via Harness Expert indicando quem está com a bola.

---

## Quality Rules

- O Handoff deve ser enxuto.
- A comunicação é feita exclusivamente através do Markdown salvo no repositório.

---

## Prohibited

- NUNCA tentar fazer um único agente assumir o papel de Arquiteto e Desenvolvedor simultaneamente em épicos complexos. Dividir para conquistar.
- NUNCA passar o bastão sem atualizar o `STATE.md` via `harness-expert` ou script.

