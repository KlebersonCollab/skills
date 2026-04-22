# Django Deployment & Verification Reference

## 1. Static Analysis Loop
Mantenha o código limpo e tipado.

### Linting (Ruff)
Ruff substitui Flake8, Isort e Black com 10-100x mais velocidade.
```bash
uv run ruff check . --fix
uv run ruff format .
```

### Type Checking (Mypy)
```bash
uv run mypy .
```

## 2. Security Auditing
NUNCA deploye com vulnerabilidades conhecidas.

```bash
# Auditoria de dependências
uv run pip-audit

# Auditoria de código (Bandit)
uv run bandit -r apps/
```

## 3. Pre-Deployment Checklist
Execute antes de cada merge em `main` ou deploy em `prod`.

- [ ] `python manage.py check --deploy` (Zero erros)
- [ ] `python manage.py makemigrations --check` (Nenhuma migration faltando)
- [ ] `pytest --cov` (Cobertura mínima atingida)
- [ ] `DEBUG=False` em ambiente de produção.
- [ ] Variáveis críticas (`SECRET_KEY`, `DATABASE_URL`) configuradas no ambiente.

## 4. GitHub Actions Template
Boilerplate para verificação automática.

```yaml
name: Django Verification
on: [push, pull_request]
jobs:
  verify:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Install uv
        run: curl -LsSf https://astral.sh/uv/install.sh | sh
      - name: Install dependencies
        run: uv sync
      - name: Lint & Type Check
        run: |
          uv run ruff check .
          uv run mypy .
      - name: Run Tests
        run: uv run pytest --cov
```

## 5. Performance Verification
- **N+1 Queries**: Use `django-debug-toolbar` em dev. Alvo: < 50 queries/página.
- **Indexes**: Verifique `db_index=True` em campos usados em filtros frequentes.
