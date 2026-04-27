# Contract: Enrichment of Swarm Facilitator

## Agreement
The Implementer will enrich the `swarm-facilitator` skill with technical governance patterns, ensuring that multi-agent sessions are secure, synchronized, and resilient across different terminal backends.

## Sensors (Validation)
1.  **Backend Mention**: `SKILL.md` must mention iTerm2, Tmux, and In-Process runners.
2.  **Permission Sync**: `SKILL.md` must describe the "Permission Bridge" protocol.
3.  **Lifecycle**: `SKILL.md` must include a section on Teammate Heartbeat/Reconnection.
4.  **Hub Audit**: `hb audit swarm-facilitator` must pass.
