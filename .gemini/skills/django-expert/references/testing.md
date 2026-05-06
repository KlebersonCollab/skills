# Django Testing Excellence Reference

## 1. Fast Infrastructure
Reduce developer feedback time.

### Pytest Configuration
```ini
# pytest.ini
[pytest]
DJANGO_SETTINGS_MODULE = config.settings.test
addopts = --reuse-db --nomigrations --cov=apps --cov-report=term-missing
```

### Test Settings (In-Memory)
```python
# config/settings/test.py
DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.sqlite3',
        'NAME': ':memory:',
    }
}
PASSWORD_HASHERS = ['django.contrib.auth.hashers.MD5PasswordHasher']
```

## 2. Advanced Factories (Factory Boy)
Eliminate manual data creation.

```python
class ProductFactory(factory.django.DjangoModelFactory):
    class Meta:
        model = Product

    name = factory.Faker('word')
    category = factory.SubFactory(CategoryFactory)
    
    @factory.post_generation
    def tags(self, create, extracted, **kwargs):
        if not create: return
        if extracted:
            for tag in extracted: self.tags.add(tag)
```

## 3. Mocking & External Services
Isolate your application from the external world.

```python
@patch('apps.payments.services.stripe.Charge.create')
def test_successful_payment(self, mock_stripe, client):
    mock_stripe.return_value = {'id': 'ch_123', 'status': 'succeeded'}
    # ... executes view ...
    mock_stripe.assert_called_once()
```

## 4. DRF & API Testing
```python
@pytest.fixture
def authenticated_api_client(user):
    from rest_framework.test import APIClient
    client = APIClient()
    client.force_authenticate(user=user)
    return client

def test_create_product(authenticated_api_client):
    response = authenticated_api_client.post('/api/products/', {'name': 'New'})
    assert response.status_code == 201
```

## 5. Coverage Goals
Keep quality visible.

| Layer | Coverage Goal |
|--------|-------------------|
| Models | 90%+ |
| Services | 90%+ |
| Serializers | 85%+ |
| Views | 80%+ |


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.371812Z"
evidence_checksum: "NONE"
```
