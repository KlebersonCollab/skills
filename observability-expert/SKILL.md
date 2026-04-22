---
name: observability-expert
version: 1.0.0
description: "Skill para especialista em SRE e Observabilidade. Foca em Logs Estruturados, OpenTelemetry, SLIs/SLOs e monitoramento proativo para garantir a resiliência de sistemas."
category: site-reliability-engineering
---

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar qualquer execução técnica:
1. **Context Check**: Você reidratou o contexto lendo `STATE.md`, `MEMORY.md` e `LEARNINGS.md`?
2. **Spec Check**: O arquivo `spec.md` existe com requisitos e Critérios de Aceitação (ACs) claros? (BDD mandatório para Medium+).
3. **Plan Check**: O arquivo `plan.md` define a arquitetura, schemas e inclui diagramas **Mermaid**?
4. **Contract Check**: O arquivo `contract.md` foi estabelecido com os sensores de validação?
5. **Task Check**: A lista de tarefas em `tasks.md` está detalhada e atomizada?

---
# Observability Expert

> "Se você não pode medir, você não pode gerenciar." — Peter Drucker. O foco aqui é transformar dados brutos em insights acionáveis e sistemas resilientes.

---

## Goal

Capacitar o agente a projetar e implementar ecossistemas de observabilidade de alta performance, utilizando os pilares de **Logs, Métricas e Tracing**. A skill assegura que o sistema seja transparente para o time de operação, permitindo detecção rápida de falhas, análise de causa raiz e evolução baseada em dados reais de uso.

---

## Workflow (4 Fases)

### Fase 1: INSTRUMENT — Codificação para Visibilidade
1.  **Mapear Contexto**: Identificar os metadados cruciais (TraceID, UserID, TenantID) que devem acompanhar cada operação.
2.  **Logging Design**: Aplicar o padrão de **Structured Logging** (JSON) em pontos críticos do código (entrada de API, queries, erros).
3.  **SDK Setup**: Sugerir a integração de SDKs do **OpenTelemetry** para captura automática de spans e métricas de sistema.

### Fase 2: COLLECT — Orquestração de Telemetria
1.  **Pipeline de Dados**: Desenhar o fluxo dos dados desde a aplicação até o armazenamento (ex: OTel Collector -> Prometheus/Loki).
2.  **Sampling Strategy**: Definir estratégias de amostragem para traces para equilibrar custo e visibilidade em alta escala.
3.  **Metrics Export**: Configurar a exportação de métricas padrão (RED: Rate, Errors, Duration) e métricas de negócio.

### Fase 3: VISUALIZE — Dashboards & Insights
1.  **Dashboard Layout**: Projetar painéis que contem uma história (do macro para o micro).
2.  **SLI/SLO Mapping**: Traduzir requisitos de negócio em indicadores técnicos monitoráveis.
3.  **Correlation**: Garantir que, ao clicar em um erro no log, seja possível pular direto para o trace relacionado.

### Fase 4: REACT — Gestão de Incidentes e Alertas
1.  **Alerting Rules**: Definir alertas baseados em queima de **Error Budget** em vez de limiares estáticos simples.
2.  **Actionable Alerts**: Garantir que todo alerta tenha um link para o playbook de correção ou dashboard de diagnóstico.
3.  **Feedback Loop**: Utilizar dados de performance para sugerir otimizações de código ou infraestrutura.

---

## Quality Rules

- **Structured-First**: Todo log de produção **DEVE** ser emitido em formato estruturado (JSON).
- **Vendor Agnostic**: Preferir padrões abertos (**OpenTelemetry**, **Prometheus**) sobre soluções proprietárias de cloud.
- **Trace Propagation**: Toda chamada entre serviços deve propagar o cabeçalho de contexto de rastreio (W3C Trace Context).
- **Actionable Only**: Não crie alertas que não exijam uma ação imediata; evite a fadiga de alertas.

## Prohibited

- **NUNCA** logar dados sensíveis (PII, senhas, tokens) em texto claro.
- **NUNCA** usar strings puras para logs de sistema; utilize chaves fixas para facilitar a busca.
- **NUNCA** ignorar o custo de armazenamento de telemetria; use amostragem (sampling) quando necessário.
- **NUNCA** considerar um sistema "pronto para produção" sem um SLO definido e monitorado.

---

## Reference Documentation

Esta skill inclui documentação de referência detalhada:

1. **[Structured Logging](references/structured-logging.md)** — Padrões JSON e metadados contextuais.
2. **[Metrics & Tracing](references/metrics-and-tracing.md)** — Guia prático de OpenTelemetry e Prometheus.
3. **[SLIs, SLOs & Alerting](references/sli-slo-alerting.md)** — Metodologia de confiabilidade orientada a dados.
4. **[Framework Monitoring](references/framework_monitoring.md)** — Monitoramento específico para FastAPI e Django (Prometheus/Grafana).

---

## Output Structure

A execução desta skill resulta nos seguintes artefatos padronizados:

| Artefato | Formato | Descrição |
|----------|---------|-----------|
| **Logging Schema** | `.json` | Definição dos campos e tipos para logs estruturados. |
| **OTel Blueprint** | Mermaid | Diagrama do pipeline de coleta e propagação de contexto. |
| **SLO Report** | `.md` | Definição formal de indicadores, objetivos e orçamentos de erro. |
| **Alerting Policy** | `.yaml` | Regras de alerta e pautas de notificação. |
