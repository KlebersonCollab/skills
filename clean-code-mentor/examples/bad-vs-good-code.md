# Examples: Bad vs. Good Code Design

Practical comparisons of Clean Code and SOLID principles application.

---

## 1. Single Responsibility Principle (SRP)

### ❌ Bad (God Class)
```dart
class UserProfile {
  void saveToDatabase() { /* ... */ }
  void uploadAvatar() { /* ... */ }
  void validateEmail() { /* ... */ }
  void renderProfileWidget() { /* ... */ }
}
```

### ✅ Good (Specialization)
```dart
class UserRepository { void save(User user) { /* ... */ } }
class AvatarService { void upload(File file) { /* ... */ } }
class UserValidator { bool isValid(User user) { /* ... */ } }
class ProfileScreen extends StatelessWidget { /* UI only */ }
```

---

## 2. DRY vs. Knowledge Duplication

### ❌ Bad (Duplicated Logic)
```python
# in auth.py
if "@" not in email: raise Error("Invalid email")

# in newsletter.py
if "@" not in email: raise Error("Invalid email")
```

### ✅ Good (Centralization)
```python
# in utils/validators.py
def validate_email(email):
    if "@" not in email: raise Error("Invalid email")

# Reuse
from utils.validators import validate_email
validate_email(email)
```

---

## 3. KISS & YAGNI (Excessive Abstraction)

### ❌ Bad (Over-engineering)
```typescript
interface IStringConcatenator {
  concatenate(a: string, b: string): string;
}

class StandardConcatenator implements IStringConcatenator {
  concatenate(a: string, b: string): string { return a + b; }
}

// Usage:
const concat = new StandardConcatenator();
console.log(concat.concatenate("Hello", "World"));
```

### ✅ Good (Straight to the point)
```typescript
const fullName = `${firstName} ${lastName}`;
```

---

## 4. Dependency Inversion Principle (DIP)

### ❌ Bad (Concrete Dependency)
```dart
class OrderService {
  final SqlDatabase db = SqlDatabase(); // Tight coupling
  
  void saveOrder(Order order) {
    db.insert(order);
  }
}
```

### ✅ Good (Abstraction Injection)
```dart
abstract class IDatabase { void insert(dynamic data); }

class OrderService {
  final IDatabase db;
  OrderService(this.db); // Receives any database that implements IDatabase
  
  void saveOrder(Order order) {
    db.insert(order);
  }
}
```
