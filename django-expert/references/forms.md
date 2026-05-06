# Django Forms & Validation

Forms are the first line of defense against invalid or malicious data.

## 1. Clean Methods

Always use `clean_<field>` methods for specific validations and the `clean()` method for cross-validations.

```python
class RegistrationForm(forms.ModelForm):
    def clean_email(self):
        email = self.cleaned_data.get('email')
        if User.objects.filter(email=email).exists():
            raise forms.ValidationError("This email is already in use.")
        return email

    def clean(self):
        cleaned_data = super().clean()
        password = cleaned_data.get("password")
        confirm = cleaned_data.get("confirm_password")
        if password != confirm:
            raise forms.ValidationError("Passwords do not match.")
```

## 2. Widgets and CSS
Do not put styling logic in HTML. Use widgets to manage CSS classes.

```python
widgets = {
    'name': forms.TextInput(attrs={'class': 'form-control', 'placeholder': 'Name'}),
}
```

## 3. Best Practices
- **ModelForms**: Use whenever the form maps directly to a model.
- **CSRF**: Never forget the `{% csrf_token %}` tag in the template.
- **Error Messages**: Customize error messages to improve UX.


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.371335Z"
evidence_checksum: "NONE"
```
