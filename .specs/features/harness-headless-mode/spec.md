
# Feature: Harness Headless/CI Mode

## Context
Harness atualmente só roda em modo interativo. Approval gate pede confirmação via stdin, quebrando em CI. Precisamos de flags CLI para bypass e execução direta.

## Functional Requirements
- **FR-1**: Flag `--yes` (ou `-y`): bypass automático do approval gate para comandos destrutivos.
- **FR-2**: Flag `--execute "comando"` (ou `-e`): executa uma query diretamente sem entrar em loop interativo, depois sai.
- **FR-3**: Modo headless não deve entrar em loop; deve imprimir resposta e exit(0).
- **FR-4**: Compatibilidade total com modo interativo existente (não quebrar).

## Acceptance Criteria
- **AC-1**: `harness -e "echo hello"` executa, imprime resposta do LLM, e sai (exit 0).
- **AC-2**: `harness --yes -e "rm -rf /"` não pede confirmação e executa diretamente.
- **AC-3**: `harness --yes -e "echo ok"` imprime resposta normalmente.
- **AC-4**: `harness` sem flags entra em modo interativo normal.
- **AC-5**: `go vet` e `go test` passam.
