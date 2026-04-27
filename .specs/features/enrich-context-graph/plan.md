# Plan: Enrichment of Context Graph (MemDir Protocol)

## 1. Architecture
Enrich `context-graph/SKILL.md` by merging conceptual memory (episodic/semantic) with the technical implementation protocol from `src/memdir`.

## 2. Technical Protocol Details
- **Index Entry Format**: `- [Title](file.md) — One-line hook (max 150 chars)`.
- **Topic File Format**:
  ```markdown
  ---
  name: Topic Name
  description: Short summary
  type: User | Project | Local | Managed | AutoMem
  ---
  Detailed content here.
  ```
- **Size Caps**: 200 lines / 25,000 bytes for `MEMORY.md`.

## 3. Implementation Steps
1.  **Enrich `SKILL.md`**: Add "Technical Implementation (MemDir Protocol)" section.
2.  **Update `CHANGELOG.md`**: Record v1.1.0.
3.  **Update `MEMORY.md`** of the skill (optional but good for dogfooding).
4.  **Audit**: `hb audit context-graph`.

## 4. Mermaid Diagram (Persistence Hierarchy)
```mermaid
graph TD
    A[Information Type] --> B{Lifetime?};
    B -- "Future Conversations" --> C[Memory (Topic File)];
    B -- "Current Conversation Only" --> D{Complexity?};
    D -- "Implementation/Alignment" --> E[Plan];
    D -- "Discrete Steps" --> F[Tasks];
    C --> G[Update MEMORY.md Index];
```
