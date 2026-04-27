---
name: azure-devops
version: 1.1.0
description: "Skill for professional management of Azure DevOps (AzDO). Allows managing Boards, Repos, Pipelines, Artifacts, Governance (Variable Groups, Service Connections), and Administration efficiently."
category: devops-automation
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the guidelines of the `token-distiller` skill.
1. **Context Check**: Have you rehydrated the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define the architecture, schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---
# Azure DevOps (AzDO)

> Professional interface for orchestrating the development lifecycle and platform governance in Azure DevOps.

---

## Goal

Empower the agent to comprehensively manage Azure DevOps resources, from task and code management to infrastructure automation (Service Connections), variable governance, and user/permission administration, using best practices for security and CLI automation.

---

## Workflow (5 Phases)

### Phase 1: AUTH — Connection and Security
1.  **Verify Credentials**: Ensure the PAT is available and configured.
2.  **Configure Context**: Use `az devops configure --defaults` to fix organization and project.

### Phase 2: INFRA — Governance and IaC
1.  **Service Connections**: Create or audit Service Endpoints for external integrations.
2.  **Global Variables**: Manage shared Variable Groups and integrate them with Key Vaults.
3.  **Pipeline Security**: Authorize the use of critical resources (Pools, Queues, Variable Groups) for specific pipelines.

### Phase 3: PLANNING — Boards and Work Items
1.  **WIQL Queries**: Execute complex filters for backlog triage.
2.  **Item Management**: Create and update Work Items via JSON Patch abstraction.

### Phase 4: DELIVERY — Repos and Code Flow
1.  **PR Lifecycle**: Manage the complete cycle of Pull Requests and reviews.
2.  **Branch Policies**: Validate and apply branch protection rules.

### Phase 5: OPS & ADMIN — Management and Monitoring
1.  **CI/CD Ops**: Trigger builds, releases, and monitor execution logs.
2.  **User Admin**: Manage users, teams, and security permissions.
3.  **Extensions & Wiki**: Administer installed extensions and maintain project documentation via Wiki.

---

## AzDO CLI Commands Reference

| Area | Main Command | Purpose |
|------|-------------------|-----------|
| **Infra** | `az pipelines variable-group` | Management of variable groups |
| **Infra** | `az devops service-endpoint` | Management of external connections |
| **Admin** | `az devops user / team` | People and teams administration |
| **Admin** | `az devops security` | Permissions and security groups |
| **Low-Level**| `az devops invoke` | Direct calls to the REST API |

---

## Quality Rules

- **Security First**: Administrative operations require double confirmation of the target organization.
- **Least Privilege**: Suggestions for variable groups and connections should always limit scope to necessary pipelines only.
- **Standard Syntax**: Use official `az devops` CLI syntax over raw REST calls unless using `invoke`.
- **Infrastructure-as-Code**: Prefer YAML-based pipelines and Variable Groups over UI-defined configurations.

## Prohibited

- **NEVER** suggest creating Service Connections with "Owner" or "Contributor" global permissions without justification.
- **NEVER** ignore the `configure --defaults` command when starting a new administrative session.
- **NEVER** delete Variable Groups or Service Endpoints without checking dependencies in active pipelines.

---

## Reference Documentation

This skill includes detailed reference documentation:

1. **[Authentication](references/authentication.md)** — PAT, contexts, and `configure`.
2. **[Infrastructure & Variables](references/infrastructure-and-variables.md)** — Variable Groups and Service Connections.
3. **[Administration & Security](references/administration-and-security.md)** — Users, Teams, and Permissions.
4. **[Boards Management](references/boards-management.md)** — Work Items and WIQL.
5. **[Repos & Code Flow](references/repos-code-flow.md)** — Pull Requests and Policies.
6. **[Pipelines CI/CD](references/pipelines-cicd.md)** — Execution and monitoring.
7. **[Advanced CLI Commands](references/advanced-cli-commands.md)** — `invoke` and `wiki`.

---

## Output Structure

The execution of this skill results in the following standardized artifacts:

| Artifact | Format | Description |
|----------|---------|-----------|
| **Infra Spec** | `.json` / `.yaml`| Definition of Variable Groups or Endpoints. |
| **Security Audit**| `.md` | Users, permissions, and policies report. |
| **API Invoke** | `.sh` | Direct call scripts via `az devops invoke`. |
| **WIQL Query** | `.wiql` | Structured queries for Azure Boards. |
