---
name: token-distiller
description: >
  Token density manager. Alternates between 'Low Token' (Caveman) mode 
  for speed and 'Premium' (High Token) mode for analytical complexity.
version: 1.2.0
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
- **Focus**: Use `hb harness focus "Task"` to track optimization milestones.
- **Optimization**: Use `hb harness distill` if the context exceeds 80% of the token limit (Snip, Micro-compact, Collapse).

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

## 2. Context Maintenance Protocols (Micro-Compaction)

Independente do modo de linguagem, o agente deve gerenciar a "saúde" do contexto seguindo estes protocolos técnicos:

### 🗜️ Micro-Compaction (Stubbing)
- **O que:** Substituir resultados de ferramentas antigos ou excessivamente longos por "stubs" (resumos técnicos).
- **Gatilho:** Resultados de ferramentas (bash, read_file, grep) com mais de 10 turns de idade E mais de 5000 caracteres.
- **Ação:** Substituir o conteúdo por: `[Old tool result content cleared to save tokens. Tool: {name}. Summary: {breve resumo de 1 linha}].`
- **Exceção:** Nunca limpe resultados que contenham dados críticos para a tarefa atual em execução.

### 📦 API-Round Grouping
- **O que:** Agrupar raciocínio por rodadas de interação.
- **Protocolo:** A cada 5 interações (API Round-trips), o agente deve fazer um "Self-Summary" interno para consolidar o progresso e descartar caminhos de raciocínio obsoletos.

### 🖼️ Media & Noise Pruning
- **Imagens:** Após 5 turns, referenciar imagens por ID em vez de manter dados base64 (se aplicável).
- **Lixo de Terminal:** Remover saídas repetitivas de comandos `ls`, `grep` ou logs de build após a extração da informação necessária.

## 3. Persistence and Toggle

- The mode persists until the end of the session or until changed.
- Manual triggers: `/mode low`, `/mode high`.
- Automatic detection based on SDD **Task Sizing**.

## 4. Auto-Clarity (Safety Rule)

**The agent MUST abandon Low Token mode and use clear prose for:**
- Safety warnings.
- Confirmations of irreversible actions (e.g., delete database).
- Multi-step sequences where fragment ambiguity might cause execution errors.
- When the user asks for clarification.

## 5. Limits

- **Code:** Always write in a normal and readable way, regardless of the mode.
- **Commits:** Follow the `git-workflow` (Conventional Commits), but prefer extreme brevity in Low Token mode.
