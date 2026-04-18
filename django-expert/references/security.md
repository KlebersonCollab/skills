# Django Security Reference

Django é "secure by default", mas configurações incorretas em produção podem expor sua aplicação.

## 1. Deployment Checklist

Sempre rode o comando de verificação antes de qualquer deploy:
```bash
uv run python manage.py check --deploy
```

## 2. Configurações Críticas (`settings.py`)

Em produção, garanta que estas variáveis estejam configuradas corretamente (usando `django-environ`):

- `DEBUG = False`
- `ALLOWED_HOSTS = ['seu-dominio.com']`
- `SECRET_KEY` (Deve ser forte e vir de variável de ambiente)
- `SECURE_SSL_REDIRECT = True`
- `SESSION_COOKIE_SECURE = True`
- `CSRF_COOKIE_SECURE = True`

## 3. Proteção CSRF e CORS
Ao usar HTMX ou APIs, certifique-se de configurar o `CSRF_TRUSTED_ORIGINS`.

```python
CSRF_TRUSTED_ORIGINS = ['https://seu-dominio.com']
```

## 4. Database Security
- Use sempre variáveis de ambiente para a `DATABASE_URL`.
- Garanta que o usuário do banco tenha permissões limitadas (mínimo privilégio).
