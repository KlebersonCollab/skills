
# Plan: Harness Ripgrep Integration

## Architecture

### 1. Ripgrep Detection (tools.go)
```go
func isRipgrepAvailable() bool
```
Executa `rg --version` e verifica exit code 0.

### 2. SearchFiles Refactor (tools.go)
```go
func SearchFiles(pattern string, query string) (string, error) {
    if isRipgrepAvailable() {
        return searchWithRipgrep(pattern, query)
    }
    return searchWithGoWalk(pattern, query) // existing code
}
```

### 3. searchWithRipgrep (tools.go)
- Monta comando `rg --json -g '<pattern>' '<query>'` ou similar
- Processa saída JSON linha linha
- Formata no mesmo formato do SearchFiles: `relpath:linha: preview`

### 4. Arquivos Alterados
- `bin/harness/tools.go` — refatorar SearchFiles, adicionar searchWithRipgrep, isRipgrepAvailable
- `bin/harness/harness_test.go` — adicionar teste para detecção e fallback

## Mermaid Diagram
```mermaid
flowchart TD
    [SearchFiles] --> B{Is ripgrep available?}
    B -->|Yes| C[searchWithRipgrep]
    B -->|No| D[searchWithGoWalk]
    C --> E[Execute rg with args]
    E --> F[Parse JSON output]
    F --> G[Format as relpath:lineno:preview]
    D --> H[WalkDir + read files]
    H --> G
```
