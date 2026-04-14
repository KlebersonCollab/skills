---
name: python-uv
version: 1.0.0
description: "Skill para uso e boas práticas de Python utilizando UV como gerenciador de pacotes e versionamento"
category: development-workflow
---

# Python com UV

> Capacitar desenvolvedores a usar Python moderno com UV, o gerenciador de pacotes rápido e confiável da Astral.

---

## Goal

Capacitar desenvolvedores a adotar Python moderno utilizando UV como gerenciador de pacotes e versionamento, seguindo boas práticas de desenvolvimento, gerenciamento de dependências e workflows eficientes.

---

## Output Structure

Esta skill produz:
- Guias práticos para instalação e configuração do UV
- Comandos essenciais para gerenciamento de pacotes e ambientes
- Configurações recomendadas para projetos Python modernos
- Workflows de desenvolvimento, CI/CD e packaging
- Referências à documentação oficial do UV

---

## Quality Rules

- **Referência Oficial**: Sempre referenciar a documentação oficial do UV (https://docs.astral.sh/uv/)
- **Exemplos Práticos**: Todos os exemplos devem ser testáveis e reproduzíveis
- **Cross-Platform**: Comandos e configurações devem funcionar em macOS, Linux e Windows
- **Boas Práticas Modernas**: Promover type hints, ruff, pytest, mypy e outras ferramentas modernas
- **Versionamento Semântico**: Seguir Semantic Versioning para configurações de projeto
- **Lockfiles**: Sempre usar e versionar uv.lock para reprodução exata de ambientes

## Prohibited

- **Não assumir SO**: Não assumir configurações específicas de sistema operacional
- **Não duplicar documentação**: Não copiar conteúdo da documentação oficial, apenas referenciar
- **Não promover práticas obsoletas**: Não ensinar práticas deprecated como usar pip sem venv
- **Não ignorar type safety**: Não promover código Python sem type hints
- **Não pular validação**: Não pular linting, formatting ou type checking