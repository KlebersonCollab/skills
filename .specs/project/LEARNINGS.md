# Project Learnings: Skills Hub

## Technical Patterns
- **SDD v2.3.0 Integration**: Observed that modularizing skills into independent folders with their own `SKILL.md` allows for seamless agentic consumption.
- **Hybrid Relationship Mapping**: Combining global mandates (governance) with local content analysis (mentions) creates a high-fidelity visual representation of skill interdependencies.
- **Mandate Synchronization**: Using `bin/sync-skills.sh` to propagate `GLOBAL_MANDATES.md` to root `AGENTS.md` and hidden governance directories (`.gemini`, `.claude`) ensures a single source of truth for engineering mandates.

## Bug Fixes & Gotchas
- **Vite Path Resolution**: When building dashboards inside subdirectories, ensuring correct relative paths to the root `.agents` folder is critical for data extraction scripts.

## Process Insights
- **Ripgrep Integration**: `isRipgrepAvailable()` verifica `rg --version` silenciosamente. `searchWithRipgrep` monta args dinamicamente: se há query, usa `rg --json --glob`; se só pattern, usa `rg --files --glob`. Fallback puro Go garante funcionamento sem ripgrep. Testes validam ambos os caminhos.
- **Streaming SSE**: `CallLLMStream` usa `bufio.Scanner` + `http.Flusher` para ler tokens em tempo real. Funciona com qualquer API que retorne SSE (`data: {...}`). O path de extração é configurável via `StreamResponsePath`. O teste mocka o SSE com `httptest.Server` e verifica tanto o texto completo quanto a saída parcial no stdout.
- **Sobrescrita de arquivo com write_file pode truncar se conteúdo for muito grande**: Ao recriar main.go (700+ linhas), a ferramenta `write_file` do harness pode truncar o conteúdo. Sempre verificar com `wc -l` após escrever arquivos grandes. Prefira múltiplos patches ou escrita em etapas para arquivos >500 linhas.
- **Global vars entre arquivos**: `globalHeadlessConfig` definido em `main.go` e usado em `tools.go` funciona porque ambos são `package main`. Mas se um dos arquivos for deletado, o outro quebra silenciosamente. Considerar mover para um arquivo separado `config.go` no futuro.
- **Parser Tokenizer vs Regex**: Substituir regex `(?s)` frágil por parser state-machine baseado em strings.Index resolveu parsing de conteúdo com `>` aninhado. `strings.Index` é O(n) mas suficiente para tool calls (quantidade pequena por response). Ganho principal: eliminação de falsos negativos e positivos.
- **Approval Gate Pattern**: `isDestructiveCommand()` com lista de patterns + confirmação stdin é leve e eficaz. Lista em `distiller.go` centraliza patterns reutilizáveis. Para modo headless futuro, adicionar flag `--yes` ou `--force` para bypass.
- **Go build sem regex dead code**: Ao remover regexes globais de `main.go`, o pacote `regexp` deixou de ser importado lá → `go vet` acusa unused import. Remover import resolveu. Isso reforça check de dependências mínimas.
- **Manual Bootstrapping**: Discovered that manual creation of `.specs/` structure ensures the agent fully understands the governance layer before implementation.
- **Purist Transition**: Removing CLI dependencies simplifies the environment and reduces "magic" failures, forcing the agent to rely on explicit Markdown contracts which are more resilient and traceable.
- **Governance as Code (GaC)**: Automating metadata and evidence validation via `make audit` ensures 100% compliance and prevents governance drift.
- **Centralized Knowledge Mapping**: Maintaining a root `KNOWLEDGE-MAP.mermaid` provides a consistent dependency view across all project phases.
- **Automated Link Verification**: Using simple scratch scripts to verify the presence of `SKILL.md` and directory structures across a large number of modules is highly effective for maintaining repository integrity at scale.
- **Grilling Alignment Value**: Integrating the glossary (`CONTEXT.md`) and Architectural Decision Records (ADRs) with an interactive "sabatina" (grill-with-docs) drastically reduces semantic noise and design misalignment before the specification and implementation phases begin.
- **De-ambiguation via File Consolidation**: Observed that a large number of reference files (e.g. 24 files in a skill directory) increases token cost and can lead to semantic confusion. Fusing files of the same nature (like phase and session handoffs into a single `handoff-protocol.md`, and `quick-mode`/`context-limits` into `operational-guidelines.md`) clarifies rules and maximizes cognitive efficiency for LLMs.
- **Root Mandate Synchronization**: Elevating local SDD developments (such as context budget check zones, Document Purity, and Phase 0 Grilling) to global mandates ensures that all active sub-agents (Gemini, Claude, and Core) execute with optimized focus and context-awareness.
- **Go Concurrency for CLI Animations**: Using lightweight background goroutines with `select` loops and a stop channel (`chan struct{}`) allows rendering real-time Unicode loading loaders (spinners) and timers in Go without blocking standard HTTP REST calls, ensuring smooth and modern terminal UX.
- **ANSI Card Layouts**: Wrapping text lines with vertical border symbols (`│`) and using mathematical line repeats for headers/footers (`strings.Repeat`) provides a clean card-based layout structure in command-line terminals that matches premium web-dashboard aesthetics.
- **Headless/CI Pattern**: `globalHeadlessConfig` var global + flag `--yes` permite bypass de approval gate sem modificar funções internas. Funciona como injeção de dependência simples. Para futuro: usar context ou argumento explícito em vez de global.
- **Dynamic Timeout Testing in Go**: When testing absolute and long HTTP request timeouts (such as a 2-minute limit) without slowing down the test suite or CI pipelines, exposing a package-level mutable variable (e.g. `var llmRequestTimeout = 2 * time.Minute`) that can be overridden in test files (e.g. to `15 * time.Millisecond` using a `defer` function to restore it) is a premium, clean-code pattern to validate complex connection retries instantly.

