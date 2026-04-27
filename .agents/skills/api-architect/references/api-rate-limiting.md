# API Rate Limiting — Complete Reference

Guide to strategies for rate control and API resource protection.

---

## 1. Why Implement Rate Limiting?

- **Prevent Abuse**: Avoid Denial of Service (DoS) attacks.
- **Fair Usage**: Ensure that a single client does not consume all server resources.
- **Cost Control**: Limit the use of expensive computational resources (e.g., LLMs, third-party APIs).
- **Stability**: Maintain API performance under high load.

---

## 2. Algorithm Strategies

### 2.1 Token Bucket
A "bucket" of tokens that fills at a constant rate. Each request consumes a token.
- **Advantage**: Allows moderate traffic bursts.
- **Usage**: Market standard for most APIs.

### 2.2 Fixed Window Counter
Counts the number of requests in fixed time windows (e.g., 1 minute).
- **Advantage**: Simple to implement.
- **Disadvantage**: Can allow double the traffic at the edges of the windows.

### 2.3 Sliding Window Log / Sliding Window Counter
Uses a timestamp for each request to calculate the exact rate in real-time.
- **Advantage**: More accurate, avoids the fixed window edge problem.
- **Disadvantage**: Consumes more memory/CPU.

---

## 3. Implementation and Response

### 3.1 Status Code
Always return **429 Too Many Requests** when the limit is exceeded.

### 3.2 Standardized Headers
Inform the client about their current situation:

| Header | Description |
|--------|-----------|
| `X-RateLimit-Limit` | The maximum number of requests allowed in the period. |
| `X-RateLimit-Remaining` | How many requests are still left in the current window. |
| `X-RateLimit-Reset` | Timestamp (epoch) of when the limit will be reset. |
| `Retry-After` | Seconds the client must wait before trying again (optional). |

---

## 4. Where to Apply the Limit?

- **By IP**: Generic protection against bots.
- **By API Key / User**: Business model (e.g., Free tier vs. Premium).
- **By Endpoint**: Expensive endpoints (e.g., `/search`) should have lower limits than cheap endpoints (e.g., `/health`).

---

## 5. Response Example (JSON)

```json
{
  "error": {
    "code": "RATE_LIMIT_EXCEEDED",
    "message": "You have exceeded the request limit. Try again in 45 seconds.",
    "details": {
      "limit": 1000,
      "remaining": 0,
      "reset_at": 1713128400
    }
  }
}
```

---

## 6. Best Practices

1. **Graceful Degradation**: Warn the user before cutting them off completely (optional).
2. **Whitelist**: Allow known IPs or internal services without limits (e.g., health checks).
3. **Distributed Cache**: Use Redis to store rate limit counters in multi-server environments.
4. **Communication**: Clearly document limits in your API specification.
