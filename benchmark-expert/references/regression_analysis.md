# Análise de Regressão de Performance

Este guia explica como interpretar as mudanças de performance entre diferentes estados do código.

## Por que analisar regressões?
Código novo pode introduzir lentidão imperceptível em ambientes de desenvolvimento, mas crítica em produção. A análise de regressão quantifica esse impacto.

## Metodologia de Comparação
1. **Fator de Confiança**: Sempre execute o benchmark 3 vezes para eliminar ruídos estatísticos.
2. **Delta de Mudança**: Diferença percentual entre `After` e `Before`.
3. **Threshold de Alerta**:
   - < 5%: Variação normal.
   - 5% - 15%: Alerta (WARNING). Investigue a causa.
   - > 15%: Regressão Crítica (FAIL). Reverter ou otimizar.

## O Que Observar
- **JS Bundle Size**: O aumento do bundle impacta diretamente o TTI (Time to Interactive).
- **TTFB (Time to First Byte)**: Se subir, indica regressão no backend ou banco de dados.
- **LCP**: Se subir, indica problemas de priorização de recursos ou renderização pesada.

## Como Reportar
Use tabelas claras que destaquem as mudanças negativas em vermelho e positivas em verde.
Consulte [examples/comparison_report.md](../examples/comparison_report.md) para um modelo.
