# Specification: HB-CLI Expansion Pack

## 1. Overview
Expand the HB-CLI (Harness Binary) with specialized tools to automate Hub governance, decision recording, and AI context optimization.

## 2. Core Features

### 2.1. HB ADR (Architecture Decision Records)
- **`hb adr new "Title"`**:
    - Generates a new Markdown file in `.specs/architecture/`.
    - Format: `ADR-XXXX-title-in-kebab-case.md`.
    - Includes metadata: Status (Proposed), Date, Context, Decision, Consequences.
- **`hb adr list`**:
    - Scans `.specs/architecture/` and displays a table of ADRs with their status.

### 2.2. HB Audit
- **`hb audit`**:
    - Performs local static analysis.
    - Python: `ruff check`.
    - Go: `go vet` and `golangci-lint` (if available).
    - Hub Compliance: Checks for mandatory files in skills (`SKILL.md`, `CHANGELOG.md`).
    - Output: A combined health score.

### 2.3. HB Learn
- **`hb learn "Subject" --desc "Description" --cat "Category"`**:
    - Appends a structured entry to `.specs/project/LEARNINGS.md`.
    - Format: Date, Subject, Description, Category (Pattern, Bug, Optimization).

### 2.4. HB Skill
- **`hb skill new <name>`**:
    - Scaffolds a new skill folder with:
        - `SKILL.md` (template with frontmatter).
        - `CHANGELOG.md`.
        - `README.md`.
- **`hb skill validate`**:
    - Validates the skill structure against Hub standards.

### 2.5. HB Distill
- **`hb distill <file>`**:
    - Strips comments, extra whitespaces, and optimizes code for context.
- **`hb distill session`**:
    - Summarizes the current interaction into a `MEMORY.md` block.

### 2.6. HB Doctor
- **`hb doctor`**:
    - Checks:
        - HB binary version.
        - Existence of `.cursorrules`.
        - Consistency of agent folders.
        - Symlink health.
    - Fix: `--fix` flag to automatically resolve common issues.

### 2.7. HB SDD Review
- **`hb sdd review <feature>`**:
    - Automated sanity check: Does the feature have all 6 SDD artifacts?
    - Content Check: Do tasks in `tasks.md` match the implementation status?

## 3. Technical Requirements
- Language: Go.
- CLI Framework: Cobra.
- UI: Fatih Color for styled output.
- Reliability: 100% test coverage for core logic (internal/ logic).
