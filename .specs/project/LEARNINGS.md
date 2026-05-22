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

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HARNESS-TIMEOUT-RETRY"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-22T15:42:00Z"
evidence_checksum: "go-test-pass-retry-002"
```
