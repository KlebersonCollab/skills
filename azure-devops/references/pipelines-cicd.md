# Pipelines & CI/CD: Azure DevOps

Azure Pipelines fornece ferramentas para automação de build, teste e deploy em qualquer linguagem ou plataforma.

## 1. Definições de Build

Listar e enfileirar builds:

`GET https://dev.azure.com/{org}/{project}/_apis/build/definitions?api-version=7.1`

### Enfileirando um novo Build (POST)
```json
{
  "definition": {
    "id": 42
  },
  "sourceBranch": "refs/heads/main",
  "parameters": "{\"debug\": \"true\"}"
}
```

## 2. Monitorando Logs e Status

Verificar o status de uma execução de build:

`GET https://dev.azure.com/{org}/{project}/_apis/build/builds/{buildId}?api-version=7.1`

**Status comuns**: `queued`, `inProgress`, `completed`, `failed`, `succeeded`.

## 3. Releases e Ambientes

Gerenciar o fluxo de entrega em diferentes estágios (Dev, Staging, Prod):

`GET https://dev.azure.com/{org}/{project}/_apis/release/definitions?api-version=7.1`

A criação de uma nova release dispara o fluxo de deploy definido na release definition.

## 4. Variáveis e Conexões (Governance)

- **Variable Groups**: Agrupam segredos e configurações compartilhadas entre pipelines.
- **Service Connections**: Conectam o AzDO com serviços externos (Azure, AWS, GCP, Docker Hub, Kubernetes).

Você pode listar conexões de serviço:

`GET https://dev.azure.com/{org}/{project}/_apis/serviceendpoint/endpoints?api-version=7.1`

## 5. YAML Pipelines (Modern CI)

O padrão em 2026 é definir pipelines via arquivos YAML (`azure-pipelines.yml`) versionados no repositório.

Exemplo de estrutura:
```yaml
trigger:
  - main

pool:
  vmImage: 'ubuntu-latest'

steps:
  - script: echo "Hello World"
    displayName: 'Exibir mensagem'
```

## 6. Boas Práticas

- **Templating**: Utilize templates YAML para reutilizar etapas de build e deploy em diferentes pipelines.
- **Secrets Management**: Nunca coloque segredos diretamente no YAML. Use Variable Groups integrados ao **Azure Key Vault**.
- **Agent Pools**: Monitore a utilização dos agentes para evitar filas excessivas.
- **Gates & Approvals**: Utilize aprovações manuais e gates de verificação (ex: consulta via API) antes de promover deploys para produção.
