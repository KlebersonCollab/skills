# Django Testing Reference

O Django possui um ecossistema de testes robusto. Na `django-expert`, padronizamos o uso do `pytest`.

## 1. Setup com Pytest-Django

Evite o `manage.py test` para testes complexos. O `pytest` oferece fixtures muito mais flexíveis.

```bash
uv add --dev pytest-django factory-boy pytest-cov
```

No arquivo `pytest.ini`:
```ini
[pytest]
DJANGO_SETTINGS_MODULE = core.settings
python_files = tests.py test_*.py
```

## 2. Factories (Factory Boy)

NUNCA use fixtures em JSON/YAML. Use Factories para gerar dados dinâmicos e legíveis.

```python
import factory
from apps.users.models import User

class UserFactory(factory.django.DjangoModelFactory):
    class Meta:
        model = User

    username = factory.Faker('user_name')
    email = factory.Faker('email')
```

## 3. Testando Views e HTMX

Use o `RequestFactory` ou o `Client` do Django para simular requisições.

```python
@pytest.mark.django_db
def test_htmx_partial_return(client):
    response = client.get('/items/', HTTP_HX_REQUEST='true')
    assert response.status_code == 200
    assert 'partials/item_rows.html' in [t.name for t in response.templates]
```
