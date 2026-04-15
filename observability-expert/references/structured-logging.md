# Structured Logging: Machine-Readable Insights

O logging estruturado é a prática de emitir logs em um formato fixo (geralmente JSON) para que possam ser facilmente consultados por ferramentas como ELK, Loki ou Splunk.

---

## 1. Por que não strings simples?
Strings puras exigem expressões regulares complexas para análise e são difíceis de filtrar. JSON permite que você filtre instantaneamente por `user_id` ou `trace_id` em milhões de entradas.

## 2. Esquema Padrão Sugerido

Todo log deve conter estes campos base:

```json
{
  "timestamp": "2026-04-14T20:00:00.000Z",
  "severity": "ERROR",
  "trace_id": "abc-123-def",
  "message": "Falha ao processar pagamento",
  "context": {
    "user_id": "42",
    "order_id": "999",
    "error_code": "INSUFFICIENT_FUNDS"
  },
  "service": "payment-gateway",
  "version": "1.2.0"
}
```

## 3. Níveis de Severidade (Geral)

- **DEBUG**: Detalhes técnicos para desenvolvimento.
- **INFO**: Eventos normais do sistema (ex: "Serviço iniciado").
- **WARN**: Algo inesperado, mas o sistema continua operando.
- **ERROR**: Falha em uma operação específica que exige atenção.
- **CRITICAL**: Falha sistêmica total.

## 4. O que logar?

- **Entrada/Saída de API**: Status code, tempo de resposta e IDs de correlação.
- **Queries de Banco**: Apenas quando lentas ou em erro.
- **Integrações Externas**: Falhas de rede e timeouts.
- **Mudanças de Estado**: Ex: "Pedido movido de PENDING para PAID".

## 5. O que NUNCA logar (Prohibited)

- **PII (Personally Identifiable Information)**: CPF, nome completo, endereço.
- **Segredos**: Senhas, tokens JWT, chaves de API.
- **Dados Financeiros**: Número de cartão de crédito, CVV.

## 6. Ferramentas Recomendadas

- **Loki**: Excelente para ambientes Kubernetes (Grafana stack).
- **Elasticsearch (ELK)**: O padrão de ouro para busca complexa.
- **Zap (Go)**, **Pino (Node)**, **Loguru (Python)**: Bibliotecas de alto desempenho.
