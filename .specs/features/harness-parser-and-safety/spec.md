
# Feature: Harness XML Parser & Safety Gate

## Context
Harness CLI usa regex frágil para parsear tool XML. Caso conteúdo contenha `>` aninhado, regex quebra. Além disso, `ExecuteCommand` não tem proteção contra comandos destrutivos (rm -rf, DROP, etc.).

## Functional Requirements
- **FR-1**: Substituir regex `(?s)` por parser tokenizer simples baseado em stack de tags XML.
- **FR-2**: Validar malformed XML (tags não fechadas, aninhamento incorreto) com erro claro.
- **FR-3**: Implementar approval gate opcional p/ comandos destrutivos listados em `distiller.go`.
- **FR-4**: Preservar compatibilidade total com formato atual de tool XML.
- **FR-5**: Limpar código morto (`compactHistory` não usada em `main.go`).

## Acceptance Criteria
- **AC-1**: `write_file` com conteúdo contendo `>` e `<` dentro de CDATA ou texto é parseado corretamente.
- **AC-2**: `patch_file` com target contendo `>` é parseado corretamente.
- **AC-3**: `execute_command` com conteúdo contendo `>` é parseado corretamente.
- **AC-4**: Se XML estiver malformed (ex: tag <tool:write_file sem fechamento), retorna erro descritivo.
- **AC-5**: Se `execute_command` contiver pattern destrutivo (ex: `rm -rf`), Harness pergunta confirmação via stdin antes de executar.
- **AC-6**: Código morto removido sem quebrar build/go vet.
- **AC-7**: Todos os testes existentes passam.

## Non-Requirements
- Não implementar parser XML completo (ns, namespaces, atributos com aspas duplas aninhadas). Apenas o necessário p/ tags tool.
- Approval gate não deve ser implementado com confirmação automática em modo headless (future).
