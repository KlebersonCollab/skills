# Pilar 1: Foundations & Style

> A base de um código Expert em Go reside na sua simplicidade e na adesão rigorosa aos padrões idiomáticos da linguagem.

---

## 🏗️ Project Layout (Organização de Código)

Seguir o padrão `golang-standards/project-layout` para projetos de grande escala:
- `/cmd`: Ponto de entrada das aplicações (ex: `cmd/api/main.go`).
- `/internal`: Código privado que não deve ser importado por outros projetos.
- `/pkg`: Bibliotecas reutilizáveis que podem ser consumidas externamente.
- `/api`: Definições de contratos (Swagger, Proto files).
- `/web`: Ativos de interface (se houver).

---

## 🏷️ Naming Conventions (Nomenclatura)

Em Go, a clareza deve prevalecer sobre a brevidade excessiva, mas respeitando o contexto:
- **Variáveis Locais**: Devem ser curtas e significativas (ex: `idx` para índice, `err` para erro).
- **Exportação**: Somente o que começa com letra maiúscula é visível fora do pacote (`Public` vs `private`).
- **Interfaces**: Geralmente terminam em "er" (ex: `Reader`, `Writer`, `Validator`).
- **Pacotes**: Nomes curtos, em minúsculas, sem underscores ou CamelCase (ex: `http`, `json`).

---

## ✍️ Code Style & Documentation

- **Idiomaticity**: Utilize `go fmt` para formatação automática e `go lint` para análise de estilo.
- **Documentação**: Comente todas as funções exportadas começando com o nome da própria função.
  ```go
  // FetchUser retrieves a user from the database by ID.
  func FetchUser(id string) (*User, error) { ... }
  ```
- **Comments**: Comentários devem explicar o "porquê" e não o "como" (que o código já deve mostrar).

---

## 🛠️ Structs & Interfaces (Sistema de Tipos)

- **Composition over Inheritance**: Go utiliza composição e interfaces implícitas via embedding.
- **Interfaces Pequenas**: Defina interfaces focadas (ex: uma única função). Elas são mais fáceis de compor e testar.
- **Accept Interfaces, Return Structs**: Funções devem aceitar interfaces como parâmetros (máxima flexibilidade) e retornar tipos concretos/structs (facilidade de uso).
- **Consumer-Defined Interfaces**: Defina a interface no pacote que a consome. O provedor não precisa saber que a interface existe (decoupling).
- **Make the Zero Value Useful**: Desenhe tipos para que funcionem sem inicialização explícita.
  ```go
  // Bom: Zero value pronto para uso
  type Counter struct {
      mu    sync.Mutex
      count int
  }
  func (c *Counter) Inc() {
      c.mu.Lock()
      c.count++
      c.mu.Unlock()
  }
  ```
- **Functional Options Pattern**: Ideal para configurar objetos complexos com valores padrão.
  ```go
  type Server struct {
      timeout time.Duration
  }
  type Option func(*Server)
  func WithTimeout(t time.Duration) Option {
      return func(s *Server) { s.timeout = t }
  }
  func NewServer(opts ...Option) *Server {
      s := &Server{timeout: 30 * time.Second}
      for _, opt := range opts { opt(s) }
      return s
  }
  ```
- **Tags de Struct**: Utilize tags para mapeamento de dados (JSON, DB, Validation).
  ```go
  type User struct {
      ID   string `json:"id" db:"user_id"`
      Name string `json:"name" validate:"required"`
  }
  ```

---

## ⚠️ Error Handling (Tratamento de Erros)

Tratar erros em Go é controle de fluxo, não exceção:
- **Sempre Verifique**: Nunca ignore erros com `_`.
- **Wrapping**: Utilize `%w` com `fmt.Errorf` para adicionar contexto mantendo o erro original para inspeção.
- **Errors Is/As**: Use `errors.Is(err, target)` e `errors.As(err, &target)` para verificar erros específicos em Go 1.13+.
- **Sentinelas vs Tipos**: Prefira `errors.New` para erros simples e tipos customizados para erros que carregam dados extras.

---

## 📚 Bibliotecas de Apoio (Foundations)
- `samber/oops`: Para erros ricos com stack traces e metadados contextuais.
- `samber/ro`: Para garantir a imutabilidade de coleções (Read-Only).
