# Spec: Governance Hardening (Git & Testing Compliance)

## 1. Visão Geral
Esta especificação visa eliminar a discricionariedade do agente no cumprimento dos workflows de Git e Testes dentro do framework SDD. O objetivo é tornar mandatório o ciclo `Task -> Test -> Commit` antes de qualquer marcação de progresso no `tasks.md`.

## 2. Requisitos de Governança (GR)

- **GR-1**: Inserir mandato proibitivo na Fase 3 do SDD sobre a atomicidade da tarefa.
- **GR-2**: Vincular a conclusão de tarefas obrigatoriamente à aprovação em testes (100%) e ao commit Git.
- **GR-3**: Padronizar o uso de sensores de conformidade Git nos `contract.md` das features.
- **GR-4**: Reforçar o "Exit Gate" no Onboarding Navigator para auditoria de commits e branches.
- **GR-5**: Tornar mandatório o uso de feature branches (`feat/`, `fix/`, etc.) para qualquer desenvolvimento, protegendo a `main`.

## 3. Critérios de Aceitação (AC) - BDD

### Cenário 1: Marcação de Tarefa Concluída
**Given** que o agente está na Fase 3 (Implementação) do SDD
**When** ele considerar que uma tarefa foi finalizada no código
**Then** ele NUNCA deve marcar o checkbox no `tasks.md` sem antes:
    1. Executar a suíte de testes e obter 100% de sucesso.
    2. Realizar o commit das mudanças seguindo a `git-workflow`.
**And** o commit deve obrigatoriamente mencionar o ID da task.

### Cenário 2: Validação de Contrato
**Given** a criação de um novo `contract.md`
**When** o agente definir os sensores de validação
**Then** ele deve incluir obrigatoriamente um sensor de **Git Compliance** para verificar a atomicidade dos commits.

## 4. Restrições Técnicas
- As alterações devem ser aplicadas em `sdd/SKILL.md` e `onboarding-navigator/SKILL.md`.
- O idioma deve ser Português Brasileiro (PT-BR).
- A numeração de versões das skills afetadas deve ser incrementada.
