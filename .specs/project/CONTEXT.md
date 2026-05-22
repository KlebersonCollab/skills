# Context - Skills Harness

This context defines the precise domain model and terminology used to transform our repository into an active agentic execution harness.

## Language

**Harness**:
A lightweight, terminal-based local runtime that orchestrates Language Models to interact with the local filesystem and terminal interface.
_Avoid_: platform, framework, wrapper

**Session Tree (Árvore de Sessão)**:
A directed acyclic graph (DAG) of reasoning elements, which records chat history and allows branching from any historical point.
_Avoid_: chat log, history list

**Node (Nó)**:
A single element in the Session Tree, representing a message (system, user, assistant) and any actions or tool executions.
_Avoid_: message, step, entry

**Active Skill (Skill Ativa)**:
A module folder within the Hub containing guidelines and local scripts that the Harness dynamically loads to direct the agent's reasoning.
_Avoid_: rule file, document, template

**Minimal Toolset (Conjunto Mínimo de Ferramentas)**:
The set of four pristine capabilities (read, write, patch, bash) exposed by the Harness to manipulate codebases.
_Avoid_: operations, APIs

**Provider Template (Modelo de Provedor)**:
A configuration structure defining the API endpoints, headers, and request body formats for dynamic LLM integration.
_Avoid_: model SDK, client wrapper

---

## Term Interaction Example

> **Developer**: Type `harness run "fix alignment"` to boot the **Harness**.
> **Harness**: The **Harness** loads the **Active Skill** `clean-code-mentor` to enforce styling, creating a **Session Tree** with a root **Node**.
> **Harness**: Using its **Minimal Toolset**, the **Harness** executes `patch` to modify the CSS and `bash` to verify the build, adding subsequent **Nodes** to the **Session Tree**.
> **Developer**: The developer selects a previous **Node** and types `/branch` to try another design, creating a new path in the **Session Tree**.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HARNESS-TRANSFORMATION"
phase: "ALIGN"
status: "COMPLETED"
last_update: "2026-05-22T14:41:00Z"
evidence_checksum: "ffb546e"
```
