# Flutter Security Guide: OWASP Mobile Top 10

Este guia apresenta práticas de segurança para aplicações Flutter baseadas no OWASP Mobile Top 10 (2024).

## 1. M1: Improper Credential Usage

### Padrões de Armazenamento Inseguro
**PROIBIDO:**
```dart
// NUNCA faça isso:
class ApiService {
  static const String apiKey = 'sk_live_1234567890abcdef'; // Hardcoded
  static const String secret = 'supersecret123';
}
```

**RECOMENDADO:**
```dart
// Use variáveis de ambiente ou secure storage
import 'package:flutter_dotenv/flutter_dotenv';

class ApiService {
  final String apiKey;
  
  ApiService() : apiKey = dotenv.get('API_KEY');
}

// Ou use flutter_secure_storage para tokens
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

final storage = FlutterSecureStorage();
await storage.write(key: 'auth_token', value: token);
```

### Configurações Sensíveis
**Verifique:**
- `google-services.json` (Firebase) - Não versionar chaves de produção
- `GoogleService-Info.plist` - iOS Firebase config
- `appcenter-config.json` - App Center secrets
- Qualquer arquivo `.env` ou `.config` com credenciais

## 2. M2: Inadequate Supply Chain Security

### Análise de Dependências
**Verifique seu `pubspec.yaml`:**
```yaml
dependencies:
  # BOM - Versão específica
  http: ^1.1.0
  
  # RUIM - "any" permite qualquer versão
  some_package: any  # EVITE ISSO!
  
  # RUIM - Range muito amplo
  another_package: '>=1.0.0 <3.0.0'  # Muito permissivo
```

**Comandos de verificação:**
```bash
# Listar dependências desatualizadas
fvm flutter pub outdated

# Verificar vulnerabilidades conhecidas
# (Use ferramentas como `dart pub global activate pana` e `pana .`)
```

### Auditoria de Pacotes
**Checklist:**
- [ ] Pacotes mantidos ativamente (última atualização < 6 meses)
- [ ] Popularidade e número de stars no GitHub
- [ ] Issues abertos vs. fechados
- [ ] Presença de testes no repositório
- [ ] Licença compatível com seu projeto

## 3. M3: Insecure Authentication/Authorization

### Gerenciamento de Tokens
**PROIBIDO:**
```dart
// NUNCA armazene tokens em SharedPreferences sem criptografia
final prefs = await SharedPreferences.getInstance();
prefs.setString('auth_token', token); // Inseguro!
```

**RECOMENDADO:**
```dart
// Use flutter_secure_storage
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

final storage = FlutterSecureStorage();
await storage.write(key: 'auth_token', value: token);

// Com expiração automática
await storage.write(
  key: 'auth_token',
  value: token,
  iOptions: const IOSOptions(
    accessibility: KeychainAccessibility.first_unlock_this_device,
  ),
  aOptions: const AndroidOptions(
    encryptedSharedPreferences: true,
  ),
);
```

### Refresh Tokens
```dart
class AuthService {
  final SecureStorage storage;
  final ApiClient api;
  Timer? _refreshTimer;
  
  Future<void> scheduleTokenRefresh(String token) async {
    final payload = _decodeJwt(token);
    final expiry = DateTime.fromMillisecondsSinceEpoch(payload['exp'] * 1000);
    final refreshTime = expiry.subtract(const Duration(minutes: 5));
    
    _refreshTimer = Timer(
      refreshTime.difference(DateTime.now()),
      () => _refreshToken(),
    );
  }
  
  Future<void> _refreshToken() async {
    try {
      final newToken = await api.refreshToken();
      await storage.write(key: 'auth_token', value: newToken);
      await scheduleTokenRefresh(newToken);
    } catch (e) {
      // Forçar logout
      await logout();
    }
  }
}
```

## 4. M4: Insufficient Input/Output Validation

### Validação de Entrada
```dart
class UserInputValidator {
  static final _emailRegex = RegExp(
    r'^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$'
  );
  
  static final _sqlInjectionRegex = RegExp(
    r"(\b(SELECT|INSERT|UPDATE|DELETE|DROP|UNION|EXEC)\b|--|;)",
    caseSensitive: false,
  );
  
  static ValidationResult validateEmail(String email) {
    if (email.isEmpty) {
      return ValidationResult.error('Email é obrigatório');
    }
    
    if (!_emailRegex.hasMatch(email)) {
      return ValidationResult.error('Email inválido');
    }
    
    if (_sqlInjectionRegex.hasMatch(email)) {
      return ValidationResult.error('Entrada suspeita detectada');
    }
    
    return ValidationResult.success();
  }
  
  static ValidationResult sanitizeForDatabase(String input) {
    // Remove caracteres perigosos
    final sanitized = input
      .replaceAll("'", "''")
      .replaceAll(';', '')
      .replaceAll('--', '')
      .trim();
    
    return ValidationResult.success(sanitized);
  }
}
```

## 5. M5: Insecure Communication

### HTTPS Obrigatório
**PROIBIDO:**
```dart
// NUNCA use HTTP em produção
final response = await http.get(Uri.parse('http://api.example.com/data'));
```

