---
name: flutter-fvm
version: 2.2.0
description: "Skill for professional Flutter development with FVM. Includes Dart 3+ patterns, optimized performance, mandatory accessibility, OWASP security, and resilience management (Error Handling)."
category: development-workflow
---

# Flutter with FVM & Modern Patterns

> Flutter version manager combined with modern architecture, performance, and accessibility patterns for high-fidelity applications.

---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify `.hub-mode` and apply `token-distiller` guidelines.
1. **Context Check**: Rehydrate state by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`.
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define the architecture and schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---
## Goal

**FVM** ensures the environment. This skill ensures **Technical Excellence** through:

| Pillar | Main Focus |
|-------|----------------|
| **Dart 3+** | Sealed classes, Pattern Matching, and Records for type safety. |
| **Performance** | Rebuild optimization (`MediaQuery.sizeOf`), RepaintBoundaries, and Decomposition. |
| **UX/Accessibility** | Semantics, touch targets (48x48), and dynamic contrast. |
| **Resilience** | Global error capture and graceful failure handling. |
| **Security** | OWASP Mobile Top 10 and secure Secrets management (`--dart-define`). |

## Version Awareness

**Recommended Version: Flutter Stable (latest available via FVM)**

Check the installed and active project versions:

```bash
fvm list
fvm current
```

## When to Use This Skill

Use this skill when:
- Initializing a new Flutter project or migrating to Dart 3+.
- Implementing complex business logic with State Management (Riverpod/Bloc).
- Optimizing screen performance with complex lists or animations.
- Ensuring the app is accessible via screen readers (TalkBack/VoiceOver).
- Setting up robust error handling and crash reporting.
- Managing secrets and API keys securely.
- Performing OWASP Mobile Top 10 security analysis.

Skip this skill when:
- The project does NOT use Flutter/Dart.

---

## Workflow (4 Phases)

### Phase 1: ENVIRONMENT — Configuration and Alignment
1.  **SDK Version**: Set with `fvm use stable`.
2.  **Rigorous Linting**: Configure `analysis_options.yaml` with `strict-casts: true`, `strict-inference: true`, and `strict-raw-types: true`.

### Phase 2: PROJECT — Structure and Modernization
1.  **State Modeling**: Use `sealed class` to represent states (Initial, Loading, Data, Error).
2.  **Secrets Management**: Create structure to receive keys via `--dart-define` or `.env` (excluded from git).
3.  **Error Boundaries**: Configure `FlutterError.onError` in `main.dart`.

### Phase 3: DEVELOP — Quality and Performance
1.  **Decomposition**: Keep `build()` methods under 100 lines.
2.  **Selective Rebuilds**: Prefer `MediaQuery.sizeOf(context)` over `MediaQuery.of(context)`.
3.  **Accessibility**: Add `semanticLabel` to all critical visual elements.
4.  **Testing**: Cover sealed state transitions and error flows.

### Phase 4: DEPLOY — Security and Resilience
1.  **Crash Reporting**: Ensure integration with Sentry/Crashlytics.
2.  **Hardening**: Run builds with `--obfuscate --split-debug-info`.
3.  **Validation**: Verify that no `print` statements remain (use `debugPrint` or `log`).

## Output Structure

The execution of this skill results in the generation or update of the following artifacts:

| Artifact | Location | Description |
|----------|-------------|-----------|
| **FVM Config** | `.fvm/fvm_config.json` | Definition of the SDK version for the team. |
| **Analysis Options** | `analysis_options.yaml` | Linting and static analysis rules. |
| **Sealed States** | `lib/features/*/domain/` | Exhaustive state modeling. |
| **Secure Secrets** | `.env` / `main.dart` | Key management via environment. |

---

## Quality Rules

**DO:**
- Use `sealed` classes to ensure exhaustiveness in state handling.
- Use `const` constructors aggressively to reduce GC pressure.
- Ensure touch targets are at least 48x48 dp.
- Use `RepaintBoundary` on animations or heavy drawing widgets.
- Validate `context.mounted` after any `await`.

**DON'T:**
- Do not use multiple booleans (`isLoading`, `hasError`) for screen states.
- Do not use `print()` in production; use the `logging` package or `dart:developer`.
- Do not ignore errors from `fvm flutter analyze`.
- Do not store secrets directly in Dart code (hardcoded).

---

## Reference Documentation

1. **[Modern Dart Patterns](references/modern-dart-patterns.md)** — Sealed classes, Records, and Pattern Matching.
2. **[Performance & Optimization](references/performance-and-optimization.md)** — Rebuilds, RepaintBoundaries, and GC.
3. **[Accessibility Guide](references/accessibility-guide.md)** — Semantics, Hit Targets, and Inclusive UX.
4. **[Environment Setup](references/environment-setup.md)** — FVM and Strict Analyzer.
5. **[Flutter Security Guide](references/flutter-security-guide.md)** — OWASP and Secrets.
6. **[Testing & Quality](references/testing-and-quality.md)** — State and Logic testing.

---

## Prohibited

- **Do Not Skip Accessibility**: Features without semantic labels or small targets are considered incomplete.
- **Do Not Ignore the Analyzer**: Analysis errors prevent deployment.
- **Do Not Use Global Flutter**: Maintain project isolation via FVM.

---
*Updated on April 22, 2026, to align with ECC standards.*
