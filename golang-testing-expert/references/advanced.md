# Pilar 2: Advanced Testing

> Testes avançados envolvem o isolamento de dependências externas e a validação de interfaces complexas como APIs HTTP.

---

## 🎭 Mocking & Dependency Injection

Em Go, a melhor forma de mockar é através de interfaces. Crie mocks leves manualmente para evitar dependências pesadas de frameworks quando possível.

```go
type MockStore struct {
    GetFunc func(id string) (User, error)
}

func (m *MockStore) Get(id string) (User, error) {
    return m.GetFunc(id)
}
```

---

## 🌐 HTTP Testing (Package `httptest`)

Utilize `net/http/httptest` para testar handlers sem precisar subir um servidor real.

- **httptest.NewRecorder()**: Captura a resposta do handler.
- **httptest.NewRequest()**: Cria uma requisição mockada.
- **httptest.NewServer()**: Sobe um servidor real em uma porta aleatória (ideal para testar clientes HTTP).

```go
func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/health", nil)
    w := httptest.NewRecorder()
    HealthHandler(w, req)
    
    if w.Code != http.StatusOK {
        t.Errorf("got status %d", w.Code)
    }
}
```

---

## 🏆 Golden Files

Para testar saídas grandes (JSON complexo, HTML, arquivos gerados), salve o "resultado esperado" em um arquivo na pasta `testdata/`.

- Compare o resultado atual com o conteúdo do arquivo golden.
- Utilize uma flag (`-update`) para atualizar os arquivos golden quando o comportamento mudar intencionalmente.

---

## 📦 TestData & Fixtures

- Utilize a pasta `testdata/` para armazenar arquivos estáticos de teste. O Go ignora esta pasta durante a compilação.
- Prefira injetar dados via código (Factory Pattern) do que depender de estados complexos em arquivos externos quando possível.
