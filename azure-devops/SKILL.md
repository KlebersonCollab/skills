---
name: azure-devops
version: 1.1.0
description: "Skill para gerenciamento profissional do Azure DevOps (AzDO). Permite gerenciar Boards, Repos, Pipelines, Artifacts, Governança (Variable Groups, Service Connections) e Administração de forma eficiente."
category: devops-automation
---

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar qualquer execução técnica:
1. **Context Check**: Você reidratou o contexto lendo `STATE.md`, `MEMORY.md` e `LEARNINGS.md`?
2. **Spec Check**: O arquivo `spec.md` existe com requisitos e Critérios de Aceitação (ACs) claros? (BDD mandatório para Medium+).
3. **Plan Check**: O arquivo `plan.md` define a arquitetura, schemas e inclui diagramas **Mermaid**?
4. **Contract Check**: O arquivo `contract.md` foi estabelecido com os sensores de validação?
5. **Task Check**: A lista de tarefas em `tasks.md` está detalhada e atomizada?

---
# Azure DevOps (AzDO)

> Interface profissional para orquestrar o ciclo de vida de desenvolvimento e a governança de plataforma no Azure DevOps.

---

## Goal

Capacitar o agente a gerenciar de forma abrangente os recursos do Azure DevOps, desde a gestão de tarefas e código até a automação de infraestrutura (Service Connections), governança de variáveis e administração de usuários/permissões, utilizando as melhores práticas de segurança e automação CLI.

---

## Workflow (5 Fases)

### Fase 1: AUTH — Conexão e Segurança
1.  **Verificar Credenciais**: Garantir que o PAT está disponível e configurado.
2.  **Configurar Contexto**: Utilizar `az devops configure --defaults` para fixar organização e projeto.

### Fase 2: INFRA — Governança e IaC
1.  **Conexões de Serviço**: Criar ou auditar Service Endpoints para integrações externas.
2.  **Variáveis Globais**: Gerenciar Variable Groups compartilhados e integrá-los a Key Vaults.
3.  **Segurança de Pipeline**: Autorizar o uso de recursos críticos (Pools, Queues, Variable Groups) para pipelines específicos.

### Fase 3: PLANNING — Boards e Work Items
1.  **Consultas WIQL**: Executar filtros complexos para triagem de backlog.
2.  **Gestão de Itens**: Criar e atualizar Work Items via abstração JSON Patch.

### Fase 4: DELIVERY — Repos e Code Flow
1.  **PR Lifecycle**: Gerenciar o ciclo completo de Pull Requests e revisões.
2.  **Políticas de Branch**: Validar e aplicar regras de proteção de branch.

### Fase 5: OPS & ADMIN — Gestão e Monitoramento
1.  **CI/CD Ops**: Disparar builds, releases e monitorar logs de execução.
2.  **User Admin**: Gerenciar usuários, times e permissões de segurança.
3.  **Extensões & Wiki**: Administrar extensões instaladas e manter a documentação do projeto via Wiki.

---

## AzDO CLI Commands Reference

| Área | Comando Principal | Propósito |
|------|-------------------|-----------|
| **Infra** | `az pipelines variable-group` | Gestão de grupos de variáveis |
| **Infra** | `az devops service-endpoint` | Gestão de conexões externas |
| **Admin** | `az devops user / team` | Administração de pessoas e times |
| **Admin** | `az devops security` | Permissões e grupos de segurança |
| **Low-Level**| `az devops invoke` | Chamadas diretas à API REST |

---

## Quality Rules

- **Security First**: Administrative operations require double confirmation of the target organization.
- **Lease Privilege**: Suggestions for variable groups and connections should always limit scope to necessary pipelines only.
- **Standard Syntax**: Use official `az devops` CLI syntax over raw REST calls unless using `invoke`.
- **Infrastructure-as-Code**: Prefer YAML-based pipelines and Variable Groups over UI-defined configurations.

## Prohibited

- **NUNCA** sugerir a criação de Service Connections com permissões de "Owner" ou "Contributor" global sem justificativa.
- **NUNCA** ignorar o comando `configure --defaults` ao iniciar uma nova sessão administrativa.
- **NUNCA** apagar Variable Groups ou Service Endpoints sem verificar dependências em pipelines ativos.

---

## Reference Documentation

Esta skill inclui documentação de referência detalhada:

1. **[Authentication](references/authentication.md)** — PAT, contextos e `configure`.
2. **[Infrastructure & Variables](references/infrastructure-and-variables.md)** — Variable Groups e Service Connections.
3. **[Administration & Security](references/administration-and-security.md)** — Users, Teams e Permissions.
4. **[Boards Management](references/boards-management.md)** — Work Items e WIQL.
5. **[Repos & Code Flow](references/repos-code-flow.md)** — Pull Requests e Políticas.
6. **[Pipelines CI/CD](references/pipelines-cicd.md)** — Execução e monitoramento.
7. **[Advanced CLI Commands](references/advanced-cli-commands.md)** — `invoke` e `wiki`.

---

## Output Structure

A execução desta skill resulta nos seguintes artefatos padronizados:

| Artefato | Formato | Descrição |
|----------|---------|-----------|
| **Infra Spec** | `.json` / `.yaml`| Definição de Variable Groups ou Endpoints. |
| **Security Audit**| `.md` | Relatório de usuários, permissões e políticas. |
| **API Invoke** | `.sh` | Scripts de chamada direta via `az devops invoke`. |
| **WIQL Query** | `.wiql` | Consultas estruturadas para o Azure Boards. |
