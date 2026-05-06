# Repos & Code Flow: Azure DevOps

Azure Repos provides version control tools (Git) and a robust workflow for code review through **Pull Requests (PRs)**.

## 1. Repository Management

List, create, and delete repositories in a project:

`GET https://dev.azure.com/{org}/{project}/_apis/git/repositories?api-version=7.1`

**CLI Command**:
`az repos list --project {project}`

## 2. Pull Requests (PRs)

The PR life cycle includes:
- **Creation**: Define origin and target branch, title, description, and reviewers.
- **Review**: Add comments, discussion threads, and votes.
- **Completion**: Merge into the target branch with the chosen strategy (Merge, Squash, Rebase).

### Creating a PR (POST)
```json
{
  "sourceRefName": "refs/heads/feature/auth",
  "targetRefName": "refs/heads/develop",
  "title": "Add login via OAuth2",
  "description": "Implementation based on issue #42",
  "reviewers": [
    {
      "id": "uuid-reviewer-1"
    }
  ]
}
```

## 3. Comments and Discussions

PR discussions are organized into **Threads**. You can list and interact with them programmatically:

`GET https://dev.azure.com/{org}/{project}/_apis/git/repositories/{repoId}/pullRequests/{prId}/threads?api-version=7.1`

## 4. Branch Policies

Policies ensure code quality before merging. Examples:
- Minimum number of reviewers.
- Successful pipeline build.
- Resolution of all comments.
- Branch name standardization.

You can query policies applied to a branch:

`GET https://dev.azure.com/{org}/{project}/_apis/policy/evaluations?api-version=7.1&targetRefName=refs/heads/main`

## 5. Merge Strategies

When completing a PR, you can choose the strategy:
- **No-fast-forward (Merge)**: Maintains full history.
- **Squash merge**: Consolidates all commits into one.
- **Rebase and merge**: Re-applies commits on top of the target branch.

## 6. Best Practices

- **Feature Branching**: Never work directly on the primary branch (`main` or `develop`).
- **Clear Descriptions**: Always link the PR to a Board Work Item (e.g., `refs/heads/feature/42-auth`).
- **Small PRs**: Smaller Pull Requests are easier to review and decrease the chance of bugs.
- **Automation**: Use pipelines to run linters and unit tests automatically when creating a PR.
