# Django Background Tasks (Celery)

Heavy tasks (email, image processing, external integrations) must be executed outside the request-response cycle.

## 1. Setup Celery with Redis

Installation via UV:
```bash
uv add celery redis
```

Recommended structure:
- `core/celery.py`: Celery app configuration.
- `apps/<app>/tasks.py`: Task definitions.

## 2. retry_backoff Pattern

Always use retry for tasks that depend on external services.

```python
@shared_task(bind=True, max_retries=3)
def send_external_notification(self, user_id):
    try:
        # Task logic
        ...
    except Exception as exc:
        raise self.retry(exc=exc, countdown=60 * 5) # Retry in 5 min
```

## 3. Best Practices
- **Atomic Tasks**: Ensure tasks are idempotent (they can run twice without issues).
- **Avoid Large Payloads**: Pass only the object ID, not the entire object or large strings.
- **Monitoring**: Use **Flower** to monitor queues in real-time.


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.371012Z"
evidence_checksum: "NONE"
```
