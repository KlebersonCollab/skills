---
name: harness-science
version: 1.0.0
description: "Expert in Agentic Behavior Science and Ablation Testing. Focuses on measuring the contribution of specific agentic capabilities to overall success rates using deterministic baselines."
category: intelligence-governance
---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Did you rehydrate the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Benchmark Link**: This skill requires integration with `benchmark-expert` for quantitative data.
3. **Spec Check**: Does the experiment have a clear hypothesis in `spec.md`?

---

# Harness Science (Ablation & Behavior Engineering)

> "To understand the intelligence of an agent, you must see how it fails when its tools are stripped away." — Protocol for rigorous agentic evaluation.

---

## 1. Goal
The goal of this skill is to transform agent development from "vibe-based" adjustments into a rigorous scientific discipline. It uses **Ablation Testing** to determine the exact value of every feature, tool, and reasoning step in the system.

---

## 2. Core Methodology: Ablation Testing
Ablation is the systematic removal of components to observe the impact on performance.

### 🚩 Standard Ablation Flags
When conducting experiments, use these flags to enforce behavior:
- **`CLAUDE_CODE_DISABLE_THINKING`**: Disables internal reasoning chains (CoT). Measures if "thinking" actually improves the outcome or just increases latency.
- **`DISABLE_BACKGROUND_TASKS`**: Forces sequential execution. Measures the risk/benefit of parallel agentic actions.
- **`CLAUDE_CODE_SIMPLE`**: The absolute baseline. Minimal toolset, minimal prompt. Used to detect regression in core logic.
- **`DISABLE_MEMORY_RETRIEVAL`**: Disables access to `MEMORY.md`. Measures the importance of institutional memory vs. zero-shot capabilities.

---

## 3. The Scientific Cycle
1. **Baseline (Harness L0)**: Measure performance in the simplest possible configuration.
2. **Feature Injection**: Add the new capability (e.g., a new skill or tool).
3. **Full Run**: Measure performance with all features active.
4. **Counter-factual Ablation**: Disable *only* the new feature and re-measure.
5. **Delta Analysis**: If `Score(Full) - Score(Ablated) <= 0`, the feature is technical debt and should be removed.

---

## 4. Operational Baselines (Ground Truth)
Maintain a set of "Ground Truth" tasks that the agent must *always* succeed at in `SIMPLE_MODE`. These are your regression shields.

- **Baseline A (Logic)**: Basic code modification without external tools.
- **Baseline B (Navigation)**: Finding information in a multi-file project.
- **Baseline C (Compliance)**: Adherence to global mandates under constraints.

---

## 5. Metrics of Success
- **Success Rate (SR)**: Did the task reach completion according to ACs?
- **Token Efficiency (TE)**: How many tokens were consumed per unit of success?
- **Reasoning Depth (RD)**: Complexity of the path taken (measured via trace length).
- **Ablation Impact (AI)**: The percentage of SR lost when the feature is removed.

---

## 6. Prohibited
- **NEVER** promote a feature to "Production" without an ablation report.
- **NEVER** ignore "Negative Gain" (when a feature makes the agent slower or less accurate).
- **NEVER** use anecdotal evidence ("it feels smarter") as a substitute for benchmarking.
- **NEVER** skip the baseline measurement; without a floor, you cannot measure the ceiling.

---

## 7. References
- [Ablation Studies in Deep Learning](https://towardsdatascience.com/ablation-studies-in-deep-learning-70337c76899b)
- [Harness Science L0 Ablation Baseline - src/entrypoints/cli.tsx]
