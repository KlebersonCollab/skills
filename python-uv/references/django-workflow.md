# Django Professional Workflow with UV — Referência Completa

Guia de padrões, segurança, performance e arquitetura avançada para aplicações Django de alta performance usando o ecossistema `uv`.

---

## 1. Setup e Gerenciamento Otimizado

O `uv` elimina a lentidão histórica do gerenciamento de dependências no Django e facilita a execução de comandos administrativos.

### Inicialização Profissional
```bash
uv init meu-projeto-django
cd meu-projeto-django
uv add django djangorestframework django-environ django-filter
uv add --dev pytest-django ruff bandit django-debug-toolbar django-extensions

# Iniciar projeto Django
uv run django-admin startproject config .
```

### Comandos de Produtividade (sem ativar venv)
| Ação | Comando UV |
|------|------------|
| Servidor de Dev | `uv run python manage.py runserver` |
| Servidor ASGI (Async) | `uv run uvicorn config.asgi:application --reload` |
| Shell Plus (IPython) | `uv run python manage.py shell_plus` |
| Listar URLs | `uv run python manage.py show_urls` |
| Testes com Coverage | `uv run pytest --cov` |

---

## 2. Performance e Otimização do ORM

Evite gargalos comuns como o problema de N+1 queries e exaustão de memória.

### Estratégias contra N+1
Sempre utilize carregamento antecipado (Eager Loading) em views e serializers.

- **`select_related(*fields)`**: Para chaves estrangeiras (ForeignKey) e relações 1-para-1 (JOIN SQL).
- **`prefetch_related(*fields)`**: Para relações muitos-para-muitos (ManyToMany) ou relações reversas (Queries separadas com IN).

```python
# ✅ Padrão de Performance em ViewSets
class OrderViewSet(viewsets.ReadOnlyModelViewSet):
    queryset = Order.objects.select_related('user', 'shipping_address') \
                            .prefetch_related('items__product')
```

### Operações em Lote (Bulk)
Evite salvar objetos um a um dentro de loops.

```python
# ❌ Lento
for item in items:
    Product.objects.create(**item)

# ✅ Rápido (Bulk Create)
Product.objects.bulk_create([Product(**item) for item in items])

# ✅ Bulk Update
Product.objects.filter(category='old').update(status='retired')
```

### Processamento de Grandes Volumes
Use `.iterator()` para evitar carregar milhares de objetos na memória RAM.

```python
# ✅ Seguro para memória (processa em chunks)
for user in User.objects.all().iterator(chunk_size=1000):
    send_email(user)
```

---

## 3. Profiling e Diagnóstico via UV

Use ferramentas de inspeção para identificar queries lentas durante o desenvolvimento.

### Django Debug Toolbar
Configure no `settings.py` e use o `uv` para garantir que as dependências de dev estão presentes.
```bash
# Rodar servidor com debug toolbar ativo (se configurado)
uv run python manage.py runserver
```

### Analisando Queries no Shell
```python
# uv run python manage.py shell_plus --print-sql
from django.db import connection
print(len(connection.queries)) # Ver contagem de queries
```

---

## 4. Segurança e Revisão de Acesso (Audit)

Baseado nos padrões de **Access Review**, proteja seu código contra falhas como **BOLA/IDOR**.

### Busca de Padrões Inseguros
```bash
# Buscar views que podem estar expondo dados sem filtro de dono (BOLA)
uv run grep -rn ".objects.all()" . --include="views.py"
```

### Padrão de Queryset Seguro
```python
# views.py (DRF)
class ProjectViewSet(viewsets.ModelViewSet):
    def get_queryset(self):
        # ✅ Padrão Seguro: Escopo restrito ao usuário logado
        return Project.objects.filter(owner=self.request.user)
```

---

## 5. Arquitetura Avançada (Django Pro)

### Service Layer (Lógica de Negócio)
Evite "Fat Models" ou "Fat Views". Coloque a lógica complexa em serviços.

```python
# services.py
def process_order_payment(order_id, payment_data):
    # Lógica complexa aqui
    pass
```

### Async Django (ASGI)
Para operações I/O bound (chamadas de API externas, WebSockets), use views assíncronas.

```python
async def my_async_view(request):
    async with httpx.AsyncClient() as client:
        response = await client.get("https://api.external.com")
    return JsonResponse(response.json())
```

---

## 6. Checklist de Performance e Qualidade

- [ ] **N+1**: Verificou se todas as chaves estrangeiras em listas usam `select_related`?
- [ ] **Indexes**: Campos usados em `filter()` ou `order_by()` possuem `db_index=True`?
- [ ] **Pagination**: Todos os endpoints de listagem possuem paginação obrigatória?
- [ ] **Bulk**: Operações de criação/atualização múltipla usam métodos `bulk_`?
- [ ] **Security**: O `get_queryset` filtra os dados pelo usuário/tenant logado?
- [ ] **Async**: Operações de rede externas estão usando views `async` e `httpx`?
