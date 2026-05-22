# Plano de Implementação - Upgrade Visual e UX do Harness CLI (harness-ui-upgrade)

Aprimoramento estético e de usabilidade do terminal interativo do Harness em Go (`bin/harness`). O terminal atual, embora funcional, é puramente textual e carece de feedback visual de tempo de execução, formatação de blocos e indicadores de loading dinâmicos. Inspirando-se no `pi.dev` e em ferramentas CLI modernas, traremos uma experiência visual premium baseada em caracteres Unicode, cores ANSI TrueColor e animações assíncronas em Go puro, sem qualquer dependência externa de terceiros.

---

## Proposed Changes

### 1. Novo Design do Prompt de Comando (Starship-like)
A linha de prompt interativa do Harness será estilizada com cores ANSI vibrantes e separadores de estilo Starship, tornando a indicação do nó ativo e o nome do agente extremamente elegantes.
*   **Formato:** `harness ❯ [node-ID] ❯ `
*   **Cores:** `harness` em roxo/azul brilhante (`\033[1;38;5;99m`), setas `❯` em magenta brilhante (`\033[38;5;198m`), e o ID do nó em ciano/azul claro (`\033[36m`).

### 2. Spinner Animado Assíncrono com Contador de Tempo
Durante chamadas à API da LLM ou execução de ferramentas, uma animação Unicode (`⠋`, `⠙`, `⠹`...) com contagem em tempo real de segundos (ex: `2.4s`) será exibida em background, limpando-se do console assim que a execução concluir.
*   **Módulo:** [main.go](file:///home/kleberson/Documentos/skills/bin/harness/main.go)
*   **Função Nova:** `startSpinner(suffix string) chan struct{}`

### 3. Cards de Mensagens do Assistente (Balão de Diálogo Elegante)
Todas as respostas textuais geradas pela LLM serão envelopadas em uma moldura de console com bordas Unicode e linhas laterais, apresentando informações de metadados como ID do nó, horário e tokens aproximados.
*   **Caracteres de Borda:** `┌──`, `│`, `└──`

### 4. Cards de Visualização de Ferramentas (Tool Cycle Boxes)
O uso de ferramentas como `read_file`, `write_file`, `patch_file` e `execute_command` será envelopado em blocos visuais distintos, com sub-painéis para exibição de erros, saída do console bash e relatórios de sucesso.

---

## proposed-changes-per-file

### [Component: Harness UI]

#### [MODIFY] [main.go](file:///home/kleberson/Documentos/skills/bin/harness/main.go)
*   Adicionar função utilitária `startSpinner(message string) chan struct{}` para inicializar o loader concorrente.
*   Refatorar a função `agentExecutionLoop` para:
    *   Envelopar chamadas de `CallLLM` com o spinner dinâmico em background.
    *   Refatorar as impressões de ferramentas (`read_file`, `write_file`, `patch_file`, `execute_command`) em caixas estilizadas de console com linhas laterais Unicode.
    *   Formatatar a saída final do Assistente com um card estruturado com borda verde/cyan, exibindo cabeçalho e rodapé decorados.
*   Refatorar a linha de prompt interativa de `runInteractiveLoop` para usar o padrão colorido Starship.

#### [MODIFY] [session.go](file:///home/kleberson/Documentos/skills/bin/harness/session.go)
*   Ajustar a visualização da árvore de histórico no comando `/history` para usar paleta de cores harmonizada com o novo tema visual.

---

## Verification Plan

### Automated Tests
*   **Go Test Suite:**
    *   Executar `go test -v ./bin/harness/...` para garantir que as alterações não introduziram quebras de compilação ou de lógica.
*   **Compliance Audit:**
    *   Executar `make audit` para garantir 100% de conformidade com as diretrizes do SDD v2.3.0 do Skills Hub.

### Manual Verification
*   Compilar o harness atualizado: `go build -o bin/harness/harness bin/harness/*.go` (ou usar o target correspondente do Makefile).
*   Iniciar o harness interativo: `./bin/harness/harness` e executar comandos para verificar o alinhamento visual, o spinner assíncrono durante chamadas LLM e a elegância dos cards impressos.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HARNESS-UI-UPGRADE"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-22T15:07:00Z"
evidence_checksum: "go-test-and-make-audit-pass-002"
```
