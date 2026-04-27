# Brownfield Mapping (TECHNICAL-MAP)

**Trigger:** "Map codebase", "Analyze existing code", "Document current architecture"

**Purpose:** Understand the existing project structure and identify risks before planning features.

## The Output: `.specs/codebase/TECHNICAL-MAP.md`

Instead of multiple fragmented files, the Explorer consolidates knowledge into a single high-fidelity map.

### Section 1: Tech Stack & Ecosystem
- **Core:** Languages, frameworks, and runtimes with versions.
- **Tools:** Package managers, build tools, and CLI utilities.
- **Testing:** Unit, Integration, and E2E frameworks.

### Section 2: Architecture & Patterns
- **High-Level Structure:** Directory organization and layer definitions.
- **Observed Patterns:** Concrete examples of how logic is decoupled (e.g., Services, Repositories, Hooks).
- **Data Flow:** Mermaid diagram of primary request/response cycles.

### Section 3: Core Conventions
- **Naming:** Files, functions, and variables based on codebase evidence.
- **Style:** Error handling patterns, type safety approach, and commenting style.

### Section 4: Critical Risks (Concerns)
*This is the most important section for the Safety Valve protocol.*
- **Technical Debt:** Shortcuts that impact maintainability.
- **Fragile Areas:** Components that break easily or have low test coverage.
- **Security & Performance:** Known bottlenecks or exposed patterns.
- **Mandatory:** Every concern must include file paths as evidence.

### Section 5: External Bridges
- **Integrations:** Third-party APIs, databases, and message queues.
- **Authentication:** How external services are secured.

---

## Analysis Protocol

1. **Systematic Scan:** Use `ls -R` and dependency manifests to build Section 1.
2. **Pattern Extraction:** Analyze 5-10 representative files to define Sections 2 and 3.
3. **Risk Discovery:** Search for "TODO", "FIXME", "HACK", or complex legacy structures for Section 4.
4. **Integration Audit:** Check environment variables and client initializations for Section 5.

## Quality Rules
- **Evidence-Based:** Reference specific file paths for every claim.
- **Current State Only:** Document what IS, not what you think it should be.
- **Concise:** Use bullet points and tables. Avoid prose.
