# Knowledge Architect

> Local knowledge architecture via relational graphs (Local GraphRAG). Transforms static documentation into a connection map for intelligent context navigation.

[![Version](https://img.shields.io/badge/Version-1.0.0-blue)](#changelog)
[![Sub-skills](https://img.shields.io/badge/Sub--skills-0-brightgreen)](#)

---

## 📖 Overview

The **Knowledge Architect** is the intelligence behind the Hub's semantic organization. While other skills focus on *doing* (implementing, testing, orchestrating), this skill focuses on *understanding* the structural relationships of the project.

Inspired by **GraphRAG** concepts, it allows the agent to build a "second brain" for the repository, mapping how business requirements transform into technical decisions and, finally, into code and tests. This resolves the "tunnel vision" problem where the agent understands an isolated file but forgets the systemic impact of its changes.

---

## ⚙️ How to Use

This skill should be invoked in scenarios of:
1.  **Codebase Onboarding**: To quickly map the dependencies of a legacy project.
2.  **Impact Analysis**: Before large refactorings, to see who will be affected.
3.  **Feature Discovery (SDD)**: To connect new requirements to the existing knowledge graph.
4.  **Relational Rehydration (Harness)**: To load "neighboring" context intelligently.

**Invocation Example:**
*"Use the knowledge-architect skill to map the relationships between the new authentication feature and the existing database modules, generating a KNOWLEDGE-MAP.mermaid."*

---

## 📝 Changelog

See [CHANGELOG.md](CHANGELOG.md) for full version history.
