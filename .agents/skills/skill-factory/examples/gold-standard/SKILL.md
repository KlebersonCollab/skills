---
name: gold-standard-example
version: 1.0.0
description: "A perfect example of a v2.3.0 compliant skill following the SDD Purist standard."
category: examples
---

## 🔒 Prerequisites (Mandatory)
Before execution:
0. **Mode Check**: Verify `.hub-mode` and apply `token-distiller` guidelines.
1. **Context Check**: Read `.specs/project/STATE.md`, `.specs/project/MEMORY.md`, and `.specs/project/LEARNINGS.md`.
2. **Knowledge Check**: Follow the **Knowledge Verification Chain**.

---

# Gold Standard Example (v2.3.0)

> "Purity in structure leads to precision in execution."

---

## 🧩 Delegation Matrix
| Phase | Sub-Skill | Primary Artifact | Purpose |
|---|---|---|---|
| **DESIGN** | `example-sub-skill` | `blueprint.md` | Logic design. |

---

## 🛠️ Operational Protocols

### 1. Knowledge Verification Chain
1. **Existing Code**: Scan for established patterns in the repository.
2. **Internal Specs**: Consult `.specs/features/` and `STATE.md`.
3. **Global Mandates**: Read `.specs/codebase/GLOBAL_MANDATES.md`.
4. **Project Memory**: Check `.specs/project/MEMORY.md`.
5. **Flag Uncertainty**: Never assume.

### 2. Safety Valve
If task complexity exceeds 5 interdependent modules, escalate to `sdd-orchestrator`.

---

### 3. Observable Governance (Mandatory v2.3.0)
Every artifact generated MUST include a structured metadata block.
- **Zero Local Memory**: DO NOT create `STATE.md`, `MEMORY.md`, `LEARNINGS.md`, or `tasks.md` in this directory. Use `.specs/` instead.

#### Metadata Protocol
Every artifact MUST include this block.

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
- NEVER create local memory files (`tasks.md`, `STATE.md`, etc.).
- NEVER skip the Knowledge Verification Chain.
- NEVER use placeholders.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "GOLD-STANDARD-V230"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T10:30:00Z"
evidence_checksum: "NONE"
```
