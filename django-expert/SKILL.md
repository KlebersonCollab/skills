---
name: django-expert
version: 1.5.0
description: "Expert level Django development focused on production architecture, security hardening, testing excellence, and deployment readiness."
category: development-python
---

# Django Expert: Professional Web Systems

> "The web framework for perfectionists with deadlines." - Esta skill garante que você cumpra os prazos sem sacrificar a qualidade técnica.

---

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar a implementação:
1. **Spec Check**: O arquivo `spec.md` existe e contém os critérios de aceitação (ACs)?
2. **Plan Check**: O arquivo `plan.md` define a arquitetura e schemas?
3. **Task Check**: A lista de tarefas em `tasks.md` está detalhada?

---

## Goal

Prover um framework de decisão para o desenvolvimento de aplicações Django robustas, focando em performance de banco de dados (ORM), interfaces reativas com HTMX e gerenciamento de dependências via `python-uv`.

---

## Workflow (12 Phases)

### Phase 1: SETUP & ARCHITECTURE
Configuração do ecossistema Django via UV e definição da arquitetura.
- **Rule**: Sempre usar `uv init` e gerenciar dependências via UV.
- **Architecture**: Adotar o layout `config/` e `apps/` com **Split Settings**.
- **Mandate**: Configurar `pyproject.toml` com seções para ferramentas (Ruff, Pytest).
- **Reference**: Ver [Production Architecture](references/architecture_prod.md).

### Phase 2: DATA_MODELING (ORM)
Definição de modelos e migrações.
- **Rule**: NUNCA usar `null=True` em campos de texto (usar string vazia por padrão).
- **Reference**: Ver [ORM Performance](references/orm_performance.md).

### Phase 3: LOGIC & FORMS
Implementação da regra de negócio.
- **Rule**: Preferir a lógica em `Services` ou `Managers` em vez de Views ou Modelos "gordos".
- **Mandate**: Utilizar `Django Forms/ModelForms` para validação de entrada, sempre.

### Phase 4: UI & REACTIVITY (HTMX)
Criação da interface de usuário.
- **Rule**: Usar **HTMX** para interações dinâmicas sem JavaScript customizado pesado.
- **Reference**: Ver [HTMX Patterns](references/htmx_patterns.md).

### Phase 5: TESTING & QUALITY
Garantia de estabilidade inicial.
- **Rule**: Usar `pytest-django` e `factory-boy` para testes rápidos e isolados.
- **Reference**: Ver [Testing Guide](references/testing.md).

### Phase 6: SECURITY & DEPLOY
Preparação básica para o mundo real.
- **Check**: Rodar `python manage.py check --deploy`.
- **Mandate**: Configurar CSRF, HSTS e Session Security.

### Phase 7: MAINTENANCE & DEBUGGING
Cultura de depuração e monitoramento.
- **Rule**: Seguir o [Systematic Debugging](references/debugging.md).
- **Logic**: Implementar logs estruturados e monitoramento de tarefas.

### Phase 8: API DEVELOPMENT (DRF)
Criação de APIs robustas com Django REST Framework.
- **Rule**: Usar ViewSets e Serializers explícitos para garantir contratos claros.
- **Mandate**: Implementar autenticação via JWT ou Token e permissões granulares.
- **Reference**: Ver [API Patterns](references/api_patterns.md).

### Phase 9: CROSS-CUTTING CONCERNS
Implementação de funcionalidades transversais.
- **Rule**: Usar Caching para endpoints caros e Signals para efeitos colaterais.
- **Reference**: Ver [Cross-Cutting Concerns](references/cross_cutting.md).

### Phase 10: SECURITY HARDENING
Fortalecimento da aplicação contra vulnerabilidades comuns.
- **Rule**: Implementar Argon2, validadores de senha de 12+ caracteres e RBAC.
- **Mandate**: Validar tipo e tamanho de todos os arquivos enviados pelo usuário.
- **Reference**: Ver [Security Hardening](references/security.md).

### Phase 11: CONTINUOUS TESTING & COVERAGE
Garantia de estabilidade e prevenção de regressões.
- **Rule**: Seguir o workflow TDD (Red-Green-Refactor) via Pytest.
- **Mandate**: NUNCA usar fixtures JSON/YAML; usar exclusivamente Factories.
- **Coverage**: Manter cobertura > 90% em Models e Services.
- **Reference**: Ver [Testing Excellence](references/testing.md).

### Phase 12: DEPLOYMENT READINESS & VERIFICATION
Loop final de qualidade antes da entrega.
- **Rule**: Rodar checklist de deploy e auditoria de segurança (pip-audit).
- **Mandate**: Nenhuma PR deve ser aceita sem passar no `ruff check` e `mypy`.
- **Reference**: Ver [Deployment & Verification](references/verification.md).

---

## Key Patterns

### 1. N+1 Avoidance
Sempre utilize `select_related` (FKs) e `prefetch_related` (M2M) para otimizar queries.

### 2. Modern UI with HTMX
Interfaces reativas sem a complexidade de SPAs.

### 3. API Patterns (DRF)
Contratos claros e segurança em primeiro lugar.

---

## Quality Rules

- **Code Quality**: Usar **Ruff** para linting e **Mypy** para verificação de tipos.
- **DRY ORM**: Consultas complexas devem residir em Custom Managers.
- **Service Layer**: Toda lógica complexa deve estar em `services.py`.
- **HTMX Fragments**: Templates modulares para retornos parciais.
- **Strict Testing**: Todo Bug fix deve vir acompanhado de um teste de regressão.

## Prohibited

- NUNCA usar `null=True` em CharField ou TextField (usar default vazio).
- NUNCA realizar queries dentro de loops (usar `select_related` / `prefetch_related`).
- NUNCA colocar chaves secretas ou credenciais no código (usar `django-environ`).
- NUNCA importar signals fora do método `ready()` do `AppConfig`.
- NUNCA usar `mark_safe` em input de usuário sem escape prévio.
- NUNCA concatenar variáveis diretamente em queries SQL brutas.
- NUNCA submeter código sem testes de unidade para novas funcionalidades.
- NUNCA realizar chamadas externas reais em ambiente de teste (usar Mocks).
- NUNCA expor o `id` sequencial do banco em URLs públicas se a segurança for crítica (usar UUID).

---

## Output Structure

| Guia | Descrição |
|------|-----------|
| [ORM Performance](references/orm_performance.md) | Otimização de queries e Managers. |
| [Architecture Prod](references/architecture_prod.md) | Layout apps/ e Split Settings. |
| [API Patterns](references/api_patterns.md) | DRF, ViewSets e Serializers. |
| [Security Hardening](references/security.md) | Argon2, RBAC e File Security. |
| [Testing Excellence](references/testing.md) | Factories, Pytest e Mocking. |
| [Deployment & Verification](references/verification.md) | Ruff, Mypy e Checklists. |
| [Cross-Cutting](references/cross_cutting.md) | Caching, Signals e Middleware. |
| [HTMX Patterns](references/htmx_patterns.md) | Fragmentos e interações dinâmicas. |
| [Background Tasks](references/background_tasks.md) | Celery, Redis e Idempotência. |
| [Forms & Validation](references/forms.md) | Validação avançada e ModelForms. |
| [Debugging](references/debugging.md) | Depuração sistemática e Logs. |
