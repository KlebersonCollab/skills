# Exemplo de Relatório de Comparação de Performance

**Data**: 2026-04-22
**Branch**: `feature/optimize-auth-flow`
**Baseline**: `main` (2026-04-21)

## Resumo Executivo
A refatoração do fluxo de autenticação resultou em uma melhoria de **12% no LCP** e redução de **45KB no bundle JS**.

## Tabela de Métricas

| Métrica | Antes (Baseline) | Depois (PR) | Delta | Veredito |
|---------|------------------|-------------|-------|----------|
| **LCP** | 2.1s             | 1.85s       | -250ms | ✓ BETTER |
| **FCP** | 1.2s             | 1.1s        | -100ms | ✓ BETTER |
| **CLS** | 0.05             | 0.04        | -0.01 | PASS |
| **Bundle Size** | 240KB   | 195KB       | -45KB | ✓ BETTER |
| **API p95** | 120ms         | 115ms       | -5ms  | PASS |

## Análise Técnica
A redução do bundle foi alcançada através do uso de `dynamic imports` nos componentes de modais de login, que agora são carregados apenas sob demanda. A melhoria no LCP é reflexo direto do bundle menor e da otimização de prioridade da imagem de logo no header.

## Conclusão
O PR está **Aprovado** do ponto de vista de performance.
