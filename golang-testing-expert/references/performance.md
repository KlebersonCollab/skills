# Pilar 3: Performance & Fuzzing

> Validar não apenas a correção lógica, mas também a eficiência e a resiliência contra entradas inesperadas.

---

## ⏱️ Benchmarking

O Go possui suporte nativo a benchmarks integrados ao comando `go test`.

- Funções devem começar com `Benchmark`.
- Utilize o loop `b.N` para execuções repetidas controladas pelo runtime.
- **b.ResetTimer()**: Use após setups pesados para não distorcer o tempo de execução.
- **b.ReportAllocs()**: Ativa o relatório de alocações de memória por operação.

```go
func BenchmarkProcess(b *testing.B) {
    data := setup()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        Process(data)
    }
}
```

---

## 🧪 Fuzz Testing (Go 1.18+)

O Fuzzing gera inputs aleatórios para encontrar falhas de segurança ou crashes que testes manuais não pegariam.

- Utilize `f.Add()` para fornecer exemplos iniciais (seed corpus).
- O Go tentará mutar esses inputs para quebrar seu código.

```go
func FuzzParse(f *testing.F) {
    f.Add("input inicial")
    f.Fuzz(func(t *testing.T, str string) {
        Parse(str) // Deve ser resiliente a qualquer string
    })
}
```

---

## 📊 Test Coverage (Cobertura)

- **go test -cover**: Exibe a porcentagem de cobertura no console.
- **go test -coverprofile=c.out**: Gera um arquivo de profile.
- **go tool cover -html=c.out**: Abre um visualizador no browser destacando as linhas não testadas.

---

## 🏁 Race Detection

Sempre rode seus testes com a flag `-race` em ambientes de CI e durante o desenvolvimento local de código concorrente.

```bash
go test -race ./...
```
