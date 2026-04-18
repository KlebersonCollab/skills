# Django Background Tasks (Celery)

Tarefas pesadas (e-mail, processamento de imagens, integrações externas) devem ser executadas fora do ciclo request-response.

## 1. Setup Celery com Redis

Instalação via UV:
```bash
uv add celery redis
```

Estrutura recomendada:
- `core/celery.py`: Configuração do app Celery.
- `apps/<app>/tasks.py`: Definição das tarefas.

## 2. retry_backoff Pattern

Sempre utilize retry para tarefas que dependem de serviços externos.

```python
@shared_task(bind=True, max_retries=3)
def send_external_notification(self, user_id):
    try:
        # Lógica da tarefa
        ...
    except Exception as exc:
        raise self.retry(exc=exc, countdown=60 * 5) # Retry em 5 min
```

## 3. Best Practices
- **Atomic Tasks**: Garanta que as tarefas sejam idempotentes (podem rodar 2x sem problemas).
- **Avoid Large Payloads**: Passe apenas o ID do objeto, não o objeto inteiro ou grandes strings.
- **Monitoring**: Use o **Flower** para monitorar as filas em tempo real.
