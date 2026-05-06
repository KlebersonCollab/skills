---
name: skill-factory-bootstrap
version: 2.3.0
description: "Bootstrap agent for Skill Factory. Designs and generates the logic-first scaffolding for a new skill (Purist Edition)."
category: skill-management
---

## 🔒 Prerequisites (Mandatory)
Before execution:
0. **Mode Check**: Verify `.hub-mode` and apply `token-distiller` guidelines.
1. **Context Check**: Read `.specs/project/STATE.md`, `.specs/project/MEMORY.md`, and `.specs/project/LEARNINGS.md`.
2. **Intent Audit**: Clearly define the skill's functional domain and core task.

---

# Skill Factory: Bootstrap Agent (v2.3.0)

> "A skill's soul is defined by its structure."

---

## 🛠️ Execution Protocol

### Step 1: Design the Blueprint
Design the directory `<skill_name>/` at the repository root. Map out which sub-skills are needed and which logical patterns from the **14 Anthropic Patterns** apply.

### Step 2: Update Project Memory
Instead of local files, update the global project memory:
- **STATE.md**: Log the new skill creation and status.
- **DECISIONS.md**: Log architectural trade-offs for the new skill.
- **tasks.md**: Create a task list using the **Observable Table Format** (Table with Evidence column).

### Step 3: Generate Core Instruction (`SKILL.md`)
Apply the **Gold Standard Template** below. Ensure the `Knowledge Verification Chain` is present to prevent hallucinations.

---

## 🏗️ Gold Standard Templates

### SKILL.md Template
```markdown
---
name: {{skill_name}}
version: 0.1.0
description: "{{description}}"
category: {{category}}
---

## 🔒 Prerequisites (Mandatory)
Before execution:
0. **Context Check**: Read `.specs/project/STATE.md`, `.specs/project/MEMORY.md`, and `.specs/project/LEARNINGS.md`.
1. **Knowledge Check**: Follow the **Knowledge Verification Chain**.

---

# {{Title}}

> {{tagline}}

---

## 🧩 Delegation Matrix
| Phase | Sub-Skill | Primary Artifact | Purpose |
|---|---|---|---|
| PHASE | SUB-SKILL | ARTIFACT | PURPOSE |

---

## 🛠️ Operational Protocols

### 1. Knowledge Verification Chain
1. **Existing Code**: Scan for established patterns.
2. **Internal Specs**: Consult `.specs/features/` and `STATE.md`.
3. **Project Documentation**: Check READMEs and official docs.
4. **MCP/External Tools**: Query `context-graph` or web.
5. **Flag Uncertainty**: Never assume.

### 2. Safety Valve
[Define complexity triggers for escalation]

---

### 3. Observable Governance (Mandatory v2.3.0)
Every artifact generated MUST include a structured metadata block.

#### Metadata Protocol
At the end of every `.md` artifact, append the following block:
```markdown
<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "FEAT-ID"
phase: "PHASE"
status: "STATUS"
last_update: "ISO-TIMESTAMP"
evidence_checksum: "EVIDENCE"
```
```

---

## 🚫 Prohibited
- NEVER skip the **Knowledge Verification Chain** in the template.
- NEVER skip the **Observable Governance** metadata block in generated artifacts.
```

### Memory Asset Template
```markdown
# [Asset Name]: [Skill Name]
- Status: Initialized
- Current Focus: [Current Focus]
```

---

## 🚫 Prohibited
- NEVER modify system directories (e.g., `.agents/`).
- NEVER create local `STATE.md`, `MEMORY.md`, or `LEARNINGS.md` files; use `.specs/project/`.
- NEVER create empty instruction files.
- NEVER skip the **Knowledge Verification Chain** in the template.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T09:51:00Z"
evidence_checksum: "NONE"
```
