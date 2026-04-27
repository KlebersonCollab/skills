---
name: python-uv
version: 3.0.0
description: "EXPERT level skill for professional Python development with UV. Encompasses environment management, idiomatic patterns (Clean Code), and advanced testing with Pytest."
category: development-workflow
---

# Python with UV

> Extremely fast Python package and project manager, written in Rust. Replaces pip, pipx, poetry, pyenv, and more — with 10-100x superior performance.

---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Have you rehydrated the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define the architecture, schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---
## Goal

Empower the agent to develop professional-level Python applications, using **UV** for extreme performance and applying excellence standards in coding (Idioms/Patterns) and technical rigor in testing (TDD/Pytest). The skill unifies package management, idiomatic patterns, and software quality into a single, deterministic workflow.

---

## Version Awareness

**Recommended Version: UV 0.9.7+**

Before starting, check the installed version:

```bash
uv --version
```

**Changes by version:**

| Version | Change |
|--------|---------|
| **0.9.6+** | Python 3.14 is the default (previously 3.13) |
| **0.9.6+** | Free-threaded Python 3.14+ without explicit opt-in |
| **0.9.6+** | `uv build --clear` to clear build artifacts |
| **0.9.7+** | Security updates for tar/ZIP handling |

If the version is earlier than 0.9.0, update:

```bash
# macOS / Linux
curl -LsSf https://astral.sh/uv/install.sh | sh

# macOS via Homebrew
brew install uv

# Windows
powershell -ExecutionPolicy ByPass -c "irm https://astral.sh/uv/install.ps1 | iex"
```

See [Recent Changes](references/python-environment.md) for migration details.

## When to Use This Skill

Use this skill when:
- Creating or initializing a new Python project
- Managing dependencies with `pyproject.toml`
- Installing or managing Python versions
- Configuring virtual environments
- Installing development CLI tools (black, ruff, mypy, pytest)
- Executing Python scripts with inline dependencies (PEP 723)
- Configuring CI/CD with UV (GitHub Actions)
- Deciding between `uv tool install` vs `uvx`
- Migrating from pip, pipx, or poetry to UV
- Resolving errors related to UV

Skip this skill when:
- The project uses Python 2.x (UV does not support it)
- The project does NOT use Python as a language
- You only need basic pip documentation (without UV)

---

## Workflow (4 Phases)

### Phase 1: ENVIRONMENT — Configuration and Alignment
1.  **Check Version**: Run `uv --version` to ensure compatibility (Recommended 0.9.7+).
2.  **Define Python Version**: If `.python-version` does not exist, use `uv python pin 3.12` (or desired version).
3.  **Isolated Environment**: Ensure `.venv/` exists or will be created via `uv venv`.

### Phase 2: PROJECT — Initialization and Structure
1.  **Scaffolding**: For new projects, use `uv init`. For migrations, use `uv init --bare`.
2.  **Declare Dependencies**: Add packages via `uv add` (production) and `uv add --dev` (tools).
3.  **Global Sync**: Run `uv sync` to create `uv.lock` and synchronize the virtual environment.

### Phase 3: DEVELOP — Quality and Execution
1.  **Configure Tooling**: Define rules in `pyproject.toml` for Ruff and Mypy.
2.  **Safe Execution**: Always use `uv run <command>` to ensure code runs in the correct environment.
3.  **Incremental Quality**: Run `uv run ruff check .` and `uv run pytest` periodically.

### Phase 4: DEPLOY — Build and Distribution
1.  **Reproducible Lock**: Ensure `uv.lock` is updated and under version control.
2.  **Compilation**: Use `uv build --clear` to generate clean distributions (sdist/wheel).
3.  **Publication**: Run `uv publish` to distribute to PyPI or a private repository.

---

## UV Commands Overview

| Command | Purpose | Example |
|---------|-----------|---------|
| `uv init` | Create new project | `uv init my-project` |
| `uv add` | Add dependency | `uv add requests pydantic` |
| `uv remove` | Remove dependency | `uv remove requests` |
| `uv sync` | Synchronize environment | `uv sync --dev` |
| `uv lock` | Generate/update lockfile | `uv lock` |
| `uv run` | Execute in virtual environment | `uv run pytest` |
| `uv pip install` | Install packages (pip mode) | `uv pip install requests` |
| `uv tool install` | Install isolated global CLI tool | `uv tool install black` |
| `uvx` | Execute package in temporary environment | `uvx ruff check .` |
| `uv venv` | Create virtual environment | `uv venv .venv` |
| `uv python install` | Install Python version | `uv python install 3.12` |
| `uv python pin` | Pin Python version in the project | `uv python pin 3.12` |
| `uv build` | Build distribution | `uv build --clear` |
| `uv publish` | Publish to PyPI | `uv publish` |

