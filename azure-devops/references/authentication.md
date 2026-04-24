# Authentication & Setup: Azure DevOps

Azure DevOps uses the **Personal Access Token (PAT)** as the primary authentication mechanism for its REST API and CLI. This guide describes how to configure and manage your connection securely.

## 1. Generating a PAT

1.  Access your Azure DevOps (`https://dev.azure.com/{your-organization}`).
2.  In the top right corner, click **User settings** > **Personal access tokens**.
3.  Click **New Token**.
4.  Define a name (e.g., `Gemini-CLI`) and expiration.
5.  **Scopes**: Select the minimum required scopes:
    - **Build**: Read & Execute
    - **Code**: Read, Write, & Manage
    - **Release**: Read, Write, & Execute
    - **Test Management**: Read & Write
    - **Work Items**: Read, Write, & Manage

## 2. Configuring Environment Variables

For security, never store the PAT in code files. Use environment variables in your shell or a `.env` file.

```bash
# Add to your ~/.zshrc or ~/.bashrc
export AZURE_DEVOPS_EXT_PAT="your-pat-here"
export AZURE_DEVOPS_ORG="https://dev.azure.com/your-organization"
export AZURE_DEVOPS_PROJECT="your-project"
```

## 3. HTTP Header Pattern

If you are performing manual calls via cURL or scripts, use the Base64-encoded PAT in the `Authorization` header.

```bash
# The format is ":{PAT}" (the colon before the PAT is mandatory)
TOKEN=$(echo -n ":$AZURE_DEVOPS_EXT_PAT" | base64)

curl -H "Authorization: Basic $TOKEN" \
     -H "Content-Type: application/json" \
     "https://dev.azure.com/{org}/{project}/_apis/projects?api-version=7.1"
```

## 4. Multi-Tenancy (Multiple Orgs/Projects)

If you work across multiple organizations, you can configure profiles or use the AzDO CLI configuration command:

```bash
# Configure default via official CLI
az devops configure --defaults organization=https://dev.azure.com/org1 project=proj1
```

## 5. Security Best Practices

- **Rotation**: Set a short expiration (e.g., 30 to 90 days) and rotate the token periodically.
- **Least Privilege**: Grant only the scopes necessary for the current task.
- **Secret Management**: Use tools like AWS Secrets Manager, Azure Key Vault, or 1Password CLI to inject the PAT dynamically.
