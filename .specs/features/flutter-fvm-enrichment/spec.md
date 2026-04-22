# Feature Spec: Flutter FVM Skill Enrichment (v1.3.0)

## Overview
Enriquecer a skill `flutter-fvm` com padrões avançados de desenvolvimento Flutter/Dart, performance, acessibilidade e resiliência, baseados no ecossistema ECC.

## Requirements (Functional)
- **FR-1**: Suporte a Dart 3+ (Sealed Classes, Pattern Matching, Records).
- **FR-2**: Diretrizes de Acessibilidade (Semantics, Hit Targets).
- **FR-3**: Otimização de Performance (Selective Rebuilds, RepaintBoundary).
- **FR-4**: Protocolos de Error Handling global.
- **FR-5**: Gestão segura de Secrets via `--dart-define`.

## Acceptance Criteria (AC) - BDD Format

### AC 1: Modernização Dart 3+ (Sealed Classes)
**Given** que o desenvolvedor precisa modelar estados de autenticação
**When** ele consultar a skill `flutter-fvm`
**Then** ele deve encontrar instruções para usar `sealed class` para representar estados mutuamente exclusivos (Initial, Loading, Authenticated, Error)
**And** um exemplo de switch expression para tratar esses estados de forma exaustiva.

### AC 2: Acessibilidade (TalkBack/VoiceOver)
**Given** que um widget interativo está sendo criado
**When** a skill for aplicada
**Then** deve ser exigido o uso de `Semantics` com labels descritivos
**And** a área de toque deve ser validada para o mínimo de 48x48 dp.

### AC 3: Performance (Selective Rebuilds)
**Given** um widget que depende apenas do tamanho da tela
**When** o desenvolvedor utilizar a skill
**Then** ele deve ser instruído a usar `MediaQuery.sizeOf(context)` em vez de `MediaQuery.of(context)` para evitar rebuilds desnecessários.

### AC 4: Segurança (Secrets)
**Given** a necessidade de integrar uma API Key
**When** a skill for consultada
**Then** deve ser proibido o hardcoding da chave
**And** deve ser recomendado o uso de `--dart-define` ou `.env` via `flutter_dotenv`.

### AC 5: Resiliência (Global Catchers)
**Given** uma aplicação Flutter profissional
**When** configurada via skill
**Then** deve existir um hook no `main.dart` para `FlutterError.onError` e `PlatformDispatcher.instance.onError`.

## Constraints
- **C-1**: Compatibilidade com FVM mandatória.
- **C-2**: Nenhuma quebra de retrocompatibilidade com as versões 1.1.x e 1.2.x.
- **C-3**: Documentação em Português Brasileiro (conforme regra do usuário).
