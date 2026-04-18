# Django Forms & Validation

Formulários são a primeira linha de defesa contra dados inválidos ou maliciosos.

## 1. Clean Methods

Sempre use métodos `clean_<field>` para validações específicas e o método `clean()` para validações cruzadas.

```python
class RegistrationForm(forms.ModelForm):
    def clean_email(self):
        email = self.cleaned_data.get('email')
        if User.objects.filter(email=email).exists():
            raise forms.ValidationError("Este e-mail já está em uso.")
        return email

    def clean(self):
        cleaned_data = super().clean()
        password = cleaned_data.get("password")
        confirm = cleaned_data.get("confirm_password")
        if password != confirm:
            raise forms.ValidationError("Senhas não conferem.")
```

## 2. Widgets e CSS
Não coloque lógica de estilo no HTML. Use widgets para gerenciar classes CSS.

```python
widgets = {
    'name': forms.TextInput(attrs={'class': 'form-control', 'placeholder': 'Nome'}),
}
```

## 3. Best Practices
- **ModelForms**: Use sempre que o formulário mapear diretamente para um modelo.
- **CSRF**: Nunca esqueça a tag `{% csrf_token %}` no template.
- **Error Messages**: Personalize mensagens de erro para melhorar a UX.
