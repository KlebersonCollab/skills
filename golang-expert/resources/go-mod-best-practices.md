# Best Practices: Go Modules

Practical guide for dependency management and versioning in Go projects.

## 1. Core Commands

- `go mod init <module-name>`: Initializes a new module. The name should be the repository URL (e.g., `github.com/user/repo`).
- `go mod tidy`: Removes unused dependencies and adds necessary ones. **Always execute before committing.**
- `go mod vendor`: Creates a local copy of all dependencies in a `vendor/` folder. Useful for offline builds or restricted environments.
- `go list -m all`: Lists all direct and indirect dependencies.

## 2. Versioning and Semantic Versioning (SemVer)

Go Modules follow SemVer (`vX.Y.Z`):
- **Patch (Z)**: Bug fixes.
- **Minor (Y)**: New backward-compatible features.
- **Major (X)**: Changes that break compatibility.

**Note**: Major versions >= 2 must include the suffix in the module name:
```go
module github.com/user/repo/v2
```

## 3. The `go.sum` file
NEVER edit `go.sum` manually. It contains the cryptographic hashes of dependencies to ensure the code downloaded today is identical to the one downloaded tomorrow (Security and Reproducibility).

## 4. Local Replacements (Replace)
To test changes in a dependency locally before publishing it:
```bash
go mod edit -replace github.com/dependency/project=../local/path
```
**Remember**: Remove the `replace` before pushing to the main repository.

## 5. Commit Strategy
- Commit `go.mod` and `go.sum`.
- The use of the `vendor/` folder is optional, but if used, it must be committed entirely.

## Pro Tips
1. **Keep dependencies updated**: Use `go get -u ./...` to update dependencies (with caution).
2. **Avoid Unnecessary Dependencies**: Go values simplicity. Sometimes, 10 lines of own code are better than a heavy external dependency.
3. **Internal Packages**: Use the `internal/` folder for packages that should not be exported to other modules.
