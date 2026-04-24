# Logging Standards

Guide for implementing consistent logs and integration with modern observability (OpenTelemetry).

## 1. Output Format (JSON)
All logs in production **must** be generated in structured JSON format to facilitate aggregation by ELK, Splunk, or Datadog.

### Base Structure
```json
{
  "timestamp": "ISO-8601",
  "level": "INFO",
  "logger": "app.service.name",
  "message": "Descriptive message",
  "context": {
    "trace_id": "opentelemetry-trace-id",
    "span_id": "opentelemetry-span-id",
    "user_id": "optional-user-context",
    "request_id": "correlation-id"
  },
  "payload": {
    "custom_data_field": "value"
  }
}
```

## 2. Log Levels
- **DEBUG**: Technical details useful for local development.
- **INFO**: Significant system events (e.g., start/end of processing).
- **WARN**: Unexpected events that do not interrupt the flow but require attention.
- **ERROR**: Critical failures that impact the execution of an operation.
- **FATAL**: Errors that prevent the application from continuing to run.

## 3. OpenTelemetry (OTel) Integration
To ensure distributed tracing, automatically inject the trace context into the logs.

### Best Practices:
1. **Correlation**: Always include `trace_id` and `span_id` in logs generated during a request.
2. **Resource Attributes**: Identify the service, version, and environment (staging/prod) in the OTel resource metadata.
3. **Events**: Prefer using OpenTelemetry "Span Events" for highly granular logs instead of conventional logs.

## 4. What NOT to Log (Security)
- [ ] **Secrets**: Passwords, API tokens, private keys.
- [ ] **PII (Personal Identifiable Information)**: SSN/Tax IDs, full email without anonymization, medical data, etc.
- [ ] **Giant Payloads**: Avoid logging binary request bodies or JSONs with thousands of lines.
