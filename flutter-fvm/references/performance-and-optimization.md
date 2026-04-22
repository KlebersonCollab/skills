# Performance & Otimização

Garantir 60/120 FPS exige controle rigoroso sobre o pipeline de renderização.

## 1. Selective Rebuilds (MediaQuery)
**Conceito**: Ouvir apenas o que é necessário do contexto para evitar reconstruções globais.

### ❌ Ruim
```dart
// Reconstrói se qualquer coisa no MediaQuery mudar (ex: teclado abre)
final size = MediaQuery.of(context).size;
```

### ✅ Bom
```dart
// Reconstrói APENAS se o tamanho mudar
final size = MediaQuery.sizeOf(context);
```

## 2. Widget Decomposition
**Conceito**: Isolar partes da UI em classes `StatelessWidget` para aproveitar o cache do framework.

### ❌ Ruim (Helper Methods)
```dart
Widget _buildHeader() { // Reconstruído sempre que o pai reconstrói
  return Text('Header');
}
```

### ✅ Bom (Widget Classes)
```dart
class Header extends StatelessWidget { // Cacheável e suporta const
  const Header({super.key});
  @override
  Widget build(BuildContext context) => const Text('Header');
}
```

## 3. RepaintBoundary
**Conceito**: Cria uma camada separada no canvas para subárvores que mudam independentemente.

### ✅ Bom (Uso em Animações)
```dart
RepaintBoundary(
  child: CircularProgressIndicator(), // O resto da tela não repinta
)
```

## Checklist de Performance
- [ ] Usou `MediaQuery.sizeOf` em vez de `.of`?
- [ ] Métodos `build` têm menos de 100 linhas?
- [ ] Funções que retornam widgets foram transformadas em classes?
- [ ] Usou `const` em todos os widgets estáticos?
- [ ] Imagens pesadas usam `cacheWidth` / `cacheHeight`?
