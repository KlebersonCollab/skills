# Pipelines & CI/CD: Azure DevOps

Azure Pipelines provides tools for build, test, and deployment automation in any language or platform.

## 1. Build Definitions

List and enqueue builds:

`GET https://dev.azure.com/{org}/{project}/_apis/build/definitions?api-version=7.1`

### Enqueueing a new Build (POST)
```json
{
  "definition": {
    "id": 42
  },
  "sourceBranch": "refs/heads/main",
  "parameters": "{\"debug\": \"true\"}"
}
```

## 2. Monitoring Logs and Status

Check the status of a build execution:

`GET https://dev.azure.com/{org}/{project}/_apis/build/builds/{buildId}?api-version=7.1`

**Common Statuses**: `queued`, `inProgress`, `completed`, `failed`, `succeeded`.

## 3. Releases and Environments

Manage the delivery flow across different stages (Dev, Staging, Prod):

`GET https://dev.azure.com/{org}/{project}/_apis/release/definitions?api-version=7.1`

Creating a new release triggers the deployment flow defined in the release definition.

## 4. Variables and Connections (Governance)

- **Variable Groups**: Group secrets and configurations shared among pipelines.
- **Service Connections**: Connect AzDO with external services (Azure, AWS, GCP, Docker Hub, Kubernetes).

You can list service connections:

`GET https://dev.azure.com/{org}/{project}/_apis/serviceendpoint/endpoints?api-version=7.1`

## 5. YAML Pipelines (Modern CI)

The standard in 2026 is to define pipelines via YAML files (`azure-pipelines.yml`) versioned in the repository.

Example structure:
```yaml
trigger:
  - main

pool:
  vmImage: 'ubuntu-latest'

steps:
  - script: echo "Hello World"
    displayName: 'Display message'
```

## 6. Best Practices

- **Templating**: Use YAML templates to reuse build and deployment steps across different pipelines.
- **Secrets Management**: Never put secrets directly in the YAML. Use Variable Groups integrated with **Azure Key Vault**.
- **Agent Pools**: Monitor agent utilization to avoid excessive queues.
- **Gates & Approvals**: Use manual approvals and verification gates (e.g., API query) before promoting deployments to production.
