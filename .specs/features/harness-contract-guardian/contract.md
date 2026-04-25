# Contract: Harness Contract Guardian

## CLI Interface

| Command | Args | Description |
| :--- | :--- | :--- |
| `hb harness contract` | `<feature>` | Valida a implementação contra o contract.md da feature. |

## Expected Behavior

1.  **Input**: A feature directory containing `contract.md`.
2.  **Process**:
    *   Find tables with headers like `| Command |` or `| Interface |`.
    *   Extract the first column.
    *   Search for these names in the project root.
3.  **Output**:
    *   Status list: `✅ Found` or `❌ Missing`.

## Exit Codes
- `0`: All contracts implemented.
- `1`: Missing implementation for one or more contracts.
