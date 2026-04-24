#!/bin/bash

# Flutter Security Scanner
# Based on OWASP Mobile Top 10 (2024)
# Usage: ./security-scan.sh

set -e

echo "=== Flutter Security Scan (OWASP Mobile Top 10) ==="
echo ""

# 1. M1: Improper Credential Usage - Hardcoded Credentials
echo "1. 🔍 Scanning for hardcoded secrets (M1)..."
SECRET_PATTERNS="sk_live_|api[_-]key|secret[_-]|password|token|auth[_-]|bearer|jwt|firebase"
if grep -r -i -E "$SECRET_PATTERNS" lib/ --include="*.dart" | grep -v "//\|test\|TODO\|FIXME" > /dev/null 2>&1; then
    echo "   ⚠️  WARNING: Potential hardcoded secrets found:"
    grep -r -i -E "$SECRET_PATTERNS" lib/ --include="*.dart" | grep -v "//\|test\|TODO\|FIXME"
else
    echo "   ✅ No obvious hardcoded secrets found"
fi

# 2. M2: Inadequate Supply Chain Security - Dependencies
echo ""
echo "2. 📦 Checking dependencies (M2)..."
if [ -f "pubspec.yaml" ]; then
    echo "   Checking for 'any' version constraints..."
    if grep -q "any" pubspec.yaml; then
        echo "   ⚠️  WARNING: Found 'any' version constraints in pubspec.yaml"
        grep -n "any" pubspec.yaml
    else
        echo "   ✅ No 'any' version constraints found"
    fi

    echo "   Checking for outdated packages..."
    fvm flutter pub outdated || echo "   ℹ️  Run 'fvm flutter pub outdated' manually for details"
else
    echo "   ❌ pubspec.yaml not found"
fi

# 3. M5: Insecure Communication - HTTP URLs
echo ""
echo "3. 🌐 Checking for HTTP URLs (M5)..."
if grep -r "http://" lib/ --include="*.dart" | grep -v "//\|test\|TODO\|FIXME\|localhost\|127.0.0.1" > /dev/null 2>&1; then
    echo "   ⚠️  WARNING: HTTP URLs found (use HTTPS in production):"
    grep -r "http://" lib/ --include="*.dart" | grep -v "//\|test\|TODO\|FIXME\|localhost\|127.0.0.1"
else
    echo "   ✅ No HTTP URLs found (excluding localhost)"
fi

# 4. M9: Insecure Data Storage - SharedPreferences usage
echo ""
echo "4. 💾 Checking storage patterns (M9)..."
if grep -r "SharedPreferences\|shared_preferences" lib/ --include="*.dart" > /dev/null 2>&1; then
    echo "   ℹ️  SharedPreferences usage detected. Ensure sensitive data is encrypted."
    echo "   Recommended: Use flutter_secure_storage for sensitive data"
fi

# 5. Build configuration check
echo ""
echo "5. 🛠️  Checking build configuration..."
if [ -f "android/app/build.gradle" ]; then
    echo "   Checking Android release build settings..."
    if grep -q "minifyEnabled true" android/app/build.gradle && \
       grep -q "shrinkResources true" android/app/build.gradle; then
        echo "   ✅ Android minifyEnabled and shrinkResources enabled"
    else
        echo "   ⚠️  Consider enabling minifyEnabled and shrinkResources in release builds"
    fi
fi

# 6. Analysis check
echo ""
echo "6. 📊 Running static analysis..."
fvm flutter analyze --no-pub --no-fatal-infos || true

echo ""
echo "=== Security Scan Complete ==="
echo ""
echo "Recommended Actions:"
echo "1. Review any warnings above"
echo "2. Run 'fvm flutter pub upgrade' to update dependencies"
echo "3. Consider adding certificate pinning for production APIs"
echo "4. Use flutter_secure_storage for sensitive data"
echo "5. Enable obfuscation in release builds:"
echo "   fvm flutter build apk --release --obfuscate --split-debug-info=./build/symbols/"
echo ""
echo "For detailed security guidance, see references/flutter-security-guide.md"