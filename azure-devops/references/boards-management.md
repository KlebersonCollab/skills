# Boards Management: Azure DevOps

A gestão de trabalho no Azure DevOps (Azure Boards) gira em torno de **Work Items** e consultas estruturadas via **WIQL**.

## 1. Tipos de Work Items

Os tipos padrão dependem do processo (Agile, Scrum, Basic):
- **User Story / PBI**: Requisito de negócio.
- **Bug**: Erro ou falha.
- **Task**: Trabalho técnico decomposto.
- **Epic / Feature**: Níveis superiores de agrupamento.

## 2. Consultas via WIQL (Work Item Query Language)

WIQL é uma linguagem semelhante ao SQL para filtrar Work Items.

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

**Macros úteis**: `@me`, `@project`, `@today`.

## 3. Criando Work Items (JSON Patch)

O Azure DevOps exige o formato `application/json-patch+json` para criar ou atualizar itens.

### Exemplo de Payload (POST)
```json
[
  {
    "op": "add",
    "path": "/fields/System.Title",
    "from": null,
    "value": "Implementar autenticação OAuth2"
  },
  {
    "op": "add",
    "path": "/fields/System.Description",
    "from": null,
    "value": "Necessário para a integração com o novo portal."
  },
  {
    "op": "add",
    "path": "/fields/System.AssignedTo",
    "from": null,
    "value": "usuario@exemplo.com"
  }
]
```

## 4. Atualizando Itens (PATCH)

Para atualizar, use o ID do item e envie as operações desejadas (`add`, `replace`, `remove`).

```json
[
  {
    "op": "replace",
    "path": "/fields/System.State",
    "value": "Active"
  }
]
```

## 5. Gerenciando Iterações (Sprints)

Você pode listar iterações de um time para planejar o trabalho:

`GET https://dev.azure.com/{org}/{project}/{team}/_apis/work/teamsettings/iterations?api-version=7.1`

## 6. Boas Práticas

- **Abstração**: Use funções helper que montem o array JSON Patch para você, reduzindo erros de sintaxe.
- **WIQL**: Teste suas queries no editor visual do Azure Boards antes de implementá-las na CLI.
- **Fields**: Fique atento aos campos customizados (ex: `Custom.Severity`). Eles variam de acordo com o processo do projeto.
