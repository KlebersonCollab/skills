---
name: swarm-facilitator
version: 1.2.0
description: "Orchestrator for Multi-Agent (Swarm) workflows. Defines handoff protocols, personas, and context passing using the SDD structure."
category: orchestration
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Have you rehydrated the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define architecture, schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---
# Swarm Facilitator

> "Conducting multiple agents requires choreography, not just programming. The repository is the board, and Markdown is the inter-agent communication language."

---

## Goal

Enable AI agents to act in a choreographed flow (Swarm/Crew), assuming specialized personas (e.g., Architect, Developer, Reviewer) and utilizing Handoff protocols based on the SDD ecosystem. This skill transforms interaction into an invisible production line pipeline.

---

## Supported Personas

The orchestrator can instruct the current LLM or direct the user to start a new thread with one of these personas:

1. **Solution Architect**: Focuses on Phases 1 and 2 of SDD. Uses `architecture`, `api-architect`, and `brainstorming` skills. Produces PRDs and ADRs.
2. **Lead Developer (Engineer)**: Focuses on Phase 3 of SDD. Consumes the Architect's specifications. Uses stack-specific skills (`fastapi-expert`, `flutter-fvm`). Writes the production code.
3. **Quality Assurance / SRE (Reviewer)**: Focuses on Phase 4 of SDD. Uses `clean-code-mentor` and `observability-expert`. Performs code review and audits telemetry.

---

## Workflow (Handoff Protocol)

To orchestrate multiple agents in the same local project without context loss, execute this protocol:

### 1. State Definition (State is King)
The Architect must update `STATE.md` with the macro action plan and save tasks in `tasks.md`.

### 2. Context Dump (Save Game)
Before finishing the session of Agent 1 (Architect), it **must** compile recent decisions into a handoff file in the `.specs/features/<feature>/handoff.md` folder.
*Example handoff format:*
- What was decided.
- Where the specification is (absolute path).
- Which skill the next agent MUST invoke upon waking up.

### 3. Handoff Message (Passing the Torch)
The last message of Agent 1's session must be a formatted prompt for the user to copy and paste into Agent 2's new session.
**Template**: 
> "Start development for Feature X. I am the Architect and I left the specs in `.specs/...`. Invoke the `fastapi-expert` skill and follow steps 1 and 2."

---

## Output Structure

- `handoff.md`: Document containing guidelines for the next persona.
- `STATE.md`: Updated via Harness Expert indicating who has the ball.

---

## Quality Rules

- The Handoff must be lean.
- Communication is done exclusively through Markdown saved in the repository.

---

## Prohibited

- NEVER attempt to have a single agent assume the role of Architect and Developer simultaneously in complex epics. Divide and conquer.
- NEVER pass the torch without updating `STATE.md` via `harness-expert` or script.

## 5. Technical Governance & Synchronization (Teammates)

When orchestrating multiple agents in parallel (via terminal backends), the Leader agent must ensure technical parity and security across all sessions.

### 🚩 Terminal Backends
- **iTerm2 (Visual)**: Use for local debugging where the user needs to see the teammate's terminal state. Requires `It2SetupPrompt`.
- **Tmux (Headless)**: Use for long-running, multi-agent background tasks. Enables persistent sessions that survive disconnections.
- **In-Process (Ephemeral)**: Use for lightweight sub-tasks (e.g., file reading, simple linting) that don't require a full terminal environment.

### 🛡️ Permission Bridge (Security Sync)
The Leader **MUST** synchronize its security guardrails with all teammates:
1. **Blocked Patterns**: Propagate the `DANGEROUS_BASH_PATTERNS` list (see `devsecops-expert`) to teammates via the initialization prompt.
2. **Auto-Mode Parity**: If the Leader is in `auto-mode`, teammates should assume the same level of autonomy but with restricted write access to global config files.

### 🔄 Teammate Lifecycle
1. **Initialization**: Use `teammateInit.ts` patterns to define the sub-agent's restricted identity.
2. **Heartbeat**: The Leader should periodically query the teammate's status (running/done/error).
3. **Reconnection**: Use the `SessionID` to re-attach to a Tmux session if the connection is lost.

### 📊 Layout Management
Organize the workspace visually:
- **Leader**: Main pane.
- **Teammates**: Split panes or dedicated tabs, labeled by persona (e.g., `[Teammate-Reviewer]`).
