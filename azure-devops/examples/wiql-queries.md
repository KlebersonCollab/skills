# Examples: WIQL Queries (Azure Boards)

Modelos de consultas estruturadas em **WIQL** para uso programático no Azure Boards.

## 1. Bugs Ativos (Sprint Atual)

Retorna todos os bugs que não estão fechados ou resolvidos na iteração atual do projeto.

```sql
SELECT
    [System.Id],
    [System.Title],
    [System.State],
    [System.AssignedTo]
FROM workitems
WHERE
    [System.TeamProject] = @project
    AND [System.WorkItemType] = 'Bug'
    AND [System.State] NOT IN ('Closed', 'Resolved')
    AND [System.IterationPath] = @currentIteration
ORDER BY [Microsoft.VSTS.Common.Priority] ASC, [System.CreatedDate] DESC
```

## 2. Minhas Tarefas (Backlog)

Lista todas as tarefas (Tasks) atribuídas ao usuário autenticado que ainda não foram iniciadas.

```sql
SELECT
    [System.Id],
    [System.Title],
    [System.State]
FROM workitems
WHERE
    [System.TeamProject] = @project
    AND [System.WorkItemType] = 'Task'
    AND [System.AssignedTo] = @me
    AND [System.State] = 'To Do'
ORDER BY [System.Id] ASC
```

## 3. Histórias de Usuário por Tag

Filtra User Stories que contenham uma tag específica (ex: 'Refatoração').

```sql
SELECT
    [System.Id],
    [System.Title],
    [System.Tags]
FROM workitems
WHERE
    [System.TeamProject] = @project
    AND [System.WorkItemType] = 'User Story'
    AND [System.Tags] CONTAINS 'Refatoração'
```

## 4. Itens Alterados nas Últimas 24h

Monitorar mudanças recentes no projeto.

```sql
SELECT
    [System.Id],
    [System.Title],
    [System.ChangedBy],
    [System.ChangedDate]
FROM workitems
WHERE
    [System.TeamProject] = @project
    AND [System.ChangedDate] >= @today - 1
ORDER BY [System.ChangedDate] DESC
```

## 5. Work Items Sem Atribuição

Ideal para triagem de backlog.

```sql
SELECT
    [System.Id],
    [System.Title],
    [System.WorkItemType]
FROM workitems
WHERE
    [System.TeamProject] = @project
    AND [System.AssignedTo] = ''
    AND [System.State] NOT IN ('Closed', 'Removed')
```
