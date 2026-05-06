# Swarm Execution: Infrastructure & Personas

This reference defines how to execute SDD workflows in multi-agent environments, focusing on infrastructure stability and persona specialization.

---

## 🎭 Multi-Agent Personas
In a Swarm/Crew execution, agents assume specialized roles to maintain focus and context economy:

| Persona | SDD Phase | Core Skills | Primary Mission |
|---|---|---|---|
| **Solution Architect** | Phase 1 & 2 | `architecture`, `brainstorming` | Define PRDs, ADRs, Specs, and Plans. |
| **Lead Developer** | Phase 3 | `fastapi-expert`, `frontend-expert` | Implementation of atomic tasks and TDD. |
| **Reviewer/SRE** | Phase 4 | `clean-code-mentor`, `observability` | Audit delivery against Contract and ACs. |

---

## 🚩 Terminal Backends
To maintain persistent or visual sessions across agent handoffs:

1.  **Tmux (Headless/Persistent)**:
    *   **Use Case**: Long-running background tasks.
    *   **Benefit**: Survives network disconnections. Use `SessionID` to re-attach.
2.  **iTerm2 (Visual)**:
    *   **Use Case**: Local debugging where the user must see the terminal state.
    *   **Requirement**: Use `It2SetupPrompt` logic.
3.  **In-Process (Ephemeral)**:
    *   **Use Case**: Simple reads, linting, or quick commands.

---

## 🛡️ Permission Bridge (Security Sync)
When multiple agents act in parallel, the **Leader** agent must propagate security guardrails:

1.  **Safety Parity**: Synchronize `DANGEROUS_BASH_PATTERNS` across all sub-sessions.
2.  **Autonomy Sync**: If the Leader is in `auto-mode`, teammates assume the same level of autonomy but with restricted access to global configuration files.
3.  **Lifecycle Management**: Use heartbeats to verify teammate status (running/done/error).

---

## 📊 Layout & Workspace Management
Organize the workspace to prevent cognitive load during swarm operations:

- **Primary Pane**: Leader agent (Orchestrator).
- **Secondary Panes/Tabs**: Teammates, labeled by persona (e.g., `[Reviewer]`).
- **Standard Communication**: All inter-agent communication MUST happen via Markdown artifacts in the project repository (`STATE.md`, `tasks.md`, `handoff.md`).

---

> **Law of the Swarm**: Coordination is more important than speed. A synchronized swarm is unstoppable.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "SDD-CORE"
phase: "MANAGEMENT"
status: "COMPLETED"
last_update: "2026-05-06T09:42:30Z"
evidence_checksum: "NONE"
```
