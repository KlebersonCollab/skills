# Pilar 2: Concurrency & Advanced

> Concorrência em Go não é sobre paralelismo, mas sobre a composição de tarefas independentes que se comunicam via channels e sinais de cancelamento.

---

## 🏎️ Concurrency Fundamentals (Goroutines & Channels)

- **Share Memory by Communicating**: Siga o mantra do Go — não comunique compartilhando memória, compartilhe memória comunicando.
- **Channels**: Utilize channels para sincronização e transferência de dados entre goroutines.
  - Channels com buffer para desacoplamento temporal controlado.
  - Channels sem buffer para sincronização rígida.
- **Select Statement**: Utilize o `select` para orquestrar múltiplos channels e timeouts.

---

## 🚦 Context Management (Pacote `context`)

O pacote `context` é a espinha dorsal de qualquer sistema resiliente em Go:
- **Propagation**: Sempre passe o `context.Context` como o primeiro argumento de funções de longa duração (IO, DB, HTTP).
- **Cancellation**: Utilize `WithCancel`, `WithTimeout` ou `WithDeadline` para evitar goroutines "zumbis" que rodam indefinidamente.
- **Values**: Utilize `WithValue` apenas para dados transientes e globais (ex: Request ID, Logs), nunca para lógica de negócio.

---

## 🛡️ Concurrency Safety (Segurança)

- **Race Conditions**: Utilize `go test -race` religiosamente para detectar condições de corrida.
- **WaitGroups**: Utilize `sync.WaitGroup` para esperar a finalização de um conjunto de goroutines.
- **errgroup**: Utilize `golang.org/x/sync/errgroup` para coordenar grupos de goroutines que podem retornar erros, permitindo o cancelamento do grupo inteiro se uma falhar.
  ```go
  g, ctx := errgroup.WithContext(ctx)
  for _, url := range urls {
      url := url
      g.Go(func() error {
          return fetch(ctx, url)
      })
  }
  if err := g.Wait(); err != nil { return err }
  ```
- **Mutexes**: Utilize `sync.Mutex` ou `sync.RWMutex` para proteger estados compartilhados quando channels não forem a melhor opção.
- **Atomic Operations**: Para contadores simples ou flags de estado, utilize `sync/atomic` para máxima performance sem travas.

---

## 🧩 Advanced Design Patterns

- **Fan-out/Fan-in**: Distribua trabalho pesado em múltiplas goroutines e consolide os resultados em um único channel.
- **Worker Pools**: Limite o número de goroutines simultâneas para evitar sobrecarga de recursos.
- **Graceful Shutdown**: Capture sinais do SO (SIGINT, SIGTERM) para fechar recursos e finalizar goroutines antes de encerrar o processo.
- **Circuit Breaker**: Implemente padrões de resiliência para falhas em cascata em serviços externos.
- **Singleton**: Utilize `sync.Once` para inicializações únicas seguras para múltiplas goroutines.

---

## 💉 Dependency Injection (DI)

- **Interface-based**: Injete dependências como interfaces, não como tipos concretos, para facilitar mocks em testes.
- **Constructor Pattern**: Utilize funções `New...` para inicializar structs com suas dependências.
- **DI Containers**: Para projetos de larga escala, utilize containers de injeção de dependência para evitar o "Constructor Hell".

---

## 📚 Bibliotecas de Apoio (Concurrency & Advanced)
- `samber/do`: Uma biblioteca de injeção de dependência baseada em tipos, simples e poderosa.
- `samber/hot`: Para recarregamento dinâmico de configurações ou código sem interromper o serviço.
