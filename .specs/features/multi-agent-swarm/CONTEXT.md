
# Glossary — Multi-Agent Swarm

## Domain Terms

| Term | Definition |
|------|------------|
| **Swarm** | A coordinated group of AI agents working together on a shared goal, with defined roles, communication protocols, and orchestration. |
| **Agent** | An autonomous AI entity with specific capabilities (tools, skills, LLM provider) that can execute tasks within the swarm. |
| **Orchestrator** | The central agent responsible for decomposing tasks, delegating to worker agents, and merging results. |
| **Worker** | A specialized agent that executes specific subtasks and reports results back to the orchestrator. |
| **Generator** | An agent role focused on creating/proposing solutions (GAN pattern). |
| **Discriminator** | An agent role focused on auditing/critiquing solutions (GAN pattern). |
| **Swarm Topology** | The communication structure of the swarm (e.g., centralized, decentralized, hierarchical). |
| **Baton Pass** | The handoff protocol for passing control between agents in a sequence. |

## Current State

The Harness Agent currently operates as a **single agent** — one LLM call, one execution loop, one session tree. Multi-Agent Swarm would extend this to support multiple agents collaborating.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "multi-agent-swarm"
phase: "DISCOVERY"
status: "IN_PROGRESS"
last_update: "2026-05-23T17:50:00Z"
evidence_checksum: "discovery-glossary"
```
