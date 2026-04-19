# Padrões de Logging Estruturado (v1)

Guia para implementação de logs consistentes e integração com observabilidade moderna (OpenTelemetry).

## 1. Formato de Saída (JSON)
Todos os logs em produção **devem** ser gerados em formato JSON estruturado para facilitar a agregação por ELK, Splunk ou Datadog.

### Estrutura Base
```json
{
  "timestamp": "ISO-8601",
  "level": "INFO",
  "logger": "app.service.name",
  "message": "Mensagem descritiva",
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

## 2. Níveis de Log (Log Levels)
- **DEBUG**: Detalhes técnicos úteis para desenvolvimento local.
- **INFO**: Eventos significativos do sistema (ex: início/fim de processamento).
- **WARN**: Eventos inesperados que não interrompem o fluxo, mas exigem atenção.
- **ERROR**: Falhas críticas que impactam a execução de uma operação.
- **FATAL**: Erros que impedem a aplicação de continuar rodando.

## 3. Integração OpenTelemetry (OTel)
Para garantir o rastreamento distribuído, injete automaticamente o contexto de trace nos logs.

### Melhores Práticas:
1. **Correlation**: Sempre inclua `trace_id` e `span_id` em logs gerados durante uma requisição.
2. **Atributos de Recurso**: Identifique o serviço, versão e ambiente (staging/prod) nos metadados do recurso OTel.
3. **Eventos (Events)**: Prefira usar "Span Events" do OpenTelemetry para logs altamente granulares em vez de logs convencionais.

## 4. O que NÃO logar (Segurança)
- [ ] **Secrets**: Senhas, tokens API, chaves privadas.
- [ ] **PII (Personal Identifiable Information)**: CPF, e-mail completo sem anonimização, dados médicos, etc.
- [ ] **Payloads Gigantes**: Evite logar corpos de requisições binárias ou JSONs com milhares de linhas.
