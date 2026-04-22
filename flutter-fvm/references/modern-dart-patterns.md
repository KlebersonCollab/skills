# Padrões Modernos de Dart (v3.0+)

O Dart 3 transformou o Flutter de um framework puramente imperativo para um modelo rico em tipos e padrões funcionais.

## 1. Sealed Classes (Unions)
**Conceito**: Garante que o compilador conheça todos os subtipos possíveis de uma classe, permitindo exaustividade em switches.

### ❌ Ruim (Boolean Flags)
```dart
class UserState {
  final bool isLoading;
  final String? error;
  final User? data;
  // Permite estados impossíveis como isLoading=true E data!=null
}
```

### ✅ Bom (Sealed)
```dart
sealed class UserState {}
class UserInitial extends UserState {}
class UserLoading extends UserState {}
class UserSuccess extends UserState { final User user; UserSuccess(this.user); }
class UserError extends UserState { final String msg; UserError(this.msg); }
```

## 2. Pattern Matching
**Conceito**: Extração de dados e controle de fluxo baseado na estrutura do objeto.

### ✅ Bom (Switch Expression)
```dart
final message = switch (state) {
  UserInitial() => 'Bem-vindo',
  UserLoading() => 'Carregando...',
  UserSuccess(user: var u) => 'Olá, ${u.name}',
  UserError(msg: var m) => 'Erro: $m',
};
```

## 3. Records
**Conceito**: Tipos anônimos para retornos múltiplos sem necessidade de classes pesadas.

### ✅ Bom
```dart
(double, double) getCoordinates() => (-23.5, -46.6);

// Destructuring
final (lat, lng) = getCoordinates();
```

## Checklist Modern Dart
- [ ] Usou `sealed` para estados de UI?
- [ ] O `switch` é uma expressão (com retorno) sempre que possível?
- [ ] Evitou `dynamic` habilitando `strict-casts` no analyzer?
- [ ] Usou `records` para retornos múltiplos simples?
