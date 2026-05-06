# Branching Strategies

Choosing the right strategy depends on team size, release cadence, and automation maturity.

## 1. GitHub Flow
**Best for**: SaaS, web applications, small teams, and continuous delivery.

### Workflow:
1.  **main Branch**: Represents the production state. Must always be stable.
2.  **Feature Branches**: Any change (new feature or fix) must be made in a branch created from `main`.
    - Naming: `feature/<desc>` or `fix/<desc>`.
3.  **Pull Request (PR)**: When the branch is ready, a PR is opened for review.
4.  **Discussion and Review**: Team feedback and adjustments.
5.  **Merge**: After approval and CI pass, the branch is integrated into `main`.
6.  **Deploy**: Deployment is done immediately after the merge.

---

## 2. Trunk-Based Development
**Best for**: Senior teams, high development velocity, microservices architecture.

### Workflow:
- Developers work on a single main branch (`trunk` or `main`).
- Branches are very short-lived (hours or 1-2 days).
- **Feature Flags**: Incomplete functionalities are sent to production but hidden behind configuration keys.
- **Continuous Integration**: Code is integrated several times a day.

---

## 3. GitFlow
**Best for**: Legacy projects, regulated environments, scheduled releases (e.g., Mobile Apps in stores).

### Workflow:
- **main**: Code in production.
- **develop**: Integration branch for the next release.
- **feature/**: Created from `develop`, integrated back into `develop`.
- **release/**: Created from `develop` for final release preparation. Merged into `main` and `develop`.
- **hotfix/**: Created from `main` for immediate critical fixes. Merged into `main` and `develop`.

---

## Comparison Table

| Criterion | GitHub Flow | Trunk-Based | GitFlow |
|----------|-------------|-------------|---------|
| **Complexity** | Low | Medium | High |
| **Velocity** | High | Very High | Moderate |
| **Conflict Risk** | Low | Moderate | High |
| **Release Control** | Continuous | Continuous | Scheduled |
