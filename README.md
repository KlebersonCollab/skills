# 🧠 AI Agent Skills Hub (v4.7.0)

> Centralized hub for modular **Skills** and **Intelligence Governance**. Powered by the **HB-CLI**.

[![Skills](https://img.shields.io/badge/Skills-26-brightgreen)](#-available-skills)
[![License](https://img.shields.io/badge/License-MIT-blue)](LICENSE)
[![Build Status](https://github.com/KlebersonCollab/skills/actions/workflows/validate.yml/badge.svg)](https://github.com/KlebersonCollab/skills/actions)
[![Status](https://img.shields.io/badge/Status-Active-brightgreen)](#-roadmap)

---

## 📖 About the Project

This repository is the **source of truth** for modular AI agent skills. Every skill is an independent, versioned, and documented module that follows the **[SDD](sdd/)** (Spec-Driven Development) methodology.

Instead of manual downloads, we provide a high-performance CLI tool (**HB-CLI**) to manage your agent's intelligence with a single command.

---

## 🚀 HB-CLI: The Intelligence Manager

The **HB-CLI** (Harness Binary) is a standalone tool that makes the Hub **Plug & Play**. It handles installations, recursive dependencies, and agent rule distribution.

### 📥 Download the CLI
Download the binary for your OS and add it to your PATH:

| Platform | Download |
|----------|----------|
| **macOS (M1/M2/M3)** | [📥 hb-darwin-arm64](https://github.com/KlebersonCollab/skills/releases/latest/download/hb-darwin-arm64) |
| **macOS (Intel)** | [📥 hb-darwin-amd64](https://github.com/KlebersonCollab/skills/releases/latest/download/hb-darwin-amd64) |
| **Linux (x64)** | [📥 hb-linux-amd64](https://github.com/KlebersonCollab/skills/releases/latest/download/hb-linux-amd64) |
| **Windows (x64)** | [📥 hb-windows-amd64.exe](https://github.com/KlebersonCollab/skills/releases/latest/download/hb-windows-amd64.exe) |

> [!TIP]
> **macOS Users**: If you see a "malware" or "unverified developer" warning, run:
> `xattr -d com.apple.quarantine /path/to/hb-darwin-arm64`
> Or go to **System Settings > Privacy & Security** and click **"Allow Anyway"**.

---

## 🛠️ Quick Start (Plug & Play)

Once you have the `hb` binary, you can bootstrap your environment in seconds:

```bash
# 1. List available skills on GitHub
hb listskills

# 2. List locally registered skills
hb list

# 3. Install a skill (it will automatically install all dependencies!)
# Example: Architecture depends on SDD and Token-Distiller
hb install architecture --remote

# 4. Distribute rules to your agents (.gemini, .claude, .agent)
hb rules

# 5. Check ecosystem integrity
hb check
```

### 🧠 Intelligent Dependency Management
When you install a skill using `hb install <skill> --remote`, the CLI reads the metadata and **automatically downloads all dependencies**. You get a ready-to-use intelligence package with no manual work.

---

## 🧭 SDD Governance (Spec-Driven Development)

The HB-CLI replaces legacy scripts for managing the SDD lifecycle:

| Command | Purpose |
|---------|---------|
| `hb sdd start <name>` | Initializes a new feature with full SDD structure (6 artifacts). |
| `hb sdd status` | Shows real-time progress of all features in development. |
| `hb sdd bootstrap` | Initializes the project's **Triad of Memory** (State, Memory, Learnings). |
| `hb map` | Generates the **Knowledge Map** of the whole ecosystem (Mermaid). |
| `hb bench` | Runs **Performance Benchmarks** and compares with historical baseline. |

---

## 🛠️ Available Skills (26 Total)

| Skill | Purpose | Version |
|-------|---------|---------|
| **[SDD](sdd/)** | Spec-Driven Development Framework. | `1.5.0` |
| **[Architecture](architecture/)** | Systems Architect & ADR Management. | `2.0.1` |
| **[Context Graph](context-graph/)** | Meta-Knowledge & Decision Intelligence. | `1.0.0` |
| **[Harness Expert](harness-expert/)** | Agentic support for Sync & Rehydration. | `2.1.0` |
| **[Golang Expert](golang-expert/)** | High-performance Go development. | `1.2.0` |
| **[Token Distiller](token-distiller/)** | Token optimization & dual modes. | `1.1.0` |
| ... and 20 more in the [Full Catalog](onboarding-navigator/references/skills-catalog.md). | | |

---

## 📄 License
This project is licensed under the [MIT License](LICENSE).

<div align="center">
**Built with 🧠 by [Kleberson Romero](https://github.com/KlebersonCollab)**
</div>
