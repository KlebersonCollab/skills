# Changelog - Django Expert

## [1.5.0] - 2026-04-22

### Added
- **Deployment Readiness & Verification** section (Phase 12).
- Mandate to use **Ruff** for linting and **Mypy** for Type Checking.
- Pre-deployment checklists and security audit.
- **GitHub Actions** template for automated verification.
- Consolidated redundancies in `SKILL.md`.

## [1.4.0] - 2026-04-22

### Added
- **Testing Excellence** section (Phase 11).
- Mandates for using **Factories** (Factory Boy) instead of static fixtures.
- Ultra-fast test infrastructure configurations (SQLite in-memory, `--nomigrations`).
- **Mocking** patterns for external services and APIs.
- Explicit code coverage goals per layer.

## [1.3.0] - 2026-04-22

### Added
- **Security Hardening** section (Phase 10).
- Hashing patterns with Argon2 and 12+ character password validators.
- Mandates for file type and size validation.
- RBAC (Role-Based Access Control) and Throttling patterns for APIs.
- Prohibitive rules against SQL Injection and unsafe use of `mark_safe`.

## [1.2.0] - 2026-04-22

### Added
- Production architecture patterns (Split Settings and `apps/` layout).
- API development section with Django REST Framework (DRF).
- Detailed references for Caching, Signals, and Middleware.
- Advanced ORM patterns (`as_manager()` and database constraints).

## [1.1.0] - 2026-04-20

### Added
- `🔒 Prerequisites (Mandatory)` section linking the skill mandatorily to the SDD framework.

## [1.0.0] - 2026-04-18

### Added
- Initialization of the `django-expert` skill.
- ORM performance guide (N+1 solutions).
- HTMX integration and reactivity patterns.
- Native support for the `python-uv` workflow.
- Recommended layer structure (Services/Managers).
