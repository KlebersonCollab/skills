# Contract: SDD Reverse Reconciler

## CLI Interface

| Command | Args | Flags | Description |
| :--- | :--- | :--- | :--- |
| `hb sdd reconcile` | `[feature]` | `--apply` | Sincroniza docs com a implementação real. |

## Expected Behavior

1.  **Scanning**: O auditor identifica itens pendentes em features com progresso < 100%.
2.  **Evidence**: Busca no código sinais de que o item foi implementado.
3.  **Sync**: 
    *   Marca `[x]` em `tasks.md`.
    *   Cria `contract.md` se inexistente.
4.  **Verification**: O progresso no `hb sdd audit` deve subir após a execução.
