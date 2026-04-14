# Changelog — Python com UV

Todas as mudanças notáveis desta skill serão documentadas neste arquivo.

O formato segue o padrão [Keep a Changelog](https://keepachangelog.com/pt-BR/1.1.0/), e o versionamento segue [Semantic Versioning](https://semver.org/lang/pt-BR/).

---

## [2.0.0] — 2026-04-14

### Adicionado
- **SKILL.md**: Reescrita completa como instrução acionável para agentes AI.
  - Overview com tabela de ferramentas substituídas
  - Version Awareness com tabela de features por versão
  - When to Use / Skip com triggers explícitos
  - UV Commands Overview com 14 comandos documentados
  - Tool vs UVX Decision Tree
  - Project Initialization workflow completo
  - Inline Script Metadata (PEP 723) com exemplos
  - Best Practices DO/DON'T por categoria
  - Troubleshooting com 4 cenários comuns
  - Migration Guides (pip, poetry, pipx)
  - Links para Reference Documentation
- **references/installation-and-setup.md**: Instalação cross-platform, virtual environments, troubleshooting de PATH.
- **references/project-management.md**: Init, pyproject.toml, dependências, lockfile, workspaces, ferramentas de qualidade.
- **references/tool-management.md**: `uv tool install` vs `uvx`, feature matrix, anti-patterns, manutenção.
- **references/inline-script-metadata.md**: PEP 723, sintaxe, exemplos (scraper, FastAPI, CLI, data analysis).
- **references/python-environment.md**: Versões Python, Python 3.14 default, free-threaded Python, paths por plataforma.
- **references/ci-cd-workflows.md**: GitHub Actions (básico, matrix, cache), Docker (básico, multi-stage), packaging.
- **examples/pyproject-toml-example.md**: Template completo e anotado com todas as seções de configuração.
- **examples/github-actions-example.md**: Workflow CI/CD de produção com quality checks, matrix teste e publicação.
- **README.md**: Reescrita seguindo padrão do skill-factory com tabelas de references e examples.

### Removido
- Conteúdo tutorial do README (movido para references modulares)
- Seções genéricas sem instruções acionáveis

### Alterado
- **SKILL.md**: De resumo genérico para instrução direta ao agente com decision trees, tabelas e code blocks
- **README.md**: De tutorial humano para documentação do skill com links modulares
- Versão: `1.0.0` → `2.0.0` (breaking change — reescrita completa)

---

## [1.0.0] — 2026-04-14

### Adicionado
- Skill inicial para uso de Python com UV
- Documentação completa de instalação cross-platform
- Guia de configuração inicial com `uv init`
- Gerenciamento de dependências com `pyproject.toml`
- Boas práticas de Python moderno (type hints, ruff, mypy, pytest)
- Workflows de desenvolvimento, CI/CD e packaging
- Comandos essenciais do UV
- Referências à documentação oficial
- Exemplos práticos e testáveis