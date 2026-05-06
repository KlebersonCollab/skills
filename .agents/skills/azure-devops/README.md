# Azure DevOps Skill

This skill provides a professional interface for managing **Azure DevOps (AzDO)** resources through the Gemini CLI. It abstracts the complexity of the REST API and JSON Patch operations, allowing you to manage Boards, Repos, Pipelines, and Artifacts efficiently.

## Features

- **Azure Boards**: Work Items CRUD and queries via WIQL.
- **Azure Repos**: Pull Request management and branch policies.
- **Azure Pipelines**: Execution and monitoring of builds and releases.
- **Azure Artifacts**: Feed and package management.
- **Security**: Robust authentication via Personal Access Tokens (PAT).

## How to Use

To activate the guidelines of this skill in your current session, use:

```bash
activate_skill azure-devops
```

## Skill Structure

- `SKILL.md`: Main document with workflow and commands.
- `references/`: Detailed guides on each AzDO module.
- `examples/`: WIQL query templates and JSON Patch payloads.

---

Based on the `python-uv` and `flutter-fvm` skill pattern.
