
# Tasks: Harness Streaming SSE (Real-time LLM Token Feedback)

## Status: IMPLEMENT COMPLETED ✅

- [x] TASK-1: Adicionar campo `StreamResponsePath string` e `Streaming bool` em `ProviderConfig` (client.go)
- [x] TASK-2: Implementar `CallLLMStream` — leitura chunked SSE, extração parcial, callback onChunk (client.go)
- [x] TASK-3: Implementar `extractStreamToken` — parser genérico para SSE (Gemini, DeepSeek, Ollama) com fallback automático (client.go)
- [x] TASK-4: Adicionar flag `--stream` / `-s` e modificar `agentExecutionLoop` para usar streaming (main.go)
- [x] TASK-5: Atualizar wizard `init` para salvar `Streaming` e `StreamResponsePath` por provider (main.go)
- [x] TASK-6: Testes unitários — 11 testes: extração de tokens, integração mock SSE, DeepSeek delta, Ollama NDJSON, fallback, erro HTTP, stream vazio (harness_test.go)
- [x] TASK-7: `go vet` ✅, `go test` ✅ (100% passing), `go build` ✅

## Evidências

| Tarefa | Evidência |
|--------|-----------|
| TASK-1 | `client.go:12-17` — `Streaming bool`, `StreamResponsePath string` adicionados |
| TASK-2 | `client.go:80-215` — `CallLLMStream` implementada |
| TASK-3 | `client.go:217-243` — `extractStreamToken` com 3 formatos (Gemini, DeepSeek, Ollama) |
| TASK-4 | `main.go:22-42` — Flag parsing; `main.go:275-320` — streaming loop |
| TASK-5 | `main.go:154-178` — Gemini/Ollama/DeepSeek com streaming habilitado |
| TASK-6 | `harness_test.go` — 11 testes passando |
| TASK-7 | `go vet` passou, `go build` passou, `go test` passou (0.557s) |

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "harness-streaming-sse"
phase: "IMPLEMENT"
status: "COMPLETED"
last_update: "2026-05-23T15:15:00Z"
evidence_checksum: "go-test-pass-0.557s"
```
