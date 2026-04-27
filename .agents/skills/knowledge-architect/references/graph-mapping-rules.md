# Mapping Rules: Knowledge Map (Mermaid)

The Knowledge Map (`KNOWLEDGE-MAP.mermaid`) is the visual representation of relations between entities in the Hub.

## 📐 Connection Rules

### 1. Node Types
- **Skill**: `node_id[Skill Name]`
- **Feature**: `feature_id((Feature Name))`
- **Architecture**: `arch_id{{Component}}`

### 2. Relation Definition
Use arrows to indicate dependency or flow:
- `A --> B`: A depends on B.
- `A -.-> B`: A references B (weak relation).
- `A === B`: Strong/bidirectional connection.

### 3. Styling (Classes)
Maintain visual consistency using standard classes:
```mermaid
classDef skill fill:#f9f,stroke:#333,stroke-width:2px;
classDef feature fill:#bbf,stroke:#333,stroke-width:1px;
```

## 📝 Best Practices
- **Granularity**: Do not map every code function; map modules, patterns, and design decisions.
- **Top-Down**: Organize the graph top-down (TD) for easy reading.
- **Labels**: Always use labels on arrows to explain the relationship (e.g., `A -- implements --> B`).

## 🔄 Update
Every approved and completed new feature must have its main entities added to the map to maintain the integrity of the Local Knowledge Graph (LKG).
