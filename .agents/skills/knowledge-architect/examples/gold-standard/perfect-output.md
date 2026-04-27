# Gold Standard Output: Knowledge Mapping

## 1. Relational Mapping (KNOWLEDGE-MAP.mermaid)
The agent generated a clear mapping focusing on cause and effect (Impact Path):

```mermaid
graph LR
    REQ_REFUND["📄 Req: Partial Refund"] --(imp)--> SVC_PAY["⚙️ PaymentService"]
    SVC_PAY --(uses)--> GATEWAY["🔌 StripeGateway"]
    SVC_PAY --(updates)--> DB_TRANS["🗄️ TransactionsTable"]
    
    style REQ_REFUND fill:#f9f,stroke:#333
    style SVC_PAY fill:#bbf,stroke:#333
```

## 2. Rationale
This output is Gold Standard because:
- Identifies **Entities** (Requirements, Services, Gateways).
- Defines clear **Relationship Verbs** (`imp`, `uses`, `updates`).
- Uses Mermaid **Styling** to differentiate node types.
- Provides a **Cause and Effect** view (Impact Path).
