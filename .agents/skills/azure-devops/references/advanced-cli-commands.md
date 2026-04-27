# Advanced CLI Commands: Azure DevOps

This guide covers utility and low-level commands for advanced automation in Azure DevOps.

---

## 1. Context Management (`configure`)

Avoid passing `--org` and `--project` in every command by fixing the defaults for the current session.

### Configure Defaults
`az devops configure --defaults organization=https://dev.azure.com/MyOrg project=MyProject`

### Clear Defaults
`az devops configure --defaults organization='' project=''`

## 2. Direct API Invocation (`invoke`)

When there is no CLI command for a specific functionality, you can invoke the REST API directly while maintaining the PAT authentication context.

### Base Syntax
`az devops invoke --area {area} --resource {resource} --route-values {key=val} --query-parameters {key=val}`

### Example: List Audit Definitions
```bash
az devops invoke \
    --area "audit" \
    --resource "auditlog" \
    --route-values project="MyProject" \
    --http-method GET
```

## 3. Wiki Management

Manage technical documentation integrated into the repository or as a project Wiki.

### List Wikis
`az devops wiki list`

### Create a Wiki
`az devops wiki create --name "Technical-Manual" --type projectwiki`

### Manage Pages
- `az devops wiki page create`: Add new content.
- `az devops wiki page update`: Update existing documentation.

## 4. Agent Pools & Queues

Manage the runner infrastructure for pipelines.

### List Agent Pools
`az pipelines pool list`

### List Execution Queues
`az pipelines queue list`

## 5. Automation Tip (JSON Output)

All commands support the `--output json` flag. Combine with `jq` to extract IDs and values for complex scripts:

```bash
# Example: Get the ID of the first active PR
PR_ID=$(az repos pr list --status active --query "[0].pullRequestId" -o tsv)
echo "Processing PR #$PR_ID"
```
