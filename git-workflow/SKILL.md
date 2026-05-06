---
name: git-workflow
version: 2.2.0
description: "Git workflow standards, including branching strategies, commit conventions, and SDD integration."
category: development-workflow
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Did you rehydrate the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define the architecture, schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---

## Goal

Establish version control standards, branching strategies, and collaborative development that ensure traceability, facilitate code reviews, and maintain project history integrity, especially when operating under the SDD framework.

---

## Output Structure

Execution of this skill results in the following artifacts:

| Artifact | Format | Description |
|----------|---------|-----------|
| **Clean History** | Git Log | Linear and readable commit history. |
| **Pull Request** | Markdown | Documented PR following the project template. |
| **Git Config** | `.gitmessage` | Local commit template configuration. |

---

## Branching Strategies

### 1. GitHub Flow (Simple and Recommended)
Ideal for continuous deployment and small to medium-sized teams.
- `main` is always production-ready.
- Create `feature/` branches from `main`.
- Open a Pull Request for review.
- After approval and testing, merge into `main`.

### 2. Trunk-Based Development (High Velocity)
For teams with robust CI/CD and Feature Flags.
- Short, frequent commits directly to `main` (or very short-lived branches).
- CI must pass before any integration.

### 3. GitFlow (Enterprise Release Cycle)
For scheduled releases and complex projects.
- `main` (production), `develop` (integration).
- `feature`, `release`, and `hotfix` branches.

---

## Conventional Commits

Mandatory format: `<type>(<scope>): <subject>`

| Type | Usage | Example |
|------|-----|---------|
| `feat` | New functionality | `feat(auth): add SSO support` |
| `fix` | Bug fix | `fix(api): handle connection timeout` |
| `docs` | Documentation | `docs(readme): add setup steps` |
| `refactor` | Code refactoring | `refactor(db): optimize query logic` |
| `test` | Tests | `test(unit): add auth validation tests` |

*See the full guide at: [references/conventional-commits.md](references/conventional-commits.md)*

---

## Git in the SDD Workflow (Mandatory)

Integrating Git with the SDD cycle is fundamental for traceability and history integrity:

1.  **Discovery Phase**: Before starting, ensure your local `main` is synchronized with the remote using **`git pull --rebase origin main`**.
2.  **Implementation Phase**: 
    - **Atomic Branches**: Each feature/fix must have its own branch. Initialize it with **`git checkout -b feature/<name>`**.
    - **Commits per Task**: Each completed task in `tasks.md` must generate at least one clear commit.
    - **Conventional Commits**: Use **`git commit -m "<type>(<scope>): <message>"`**.
    - **Language**: All commit messages **MUST** be written in **English**, even if technical documentation is in Portuguese.
    - **Atomicity**: Do not mix changes from different tasks in the same commit.
3.  **Review Phase**:
    - **Linearity**: Always run **`git pull --rebase origin main`** before pushing to keep history clean and resolve conflicts early via rebase.
    - **Evidence**: The Pull Request link must be included in `validation-report.md`.

### SDD Commit Cycle Diagram
```mermaid
graph TD
    A[Start Task in tasks.md] --> B[Code Implementation]
    B --> C[Local Test Execution]
    C --> D{Passed?}
    D -- No --> B
    D -- Yes --> E["git commit -m 'type: description'"]
    E --> F[Update tasks.md as Completed]
    F --> G{More tasks?}
    G -- Yes --> B
    G -- No --> H[Open Pull Request]
```

---

## Integration Mandates: Merge vs Rebase

### Use REBASE for:
- Updating your local branch with the latest changes from `main`.
- Cleaning up intermediate commits (squash) before sending for review.
- Maintaining a linear history without unnecessary merge commit noise.

### Use MERGE for:
- Integrating a completed feature branch into `main`.
- When the branch is shared with other developers (to avoid force pushing to public branches).

---

## Quality Rules

- **English-Only**: All commit messages must be written in English.
- **Conventional-Commits**: Mandatory use of the Conventional Commits standard.
- **Linear-History**: Preference for `rebase` to maintain a clean history.
- **Atomic-Commits**: Small commits focused on a single change.

---

## Prohibited

- **NEVER** perform direct commits to `main` without review (except in Quick SDD mode).
- **NEVER** commit secrets, API keys, or `.env` files.
- **NEVER** use vague commit messages (e.g., "update", "fix").
- **NEVER** send giant Pull Requests (>500 lines).
- **NEVER** write commit messages in any language other than English.

---

## References and Examples
- [Conventional Commits Guide](references/conventional-commits.md)
- [Detailed Branching Strategies](references/branching-strategies.md)
- [Git Message Template](examples/gitmessage.template)
- [Pull Request Template](examples/pull-request.md)


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.362355Z"
evidence_checksum: "NONE"
```
