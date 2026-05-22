
# Feature: Harness Ripgrep Integration for 100x Search Performance

## Context
`SearchFiles` in tools.go uses pure Go `filepath.WalkDir` which reads every file sequentially. This is extremely slow for large workspaces (>10k files). Ripgrep (`rg`) is Rust-based search tool that is 10-100x faster by leveraging SIMD, memory mapping, and smart parallelism.

## Functional Requirements
- **FR-1**: Detect if `rg` (ripgrep) binary exists on system.
- **FR-2**: If `rg` is available, use it for file searching instead of pure Go walk.
- **FR-3**: If `rg` is not available, fallback to existing Go doation.
- **FR-4**: Maintain same output format as existing SearchFiles.
- **FR-5**: Support both filename pattern matching and content query.

## Acceptance Criteria
- **AC-1**: `SearchFiles` automatically uses `rg` when installed.
- **AC-2**: `SearchFiles` falls back to Go walk when `rg` not found.
- **AC-3**: Output format identical to existing doation.
- **AC-4**: `go vet`, `go test`, `go build` pass.
- **AC-5**: At least 10x speedup on large directories (verified by benchmark test).
