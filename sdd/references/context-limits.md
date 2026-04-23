# Limites de Contexto e Tamanho de Artefatos

A gestão de contexto é vital para manter a capacidade de raciocínio lógico em tarefas longas. Siga rigorosamente estes limites.

## Limites de Arquivo

| Arquivo         | Max Tokens | ~Palavras | Aviso em (Warning) |
| --------------- | ---------- | --------- | ------------------ |
| PROJECT.md      | 2.000      | 1.200     | 1.600 (80%)        |
| ROADMAP.md      | 3.000      | 1.800     | 2.400              |
| STATE.md        | 10.000     | 6.000     | 7.000 (70%)        |
| spec.md         | 5.000      | 3.000     | 4.000              |
| design.md       | 8.000      | 4.800     | 6.400              |
| tasks.md        | 10.000     | 6.000     | 8.000              |
| STACK.md        | 2.000      | 1.200     | 1.600              |
| ARCHITECTURE.md | 4.000      | 2.400     | 3.200              |
| CONVENTIONS.md  | 3.000      | 1.800     | 2.400              |
| STRUCTURE.md    | 2.000      | 1.200     | 1.600              |
| TESTING.md      | 4.000      | 2.400     | 3.200              |
| INTEGRATIONS.md | 5.000      | 3.000     | 4.000              |

## Zonas de Contexto
- 🟢 **Saudável** (<40k tokens totais): Silencioso.
- 🟡 **Moderado** (40-60k): Aviso discreto no rodapé.
- 🔴 **Crítico** (>60k): Aviso ativo, sugere otimização do contexto (ex: arquivar logs no STATE.md).

## Princípios de Operação
1. **Alvo:** Manter sempre menos de 40k tokens ativos carregados na memória de contexto.
2. **Reserva Operacional:** Garantir que o modelo sempre tenha mais de 160k tokens disponíveis livres para raciocínio e saída.
3. **Carregamento Sob Demanda:** Evite carregar múltiplas specs de features diferentes ao mesmo tempo. Carregue artefatos apenas quando forem estritamente necessários para a tarefa atual.
