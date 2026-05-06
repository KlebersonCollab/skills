---
name: benchmark-expert
version: 2.2.0
description: "Expert Skill for measuring performance baselines, detecting regressions, and comparing stacks."
category: performance
---

# Benchmark Expert — Performance Baseline & Regression Detection

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any benchmark:
0. **Mode Check**: Verify `.hub-mode` and apply `token-distiller` guidelines.
1. **Context Check**: Rehydrate state by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`.
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define the architecture and schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?
6. **Phase 4 Integration**: Benchmarks should preferably be executed in **Phase 4 (Review & Persistence)** of the SDD to validate the implementation impact.
7. **Clean Environment**: Ensure the test environment is isolated and without heavy background processes.

---

## Goal

The goal of this skill is to provide a standardized framework for performance measurement, allowing developers to establish reliable baselines, detect regressions early, and compare technical alternatives based on quantitative data.

---

## Output Structure

The execution of this skill produces the following artifacts:

| Artifact | Location | Description |
|----------|-------------|-----------|
| **Baseline JSON** | `.benchmarks/baseline-*.json` | Raw measurement data for persistence in Git. |
| **Comparison Report** | `artifacts/benchmark-report.md` | Comparative table with deltas and performance verdict. |
| **Performance Logs** | Console | Detailed benchmark execution logs. |

---

## When to Use

- Before and after a PR to measure the performance impact.
- Setting up initial baselines for new projects.
- When users report that the system "is slow."
- Before an official release to ensure targets have been met.
- Comparing different libraries or stacks (e.g., Vite vs. Webpack).

---

## Operating Modes

### Mode 1: Page Performance (Browser)
Measures real browser metrics via automation tools.

**Workflow:**
1. Navigate to target URLs.
2. Measure **Core Web Vitals**:
   - **LCP (Largest Contentful Paint)**: Target < 2.5s.
   - **CLS (Cumulative Layout Shift)**: Target < 0.1.
   - **INP (Interaction to Next Paint)**: Target < 200ms.
   - **FCP (First Contentful Paint)**: Target < 1.8s.
   - **TTFB (Time to First Byte)**: Target < 800ms.
3. Measure resource sizes:
   - Total page weight (Target < 1MB).
   - JS Bundle (Target < 200KB gzipped).
   - CSS, Images, and third-party scripts.

### Mode 2: API Performance
Benchmarks API endpoints for latency and robustness.

**Workflow:**
1. Run 100+ requests on the target endpoint.
2. Measure percentiles: **p50, p95, p99**.
3. Monitor status codes and response sizes.
4. Test under concurrent load (e.g., 10-50 simultaneous connections).

### Mode 3: Build & Dev Loop Performance
Measures the agility of the development ecosystem.

**Metrics:**
1. **Cold Build**: Build time from scratch.
2. **Hot Reload (HMR)**: Time between saving a file and reflecting in the UI.
3. **Test Duration**: Total time for the test suite.
4. **Tooling Latency**: Time for Lint and Type Check.

### Mode 4: Before/After Comparison (Regression)
The heart of the skill for PR validation.

```bash
# Save current state as baseline
/benchmark-expert baseline

# ... apply code changes ...

# Compare current state against saved baseline
/benchmark-expert compare
```

---

## Quality Rules

- **Deterministic Measurements**: Always run the benchmark 3 times and use the average.
- **Context Awareness**: Document the specifications of the machine where the benchmark was run.
- **Fail Fast**: If a critical metric (LCP) delta increases by > 15%, the verdict must be a mandatory **FAIL**.
- **Git Persistence**: Critical baselines must be committed to the `.benchmarks/` directory.

---

## Prohibited

- **NEVER** ignore performance regressions in PRs without documented technical justification.
- **NEVER** run benchmarks on machines with high CPU/Memory load (instability).
- **NEVER** use benchmark tools that do not measure p99 for APIs.

---

## References
- [Performance Targets](./references/performance_targets.md)
- [Regression Analysis](./references/regression_analysis.md)


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.358445Z"
evidence_checksum: "NONE"
```
