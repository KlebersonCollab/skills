---
name: observability-expert
version: 2.2.0
description: "Skill for SRE and Observability specialist. Focuses on Structured Logs, OpenTelemetry, SLIs/SLOs, and proactive monitoring to ensure system resilience."
category: site-reliability-engineering
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify `.hub-mode` and apply `token-distiller` guidelines.
1. **Context Check**: Rehydrate state by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`.
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define the architecture, schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---
# Observability Expert

> "If you can't measure it, you can't manage it." — Peter Drucker. The focus here is on transforming raw data into actionable insights and resilient systems.

---

## Goal

Empower the agent to design and implement high-performance observability ecosystems using the pillars of **Logs, Metrics, and Tracing**. The skill ensures that the system is transparent for the operations team, allowing rapid failure detection, root cause analysis, and evolution based on real usage data.

---

## Workflow (4 Phases)

### Phase 1: INSTRUMENT — Coding for Visibility
1.  **Map Context**: Identify crucial metadata (TraceID, UserID, TenantID) that must accompany each operation.
2.  **Logging Design**: Apply the **Structured Logging** (JSON) standard at critical code points (API entry, queries, errors).
3.  **SDK Setup**: Suggest the integration of **OpenTelemetry** SDKs for automatic capture of spans and system metrics.

### Phase 2: COLLECT — Telemetry Orchestration
1.  **Data Pipeline**: Design data flow from application to storage (e.g., OTel Collector -> Prometheus/Loki).
2.  **Sampling Strategy**: Define sampling strategies for traces to balance cost and visibility at high scale.
3.  **Metrics Export**: Configure export of standard metrics (RED: Rate, Errors, Duration) and business metrics.

### Phase 3: VISUALIZE — Dashboards & Insights
1.  **Dashboard Layout**: Design panels that tell a story (from macro to micro).
2.  **SLI/SLO Mapping**: Translate business requirements into monitorable technical indicators.
3.  **Correlation**: Ensure that by clicking an error in the log, it's possible to jump directly to the related trace.

### Phase 4: REACT — Incident and Alert Management
1.  **Alerting Rules**: Define alerts based on **Error Budget** burn rate instead of simple static thresholds.
2.  **Actionable Alerts**: Ensure every alert has a link to a fix playbook or diagnostic dashboard.
3.  **Feedback Loop**: Use performance data to suggest code or infrastructure optimizations.

---

## Quality Rules

- **Structured-First**: Every production log **MUST** be emitted in structured format (JSON).
- **Vendor Agnostic**: Prefer open standards (**OpenTelemetry**, **Prometheus**) over proprietary cloud solutions.
- **Trace Propagation**: Every call between services must propagate the trace context header (W3C Trace Context).
- **Actionable Only**: Do not create alerts that do not require immediate action; avoid alert fatigue.

## Prohibited

- **NEVER** log sensitive data (PII, passwords, tokens) in plain text.
- **NEVER** use pure strings for system logs; use fixed keys to facilitate search.
- **NEVER** ignore telemetry storage cost; use sampling when necessary.
- **NEVER** consider a system "production-ready" without a defined and monitored SLO.

---

## Reference Documentation

This skill includes detailed reference documentation:

1. **[Structured Logging](references/structured-logging.md)** — JSON patterns and contextual metadata.
2. **[Metrics & Tracing](references/metrics-and-tracing.md)** — Practical guide to OpenTelemetry and Prometheus.
3. **[SLIs, SLOs & Alerting](references/sli-slo-alerting.md)** — Data-driven reliability methodology.
4. **[Framework Monitoring](references/framework_monitoring.md)** — Specific monitoring for FastAPI and Django (Prometheus/Grafana).

---

## Output Structure

Execution of this skill results in the following standardized artifacts:

| Artifact | Format | Description |
|----------|---------|-----------|
| **Logging Schema** | `.json` | Definition of fields and types for structured logs. |
| **OTel Blueprint** | Mermaid | Diagram of context collection and propagation pipeline. |
| **SLO Report** | `.md` | Formal definition of indicators, objectives, and error budgets. |
| **Alerting Policy** | `.yaml` | Alerting rules and notification guidelines. |


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.375595Z"
evidence_checksum: "NONE"
```
