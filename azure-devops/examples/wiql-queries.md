# Examples: WIQL Queries (Azure Boards)

Structured **WIQL** query templates for programmatic use in Azure Boards.

## 1. Active Bugs (Current Sprint)

Returns all bugs that are not closed or resolved in the project's current iteration.

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

## 2. My Tasks (Backlog)

Lists all tasks assigned to the authenticated user that haven't started yet.

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

## 3. User Stories by Tag

Filters User Stories containing a specific tag (e.g., 'Refactoring').

```sql
SELECT
    [System.Id],
    [System.Title],
    [System.Tags]
FROM workitems
WHERE
    [System.TeamProject] = @project
    AND [System.WorkItemType] = 'User Story'
    AND [System.Tags] CONTAINS 'Refactoring'
```

## 4. Items Changed in the Last 24h

Monitor recent changes in the project.

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

## 5. Unassigned Work Items

Ideal for backlog grooming.

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
