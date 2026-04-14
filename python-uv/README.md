# Python com UV

> Skill para desenvolvimento Python profissional com UV — o gerenciador de pacotes e projetos extremamente rápido.

[![Versão](https://img.shields.io/badge/Versão-2.0.0-blue)](#changelog)
[![References](https://img.shields.io/badge/References-6-brightgreen)](#-documentação-de-referência)
[![Examples](https://img.shields.io/badge/Examples-2-orange)](#-exemplos)

---

## 📖 Visão Geral

**UV** é um gerenciador de pacotes e projetos Python da [Astral](https://astral.sh/), escrito em Rust. Substitui pip, pipx, poetry, pyenv e mais — com performance **10-100x superior**.

Esta skill capacita agentes AI a:
- Inicializar e gerenciar projetos Python com UV
- Gerenciar dependências, lockfiles e ambientes virtuais
- Instalar e manter ferramentas de desenvolvimento (ruff, mypy, pytest)
- Executar scripts inline com PEP 723
- Configurar CI/CD com GitHub Actions
- Decidir entre `uv tool install` vs `uvx`
- Migrar de pip, pipx ou poetry para UV
- Resolver erros comuns do UV

---

## ⚙️ Como Usar

Esta skill é ativada automaticamente quando o agente detecta necessidade de:
- Criar ou configurar projetos Python
- Gerenciar dependências com `pyproject.toml`
- Instalar versões de Python ou ferramentas CLI
- Configurar pipelines de CI/CD para Python
- Resolver problemas relacionados a UV

O agente consulta o `SKILL.md` para instruções diretas e os `references/` para detalhamento aprofundado.

---

## 📚 Documentação de Referência

| # | Arquivo | Tópico |
|---|---------|--------|
| 1 | [installation-and-setup.md](references/installation-and-setup.md) | Instalação cross-platform, virtual environments, troubleshooting |
| 2 | [project-management.md](references/project-management.md) | Init, dependências, lockfile, workspaces, estrutura de projeto |
| 3 | [tool-management.md](references/tool-management.md) | `uv tool install` vs `uvx`, decision matrix, manutenção |
| 4 | [inline-script-metadata.md](references/inline-script-metadata.md) | PEP 723, scripts self-contained, exemplos práticos |
| 5 | [python-environment.md](references/python-environment.md) | Versões, Python 3.14, free-threaded Python, paths |
| 6 | [ci-cd-workflows.md](references/ci-cd-workflows.md) | GitHub Actions, Docker, packaging, publicação |

---

## 📋 Exemplos

| Arquivo | Descrição |
|---------|-----------|
| [pyproject-toml-example.md](examples/pyproject-toml-example.md) | Template completo e anotado de `pyproject.toml` moderno |
| [github-actions-example.md](examples/github-actions-example.md) | Workflow CI/CD completo com matrix, caching e publicação |

---

## 🔗 Referências Externas

- **UV Official Docs**: https://docs.astral.sh/uv/
- **UV GitHub**: https://github.com/astral-sh/uv
- **Ruff**: https://docs.astral.sh/ruff/
- **PEP 723**: https://peps.python.org/pep-0723/

---

## 📝 Changelog

Consulte o [CHANGELOG.md](CHANGELOG.md) para o histórico completo de versões.