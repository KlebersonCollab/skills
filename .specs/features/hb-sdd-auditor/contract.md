# Contract: SDD Auditor

## CLI Interface

| Command | Args | Description |
| :--- | :--- | :--- |
| `hb sdd audit` | None | Audita todas as features e exibe relatório de conformidade. |

## Expected Output Format

| Feature | Spec | Plan | Contract | Tasks | Progress |
| :--- | :--- | :--- | :--- | :--- | :--- |
| feature-name | ✅ | ✅ | ❌ | ✅ | 85% |
| legacy-feat | ⚠️ | ✅ | ✅ | ✅ | 100% |

- ✅: File exists.
- ❌: File missing.
- ⚠️: Legacy file used (e.g. PRD.md instead of spec.md).
