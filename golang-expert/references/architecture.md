# Pilar 4: Systems & Architecture

> Construir sistemas resilientes e escaláveis em Go requer uma compreensão profunda da infraestrutura e das ferramentas que sustentam aplicações de missão crítica.

---

## 🔌 Communication & gRPC (Microserviços)

Go é a escolha dominante para microserviços de alta performance:
- **gRPC & Protobuf**: Utilize para comunicação inter-serviços de baixa latência e contratos tipados e versionáveis.
- **RESTful APIs**: Utilize o pacote `net/http` ou frameworks leves como `chi` ou `echo`.
- **API Versioning**: Planeje o versionamento desde o início, seja via cabeçalhos, URL ou pacotes Protobuf.

---

## 🗄️ Database Interoperability (Bancos de Dados)

O ecossistema Go oferece flexibilidade total para acesso a dados:
- **database/sql**: O pacote nativo para interação com SQL. Utilize-o diretamente para controle total.
- **ORMs & Query Builders**: Considere bibliotecas como `GORM` para produtividade ou `sqlc` para gerar código Go idiomático a partir de SQL puro.
- **NoSQL**: Utilize drivers oficiais para MongoDB, Redis, Cassandra, etc., sempre injetando dependências via interfaces.
- **Transaction Management**: Garanta a atomicidade utilizando padrões de gerenciamento de transações no repositório.

---

## 📊 Observability (Logs, Metrics, Tracing)

Um sistema Expert deve ser visível internamente:
- **Structured Logging**: Utilize o pacote nativo `slog` para logs em JSON estruturado, facilitando a ingestão por sistemas como ELK ou Datadog.
- **Metrics**: Utilize bibliotecas como `prometheus/client_golang` para expor métricas de performance e negócio.
- **Distributed Tracing**: Implemente `OpenTelemetry` para rastrear requisições através de múltiplos microserviços.

---

## ⚡ Performance & Benchmarking (Otimização)

Go foi projetada para ser rápida. Prove que seu código é rápido:
- **Micro-benchmarks**: Utilize `testing.B` para medir funções específicas.
  ```go
  func BenchmarkSum(b *testing.B) {
      for i := 0; i < b.N; i++ { Sum(1, 2) }
  }
  ```
- **Heap vs Stack**: Entenda a análise de escape (`go build -gcflags="-m"`) para minimizar alocações no heap e reduzir a carga do Garbage Collector.
- **Optimization Strategy**: Meça primeiro (`pprof`), identifique o gargalo e otimize apenas o necessário (YAGNI).

---

## 🛡️ Security & CI/CD (Governança)

- **Input Validation**: Nunca confie no input do usuário. Valide todos os dados na entrada da API.
- **Secrets Management**: Nunca comite chaves de API ou segredos. Utilize variáveis de ambiente ou gerenciadores de segredos (HashiCorp Vault, AWS Secrets Manager).
- **CI Pipelines**: Configure pipelines automatizados para rodar testes, linters e scans de segurança em cada commit.
- **Dockerization**: Crie imagens Docker otimizadas utilizando `distroless` ou imagens multi-stage para reduzir a superfície de ataque.

---

## 📚 Bibliotecas de Apoio (Architecture)
- `samber/slog`: Handlers e utilitários para enriquecer o log estruturado nativo `slog`.
- `samber/do`: Para orquestrar a injeção de dependências em microserviços complexos.
