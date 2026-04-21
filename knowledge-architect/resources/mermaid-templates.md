# Mermaid Templates for Knowledge Mapping

## Template 1: Feature Dependency Graph

```mermaid
graph TD
    subgraph "Features"
        F1[Feature A]
        F2[Feature B]
        F3[Feature C]
    end

    subgraph "Components"
        C1[Component X]
        C2[Component Y]
        C3[Component Z]
    end

    subgraph "Data"
        D1[(Database)]
        D2[(Cache)]
    end

    F1 --> C1
    F1 --> C2
    F2 --> C2
    F3 --> C3
    C1 --> D1
    C2 --> D1
    C3 --> D2
```

## Template 2: Architecture Decision Impact

```mermaid
graph LR
    A[Decision: Use Event Sourcing] --> B[Benefits]
    A --> C[Costs]
    A --> D[Affected Components]

    B --> B1[Audit Trail]
    B --> B2[Temporal Queries]

    C --> C1[Complexity]
    C --> C2[Storage Overhead]

    D --> D1[Event Store]
    D --> D2[Projections]
    D --> D3[Event Handlers]
```

## Template 3: Module Relationships

```mermaid
graph TD
    subgraph "Domain Layer"
        Entity[Entities]
        Value[Value Objects]
        Aggregate[Aggregates]
    end

    subgraph "Application Layer"
        Service[Services]
        UseCase[Use Cases]
    end

    subgraph "Infrastructure Layer"
        Repo[Repositories]
        External[External Services]
    end

    UseCase --> Service
    Service --> Aggregate
    Aggregate --> Entity
    Service --> Repo
    Repo --> External
```

## Template 4: Data Flow

```mermaid
sequenceDiagram
    participant User
    participant API
    participant Service
    participant DB

    User->>API: Request
    API->>Service: Process
    Service->>DB: Query
    DB-->>Service: Data
    Service-->>API: Response
    API-->>User: Result
```

## Template 5: State Transitions

```mermaid
stateDiagram-v2
    [*] --> Created
    Created --> Pending
    Pending --> Processing
    Processing --> Completed
    Processing --> Failed
    Failed --> Pending
    Completed --> [*]
```

## Template 6: Component Hierarchy

```mermaid
graph TD
    App[Application] --> Layout[Layout]
    Layout --> Header[Header]
    Layout --> Main[Main]
    Layout --> Footer[Footer]

    Main --> Dashboard[Dashboard]
    Main --> Settings[Settings]

    Dashboard --> Widget1[Widget 1]
    Dashboard --> Widget2[Widget 2]

    Settings --> Profile[Profile]
    Settings --> Preferences[Preferences]
```

## Template 7: API Endpoint Relationships

```mermaid
graph LR
    A[GET /api/users] --> B[GET /api/users/:id]
    B --> C[POST /api/users]
    C --> D[PUT /api/users/:id]
    D --> E[DELETE /api/users/:id]

    B --> F[GET /api/users/:id/posts]
    F --> G[POST /api/users/:id/posts]
```

## Template 8: Technology Stack

```mermaid
graph TD
    subgraph "Frontend"
        FE1[React]
        FE2[Next.js]
        FE3[Tailwind]
    end

    subgraph "Backend"
        BE1[FastAPI]
        BE2[PostgreSQL]
        BE3[Redis]
    end

    subgraph "Infrastructure"
        INF1[Docker]
        INF2[Kubernetes]
        INF3[GitHub Actions]
    end

    FE2 --> BE1
    BE1 --> BE2
    BE1 --> BE3
    BE1 --> INF1
    INF1 --> INF2
    INF2 --> INF3
```

## Template 9: Risk Impact Analysis

```mermaid
graph TD
    R[Risk: Third-Party API Deprecation] --> I1[Impact Areas]
    R --> I2[Mitigation Strategies]

    I1 --> IA1[Feature Breakage]
    I1 --> IA2[Data Loss]
    I1 --> IA3[User Experience]

    I2 --> IS1[API Versioning]
    I2 --> IS2[Fallback Mechanism]
    I2 --> IS3[Monitoring]
```

## Template 10: Knowledge Clusters

```mermaid
graph TD
    subgraph "Cluster: User Management"
        UM1[User Entity]
        UM2[Auth Service]
        UM3[Profile Component]
    end

    subgraph "Cluster: Content Management"
        CM1[Content Entity]
        CM2[Editor Service]
        CM3[Display Component]
    end

    subgraph "Cluster: Analytics"
        AN1[Event Tracking]
        AN2[Reporting Service]
        AN3[Dashboard Component]
    end

    UM2 --> AN1
    CM2 --> AN1
```

## Usage Guidelines

1. **Choose the right template** based on what you're mapping
2. **Customize nodes** with your specific entities
3. **Add relationships** that matter for your context
4. **Use subgraphs** to group related concepts
5. **Keep it simple** - don't overcomplicate

## Styling Tips

```mermaid
graph TD
    A[Critical Node]:::critical
    B[Warning Node]:::warning
    C[Success Node]:::success

    classDef critical fill:#fee2e2,stroke:#dc2626,stroke-width:2px
    classDef warning fill:#fef3c7,stroke:#d97706,stroke-width:2px
    classDef success fill:#dcfce7,stroke:#16a34a,stroke-width:2px
```

## Best Practices

1. **Consistent naming** - Use clear, descriptive names
2. **Logical grouping** - Group related items together
3. **Directional flow** - Show the direction of relationships
4. **Color coding** - Use colors to indicate importance or type
5. **Documentation** - Add notes for complex relationships