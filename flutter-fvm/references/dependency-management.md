# Dependency Management: Flutter with FVM

Gerenciar dependências no Flutter com FVM é simples e segue os padrões do ecossistema Dart (Pub), mas sempre prefixado com `fvm`.

## 1. Comandos Básicos do Pub via FVM

| Ação | Comando |
|------|---------|
| Instalar dependências | `fvm flutter pub get` |
| Adicionar pacote | `fvm flutter pub add <package>` |
| Remover pacote | `fvm flutter pub remove <package>` |
| Atualizar pacotes | `fvm flutter pub upgrade` |
| Verificar vulnerabilidades | `fvm flutter pub deps` |

## 2. Dependências de Produção vs. Desenvolvimento

No `pubspec.yaml`:

```yaml
dependencies:
  flutter:
    sdk: flutter
  http: ^1.2.0
  riverpod: ^2.5.1

dev_dependencies:
  flutter_test:
    sdk: flutter
  flutter_lints: ^3.0.1
  mocktail: ^1.0.3
  build_runner: ^2.4.8 # Para geração de código
```

## 3. Versionamento Semântico (SemVer)

- `^1.2.3`: Permite atualizações que não quebrem (major = 1).
- `1.2.3`: Fixa exatamente na versão 1.2.3.
- `any`: Qualquer versão (não recomendado para produção).

## 4. O arquivo `pubspec.lock`

Este arquivo é gerado automaticamente e contém as versões EXATAS de cada pacote instalado. 

**REGRA DE OURO**: Sempre adicione o `pubspec.lock` ao controle de versão (Git). Isso garante que todo o time e o CI/CD usem as mesmas versões de dependências.

## 5. Geração de Código (`build_runner`)

Muitos pacotes populares (Freezed, JSON Serializable, Riverpod Generator) dependem da geração de código.

Execute sempre via FVM:

```bash
# Rodar uma única vez
fvm flutter pub run build_runner build --delete-conflicting-outputs

# Rodar em modo watch (desenvolvimento ativo)
fvm flutter pub run build_runner watch --delete-conflicting-outputs
```

## 6. Cache do Pub

O Pub mantém um cache global de pacotes. Se você encontrar erros estranhos de dependência, tente limpar o cache:

```bash
fvm flutter pub cache clean
fvm flutter pub get
```
