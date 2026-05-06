# Flutter Security Guide (OWASP)

This guide presents security practices for Flutter applications based on the OWASP Mobile Top 10 (2024).

## 1. M1: Improper Credential Usage

### Insecure Storage Patterns
**FORBIDDEN:**
```dart
// NEVER do this:
class ApiService {
  static const String apiKey = 'sk_live_1234567890abcdef'; // Hardcoded
  static const String secret = 'supersecret123';
}
```

**RECOMMENDED:**
```dart
// Use environment variables or secure storage
import 'package:flutter_dotenv/flutter_dotenv.dart';

class ApiService {
  final String apiKey;
  
  ApiService() : apiKey = dotenv.get('API_KEY');
}

// Or use flutter_secure_storage for tokens
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

final storage = FlutterSecureStorage();
await storage.write(key: 'auth_token', value: token);
```

### Sensitive Configurations & Secrets
**RECOMMENDED:**
- **--dart-define**: Use to pass keys at compile time without leaving them in the code.
```dart
// In code:
static const apiKey = String.fromEnvironment('API_KEY');
```
```bash
# In build command:
fvm flutter build apk --dart-define=API_KEY=your_key_here
```
- **.env files**: Use the `flutter_dotenv` package and ensure that `.env` is in `.gitignore`.
- **Firebase Config**: Do not version `google-services.json` (Android) or `GoogleService-Info.plist` (iOS) if they contain production keys.

**Check:**
- `.env` or `.config` files with credentials.
- Old commits containing secrets (use `trufflehog` or `git-secrets`).

## 2. M2: Inadequate Supply Chain Security

### Dependency Analysis
**Check your `pubspec.yaml`:**
```yaml
dependencies:
  # GOOD - Specific version
  http: ^1.1.0
  
  # BAD - "any" allows any version
  some_package: any  # AVOID THIS!
  
  # BAD - Range too wide
  another_package: '>=1.0.0 <3.0.0'  # Too permissive
```

**Verification commands:**
```bash
# List outdated dependencies
fvm flutter pub outdated

# Check for known vulnerabilities
# (Use tools like `dart pub global activate pana` and `pana .`)
```

### Package Audit
**Checklist:**
- [ ] Actively maintained packages (last update < 6 months)
- [ ] Popularity and number of stars on GitHub
- [ ] Open vs. closed issues
- [ ] Presence of tests in the repository
- [ ] License compatible with your project

## 3. M3: Insecure Authentication/Authorization

### Token Management
**FORBIDDEN:**
```dart
// NEVER store tokens in SharedPreferences without encryption
final prefs = await SharedPreferences.getInstance();
prefs.setString('auth_token', token); // Insecure!
```

**RECOMMENDED:**
```dart
// Use flutter_secure_storage
import 'package:flutter_secure_storage/flutter_secure_storage.dart';

final storage = FlutterSecureStorage();
await storage.write(key: 'auth_token', value: token);

// With automatic expiration
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
      // Force logout
      await logout();
    }
  }
}
```

## 4. M4: Insufficient Input/Output Validation

### Input Validation
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
      return ValidationResult.error('Email is required');
    }
    
    if (!_emailRegex.hasMatch(email)) {
      return ValidationResult.error('Invalid email');
    }
    
    if (_sqlInjectionRegex.hasMatch(email)) {
      return ValidationResult.error('Suspicious input detected');
    }
    
    return ValidationResult.success();
  }
  
  static ValidationResult sanitizeForDatabase(String input) {
    // Remove dangerous characters
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

### Mandatory HTTPS
**FORBIDDEN:**
```dart
// NEVER use HTTP in production
final response = await http.get(Uri.parse('http://api.example.com/data'));
```

**RECOMMENDED:**
```dart
// Always use HTTPS
final response = await http.get(
  Uri.https('api.example.com', '/data'),
  headers: {'Authorization': 'Bearer $token'},
);

// With certificate pinning (Android)
// Add in android/app/src/main/res/xml/network_security_config.xml:
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

### Secure Storage
**FORBIDDEN:**
```dart
// NEVER store sensitive data in plain text
final prefs = await SharedPreferences.getInstance();
prefs.setString('user_ssn', '123-45-6789');
prefs.setString('credit_card', '4111111111111111');
```

**RECOMMENDED:**
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

## 7. Secure Builds

### Obfuscation and Protection
```bash
# Android with obfuscation
fvm flutter build apk --release \
  --obfuscate \
  --split-debug-info=./build/symbols/ \
  --target-platform android-arm,android-arm64

# iOS
fvm flutter build ios --release \
  --obfuscate \
  --split-debug-info=./build/symbols/
```

### Release Settings
**Android (`android/app/build.gradle`):**
```groovy
android {
    buildTypes {
        release {
            signingConfig signingConfigs.release
            minifyEnabled true
            shrinkResources true
            proguardFiles getDefaultProguardFile('proguard-android.txt'), 'proguard-rules.pro'
            
            // Disable logging in release
            buildConfigField "boolean", "ENABLE_LOGGING", "false"
        }
    }
}
```

**iOS (`ios/Runner.xcworkspace`):**
- Enable Bitcode: NO (recommended for Flutter)
- Enable Dead Code Stripping: YES
- Optimization Level: Fastest, Smallest [-Os]

## 8. Security Scanner

### Automatable Checklist
```bash
#!/bin/bash
# security_scan.sh

echo "=== Flutter Security Scan ==="

# 1. Check for hardcoded credentials
echo "1. Scanning for hardcoded secrets..."
grep -r "sk_live\|api[_-]key\|secret\|password\|token" lib/ --include="*.dart" | grep -v "//\|test" || echo "✅ No obvious hardcoded secrets found"

# 2. Check for HTTP URLs
echo "2. Scanning for HTTP URLs..."
grep -r "http://" lib/ --include="*.dart" | grep -v "//\|test" || echo "✅ No HTTP URLs found"

# 3. Check dependencies
echo "3. Checking dependencies..."
fvm flutter pub outdated

# 4. Static analysis
echo "4. Running static analysis..."
fvm flutter analyze

echo "=== Scan Complete ==="
```

## 9. OWASP Security Checklist

**M1: Credentials**
- [ ] No hardcoded API keys in Dart code
- [ ] Sensitive settings in environment variables
- [ ] Firebase config not versioned in git
- [ ] Tokens stored with flutter_secure_storage

**M2: Supply Chain**
- [ ] `pubspec.yaml` without `any` versioning
- [ ] Dependencies updated regularly
- [ ] Packages audited for security
- [ ] `pubspec.lock` versioned

**M3: Authentication**
- [ ] Tokens with expiration
- [ ] Refresh tokens implemented
- [ ] Logout clears all sensitive data
- [ ] Biometrics/face ID for sensitive data

**M4: Input Validation**
- [ ] Email/password validation on client
- [ ] Sanitization of inputs for database
- [ ] Protection against SQL injection
- [ ] Input size limits

**M5: Communication**
- [ ] HTTPS in all API calls
- [ ] Certificate pinning for critical APIs
- [ ] Security headers (CSP, HSTS)
- [ ] Timeouts configured in requests

**M9: Storage**
- [ ] Sensitive data encrypted
- [ ] SharedPreferences only for non-sensitive data
- [ ] Automatic cache cleanup
- [ ] Backup exclude for sensitive data

**Build & Deploy**
- [ ] Obfuscation enabled in release
- [ ] Debug disabled in production
- [ ] Logging reduced in production
- [ ] Code signing configured correctly
