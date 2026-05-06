# Boards Management: Azure DevOps

Work management in Azure DevOps (Azure Boards) revolves around **Work Items** and structured queries via **WIQL**.

## 1. Work Item Types

Standard types depend on the process (Agile, Scrum, Basic):
- **User Story / PBI**: Business requirement.
- **Bug**: Error or failure.
- **Task**: Decomposed technical work.
- **Epic / Feature**: Higher-level grouping levels.

## 2. Queries via WIQL (Work Item Query Language)

WIQL is a SQL-like language for filtering Work Items.

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
    AND [System.State] <> 'Closed'
ORDER BY [System.ChangedDate] DESC
```

**Useful Macros**: `@me`, `@project`, `@today`.

## 3. Creating Work Items (JSON Patch)

Azure DevOps requires the `application/json-patch+json` format to create or update items.

### Example Payload (POST)
```json
[
  {
    "op": "add",
    "path": "/fields/System.Title",
    "from": null,
    "value": "Implement OAuth2 authentication"
  },
  {
    "op": "add",
    "path": "/fields/System.Description",
    "from": null,
    "value": "Required for integration with the new portal."
  },
  {
    "op": "add",
    "path": "/fields/System.AssignedTo",
    "from": null,
    "value": "user@example.com"
  }
]
```

## 4. Updating Items (PATCH)

To update, use the item ID and send the desired operations (`add`, `replace`, `remove`).

```json
[
  {
    "op": "replace",
    "path": "/fields/System.State",
    "value": "Active"
  }
]
```

## 5. Managing Iterations (Sprints)

You can list a team's iterations to plan work:

`GET https://dev.azure.com/{org}/{project}/{team}/_apis/work/teamsettings/iterations?api-version=7.1`

## 6. Best Practices

- **Abstraction**: Use helper functions that assemble the JSON Patch array for you, reducing syntax errors.
- **WIQL**: Test your queries in the Azure Boards visual editor before implementing them in the CLI.
- **Fields**: Be aware of custom fields (e.g., `Custom.Severity`). They vary according to the project process.
