# Project Persistent Memory

## Convenções de Skills
- **Marker**: Uma pasta é considerada uma skill se contiver um arquivo `SKILL.md` na sua raiz.
- **Estrutura Padrão**: Uma skill deve conter `CHANGELOG.md`, `README.md`, `SKILL.md` e opcionalmente as pastas `references/` e `resources/`.
- **Mandatory Hook**: Skills de execução DEVEM conter a seção `🔒 Prerequisites (Mandatory)` vinculando-as ao SDD.

## Mapeamento de Agentes
- **Claude**: Pasta `.claude/`, arquivo mestre `CLAUDE.md`.
- **Gemini**: Pasta `.gemini/`, arquivo mestre `GEMINI.md`.
- **Agent**: Pasta `.agent/`, arquivo mestre `AGENT.md`.

## Pipeline de Distribuição
- Os artefatos são gerados dinamicamente coletando as skills da raiz.
- A estrutura final nos ZIPs preserva o nome da pasta do agente (ex: `.claude/`) para facilitar a instalação direta.

## Governança Operacional
- **Uso Mandatório de Skills**: Conforme `GLOBAL_MANDATES.md`, qualquer execução técnica **DEVE** ser mediada por uma skill específica do Hub.
- **Deterministic Lifecycle**: Todo agente operando no Hub deve obrigatoriamente executar o **Session Bootstrap** (início) e o **Session Exit Gate** (fim).

