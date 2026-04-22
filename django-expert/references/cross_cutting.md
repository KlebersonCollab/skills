# Cross-Cutting Concerns

## Caching
Proteja recursos caros.

- **View Cache**: `@method_decorator(cache_page(60 * 15))`.
- **Fragment Cache**: `{% cache 500 sidebar %}`.
- **Low-level Cache**: `cache.get_or_set('key', expensive_func)`.

## Signals
Use sinais para desacoplar efeitos colaterais (ex: criar Perfil ao criar Usuário).
**Mandato**: Importe os sinais no `ready()` do `AppConfig`.

```python
# apps/users/apps.py
def ready(self):
    import apps.users.signals
```

## Middleware
Utilize para processamento global de Request/Response (Logs, Auth, Headers).

```python
class RequestLoggingMiddleware:
    def __init__(self, get_response):
        self.get_response = get_response

    def __call__(self, request):
        # Pré-processamento
        response = self.get_response(request)
        # Pós-processamento
        return response
```
