# Contract: Governance Hardening

## 1. Acordo de Entrega
O agente compromete-se a endurecer as regras de conformidade de Git e Testes no Hub, garantindo que o ciclo atômico de desenvolvimento seja impossível de ignorar sem violar mandatos explícitos.

## 2. Sensores de Validação

| ID | Sensor | Critério de Sucesso | Evidência |
|----|--------|---------------------|-----------|
| **S-1** | SDD Mandate | Mandato proibitivo presente na Fase 3 da `sdd`. | [sdd/SKILL.md] |
| **S-2** | Exit Gate | Onboarding Navigator exige auditoria de commits. | [onboarding-navigator/SKILL.md] |
| **S-3** | Versioning | Versões das skills afetadas foram incrementadas. | frontmatter YAML |
| **S-4** | Git Compliance | Esta feature deve ter commits atômicos por task. | `git log` |

## 3. Definição de Pronto (DoD)
- Skills Core atualizadas com os mandatos.
- Relatório de validação emitido com aprovação de 100%.
- Aprendizados registrados sobre a falha de conformidade anterior.
