# Project Learnings: Skills Hub

## Technical Patterns
- **SDD v2.2.0 Integration**: Observed that modularizing skills into independent folders with their own `SKILL.md` allows for seamless agentic consumption.

## Bug Fixes & Gotchas
- [TBD] Initializing learning log.

## Process Insights
- **Manual Bootstrapping**: Discovered that manual creation of `.specs/` structure ensures the agent fully understands the governance layer before implementation.
- **Purist Transition**: Removing CLI dependencies (hb cli) simplifies the environment and reduces "magic" failures, forcing the agent to rely on explicit Markdown contracts which are more resilient and traceable.

