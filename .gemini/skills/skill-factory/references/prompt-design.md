# Prompt Design Guide for Skills & Sub-Skills

Designing effective prompts for skills is crucial for ensuring AI agents can operate autonomously and reliably. Use these best practices when creating or updating skills.

## 1. Identity & Role definition
Define a clear persona and mission for the skill or sub-skill.
- **Rule**: Start with a high-level goal that sets the boundaries of the skill.
- **Good Example**: "You are the 'Architecture Expert', responsible for maintaining the structural integrity of the project and enforcing architectural patterns."

## 2. Contextual Grounding
Provide the necessary context for the skill to operate correctly.
- **Rule**: Use absolute paths for file operations and explain the environment structure.
- **Good Example**: "This skill operates within the project root directory. Always verify the existence of the `pyproject.toml` file before making changes."

## 3. Instruction Clarity (The 'Core Instructions')
Break down instructions into logical sections: Goal, Output Structure, Quality Rules, and Prohibited actions.
- **Rule**: Use Markdown headers and bullet points for readability.
- **Good Example**: 
  - **Quality Rules**: Follow DRY and SOLID principles.
  - **Prohibited**: Never modify files in the `.git` directory.

## 4. Constraint-Based Engineering
Clearly state what the skill **cannot** do to prevent hallucination or unsafe actions.
- **Rule**: Use the `## Prohibited` section for all critical constraints.
- **Good Example**: "NEVER rewrite an entire file if a surgical change (replace) is sufficient."

## 5. Output Structure
Define the expected artifacts or responses.
- **Rule**: Specify the file format and the key information it must contain.
- **Good Example**: "Executing this skill results in an `ADR.md` following the official template."

## 6. Structural Rigor
- Use clear headers and sections.
- Use Markdown for structure and clarity.
- Define explicit roles, constraints, and objectives.

## 7. Few-Shot Examples
- Provide 3-5 high-quality examples of desired inputs and outputs.
- This significantly improves the agent's ability to follow complex patterns.

## 8. Iterative Refinement
Prompts are never final. Use the global `.specs/project/LEARNINGS.md` to identify failures and improve the instructions.
- **Rule**: Update the prompt based on real-world edge cases encountered by the agents.
- **Good Example**: "Added a rule to ignore `.DS_Store` files after a common listing error."

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "REFACTOR-SKILL-FACTORY-V230"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T10:28:00Z"
evidence_checksum: "NONE"
```
