# Conventional Commits Guide

The **Conventional Commits** standard is a lightweight convention for commit messages. **Project Mandate**: All messages must be written in **English**.

## Message Structure

```
<type>(<scope>): <subject>

[optional body]

[optional footer]
```

### Types (Mandatory)

| Type | Description |
|------|-----------|
| **feat** | A new feature for the user, not a new feature for the build script. |
| **fix** | A bug fix for the user, not a fix for the build script. |
| **docs** | Documentation changes. |
| **style** | Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc). |
| **refactor** | A code change that neither fixes a bug nor adds a feature. |
| **perf** | A code change that improves performance. |
| **test** | Adding missing tests or correcting existing tests. |
| **build** | Changes that affect the build system or external dependencies (e.g., gulp, broccoli, npm). |
| **ci** | Changes to our CI configuration files and scripts (e.g., Travis, Circle, BrowserStack, SauceLabs). |
| **chore** | Other changes that don't modify `src` or test files. |
| **revert** | Reverts a previous commit. |

### Scope (Optional)
The scope can be anything specifying the place of the commit change. For example: `auth`, `api`, `ui`, `database`, etc.

### Subject (Mandatory)
The subject contains a succinct description of the change:
- Use the imperative, present tense: "change" not "changed" nor "changes".
- Don't capitalize the first letter.
- No dot (.) at the end.

## Message Examples

### Simple Commit
```bash
feat(api): add oauth2 support for external partners
```

### Commit with Body and Footer
```bash
fix(ui): resolve race condition in login form

Added a loading state to prevent multiple submissions while the request
is in flight.

Closes #456
```

### Breaking Change
```bash
feat(db):! migrate to postgresql

BREAKING CHANGE: The database layer now requires a PostgreSQL connection string.
```


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.363366Z"
evidence_checksum: "NONE"
```
