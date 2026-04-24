# Brownfield Mapping (Existing Projects)

**Trigger:** "Map codebase", "Analyze existing code", "Document current architecture"
**Purpose:** Understand the existing project structure before adding new features.

## Process
Before starting, use code analysis tools (ast-grep, ripgrep) to systematically explore:
1. Directory structure
2. Technology stack from dependency manifests
3. Patterns extracted from representative code examples
4. External integrations and dependencies
5. Points of attention (Tech debt, bugs, security risks)

## Artifacts Generated (in `.specs/codebase/`)

### 1. STACK.md
Documents the technology stack.
- **Size Limit:** ~1,200 words
- **Content:** Language, Framework, Runtime, Database, Test Libraries.

### 2. ARCHITECTURE.md
Documents architectural patterns and data flow.
- **Size Limit:** ~2,400 words
- **Content:** Structural diagram (Mermaid), identified patterns, main data flow.

### 3. CONVENTIONS.md
Documents code style and naming conventions.
- **Size Limit:** ~1,800 words
- **Content:** Filenames, functions, error handling, and typing based on evidence from code.

### 4. STRUCTURE.md
Documents the directory layout.
- **Size Limit:** ~1,200 words
- **Content:** Directory tree (max 3 levels), module organization.

### 5. TESTING.md
Documents the testing infrastructure.
- **Size Limit:** ~2,400 words
- **Content:** Test frameworks, organization, execution commands, coverage, and parallelism support.

### 6. INTEGRATIONS.md
Documents integrations with external services.
- **Content:** APIs, Webhooks, background jobs, and their authentication methods.

### 7. CONCERNS.md
Surface of actionable warnings about the codebase.
- **Size Limit:** ~3,000 words
- **Content:** Technical Debt, known bugs, security flaws, performance bottlenecks, and fragilities. *Mandatory to always include file/path as evidence.*
