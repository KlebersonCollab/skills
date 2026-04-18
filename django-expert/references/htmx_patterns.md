# Django + HTMX Patterns

O HTMX permite criar interfaces ricas usando apenas HTML, mantendo o estado no servidor.

## 1. Retornando Fragmentos (Partials)

No Django, você deve configurar suas views para detectar requisições HTMX e retornar apenas o pedaço do HTML necessário.

```python
def item_list(request):
    items = Item.objects.all()
    if request.headers.get('HX-Request'):
        return render(request, 'partials/item_rows.html', {'items': items})
    return render(request, 'items.html', {'items': items})
```

## 2. Padrões de Interação Common

### Inline Validation
Validar campos de formulário enquanto o usuário digita.
```html
<input name="email" hx-post="/validate-email/" hx-trigger="blur">
```

### Infinite Scroll
```html
<div hx-get="/items/?page=2" hx-trigger="revealed" hx-swap="afterend">
    ...
</div>
```

## 3. Middleware Úteis
Recomenda-se o uso do `django-htmx` para facilitar a manipulação de headers.
- `request.htmx`: Verifica se a request veio do HTMX.
- `HX-Trigger`: Disparar eventos no frontend a partir do backend.
