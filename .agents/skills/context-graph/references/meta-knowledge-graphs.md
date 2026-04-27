# Reference: Meta Knowledge Graphs (Context Graphs)

**Source**: [Substack - Firat Tekiner](https://firattekiner.substack.com/p/meta-knowledge-graphscontext-graphs)
**Topic**: The Memory Graph Layer for Agents

## Key Takeaways

1.  **The Memory Gap**: Most AI architectures lack long-term memory. Every project starts fresh, and lessons learned are lost.
2.  **Domain vs. Context**:
    *   **Domain Graph**: Records *what* exists (e.g., "Company X acquired Company Y").
    *   **Context Graph**: Records *how* and *why* it was built (e.g., "We chose this schema because it handles M&A events better").
3.  **The Chef Metaphor**:
    *   A good chef follows a recipe (Extraction pipeline).
    *   A great chef brings years of context (Meta-Knowledge). The **Context Graph** is the chef's notebook.
4.  **Types of Memory**:
    *   **Episodic**: Specific decision traces (events).
    *   **Semantic**: Generalized patterns and tribal knowledge (facts).
    *   **Procedural**: Proven behavior strategies (skills/workflows).
5.  **Multi-Agent Coordination**: The Context Graph acts as the "Shared Memory" (Blackboard pattern) for a swarm of specialized agents (Document Agent, Schema Agent, extraction Agent, etc.).
6.  **Layered Context**: Guidance should flow from Specific (Project) -> General (Industry/Universal).

## Implementation Ideas for Skills Hub

*   Maintain a `DECISIONS.md` file in every project to track the Context Graph manually until automated tools are built.
*   Use Mermaid diagrams to visualize the relationships between decisions and outcomes.
*   Query past project `LEARNINGS.md` as part of the Context Graph retrieval phase.
