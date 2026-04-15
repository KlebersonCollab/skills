# Authentication & Setup: Azure DevOps

O Azure DevOps utiliza o **Personal Access Token (PAT)** como o principal mecanismo de autenticação para sua API REST e CLI. Este guia descreve como configurar e gerenciar sua conexão de forma segura.

## 1. Gerando um PAT

1.  Acesse seu Azure DevOps (`https://dev.azure.com/{sua-organizacao}`).
2.  No canto superior direito, clique em **User settings** > **Personal access tokens**.
3.  Clique em **New Token**.
4.  Defina um nome (ex: `Gemini-CLI`) e a expiração.
5.  **Escopos (Scopes)**: Selecione os escopos mínimos necessários:
    - **Build**: Read & Execute
    - **Code**: Read, Write, & Manage
    - **Release**: Read, Write, & Execute
    - **Test Management**: Read & Write
    - **Work Items**: Read, Write, & Manage

## 2. Configurando Variáveis de Ambiente

Para segurança, nunca armazene o PAT em arquivos de código. Utilize variáveis de ambiente no seu shell ou arquivo `.env`.

```bash
# Adicione ao seu ~/.zshrc ou ~/.bashrc
export AZURE_DEVOPS_EXT_PAT="seu-pat-aqui"
export AZURE_DEVOPS_ORG="https://dev.azure.com/sua-organizacao"
export AZURE_DEVOPS_PROJECT="seu-projeto"
```

## 3. Padrão de Cabeçalho HTTP

Se você estiver realizando chamadas manuais via cURL ou scripts, utilize o PAT codificado em Base64 no cabeçalho `Authorization`.

```bash
# O formato é ":{PAT}" (o colon antes do PAT é obrigatório)
TOKEN=$(echo -n ":$AZURE_DEVOPS_EXT_PAT" | base64)

curl -H "Authorization: Basic $TOKEN" \
     -H "Content-Type: application/json" \
     "https://dev.azure.com/{org}/{projeto}/_apis/projects?api-version=7.1"
```

## 4. Multi-Tenancy (Várias Orgs/Projetos)

Se você trabalha em múltiplas organizações, pode configurar perfis ou utilizar o comando de configuração do AzDO CLI:

```bash
# Configurar default via CLI oficial
az devops configure --defaults organization=https://dev.azure.com/org1 project=proj1
```

## 5. Melhores Práticas de Segurança

- **Rotação**: Defina uma expiração curta (ex: 30 a 90 dias) e rotacione o token periodicamente.
- **Least Privilege**: Conceda apenas os escopos necessários para a tarefa atual.
- **Secret Management**: Utilize ferramentas como AWS Secrets Manager, Azure Key Vault ou 1Password CLI para injetar o PAT dinamicamente.
