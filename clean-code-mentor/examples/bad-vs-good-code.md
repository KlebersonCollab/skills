# Examples: Bad vs. Good Code Design

Comparações práticas de aplicação de princípios Clean Code e SOLID.

---

## 1. Single Responsibility Principle (SRP)

### ❌ Ruim (Classe God)
```dart
class UserProfile {
  void saveToDatabase() { /* ... */ }
  void uploadAvatar() { /* ... */ }
  void validateEmail() { /* ... */ }
  void renderProfileWidget() { /* ... */ }
}
```

### ✅ Bom (Especialização)
```dart
class UserRepository { void save(User user) { /* ... */ } }
class AvatarService { void upload(File file) { /* ... */ } }
class UserValidator { bool isValid(User user) { /* ... */ } }
class ProfileScreen extends StatelessWidget { /* UI only */ }
```

---

## 2. DRY vs. Duplicação de Conhecimento

### ❌ Ruim (Lógica Duplicada)
```python
# em auth.py
if "@" not in email: raise Error("Email inválido")

# em newsletter.py
if "@" not in email: raise Error("Email inválido")
```

### ✅ Bom (Centralização)
```python
# em utils/validators.py
def validate_email(email):
    if "@" not in email: raise Error("Email inválido")

# Reuso
from utils.validators import validate_email
validate_email(email)
```

---

## 3. KISS & YAGNI (Abstração Excessiva)

### ❌ Ruim (Over-engineering)
```typescript
interface IStringConcatenator {
  concatenate(a: string, b: string): string;
}

class StandardConcatenator implements IStringConcatenator {
  concatenate(a: string, b: string): string { return a + b; }
}

// Uso:
const concat = new StandardConcatenator();
console.log(concat.concatenate("Hello", "World"));
```

### ✅ Bom (Direto ao ponto)
```typescript
const fullName = `${firstName} ${lastName}`;
```

---

## 4. Dependency Inversion Principle (DIP)

### ❌ Ruim (Dependência Concreta)
```dart
class OrderService {
  final SqlDatabase db = SqlDatabase(); // Acoplamento forte
  
  void saveOrder(Order order) {
    db.insert(order);
  }
}
```

### ✅ Bom (Injeção de Abstração)
```dart
abstract class IDatabase { void insert(dynamic data); }

class OrderService {
  final IDatabase db;
  OrderService(this.db); // Recebe qualquer banco que implemente IDatabase
  
  void saveOrder(Order order) {
    db.insert(order);
  }
}
```
