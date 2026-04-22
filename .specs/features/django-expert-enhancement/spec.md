# Spec: Django Expert Enhancement (DRF & Production Patterns)

## Context
A skill `django-expert` atual foca em performance ORM, HTMX e UV. Identificamos a necessidade de integrar padrões de arquitetura de produção e suporte a APIs REST (DRF) baseados na skill `django-patterns` do ECC.

## Objectives
1.  **Integrar Layout de Produção**: Adotar a estrutura `apps/` e `config/` (Split Settings).
2.  **Adicionar Suporte a DRF**: Incluir padrões de ViewSets e Serializers.
3.  **Expandir Padrões de ORM**: Adicionar `QuerySet.as_manager()` e Managers customizados.
4.  **Implementar Cross-Cutting Concerns**: Caching, Signals e Middleware.
5.  **Security Hardening**: Integrar padrões de autenticação forte, RBAC e validação de arquivos.
6.  **Testing Excellence (TDD)**: Implementar padrões de Factory Boy, Pytest e Mocking avançado.
7.  **Deployment Readiness & Verification**: Estabelecer loops de auditoria (Ruff, Mypy, pip-audit) e checklists de deploy.

## Acceptance Criteria (AC)
- [x] `SKILL.md` atualizado com o novo Recommended Layout.
- [x] Novos mandatos de "Split Settings" adicionados à Fase 1.
- [x] Nova seção "API Development (DRF)" adicionada com padrões de ViewSet/Serializer.
- [x] Seção de ORM expandida com mandatos de Managers e Constraints.
- [x] Seção de "Cross-Cutting Concerns" (Cache, Signals, Middleware) adicionada.
- [x] Nova Fase 10: "Security Hardening" adicionada ao `SKILL.md`.
- [x] Referência `references/security.md` expandida com padrões de Hardening.
- [x] Nova Fase 11: "Testing Excellence" adicionada ao `SKILL.md`.
- [x] Referência `references/testing.md` expandida com Factories e Mocks.
- [x] Nova Fase 12: "Deployment Readiness & Verification" adicionada ao `SKILL.md`.
- [x] Referência `references/verification.md` criada com checklists de deploy e auditoria.
- [x] Mandatos de uso de Ruff e Mypy incluídos nas Quality Rules.
- [x] Arquivos de referência criados em `django-expert/references/` para os novos tópicos.
- [x] Validação final via `skill-factory-validator` (ou auditoria manual rigorosa).

## Technical Requirements
- Manter compatibilidade com **SDD**.
- Focar em **Python-UV**.
- Seguir as regras globais de design do Hub.
