# Implementation Plan: Harness Contract Guardian

## Architecture

### 1. CLI Layer
- Subcommand `contract [feature]` in `hb/cmd/harness.go`.

### 2. Domain Layer (`hb/internal/harness/contract.go`)
- `ExecuteContractCheck(root, feature string) error`
- `Parser`: Logic to extract tables from Markdown.
- `Validator`: Logic to verify existence of symbols and files.

### 3. Implementation Details
- **Markdown Parsing**: Use regex to find markdown tables `| Interface | ... |`.
- **Symbol Search**: 
    - Use `grep` (via `exec.Command`) to look for strings like `var <name> = &cobra.Command` or `func <name>`.
- **Feature Detection**: Auto-detect feature by looking into `.specs/features/`.

## Mermaid Diagram
```mermaid
graph TD
    CLI[hb harness contract] --> Parser[Contract Parser]
    Parser --> ContractMD[.specs/features/contract.md]
    Parser --> List[Expected Interfaces/Files]
    
    List --> Validator[Contract Validator]
    Validator --> Filesystem[File Check]
    Validator --> Codebase[Grep/AST Symbol Check]
    
    Validator --> Report[Compliance Report]
    Report --> SyncGate[Block/Allow Sync]
```
