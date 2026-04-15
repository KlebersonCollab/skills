# Onboarding Checklists: Action Plans

Checklists padronizados para garantir que nenhum passo importante seja esquecido durante a entrada de membros ou início de novos projetos.

---

## 1. Novo Membro do Time (Skynet Developer)

### A. Acessos e Ambiente
- [ ] Clonar o repositório central de skills.
- [ ] Instalar o Gemini CLI.
- [ ] Configurar variáveis de ambiente (ex: `AZURE_DEVOPS_EXT_PAT`).
- [ ] Setup de ambiente Python via **UV** (`python-uv`).
- [ ] Setup de ambiente Flutter via **FVM** (`flutter-fvm`).

### B. Leituras Obrigatórias (Fonte da Verdade)
- [ ] `README.md` (Visão Geral).
- [ ] `sdd/SKILL.md` (Metodologia de Trabalho).
- [ ] `.specs/codebase/STACK.md` (O que usamos).
- [ ] `.specs/project/PROJECT.md` (Para onde vamos).

### C. Primeiro Commit
- [ ] Escolher uma tarefa "Small" ou "Bug fix" no Roadmap.
- [ ] Executar o workflow SDD completo (Spec -> Plan -> Tasks -> Impl -> Review).

---

## 2. Início de Nova Feature (Módulo ou Skill)

### A. Planejamento (Discovery & Specify)
- [ ] Criar pasta em `.specs/features/`.
- [ ] Auditar PRD ou Requisitos iniciais.
- [ ] Escrever `spec.md` com BDD.
- [ ] Criar `plan.md` com Diagrama Mermaid.
- [ ] Gerar lista de tarefas atômicas em `tasks.md`.

### B. Arquitetura
- [ ] Avaliar se a mudança exige um ADR (`architecture`).
- [ ] Revisar princípios SOLID e DRY (`clean-code-mentor`).

### C. Execução e Fechamento
- [ ] Implementar com testes (TDD).
- [ ] Auditar via `skill-factory-validator` (se for uma skill).
- [ ] Atualizar proativamente o ROADMAP e os logs de memória.

---

## 3. Dica do Navigator
"O checklist não é um peso, é sua rede de segurança. Siga os passos e garanta que sua entrega seja inquestionável."
