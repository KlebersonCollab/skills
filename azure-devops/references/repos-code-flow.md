# Repos & Code Flow: Azure DevOps

Azure Repos fornece ferramentas de controle de versão (Git) e um fluxo de trabalho robusto para revisão de código através de **Pull Requests (PRs)**.

## 1. Gestão de Repositórios

Listar, criar e deletar repositórios em um projeto:

`GET https://dev.azure.com/{org}/{project}/_apis/git/repositories?api-version=7.1`

**Comando CLI**:
`az repos list --project {project}`

## 2. Pull Requests (PRs)

O ciclo de vida de um PR inclui:
- **Criação**: Definir a branch de origem e destino, título, descrição e revisores.
- **Revisão**: Adicionar comentários, threads de discussão e votos.
- **Conclusão**: Merge na branch de destino com a estratégia escolhida (Merge, Squash, Rebase).

### Criando um PR (POST)
```json
{
  "sourceRefName": "refs/heads/feature/auth",
  "targetRefName": "refs/heads/develop",
  "title": "Adicionar login via OAuth2",
  "description": "Implementação baseada na issue #42",
  "reviewers": [
    {
      "id": "uuid-revisor-1"
    }
  ]
}
```

## 3. Comentários e Discussões

As discussões em PRs são organizadas em **Threads**. Você pode listar e interagir com elas programaticamente:

`GET https://dev.azure.com/{org}/{project}/_apis/git/repositories/{repoId}/pullRequests/{prId}/threads?api-version=7.1`

## 4. Políticas de Branch (Branch Policies)

Políticas garantem a qualidade do código antes do merge. Exemplos:
- Número mínimo de revisores.
- Build de pipelines bem-sucedido.
- Resolução de todos os comentários.
- Padronização de nomes de branch.

Você pode consultar políticas aplicadas a uma branch:

`GET https://dev.azure.com/{org}/{project}/_apis/policy/evaluations?api-version=7.1&targetRefName=refs/heads/main`

## 5. Estratégias de Merge

Ao completar um PR, você pode escolher a estratégia:
- **No-fast-forward (Merge)**: Mantém o histórico completo.
- **Squash merge**: Consolida todos os commits em um só.
- **Rebase and merge**: Reaplica commits na ponta da branch alvo.

## 6. Melhores Práticas

- **Feature Branching**: Nunca trabalhe diretamente na branch principal (`main` ou `develop`).
- **Descrições Claras**: Sempre vincule o PR a um Work Item do Board (ex: `refs/heads/feature/42-auth`).
- **Small PRs**: Pull Requests menores são mais fáceis de revisar e diminuem a chance de bugs.
- **Automation**: Utilize pipelines para executar linters e testes unitários automaticamente ao criar um PR.
