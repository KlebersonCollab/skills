---
name: scaffolding-expert
version: 1.0.0
description: "Gerador dinâmico de projetos. Utiliza CLI tools (uvx copier) para inicializar boilerplates e arquiteturas complexas instantaneamente."
category: automation
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
# Scaffolding Expert

> "Não escreva boilerplate lendo texto, gere a arquitetura executando templates."

---

## Goal

Substituir o processo lento, custoso e propenso a erros de codificação manual de arquivos boilerplate (como `pyproject.toml`, Dockerfiles, `Makefile` base, roteadores padrões) por geração instantânea através de motores de templates executáveis, especificamente via **Copier**.

---

## Como Funciona

A ferramenta principal escolhida para a stack é o `copier`, executado zero-install via UV.

### Comando Padrão de Scaffold:
```bash
uvx copier copy gh:KlebersonCollab/skills/scaffolding/<nome-do-template> <destino>
```
*Nota: Caso exista um diretório local com templates no Hub, utilize o caminho local em vez da URL do GitHub.*

---

## Workflow para Novas Features/Projetos

1. **Verify**: Identifique qual stack o usuário deseja (ex: Python/FastAPI, Python/Django).
2. **Execute**: Antes de escrever qualquer arquivo de código `.py` manualmente, execute o comando de scaffolding no terminal.
3. **Configure**: O `copier` (se configurado com flags non-interactive ou respondido via STDIN) irá perguntar o nome do projeto. Se necessário, rode com flags explícitas ou crie um arquivo `.copier-answers.yml`.
4. **Build**: Após a geração do diretório, navegue para ele e rode `uv sync` para validar a instalação.

---

## Output Structure

- Diretório gerado contendo todo o boilerplate da stack solicitada (ex: `pyproject.toml`, Dockerfile).

---

## Quality Rules

- **Prioridade Absoluta**: Se existir um template de scaffolding, ele DEVE ser utilizado em detrimento de escrever código manual.
- **Customização**: Você pode (e deve) alterar o código *após* o scaffolding, mas a estrutura raiz sempre vem do template.

---

## Prohibited

- NUNCA escrever `pyproject.toml` na mão se houver um template disponível.
- NUNCA instalar frameworks de scaffolding globalmente; sempre use ferramentas efêmeras como `uvx`.

