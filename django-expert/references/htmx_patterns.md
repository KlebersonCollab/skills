# Django + HTMX Patterns

HTMX allows creating rich interfaces using only HTML, keeping the state on the server.

## 1. Returning Fragments (Partials)

In Django, you must configure your views to detect HTMX requests and return only the necessary HTML fragment.

```python
def item_list(request):
    items = Item.objects.all()
    if request.headers.get('HX-Request'):
        return render(request, 'partials/item_rows.html', {'items': items})
    return render(request, 'items.html', {'items': items})
```

## 2. Common Interaction Patterns

### Inline Validation
Validate form fields while the user types.
```html
<input name="email" hx-post="/validate-email/" hx-trigger="blur">
```

### Infinite Scroll
```html
<div hx-get="/items/?page=2" hx-trigger="revealed" hx-swap="afterend">
    ...
</div>
```

## 3. Useful Middleware
It is recommended to use `django-htmx` to facilitate header manipulation.
- `request.htmx`: Checks if the request came from HTMX.
- `HX-Trigger`: Trigger events on the frontend from the backend.


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.371439Z"
evidence_checksum: "NONE"
```
