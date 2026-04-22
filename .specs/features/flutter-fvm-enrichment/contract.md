# Development Contract: Flutter FVM Skill Enrichment

## Agreement
Este contrato define os sensores de qualidade e critérios de entrega para a feature de enriquecimento da skill `flutter-fvm` (v1.3.0).

## Deliverables
- [ ] `flutter-fvm/SKILL.md` atualizado.
- [ ] `flutter-fvm/references/modern-dart-patterns.md` (Novo).
- [ ] `flutter-fvm/references/performance-and-optimization.md` (Novo).
- [ ] `flutter-fvm/references/accessibility-guide.md` (Novo).
- [ ] `flutter-fvm/references/testing-and-quality.md` atualizado.
- [ ] `flutter-fvm/references/flutter-security-guide.md` atualizado.
- [ ] `flutter-fvm/references/project-structure.md` atualizado.

## Validation Sensors (Quality Gates)
| Sensor | Method | Threshold |
|---|---|---|
| **Structural Integrity** | `ls` e `view_file` | Todos os 7 arquivos de referência devem existir e estar linkados no `SKILL.md`. |
| **Language Consistency** | `grep` | Proibido o uso de termos em inglês onde houver tradução consagrada (ex: usar "Alvos de Toque" em vez de apenas "Hit Targets"). |
| **FVM Compliance** | `grep` | Todos os comandos sugeridos DEVEM usar o prefixo `fvm`. |
| **Modern Dart** | `grep` | Deve haver menção explícita a `sealed class` e `MediaQuery.sizeOf`. |
| **A11y Check** | `grep` | Menção obrigatória a `Semantics` e `48x48`. |

## Definition of Done (DoD)
1. Todas as tarefas em `tasks.md` marcadas como `DONE`.
2. Todos os sensores acima PASS.
3. Changelog atualizado.
4. `STATE.md` e `LEARNINGS.md` sincronizados.
