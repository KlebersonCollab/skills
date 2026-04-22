# Django ORM Performance Reference

O ORM do Django é poderoso, mas pode ser uma armadilha de performance se não for usado com rigor.

## 1. O Problema do N+1

Este é o erro mais comum. Ocorre quando você itera sobre um queryset e acessa uma relação que não foi carregada previamente.

### Soluções:
- **`select_related(*fields)`**: Faz um `JOIN` SQL. Use para `ForeignKey` e `OneToOneField`.
- **`prefetch_related(*fields)`**: Faz uma query separada e faz o "join" em Python. Use para `ManyToManyField` e relações reversas.

```python
# Exemplo de uso em View
queryset = Order.objects.select_related('customer').prefetch_related('items__product')
```

## 2. Somente o que você precisa

NUNCA carregue o objeto inteiro se precisar apenas de alguns campos.

- **`only(*fields)`**: Carrega apenas os campos especificados.
- **`defer(*fields)`**: Carrega todos exceto os especificados.
- **`values()` / `values_list()`**: Retorna dicionários ou tuplas (muito mais rápido se não precisar dos métodos do modelo).

## 3. QuerySet & Managers (Advanced)

Utilize `QuerySet.as_manager()` para encapsular lógica de filtro reutilizável.

```python
class ProductQuerySet(models.QuerySet):
    def active(self):
        return self.filter(is_active=True)

    def with_category(self):
        return self.select_related('category')

class Product(models.Model):
    # ...
    objects = ProductQuerySet.as_manager()
```

### Manager Helper Methods
Adicione métodos utilitários para evitar repetição de `try/except`.
```python
class BaseManager(models.Manager):
    def get_or_none(self, **kwargs):
        try:
            return self.get(**kwargs)
        except self.model.DoesNotExist:
            return None
```

## 4. Database Hardening
Use constraints e índices para garantir integridade e performance no banco.

```python
class Meta:
    indexes = [
        models.Index(fields=['slug']),
        models.Index(fields=['category', 'is_active']),
    ]
    constraints = [
        models.CheckConstraint(
            check=models.Q(price__gte=0),
            name='price_non_negative'
        )
    ]
```

## 5. Bulk Operations
Sempre use `bulk_create`, `bulk_update` e `QuerySet.update()` para manipulação de grandes volumes de dados.

