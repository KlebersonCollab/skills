---
name: django-expert
version: 1.0.0
description: "Expert level Django development focused on ORM performance, HTMX patterns, and modern Python-UV deployment."
category: development-python
---

# Django Expert: Professional Web Systems

> "The web framework for perfectionists with deadlines." - Esta skill garante que você cumpra os prazos sem sacrificar a qualidade técnica.

---

## Goal

Prover um framework de decisão para o desenvolvimento de aplicações Django robustas, focando em performance de banco de dados (ORM), interfaces reativas com HTMX e gerenciamento de dependências via `python-uv`.

---

## Workflow (6 Phases)

### Phase 1: ENVIRONMENT & SETUP
Configuração do ecossistema Django via UV.
- **Rule**: Sempre usar `uv init` e gerenciar `django` e `django-environ` via UV.
- **Mandate**: Configurar `pyproject.toml` com seções para ferramentas (Ruff, Pytest).

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
Garantia de estabilidade.
- **Rule**: Usar `pytest-django` e `factory-boy` para testes rápidos e isolados.
- **Reference**: Ver [Testing Guide](references/testing.md).

### Phase 6: SECURITY & DEPLOY
Preparação para o mundo real.
- **Check**: Rodar `python manage.py check --deploy`.
- **Mandate**: Configurar CSRF, HSTS e Session Security.

### Phase 7: MAINTENANCE & DEBUGGING
Cultura de depuração e monitoramento.
- **Rule**: Seguir o [Systematic Debugging](references/debugging.md).
- **Logic**: Implementar logs estruturados e monitoramento de tarefas.

---

## Key Patterns

### 1. N+1 Avoidance
Sempre utilize `select_related` (FKs) e `prefetch_related` (M2M) para otimizar queries.
```python
# DO THIS
users = User.objects.select_related('profile').all()

# NOT THIS
users = User.objects.all() # Cada acesso a user.profile gerará uma nova query
```

### 2. Modern UI with HTMX
```html
<button hx-post="/add-item/" hx-target="#item-list" hx-swap="beforeend">
    Add Item
</button>
```

---

## Quality Rules

- **Code Style (Ruff)**: Utilizar Ruff para linting e formatação automática.
- **DRY ORM**: Consultas complexas devem residir em Custom Managers.
- **Service Layer**: Toda lógica que envolve múltiplos modelos ou efeitos colaterais deve estar em `services.py`.
- **HTMX Fragments**: Templates devem ser modulares para permitir retornos parciais.

## Prohibited

- NUNCA usar `null=True` em CharField ou TextField (usar default vazio).
- NUNCA realizar queries dentro de loops (usar `select_related` / `prefetch_related`).
- NUNCA colocar chaves secretas ou credenciais no código (usar `django-environ`).
- NUNCA ignorar falhas de tarefas em background (implementar Retries).

## Output Structure

| Artefato | Localização | Descrição |
|----------|-------------|-----------|
| **Core Settings** | `core/settings.py` | Configurações divididas via variáveis de ambiente. |
| **Services** | `apps/<app>/services.py` | Camada de lógica de negócio pura. |
| **Partial Templates** | `templates/partials/` | Fragmentos HTML para uso com HTMX. |
| **Tasks** | `apps/<app>/tasks.py` | Definição de tarefas Celery. |

## Detailed References

| Guia | Tópico |
|------|--------|
| [ORM Performance](references/orm_performance.md) | Otimização de queries e Managers. |
| [HTMX Patterns](references/htmx_patterns.md) | Fragmentos e interações dinâmicas. |
| [Background Tasks](references/background_tasks.md) | Celery, Redis e Idempotência. |
| [Forms & Validation](references/forms.md) | Validação avançada e ModelForms. |
| [Testing](references/testing.md) | Pytest-django e Factories. |
| [Debugging](references/debugging.md) | Depuração sistemática e Logs. |
| [Security](references/security.md) | Checklist de deploy e segurança. |

## Quality Rules

- **DRY ORM**: Consultas complexas devem residir em Custom Managers.
- **Service Layer**: Toda lógica que envolve múltiplos modelos ou efeitos colaterais deve estar em `services.py`.
- **HTMX Fragments**: Templates devem ser modulares para permitir retornos parciais.
- **Strict Testing**: Todo Bug fix deve vir acompanhado de um teste de regressão.

## Prohibited

- NUNCA usar `null=True` em CharField ou TextField (usar default vazio).
- NUNCA realizar queries dentro de loops (usar `select_related` / `prefetch_related`).
- NUNCA colocar chaves secretas ou credenciais no código (usar `django-environ`).
- NUNCA expor o `id` sequencial do banco em URLs públicas se a segurança for crítica (usar UUID).
