# Django ORM Performance Reference

The Django ORM is powerful, but it can be a performance trap if not used strictly.

## 1. The N+1 Problem

This is the most common error. It occurs when you iterate over a queryset and access a relationship that was not previously loaded.

### Solutions:
- **`select_related(*fields)`**: Does a SQL `JOIN`. Use for `ForeignKey` and `OneToOneField`.
- **`prefetch_related(*fields)`**: Does a separate query and performs the "join" in Python. Use for `ManyToManyField` and reverse relationships.

```python
# Example usage in View
queryset = Order.objects.select_related('customer').prefetch_related('items__product')
```

## 2. Only what you need

NEVER load the entire object if you only need a few fields.

- **`only(*fields)`**: Loads only the specified fields.
- **`defer(*fields)`**: Loads all except those specified.
- **`values()` / `values_list()`**: Returns dictionaries or tuples (much faster if you don't need the model methods).

## 3. QuerySet & Managers (Advanced)

Use `QuerySet.as_manager()` to encapsulate reusable filter logic.

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
Add utility methods to avoid repeating `try/except`.
```python
class BaseManager(models.Manager):
    def get_or_none(self, **kwargs):
        try:
            return self.get(**kwargs)
        except self.model.DoesNotExist:
            return None
```

## 4. Database Hardening
Use constraints and indexes to ensure database integrity and performance.

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
Always use `bulk_create`, `bulk_update`, and `QuerySet.update()` for handling large volumes of data.
