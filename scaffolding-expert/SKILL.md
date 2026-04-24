---
name: scaffolding-expert
version: 1.0.0
description: "Dynamic project generator. Uses CLI tools (uvx copier) to initialize boilerplates and complex architectures instantly."
category: automation
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the guidelines of the `token-distiller` skill.
1. **Context Check**: Have you rehydrated the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (Mandatory BDD for Medium+).
3. **Plan Check**: Does the `plan.md` file define the architecture and schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---
# Scaffolding Expert

> "Do not write boilerplate by reading text; generate the architecture by executing templates."

---

## Goal

Replace the slow, costly, and error-prone process of manually coding boilerplate files (such as `pyproject.toml`, Dockerfiles, base `Makefile`, default routers) with instant generation through executable template engines, specifically via **Copier**.

---

## How It Works

The primary tool chosen for the stack is `copier`, executed zero-install via UV.

### Standard Scaffold Command:
```bash
uvx copier copy gh:KlebersonCollab/skills/scaffolding/<template-name> <destination>
```
*Note: If there is a local directory with templates in the Hub, use the local path instead of the GitHub URL.*

---

## Workflow for New Features/Projects

1. **Verify**: Identify which stack the user wants (e.g., Python/FastAPI, Python/Django).
2. **Execute**: Before writing any `.py` code file manually, run the scaffolding command in the terminal.
3. **Configure**: `copier` (if configured with non-interactive flags or answered via STDIN) will ask for the project name. If necessary, run with explicit flags or create a `.copier-answers.yml` file.
4. **Build**: After generating the directory, navigate to it and run `uv sync` to validate the installation.

---

## Output Structure

- Generated directory containing all boilerplate for the requested stack (e.g., `pyproject.toml`, Dockerfile).

---

## Quality Rules

- **Absolute Priority**: If a scaffolding template exists, it MUST be used instead of writing manual code.
- **Customization**: You can (and should) modify the code *after* scaffolding, but the root structure always comes from the template.

---

## Prohibited

- NEVER write `pyproject.toml` by hand if a template is available.
- NEVER install scaffolding frameworks globally; always use ephemeral tools like `uvx`.
