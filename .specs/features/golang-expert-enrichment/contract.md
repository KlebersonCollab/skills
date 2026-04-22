# Contract: Golang Expert Skill Enrichment

## 1. Acordo de Entrega
O Implementador compromete-se a entregar a skill `golang-expert` enriquecida com os padrões do ECC, garantindo que toda a documentação esteja em PT-BR e siga os mandatos globais do Hub.

## 2. Sensores de Validação

| ID | Sensor | Critério de Sucesso | Evidência |
|----|--------|---------------------|-----------|
| **S-1** | Arquitetura | `SKILL.md` contém referências a Interfaces, Functional Options e Zero Value. | [SKILL.md:L70-85] |
| **S-2** | Concorrência | `concurrency.md` contém exemplos de `errgroup` e shutdown. | [concurrency.md] |
| **S-3** | Restrições | Seção `Prohibited` veta Context em structs. | [SKILL.md:L80-90] |
| **S-4** | Idioma | Todos os textos explicativos estão em Português Brasileiro. | Inspeção Manual |
| **S-5** | Integridade | Todos os links de referências estão funcionais. | Inspeção Manual |

## 3. Definição de Pronto (DoD)
- Spec, Plan e Contract aprovados.
- Implementação concluída em todos os arquivos alvo.
- Nenhuma violação de mandatos globais detectada.
- Registro da evolução no `STATE.md` e `LEARNINGS.md`.
- Relatório de validação (`validation-report.md`) gerado com Score >= 95.