**RECOMENDADO:**
```dart
// Sempre use HTTPS
final response = await http.get(
  Uri.https('api.example.com', '/data'),
  headers: {'Authorization': 'Bearer $token'},
);

// Com certificate pinning (Android)
// Adicione no android/app/src/main/res/xml/network_security_config.xml:
/*
<network-security-config>
  <domain-config>
    <domain includeSubdomains="true">api.example.com</domain>
    <pin-set>
      <pin digest="SHA-256">BASE64_HASH_OF_CERTIFICATE</pin>
      <pin digest="SHA-256">BACKUP_CERTIFICATE_HASH</pin>
    </pin-set>
  </domain-config>
</network-security-config>
*/
```

## 6. M9: Insecure Data Storage

### Armazenamento Seguro
**PROIBIDO:**
```dart
// NUNCA armazene dados sensíveis em plain text
final prefs = await SharedPreferences.getInstance();
prefs.setString('user_ssn', '123-45-6789');
prefs.setString('credit_card', '4111111111111111');
```

**RECOMENDADO:**
```dart
import 'package:encrypt/encrypt.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

class SecureDataStorage {
  final FlutterSecureStorage _secureStorage;
  final Encrypter _encrypter;
  
  Future<void> storeSensitiveData(String key, String value) async {
    final encrypted = _encrypter.encrypt(value);
    await _secureStorage.write(
      key: key,
      value: encrypted.base64,
    );
  }
  
  Future<String?> retrieveSensitiveData(String key) async {
    final encryptedBase64 = await _secureStorage.read(key: key);
    if (encryptedBase64 == null) return null;
    
    final encrypted = Encrypted.fromBase64(encryptedBase64);
    return _encrypter.decrypt(encrypted);
  }
}
```

## 7. Builds Seguros

### Ofuscação e Proteção
```bash
# Android com ofuscação
fvm flutter build apk --release \
  --obfuscate \
  --split-debug-info=./build/symbols/ \
  --target-platform android-arm,android-arm64

# iOS
fvm flutter build ios --release \
  --obfuscate \
  --split-debug-info=./build/symbols/
```

### Configurações de Release
**Android (`android/app/build.gradle`):**
```groovy
android {
    buildTypes {
        release {
            signingConfig signingConfigs.release
            minifyEnabled true
            shrinkResources true
            proguardFiles getDefaultProguardFile('proguard-android.txt'), 'proguard-rules.pro'
            
            // Desabilitar logging em release
            buildConfigField "boolean", "ENABLE_LOGGING", "false"
        }
    }
}
```

**iOS (`ios/Runner.xcworkspace`):**
- Habilitar Bitcode: NO (recomendado para Flutter)
- Habilitar Dead Code Stripping: YES
- Nível de Otimização: Fastest, Smallest [-Os]

## 8. Scanner de Segurança

### Checklist Automatizável
```bash
#!/bin/bash
# security_scan.sh

echo "=== Flutter Security Scan ==="

# 1. Verificar credenciais hardcoded
echo "1. Scanning for hardcoded secrets..."
grep -r "sk_live\|api[_-]key\|secret\|password\|token" lib/ --include="*.dart" | grep -v "//\|test" || echo "✅ No obvious hardcoded secrets found"

# 2. Verificar HTTP URLs
echo "2. Scanning for HTTP URLs..."
grep -r "http://" lib/ --include="*.dart" | grep -v "//\|test" || echo "✅ No HTTP URLs found"

# 3. Verificar dependências
echo "3. Checking dependencies..."
fvm flutter pub outdated

# 4. Análise estática
echo "4. Running static analysis..."
fvm flutter analyze

echo "=== Scan Complete ==="
```

## 9. Checklist de Segurança OWASP

**M1: Credenciais**
- [ ] Nenhuma chave API hardcoded no código Dart
- [ ] Configurações sensíveis em variáveis de ambiente
- [ ] Firebase config não versionado no git
- [ ] Tokens armazenados com flutter_secure_storage

**M2: Supply Chain**
- [ ] `pubspec.yaml` sem versionamento `any`
- [ ] Dependências atualizadas regularmente
- [ ] Pacotes auditados por segurança
- [ ] `pubspec.lock` versionado

**M3: Autenticação**
- [ ] Tokens com expiração
- [ ] Refresh tokens implementados
- [ ] Logout limpa todos os dados sensíveis
- [ ] Biometria/face ID para dados sensíveis

**M4: Input Validation**
- [ ] Validação de email/senha no client
- [ ] Sanitização de inputs para database
- [ ] Proteção contra SQL injection
- [ ] Limites de tamanho de input

**M5: Comunicação**
- [ ] HTTPS em todas as chamadas API
- [ ] Certificate pinning para APIs críticas
- [ ] Headers de segurança (CSP, HSTS)
- [ ] Timeout configurados em requisições

**M9: Armazenamento**
- [ ] Dados sensíveis criptografados
- [ ] SharedPreferences apenas para dados não-sensíveis
- [ ] Limpeza automática de cache
- [ ] Backup exclude de dados sensíveis

**Build & Deploy**
- [ ] Ofuscação habilitada em release
- [ ] Debug desabilitado em produção
- [ ] Logging reduzido em produção
- [ ] Code signing configurado corretamente