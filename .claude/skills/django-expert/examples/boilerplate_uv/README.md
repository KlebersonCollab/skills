# Django Expert Boilerplate

This is a professional boilerplate following the `django-expert` skill.

## Requirements
- [UV](https://github.com/astral-sh/uv)

## How to run

1. **Install dependencies:**
   ```bash
   uv sync
   ```

2. **Configure environment:**
   Create a `.env` file based on `.env.example`.

3. **Run migrations:**
   ```bash
   uv run python manage.py migrate
   ```

4. **Start the server:**
   ```bash
   uv run python manage.py runserver
   ```

## Included Features
- **HTMX:** Server-side reactivity without heavy JS.
- **UV:** Ultra-fast package management.
- **Celery Ready:** Configured for background tasks.
- **Security:** Configurations via environment variables.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T14:05:00Z"
evidence_checksum: "8e52f6a"
```
