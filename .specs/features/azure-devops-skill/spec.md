# Spec: Azure DevOps Skill

## 1. Goal
Create a specialized skill for managing Azure DevOps (AzDO) resources through the Gemini CLI. This skill will provide a professional and efficient interface for Boards, Repos, Pipelines, Test Plans, and Artifacts, abstracting the complexity of the AzDO REST API and JSON Patch operations.

## 2. Target Audience
Software engineers and DevOps specialists using Gemini CLI to automate and manage their Azure DevOps environments.

## 3. Core Requirements (Functional)
- **Azure Boards**: 
    - CRUD for Work Items (Bugs, Tasks, Stories).
    - Query execution via WIQL (Work Item Query Language).
    - Backlog and iteration management.
- **Azure Repos**:
    - Repository listing and management.
    - Pull Request (PR) lifecycle: create, list, comment, approve.
    - Branch policies check.
- **Azure Pipelines**:
    - List and run build/release pipelines.
    - Monitor logs and download artifacts.
    - Manage Variable Groups and Service Connections.
- **Azure Test Plans & Artifacts**:
    - Manage test plans and suites.
    - List feeds and package versions (NuGet, npm).

## 4. Technical Requirements (Non-Functional)
- **Security**: Mandatory use of Personal Access Tokens (PAT) via environment variables (e.g., `AZURE_DEVOPS_EXT_PAT`).
- **JSON Patch Abstraction**: Helper methods to simplify `add`, `replace`, and `remove` operations in Work Items.
- **Multi-Tenancy**: Easy switching between organizations and projects using a config file or environment variables.
- **Resilience**: Implement retry logic for rate limiting (429 status code) and clear error messaging.
- **Consistency**: Follow the established pattern of `python-uv` and `flutter-fvm` skills.

## 5. Acceptance Criteria (AC)
- [ ] `SKILL.md` exists in an `azure-devops/` directory with `Goal`, `Output Structure`, `Quality Rules`, and `Prohibited` sections.
- [ ] `README.md` and `CHANGELOG.md` are present.
- [ ] `references/` directory contains guides for Boards, Repos, Pipelines, and Authentication.
- [ ] The skill includes a "Workflow" section with 4 phases: AUTH, PLANNING, DELIVERY, and OPS.
- [ ] A command comparison table between `az devops` CLI and this skill is included.
- [ ] Helper templates for JSON Patch operations are provided in `examples/`.
- [ ] The skill is registered in the root `README.md`.

## 6. Architecture & Design
The skill will be organized as follows:
- `azure-devops/SKILL.md`: Main entry point.
- `azure-devops/references/`:
    - `authentication.md`: PAT setup, scopes, and environment variables.
    - `boards-management.md`: Work items, WIQL, and Backlogs.
    - `repos-code-flow.md`: PRs, branches, and git commands.
    - `pipelines-cicd.md`: Builds, releases, and environments.
    - `artifacts-and-tests.md`: Feeds and test plans.
- `azure-devops/examples/`:
    - `wiql-queries.md`: Common SELECT statements.
    - `json-patch-payloads.json`: Templates for work item updates.
    - `pipeline-triggers.yml`: Examples for automated triggers.
