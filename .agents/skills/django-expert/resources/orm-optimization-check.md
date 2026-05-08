# Checklist: Django ORM Optimization

Use this checklist to identify and fix performance issues in your Django queries.

## 1. N+1 Problems
- [ ] **Identification**: Does the number of queries increase proportionally to the number of items in the list?
- [ ] **Solution (ForeignKey/OneToOne)**: Use `.select_related('field_name')` to perform a SQL JOIN.
- [ ] **Solution (ManyToMany/Reverse ForeignKey)**: Use `.prefetch_related('field_name')` to perform separate queries and join in Python.

## 2. Field Selection (Payload)
- [ ] **Only necessary fields**: Use `.only('field1', 'field2')` to load only specific columns.
- [ ] **Exclude heavy fields**: Use `.defer('heavy_field')` for fields like `TextField` or `BinaryField` that are not needed immediately.
- [ ] **Dictionaries/Tuples**: Use `.values()` or `.values_list()` when you don't need model instances (much faster).

## 3. Aggregations and Calculations
- [ ] **Database Calculation**: Use `.annotate()` with aggregation functions (`Sum`, `Count`, `Avg`) instead of calculating in Python loops.
- [ ] **Simple Count**: Use `.exists()` to check for presence and `.count()` instead of `len(queryset)`.

## 4. Lazy Querysets
- [ ] **Avoid premature evaluation**: Querysets are lazy. Do not convert them to lists (`list(qs)`) until strictly necessary.
- [ ] **Efficient Iteration**: Use `.iterator()` to process large volumes of data without loading everything into memory.

## 5. Debugging Tools
- [ ] **django-debug-toolbar**: Essential for seeing generated SQL in real-time in the browser.
- [ ] **nplusone**: Library to automatically detect N+1 issues during development.
- [ ] **Query Inspector**: Log SQL to the console during testing or local development:
  ```python
  import logging
  l = logging.getLogger('django.db.backends')
  l.setLevel(logging.DEBUG)
  l.addHandler(logging.StreamHandler())
  ```

## Golden Rule
> "Move data logic to the database and presentation logic to Python."

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T14:05:00Z"
evidence_checksum: "8e52f6a"
```
