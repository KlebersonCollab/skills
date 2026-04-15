# Advanced CLI Commands: Azure DevOps

Este guia cobre comandos de utilidade e baixo nível para automação avançada no Azure DevOps.

---

## 1. Gestão de Contexto (`configure`)

Evite passar `--org` e `--project` em todos os comandos fixando os defaults para a sessão atual.

### Configurar Defaults
`az devops configure --defaults organization=https://dev.azure.com/MinhaOrg project=MeuProjeto`

### Limpar Defaults
`az devops configure --defaults organization='' project=''`

## 2. Invocação Direta de API (`invoke`)

Quando não existe um comando CLI para uma funcionalidade específica, você pode invocar a API REST diretamente mantendo o contexto de autenticação do PAT.

### Sintaxe Base
`az devops invoke --area {area} --resource {resource} --route-values {key=val} --query-parameters {key=val}`

### Exemplo: Listar Definições de Auditoria
```bash
az devops invoke \
    --area "audit" \
    --resource "auditlog" \
    --route-values project="MeuProjeto" \
    --http-method GET
```

## 3. Wiki Management

Gerenciar documentação técnica integrada ao repositório ou como Wiki de projeto.

### Listar Wikis
`az devops wiki list`

### Criar uma Wiki
`az devops wiki create --name "Manual-Tecnico" --type projectwiki`

### Gerenciar Páginas
- `az devops wiki page create`: Adicionar novo conteúdo.
- `az devops wiki page update`: Atualizar documentação existente.

## 4. Agent Pools & Queues

Gerenciar a infraestrutura de runners para pipelines.

### Listar Pools de Agentes
`az pipelines pool list`

### Listar Filas de Execução
`az pipelines queue list`

## 5. Dica de Automação (JSON Output)

Todos os comandos suportam a flag `--output json`. Combine com `jq` para extrair IDs e valores para scripts complexos:

```bash
# Exemplo: Pegar o ID do primeiro PR aberto
PR_ID=$(az repos pr list --status active --query "[0].pullRequestId" -o tsv)
echo "Processando PR #$PR_ID"
```
