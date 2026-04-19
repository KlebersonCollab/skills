# Pubspec Standards Guide

Este guia define os padrões de organização, versionamento e manutenção do arquivo `pubspec.yaml` em projetos que utilizam o ecossistema do Hub.

## 1. Organização de Seções
O arquivo deve seguir a ordem padrão do SDK do Flutter, mantendo seções claras:
1. Meta-informações (name, description, version, etc.)
2. Environment (sdk, flutter version)
3. Dependencies (ordenadas alfabeticamente)
4. Dev Dependencies (ordenadas alfabeticamente)
5. Flutter Assets/Fonts/etc.

## 2. Regras de Versionamento
- **Dependências de Produção**: Use o prefixo careta (`^`) para permitir atualizações de patch e minor, garantindo compatibilidade.
- **Dependências Críticas**: Fixe a versão (ex: `1.2.3`) apenas em casos de breaking changes conhecidas.
- **FVM**: Utilize `fvm flutter pub get` para garantir que as versões das dependências sejam resolvidas com o SDK correto.

## 3. Comentários e Agrupamento
Agrupe dependências relacionadas com comentários:
```yaml
dependencies:
  # Core
  flutter:
    sdk: flutter
  cupertino_icons: ^1.0.2

  # State Management
  flutter_bloc: ^8.1.0
  riverpod: ^2.3.0

  # Network & API
  dio: ^5.1.0
  json_annotation: ^4.8.0
```

## 4. Manutenção
- Execute `flutter pub outdated` mensalmente.
- Evite dependências não utilizadas (DRY).
- Sempre verifique a pontuação no Pub.dev antes de adicionar novas bibliotecas.
