# Project Structure: Flutter with FVM (2026 Edition)

A organização por funcionalidades (**Feature-First**) isola o escopo e facilita a manutenção de times grandes.

## 1. Estrutura de Diretórios
```text
lib/
├── core/                # Código transversal (di, network, theme)
├── features/            # Funcionalidades de negócio
│   ├── [feature_name]/  
│   │   ├── data/        # Repositories, API Sources
│   │   ├── domain/      # Sealed States, Entities
│   │   └── presentation/# ViewModels, Views, Widgets
├── main_dev.dart        # Entry points por flavor
└── main_prod.dart
```

## 2. Padrão Arquitetural: MVVM + Dart 3
O Dart 3 permite que o ViewModel exponha estados imutáveis e exaustivos.

### ❌ Ruim (Lógica Espalhada)
- Widgets fazendo chamadas de API diretamente.
- Uso de `if (loading) ... else if (error) ...`.

### ✅ Bom (Reativo e Tipado)
- **ViewModel**: Usa `AsyncNotifier` (Riverpod) ou `Bloc` e emite uma `sealed class`.
- **View**: Usa `switch (state)` para renderizar a UI.

## 3. Gestão de Secrets via Flavors
Use `--dart-define-from-file` para carregar configurações JSON.

```bash
fvm flutter run --flavor dev --dart-define-from-file=env/dev.json
```

## Checklist de Estrutura
- [ ] A feature está isolada (não importa widgets privados de outras)?
- [ ] Usou `sealed classes` no `domain/` para representar estados?
- [ ] O `main.dart` está limpo, delegando a lógica para o `core/`?
- [ ] Os entry points por flavor (`main_dev.dart`, etc) estão configurados?
