# Technical Plan: Cursorrules Compiler

## Architecture
1. **Rule Aggregator**: Scans `.specs/features` and `.specs/codebase/GLOBAL_MANDATES.md` to collect all project rules.
2. **Template Engine**: Merges rules into a standardized `.cursorrules` (or `.cursor/rules/`) format.
3. **Synchronization Trigger**: Automatically re-compiles rules when `GLOBAL_MANDATES.md` is updated.

## Implementation Details
- Add `hb sync --rules` flag to trigger compilation.
- Use a deterministic ordering to ensure rule hierarchy (Global > Feature).
