---
name: django-expert
version: 1.5.0
description: "Expert level Django development focused on production architecture, security hardening, testing excellence, and deployment readiness."
category: development-python
---

# Django Expert: Professional Web Systems

> "The web framework for perfectionists with deadlines." - This skill ensures you meet deadlines without sacrificing technical quality.

---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the guidelines of the `token-distiller` skill.
1. **Context Check**: Have you rehydrated the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define the architecture, schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---
## Goal

Provide a decision framework for developing robust Django applications, focusing on database performance (ORM), reactive interfaces with HTMX, and dependency management via `python-uv`.

---

## Workflow (12 Phases)

### Phase 1: SETUP & ARCHITECTURE
Configuration of the Django ecosystem via UV and architecture definition.
- **Rule**: Always use `uv init` and manage dependencies via UV.
- **Architecture**: Adopt the `config/` and `apps/` layout with **Split Settings**.
- **Mandate**: Configure `pyproject.toml` with sections for tools (Ruff, Pytest).
- **Reference**: See [Production Architecture](references/architecture_prod.md).

### Phase 2: DATA_MODELING (ORM)
Model and migration definition.
- **Rule**: NEVER use `null=True` on text fields (use empty string by default).
- **Reference**: See [ORM Performance](references/orm_performance.md).

### Phase 3: LOGIC & FORMS
Business rule implementation.
- **Rule**: Prefer logic in `Services` or `Managers` instead of "fat" Views or Models.
- **Mandate**: Always use `Django Forms/ModelForms` for input validation.

### Phase 4: UI & REACTIVITY (HTMX)
User interface creation.
- **Rule**: Use **HTMX** for dynamic interactions without heavy custom JavaScript.
- **Reference**: See [HTMX Patterns](references/htmx_patterns.md).

### Phase 5: TESTING & QUALITY
Initial stability guarantee.
- **Rule**: Use `pytest-django` and `factory-boy` for fast, isolated tests.
- **Reference**: See [Testing Guide](references/testing.md).

### Phase 6: SECURITY & DEPLOY
Basic real-world preparation.
- **Check**: Run `python manage.py check --deploy`.
- **Mandate**: Configure CSRF, HSTS, and Session Security.

### Phase 7: MAINTENANCE & DEBUGGING
Debugging and monitoring culture.
- **Rule**: Follow [Systematic Debugging](references/debugging.md).
- **Logic**: Implement structured logs and task monitoring.

### Phase 8: API DEVELOPMENT (DRF)
Creating robust APIs with Django REST Framework.
- **Rule**: Use ViewSets and explicit Serializers to ensure clear contracts.
- **Mandate**: Implement authentication via JWT or Token and granular permissions.
- **Reference**: See [API Patterns](references/api_patterns.md).

### Phase 9: CROSS-CUTTING CONCERNS
Implementation of cross-cutting features.
- **Rule**: Use Caching for expensive endpoints and Signals for side effects.
- **Reference**: See [Cross-Cutting Concerns](references/cross_cutting.md).

### Phase 10: SECURITY HARDENING
Strengthening the application against common vulnerabilities.
- **Rule**: Implement Argon2, 12+ character password validators, and RBAC.
- **Mandate**: Validate the type and size of all user-uploaded files.
- **Reference**: See [Security Hardening](references/security.md).

### Phase 11: CONTINUOUS TESTING & COVERAGE
Ensuring stability and preventing regressions.
- **Rule**: Follow the TDD workflow (Red-Green-Refactor) via Pytest.
- **Mandate**: NEVER use JSON/YAML fixtures; use Factories exclusively.
- **Coverage**: Maintain > 90% coverage on Models and Services.
- **Reference**: See [Testing Excellence](references/testing.md).

### Phase 12: DEPLOYMENT READINESS & VERIFICATION
Final quality loop before delivery.
- **Rule**: Run deployment checklist and security audit (pip-audit).
- **Mandate**: No PR should be accepted without passing `ruff check` and `mypy`.
- **Reference**: See [Deployment & Verification](references/verification.md).

---

## Key Patterns

### 1. N+1 Avoidance
Always use `select_related` (FKs) and `prefetch_related` (M2M) to optimize queries.

### 2. Modern UI with HTMX
Reactive interfaces without the complexity of SPAs.

### 3. API Patterns (DRF)
Clear contracts and security first.

---

## Quality Rules

- **Code Quality**: Use **Ruff** for linting and **Mypy** for type checking.
- **DRY ORM**: Complex queries must reside in Custom Managers.
- **Service Layer**: All complex logic must be in `services.py`.
- **HTMX Fragments**: Modular templates for partial returns.
- **Strict Testing**: Every bug fix must be accompanied by a regression test.

## Prohibited

- NEVER use `null=True` on CharField or TextField (use empty default).
- NEVER perform queries inside loops (use `select_related` / `prefetch_related`).
- NEVER put secret keys or credentials in code (use `django-environ`).
- NEVER import signals outside the `ready()` method of `AppConfig`.
- NEVER use `mark_safe` on user input without prior escaping.
- NEVER concatenate variables directly into raw SQL queries.
- NEVER submit code without unit tests for new features.
- NEVER perform real external calls in a testing environment (use Mocks).
- NEVER expose sequential database `id` in public URLs if security is critical (use UUID).

---

## Output Structure

| Guide | Description |
|------|-----------|
| [ORM Performance](references/orm_performance.md) | Query and Manager optimization. |
| [Architecture Prod](references/architecture_prod.md) | apps/ layout and Split Settings. |
| [API Patterns](references/api_patterns.md) | DRF, ViewSets, and Serializers. |
| [Security Hardening](references/security.md) | Argon2, RBAC, and File Security. |
| [Testing Excellence](references/testing.md) | Factories, Pytest, and Mocking. |
| [Deployment & Verification](references/verification.md) | Ruff, Mypy, and Checklists. |
| [Cross-Cutting](references/cross_cutting.md) | Caching, Signals, and Middleware. |
| [HTMX Patterns](references/htmx_patterns.md) | Fragments and dynamic interactions. |
| [Background Tasks](references/background_tasks.md) | Celery, Redis, and Idempotency. |
| [Forms & Validation](references/forms.md) | Advanced validation and ModelForms. |
| [Debugging](references/debugging.md) | Systematic debugging and Logs. |