---

## Harness SOLID Refactor — Key Learnings (2026-05-22)

### SRP na Prática
- **Before**: `main.go` (396 linhas) fazia parsing de flags, wizard de init, loop interativo, árvore de sessão, cards de UI.
- **After**: `main.go` (272 linhas) faz APENAS parsing de flags + roteamento. Wizard → `cmd_init.go`, Swarm → `cmd_swarm.go`, MCP → `cmd_mcp.go`, UI → `internal/ui/`, session → `internal/session/`, LLM → `internal/llm/`.
- **Ganho**: Cada arquivo tem exatamente uma responsabilidade. Manutenção drasticamente mais fácil.

### OCP via Registry Pattern
- **Before**: `handleRequest()` no MCP usava um switch monstruoso para cada método. Adicionar método = modificar o switch.
- **After**: `mcp.Server.RegisterMethod()` permite registrar handlers. Novo método = `server.RegisterMethod("tools/execute", myHandler)`. O código do server nunca precisa ser modificado.
- **Padrão análogo**: `tools.Registry.Register()` para ferramentas, `llm.ProviderRegistry.Register()` para providers.

### LSP — SSETransport Corrigido
- **Before**: `SSETransport.Receive()` sempre retornava `"SSE transport receive not implemented"` — violação clara de LSP.
- **After**: `SSETransport` usa um canal interno (`httpCh`) que recebe dados da resposta do POST via goroutine. `Receive()` agora espera nesse canal. A interface é honesta.

### ISP — Interfaces Segregadas
- **Before**: `MCPTransport` forçava `Send`, `Receive`, `Close` em todos os transports.
- **After**: Interfaces segregadas `Sender`, `Receiver`, `Closer`. `Transport` combina as três, mas implementações parciais são possíveis.
- **LLMProvider**: Interface limpa com `Complete()` e `Stream()` — sem acoplamento a HTTP.

### DIP — Injeção de Dependência
- **Before**: `CallLLM(config, prompt, history)` — função de package que lia diretamente de `http.Client`, fazia parsing JSON, etc.
- **After**: `agent.Run(ctx, cfg, logger, provider, toolReg, tree, ...)` — recebe interfaces. O core do agente não sabe se o provider é Gemini, DeepSeek, Ollama ou mock.
- **Testabilidade**: `mockProvider` implementa `LLMProvider` em 3 linhas. Testes de swarm não precisam de HTTP real.

### Closure Bug no Swarm (Go específico)
- **Problema**: `for name := range orch.Agents { orch.BindLLM(agentName, func(...){ ... agentName ... }) }` — todas as closures capturam a MESMA variável `name`, que recebe o último valor do map.
- **Solução**: Criar cópia local `agentName := name` dentro do loop, ou iterar sobre slice em vez de map.
- **Aplicada em**: `cmd_swarm.go:runSwarm()` — `for _, agentCfg := range cfg.Agents { agentName := agentCfg.Name }`

### Interface Segregation para CompactHistory
- **Problema**: `CompactHistory` genérico com genérica (`[T any]`) era complexo e não funcionava com tipos concretos de `session.Node`.
- **Solução**: Criar `distiller.HasContent` interface com `GetContent()`, `SetContent()`, `GetRole()`. `Compact()` aceita `[]HasContent`.
- **Flexibilidade**: Qualquer struct que implemente essas 3 funções pode ser compactada.

### Migração Incremental Funciona
- Estratégia de criar `internal/` primeiro, manter `package main` intacto, depois deletar arquivos velhos UM A UM.
- A cada deleção, compilar e testar. Isso evitou refactoring catastrófico e manteve o binário funcional durante todo o processo.

### Encapsulamento de PCM em WAV para Gemini TTS (2026-05-23)
- **Problema**: O modelo `gemini-3.1-flash-tts-preview` retorna o áudio de síntese em formato **PCM bruto (24kHz, 16-bit, mono)** dentro do campo `inlineData`. Ao salvar esses bytes diretamente em arquivo com extensão `.wav` e reproduzir em mixadores modernos (como PipeWire ou PulseAudio do Zorin OS via `mpv` ou `pw-play`), o player falha com erro de formato não reconhecido ou reproduz apenas ruído branco devido à falta de metadados.
- **Solução**: Implementar uma função auxiliar pura em Go (`addWavHeader`) que monta o cabeçalho RIFF/WAVE padrão de 44 bytes em Little-Endian com as propriedades corretas da amostragem (24000 Hz, 1 channel, 16 bits por sample) e concatena os bytes do PCM bruto.
- **Resultado**: Compatibilidade nativa imediata com todos os players (`mpv`, `pw-play`, `aplay`) de forma limpa e com excelente fidelidade, eliminando ruídos.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HARNESS-VOICE-ENHANCEMENTS"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-23T00:37:00Z"
evidence_checksum: "go-build-and-test-pass"
```
