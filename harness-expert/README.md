# Harness Expert

> "If you are not the model, then you are the harness." — Support infrastructure to transform LLMs into operational agents.

[![Version](https://img.shields.io/badge/Version-1.1.0-blue)](#changelog)
[![Sub-skills](https://img.shields.io/badge/Sub--skills-3-brightgreen)](#-sub-skills)

---

## 📖 Overview

**Harness Expert** is the skill that orchestrates the infrastructure necessary for a language model (LLM) to behave as a reliable, secure, and persistent agent. Inspired by "Harness Engineering" concepts, it focuses on ensuring the agent has access to the right tools, maintains its state between sessions, and rigorously validates its own outputs.

In this Hub's ecosystem, the Harness Expert acts as the "caretaker" of the SDD (Spec-Driven Development) workflow, ensuring that the project's long-term memory (specs, logs, and states) is always synchronized with the agent's actions.

### 🧠 Harness Engineering Principles

1.  **Feed Forward (Preventive)**: Using Specs, Plans, and Contracts to guide the agent before action.
2.  **Sensors & Feedback (Corrective)**: Mandatory use of tools (linters, tests) to validate work. The agent is not the judge; the sensors are.
3.  **Bootstrapping**: Ensuring the operational environment (dependencies, scripts) is always ready for the agent to operate with precision.
4.  **Contract-based Review**: Every delivery is evaluated against a pre-established contract with a compliance Score.

---

## ⚙️ How to Use

This skill is automatically activated whenever an implementation task begins, or it can be manually invoked to perform state audits or progress synchronization.

**Application Scenarios:**
1.  **Session Start**: Rehydrate context by reading specification files.
2.  **During Implementation**: Synchronize completed sub-tasks in `tasks.md`.
3.  **Finalization**: Generate validation reports and persist lessons learned.

---

## 🧩 Sub-skills

| Sub-skill | File | Responsibility |
|-----------|---------|------------------|
| **Harness Sync** | [harness-expert-sync.skill.md](harness-expert-sync.skill.md) | Synchronizes task progress and records learnings in the File System. |
| **Harness Rehydrate** | [harness-expert-rehydrate.skill.md](harness-expert-rehydrate.skill.md) | Rebuilds the operational context by reading persisted specs and states. |
| **Harness Compress** | [harness-expert-compress.skill.md](harness-expert-compress.skill.md) | Compresses completed task history in STATE.md for token efficiency. |

---

## 📝 Changelog

Consult [CHANGELOG.md](CHANGELOG.md) for full version history.
