# Cross-Cutting Concerns

## Caching
Protect expensive resources.

- **View Cache**: `@method_decorator(cache_page(60 * 15))`.
- **Fragment Cache**: `{% cache 500 sidebar %}`.
- **Low-level Cache**: `cache.get_or_set('key', expensive_func)`.

## Signals
Use signals to decouple side effects (e.g., creating a Profile when creating a User).
**Mandate**: Import signals in the `ready()` method of `AppConfig`.

```python
# apps/users/apps.py
def ready(self):
    import apps.users.signals
```

## Middleware
Use for global Request/Response processing (Logs, Auth, Headers).

```python
class RequestLoggingMiddleware:
    def __init__(self, get_response):
        self.get_response = get_response

    def __call__(self, request):
        # Pre-processing
        response = self.get_response(request)
        # Post-processing
        return response
```


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.371122Z"
evidence_checksum: "NONE"
```
