---
name: flutter-fvm
version: 1.3.0
description: "Skill para desenvolvimento Flutter profissional com FVM. Inclui padrões Dart 3+, performance otimizada, acessibilidade mandatória, segurança OWASP e gestão de resiliência (Error Handling)."
category: development-workflow
---

# Flutter com FVM & Modern Patterns

> Gerenciador de versões do Flutter combinado com padrões modernos de arquitetura, performance e acessibilidade para aplicações de alta fidelidade.

---

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar qualquer execução técnica:
1. **Context Check**: Você reidratou o contexto lendo `STATE.md`, `MEMORY.md` e `LEARNINGS.md`?
2. **Spec Check**: O arquivo `spec.md` existe com requisitos e Critérios de Aceitação (ACs) claros? (BDD mandatório para Medium+).
3. **Plan Check**: O arquivo `plan.md` define a arquitetura, schemas e inclui diagramas **Mermaid**?
4. **Contract Check**: O arquivo `contract.md` foi estabelecido com os sensores de validação?
5. **Task Check**: A lista de tarefas em `tasks.md` está detalhada e atomizada?

---
## Goal

**FVM** garante o ambiente. Esta skill garante a **Excelência Técnica** através de:

| Pilar | Foco Principal |
|-------|----------------|
| **Dart 3+** | Sealed classes, Pattern Matching e Records para segurança de tipos. |
| **Performance** | Otimização de rebuilds (`MediaQuery.sizeOf`), RepaintBoundaries e Decomposição. |
| **UX/Acessibilidade** | Semantics, alvos de toque (48x48) e contraste dinâmico. |
| **Resiliência** | Captura global de erros e tratamento gracioso de falhas. |
| **Segurança** | OWASP Mobile Top 10 e gestão segura de Secrets (`--dart-define`). |

## Version Awareness

**Versão Recomendada: Flutter Stable (última disponível via FVM)**

Verifique as versões instaladas e a ativa no projeto:

```bash
fvm list
fvm current
```

## When to Use This Skill

Use this skill when:
- Inicializando um novo projeto Flutter ou migrando para Dart 3+.
- Implementando lógica de negócio complexa com Gerenciamento de Estado (Riverpod/Bloc).
- Otimizando a performance de telas com listas complexas ou animações.
- Garantindo que o app seja acessível via leitores de tela (TalkBack/VoiceOver).
- Configurando tratamento de erros robusto e crash reporting.
- Gerenciando segredos e chaves de API com segurança.
- Realizando análise de segurança OWASP Mobile Top 10.

Skip this skill when:
- O projeto NÃO usa Flutter/Dart.

---

## Workflow (4 Fases)

### Fase 1: ENVIRONMENT — Configuração e Alinhamento
1.  **Versão SDK**: Fixar com `fvm use stable`.
2.  **Linting Rigoroso**: Configurar `analysis_options.yaml` com `strict-casts: true`, `strict-inference: true` e `strict-raw-types: true`.

### Fase 2: PROJECT — Estrutura e Modernização
1.  **State Modeling**: Usar `sealed class` para representar estados (Initial, Loading, Data, Error).
2.  **Secrets Management**: Criar estrutura para receber chaves via `--dart-define` ou `.env` (excluído do git).
3.  **Error Boundaries**: Configurar `FlutterError.onError` no `main.dart`.

### Fase 3: DEVELOP — Qualidade e Performance
1.  **Decomposição**: Manter métodos `build()` abaixo de 100 linhas.
2.  **Selective Rebuilds**: Preferir `MediaQuery.sizeOf(context)` sobre `MediaQuery.of(context)`.
3.  **Acessibilidade**: Adicionar `semanticLabel` em todos os elementos visuais críticos.
4.  **Testes**: Cobrir transições de estado selados e fluxos de erro.

### Fase 4: DEPLOY — Segurança e Resiliência
1.  **Crash Reporting**: Garantir integração com Sentry/Crashlytics.
2.  **Hardening**: Rodar builds com `--obfuscate --split-debug-info`.
3.  **Validation**: Verificar se nenhum `print` restou (usar `debugPrint` ou `log`).

## Output Structure

A execução desta skill resulta na geração ou atualização dos seguintes artefatos:

| Artefato | Localização | Descrição |
|----------|-------------|-----------|
| **FVM Config** | `.fvm/fvm_config.json` | Definição da versão do SDK para o time. |
| **Analysis Options** | `analysis_options.yaml` | Regras de linting e análise estática. |
| **Sealed States** | `lib/features/*/domain/` | Modelagem de estados exaustivos. |
| **Secure Secrets** | `.env` / `main.dart` | Gestão de chaves via environment. |

---

## Quality Rules

**DO:**
- Use `sealed` classes para garantir exaustividade no tratamento de estados.
- Use `const` construtores agressivamente para reduzir pressão no GC.
- Garanta alvos de toque de no mínimo 48x48 dp.
- Use `RepaintBoundary` em animações ou widgets de desenho pesado.
- Valide `context.mounted` após qualquer `await`.

**DON'T:**
- Não use booleanos múltiplos (`isLoading`, `hasError`) para estados de tela.
- Não use `print()` em produção; use o pacote `logging` ou `dart:developer`.
- Não ignore erros do `fvm flutter analyze`.
- Não armazene segredos diretamente no código Dart (hardcoded).

---

## Reference Documentation

1. **[Padrões Modernos de Dart](references/modern-dart-patterns.md)** — Sealed classes, Records e Pattern Matching.
2. **[Performance & Otimização](references/performance-and-optimization.md)** — Rebuilds, RepaintBoundaries e GC.
3. **[Guia de Acessibilidade](references/accessibility-guide.md)** — Semantics, Hit Targets e UX Inclusiva.
4. **[Configuração do Ambiente](references/environment-setup.md)** — FVM e Strict Analyzer.
5. **[Guia de Segurança Flutter](references/flutter-security-guide.md)** — OWASP e Secrets.
6. **[Testes & Qualidade](references/testing-and-quality.md)** — State and Logic testing.

---

## Prohibited

- **Não pular Acessibilidade**: Features sem labels semânticos ou alvos pequenos são consideradas incompletas.
- **Não ignorar o Analyzer**: Erros de análise impedem o deploy.
- **Não usar o Flutter Global**: Mantenha o isolamento do projeto via FVM.

---
*Atualizada em 22 de Abril de 2026 para alinhamento com padrões ECC.*
