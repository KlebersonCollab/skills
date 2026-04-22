---
name: golang-testing-expert
version: 1.1.0
description: "Expert level Go testing patterns including TDD, table-driven tests, benchmarks, fuzzing, and HTTP testing."
category: development-testing
---

# Golang Testing Expert

> Expert level Go testing patterns including TDD, table-driven tests, benchmarks, fuzzing, and HTTP testing.

---

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar qualquer execução técnica:
1. **Context Check**: Você reidratou o contexto lendo `STATE.md`, `MEMORY.md` e `LEARNINGS.md`?
2. **Spec Check**: O arquivo `spec.md` existe com requisitos e Critérios de Aceitação (ACs) claros? (BDD mandatório para Medium+).
3. **Plan Check**: O arquivo `plan.md` define a arquitetura, schemas e inclui diagramas **Mermaid**?
4. **Contract Check**: O arquivo `contract.md` foi estabelecido com os sensores de validação?
5. **Task Check**: A lista de tarefas em `tasks.md` está detalhada e atomizada?

---

## Goal

Atuar como um especialista em QA e Testes para Go (Golang). Esta skill fornece diretrizes rigorosas para implementar o ciclo TDD, garantir alta cobertura de código, validar performance via benchmarks e aumentar a resiliência via fuzzing, seguindo os padrões mais modernos da linguagem.

---

## 🛠️ Expert Workflows

### 1. The TDD Cycle (Red-Green-Refactor)
Ao desenvolver novas funcionalidades, siga rigorosamente:
1. **RED**: Escreva um teste que falha primeiro.
2. **GREEN**: Escreva o código mínimo necessário para passar no teste.
3. **REFACTOR**: Melhore o código mantendo os testes verdes.

### 2. Table-Driven Testing
Utilize structs anônimas para definir múltiplos casos de teste em uma única função, garantindo cobertura de casos de borda e erros.

### 3. Integration Testing with Testcontainers
Para testes que dependem de infraestrutura (Banco de Dados, Redis, etc.), utilize **testcontainers-go** para subir instâncias efêmeras em Docker, garantindo isolamento total.

### 4. Verification Commands
Utilize os seguintes comandos para garantir a qualidade:

| Cenário | Comando |
|---------|---------|
| **Unitários** | `go test ./internal/... -short -count=1` |
| **Integração** | `go test ./internal/repository/... -count=1 -timeout 120s` |
| **Race Detector** | `go test ./... -race -count=1` |
| **Cobertura** | `go test ./... -coverprofile=coverage.out && go tool cover -func=coverage.out` |

---

## ⚡ Testing Checklist
- [ ] O ciclo RED-GREEN-REFACTOR foi seguido?
- [ ] Foram utilizados Table-Driven Tests para casos de sucesso e erro?
- [ ] Subtests (`t.Run`) foram utilizados para isolar os casos?
- [ ] Helpers chamam `t.Helper()` para logs de erro precisos?
- [ ] Recursos são limpos via `t.Cleanup()` ou `defer`?
- [ ] A cobertura de código (`go test -cover`) foi verificada?
- [ ] O race detector (`go test -race`) foi executado?

---

## Output Structure

A execução desta skill resulta nos seguintes artefatos e orientações:

| Artefato | Descrição |
|----------|-----------|
| **Test Suite** | Arquivos `_test.go` robustos e idiomáticos. |
| **Mocks & Stubs** | Implementações leves para isolamento de dependências. |
| **Benchmarks** | Análises de performance e alocação de memória. |
| **Fuzz Tests** | Testes de mutação para validação de inputs complexos. |
| **Coverage Report** | Relatórios detalhados de cobertura de código. |

---

## Quality Rules

- **Behavior over Implementation**: Teste o comportamento público, não os detalhes internos.
- **Fast Feedback**: Mantenha os testes unitários extremamente rápidos.
- **Meaningful Names**: Use nomes de subtests que expliquem o cenário (ex: "success_valid_user").
- **No Sleep**: NUNCA use `time.Sleep()` em testes; utilize channels ou sinais de sincronização.
- **TDD First**: Sempre escreva o teste antes do código de produção.

## Prohibited

- **NUNCA** utilizar `time.Sleep()` para sincronização em testes.
- **NUNCA** ignorar erros retornados em funções de cleanup.
- **NUNCA** deixar de usar `t.Helper()` em funções auxiliares de teste.
- **NUNCA** realizar commits de código de produção sem a suíte de testes correspondente passar.
- **NUNCA** usar frameworks de mock pesados quando uma interface simples resolve o problema.

---

## 📚 References

Esta skill está organizada em 3 pilares fundamentais:

1. [Testing Foundations](references/foundations.md) (TDD, Table-driven, Subtests)
2. [Advanced Testing](references/advanced.md) (Mocks, Helpers, Golden Files, HTTP)
3. [Performance & Fuzzing](references/performance.md) (Benchmarks, Fuzzing, Coverage)
