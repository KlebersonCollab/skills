---
name: skill-factory
version: 2.0.0
description: "Core Framework for high-performance Skill Authoring. Implements Anthropics's 14 patterns (Activation Metadata, Progressive Disclosure, Execution Checklist) for standardized, efficient, and robust agent extensions."
category: skill-management
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Have you rehydrated the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)?
3. **Plan Check**: Does the `plan.md` file define architecture, schemas, and include **Mermaid** diagrams?

---
# Skill Factory: Core Framework (v2.0.0)

> "Every skill deserves a solid foundation. Consistency breeds quality."

---

## 🏗️ Skill Authoring Patterns

### Group 1: Discovery & Selection
1.  **Activation Metadata**: Write "pushy" and clear descriptions (capped at 1024 chars) that trigger Claude correctly.
2.  **Exclusion Clause**: Use explicit "Do NOT use for..." lines to prevent skill hijacking and over-triggering.

### Group 2: Context Economy
3.  **Context Budget**: Every sentence must justify its token cost. Avoid re-explaining common knowledge.
4.  **Progressive Disclosure**: Treat `SKILL.md` as a Table of Contents; move details to sub-files (`REFERENCE.md`, `FORMS.md`).

### Group 3: Instruction Calibration
5.  **Control Tuning**: Match instruction freedom to task fragility (High freedom for reviews, Low for migrations).
6.  **Explain-the-Why**: State the rule, then explain *why* so Claude can generalize to edge cases.
7.  **Template Scaffold**: Provide explicit placeholders for reports/logs to ensure machine-parsable output.
8.  **In-Skill Examples**: Embed 2-3 Input/Output pairs to calibrate style and tone (Few-shot prompting).
9.  **Known Gotchas**: Document concrete failure modes (e.g., "Empty JSON returns on timeout") to prevent hallucinations.

### Group 4: Workflow Control
10. **Execution Checklist**: Provide a copyable checklist that Claude ticks off in the chat to track progress.
11. **Self-Correcting Loop**: Wire in explicit "produce-validate-fix" loops for quality-critical output.
12. **Plan-Validate-Execute**: Insert a verifiable plan (JSON/YAML) before performing side effects.

### Group 5: Executable Code
13. **Utility Bundle**: Ship purpose-built scripts in `scripts/` for deterministic logic.
14. **Autonomy Calibration**: Declare `allowed-tools` in frontmatter to minimize approval friction.

---

## Workflow (3 Phases)

### Phase 1: BOOTSTRAP — Designing the Soul
1.  **Intent Audit**: Define the core task and triggers (**Activation Metadata**).
2.  **Scaffolding**: Run `hb skill new <skill_name>` to generate the directory structure.
3.  **Tiering**: Decide which content goes into `SKILL.md` vs. sub-files (**Progressive Disclosure**).

### Phase 2: AUTHOR — Calibrating instructions
1.  **Gotcha Discovery**: Analyze common failures and document them.
2.  **Instruction Design**: Apply **Control Tuning** and **Explain-the-Why**.
3.  **Few-Shotting**: Select 2 representative examples for the **In-Skill Examples** section.

### Phase 3: VALIDATE — Compliance Audit
1.  **Validator Audit**: Run `skill-factory-validator` to check for pattern compliance.
2.  **Token Audit**: Verify that `SKILL.md` stays under the ~500 line target.
3.  **Registry Update**: Register the new skill in the root `README.md`.

---

## Output Structure

| Artifact | Description |
|----------|-----------|
| `SKILL.md` | The "Central Brain" of the skill (Patterns 1-12). |
| `scripts/` | Deterministic helpers (Pattern 13). |
| `references/` | Deep-dive documentation (Pattern 4). |
| `examples/` | Contextual samples for the agent. |

---

## Quality Rules

- **Zero Redundancy**: If a rule is in `GLOBAL_MANDATES.md`, do not repeat it in the skill.
- **Pushy Descriptions**: The description must explicitly tell the agent *when* to activate.
- **Human-Readable, Agent-Executable**: The skill must be clear to a human author but optimized for LLM attention.

## Prohibited

- NEVER use all-caps imperatives (MUST/ALWAYS) without an explanation (**Pattern 6**).
- NEVER stuff a `SKILL.md` with >800 lines of content; use **Pattern 4**.
- NEVER register a skill without a **Gotchas** section if the domain has known edge cases.

---

## References

- [Skill Authoring Patterns from Anthropic's Best Practices](https://generativeprogrammer.com/p/skill-authoring-patterns-from-anthropics)
- [Anthropic Skill Best Practices](https://platform.claude.com/docs/en/agents-and-tools/agent-skills/best-practices)
