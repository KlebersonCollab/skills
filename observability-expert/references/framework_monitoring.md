# Monitoring Frameworks (FastAPI & Django)

For modern APIs, monitoring just the server is not enough. We need visibility into framework and application behavior.

## 1. FastAPI Monitoring (Prometheus)

Use `prometheus-fastapi-instrumentator` for automatic RED metrics exposure.

### Implementation
```python
from prometheus_fastapi_instrumentator import Instrumentator

app = FastAPI()

@app.on_event("startup")
async def startup():
    Instrumentator().instrument(app).expose(app)
```

### Recommended Dashboard (Grafana)
- **ID 16101**: Full FastAPI dashboard with error rates per route and p99 latency.

## 2. Django Monitoring

Use `django-prometheus` to capture Middleware, ORM, and Database metrics.

### Configuration
```python
# settings.py
INSTALLED_APPS = [
    ...
    'django_prometheus',
]

MIDDLEWARE = [
    'django_prometheus.middleware.PrometheusBeforeMiddleware',
    ...
    'django_prometheus.middleware.PrometheusAfterMiddleware',
]
```

### Recommended Dashboard (Grafana)
- **ID 9528**: Django request monitoring and SQL query response times.

## 3. Custom Business Metrics

Do not just monitor technical errors. Monitor business success.

```python
from prometheus_client import Counter

orders_created = Counter('orders_total', 'Total created orders', ['payment_method'])

def create_order(method):
    # ... logic ...
    orders_created.labels(payment_method=method).inc()
```

## 4. Dashboard Checklist
- [ ] **Latency (p95, p99)**: How long is the user waiting?
- [ ] **Error Rate (5xx vs 4xx)**: Is the system broken or is the user making a mistake?
- [ ] **Request Rate (RPS)**: Are we under attack or at peak traffic?
- [ ] **Database Connections**: Is the pool saturated?


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.376475Z"
evidence_checksum: "NONE"
```
