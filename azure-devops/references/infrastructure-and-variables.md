# Infrastructure & Variables: Azure DevOps Governance

This guide covers the automation and governance of infrastructure resources in Azure DevOps using the official CLI.

---

## 1. Variable Groups

Variable groups allow centralizing configurations and secrets shared across multiple pipelines.

### List Variable Groups
`az pipelines variable-group list --project {project}`

### Create a Variable Group (CLI)
```bash
az pipelines variable-group create \
    --name "Shared-Secrets" \
    --project "MyProject" \
    --variables "DB_HOST=10.0.0.1" "DB_PORT=5432" \
    --authorize true
```

### Azure Key Vault Integration
Recommended for sensitive secrets. You can create a group that maps secrets directly from a Key Vault:
```bash
az pipelines variable-group create \
    --name "Vault-Secrets" \
    --vault-name "MyKeyVault" \
    --authorize true
```

## 2. Service Connections (Endpoints)

Service connections allow AzDO to interact with external resources (Azure, Docker, GitHub, etc.).

### List Endpoints
`az devops service-endpoint list --project {project}`

### Create an Azure Connection (Service Principal)
Usually done via JSON template for greater control:
```bash
az devops service-endpoint create --service-endpoint-configuration template.json
```

## 3. Resource Authorization

By default, new variable groups or endpoints may be blocked for pipelines. Always check the `--authorize true` flag or use:

`az pipelines variable-group variable update --group-id {id} --name {var} --value {val}`

## 4. Individual Pipeline Variables

For variables specific to a single execution or pipeline definition:

```bash
# Add variable to an existing pipeline
az pipelines variable create \
    --name "ENV_NAME" \
    --pipeline-id 42 \
    --value "Production" \
    --secret true
```

## 5. Best Practices

- **Naming**: Use consistent prefixes (e.g., `GLOBAL-`, `APP1-`, `INFRA-`).
- **Secrets**: Never use simple variables for passwords; use the `--secret true` flag or Variable Groups linked to Key Vault.
- **Least Privilege**: Authorize resources only for the pipelines that actually need them.
- **Audit**: Periodically list endpoints (`az devops service-endpoint list`) to identify obsolete connections.
