# Context Graph (Meta-Knowledge Manager)

[![Version](https://img.shields.io/badge/Version-1.0.0-blue)](#)
[![Category](https://img.shields.io/badge/Category-Intelligence--Governance-green)](#)

> "The Memory Graph Layer for Agents."

## Overview

The `context-graph` skill implements a **Context Graph** (also known as a Meta-Knowledge Graph) to serve as the long-term memory for AI agents. While traditional knowledge graphs map domain entities (the "what"), the Context Graph maps the intelligence process (the "how" and "why").

By capturing **Decision Traces**, **Patterns**, and **Tribal Knowledge**, this skill ensures that:
1. Every lesson learned is persisted.
2. Architectural rationale is traceable.
3. Successes are promoted to standard procedures.
4. Failures become searchable precedents to avoid repetition.

## Features

- **Episodic Memory**: Detailed logging of decisions, alternatives, and rationale.
- **Semantic Memory**: Distillation of domain-specific patterns and "soft" constraints.
- **Procedural Memory**: Proven strategies for complex tasks (e.g., schema design, entity resolution).
- **Outcome Tracking**: Closing the loop between decisions and their real-world impact.

## Installation

This skill is part of the **AI Agent Skills Hub**. To use it, ensure you have the `hb` CLI installed:

```bash
hb install context-graph
hb rules
```

## Usage

Enable the skill by referencing it in your agent's instructions or using the following trigger:

> "Use the `context-graph` skill to document the rationale for the following architectural change and check for any existing precedents."

## References

- Article: [Meta Knowledge Graphs (Context Graphs)](https://firattekiner.substack.com/p/meta-knowledge-graphscontext-graphs) by Firat Tekiner.
- Research: Hu et al. (2025), "Memory in the Age of AI Agents".
