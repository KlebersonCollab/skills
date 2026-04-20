---
name: flutter-fvm
version: 1.2.0
description: "Skill para desenvolvimento Flutter profissional com FVM (Flutter Version Management). Use quando precisar gerenciar múltiplas versões do SDK do Flutter por projeto, garantir consistência entre ambientes de desenvolvimento, gerenciar dependências com pubspec.yaml, configurar linting, testes avançados por camada arquitetural, segurança OWASP Mobile Top 10 e builds multiplataforma seguros."
category: development-workflow
---

# Flutter com FVM

> Gerenciador de versões do Flutter que permite fixar e gerenciar múltiplas versões do SDK por projeto, garantindo consistência total no time de desenvolvimento.

---

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar a implementação:
1. **Spec Check**: O arquivo `spec.md` existe e contém os critérios de aceitação (ACs)?
2. **Plan Check**: O arquivo `plan.md` define a arquitetura e schemas?
3. **Task Check**: A lista de tarefas em `tasks.md` está detalhada?

---

## Goal

**FVM** é o padrão ouro para gerenciamento de versões do Flutter. Ele isola o SDK por projeto, evitando o clássico "na minha máquina funciona" causado por versões globais do Flutter diferentes.

| Ferramenta / Conceito | Equivalente FVM |
|-----------------------|---------------|
| `flutter` (global)    | `fvm flutter` (local) |
| `dart` (global)       | `fvm dart` (local) |
| Versão do SDK         | `fvm use <versão>` |
| Instalação de SDK     | `fvm install <versão>` |
| Configuração Local    | `.fvm/fvm_config.json` |

**Vantagens**: Troca instantânea de versões, cache compartilhado de SDKs, configuração determinística para CI/CD e IDEs.

## Version Awareness

**Versão Recomendada: Flutter Stable (última disponível via FVM)**

Verifique as versões instaladas e a ativa no projeto:

```bash
fvm list
fvm current
```

Se o FVM não estiver instalado:

```bash
# Via Dart (recomendado)
dart pub global activate fvm

# Ou via Homebrew (macOS)
brew tap leoafarias/fvm
brew install fvm
```

## When to Use This Skill

Use this skill when:
- Inicializando um novo projeto Flutter
- Gerenciando múltiplas versões do Flutter SDK em diferentes projetos
- Adicionando dependências via `pubspec.yaml`
- Configurando linting (`analysis_options.yaml`) e formatação
- Executando testes (unit, widget, integration) seguindo padrões de camadas arquiteturais
- Implementando testes de estado com Riverpod ou outros gerenciadores de estado
- Realizando análise de segurança OWASP Mobile Top 10
- Gerando builds seguros para Android, iOS, Web ou Desktop
- Configurando CI/CD para projetos Flutter com FVM
- Resolvendo conflitos de versão do SDK
- Implementando padrões de segurança em aplicações móveis

Skip this skill when:
- O projeto NÃO usa Flutter/Dart
- O usuário prefere usar a instalação global do Flutter (não recomendado para projetos profissionais)

---

## Workflow (4 Fases)

### Fase 1: ENVIRONMENT — Configuração e Alinhamento
1.  **Verificar FVM**: Garantir que o FVM está instalado (`fvm --version`).
2.  **Fixar Versão do SDK**: Executar `fvm use stable` (ou versão específica como `3.19.0`) no root do projeto.
3.  **Configurar IDE**: Garantir que o VS Code ou Android Studio aponta para o SDK local em `.fvm/flutter_sdk`.

### Fase 2: PROJECT — Inicialização e Estrutura
1.  **Scaffolding**: Para novos projetos, use `fvm flutter create . --platforms=android,ios`.
2.  **Declarar Dependências**: Adicionar pacotes via `fvm flutter pub add <package>`.
3.  **Análise Estática**: Configurar `analysis_options.yaml` com regras rígidas (ex: `package:flutter_lints`).

### Fase 3: DEVELOP — Qualidade, Testes e Execução
1.  **Execução Segura**: Sempre utilize `fvm flutter run` para garantir que o código rode com o SDK correto.
2.  **Geração de Código**: Se usar `build_runner`, execute `fvm flutter pub run build_runner build --delete-conflicting-outputs`.
3.  **Testes por Camadas**: Seguir padrões de teste por camada arquitetural:
    - **Repositório**: Mock DAOs e APIs, testar ambos os cenários (sucesso e erro)
    - **DAO**: Usar banco de dados em memória, testar operações CRUD
    - **Provider (Riverpod)**: Usar `ProviderContainer` com overrides, testar transições de estado
    - **Serviço**: Mockar dependências, focar na lógica de negócio
    - **Widget**: Usar keys para identificação estável, testar interações e estados
4.  **Testes de UI**: Configurar dimensões de tela explícitas, usar `tester.pumpAndSettle()` para animações
5.  **Qualidade Incremental**: Rodar `fvm flutter analyze` e `fvm flutter test` periodicamente.
6.  **Análise de Segurança**: Verificar vulnerabilidades OWASP Mobile Top 10 em código e configurações.

