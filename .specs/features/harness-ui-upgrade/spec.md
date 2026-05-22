# Especificação da Feature - Harness Terminal UI Upgrade

Evolução do terminal interativo do Harness em Go para uma interface estritamente premium, funcional e elegante. 

---

## Objetivos e Requisitos

1.  **Feedback Visual de Tempo Real (Dynamic Loader)**
    *   Exibir uma animação unicode (spinner) enquanto requisições de API à LLM e chamadas à ferramentas estão em andamento.
    *   Exibir o tempo de execução decorrido em segundos junto ao spinner.
    *   Ao término da operação, o spinner deve ser apagado de forma limpa do console.

2.  **Identidade Visual Coesa (Colors & Aesthetics)**
    *   Utilizar cores ANSI sofisticadas no prompt, cabeçalhos, metadados e blocos.
    *   Harmonizar o tema visual com tons de roxo, magenta, ciano e cinza escuro.
    *   Evitar cores puras e primárias sem estilização.

3.  **Delimitação de Mensagens (Card-based Layout)**
    *   As respostas da LLM devem ser desenhadas dentro de caixas ("Cards") delimitadas por caracteres de linha Unicode (`┌`, `│`, `└`).
    *   Toda mensagem de assistente deve exibir claramente o metadados do nó ativo, timestamp e tokens aproximados.

4.  **Caixas de Status de Ferramentas (Tool Cycles Visuals)**
    *   Substituir logs brutos de execução de ferramentas por blocos Unicode coloridos correspondentes.
    *   Indentação estruturada para saídas de comando Bash (`execute_command`) para que fiquem alinhadas dentro da borda visual do painel.

---

## Critérios de Aceitação (Acceptance Criteria)

*   [ ] **AC-1:** O prompt interativo deve exibir `harness ❯ [node-ID] ❯ ` com cores elegantes baseadas em ANSI.
*   [ ] **AC-2:** A chamada para `CallLLM` deve apresentar um spinner assíncrono funcional com contador de tempo ativo.
*   [ ] **AC-3:** O spinner deve sumir e limpar a linha do console após a resposta ser recebida.
*   [ ] **AC-4:** As respostas do assistente devem vir delimitadas por caixas elegantes com a borda lateral esquerda `│`.
*   [ ] **AC-5:** Ferramentas devem imprimir cards compactos, e comandos do shell executados devem ter suas saídas capturadas e formatadas elegantemente dentro das bordas.
*   [ ] **AC-6:** O código Go deve compilar sem dependências externas (`go build` de sucesso).
*   [ ] **AC-7:** O `make audit` deve rodar com 100% de conformidade.

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
