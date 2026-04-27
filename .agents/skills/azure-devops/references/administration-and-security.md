# Administration & Security: Azure DevOps

This guide focuses on identity management, permissions, and administrative governance in Azure DevOps.

---

## 1. User Management and Access Levels

Add or remove users from the organization and define their license level (Basic, Stakeholder, etc.).

### List Users
`az devops user list`

### Add User (CLI)
```bash
az devops user add \
    --email-id "new.dev@example.com" \
    --license-type "basic" \
    --send-email true
```

## 2. Teams and Members

Organize developers into teams to facilitate backlog and permission management.

### Create a Team
`az devops team create --name "Frontend-Team" --project "MyProject"`

### Add Member to Team
`az devops team user add --team "Frontend-Team" --user "dev@example.com"`

## 3. Security and Groups

AzDO uses security groups to apply permissions in bulk.

### List Security Groups
`az devops security group list --project "MyProject"`

### Manage Permissions
Complex commands that allow defining permission flags for specific namespaces (e.g., Git Repositories, Build).
`az devops security permission update --id {namespaceId} --subject {email} --token {token} --allow-bit {bit}`

## 4. Extensions (Marketplace)

Audit and manage extensions that add functionality to the organization.

### List Installed Extensions
`az devops extension list`

### Install an Extension
`az devops extension install --extension-id "name-ext" --publisher-id "publisher"`

## 5. Security Audit

- **Branch Policies**: Periodically check if `main` branches have active PR and build policies (`az repos policy list`).
- **Access Review**: List users and their last logins to remove inactive accounts.
- **Service Connection Audit**: Check for expired credentials in critical connections.

## 6. Best Practices

- **Groups over Users**: Never assign permissions directly to a user; always use Security Groups.
- **Minimum Access**: Start with `Stakeholder` and upgrade to `Basic` only if necessary.
- **Policy Enforcement**: Use branch policies to ensure no code enters production without review.
