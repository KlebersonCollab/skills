# Accessibility Guide (Inclusive UX)

Professional apps must be usable by everyone, including screen reader users and people with low vision.

## 1. Semantics
**Concept**: Providing metadata so that TalkBack/VoiceOver can describe the UI.

### ❌ Bad
```dart
GestureDetector(
  onTap: () => print('Play'),
  child: Icon(Icons.play_arrow), // The reader will only say "Icon" or nothing.
)
```

### ✅ Good
```dart
Semantics(
  label: 'Play Video Button',
  button: true,
  child: IconButton(
    icon: Icon(Icons.play_arrow),
    onPressed: () {},
  ),
)
```

## 2. Hit Targets
**Concept**: Minimum area required for comfortable physical interaction.

### ❌ Bad
```dart
SizedBox(
  width: 20,
  height: 20,
  child: InkWell(onTap: () {}), // Very hard to hit with a finger
)
```

### ✅ Good (48x48 dp)
```dart
ConstrainedBox(
  constraints: BoxConstraints.tightFor(width: 48, height: 48),
  child: InkWell(onTap: () {}),
)
```

## 3. Text Contrast
**Concept**: Readability based on the relationship between background color and font color.

### ✅ Good
- Minimum contrast of **4.5:1**.
- Don't use color alone to indicate an error (e.g., use a warning icon ⚠️ along with a red border).

## Accessibility Checklist
- [ ] Do all interactive icons have `semanticLabel` or are they inside `Semantics`?
- [ ] Do hit targets have at least 48x48 dp?
- [ ] Was the app tested with the system font size at maximum?
- [ ] Did you use `MergeSemantics` to group labels and values?


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.364980Z"
evidence_checksum: "NONE"
```
