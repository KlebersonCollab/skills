# Pilar 1: Testing Foundations

> A base de uma aplicação resiliente em Go é uma suíte de testes que documenta o comportamento esperado e garante a estabilidade durante refatorações.

---

## 🔴🟢🔵 Ciclo TDD (Red-Green-Refactor)

O desenvolvimento dirigido por testes é o padrão de ouro para evitar bugs e over-engineering:

1. **RED**: Defina a assinatura da função e escreva um teste que falha (pode usar `panic("not implemented")`).
2. **GREEN**: Implemente o código mais simples que satisfaça o teste.
3. **REFACTOR**: Limpe o código, melhore nomes e abstrações.

---

## 📊 Table-Driven Tests (Padrão Idiomático)

Este é o padrão preferido em Go para testar múltiplas entradas e saídas sem duplicar lógica de teste.

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -1, -2, -3},
        {"zero values", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
            }
        })
    }
}
```

---

## 🏃 Subtests & Paralelismo

- **t.Run**: Permite agrupar testes relacionados e rodar casos específicos via CLI (`go test -run TestName/SubtestName`).
- **t.Parallel**: Indica que o teste pode rodar em paralelo com outros testes marcados com `t.Parallel()`. 
  > **Nota**: Sempre capture a variável de loop quando usar `t.Parallel()` dentro de um `range`.

```go
for _, tt := range tests {
    tt := tt // Captura necessária antes do Go 1.22
    t.Run(tt.name, func(t *testing.T) {
        t.Parallel()
        // ... teste aqui
    })
}
```

---

## 🧹 Helpers & Cleanup

- **t.Helper()**: Marca a função como auxiliar para que, em caso de erro, o log aponte para a linha do teste que chamou a função, e não para dentro do helper.
- **t.Cleanup()**: Define uma função para limpar recursos (fechar DB, deletar arquivos) após a conclusão do teste ou subtest.
- **t.TempDir()**: Cria um diretório temporário que é automaticamente removido após o teste.

```go
func setupDB(t *testing.T) *DB {
    t.Helper()
    db := connect()
    t.Cleanup(func() { db.Close() })
    return db
}
```
