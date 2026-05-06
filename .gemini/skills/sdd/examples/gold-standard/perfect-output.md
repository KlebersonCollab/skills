# Gold Standard Output: SDD Execution (Observable Governance v2.3.0)

## 1. Requirements (spec.md)
The agent generated a clear specification with metadata:

```markdown
### AC-1: Token Generation
**Given** valid credentials, **When** /login is called, **Then** a JWT is returned.

<!-- @sdd-state -->
{ "version": "2.3.0", "phase": "SPECIFY", "status": "COMPLETED" }
```

## 2. Technical Design (plan.md)
The plan includes metadata at the end:

```markdown
## Design
``mermaid
sequenceDiagram
    User->>API: POST /login (credentials)
    API->>DB: Validate User
    DB-->>API: User Data
    API->>API: Sign JWT
    API-->>User: 200 OK (token)
```

<!-- @sdd-state -->
{ "version": "2.3.0", "phase": "SPECIFY", "status": "COMPLETED" }
```

## 3. Atomic Tasks (tasks.md)
**Mandatory Table Format with Evidence**:

| Phase | Task | Status | Evidence (Commit/Log) |
| :--- | :--- | :---: | :--- |
| **1. AUTH** | Create JWT signing logic | [x] | commit: 4a2b1c3 |
| **2. API** | Implement /login route | [ ] | |

---

<!-- @sdd-state -->
{ "version": "2.3.0", "status": "IN_PROGRESS" }
```

## 4. Why this is Gold Standard (v2.3.0)
- **Machine-Readable**: Every artifact can be audited by tools via the `<!-- @sdd-state -->` block.
- **Evidence-Based**: Completion is proven by commit hashes, not just checkmarks.
- **Traceability**: Full lifecycle visibility from Spec to Verification.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "GOLD-STANDARD"
phase: "MANAGEMENT"
status: "COMPLETED"
last_update: "2026-05-06T09:42:00Z"
evidence_checksum: "NONE"
```
