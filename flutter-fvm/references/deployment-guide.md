# Deployment Guide: Flutter with FVM

Este guia descreve os comandos para gerar artefatos de release em diferentes plataformas utilizando o workflow do FVM.

## 1. Preparação para o Deploy

Antes de qualquer build de release, execute sempre a limpeza e a atualização de dependências:

```bash
fvm flutter clean
fvm flutter pub get
```

## 2. Android

### Gerar APK (Debug/Release)
```bash
# APK de Debug
fvm flutter build apk --debug

# APK de Release (Geralmente para distribuição interna ou testes)
fvm flutter build apk --release
```

### Gerar Android App Bundle (AAB) - Recomendado para Play Store
```bash
fvm flutter build appbundle --release
```

## 3. iOS

### Gerar o diretório do app (.app)
```bash
fvm flutter build ios --release
```

### Gerar o arquivo IPA (Xcode)
Após o comando acima, é comum abrir o projeto no Xcode para arquivamento e upload:
```bash
open ios/Runner.xcworkspace
```

Ou via CLI se você tiver o `fastlane` configurado (altamente recomendado).

## 4. Web

### Gerar os arquivos estáticos
```bash
fvm flutter build web --release
```

Isso gerará os arquivos na pasta `build/web/`. Eles podem ser servidos por qualquer servidor estático (Firebase Hosting, Netlify, Vercel, S3).

## 5. Desktop

### Windows
```bash
fvm flutter build windows --release
```

### macOS
```bash
fvm flutter build macos --release
```

## 6. Versionamento do App

A versão do app é definida no `pubspec.yaml`:

```yaml
version: 1.0.0+1
```

- `1.0.0`: Versão semântica visível para o usuário.
- `+1`: Número de build (inteiro, incrementado a cada novo deploy).

Para automatizar o incremento em CI/CD, você pode usar flags:
```bash
fvm flutter build apk --release --build-name=1.1.0 --build-number=42
```

## 7. Ofuscação e Segurança

Para proteger seu código contra engenharia reversa:

```bash
fvm flutter build apk --release --obfuscate --split-debug-info=./build/debug-info
```
