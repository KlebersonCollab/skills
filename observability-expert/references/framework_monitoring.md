# Monitoring Frameworks (FastAPI & Django)

Para APIs modernas, monitorar apenas o servidor não é suficiente. Precisamos de visibilidade sobre o comportamento do framework e da aplicação.

## 1. FastAPI Monitoring (Prometheus)

Utilize o `prometheus-fastapi-instrumentator` para exposição automática de métricas RED.

### Implementação
```python
from prometheus_fastapi_instrumentator import Instrumentator

app = FastAPI()

@app.on_event("startup")
async def startup():
    Instrumentator().instrument(app).expose(app)
```

### Dashboard Recomendado (Grafana)
- **ID 16101**: Dashboard completo para FastAPI com taxas de erro por rota e latência p99.

## 2. Django Monitoring

Utilize o `django-prometheus` para capturar métricas de Middleware, ORM e Banco de Dados.

### Configuração
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

### Dashboard Recomendado (Grafana)
- **ID 9528**: Monitoramento de requisições Django e tempos de resposta de query SQL.

## 3. Métricas Customizadas de Negócio

Não monitore apenas erros técnicos. Monitore o sucesso do negócio.

```python
from prometheus_client import Counter

orders_created = Counter('orders_total', 'Total created orders', ['payment_method'])

def create_order(method):
    # ... lógica ...
    orders_created.labels(payment_method=method).inc()
```

## 4. Checklist de Dashboards
- [ ] **Latência (p95, p99)**: Quanto tempo o usuário está esperando?
- [ ] **Taxa de Erros (5xx vs 4xx)**: O sistema está quebrado ou o usuário está errando?
- [ ] **Taxa de Requisições (RPS)**: Estamos sob ataque ou em pico de tráfego?
- [ ] **Conexões de Banco de Dados**: O pool está saturado?
