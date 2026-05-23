# Harness Web Tools — Specification

## Goal
Implementar as 3 ferramentas web que faltam para o harness se equiparar ao pi.dev:
- **web_search**: Pesquisa na web com múltiplas queries, filtros e sources
- **code_search**: Pesquisa em repositórios de código e documentação técnica
- **fetch_content**: Extrai conteúdo de URLs, YouTube, GitHub, e vídeos locais

## Acceptance Criteria

### AC-01: Web Search Tool
- `tools.Registry` deve aceitar `web_search` como tool registrável
- Suporte a queries únicas e múltiplas (até 4 queries em paralelo)
- Suporte a filtros: `numResults`, `recencyFilter` (day/week/month/year), `domainFilter`
- Retorna resposta sintetizada com citações (sources)
- Fallback entre providers: Perplexity → Exa → Gemini (via env vars `PERPLEXITY_API_KEY`, `EXA_API_KEY`, `GEMINI_API_KEY`)
- Provider seletável via parâmetro `provider` (auto/perplexity/exa/gemini)

### AC-02: Code Search Tool
- `tools.Registry` deve aceitar `code_search` como tool registrável
- Pesquisa em GitHub (via GitHub API + `GITHUB_TOKEN` opcional)
- Pesquisa em Stack Overflow (via Stack Exchange API)
- Retorna snippets de código com links e citações
- Suporte a limite de tokens (maxTokens)

### AC-03: Fetch Content Tool
- `tools.Registry` deve aceitar `fetch_content` como tool registrável
- Fetch de URL(s) única ou múltipla com extração de conteúdo como markdown
- YouTube: transcrição + thumbnails via `yt-dlp` + `ffmpeg` (fallback: invocação de comando externo)
- GitHub: clonagem leve para leitura de repositórios
- Vídeo local: extração de frames via `ffmpeg`
- Parâmetros: `prompt` (para análise de vídeo), `timestamp`, `frames`

### AC-04: Provider Configuration
- API keys lidas de environment variables
- Wizard `harness init` atualizado para configurar web search providers
- Config persistida em `.harness/config.json`

### AC-05: Testes
- Testes unitários para parsing de parâmetros
- Testes de fallback entre providers
- Testes de integração com mocks HTTP

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "harness-web-tools"
phase: "SPECIFY"
status: "IN_PROGRESS"
last_update: "2026-05-22T16:15:00Z"
```
