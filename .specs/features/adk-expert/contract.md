# Contract: ADK Expert Skill

## Agreement
The implementation of the `adk-expert` skill must strictly follow the visual reference provided in the image "The Agent Development Kit".

## Deliverables
1. **Directory**: `adk-expert/`
2. **Main File**: `adk-expert/SKILL.md`
3. **Language**: Brazilian Portuguese (PT-BR).

## Validation Sensors
| ID | Sensor | Requirement | Verification Method |
|---|---|---|---|
| S-1 | Path Check | File exists at `/adk-expert/SKILL.md` | `ls adk-expert/SKILL.md` |
| S-2 | Metadata Check | Version is >= 1.0.0 and name is `adk-expert` | `head -n 10 adk-expert/SKILL.md` |
| S-3 | Coverage Check | All 5 ADK layers detailed | `grep -E "Layer|Camada" adk-expert/SKILL.md` |
| S-4 | Language Check | Text is in Portuguese | Manual Review / Language Detection pattern |

## Sign-off
- **Implementer**: Antigravity
- **Reviewer**: User
