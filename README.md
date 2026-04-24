# 🧠 AI Agent Skills Hub

> Centralized hub for the development, storage, and evolution of modular **Skills** for AI agents.

[![Skills](https://img.shields.io/badge/Skills-25-brightgreen)](#-available-skills)
[![License](https://img.shields.io/badge/License-MIT-blue)](LICENSE)
[![GitHub Actions](https://img.shields.io/github/actions/workflow/status/KlebersonCollab/skills/generate-skills-artifacts.yml?branch=main&label=Build%20Artifacts)](https://github.com/KlebersonCollab/skills/actions)
[![Status](https://img.shields.io/badge/Status-Active-brightgreen)](#-roadmap)

---

## 📖 About the Project

This repository is the **source of truth** for all skills used by AI agents. Each skill is an independent, documented, and versioned module that can be integrated into any compatible agent.

We use the **[SDD](sdd/)** (Spec-Driven Development) methodology to ensure that each new functionality is rigorously specified, planned, and verified before implementation.

---

## 📦 Download & Quick Start

We automatically generate pre-configured packages for the main AI agents. Just download, extract to your project's root, and start using:

| Agent | Artifact | Content |
|--------|----------|----------|
| **Claude AI** | [📥 claude-skills.zip](https://github.com/KlebersonCollab/skills/releases/latest/download/claude-skills.zip) | `.claude/` folder with `CLAUDE.md` and all skills. |
| **Gemini CLI** | [📥 gemini-skills.zip](https://github.com/KlebersonCollab/skills/releases/latest/download/gemini-skills.zip) | `.gemini/` folder with `GEMINI.md` and all skills. |
| **AI Agent (Generic)** | [📥 agent-skills.zip](https://github.com/KlebersonCollab/skills/releases/latest/download/agent-skills.zip) | `.agent/` folder with `AGENT.md` and all skills. |

> 💡 *Note: The links above automatically download the latest version from the **Releases** tab.*

### 🚀 How to use the Skills
1. Download the ZIP corresponding to your agent.
2. Extract the content to your project's root directory.
3. Ensure the hidden folder (e.g., `.claude/`) was created correctly.
4. Your AI agent will automatically detect the new capabilities.

---

## 🛠️ Tools & Automation

The Hub includes the **SDD CLI**, a Python tool that automates the skill lifecycle and knowledge management:

| Command | Description |
|---------|-----------|
| `uv run sdd init <name>` | Initializes a new feature with full SDD structure. |
| `uv run sdd task <feat> <id>` | Marks task progress and synchronizes global state. |
| `uv run sdd graph` | Automatically generates the project **Knowledge Map** (Mermaid). |
| `uv run sdd sync` | Synchronizes global mandates across all agents (.gemini, .claude, .agent). |

> 💡 *Note: Requires [Python UV](https://docs.astral.sh/uv/) installed.*

---

## 🛠️ Available Skills

| # | Skill | Description | Version |
|---|-------|-----------|--------|
| 1 | **[SDD](sdd/)** | Spec-Driven Development — Modular and adaptive workflow with **Persistent Memory Protocol**. | `1.5.0` |
| 2 | **[Skill Factory](skill-factory/)** | Core Framework for standardized creation of new skills with scaffolding and validation. | `1.1.0` |
| 3 | **[Python with UV](python-uv/)** | Professional Python development with UV — 10-100x faster manager. | `2.6.0` |
| 4 | **[API Architect](api-architect/)** | API Architect — interoperable and secure design (OpenAPI, GraphQL, tRPC). | `1.3.0` |
| 5 | **[Brainstorming](brainstorming/)** | Design and problem-solving facilitator — deep exploration and validation. | `1.1.0` |
| 6 | **[Architecture](architecture/)** | Systems Architect — pragmatic design, trade-offs, and decision records (ADR). | `2.0.1` |
| 7 | **[Flutter with FVM](flutter-fvm/)** | Professional Flutter development with FVM. Includes Dart 3+, optimized performance, mandatory accessibility, and OWASP security. | `1.3.0` |
| 8 | **[Azure DevOps](azure-devops/)** | Professional management of Boards, Repos, Pipelines, and Artifacts in AzDO. | `1.1.0` |
| 9 | **[Clean Code Mentor](clean-code-mentor/)** | Technical mentor and code reviewer focused on SOLID, YAGNI, DRY, and KISS. | `1.0.0` |
| 10 | **[Observability Expert](observability-expert/)** | SRE Specialist — Structured Logging, OpenTelemetry, and Reliability (SLI/SLO). | `1.0.0` |
| 11 | **[Onboarding Navigator](onboarding-navigator/)** | Interactive culture and engineering guide for new members and projects. | `1.5.0` |
| 12 | **[YouTube Transcript](youtube-transcript/)** | Video transcript extraction with AI fallback (Whisper) and automatic cleanup. | `1.0.0` |
| 13 | **[Harness Expert](harness-expert/)** | Technical engine for Harness Engineering (Sync, Rehydrate, Automation). | `2.0.0` |
| 14 | **[Knowledge Architect](knowledge-architect/)** | Local knowledge architecture via relational graphs (Local GraphRAG). | `1.0.0` |
| 15 | **[FastAPI Expert](fastapi-expert/)** | API Excellence with FastAPI — performance, Annotated, and Pydantic V2. | `1.1.0` |
| 16 | **[Django Expert](django-expert/)** | Professional development with Django — Architecture, Security, and TDD. | `1.5.0` |
| 17 | **[Swarm Facilitator](swarm-facilitator/)** | Multi-Agent workflow orchestrator (Swarm). Defines handoff protocols. | `1.1.0` |
| 18 | **[Scaffolding Expert](scaffolding-expert/)** | Dynamic project generator. Uses CLI tools (uvx copier) to initialize boilerplates. | `1.0.0` |
| 19 | **[Golang Expert](golang-expert/)** | Go Excellence — Performance, idiomatic concurrency, and Samber ecosystem. | `1.2.0` |
| 20 | **[Frontend Expert](frontend-expert/)** | Expert in modern interfaces with React, Next.js, and TailwindCSS v4. | `1.1.0` |
| 21 | **[Git Workflow](git-workflow/)** | Git workflow patterns, branching strategies, and SDD integration. | `1.0.0` |
| 22 | **[Golang Testing Expert](golang-testing-expert/)** | QA Specialist for Go — TDD, Table-Driven Tests, Benchmarks, and Fuzzing. | `1.0.0` |
| 23 | **[Benchmark Expert](benchmark-expert/)** | Expert Skill for measuring performance baselines, detecting regressions, and comparing stacks. | `1.0.0` |
| 24 | **[DevSecOps Expert](devsecops-expert/)** | Static quality auditing, SAST/DAST security, and Hardening. | `1.0.0` |
| 25 | **[Token Distiller](token-distiller/)** | Token density manager and compression modes (Caveman/Premium). | `1.1.0` |

---

## 🏗️ Repository Structure

```
skills/
├── .github/workflows/           # CI/CD Automation (GitHub Actions)
├── .specs/                      # Project Specifications (SDD)
│   ├── project/                 # Vision, Roadmap, and State
│   └── codebase/                # Stack and Conventions
├── scripts/                     # Automation and distribution scripts
├── <skill>/                     # 🧩 Each skill in its directory
│   ├── README.md                # Detailed documentation
│   ├── SKILL.md                 # Main technical definition
│   └── CHANGELOG.md             # Version history
└── README.md                    # ← You are here
```

---

## 📐 How to Create a New Skill

Use the **[Skill Factory](skill-factory/)** skill to create new skills in a standardized way. This process ensures that each skill is generated with the correct structure and passes a quality audit before being integrated into the hub.

> 📖 See the [Skill Factory documentation](skill-factory/) for details.

---

## 🗺️ Roadmap & Project State

Follow the Hub's evolution through the planning documents:
- 🛤️ **[ROADMAP.md](.specs/project/ROADMAP.md)**: Long-term vision and goals.
- 📊 **[STATE.md](.specs/project/STATE.md)**: Current status and tasks in progress.

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).

---

<div align="center">

**Built with 🧠 by [Kleberson Romero](https://github.com/KlebersonCollab)**

*Precision at scale. Rigor when needed, speed when possible.*

</div>
