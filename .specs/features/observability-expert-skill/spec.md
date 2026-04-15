# Spec: Observability Expert Skill

## 1. Goal
Create a specialized skill that acts as a Site Reliability Engineer (SRE) and Observability specialist. This skill will guide the agent in designing and implementing systems that are highly observable, resilient, and data-driven, focusing on the three pillars of observability: **Logs, Metrics, and Tracing**.

## 2. Target Audience
SREs, DevOps engineers, and senior developers using Gemini CLI to ensure their services are monitorable and maintainable in production.

## 3. Core Requirements (Functional)
- **Structured Logging Blueprint**: Guidelines for implementing JSON-based logging with contextual metadata (TraceID, UserID, Error codes).
- **OpenTelemetry Design**: Design for distributed tracing and metrics collection using industry-standard protocols.
- **Service Level Indicators (SLIs) & Objectives (SLOs)**: Assistance in defining critical business and technical metrics.
- **Grafana-as-Code & Dashboards**: Guidance for creating visual monitoring layouts via configuration files.
- **Alerting Policies**: Strategies for establishing actionable alerts, avoiding "alert fatigue".

## 4. Acceptance Criteria (BDD Format)

### Scenario 1: Define Structured Logging
**Given** a new backend service requiring production logs
**When** the `observability-expert` skill is invoked
**Then** it must provide a structured logging schema (JSON), including mandatory fields like `trace_id`, `timestamp`, `severity`, and `context`.

### Scenario 2: Distributed Tracing Setup
**Given** a microservices architecture with latency issues
**When** the `observability-expert` analyzes the flow
**Then** it must design an OpenTelemetry-based tracing strategy, showing how context is propagated between services.

### Scenario 3: Establish SLI/SLO Alerting
**Given** a requirement for 99.9% availability
**When** the `observability-expert` skill is in the `Delivery` phase
**Then** it must define SLIs (e.g., Error Rate, Latency) and suggest alerting thresholds based on SLO error budgets.

## 5. Technical Requirements (Non-Functional)
- **Standard-Aligned**: Must follow CNCF and OpenTelemetry best practices.
- **Visualization Mandate**: Must include Mermaid diagrams for monitoring flows and architecture.
- **Vendor Neutral**: Focus on open-source standards (Prometheus, Jaeger, Loki) rather than specific cloud vendors.

## 6. Architecture & Design
- `observability-expert/SKILL.md`: Entry point with the SRE workflow.
- `observability-expert/references/`:
    - `structured-logging.md`: Guidelines for logs that machines can read.
    - `metrics-and-tracing.md`: Deep dive into Prometheus and OpenTelemetry.
    - `sli-slo-alerting.md`: Methodology for data-driven reliability.
- `observability-expert/examples/`:
    - `log-format-sample.json`: Example of a well-structured log entry.
    - `observability-stack.mermaid`: Diagram of a standard monitoring stack.
