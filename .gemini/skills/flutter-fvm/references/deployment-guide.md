# Deployment Guide: Flutter with FVM

This guide describes the commands for generating release artifacts on different platforms using the FVM workflow.

## 1. Preparation for Deployment

Before any release build, always perform cleaning and dependency update:

```bash
fvm flutter clean
fvm flutter pub get
```

## 2. Android

### Generate APK (Debug/Release)
```bash
# Debug APK
fvm flutter build apk --debug

# Release APK (Usually for internal distribution or testing)
fvm flutter build apk --release
```

### Generate Android App Bundle (AAB) - Recommended for Play Store
```bash
fvm flutter build appbundle --release
```

## 3. iOS

### Generate the app directory (.app)
```bash
fvm flutter build ios --release
```

### Generate the IPA file (Xcode)
After the above command, it's common to open the project in Xcode for archiving and uploading:
```bash
open ios/Runner.xcworkspace
```

Or via CLI if you have `fastlane` configured (highly recommended).

## 4. Web

### Generate static files
```bash
fvm flutter build web --release
```

This will generate the files in the `build/web/` folder. They can be served by any static server (Firebase Hosting, Netlify, Vercel, S3).

## 5. Desktop

### Windows
```bash
fvm flutter build windows --release
```

### macOS
```bash
fvm flutter build macos --release
```

## 6. App Versioning

The app version is defined in `pubspec.yaml`:

```yaml
version: 1.0.0+1
```

- `1.0.0`: Semantic version visible to the user.
- `+1`: Build number (integer, incremented with each new deployment).

To automate incrementing in CI/CD, you can use flags:
```bash
fvm flutter build apk --release --build-name=1.1.0 --build-number=42
```

## 7. Obfuscation and Security

To protect your code against reverse engineering:

```bash
fvm flutter build apk --release --obfuscate --split-debug-info=./build/debug-info
```


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.366113Z"
evidence_checksum: "NONE"
```
