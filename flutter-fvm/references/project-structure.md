# Project Structure: Flutter with FVM (2026 Edition)

Este guia descreve a organização de diretórios e padrões arquiteturais recomendados para projetos Flutter profissionais em 2026, utilizando o FVM.

## 1. Estrutura de Diretórios: Feature-First

A organização por funcionalidades (**Feature-First**) é o padrão de fato em 2026, pois isola o escopo, facilita a escalabilidade e a manutenção de times grandes.

```text
lib/
├── core/                # Código compartilhado e transversal
│   ├── di/              # Injeção de Dependência (Riverpod ou GetIt)
│   ├── network/         # Dio, Interceptors, API Exceptions
│   ├── router/          # GoRouter (Deep linking, Auth protection)
│   ├── theme/           # Design System (Colors, Typography, Themes)
│   └── utils/           # Extensions, Helpers, Constants
├── features/            # Funcionalidades de negócio (Dominio e UI)
│   ├── auth/            # Exemplo: Feature de Autenticação
│   │   ├── data/        # Repositories, DTOs, API Sources
│   │   ├── domain/      # Entities, Business Rules (opcional para MVVM puro)
│   │   └── presentation/# Camada de UI (MVVM)
│   │       ├── view_models/ # Gestão de Estado (Riverpod AsyncNotifier/Bloc)
│   │       ├── views/       # Widgets de Página (Declarativos)
│   │       └── widgets/     # Widgets locais da funcionalidade
├── env/                 # Arquivos de configuração por ambiente (.env, .json)
├── main_dev.dart        # Entry point para Desenvolvimento
├── main_staging.dart    # Entry point para Homologação
└── main_prod.dart       # Entry point para Produção
```

## 2. Padrão Arquitetural: MVVM (Model-View-ViewModel)

O Flutter em 2026 consolidou o **MVVM** como o equilíbrio ideal entre desacoplamento e produtividade.

- **View**: Widgets puros e declarativos. Observam o estado exposto pelo ViewModel e reagem a ele. Não possuem lógica de negócio.
- **ViewModel**: Gerencia o estado da View e se comunica com a camada de dados. Utiliza ferramentas modernas como **Riverpod (AsyncNotifier)** ou **Bloc/Cubit**.
- **Model**: Representa os dados (Entities/DTOs) e a lógica de persistência ou rede.

**Dica**: Utilize a lib `freezed` para criar modelos imutáveis e estados de UI (ex: `Success`, `Loading`, `Error`).

## 3. Gestão de Ambientes: Flavors

Não troque URLs ou chaves de API manualmente. Utilize **Flavors** para gerenciar diferentes configurações de ambiente de forma segura e automatizada.

### Como executar com Flavors via FVM:
```bash
# Rodar em Desenvolvimento
fvm flutter run --flavor dev -t lib/main_dev.dart

# Build de Release para Produção (Android)
fvm flutter build appbundle --release --flavor prod -t lib/main_prod.dart
```

### Configuração Recomendada:
Utilize a flag `--dart-define-from-file` para carregar configurações de arquivos JSON, evitando expor segredos no código:
```bash
fvm flutter run --flavor dev --dart-define-from-file=env/dev.json
```

## 4. Stack Tecnológica Recomendada (2026)

- **State Management**: **Riverpod** (nativo, sem context) ou **BLoC** (event-driven).
- **Navigation**: **GoRouter** (padrão oficial do Flutter Team).
- **Networking**: **Dio** + **Retrofit** (geração de código para APIs resilientes).
- **Local Storage**: **Isar** (sucessor do Hive, extremamente rápido) ou **Drift** (SQLite reativo).
- **Injeção de Dependência**: **Riverpod** (DI nativa) ou **GetIt** + **Injectable**.

## 5. Boas Práticas Cruciais

1.  **Feature Isolation**: Uma feature não deve importar widgets privados de outra feature.
2.  **Imutabilidade**: Use `final` por padrão e `copyWith` para atualizações de estado.
3.  **Error Handling**: Utilize o padrão **Functional Error Handling** (lib `fpdart`) com o tipo `Either` para evitar `try-catch` na UI.
4.  **Design System**: Crie uma biblioteca de widgets atômicos no `core/theme` para garantir consistência visual.
5.  **Clean Code**: Siga as diretrizes de [Effective Dart](https://dart.dev/guides/language/effective-dart).
