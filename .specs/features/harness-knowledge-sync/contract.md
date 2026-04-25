# Contract: Harness Knowledge Sync

## CLI Interface

| Command | Flags | Description |
| :--- | :--- | :--- |
| `hb harness map` | `--auto` | Analisa alterações e sugere atualizações no mapa. |
| `hb harness map` | `--apply` | Aplica as sugestões diretamente no KNOWLEDGE-MAP.mermaid. |

## Expected Output

```text
🔍 Analisando ecossistema...

Novas conexões detectadas:
 - feature(harness-audit-gate) --> skill(harness-expert) [Missing]
 - feature(harness-distill) --> skill(token-distiller) [Missing]

Deseja aplicar estas mudanças? (use --apply para automatizar)
```
