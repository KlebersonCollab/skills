# Administration & Security: Azure DevOps

Este guia foca na gestão de identidades, permissões e governança administrativa no Azure DevOps.

---

## 1. Gestão de Usuários e Níveis de Acesso

Adicionar ou remover usuários da organização e definir seu nível de licença (Basic, Stakeholder, etc.).

### Listar Usuários
`az devops user list`

### Adicionar Usuário (CLI)
```bash
az devops user add \
    --email-id "novo.dev@exemplo.com" \
    --license-type "basic" \
    --send-email true
```

## 2. Times e Membros

Organizar desenvolvedores em times para facilitar a gestão de backlogs e permissões.

### Criar um Time
`az devops team create --name "Frontend-Team" --project "MeuProjeto"`

### Adicionar Membro ao Time
`az devops team user add --team "Frontend-Team" --user "dev@exemplo.com"`

## 3. Segurança e Grupos

O AzDO utiliza grupos de segurança para aplicar permissões em massa.

### Listar Grupos de Segurança
`az devops security group list --project "MeuProjeto"`

### Gerenciar Permissões
Comandos complexos que permitem definir flags de permissão para namespaces específicos (ex: Git Repositories, Build).
`az devops security permission update --id {namespaceId} --subject {email} --token {token} --allow-bit {bit}`

## 4. Extensões (Marketplace)

Auditar e gerenciar extensões que adicionam funcionalidades à organização.

### Listar Extensões Instaladas
`az devops extension list`

### Instalar uma Extensão
`az devops extension install --extension-id "nome-ext" --publisher-id "editor"`

## 5. Auditoria de Segurança

- **Branch Policies**: Verifique periodicamente se as branches `main` possuem políticas de PR e build ativas (`az repos policy list`).
- **Access Review**: Liste usuários e seus últimos logins para remover contas inativas.
- **Service Connection Audit**: Verifique se há credenciais expiradas em conexões críticas.

## 6. Boas Práticas

- **Groups over Users**: Nunca atribua permissões diretamente a um usuário; sempre utilize Grupos de Segurança.
- **Minimum Access**: Comece com `Stakeholder` e suba para `Basic` apenas se necessário.
- **Policy Enforcement**: Utilize políticas de branch para garantir que nenhum código entre em produção sem revisão.
