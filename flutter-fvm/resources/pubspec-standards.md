# Pubspec Standards Guide

This guide defines the standards for organization, versioning, and maintenance of the `pubspec.yaml` file in projects using the Hub ecosystem.

## 1. Section Organization
The file should follow the default Flutter SDK order, maintaining clear sections:
1. Meta-information (name, description, version, etc.)
2. Environment (sdk, flutter version)
3. Dependencies (alphabetically ordered)
4. Dev Dependencies (alphabetically ordered)
5. Flutter Assets/Fonts/etc.

## 2. Versioning Rules
- **Production Dependencies**: Use the caret prefix (`^`) to allow patch and minor updates, ensuring compatibility.
- **Critical Dependencies**: Fix the version (e.g., `1.2.3`) only in cases of known breaking changes.
- **FVM**: Use `fvm flutter pub get` to ensure that dependency versions are resolved with the correct SDK.

## 3. Comments and Grouping
Group related dependencies with comments:
```yaml
dependencies:
  # Core
  flutter:
    sdk: flutter
  cupertino_icons: ^1.0.2

  # State Management
  flutter_bloc: ^8.1.0
  riverpod: ^2.3.0

  # Network & API
  dio: ^5.1.0
  json_annotation: ^4.8.0
```

## 4. Maintenance
- Run `flutter pub outdated` monthly.
- Avoid unused dependencies (DRY).
- Always check the score on Pub.dev before adding new libraries.
