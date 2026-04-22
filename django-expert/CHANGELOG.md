# Changelog - Django Expert

## [1.5.0] - 2026-04-22

### Added
- Seção de **Deployment Readiness & Verification** (Fase 12).
- Mandato de uso do **Ruff** para linting e **Mypy** para Type Checking.
- Checklists de pre-deployment e auditoria de segurança.
- Template de **GitHub Actions** para verificação automatizada.
- Consolidação de redundâncias no `SKILL.md`.

## [1.4.0] - 2026-04-22

### Added
- Seção de **Testing Excellence** (Fase 11).
- Mandatos para uso de **Factories** (Factory Boy) em vez de fixtures estáticas.
- Configurações de infraestrutura de teste ultrarrápida (SQLite in-memory, `--nomigrations`).
- Padrões de **Mocking** para serviços externos e APIs.
- Metas explícitas de cobertura de código por camada.

## [1.3.0] - 2026-04-22

### Added
- Seção de **Security Hardening** (Fase 10).
- Padrões de hashing com Argon2 e validadores de senha de 12+ caracteres.
- Mandatos para validação de tipo e tamanho de arquivos.
- Padrões de RBAC (Role-Based Access Control) e Throttling para APIs.
- Regras proibitivas contra SQL Injection e uso inseguro de `mark_safe`.

## [1.2.0] - 2026-04-22

### Added
- Padrões de arquitetura de produção (Split Settings e layout `apps/`).
- Seção de desenvolvimento de APIs com Django REST Framework (DRF).
- Referências detalhadas para Caching, Signals e Middleware.
- Padrões avançados de ORM (`as_manager()` e constraints de banco).

## [1.1.0] - 2026-04-20

### Added
- Seção `🔒 Prerequisites (Mandatory)` vinculando a skill obrigatoriamente ao framework SDD.

## [1.0.0] - 2026-04-18

### Added
- Inicialização da skill `django-expert`.
- Guia de performance ORM (N+1 solutions).
- Integração com HTMX e padrões de reatividade.
- Suporte nativo ao workflow `python-uv`.
- Estrutura recomendada de camadas (Services/Managers).
