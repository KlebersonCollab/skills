# Structured Logging: Machine-Readable Insights

Structured logging is the practice of emitting logs in a fixed format (usually JSON) so they can be easily queried by tools like ELK, Loki, or Splunk.

---

## 1. Why not simple strings?
Pure strings require complex regular expressions for analysis and are hard to filter. JSON allows you to instantly filter by `user_id` or `trace_id` across millions of entries.

## 2. Suggested Standard Schema

Every log should contain these base fields:

```json
{
  "timestamp": "2026-04-14T20:00:00.000Z",
  "severity": "ERROR",
  "trace_id": "abc-123-def",
  "message": "Failed to process payment",
  "context": {
    "user_id": "42",
    "order_id": "999",
    "error_code": "INSUFFICIENT_FUNDS"
  },
  "service": "payment-gateway",
  "version": "1.2.0"
}
```

## 3. Severity Levels (General)

- **DEBUG**: Technical details for development.
- **INFO**: Normal system events (e.g., "Service started").
- **WARN**: Something unexpected, but the system continues operating.
- **ERROR**: Failure in a specific operation that requires attention.
- **CRITICAL**: Total systemic failure.

## 4. What to log?

- **API Input/Output**: Status code, response time, and correlation IDs.
- **Database Queries**: Only when slow or in error.
- **External Integrations**: Network failures and timeouts.
- **State Changes**: e.g., "Order moved from PENDING to PAID".

## 5. What NEVER to log (Prohibited)

- **PII (Personally Identifiable Information)**: SSN, full name, address.
- **Secrets**: Passwords, JWT tokens, API keys.
- **Financial Data**: Credit card number, CVV.

## 6. Recommended Tools

- **Loki**: Excellent for Kubernetes environments (Grafana stack).
- **Elasticsearch (ELK)**: The gold standard for complex search.
- **Zap (Go)**, **Pino (Node)**, **Loguru (Python)**: High-performance libraries.


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.377036Z"
evidence_checksum: "NONE"
```
