# Django Deployment & Verification Reference

## 1. Static Analysis Loop
Keep the code clean and typed.

### Linting (Ruff)
Ruff replaces Flake8, Isort, and Black with 10-100x more speed.
```bash
uv run ruff check . --fix
uv run ruff format .
```

### Type Checking (Mypy)
```bash
uv run mypy .
```

## 2. Security Auditing
NEVER deploy with known vulnerabilities.

```bash
# Dependency audit
uv run pip-audit

# Code audit (Bandit)
uv run bandit -r apps/
```

## 3. Pre-Deployment Checklist
Execute before each merge into `main` or deploy to `prod`.

- [ ] `python manage.py check --deploy` (Zero errors)
- [ ] `python manage.py makemigrations --check` (No missing migrations)
- [ ] `pytest --cov` (Minimum coverage reached)
- [ ] `DEBUG=False` in production environment.
- [ ] Critical variables (`SECRET_KEY`, `DATABASE_URL`) configured in the environment.

## 4. GitHub Actions Template
Boilerplate for automatic verification.

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
- **N+1 Queries**: Use `django-debug-toolbar` in dev. Target: < 50 queries/page.
- **Indexes**: Verify `db_index=True` on fields used in frequent filters.
