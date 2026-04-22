# Guia de Acessibilidade (UX Inclusiva)

Apps profissionais devem ser usáveis por todos, incluindo usuários de leitores de tela e pessoas com baixa visão.

## 1. Semantics
**Conceito**: Fornecer metadados para que o TalkBack/VoiceOver consiga descrever a UI.

### ❌ Ruim
```dart
GestureDetector(
  onTap: () => print('Play'),
  child: Icon(Icons.play_arrow), // O leitor dirá apenas "Ícone" ou nada.
)
```

### ✅ Bom
```dart
Semantics(
  label: 'Botão Reproduzir Vídeo',
  button: true,
  child: IconButton(
    icon: Icon(Icons.play_arrow),
    onPressed: () {},
  ),
)
```

## 2. Alvos de Toque (Hit Targets)
**Conceito**: Área mínima necessária para interação física confortável.

### ❌ Ruim
```dart
SizedBox(
  width: 20,
  height: 20,
  child: InkWell(onTap: () {}), // Muito difícil de acertar com o dedo
)
```

### ✅ Bom (48x48 dp)
```dart
ConstrainedBox(
  constraints: BoxConstraints.tightFor(width: 48, height: 48),
  child: InkWell(onTap: () {}),
)
```

## 3. Contraste de Texto
**Conceito**: Legibilidade baseada na relação entre cor de fundo e cor da fonte.

### ✅ Bom
- Contraste mínimo de **4.5:1**.
- Não use apenas a cor para indicar erro (ex: use um ícone de aviso ⚠️ junto com a borda vermelha).

## Checklist de Acessibilidade
- [ ] Todos os ícones interativos têm `semanticLabel` ou estão dentro de `Semantics`?
- [ ] Alvos de toque têm no mínimo 48x48 dp?
- [ ] O app foi testado com o tamanho de fonte do sistema no máximo?
- [ ] Usou `MergeSemantics` para agrupar rótulos e valores?
