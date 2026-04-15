# Azure DevOps Skill

Esta skill fornece uma interface profissional para gerenciar recursos do **Azure DevOps (AzDO)** através do Gemini CLI. Ela abstrai a complexidade da API REST e das operações JSON Patch, permitindo gerenciar Boards, Repos, Pipelines e Artifacts de forma eficiente.

## Recursos

- **Azure Boards**: CRUD de Work Items e consultas via WIQL.
- **Azure Repos**: Gestão de Pull Requests e políticas de branch.
- **Azure Pipelines**: Execução e monitoramento de builds e releases.
- **Azure Artifacts**: Gestão de feeds e pacotes.
- **Segurança**: Autenticação robusta via Personal Access Tokens (PAT).

## Como Usar

Para ativar as diretrizes desta skill em sua sessão atual, utilize:

```bash
activate_skill azure-devops
```

## Estrutura da Skill

- `SKILL.md`: Documento principal com o workflow e comandos.
- `references/`: Guias detalhados sobre cada módulo do AzDO.
- `examples/`: Modelos de queries WIQL e payloads JSON Patch.

---

Baseado no padrão da skill `python-uv` e `flutter-fvm`.
