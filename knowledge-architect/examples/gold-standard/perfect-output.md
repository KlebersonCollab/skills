# Gold Standard Output: Knowledge Mapping

## 1. Relational Mapping (KNOWLEDGE-MAP.mermaid)

```mermaid
graph LR
    REQ_REFUND["📄 Req: Estorno Parcial"] --(imp)--> SVC_PAY["⚙️ PaymentService"]
    SVC_PAY --(uses)--> GATEWAY["🔌 StripeGateway"]
    SVC_PAY --(updates)--> DB_TRANS["🗄️ TransactionsTable"]
    
    style REQ_REFUND fill:#f9f,stroke:#333
    style SVC_PAY fill:#bbf,stroke:#333
```

## 2. Rationale
Este output é Gold Standard porque:
- Identifica **Entidades** (Requisitos, Serviços, Gateways).
- Define **Verbos de Relação** claros (`imp`, `uses`, `updates`).
- Utiliza **Styling** Mermaid para diferenciar tipos de nós.
- Fornece uma visão de **Causa e Efeito** (Caminhamento de Impacto).
