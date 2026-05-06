# Skill Factory: Bootstrap Sub-skill

This sub-skill is responsible for the physical creation of a new skill's infrastructure.

## 🔒 Prerequisites
- Ensure the parent directory `.agents/skills/` exists.
- Have a clear `name` and `description` for the new skill.

## 🛠️ Execution Protocol

### 1. Directory Structure
When bootstrapping a new skill, create the following:

```bash
mkdir -p .agents/skills/[name]/.specs/{features,project,codebase}
mkdir -p .agents/skills/[name]/{examples,references,resources}
touch .agents/skills/[name]/{SKILL.md,README.md,MEMORY.md,LEARNINGS.md,STATE.md,DECISIONS.md,tasks.md}
```

### 2. Mandatory Content Seeds

#### SKILL.md
Every `SKILL.md` must start with this frontmatter:
```markdown
---
name: [name]
version: 0.1.0
description: [description]
category: [category]
---
```

#### STATE.md
Initialize with a "Fresh Start" status:
```markdown
# State: [Name]
- Status: Initialized
- Current Phase: Discovery
- Next Steps: Create first feature spec.
```

### 3. Verification
After bootstrapping, call `skill-factory-validator` to ensure the skeleton is ready for use.
