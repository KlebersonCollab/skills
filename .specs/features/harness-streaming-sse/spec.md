
# Feature: Harness Streaming SSE (Server-Sent Events)

## Context
Harness atualmente espera resposta completa do LLM antes de mostrar qualquer saída. Para experiência de usuário mais responsiva, precisamos de streaming: tokens aparecem conforme são gerados.

## Functional Requirements
- **FR-1**: Adicionar `Streaming bool` e `StreamResponsePath string` em `ProviderConfig`.
- **FR-2**: Nova função `CallLLMStream` que envia requisição com streaming, lê SSE, extrai tokens parciais e imprime em tempo real.
- **FR-3**: Modificar `agentExecutionLoop` para usar streaming quando configurado.
- **FR-4**: Compatibilidade: providers sem streaming continuam funcionando normalmente.
- **FR-5**: Flag `--stream` / `-s` na CLI para ativar streaming (default se configurado no provider).

## Acceptance Criteria
- **AC-1**: `CallLLMStream` extrai tokens corretamente de respostas SSE do Gemini.
- **AC-2**: `CallLLMStream` extrai tokens corretamente de respostas SSE do DeepSeek.
- **AC-3**: `CallLLMStream` extrai tokens corretamente de respostas streaming do Ollama.
- **AC-4**: Agente imprime tokens conforme chegam (não bloqueia até o final).
- **AC-5**: No final do streaming, o texto completo é retornado para continuar o loop de ferramentas.
- **AC-6**: `go vet`, `go test`, `go build` passam.
