# Contrato de Entrega: Token Optimization

## 1. Acordo de Validação

Este documento define os critérios de sucesso para a implementação da feature **Token Optimization & Dual Mode**. A entrega será considerada completa após a validação dos sensores abaixo.

## 2. Sensores de Validação (SVP)

| ID | Sensor | Descrição | Evidência |
|---|---|---|---|
| **SVP-01** | **Linguistic Accuracy** | O modo Low Token deve reduzir a contagem de palavras em pelo menos 50% sem perder dados técnicos. | Log de chat comparativo. |
| **SVP-02** | **Toggle Reliability** | O comando `/mode low` e `/mode high` deve alterar instantaneamente o comportamento do agente. | Teste de fumaça (Smoke Test). |
| **SVP-03** | **Compression Safety** | O script `compressor.py` não deve corromper blocos de código ou IDs de tarefas. | Diff check em arquivos reais. |
| **SVP-04** | **SDD Integrity** | Todas as tarefas complexas devem permanecer no modo High Token por padrão para garantir rastreabilidade. | Verificação de estado no boot. |

## 3. Protocolo de Revisão

1. **Revisão de Código**: Validar se o `compressor.py` segue os princípios de Clean Code.
2. **Revisão de Prompt**: Validar se as instruções da skill `token-distiller` são claras o suficiente para evitar "alucinações de brevidade".
3. **Score Final**: A feature deve atingir um Score de conformidade > 90 no `validation-report.md`.

## 4. Assinatura do Contrato

- **Implementador**: Agente (Antigravity)
- **Revisor**: Usuário / SDD-Reviewer
- **Data**: 2026-04-23
