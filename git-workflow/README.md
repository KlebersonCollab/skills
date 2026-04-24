# Skill: Git Workflow

## Overview
This skill defines the version control standards and collaborative workflows for the project, ensuring a clean, traceable history integrated into the **SDD (Spec-Driven Development)** cycle.

## What You Can Do
- Choose the best **branching strategy** for your context.
- Write professional **commit messages** using Conventional Commits.
- Perform **rebases** and resolve conflicts safely.
- Open high-quality **Pull Requests** with standardized templates.
- Ensure that Git reflects task progress in SDD.

## Command Quick Guide

### Starting a Feature
```bash
git checkout main
git pull origin main
git checkout -b feature/my-feature
```

### Committing (Conventional Commits + Mandatory English)
> **Note**: Although the documentation is in Portuguese, commit messages must always be in **English**.

```bash
git add .
git commit -m "feat(auth): implement password recovery flow"
```

### Synchronizing with Main (Rebase)
```bash
git fetch origin
git rebase origin/main
# If there are conflicts, resolve them and:
# git add .
# git rebase --continue
```

## Anti-patterns (What NOT to do)
- ❌ Giant commits with many unrelated changes.
- ❌ Commit messages like "fix", "update", "wip".
- ❌ Leaving feature branches alive for weeks.
- ❌ Force pushing (`push -f`) to shared branches without warning.

## References
- [Conventional Commits](references/conventional-commits.md)
- [Branching Strategies](references/branching-strategies.md)
- [Commit Template](examples/gitmessage.template)
- [PR Template](examples/pull-request.md)
