# Performance & Optimization

Ensuring 60/120 FPS requires strict control over the rendering pipeline.

## 1. Selective Rebuilds (MediaQuery)
**Concept**: Listening only to what is necessary from the context to avoid global rebuilds.

### ❌ Bad
```dart
// Rebuilds if anything in MediaQuery changes (e.g., keyboard opens)
final size = MediaQuery.of(context).size;
```

### ✅ Good
```dart
// Rebuilds ONLY if size changes
final size = MediaQuery.sizeOf(context);
```

## 2. Widget Decomposition
**Concept**: Isolating parts of the UI into `StatelessWidget` classes to take advantage of the framework's cache.

### ❌ Bad (Helper Methods)
```dart
Widget _buildHeader() { // Rebuilt whenever parent rebuilds
  return Text('Header');
}
```

### ✅ Good (Widget Classes)
```dart
class Header extends StatelessWidget { // Cacheable and supports const
  const Header({super.key});
  @override
  Widget build(BuildContext context) => const Text('Header');
}
```

## 3. RepaintBoundary
**Concept**: Creates a separate layer on the canvas for subtrees that change independently.

### ✅ Good (Use in Animations)
```dart
RepaintBoundary(
  child: CircularProgressIndicator(), // Rest of screen doesn't repaint
)
```

## Performance Checklist
- [ ] Used `MediaQuery.sizeOf` instead of `.of`?
- [ ] Are `build` methods less than 100 lines?
- [ ] Were functions that return widgets transformed into classes?
- [ ] Used `const` in all static widgets?
- [ ] Do heavy images use `cacheWidth` / `cacheHeight`?
