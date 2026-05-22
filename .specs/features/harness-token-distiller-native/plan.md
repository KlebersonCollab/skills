
# Plano: Token Distiller Nativo no Harness

## Divisão em Tarefas

### TASK-1: Módulo distiller.go
**Arquivo**: `bin/harness/distiller.go`
**Descrição**: Implementar as funções puras de compressão baseadas no `distill.py`.
**ACs**: AC1
**Estimativa**: Pequena (~80 linhas)

### TASK-2: Testes distiller_test.go
**Arquivo**: `bin/harness/distiller_test.go`
**Descrição**: Testes unitários para todas as funções do distiller.
**ACs**: AC6
**Estimativa**: Pequena (~100 linhas)

### TASK-3: Modificar session.go
**Arquivo**: `bin/harness/session.go`
**Descrição**: Adicionar struct `SessionMode` e campo `Mode SessionMode` no `SessionTree`, com valor padrão inativo.
**ACs**: AC2
**Estimativa**: Minúscula (~10 linhas)

### TASK-4: Modificar main.go (comandos slash)
**Arquivo**: `bin/harness/main.go`
**Descrição**: Adicionar handlers para `/mode`, `/caveman`, `/premium` no loop interativo e feedback visual.
**ACs**: AC3
**Estimativa**: Média (~80 linhas)

### TASK-5: Integrar compressão e micro-compaction no agentExecutionLoop
**Arquivo**: `bin/harness/main.go`
**Descrição**: Aplicar `ToCaveman` pós-LLM e `CompactHistory` antes de montar o prompt.
**ACs**: AC4, AC5
**Estimativa**: Média (~50 linhas)

### TASK-6: Compilar e testar
**Comando**: `go build ./bin/harness/` e `go test ./bin/harness/`
**Descrição**: Garantir que tudo compila e os testes passam.
**ACs**: AC6
**Estimativa**: Minúscula

### TASK-7: Atualizar documentação do projeto
**Arquivos**: STATE.md, MEMORY.md, LEARNINGS.md
**Descrição**: Registrar feature concluída, decisões arquiteturais e aprendizados.
**ACs**: AC7
**Estimativa**: Minúscula

## Ordem de Execução
1. TASK-1 (distiller.go)
2. TASK-2 (distiller_test.go) — pode ser feito em paralelo ou depois
3. TASK-3 (session.go)
4. TASK-4 (main.go comandos)
5. TASK-5 (main.go integração)
6. TASK-6 (build + test)
7. TASK-7 (docs)

## Dependências
- Nenhuma externa; apenas pacotes padrão Go (`regexp`, `strings`, `fmt`).

## Riscos
- RegEx mal escritas podem falhar em edge cases (backticks aninhados, unicode). Testes mitigam.
- Adicionar campo `Mode` ao `SessionTree` pode quebrar sessões JSON existentes. Usar `omitempty` para compatibilidade.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HARNESS-TOKEN-DISTILLER-NATIVE"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-23T11:00:00Z"
evidence_checksum: "go-test-pass-distiller-001"
```
