# Testing & Quality: Flutter with FVM

Este guia foca em como garantir a qualidade do código e testar sua aplicação Flutter utilizando o workflow do FVM.

## 1. Pirâmide de Testes no Flutter

No Flutter, os testes são divididos em três categorias principais:

- **Unit Tests**: Testam uma única função, método ou classe. (Localizados em `test/`)
- **Widget Tests**: Testam a UI e a interação com widgets isolados. (Localizados em `test/`)
- **Integration Tests**: Testam o app completo em um device real ou emulador. (Localizados em `integration_test/`)

## 2. Executando Testes via FVM

Sempre execute os testes através do FVM para garantir que a mesma versão do SDK do time está sendo usada:

```bash
# Executar todos os testes unitários e de widget
fvm flutter test

# Executar um arquivo de teste específico
fvm flutter test test/meu_widget_test.dart

# Gerar relatório de cobertura de testes (LCOV)
fvm flutter test --coverage
```

## 3. Qualidade de Código (Linter)

O analisador estático do Dart é uma ferramenta poderosa. Configure regras no seu `analysis_options.yaml` e execute a análise frequentemente:

```bash
# Rodar análise estática
fvm flutter analyze
```

Se houver erros de análise, o comando retornará um exit code diferente de zero, o que é ideal para o CI/CD.

## 4. Formatação de Código

Manter o estilo consistente é vital. Use o formatador nativo do Dart:

```bash
# Formatar todos os arquivos no diretório atual
fvm dart format .

# Verificar se os arquivos estão formatados (sem aplicar as mudanças)
fvm dart format --set-exit-if-changed .
```

## 5. Cobertura de Testes (Coverage)

Para visualizar a cobertura de testes de forma amigável:

1. Gere o arquivo `lcov.info`:
   ```bash
   fvm flutter test --coverage
   ```
2. Visualize usando ferramentas como `lcov` (macOS via Homebrew: `brew install lcov`):
   ```bash
   genhtml coverage/lcov.info -o coverage/html
   open coverage/html/index.html
   ```

## 6. Mocks e Fakes

Para isolar seus testes, utilize pacotes de mocking:
- **mocktail**: Inspirado no Mockito, mas com uma sintaxe mais moderna para Dart.
- **mockito**: O clássico (requer geração de código).

**Dica**: Prefira `mocktail` por não precisar de geração de código para testes simples.
