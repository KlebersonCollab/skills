# Skill Factory: Bootstrap Sub-skill (Purist Edition)

This sub-skill defines the structural blueprint of a new capability. It focuses on the logical arrangement of intelligence.

## 🔒 Prerequisites
- Clear functional domain (e.g., `fastapi-expert`).
- Alignment with the Hub's naming conventions.

## 🛠️ Execution Protocol

### 1. The Blueprint
The agent must design the skill structure as a first-class citizen. This is not just about folders, but about the *reasoning flow*.

**Standard Layout:**
1.  **Entry Point**: `SKILL.md` (The "Brain").
2.  **Metadata**: `README.md` (The "Identity").
3.  **Memory Assets**:
    - `STATE.md`: Ephemeral progress.
    - `MEMORY.md`: Long-term patterns/preferences.
    - `LEARNINGS.md`: Discovered wisdom.
    - `DECISIONS.md`: Architectural log.
4.  **Verification**: `tasks.md` and `.specs/` hierarchy.

### 2. Implementation Strategy
Instead of relying on automated scripts, the agent must manually construct the environment to ensure each file is initialized with the correct context.

**Initialization Sequence:**
- [ ] Define the `SKILL.md` frontmatter.
- [ ] Establish the **Delegation Matrix**.
- [ ] Create the **Knowledge Verification Chain**.
- [ ] Initialize the **Memory Triad** with contextual seeds.

### 3. Safety Valve
If the skill design requires more than 5 sub-skills or complex external dependencies, **STOP** and re-evaluate if it should be a "Multi-Agent Orchestrator" instead.
