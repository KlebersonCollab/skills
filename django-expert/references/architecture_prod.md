# Production Architecture & Layout

## Recommended Layout
Separe a configuração dos aplicativos locais.

```
myproject/
├── config/             # Configurações globais
│   ├── settings/       # Split Settings
│   │   ├── base.py
│   │   ├── development.py
│   │   └── production.py
│   ├── urls.py
│   └── wsgi.py
├── apps/               # Todos os apps locais aqui
│   ├── users/
│   └── products/
└── manage.py
```

## Split Settings Pattern
Evite condicionais gigantes no `settings.py`.

```python
# config/settings/production.py
from .base import *

DEBUG = False
ALLOWED_HOSTS = env.list('ALLOWED_HOSTS')

# Security Headers
SECURE_SSL_REDIRECT = True
SESSION_COOKIE_SECURE = True
CSRF_COOKIE_SECURE = True
```

## App Registration
Sempre registre apps locais usando o caminho completo se estiverem em `apps/`.
```python
INSTALLED_APPS = [
    # ...
    'apps.users.apps.UsersConfig',
]
```
