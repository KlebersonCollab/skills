# Project Structure: Flutter with FVM (2026 Edition)

Organization by features (**Feature-First**) isolates scope and facilitates maintenance for large teams.

## 1. Directory Structure
```text
lib/
├── core/                # Transversal code (di, network, theme, utils)
├── features/            # Business features
│   ├── [feature_name]/  
│   │   ├── data/        # Repositories, API Sources, Models
│   │   ├── domain/      # Sealed States, Entities, Use Cases
│   │   └── presentation/# ViewModels, Views, Widgets
├── main_dev.dart        # Entry points per flavor
└── main_prod.dart
```

## 2. Architectural Pattern: MVVM + Dart 3
Dart 3 allows the ViewModel to expose immutable and exhaustive states.

### ❌ Bad (Scattered Logic)
- Widgets making direct API calls.
- Using `if (loading) ... else if (error) ...`.

### ✅ Good (Reactive and Typed)
- **ViewModel**: Uses `AsyncNotifier` (Riverpod) or `Bloc` and emits a `sealed class`.
- **View**: Uses `switch (state)` to render the UI.

## 3. Secrets Management via Flavors
Use `--dart-define-from-file` to load JSON configurations.

```bash
fvm flutter run --flavor dev --dart-define-from-file=env/dev.json
```

## Structure Checklist
- [ ] Is the feature isolated (does not import private widgets from others)?
- [ ] Are `sealed classes` used in `domain/` to represent states?
- [ ] Is `main.dart` clean, delegating logic to `core/`?
- [ ] Are entry points per flavor (`main_dev.dart`, etc.) configured?


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.367515Z"
evidence_checksum: "NONE"
```
