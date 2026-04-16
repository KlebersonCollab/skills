# Advanced Testing Patterns for Flutter

Este guia apresenta padrões avançados de teste para aplicações Flutter, inspirados nas melhores práticas de testes por camada arquitetural.

## 1. Testes por Camada Arquitetural

### Camada de Repositório
Os repositórios coordenam entre DAOs e APIs. Teste ambos os cenários de sucesso e erro:

```dart
test('repository should return data when API succeeds', () async {
  // Arrange
  final mockApi = MockApi();
  final mockDao = MockDao();
  final repository = MyRepository(api: mockApi, dao: mockDao);
  
  when(() => mockApi.fetchData()).thenAnswer((_) async => apiData);
  
  // Act
  final result = await repository.getData();
  
  // Assert
  expect(result, equals(expectedData));
  verify(() => mockDao.saveData(any())).called(1);
});

test('repository should return cached data when API fails', () async {
  // Arrange
  final mockApi = MockApi();
  final mockDao = MockDao();
  final repository = MyRepository(api: mockApi, dao: mockDao);
  
  when(() => mockApi.fetchData()).thenThrow(NetworkException());
  when(() => mockDao.getCachedData()).thenReturn(cachedData);
  
  // Act
  final result = await repository.getData();
  
  // Assert
  expect(result, equals(cachedData));
  verifyNever(() => mockDao.saveData(any()));
});
```

### Camada de DAO
DAOs interagem com o banco de dados. Use FakeDatabase para testes realistas:

```dart
test('dao should insert and retrieve data', () async {
  // Arrange
  final database = FakeDatabase();
  final dao = MyDao(database);
  
  // Act
  await dao.insert(item);
  final retrieved = await dao.getById(item.id);
  
  // Assert
  expect(retrieved, equals(item));
});

test('dao should update existing item', () async {
  // Arrange
  final database = FakeDatabase();
  final dao = MyDao(database);
  await dao.insert(originalItem);
  
  // Act
  await dao.update(updatedItem);
  final retrieved = await dao.getById(originalItem.id);
  
  // Assert
  expect(retrieved, equals(updatedItem));
});
```

### Camada de Provider (Riverpod)
Providers gerenciam estado. Teste transições de estado e fluxos async:

```dart
test('provider should emit loading then data', () async {
  // Arrange
  final container = ProviderContainer(
    overrides: [
      myRepositoryProvider.overrideWithValue(MockRepository()),
    ],
  );
  
  // Act & Assert
  expect(
    container.read(myProvider),
    const AsyncValue<MyData>.loading(),
  );
  
  await container.pump();
  
  expect(
    container.read(myProvider).value,
    equals(expectedData),
  );
});

test('provider should handle errors', () async {
  // Arrange
  final mockRepository = MockRepository();
  when(() => mockRepository.getData()).thenThrow(Exception());
  
  final container = ProviderContainer(
    overrides: [
      myRepositoryProvider.overrideWithValue(mockRepository),
    ],
  );
  
  // Act & Assert
  await container.pump();
  
  expect(
    container.read(myProvider).error,
    isA<Exception>(),
  );
});
```

### Camada de Serviço
Serviços contêm lógica de negócio. Teste a lógica, não a implementação:

```dart
test('service should apply business rules', () async {
  // Arrange
  final mockRepository = MockRepository();
  final service = MyService(repository: mockRepository);
  
  when(() => mockRepository.getUserData()).thenReturn(userData);
  
  // Act
  final result = await service.processUserData();
  
  // Assert
  expect(result.isValid, isTrue);
  expect(result.processedAt, isNotNull);
});

test('service should cache results', () async {
  // Arrange
  final mockRepository = MockRepository();
  final service = MyService(repository: mockRepository);
  
  when(() => mockRepository.getExpensiveData())
    .thenAnswer((_) async => expensiveData);
  
  // Act - First call
  final result1 = await service.getData();
  
  // Second call should use cache
  final result2 = await service.getData();
  
  // Assert
  expect(result1, equals(result2));
  verify(() => mockRepository.getExpensiveData()).called(1);
});
```

## 2. Testes de Widget com Keys

Use keys para identificação estável de widgets em testes:

