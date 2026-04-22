# Django Security Hardening Reference

## 1. Authentication & Passwords
Fortaleça a entrada do sistema.

### Password Validators
```python
AUTH_PASSWORD_VALIDATORS = [
    {'NAME': 'django.contrib.auth.password_validation.UserAttributeSimilarityValidator'},
    {'NAME': 'django.contrib.auth.password_validation.MinimumLengthValidator', 'OPTIONS': {'min_length': 12}},
    {'NAME': 'django.contrib.auth.password_validation.CommonPasswordValidator'},
    {'NAME': 'django.contrib.auth.password_validation.NumericPasswordValidator'},
]
```

### Password Hashing (Argon2)
Use Argon2 para proteção superior contra força bruta.
```python
PASSWORD_HASHERS = [
    'django.contrib.auth.hashers.Argon2PasswordHasher',
    'django.contrib.auth.hashers.PBKDF2PasswordHasher',
]
```

## 2. Authorization (RBAC)
Mantenha a lógica de permissões organizada.

### Model-Level Logic
```python
class Post(models.Model):
    def user_can_edit(self, user):
        return self.author == user or user.is_staff
```

### Admin Required Mixin
```python
class AdminRequiredMixin:
    def dispatch(self, request, *args, **kwargs):
        if not request.user.is_authenticated or not request.user.is_staff:
            raise PermissionDenied
        return super().dispatch(request, *args, **kwargs)
```

## 3. File Security
Valide rigorosamente o que o usuário envia.

```python
def validate_file_size(value):
    if value.size > 5 * 1024 * 1024: # 5MB
        raise ValidationError('File too large.')

class Document(models.Model):
    file = models.FileField(validators=[validate_file_size])
```

## 4. API Throttling (DRF)
Proteja seus recursos de abuso.

```python
REST_FRAMEWORK = {
    'DEFAULT_THROTTLE_CLASSES': [
        'rest_framework.throttling.AnonRateThrottle',
        'rest_framework.throttling.UserRateThrottle'
    ],
    'DEFAULT_THROTTLE_RATES': {
        'anon': '100/day',
        'user': '1000/day',
    }
}
```

## 5. Security Headers & CSP
```python
# Content Security Policy (Middleware customizado ou django-csp)
CSP_DEFAULT_SRC = ("'self'",)
CSP_STYLE_SRC = ("'self'", "'unsafe-inline'")

# Headers de Proteção
SECURE_CONTENT_TYPE_NOSNIFF = True
SECURE_BROWSER_XSS_FILTER = True
X_FRAME_OPTIONS = 'DENY'
```

## 6. SQL Injection Prevention
- **Regra de Ouro**: Use o ORM sempre.
- **Raw SQL**: Se for usar `.raw()`, NUNCA use f-strings ou concatenação. Use parâmetros.
```python
# DO THIS
User.objects.raw('SELECT * FROM users WHERE email = %s', [email])

# NEVER THIS
User.objects.raw(f'SELECT * FROM users WHERE email = {email}')
```
