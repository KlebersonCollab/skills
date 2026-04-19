# Regras de Mapeamento: Knowledge Map (Mermaid)

O Knowledge Map (`KNOWLEDGE-MAP.mermaid`) é a representação visual das relações entre entidades no Hub.

## 📐 Regras de Conexão

### 1. Tipos de Nós
- **Skill**: `node_id[Nome da Skill]`
- **Feature**: `feature_id((Nome da Feature))`
- **Arquitetura**: `arch_id{{Componente}}`

### 2. Definição de Relações
Use setas para indicar dependência ou fluxo:
- `A --> B`: A depende de B.
- `A -.-> B`: A referencia B (relação fraca).
- `A === B`: Conexão forte/bidirecional.

### 3. Estilização (Classes)
Mantenha a consistência visual usando as classes padrão:
```mermaid
classDef skill fill:#f9f,stroke:#333,stroke-width:2px;
classDef feature fill:#bbf,stroke:#333,stroke-width:1px;
```

## 📝 Boas Práticas
- **Granularidade**: Não mapeie cada função de código; mapeie módulos, padrões e decisões de design.
- **Top-Down**: Organize o grafo de cima para baixo (TD) para facilitar a leitura.
- **Labels**: Sempre use labels nas setas para explicar a relação (ex: `A -- implementa --> B`).

## 🔄 Atualização
Toda nova feature aprovada e concluída deve ter suas principais entidades adicionadas ao mapa para manter a integridade do Local Knowledge Graph (LKG).