---

## Tool vs UVX Decision Tree

```text
Need to execute a Python package?
│
├─ Use it daily/frequently?
│  └─ YES → `uv tool install package`
│     Examples: black, ruff, mypy, pytest
│
├─ Temporary execution / testing?
│  └─ YES → `uvx package`
│     Examples: test a new tool, compare versions
│
├─ MCP server?
│  └─ YES → `uvx package` or `uvx --from path script.py`
│     Examples: mcp-server-sqlite, custom MCP servers
│
└─ Project local script?
   └─ YES → `uvx --from . script.py`
      Examples: project-specific scripts
```

---

## Project Initialization

```bash
# Create new project
uv init my-project
cd my-project

# Generated structure:
# my-project/
# ├── .python-version
# ├── pyproject.toml
# ├── README.md
# └── src/
#     └── my_project/
#         └── __init__.py

# Add dependencies
uv add requests pydantic

# Add development dependencies
uv add --dev pytest ruff mypy

# Execute
uv run python -m my_project
```

### Dependency Management

```bash
# Add production dependency
uv add requests pydantic

# Add development dependency
uv add --dev pytest ruff mypy

# Add optional dependency
uv add --optional dev pytest-cov

# Remove dependency
uv remove requests

# Synchronize environment (installs everything from pyproject.toml)
uv sync --dev

# Generate lockfile
uv lock

# Synchronize from lockfile (exact reproduction)
uv sync --locked
```

---

## Virtual Environment Management

```bash
# UV automatically creates when using uv sync/run
# But if you need to create manually:
uv venv .venv

# Activate (macOS / Linux)
source .venv/bin/activate

# Activate (Windows Git Bash)
. .venv/Scripts/activate

# Activate (Windows CMD)
.venv\Scripts\activate.bat

# Execute WITHOUT activating (recommended)
uv run python script.py
uv run pytest
```

---

## Inline Script Metadata (PEP 723)

Self-contained Python scripts with dependencies defined in comments:

```python
# /// script
# dependencies = [
#   "requests>=2.28.0",
#   "rich>=13.0.0",
# ]
# requires-python = ">=3.11"
# ///

import requests
from rich.console import Console

console = Console()
response = requests.get("https://api.github.com")
console.print(response.json())
```

Execute with automatic dependency installation:

```bash
uv run script.py
```

**When to use**: Single scripts, utilities, automation, sharable examples.
**When NOT to use**: Multi-file projects (use `pyproject.toml`).

See [Inline Script Metadata Reference](references/inline-script-metadata.md) for full examples.

---

## Development Tools Setup

```bash
# Install development tools once
uv tool install ruff
uv tool install mypy
uv tool install pytest

# Use daily (available globally)
ruff format .
ruff check .
mypy .
pytest tests/
```

### Recommended configuration in `pyproject.toml`

```toml
[tool.ruff]
line-length = 88
target-version = "py311"

[tool.ruff.format]
quote-style = "double"

[tool.mypy]
python_version = "3.11"
strict = true
warn_return_any = true
warn_unused_configs = true

[tool.pytest.ini_options]
testpaths = ["tests"]
python_files = ["test_*.py"]
```

---

## Best Practices

### Project Management

**DO:**
- Use `uv init` for new projects (creates full structure)
- Use `uv add` / `uv remove` to manage dependencies (do not edit pyproject.toml manually)
- Always version `uv.lock` in git for exact reproduction
- Use `uv sync --locked` in CI/CD for reproducible builds
- Use `uv run` to execute commands (no need to activate venv)

**DON'T:**
- Do not install packages globally without venv
- Do not mix `pip install` and `uv add` in the same project
- Do not ignore `uv.lock` in `.gitignore`
- Do not use `@latest` in production (unstable)

### Tool Management

**DO:**
- Use `uv tool install` for daily use tools (ruff, mypy, pytest)
- Use `uvx` for temporary executions and tool testing
- Keep tools updated with `uv tool upgrade --all`

**DON'T:**
- Do not use global `pip install` for CLI tools (causes conflicts)
- Do not use `uvx` for daily tools (unnecessary overhead)
- Do not mix pip and uv tool installations

### Code Quality

**DO:**
- Always use type hints in Python code
- Configure `ruff` for linting and formatting
- Configure `mypy` in strict mode
- Write tests with `pytest`

**DON'T:**
- Do not promote code without type hints
- Do not skip linting, formatting, or type checking
- Do not use deprecated practices like `pip` without venv

---

## Troubleshooting

