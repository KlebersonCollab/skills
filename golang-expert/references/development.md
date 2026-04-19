# Pilar 3: Development & Tooling

> O desenvolvimento em Go é potencializado por ferramentas integradas que garantem a qualidade e a agilidade do ciclo de vida do software.

---

## 💻 CLI Applications (Linha de Comando)

Go é a linguagem de escolha para ferramentas CLI modernas:
- **Flag Package**: Utilize o pacote nativo `flag` para aplicações simples.
- **Cobra & Viper**: Utilize bibliotecas como `spf13/cobra` e `spf13/viper` para aplicações complexas com subcomandos e configurações flexíveis.
- **Interatividade**: Forneça feedback visual claro e suporte a pipes de entrada/saída.

---

## 📦 Dependency Management (Módulos Go)

O sistema de módulos é a base da reprodutibilidade:
- **go mod tidy**: Mantenha o arquivo `go.mod` e `go.sum` sempre limpos e sincronizados.
- **Vendoring**: Considere `go mod vendor` em ambientes corporativos onde o acesso direto à internet é restrito no CI/CD.
- **Semantic Versioning**: Siga rigorosamente o versionamento semântico ao publicar pacotes públicos.

---

## 🔍 Linting & Static Analysis (Análise de Código)

A análise estática é mandatória para um código Expert:
- **golangci-lint**: O orquestrador padrão de linters em Go. Configure-o para rodar no CI.
- **Gofmt / Goimports**: Formatação automática e gestão inteligente de pacotes importados.
- **Custom Linters**: Utilize linters específicos para segurança (`gosec`) e bugs lógicos (`errcheck`, `staticcheck`).

---

## 🧪 Testing Strategies (Testes Expert)

Em Go, o código de teste é código de primeira classe:
- **Native Testing**: Utilize o pacote `testing` e a estrutura `_test.go`.
- **Table-Driven Tests**: Padronize seus casos de teste em tabelas para maior legibilidade e cobertura.
  ```go
  func TestSum(t *testing.T) {
      cases := []struct { a, b, want int }{ {1, 2, 3}, {4, 5, 9} }
      for _, c := range cases {
          if got := Sum(c.a, c.b); got != c.want { ... }
      }
  }
  ```
- **Testify**: Utilize a biblioteca `stretchr/testify` para asserções ricas e mocks robustos.
- **Sub-tests**: Utilize `t.Run` para isolar casos de teste dentro de uma mesma função.

---

## 🛠️ Troubleshooting & Debugging (Resolução de Problemas)

- **Delve (dlv)**: O debugger padrão para Go. Utilize-o para inspeção profunda de memória e execução.
- **Stack Traces**: Entenda como ler e interpretar traces de pânico e dumps de goroutines.
- **Profiling**: Utilize `pprof` para identificar gargalos de CPU, memória e contenção de travas.

---

## 📚 Bibliotecas de Apoio (Development)
- `stretchr/testify`: Para asserções e mocks.
- `samber/lo`: Para utilitários de manipulação de slices e maps em testes.
