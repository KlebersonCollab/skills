# SDD Specification Template: spec.md

This is the source of truth for **functional requirements** and **business logic**.

---

## 1. Context & Goals
- **Problem Statement**: What problem are we solving?
- **Scope**: What is included and what is NOT included.
- **Success Criteria**: How do we know we are done?

## 2. Requirements (BDD Scenarios)
*Define behavior using Given/When/Then syntax.*

### Feature: [Feature Name]

**Scenario 1: [Short Description]**
- **Given** [Initial Context]
- **When** [Action]
- **Then** [Expected Outcome]

## 3. Constraints & Risks
- **Security**: [Considerations]
- **Performance**: [Expected limits]
- **Dependencies**: [Required libraries or services]

---
*Note: Technical implementation details go to `plan.md`.*

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "FEAT-ID"
phase: "SPECIFY"
status: "COMPLETED"
last_update: "ISO-TIMESTAMP"
evidence_checksum: "NONE"
```
