# Contract: CI Pipeline (Skill Hub Guard)

## Infrastructure Interfaces
| Interface | Command | Description |
| :--- | :--- | :--- |
| `Validator` | `make validate` | Executa a validação de estrutura de todas as skills. |
| `Version Check` | `make verify-vers` | Garante consistência de versões no ecossistema. |

## Mandatory Artifacts
- **File**: `.github/workflows/ci.yml`
- **File**: `Makefile` (com os targets validate e verify-vers)

## Pipeline Requirements
- Deve rodar em `push` e `pull_request`.
- Deve bloquear merges se qualquer interface retornar código != 0.
