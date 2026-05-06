# Changelog — API Architect

All notable changes to this skill will be documented in this file.

## [1.3.0] - 2026-04-14

### Changed
- **Workflow & References**: Improved workflow to explicitly reference the use of `references/api-style.md` during the design phase.

## [1.2.0] - 2026-04-14

### Changed
- **Renaming**: Skill renamed from `api-designer` to `api-architect` to reflect a larger scope of governance and resilience.
- **Style Guide**: Added `references/api-style.md` with standards for response envelopes, errors, and pagination.
- **Resilience**: Included Timeout, Retry, and Circuit Breaker patterns.

## [1.1.0] - 2026-04-14

### Added
- **tRPC Support**: Full reference in `references/trpc-patterns.md` for fullstack TypeScript projects.
- **Security Guide (OWASP API Top 10)**: New `references/api-security-guide.md` file focusing on BOLA, Broken Auth, Mass Assignment, etc.
- **Rate Limiting**: New guide in `references/api-rate-limiting.md` covering algorithms (Token Bucket) and header standardization (429 status).
- **Decision Tree**: Integrated into `SKILL.md` to facilitate choosing between REST, GraphQL, tRPC, and gRPC.
- **Advanced Pagination Patterns**: Support for Cursor-based pagination (Relay style) and response envelopes.

---

## [1.0.0] - 2026-04-01

### Added
- Initial version of the skill with support for REST and GraphQL.
- 4-phase workflow: Discover, Model, Specify, and Validate.
