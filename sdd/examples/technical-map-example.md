# Technical Map Example: "Task Manager" Application

> Este é um exemplo ilustrativo de como um `TECHNICAL-MAP.md` gerado pelo `sdd-explorer` deve ser estruturado.  
> Baseado em um cenário hipotético de um sistema de gerenciamento de tarefas.

## Stack

| Component | Technology | Version |
|---|---|---|
| Backend | Python + FastAPI | 0.110 |
| Database | PostgreSQL | 16 |
| ORM | SQLAlchemy | 2.0 |
| Frontend | React + TypeScript | 18 |
| Test Runner | pytest | 8 |
| Linter | ruff | 0.4 |
| CI/CD | GitHub Actions | N/A |

## Architecture

```
task-manager/
├── backend/
│   ├── app/
│   │   ├── main.py          # FastAPI entry point
│   │   ├── routers/         # API routes
│   │   ├── models/          # SQLAlchemy models
│   │   ├── schemas/         # Pydantic schemas
│   │   └── services/        # Business logic
│   ├── tests/
│   └── alembic/             # Migrations
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   ├── pages/
│   │   └── hooks/
│   └── package.json
└── docker-compose.yml
```

## Conventions

- **Naming**: `snake_case` for Python, `camelCase` for TypeScript.
- **Imports**: Absolute imports using project root as base.
- **Error Handling**: Custom exception classes with HTTP status mapping.
- **Testing**: pytest with fixtures; coverage target >80%.

## Critical Risks / Concerns

- **Authentication**: JWT tokens are stored in cookies without httpOnly flag (security risk).
- **Technical Debt**: `backend/app/services/task_service.py` has >500 lines and mixes I/O with business logic.
- **Missing Tests**: `frontend/src/hooks/useAuth.ts` has no unit tests.

## Related Files

- `spec/spec.md`: Feature specifications.
- `spec/plan.md`: Technical design documents.
- `.specs/codebase/CONVENTIONS.md`: Detailed coding conventions.
- `.specs/codebase/CONCERNS.md`: Expanded risk register.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "SDD-EXAMPLE-TECHNICAL-MAP"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-22T15:00:00Z"
evidence_checksum: "example-001"
```

