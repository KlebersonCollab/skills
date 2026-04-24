# Production Architecture & Layout

## Recommended Layout
Separate global configuration from local applications.

```
myproject/
├── config/             # Global configurations
│   ├── settings/       # Split Settings
│   │   ├── base.py
│   │   ├── development.py
│   │   └── production.py
│   ├── urls.py
│   └── wsgi.py
├── apps/               # All local apps here
│   ├── users/
│   └── products/
└── manage.py
```

## Split Settings Pattern
Avoid giant conditionals in `settings.py`.

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
Always register local apps using the full path if they are in `apps/`.
```python
INSTALLED_APPS = [
    # ...
    'apps.users.apps.UsersConfig',
]
```
