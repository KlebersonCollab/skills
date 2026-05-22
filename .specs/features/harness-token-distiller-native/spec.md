
# Feature: Token Distiller Nativo no Harness (harness-token-distiller-native)

## Descrição
Incorporar as funcionalidades de compressão de tokens da skill `token-distiller` diretamente na CLI Harness (Go), eliminando a dependência de chamadas externas para compressão de respostas e micro-compaction do histórico.

## Motivacão
- Reduzir latência (compressão nativa em Go, sem fork Python)
- Reduzir custos de tokens com compressão automática
- Dar ao usuário controle via comandos /slash (`/mode low`, `/caveman`, `/premium`)
- Micro-compaction server-side do histórico para evitar estouro de contexto

## Critérios de Aceitação (ACs)

### AC1: Módulo distiller
- [ ] Função `ToCaveman(text string, level DistillLevel) string` implementada
- [ ] `CheckAutoClarity(text string) bool` implementada (segurança e multi-step)
- [ ] Três níveis: `lite`, `full`, `ultra` (comportamento idêntico ao `distill.py`)
- [ ] Abreviações e remoção de artigos/conjunções no `full`/`ultra`
- [ ] Código entre backticks (`) NUNCA é modificado

### AC2: Integração no SessionTree
- [ ] Struct `SessionMode` adicionada em `session.go`
- [ ] Campo `Mode SessionMode` no `SessionTree`, serializado em JSON
- [ ] Valor padrão: `{Active: false, Level: "full"}`

### AC3: Comandos Slash
- [ ] `/mode low` ativa Caveman (intensity `full` por padrão)
- [ ] `/mode low lite` / `/mode low full` / `/mode low ultra`
- [ ] `/caveman` (alias para `/mode low full`)
- [ ] `/caveman lite|full|ultra`
- [ ] `/mode high` ou `/premium` desativa compressão
- [ ] Feedback visual claro no terminal ao mudar modo

### AC4: Compressão pós-LLM
- [ ] Após receber `llmResponse`, se `tree.Mode.Active == true`, aplicar `ToCaveman(response, level)`
- [ ] Respeitar `CheckAutoClarity`: se retornar true, NÃO comprimir (exibir texto original com aviso)
- [ ] A compressão ocorre antes de exibir ao usuário e antes de salvar no nó da árvore

### AC5: Micro-compaction do histórico
- [ ] Função `CompactHistory(nodes []Node) []Node` que substitui tool outputs antigos (>10 turnos e >5000 chars) por stub
- [ ] Aplicada dentro de `FormatLinearHistoryText` (ou antes de montar o prompt)
- [ ] Stub no formato: `[Old tool result cleared. Tool: system. Summary: Long tool output (%d bytes).]`

### AC6: Testes
- [ ] `distiller_test.go` com casos: lite, full, ultra, auto-clarity, backtick protection, filler removal
- [ ] `go test ./bin/harness/` passa sem erros
- [ ] `go build ./bin/harness/` compila sem erros

### AC7: Documentação
- [ ] `STATE.md` atualizado com feature `harness-token-distiller-native` como `IN_PROGRESS` e depois `COMPLETED`
- [ ] `MEMORY.md` registra a decisão arquitetural
- [ ] `LEARNINGS.md` captura padrões encontrados

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
