# Python com UV — Skill

Gerenciador de pacotes e projetos Python extremamente rápido, escrito em Rust. Substitui pip, pipx, poetry, pyenv e mais.

---

## 🚀 Visão Geral

Esta skill orienta o desenvolvimento Python moderno usando o `uv`. Ela cobre desde a inicialização de projetos simples até arquiteturas corporativas complexas (Django Pro), garantindo performance extrema e segurança.

### Capacidades Principais:
- **Gestão de Projetos**: Inicialização (`uv init`), dependências e lockfiles.
- **Django Pro**: Padrões de Service Layer, Async Django e Bulk Operations.
- **Otimização de Performance**: Resolução de N+1 (ORM), Profiling e Memory Management.
- **Desenvolvimento Async**: Padrões para FastAPI, httpx e testes assíncronos.
- **Segurança**: Auditoria de BOLA/IDOR e ferramentas de segurança integrada.
- **Scripts PEP 723**: Execução de scripts únicos com dependências declaradas.

---

## 📂 Estrutura da Skill

```text
python-uv/
├── SKILL.md                 # Definição e workflow (v4.0.0 Purista)
├── README.md                # Este guia
├── CHANGELOG.md             # Histórico de versões
├── references/              # Referências Detalhadas (Django, Async, Setup, etc.)
├── examples/                # Exemplos Práticos (CI/CD, pyproject, patterns)
└── resources/               # Templates e Ativos
```

---

## 🛠️ Como Usar

1. **Inicie um projeto**: `uv init meu-app`
2. **Otimize seu Django**: Consulte `references/django-workflow.md` para padrões de N+1.
3. **Analise Performance**: Use `uv run python manage.py shell_plus --print-sql`.

---

## 🎯 Exemplos de Comandos

- *"Refatore esta view Django para resolver o problema de N+1 queries usando select_related."*
- *"Configure uma Service Layer para este projeto Django gerenciado pelo uv."*
- *"Como uso o uv para rodar o bandit e auditar a segurança do meu projeto?"*
