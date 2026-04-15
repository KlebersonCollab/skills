# Spec: Azure DevOps Skill v1.1.0 Upgrade

## 1. Goal
Upgrade the `azure-devops` skill to version 1.1.0, expanding its capabilities from basic module management to full **Governance and Infrastructure-as-Code (IaC) Support**. This includes managing Service Connections, Variable Groups, User/Team administration, and advanced security policies.

## 2. Target Audience
DevOps Engineers and Platform Leads who need to automate the setup and governance of Azure DevOps organizations and projects.

## 3. Core Requirements (Functional)
- **Service Connections (Endpoints)**: CRUD operations for integrating external services (Azure, AWS, Docker).
- **Variable Governance**: Full support for Variable Groups (`az pipelines variable-group`) and individual pipeline variables.
- **Administration & Security**:
    - Manage Users and Teams (`az devops user/team`).
    - Audit and manage Extensions (`az devops extension`).
    - Detailed permission and security group management.
- **Advanced CLI Integration**:
    - Contextual configuration (`az devops configure --defaults`).
    - Raw API invocation (`az devops invoke`) for unmapped REST endpoints.
- **Wiki Management**: Support for project and code-based wikis.

## 4. Acceptance Criteria (BDD Format)

### Scenario 1: Setup Infrastructure Variables
**Given** a requirement to share database credentials across multiple pipelines
**When** the `azure-devops` skill is used
**Then** it must provide guidance to create a Variable Group, link it to a Key Vault if necessary, and authorize it for specific pipelines.

### Scenario 2: Audit Project Security
**Given** a security audit requirement
**When** the `azure-devops` skill analyzes the project
**Then** it must list active users, their access levels, and any branch policies that are currently failing or missing.

### Scenario 3: Invoke Custom REST API
**Given** a need for a feature not present in standard CLI commands (e.g., specific Dashboard widget update)
**When** the `azure-devops` skill is invoked
**Then** it must suggest using the `az devops invoke` command with the correct area, resource, and JSON payload.

## 5. Technical Requirements (Non-Functional)
- **Standard-Aligned**: All commands must strictly follow the `azure-devops-cli` syntax.
- **Visualization**: Must include Mermaid diagrams for Infrastructure/Variable inheritance flows.
- **Security**: Reinforce PAT token safety during administrative operations.

## 6. Architecture & Design
- `azure-devops/SKILL.md`: Update to v1.1.0 with the new Governance workflow.
- `azure-devops/references/`:
    - `infrastructure-and-variables.md`: Variable Groups and Service Connections.
    - `administration-and-security.md`: Users, Teams, Permissions, and Extensions.
    - `advanced-cli-commands.md`: `invoke`, `configure`, and `wiki`.
- `azure-devops/examples/`:
    - `service-connection-template.json`: Schema for creating endpoints.
    - `variable-group-mapping.mermaid`: Visualizing variable distribution.
    - `raw-api-invoke-sample.sh`: Examples of custom REST calls.
