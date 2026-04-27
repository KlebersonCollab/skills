---
name: cli-expert
version: 1.0.0
description: "Expert in designing and implementing AI-native Command Line Interfaces (CLIs). Focuses on Slash Commands, TUI components (React/Ink), and tool-registration standards for high-maturity agentic systems."
category: automation-ux
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Did you rehydrate the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements for the CLI tool?
3. **Plan Check**: Does the `plan.md` include command schemas and UI flow diagrams?

---

# CLI Expert (AI-Native Interface Design)

> "A CLI for an agent is not just a set of flags; it is a structured communication protocol between model and machine." — Standard for building high-maturity agentic tools.

---

## 1. Goal
The goal of this skill is to provide a deterministic framework for building CLIs that are optimized for both human developers and LLM agents. It leverages modern paradigms like **Slash Commands** and **Stateful TUIs** to provide a premium user experience.

---

## 2. Slash Command Architecture
Commands should be categorized by their execution kind to manage complexity and security:

- **PromptCommand**: A command that is primarily a pre-configured prompt (e.g., `/explain`).
- **LocalCommand**: A command that executes local code/scripts (e.g., `/audit`).
- **BundledCommand**: Built-in commands that ship with the core binary.
- **MCPCommand**: Dynamic tools provided via the Model Context Protocol.

### 🚩 Command Metadata Standard
Every command must be registered with:
- `id`: Unique identifier (e.g., `ultrareview`).
- `description`: LLM-optimized technical description.
- `userDescription`: Human-friendly explanation.
- `prompt`: (Optional) The core prompt template if it's a PromptCommand.
- `parameters`: Strictly typed JSON schema for arguments.

---

## 3. TUI Design (Terminal User Interface)
Avoid "Streaming Log Fatigue". Use **Stateful UI** components:

### ⚛️ React/Ink Paradigm
- Use **Ink** (or equivalent) to build layouts with columns, spinners, and progress bars.
- Ensure that output is "compressed" when finished to avoid polluting the terminal history.
- **Micro-animations**: Use subtle spinners and color changes to indicate state (Processing, Success, Error).

### 📏 Layout Standards
- **Header**: Command name + Version.
- **Body**: Active progress or data table.
- **Footer**: Status line or available next steps.

---

## 4. Context-Aware Visibility
Commands should only be suggested when relevant:
- **Availability Guards**: Use logic to hide commands (e.g., don't show `/git-fix` if not in a git repo).
- **History Pointers**: Use `/thinkback` or `/history` patterns to help the LLM navigate its own past actions.

---

## 5. Workflow (4 Phases)

### Phase 1: SCHEMA — Command Definition
1. Define the input parameters using JSON Schema.
2. Draft the `description` to be unambiguous for the model.

### Phase 2: LOGIC — Implementation
1. Implement the core functionality (Local or Prompt).
2. Add guardrails (see `devsecops-expert`) for any command that mutates state.

### Phase 3: UX — TUI Implementation
1. Build the terminal component using Ink.
2. Test the "Streaming" experience to ensure it doesn't break the agent's context window.

### Phase 4: REGISTRY — Deployment
1. Register the command in the central registry.
2. Update help documentation and `hb map`.

---

## 6. Prohibited
- NEVER create commands with ambiguous names (e.g., `/do-it`).
- NEVER output raw, unformatted JSON to the user; always use a TUI formatter.
- NEVER ignore command availability; a cluttered help menu confuses the model.
- NEVER skip security audits for `LocalCommands`.

---

## 7. References
- [Ink - React for CLIs](https://github.com/vadimdemedes/ink)
- [Claude Code Command Pattern](https://github.com/anthropics/claude-code)
