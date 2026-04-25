# Contract: Benchmark Expert Skill

## Command Interface
| Command | Mode | Description |
| :--- | :--- | :--- |
| `/benchmark baseline` | 1, 2, 3 | Salva o baseline de performance atual. |
| `/benchmark compare` | 4 | Compara o estado atual com o baseline salvo. |

## Mandatory Artifacts
- **File**: `SKILL.md`
- **File**: `README.md`
- **File**: `references/performance_targets.md`

## Exit Veredicts
- `SUCCESS`: Performance mantida ou melhorada.
- `WARNING`: Regressão leve detectada.
- `FAIL`: Regressão crítica ou violação de SLA.
