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

## 3. QuerySet Managers

Mantenha as queries complexas dentro de `Managers` customizados para garantir DRY.

```python
class ActiveUserManager(models.Manager):
    def get_queryset(self):
        return super().get_queryset().filter(is_active=True)

class User(models.Model):
    # ...
    objects = models.Manager()
    active_objects = ActiveUserManager()
```

## 4. Bulk Operations
Sempre use `bulk_create`, `bulk_update` e `QuerySet.update()` para manipulação de grandes volumes de dados.
