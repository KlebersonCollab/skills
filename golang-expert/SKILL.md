---
name: golang-expert
version: 1.2.0
description: "Expert level Go development skill focused on performance, idiomatic concurrency, and clean architecture."
category: development
---

# Golang Expert

> Expert level Go development skill focused on performance, idiomatic concurrency, and clean architecture.

---

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar qualquer execução técnica:
1. **Context Check**: Você reidratou o contexto lendo `STATE.md`, `MEMORY.md` e `LEARNINGS.md`?
2. **Spec Check**: O arquivo `spec.md` existe com requisitos e Critérios de Aceitação (ACs) claros? (BDD mandatório para Medium+).
3. **Plan Check**: O arquivo `plan.md` define a arquitetura, schemas e inclui diagramas **Mermaid**?
4. **Contract Check**: O arquivo `contract.md` foi estabelecido com os sensores de validação?
5. **Task Check**: A lista de tarefas em `tasks.md` está detalhada e atomizada?

---

## ⚡ Slash Commands (Operational)
Utilize estes comandos para guiar seu comportamento durante a sessão:

| Comando | Descrição | Ação Esperada |
|---------|-----------|---------------|
| `/plan` | Planejamento de Feature | Gera/atualiza `spec.md` e `plan.md` com ADRs. |
| `/go-build` | Correção de Compilação | Executa `go build ./...` e corrige erros reportados. |
| `/go-test` | Ciclo TDD | Inicia workflow TDD (Red-Green-Refactor) com `golang-testing-expert`. |
| `/go-review` | Revisão de Código | Revisa idiomas Go, tratamento de erros e concorrência. |
| `/security-scan` | Auditoria de Segurança | Busca por segredos expostos e vulnerabilidades conhecidas. |

---

## Goal

Atuar como um arquiteto e engenheiro de software sênior especializado em Go (Golang). Esta skill fornece diretrizes, padrões e melhores práticas para construir sistemas escaláveis, resilientes e de alta performance, utilizando o máximo do potencial idiomático da linguagem e seu ecossistema moderno.

---

## 🛠️ Expert Workflows

### 1. New Project Scaffolding & Architecture
Ao iniciar um novo projeto, siga a estrutura de **Clean Architecture**:
1. `go mod init <module_path>`
2. Layout de pastas:
   - `cmd/<app>/main.go`: Ponto de entrada (DI, Graceful Shutdown).
   - `internal/domain/`: Entidades, interfaces de repositório e erros sentinela.
   - `internal/service/`: Lógica de negócio pura.
   - `internal/repository/`: Acesso a dados (ex: `postgres/` usando `sqlc`).
   - `internal/handler/`: Transporte (subpastas `grpc/`, `rest/`).
   - `internal/config/`: Carregamento via env/yaml.
3. Ferramentas Mandatórias:
   - **sqlc**: Para geração de código SQL type-safe.
   - **Wire**: Para injeção de dependência explícita.
   - **golangci-lint**: Para análise estática rigorosa.

### 2. High-Concurrency Design
Para fluxos concorrentes:
1. Desenhe o fluxo de dados (Data Flow) primeiro.
2. Identifique pontos de contenção.
3. Utilize `context.Context` para propagação de cancelamento.
4. Utilize `golang.org/x/sync/errgroup` para coordenação de goroutines e propagação de erros.
5. Implemente Graceful Shutdown para todas as goroutines.
6. Valide com `go test -race`.

---

## ⚡ Performance Checklist
- [ ] O código passou por `go test -bench`?
- [ ] A análise de escape (`-gcflags="-m"`) foi verificada para evitar alocações desnecessárias no Heap?
- [ ] Slices e Maps foram pré-alocados quando o tamanho é conhecido (`make([]T, 0, cap)`)?
- [ ] Foi utilizado `strings.Builder` para concatenação de strings em loops?
- [ ] Foram utilizados pools de objetos (`sync.Pool`) para estruturas frequentemente alocadas?
- [ ] O tratamento de erros está utilizando `wrapping` (%w) de forma eficiente?
- [ ] O pacote `slog` está configurado para o nível de log correto em produção?

---

## Output Structure

A execução desta skill resulta nos seguintes artefatos e orientações:

| Artefato | Descrição |
|----------|-----------|
| **Architecture ADR** | Decisões técnicas fundamentadas para projetos Go. |
| **Idiomatic Code** | Implementações que seguem rigorosamente o `go-code-style` e `naming`. |
| **Concurrency Flow** | Desenhos de fluxos de concorrência seguros (goroutines, channels, context). |
| **Optimization Report** | Análises de performance, benchmarks e sugestões de profiling. |
| **Validation Suite** | Testes unitários e de integração robustos (com `testify`). |

---

## Quality Rules

- **Idiomatic First**: Sempre priorizar a simplicidade ("Clear is better than clever") e as convenções oficiais (`go fmt`, `go lint`).
- **Interfaces e Structs**: "Accept interfaces, return structs" (Aceite interfaces, retorne structs concretas).
- **Zero Value Useful**: Desenhar tipos de forma que seu valor zero seja utilizável sem inicialização explícita (ex: `sync.Mutex`, `bytes.Buffer`).
- **Functional Options**: Utilizar o padrão Functional Options para APIs com configurações complexas ou opcionais.
- **Consumer-Defined Interfaces**: Definir interfaces no pacote que as consome, não no que as provê, para evitar acoplamento desnecessário.
- **Safety Always**: Evitar race conditions, leaks de memória e uso indevido de `unsafe`.
- **Modern Go**: Utilizar recursos de versões recentes (Generics, `slog`, `errors.Join`, etc.).
- **Zero-Boilerplate**: Utilizar bibliotecas de utilitários como `samber/lo` e `samber/mo` para reduzir ruído de código repetitivo.
- **Documentation**: Tipos e funções exportados **devem** ter comentários de documentação.

## Prohibited

- **NUNCA** ignorar erros (`_ = err`); sempre tratar ou propagar com contexto.
- **NUNCA** utilizar `panic` para controle de fluxo normal; reserve para erros catastróficos.
- **NUNCA** iniciar goroutines sem um mecanismo claro de cancelamento ou finalização (leaks).
- **NUNCA** colocar `context.Context` dentro de structs; ele deve ser sempre o primeiro parâmetro de funções.
- **NUNCA** misturar lógica de negócio com detalhes de infraestrutura (Clean Architecture).
- **NUNCA** utilizar a função `init()` para inicialização; prefira injeção explícita no `main()`.
- **NUNCA** utilizar estado mutável global.

---

## 📚 References

Esta skill está organizada em 5 pilares fundamentais:

1. [Foundations & Style](references/foundations.md)
2. [Concurrency & Advanced](references/concurrency.md)
3. [Development & Tooling](references/development.md)
4. [Systems & Architecture](references/architecture.md)
5. [Ecosystem & Samber](references/ecosystem.md)
6. [Testing Expert (Specialized Skill)](../golang-testing-expert/SKILL.md)
