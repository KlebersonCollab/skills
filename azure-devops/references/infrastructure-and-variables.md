# Infrastructure & Variables: Azure DevOps Governance

Este guia cobre a automação e governança de recursos de infraestrutura no Azure DevOps, utilizando a CLI oficial.

---

## 1. Variable Groups (Grupos de Variáveis)

Variable groups permitem centralizar configurações e segredos compartilhados entre múltiplos pipelines.

### Listar Variable Groups
`az pipelines variable-group list --project {projeto}`

### Criar um Variable Group (CLI)
```bash
az pipelines variable-group create \
    --name "Shared-Secrets" \
    --project "MeuProjeto" \
    --variables "DB_HOST=10.0.0.1" "DB_PORT=5432" \
    --authorize true
```

### Integração com Azure Key Vault
Recomendado para segredos sensíveis. Você pode criar um grupo que mapeia segredos diretamente de um Key Vault:
```bash
az pipelines variable-group create \
    --name "Vault-Secrets" \
    --vault-name "MeuKeyVault" \
    --authorize true
```

## 2. Service Connections (Endpoints)

Conexões de serviço permitem que o AzDO interaja com recursos externos (Azure, Docker, GitHub, etc.).

### Listar Endpoints
`az devops service-endpoint list --project {projeto}`

### Criar uma Conexão com Azure (Service Principal)
Geralmente feito via JSON template para maior controle:
```bash
az devops service-endpoint create --service-endpoint-configuration template.json
```

## 3. Autorização de Recursos

Por padrão, novos grupos de variáveis ou endpoints podem estar bloqueados para pipelines. Sempre verifique a flag `--authorize true` ou utilize:

`az pipelines variable-group variable update --group-id {id} --name {var} --value {val}`

## 4. Variáveis de Pipeline Individuais

Para variáveis específicas de uma única execução ou definição de pipeline:

```bash
# Adicionar variável a um pipeline existente
az pipelines variable create \
    --name "ENV_NAME" \
    --pipeline-id 42 \
    --value "Production" \
    --secret true
```

## 5. Boas Práticas

- **Nomenclatura**: Utilize prefixos consistentes (ex: `GLOBAL-`, `APP1-`, `INFRA-`).
- **Segredos**: Nunca use variáveis simples para senhas; utilize a flag `--secret true` ou Variable Groups vinculados ao Key Vault.
- **Lease Privilege**: Autorize recursos apenas para os pipelines que realmente precisam deles.
- **Audit**: Periodicamente liste os endpoints (`az devops service-endpoint list`) para identificar conexões obsoletas.