### Fase 4: DEPLOY — Build, Segurança e Distribuição
1.  **Análise de Segurança**: Verificar vulnerabilidades antes do deploy:
    - **M1: Credenciais hardcoded**: Escanear código Dart e configurações
    - **M2: Dependências inseguras**: Verificar `pubspec.yaml` por pacotes desatualizados ou com `any`
    - **M5: Segurança de rede**: Verificar uso de HTTPS, certificate pinning
    - **M9: Armazenamento inseguro**: Verificar SharedPreferences não criptografados
2.  **Limpeza**: Executar `fvm flutter clean` antes de builds de release.
3.  **Build Seguro**: Executar builds com proteções:
    ```bash
    # Android com ofuscação
    fvm flutter build apk --release --obfuscate --split-debug-info=./build/symbols/
    
    # iOS
    fvm flutter build ios --release
    ```
4.  **CI/CD Seguro**: Utilizar o `fvm_config.json` para garantir que o runner de CI use a mesma versão do SDK, e incluir verificações de segurança no pipeline.

---

## FVM vs Flutter Commands

| Comando Tradicional | Equivalente FVM (Recomendado) | Propósito |
|---------------------|-------------------------------|-----------|
| `flutter run`       | `fvm flutter run`             | Executar o app |
| `flutter pub get`   | `fvm flutter pub get`         | Instalar dependências |
| `flutter test`      | `fvm flutter test`            | Rodar testes |
| `flutter analyze`   | `fvm flutter analyze`         | Análise estática |
| `flutter build ...` | `fvm flutter build ...`       | Gerar build de release |
| `dart format .`     | `fvm dart format .`           | Formatar código |

---

## Best Practices

### Project Management

**DO:**
- Use `fvm use` para definir a versão do SDK (isso cria o link `.fvm/flutter_sdk`)
- Adicione `.fvm/flutter_sdk` ao seu `.gitignore`
- Versionar `.fvm/fvm_config.json` no git
- Sempre use `fvm flutter` em vez de `flutter` diretamente no terminal
- Configure seu `analysis_options.yaml` logo no início do projeto
- Implemente testes por camada arquitetural (Repository, DAO, Provider, Service, Widget)
- Use keys para identificação estável em testes de widget
- Escanear regularmente por vulnerabilidades OWASP Mobile Top 10
- Usar builds ofuscados para release (`--obfuscate --split-debug-info`)

**DON'T:**
- Não versione a pasta `.fvm/flutter_sdk` (é um link simbólico ou cache local)
- Não misture comandos `flutter` globais e `fvm flutter` no mesmo projeto
- Não ignore erros do `fvm flutter analyze`
- Não armazene credenciais hardcoded no código Dart
- Não use dependências com versionamento `any` no `pubspec.yaml`
- Não use HTTP para chamadas de API em produção
- Não armazene dados sensíveis em SharedPreferences sem criptografia

---

## Reference Documentation

Esta skill inclui documentação de referência detalhada:

1. **[Environment Setup](references/environment-setup.md)** — Instalação do FVM, PATH, e fixação de versões.
2. **[Project Structure](references/project-structure.md)** — Organização de pastas, `pubspec.yaml` e `analysis_options.yaml`.
3. **[Dependency Management](references/dependency-management.md)** — Pub, versionamento semântico e `pubspec.lock`.
4. **[Testing & Quality](references/testing-and-quality.md)** — Unit tests, Widget tests, Integration tests e Coverage.
5. **[Advanced Testing Patterns](references/advanced-testing-patterns.md)** — Testes por camada arquitetural, Riverpod testing, Widget testing com keys.
6. **[Flutter Security Guide](references/flutter-security-guide.md)** — OWASP Mobile Top 10, análise de dependências, builds seguros.
7. **[Deployment Guide](references/deployment-guide.md)** — Build commands para Android, iOS e Web.
8. **[IDE Integration](references/ide-integration.md)** — Configuração do VS Code e Android Studio para usar o SDK do FVM.

---

## Output Structure

A execução desta skill em projetos Flutter deve resultar nos seguintes artefatos:

| Artefato | Arquivo | Descrição |
|----------|---------|-----------|
| **FVM Config** | `.fvm/fvm_config.json` | Define a versão do Flutter SDK para o projeto. |
| **Project Spec** | `pubspec.yaml` | Metadados, dependências e assets do projeto. |
| **Lockfile** | `pubspec.lock` | Registro determinístico de versões de pacotes. |
| **Lint Rules** | `analysis_options.yaml` | Configurações do linter e analisador estático. |
| **SDK Link** | `.fvm/flutter_sdk` | Link simbólico (ou diretório) para o SDK gerenciado pelo FVM. |

---

## Quality Rules

- **FVM-First**: Sempre preferir `fvm flutter` sobre `flutter`.
- **Análise Estática**: Nenhum commit deve ser feito com erros de `analyze`.
- **Testes**: Promover a escrita de testes para lógica de negócio (Bloc/Cubit/Providers).
- **Clean Code**: Seguir as [Dart Style Guidelines](https://dart.dev/guides/language/effective-dart/style).

## Prohibited

- **Não usar Flutter Global**: Nunca sugerir o uso do `flutter` global em projetos que possuem FVM configurado.
- **Não ignorar o Analyze**: Nunca ignorar warnings ou erros de análise estática.
- **Não versionar o SDK**: Nunca adicionar `.fvm/flutter_sdk` ao controle de versão.
- **Não pular Testes**: Nunca considerar uma feature completa sem a devida cobertura de testes de ViewModel/Logic.
