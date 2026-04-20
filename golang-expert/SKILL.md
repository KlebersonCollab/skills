---
name: golang-expert
version: 1.1.0
description: "Expert level Go development skill focused on performance, idiomatic concurrency, and clean architecture."
category: development
---

# Golang Expert

> Expert level Go development skill focused on performance, idiomatic concurrency, and clean architecture.

---

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar a implementação:
1. **Spec Check**: O arquivo `spec.md` existe e contém os critérios de aceitação (ACs)?
2. **Plan Check**: O arquivo `plan.md` define a arquitetura e schemas?
3. **Task Check**: A lista de tarefas em `tasks.md` está detalhada?

---

## Goal

Atuar como um arquiteto e engenheiro de software sênior especializado em Go (Golang). Esta skill fornece diretrizes, padrões e melhores práticas para construir sistemas escaláveis, resilientes e de alta performance, utilizando o máximo do potencial idiomático da linguagem e seu ecossistema moderno.

---

## 🛠️ Expert Workflows

### 1. New Project Scaffolding
Ao iniciar um novo projeto, siga a ordem:
1. `go mod init <module_path>`
2. Configure o layout de pastas (`/cmd`, `/internal`, `/pkg`).
3. Instale e configure o `golangci-lint`.
4. Defina o container de DI (ex: `samber/do`).
5. Configure o logging estruturado nativo (`slog`).

### 2. High-Concurrency Design
Para fluxos concorrentes:
1. Desenhe o fluxo de dados (Data Flow) primeiro.
2. Identifique pontos de contenção.
3. Utilize `context.Context` para propagação de cancelamento.
4. Implemente Graceful Shutdown para todas as goroutines.
5. Valide com `go test -race`.

---

## ⚡ Performance Checklist
- [ ] O código passou por `go test -bench`?
- [ ] A análise de escape (`-gcflags="-m"`) foi verificada para evitar alocações desnecessárias no Heap?
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
- **Safety Always**: Evitar race conditions, leaks de memória e uso indevido de `unsafe`.
- **Modern Go**: Utilizar recursos de versões recentes (Generics, `slog`, `errors.Join`, etc.).
- **Zero-Boilerplate**: Utilizar bibliotecas de utilitários como `samber/lo` e `samber/mo` para reduzir ruído de código repetitivo.

## Prohibited

- **NUNCA** ignorar erros (`_ = err`); sempre tratar ou propagar com contexto.
- **NUNCA** utilizar `panic` para controle de fluxo normal; reserve para erros catastróficos.
- **NUNCA** iniciar goroutines sem um mecanismo claro de cancelamento ou finalização (leaks).
- **NUNCA** misturar lógica de negócio com detalhes de infraestrutura (Clean Architecture).

---

## 📚 References

Esta skill está organizada em 5 pilares fundamentais:

1. [Foundations & Style](references/foundations.md)
2. [Concurrency & Advanced](references/concurrency.md)
3. [Development & Tooling](references/development.md)
4. [Systems & Architecture](references/architecture.md)
5. [Ecosystem & Samber](references/ecosystem.md)
