# Harness SOLID Refactor — Specification

## Goal
Refatorar o `bin/harness/` Go Agent para seguir SOLID, completar implementações incompletas, e se aproximar arquiteturalmente do `pi.dev` (PI Coding Agent).

## Acceptance Criteria

### AC-01: Agent Loop Completo
- `executor.go::agentExecutionLoop()` deve: montar system prompt → chamar `CallLLM`/`CallLLMStream` → parsear XML tools → executar tools → armazenar resultado no SessionTree → iterar até maxIterations.

### AC-02: Interactive Loop Completo
- `main.go::runInteractiveLoop()` deve tratar prompts do usuário (não-slash) como entrada para o `agentExecutionLoop()`.

### AC-03: SRP — Separação de Responsabilidades
- Criar pacotes: `llm/`, `session/`, `tools/`, `mcp/`, `ui/`, `distiller/`.
- `main.go` deve conter **apenas** inicialização e roteamento de comandos.

### AC-04: OCP — Registry Pattern para Providers e Comandos
- Providers LLM registráveis via interface `LLMProvider`, não por string replace de template.
- Comandos CLI registráveis via registry, não por switch.

### AC-05: LSP — MCPTransport Consistente
- `SSETransport` deve implementar `MCPTransport` sem métodos que sempre retornam erro.

### AC-06: ISP — Interfaces Segregadas
- `LLMProvider` interface com `Complete(ctx, prompt) → (string, error)` e `Stream(ctx, prompt) → (<-chan string, error)`.
- `MCPTransport` interface segregada: `Sender`, `Receiver`, `Closer`.

### AC-07: DIP — Injeção de Dependência
- `agentExecutionLoop` deve receber interfaces, não `*AppConfig` concreto.
- `SwarmOrchestrator` deve receber `LLMProvider`, não closure.

### AC-08: Swarm Bug Fix
- Closure bug do `agentName` em `runSwarm()` deve ser corrigido.

### AC-09: Distiller Integrado
- `ToCaveman()`, `CompactHistory()`, `isDestructiveCommand()` devem ser chamados no loop do agente.

### AC-10: Duplicação Eliminada
- `CallLLM` e `CallLLMStream` devem compartilhar helper de template, headers e retry.

### AC-11: go.mod Corrigido
- `go 1.26.1` → versão real do Go.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "harness-solids-refactor"
phase: "DISCOVERY"
status: "IN_PROGRESS"
last_update: "2026-05-22T16:00:00Z"
```
