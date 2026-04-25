# PRD: Multi-Agent Orchestrator

## 1. Contexto e Visão
A medida que projetos complexos evoluem (como o SecondBrain), delegar toda a carga cognitiva para um único prompt/sessão de agente torna-se insustentável. O *Multi-Agent Orchestrator* será uma nova skill que atua como um "Gerente de Projetos" ou "Coreógrafo", definindo como diferentes personas de IA interagem em sequência (Swarm/Crew patterns) utilizando os artefatos SDD como meio de comunicação.

## 2. Escopo (In Scope)
- Criação da skill `multi-agent-orchestrator`.
- Definição de papéis (ex: Arquiteto de Solução, Lead Developer, Revisor/QA).
- Protocolos de Handoff (passagem de bastão) baseados no estado do repositório (`STATE.md` e `.specs/`).

## 3. Requisitos Técnicos
1. **Workflow de Hand-off**: Instruções claras de como um agente "salva o jogo" e deixa instruções explícitas (via `tasks.md` ou comentários em PRs) para o próximo agente.
2. **Integração com SDD**: O orquestrador usará as fases do SDD. (Ex: Arquiteto faz Fase 1 e 2; Developer faz Fase 3; QA faz Fase 4).
3. **Governança**: O orquestrador não escreve código; ele lê o Roadmap e aloca/instrui a persona correta para a próxima tarefa.

## 4. Próximos Passos (Implementação)
- [ ] Criar o diretório e `SKILL.md` inicial.
- [ ] Documentar o fluxo de "Swarm" baseado no sistema de arquivos local.
- [ ] Validar a skill no `Makefile` e distribuí-la.
