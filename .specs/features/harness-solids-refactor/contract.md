# Harness SOLID Refactor — Contract

## Validation Sensors

### Sensor 1: Build
```bash
cd /home/kleberson/Documentos/skills/bin/harness && go build ./...
```
Expected: Exit code 0, no compilation errors.

### Sensor 2: Tests
```bash
cd /home/kleberson/Documentos/skills/bin/harness && go test ./... -count=1 -timeout 30s
```
Expected: All tests pass.

### Sensor 3: InteractiveLoop Completo
- `main.go` deve ter um handler para prompts não-slash que chama `agentExecutionLoop`
- Verificar com: `grep -c "agentExecutionLoop" main.go` → deve ser ≥ 1

### Sensor 4: AgentLoop Completo
- `agentExecutionLoop` deve conter pelo menos uma chamada a LLM
- Verificar com: `grep -c "CallLLM\|llmProvider.Complete\|llmProvider.Stream" internal/agent/loop.go` → deve ser ≥ 1

### Sensor 5: SRP — Package Separation
```bash
ls -d internal/*/
```
Expected: `internal/agent/ internal/config/ internal/distiller/ internal/llm/ internal/mcp/ internal/session/ internal/tools/ internal/ui/`

### Sensor 6: LSP — SSETransport
```bash
grep -c "not implemented" internal/mcp/transport.go
```
Expected: 0

### Sensor 7: Swarm Closure Bug Fixed
```bash
grep -n "agentName" internal/swarm.go
```
Expected: `agentName` must be passed as parameter to closure, not captured from loop variable.

### Sensor 8: go.mod Version
```bash
grep "^go " go.mod
```
Expected: Valid Go version (1.22, 1.23, or similar real version)

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "harness-solids-refactor"
phase: "SPECIFY"
status: "IN_PROGRESS"
last_update: "2026-05-22T16:12:00Z"
```
