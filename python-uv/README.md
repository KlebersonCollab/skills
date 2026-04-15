# Python com UV — Skill

Gerenciador de pacotes e projetos Python extremamente rápido, escrito em Rust. Substitui pip, pipx, poetry, pyenv e mais.

---

## 🚀 Visão Geral

Esta skill orienta o desenvolvimento Python moderno usando o `uv`. Ela cobre desde a inicialização de projetos simples até arquiteturas assíncronas complexas, garantindo performance e reprodutibilidade.

### Capacidades Principais:
- **Gestão de Projetos**: Inicialização (`uv init`), dependências e lockfiles.
- **Ambientes**: Gestão automática de venvs e versões de Python.
- **Desenvolvimento Async**: Padrões para FastAPI, httpx e testes assíncronos.
- **Ferramentas**: Instalação isolada de CLI tools (`ruff`, `mypy`, `pytest`).
- **Scripts PEP 723**: Execução de scripts únicos com dependências declaradas.

---

## 📂 Estrutura da Skill

```text
python-uv/
├── SKILL.md                 # Definição e workflow (v2.1.0)
├── README.md                # Este guia
├── CHANGELOG.md             # Histórico de versões
├── references/              # Referências Detalhadas
│   ├── async-development.md # NOVO!
│   ├── project-management.md
│   ├── tool-management.md
│   ├── python-environment.md
│   └── ...
└── examples/                # Exemplos Práticos
```

---

## 🛠️ Como Usar

1. **Inicie um projeto**: `uv init meu-app`
2. **Adicione dependências**: `uv add httpx`
3. **Desenvolva com Async**: Consulte `references/async-development.md` para padrões de concorrência.
4. **Teste seu código**: `uv run pytest`

---

## 🎯 Exemplos de Comandos

- *"Crie um script Python assíncrono usando httpx para fazer scraping de 5 URLs concorrentemente."*
- *"Configure um novo projeto FastAPI usando uv e configure o pytest-asyncio."*
- *"Como instalo o ruff globalmente usando o uv?"*