```dart
testWidgets('should show loading then data', (tester) async {
  // Arrange
  await tester.pumpWidget(
    ProviderScope(
      overrides: [
        myProvider.overrideWithValue(
          const AsyncValue.loading(),
        ),
      ],
      child: const MyWidget(),
    ),
  );
  
  // Assert - Loading state
  expect(find.byKey(const Key('loadingIndicator')), findsOneWidget);
  expect(find.byKey(const Key('dataList')), findsNothing);
  
  // Act - Switch to data state
  await tester.pumpWidget(
    ProviderScope(
      overrides: [
        myProvider.overrideWithValue(
          AsyncValue.data(testData),
        ),
      ],
      child: const MyWidget(),
    ),
  );
  
  // Assert - Data state
  expect(find.byKey(const Key('loadingIndicator')), findsNothing);
  expect(find.byKey(const Key('dataList')), findsOneWidget);
});

testWidgets('should handle user interactions', (tester) async {
  // Arrange
  await tester.pumpWidget(const MyApp());
  
  // Act - Tap button
  await tester.tap(find.byKey(const Key('saveButton')));
  await tester.pump();
  
  // Assert
  expect(find.text('Saved successfully'), findsOneWidget);
});
```

## 3. Padrões de Mocking e Verificação

### Use @GenerateMocks para tipos complexos
```dart
@GenerateMocks([ApiClient, Database])
void main() {
  late MockApiClient mockApi;
  late MockDatabase mockDb;
  
  setUp(() {
    mockApi = MockApiClient();
    mockDb = MockDatabase();
  });
  
  // ... tests
}
```

### Verificação de chamadas em ordem
```dart
test('should call methods in correct order', () {
  // Arrange
  final mockA = MockA();
  final mockB = MockB();
  final sut = SystemUnderTest(a: mockA, b: mockB);
  
  // Act
  sut.performOperation();
  
  // Assert
  verifyInOrder([
    () => mockA.prepare(),
    () => mockB.execute(),
    () => mockA.finalize(),
  ]);
});
```

### Teste de cenários de erro
```dart
test('should log error when operation fails', () {
  // Arrange
  final mockLogger = MockLogger();
  final mockService = MockService();
  
  when(() => mockService.performRiskyOperation())
    .thenThrow(Exception('Operation failed'));
  
  final sut = BusinessLogic(
    service: mockService,
    logger: mockLogger,
  );
  
  // Act
  expect(() => sut.execute(), throwsA(isA<BusinessException>()));
  
  // Assert
  verify(() => mockLogger.error(
    any(that: contains('Operation failed')),
  )).called(1);
});
```

## 4. Configuração de Testes de Plataforma

Teste comportamentos específicos de plataforma:

```dart
testWidgets('should show Cupertino style on iOS', (tester) async {
  // Arrange
  debugDefaultTargetPlatformOverride = TargetPlatform.iOS;
  
  // Act
  await tester.pumpWidget(const MyPlatformAwareWidget());
  
  // Assert
  expect(find.byType(CupertinoButton), findsOneWidget);
  expect(find.byType(ElevatedButton), findsNothing);
});

testWidgets('should show Material style on Android', (tester) async {
  // Arrange
  debugDefaultTargetPlatformOverride = TargetPlatform.android;
  
  // Act
  await tester.pumpWidget(const MyPlatformAwareWidget());
  
  // Assert
  expect(find.byType(CupertinoButton), findsNothing);
  expect(find.byType(ElevatedButton), findsOneWidget);
});
```

**Importante**: Sempre resetar após testes:
```dart
tearDown(() {
  debugDefaultTargetPlatformOverride = null;
});
```

## 5. Testes de Navegação com GoRouter

```dart
testWidgets('should navigate to details screen', (tester) async {
  // Arrange
  final mockGoRouter = MockGoRouter();
  
  await tester.pumpWidget(
    ProviderScope(
      overrides: [
        goRouterProvider.overrideWithValue(mockGoRouter),
      ],
      child: const MyApp(),
    ),
  );
  
  // Act
  await tester.tap(find.byKey(const Key('item_0')));
  await tester.pumpAndSettle();
  
  // Assert
  verify(() => mockGoRouter.push('/details/0')).called(1);
});
```

## 6. Checklist de Testes Avançados

- [ ] **Testes por camada**: Cada camada arquitetural testada isoladamente
- [ ] **Keys em widgets**: Identificação estável para testes de UI
- [ ] **Verificação de ordem**: Métodos chamados na sequência correta
- [ ] **Cenários de erro**: Teste de falhas e recuperação
- [ ] **Plataforma específica**: Teste de comportamentos por plataforma
- [ ] **Navegação**: Teste de rotas e transições
- [ ] **Estado async**: Teste de loading, data, error states
- [ ] **Lifecycle**: Teste de initState, dispose, etc.
- [ ] **Performance**: Teste de operações assíncronas longas
- [ ] **Acessibilidade**: Teste de semântica e labels