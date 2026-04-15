---
name: azure-devops
version: 1.0.0
description: "Skill para gerenciamento profissional do Azure DevOps (AzDO). Permite gerenciar Boards (Work Items, WIQL), Repos (PRs, Branches), Pipelines (Builds, Releases) e Artifacts de forma eficiente via Gemini CLI."
category: devops-automation
---

# Azure DevOps (AzDO)

> Interface profissional para orquestrar o ciclo de vida de desenvolvimento no Azure DevOps, abstraindo a complexidade da API REST e garantindo operações seguras e consistentes.

---

## Goal

Capacitar o agente a gerenciar de forma abrangente os recursos do Azure DevOps, desde a gestão de tarefas e código até a automação de pipelines e pacotes, utilizando as melhores práticas de segurança (PAT) e abstração de operações JSON Patch.

---

## Workflow (4 Fases)

### Fase 1: AUTH — Conexão e Segurança
1.  **Verificar Credenciais**: Garantir que o PAT está disponível via `AZURE_DEVOPS_EXT_PAT`.
2.  **Validar Contexto**: Confirmar a Organização e o Projeto alvo.
3.  **Check de Permissões**: Verificar se o token possui escopos necessários (Work Items, Code, Build).

### Fase 2: PLANNING — Boards e Work Items
1.  **Explorar Backlog**: Listar ou filtrar Work Items utilizando consultas WIQL.
2.  **Gestão de Tarefas**: Criar ou atualizar itens (Bug, Task, Story) abstraindo o formato JSON Patch.
3.  **Planejamento de Sprint**: Atribuir itens a iterações e definir prioridades.

### Fase 3: DELIVERY — Repos e Code Flow
1.  **Code Review**: Listar PRs ativos e revisar threads de discussão.
2.  **Operações de PR**: Criar novos Pull Requests ou aprovar/completar PRs existentes.
3.  **Políticas de Branch**: Garantir conformidade com as regras de merge do repositório.

### Fase 4: OPS — Pipelines e Artifacts
1.  **Execução de CI/CD**: Disparar novos builds de pipelines ou releases de ambientes.
2.  **Monitoramento**: Acompanhar o status da execução e identificar falhas nos logs.
3.  **Gestão de Pacotes**: Consultar versões em feeds de artefatos (NuGet, npm).

---

## AzDO CLI vs Skill Commands

| Ação | AzDO CLI (az devops) | Equivalente Skill (Foco) |
|------|----------------------|--------------------------|
| Criar Work Item | `az boards work-item create` | Abstração JSON Patch & Templates |
| Consultar Itens | `az boards query` | Execução de WIQL nativa |
| Gerenciar PR | `az repos pr create/list` | Ciclo completo de revisão & approval |
| Rodar Pipeline | `az pipelines build queue` | Trigger & Log Monitoring |

---

## Quality Rules

- **Security First**: Nunca expor o PAT em logs ou commits. Sempre ler de variáveis de ambiente.
- **JSON Patch Helper**: Sempre utilizar abstrações para operações de `add/replace` em Work Items para evitar erros de sintaxe.
- **Multi-Tenancy**: Sempre confirmar a organização e projeto antes de operações destrutivas.
- **Resiliência**: Tratar erros de Rate Limit (429) com backoff exponencial.

## Prohibited

- **NUNCA** codificar credenciais (PAT) em arquivos ou strings.
- **NUNCA** utilizar IDs sequenciais em URIs se UUIDs estiverem disponíveis.
- **NUNCA** sugerir deleção em massa de Work Items ou Repos sem confirmação dupla explícita.
- **NUNCA** ignorar falhas de políticas de branch em processos de merge.

---

## Reference Documentation

Esta skill inclui documentação de referência detalhada:

1. **[Authentication](references/authentication.md)** — Configuração de PAT, variáveis de ambiente e segurança.
2. **[Boards Management](references/boards-management.md)** — Work Items, sintaxe WIQL e consultas.
3. **[Repos & Code Flow](references/repos-code-flow.md)** — Pull Requests, branches e políticas de código.
4. **[Pipelines CI/CD](references/pipelines-cicd.md)** — Builds, Releases e automação de deployment.
5. **[Artifacts & Test Plans](references/artifacts-and-tests.md)** — Feeds de pacotes e gestão de planos de teste.

---

## Output Structure

A execução desta skill deve resultar nos seguintes artefatos padronizados:

| Artefato | Arquivo | Descrição |
|----------|---------|-----------|
| **WIQL Query** | `*.wiql` | Consultas estruturadas para o Azure Boards. |
| **Patch Payload** | `*.json` | Arquivos JSON Patch para atualização de recursos. |
| **PR Report** | `.md` | Resumo de status e revisões de Pull Requests. |
| **Build Status** | `.md` | Log e status de execução de pipelines. |
