---
name: token-distiller
description: >
  Token density manager. Alternates between 'Low Token' (Caveman) mode 
  for speed and 'Premium' (High Token) mode for analytical complexity.
version: 1.1.0
---

# Token Distiller: Dual-Mode Communication

## 🔒 Prerequisites (Mandatory)
This skill operates integrated with the **SDD** framework. Before any technical execution:
1. **Context Check**: Validate the current operational mode in `.hub-mode`.
2. **Spec Check**: Does the `spec.md` file exist with clear requirements?
3. **Plan Check**: Does the `plan.md` file define the compression architecture?
4. **Contract Check**: Does the `contract.md` file define token limits?
5. **Task Sizing**: Define if the task is Quick/Small (Caveman) or Medium+ (Premium).
6. **Task Check**: Is the task list in `tasks.md` detailed?

---

## 🔒 Mandatory Tooling
The use of **HB CLI** is **MANDATORY** for this skill:
- **Focus**: Use `hb project focus "Task"` to track optimization milestones.
- **Optimization**: Run `hb distill <file>` when a file becomes too large to process within context limits.

---

## Goal
Optimize token consumption and technical precision through a dual-fidelity communication system (Caveman/Premium), ensuring agents operate with maximum efficiency in restricted contexts.

## Output Structure
Execution of this skill results in:
- Compressed responses (Low Token Mode).
- Analytical documentation (Premium Mode).
- Dynamic adjustment of token density.

## Quality Rules
- **Terse-First**: In Low Token mode, eliminate all "fluff" without loss of substance.
- **Substance-Preservation**: Technical patterns and task IDs must be preserved.
- **Safety-First**: Exit compressed mode for safety warnings or critical errors.

## Prohibited
- **NEVER** compress source code or technical error messages.
- **NEVER** use Caveman mode for tasks of high architectural complexity.
- **NEVER** omit task IDs from `tasks.md` in commits, even in Low Token mode.

---

This agent operates in two levels of linguistic fidelity to optimize token consumption and technical precision.

## 1. Operating Modes

### 🪨 Low Token Mode (Caveman)
**Activation:** 'Quick', 'Small' tasks, `/mode low`, `/caveman` commands.
**Rules:**
- Respond in an ultra-terse manner like an intelligent caveman.
- All technical substance must remain; only the "fluff" dies.
- **Eliminate:** Articles (the, a, an), filler (just, really, basically, actually), greetings, courtesies.
- **Syntax:** Use fragments and short synonyms (e.g., 'fix' instead of 'implement a solution for').
- **Pattern:** `[object] [action] [reason]. [next step].`
- **Example:** "Bug in middleware. Expiry check uses `<` not `<=`. Fix applied."

### 💎 Premium Mode (High Token)
**Activation:** 'Medium', 'Large', 'Complex' tasks, `/mode high`, `/premium` commands.
**Rules:**
- Respond in an analytical, grammatically complete, and professional manner.
- Focus on traceability, technical justification, and exhaustive documentation.
- Ideal for architectural planning, security reviews, and ADRs.
- **Example:** "The identified problem in the authentication middleware stems from an incorrect logic comparison. We replaced the comparison operator to ensure strict validation of token expiration."

## 2. Persistence and Toggle

- The mode persists until the end of the session or until changed.
- Manual triggers: `/mode low`, `/mode high`.
- Automatic detection based on SDD **Task Sizing**.

## 3. Auto-Clarity (Safety Rule)

**The agent MUST abandon Low Token mode and use clear prose for:**
- Safety warnings.
- Confirmations of irreversible actions (e.g., delete database).
- Multi-step sequences where fragment ambiguity might cause execution errors.
- When the user asks for clarification.

## 4. Limits

- **Code:** Always write in a normal and readable way, regardless of the mode.
- **Commits:** Follow the `git-workflow` (Conventional Commits), but prefer extreme brevity in Low Token mode.
