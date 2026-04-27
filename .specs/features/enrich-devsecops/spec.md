# Spec: Enrichment of DevSecOps Expert with Agentic Guardrails

## 1. Goal
Integrate "Agentic Operational Safety" guidelines into the `devsecops-expert` skill, based on industry-standard dangerous pattern detection (mined from high-maturity agent architectures).

## 2. Requirements (FR)
- **FR-1**: Define a comprehensive list of "Dangerous Command Prefixes" that agents must audit before proposing.
- **FR-2**: Establish a mandatory "Pre-proposal Audit" protocol for shell commands.
- **FR-3**: Include specific guidance for Cross-Platform (Unix/Windows) code execution risks.
- **FR-4**: Add "Prohibited" rules specifically for destructive command variants (e.g., `rm -rf /`, `curl | bash`).

## 3. Acceptance Criteria (AC - BDD)

### Scenario 1: Proposing a Shell Command
**Given** that the agent is using the `devsecops-expert` skill
**When** the agent needs to propose a bash command to the user
**Then** it must check if the command prefix matches the `DANGEROUS_BASH_PATTERNS` list
**And** if a match is found, it must include a clear warning about the risks and ask for explicit confirmation highlighting the specific danger.

### Scenario 2: Windows vs Unix Parity
**Given** a cross-platform environment
**When** proposing code execution via interpreters (Python, Node, etc.)
**Then** the agent must apply the same rigor for both `Bash` and `PowerShell` aliases.

## 4. Constraints
- Must maintain 100% parity with the `src` patterns found during the Cybernetic Mining phase.
- Must not duplicate existing `devsecops-expert` content, only enrich it.
