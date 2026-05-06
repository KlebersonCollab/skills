# SDD — Spec-Driven Development

> Precision at scale. Rigor when needed, speed when possible.

[![Version](https://img.shields.io/badge/Version-2.3.0-blue)](#changelog)
[![Sub-skills](https://img.shields.io/badge/Sub--skills-5-brightgreen)](#-sub-skills)
- [Handoff Protocol](references/handoff-protocol.md)

---

## 📖 Overview

A skill for **spec-driven development** with a modular and adaptive workflow. SDD automatically adjusts the depth of the process to the complexity of the task (Auto-Sizing), ensuring rigor when necessary and speed when possible.

Version `2.3.0` integrates **Observable Governance**, mandatory metadata blocks, and evidence-based auditing via structured task tables.

> **Law of SDD**: If it's not in the spec, it doesn't exist. If it hasn't been verified, it's not done.

---

## ⚙️ Auto-Sizing

The depth of the workflow is determined by the **complexity** of the task:

| Level | Scope | Phases | Sub-skills Used |
|-------|--------|-------|-----------------------|
| **Quick** | Bug fixes, configs, <3 files | Implement → Verify | `implementer` |
| **Small** | Clear feature, <5 tasks | Spec → Impl → Verify | `orchestrator`, `implementer` |
| **Medium** | Feature + UI, <10 tasks | Discovery → Spec → Impl → Verify | `explorer`, `orchestrator` |
| **Large** | Multi-component, new module | Discovery → Spec → Impl → Verify | All |
| **Complex** | Ambiguity, high risk | All phases + Interactive UAT | All + `reviewer` |

---

## 🔄 4-Phase Workflow

The SDD follows a rigorous cycle to ensure integrity and traceability:

1.  **DISCOVERY**: Understand the terrain and align with the project vision. (Uses `sdd-explorer`, `sdd-planner`)
2.  **SPECIFY**: Define requirements (BDD), design architecture, and record decisions. (Uses `sdd-orchestrator`, `sdd-planner`)
3.  **IMPLEMENT**: Technical execution with continuous progress tracking. (Uses `sdd-implementer`)
4.  **VERIFY**: Final audit against acceptance criteria and capture learnings. (Uses `sdd-reviewer`, `sdd-planner`)

---

## 🧩 Sub-skills

| Sub-skill | File | Responsibility |
|-----------|---------|------------------|
| **Explorer** | [sdd-explorer.skill.md](sdd-explorer.skill.md) | Maps codebase into `STACK.md`, `ARCHITECTURE.md`, `CONVENTIONS.md`, and `CONCERNS.md`. |
| **Planner** | [sdd-planner.skill.md](sdd-planner.skill.md) | Manages project vision and **Persistent Memory** (`STATE.md`, `MEMORY.md`, `LEARNINGS.md`, `DECISIONS.md`). |
| **Orchestrator** | [sdd-orchestrator.skill.md](sdd-orchestrator.skill.md) | Creates specifications (`spec.md`, `plan.md`, `tasks.md`, `contract.md`) with full traceability. |
| **Implementer** | [sdd-implementer.skill.md](sdd-implementer.skill.md) | Executes atomic, test-driven code following project conventions and safety-valves. |
| **Reviewer** | [sdd-reviewer.skill.md](sdd-reviewer.skill.md) | Audits implementations against ACs with evidence-based reporting and UAT. |

---

## 🔗 Knowledge Verification Chain

When researching or deciding, follow this strict hierarchy:

1. **Codebase Docs** → Consult `.specs/codebase/`.
2. **Project Docs** → Check READMEs, `PROJECT.md`, `ROADMAP.md`.
3. **Existing Code** → Read the source code directly.
4. **Web Search** → Official documentation only.
5. **Flag Uncertainty** → If steps 1-4 fail, communicate uncertainty to the user.

---

## 📁 Directory Structure (Mandatory)

### Per Feature (`spec/` or `.specs/features/[name]/`)
- `spec.md` — Traceable requirements (FR-X) and ACs (BDD).
- `plan.md` — Technical design and schemas.
- `tasks.md` — Atomic task list with status.
- `contract.md` — Delivery agreement and validation sensors.

### Per Project (`.specs/project/`)
- `PROJECT.md`, `ROADMAP.md`, `STATE.md`, `MEMORY.md`, `LEARNINGS.md`, `DECISIONS.md`.

### Per Codebase (`.specs/codebase/`)
- `STACK.md`, `ARCHITECTURE.md`, `CONVENTIONS.md`, `CONCERNS.md`, `TECHNICAL-MAP.md`.

---

## 📋 Operational Protocols

### Safety Valve
If a task reveals hidden complexity (>3 files touched or structural design changes needed), **STOP** and re-plan.

### Persistence Protocol
Always update `STATE.md`, `MEMORY.md`, and `LEARNINGS.md` at the end of every session or after major decisions.

---

## 📝 Changelog

See [CHANGELOG.md](CHANGELOG.md) for full version history.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T09:41:00Z"
evidence_checksum: "NONE"
```

