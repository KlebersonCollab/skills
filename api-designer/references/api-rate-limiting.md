# API Rate Limiting — Referência Completa

Guia de estratégias para controle de taxa e proteção de recursos de API.

---

## 1. Por que implementar Rate Limiting?

- **Prevenir Abusos**: Evitar ataques de Denial of Service (DoS).
- **Justiça de Uso**: Garantir que um único cliente não consuma todos os recursos do servidor.
- **Controle de Custos**: Limitar o uso de recursos computacionais caros (ex: LLMs, APIs de terceiros).
- **Estabilidade**: Manter a performance da API sob carga alta.

---

## 2. Estratégias de Algoritmo

### 2.1 Token Bucket
Um "balde" de tokens que enche em uma taxa constante. Cada request consome um token.
- **Vantagem**: Permite picos moderados de tráfego.
- **Uso**: Padrão de mercado para a maioria das APIs.

### 2.2 Fixed Window Counter
Conta o número de requests em janelas de tempo fixas (ex: 1 minuto).
- **Vantagem**: Simples de implementar.
- **Desvantagem**: Pode permitir o dobro do tráfego na borda das janelas.

### 2.3 Sliding Window Log / Sliding Window Counter
Usa um timestamp para cada request para calcular a taxa exata no tempo real.
- **Vantagem**: Mais preciso, evita o problema das bordas da janela fixa.
- **Desvantagem**: Consome mais memória/CPU.

---

## 3. Implementação e Resposta

### 3.1 Status Code
Sempre retorne **429 Too Many Requests** quando o limite for excedido.

### 3.2 Headers Padronizados
Informe o cliente sobre sua situação atual:

| Header | Descrição |
|--------|-----------|
| `X-RateLimit-Limit` | O número máximo de requisições permitidas no período. |
| `X-RateLimit-Remaining` | Quantas requisições ainda restam na janela atual. |
| `X-RateLimit-Reset` | Timestamp (epoch) de quando o limite será resetado. |
| `Retry-After` | Segundos que o cliente deve esperar antes de tentar novamente (opcional). |

---

## 4. Onde Aplicar o Limite?

- **Por IP**: Proteção genérica contra bots.
- **Por API Key / Usuário**: Modelo de negócio (ex: Free tier vs Premium).
- **Por Endpoint**: Endpoints caros (ex: `/search`) devem ter limites menores que endpoints baratos (ex: `/health`).

---

## 5. Exemplo de Resposta (JSON)

```json
{
  "error": {
    "code": "RATE_LIMIT_EXCEEDED",
    "message": "Você excedeu o limite de requisições. Tente novamente em 45 segundos.",
    "details": {
      "limit": 1000,
      "remaining": 0,
      "reset_at": 1713128400
    }
  }
}
```

---

## 6. Boas Práticas

1. **Graceful Degradation**: Avise o usuário antes de cortá-lo totalmente (opcional).
2. **Whitelist**: Permita IPs conhecidos ou serviços internos sem limite (ex: health checks).
3. **Cache Distribuído**: Use Redis para armazenar os contadores de rate limit em ambientes de múltiplos servidores.
4. **Comunicação**: Documente claramente os limites na sua especificação de API.