### "UV not found" / "spawn uvx ENOENT"
**Cause**: UV/UVX is not in the PATH.
**Solution**:
```bash
# macOS/Linux — add to ~/.zshrc or ~/.bashrc
export PATH="$HOME/.local/bin:$PATH"
source ~/.zshrc

# Verify
uv --version
```

### "Package not found"
**Cause**: Incorrect package name or package not published on PyPI.
**Solution**:
```bash
# Check on PyPI
curl https://pypi.org/pypi/package-name/json | head -5

# For local projects, check pyproject.toml
ls pyproject.toml
```

### "No virtual environment found"
**Cause**: Virtual environment not activated or non-existent.
**Solution**:
```bash
# Create and synchronize
uv sync

# Or execute directly (without activating)
uv run python script.py
```

### Version conflicts
**Cause**: Incompatible dependency versions.
**Solution**:
```bash
# Check dependency tree
uv tree

# Check reverse conflicts
uv tree --reverse

# Clean cache and resync
uv cache clean
uv sync --recreate
```

---

## Migration Guides

### From pip
```bash
# Before                          # After
pip install requests      →      uv add requests
pip install -r req.txt    →      uv pip install -r req.txt
pip freeze > req.txt      →      uv lock  (uses uv.lock)
```

### From poetry
```bash
# Before                          # After
poetry add requests       →      uv add requests
poetry install            →      uv sync
poetry run pytest         →      uv run pytest
poetry build              →      uv build
```

### From pipx
```bash
# Before                          # After
pipx install black        →      uv tool install black
pipx run black .          →      uvx black .
```

---

## Reference Documentation

This skill includes detailed reference documentation:

1. **[Installation & Setup](references/installation-and-setup.md)** — Cross-platform installation, virtual environments, troubleshooting
2. **[Project Management](references/project-management.md)** — Init, dependencies, lockfile, workspaces, project structure
3. **[Tool Management](references/tool-management.md)** — `uv tool install` vs `uvx`, decision matrix, maintenance
4. **[Python Patterns](references/patterns.md)** — [NEW] Idiomatic patterns, performance, EAFP, Modern Type Hints, __slots__
5. **[Advanced Testing](references/testing.md)** — [NEW] TDD, Pytest Mastery, Fixtures, Mocking, Coverage
6. **[Inline Script Metadata](references/inline-script-metadata.md)** — PEP 723, self-contained scripts, practical examples
7. **[Python Environment](references/python-environment.md)** — Versions, Python 3.14, free-threaded Python, paths
8. **[CI/CD Workflows](references/ci-cd-workflows.md)** — GitHub Actions, Docker, packaging, publication

---

## Output Structure

Execution of this skill in Python projects must result in the creation or update of the following UV-standardized artifacts:

| Artifact | File | Description |
|----------|---------|-----------|
| **Project Spec** | `pyproject.toml` | Metadata definition, dependencies, and tool settings (Ruff, Mypy). |
| **Lockfile** | `uv.lock` | Deterministic record of all dependencies and sub-dependencies (must be versioned). |
| **Python Version** | `.python-version` | File that pins the Python version used in the project for UV. |
| **Virtual Env** | `.venv/` | Isolated virtual environment containing installed dependencies (must not be versioned). |
| **Source Code** | `src/` or root | Python package structure following layout best practices. |

---

## External Resources

- **UV Official Docs**: <https://docs.astral.sh/uv/>
- **UV GitHub**: <https://github.com/astral-sh/uv>
- **Ruff (linter/formatter)**: <https://docs.astral.sh/ruff/>
- **PEP 723 (Inline Metadata)**: <https://peps.python.org/pep-0723/>
- **Python Packaging Guide**: <https://packaging.python.org/>

---

## Quality Rules

- **Official Reference**: Always reference the official UV documentation (<https://docs.astral.sh/uv/>)
- **Practical Examples**: All examples must be testable and reproducible
- **Cross-Platform**: Commands and configurations must work on macOS, Linux, and Windows
- **Modern Best Practices**: Promote type hints, ruff, pytest, mypy
- **Semantic Versioning**: Follow SemVer for project configurations
- **Lockfiles**: Always use and version `uv.lock`
- **UV-First**: Prefer native UV commands (`uv add`, `uv run`) over legacy equivalents (`pip install`)

## Prohibited

- **Do not assume OS**: Do not assume specific OS configurations without offering alternatives
- **Do not duplicate documentation**: Do not copy content from official documentation, only reference it
- **Do not promote obsolete practices**: Do not teach `pip install` without venv, `setup.py`, or `requirements.txt` as primary management
- **Do not ignore type safety**: Do not promote Python code without type hints
- **Do not skip validation**: Do not skip linting, formatting, or type checking
- **Do not mix managers**: Do not mix pip and uv in the same project workflow
