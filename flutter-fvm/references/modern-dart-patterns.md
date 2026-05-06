# Modern Dart Patterns (v3.0+)

Dart 3 has transformed Flutter from a purely imperative framework to a model rich in types and functional patterns.

## 1. Sealed Classes (Unions)
**Concept**: Ensures the compiler knows all possible subtypes of a class, allowing exhaustiveness in switches.

### ❌ Bad (Boolean Flags)
```dart
class UserState {
  final bool isLoading;
  final String? error;
  final User? data;
  // Allows impossible states like isLoading=true AND data!=null
}
```

### ✅ Good (Sealed)
```dart
sealed class UserState {}
class UserInitial extends UserState {}
class UserLoading extends UserState {}
class UserSuccess extends UserState { final User user; UserSuccess(this.user); }
class UserError extends UserState { final String msg; UserError(this.msg); }
```

## 2. Pattern Matching
**Concept**: Data extraction and flow control based on the object's structure.

### ✅ Good (Switch Expression)
```dart
final message = switch (state) {
  UserInitial() => 'Welcome',
  UserLoading() => 'Loading...',
  UserSuccess(user: var u) => 'Hello, ${u.name}',
  UserError(msg: var m) => 'Error: $m',
};
```

## 3. Records
**Concept**: Anonymous types for multiple returns without the need for heavy classes.

### ✅ Good
```dart
(double, double) getCoordinates() => (-23.5, -46.6);

// Destructuring
final (lat, lng) = getCoordinates();
```

## Modern Dart Checklist
- [ ] Used `sealed` for UI states?
- [ ] Is the `switch` an expression (with return) whenever possible?
- [ ] Avoided `dynamic` by enabling `strict-casts` in the analyzer?
- [ ] Used `records` for simple multiple returns?


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.367171Z"
evidence_checksum: "NONE"
```
