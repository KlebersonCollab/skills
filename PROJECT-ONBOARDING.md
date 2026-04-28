# AI Agent Skills Hub - Project Governance

This document explains how to use the governance of this Hub in **new projects**.

## 1. Quick Installation (One-Liner)

To enable HB-CLI governance in any project, run:

```bash
curl -sSL https://raw.githubusercontent.com/KlebersonCollab/skills/main/scripts/install-hb.sh | bash
```

This will install `hb` in `~/.local/bin/hb`. Ensure this directory is in your PATH.

## 2. Project Bootstrapping

After installation, initialize governance in the root directory of your new project:

```bash
hb sdd bootstrap
```

This will create the **Triad of Memory** structure in `.specs/project/`:
- `STATE.md`: Progress and current focus.
- `MEMORY.md`: Long-term context and decisions.
- `LEARNINGS.md`: Technical insights and patterns.

## 3. GitHub Actions Integration

To have your project automatically audited following the Hub's standards, create a `.github/workflows/governance.yml` file in your project:

```yaml
name: Project Governance Audit

on:
  push:
    branches: [ main ]
  pull_request:

jobs:
  audit:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      # Use the official Skills Hub Action to setup HB-CLI
      - uses: KlebersonCollab/skills/.github/actions/hb-setup@main
      
      - name: Run Integrity Audit
        run: hb harness audit --deep
        
      - name: Check SDD Compliance
        run: hb sdd audit
```

## 4. Using Hub Skills

You can "equip" your project with specific intelligence from the Hub:

```bash
# Install Python support with all governance rules
hb install python-uv --remote
```

## 5. Available Templates

You can use our base templates as a starting point:
- `templates/go-service`: Base structure for Go services with native governance.
