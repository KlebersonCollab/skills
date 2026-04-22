# Testes & Qualidade (Mobile)

Qualidade no Flutter é uma combinação de análise estática rigorosa e cobertura de testes por camadas.

## 1. Pirâmide de Testes
- **Unit Tests**: Lógica pura e ViewModels.
- **Widget Tests**: Interações e Acessibilidade.
- **Golden Tests**: Regressão visual.
- **Integration Tests**: Fluxos E2E.

## 2. Testes de Acessibilidade
**Conceito**: Validar se a árvore semântica está correta.

### ✅ Bom (Semantic Test)
```dart
testWidgets('Botão deve ser acessível', (tester) async {
  final handle = tester.ensureSemantics();
  await tester.pumpWidget(const MyButton());

  expect(tester.getSemantics(find.byType(MyButton)), matchesSemantics(
    label: 'Enviar',
    isButton: true,
  ));
  handle.dispose();
});
```

## 3. Análise Estática (Strict)
Sempre execute `fvm flutter analyze` antes de qualquer push.

### ✅ Bom (`analysis_options.yaml`)
```yaml
analyzer:
  errors:
    unused_local_variable: error
    invalid_annotation_target: ignore
  language:
    strict-casts: true
    strict-inference: true
```

## Checklist de Qualidade
- [ ] Rodou `fvm flutter analyze` e corrigiu todos os warnings?
- [ ] A lógica de negócio (ViewModel) tem 80%+ de cobertura?
- [ ] Widgets críticos têm testes de Semântica?
- [ ] Goldens foram atualizados para a plataforma alvo?
- [ ] Usou `mocktail` para isolar dependências externas?
