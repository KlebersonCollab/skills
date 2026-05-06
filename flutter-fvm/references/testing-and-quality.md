# Testing & Quality (Mobile)

Quality in Flutter is a combination of rigorous static analysis and layered test coverage.

## 1. Test Pyramid
- **Unit Tests**: Pure logic and ViewModels.
- **Widget Tests**: Interactions and Accessibility.
- **Golden Tests**: Visual regression.
- **Integration Tests**: E2E flows.

## 2. Accessibility Testing
**Concept**: Validate that the semantic tree is correct.

### ✅ Good (Semantic Test)
```dart
testWidgets('Button should be accessible', (tester) async {
  final handle = tester.ensureSemantics();
  await tester.pumpWidget(const MyButton());

  expect(tester.getSemantics(find.byType(MyButton)), matchesSemantics(
    label: 'Send',
    isButton: true,
  ));
  handle.dispose();
});
```

## 3. Static Analysis (Strict)
Always run `fvm flutter analyze` before any push.

### ✅ Good (`analysis_options.yaml`)
```yaml
analyzer:
  errors:
    unused_local_variable: error
    invalid_annotation_target: ignore
  language:
    strict-casts: true
    strict-inference: true
```

## Quality Checklist
- [ ] Ran `fvm flutter analyze` and fixed all warnings?
- [ ] Business logic (ViewModel) has 80%+ coverage?
- [ ] Critical widgets have Semantics tests?
- [ ] Goldens were updated for the target platform?
- [ ] Used `mocktail` to isolate external dependencies?


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.367830Z"
evidence_checksum: "NONE"
```
