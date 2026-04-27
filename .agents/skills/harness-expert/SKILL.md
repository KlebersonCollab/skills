---
name: harness-expert
version: 1.1.0
description: "Expert in System Governance, Agentic Behavior Science, and Auto-Diagnostics. Focuses on Hub integrity (Doctor), session determinism (VCR), and ablation testing."
category: intelligence-governance
---

# Harness Expert: Governance & Resilience

> "A resilient Hub is one that can diagnose its own failures and reproduce its successes."

---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Did you rehydrate the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Integrity Check**: Run `hb doctor` if the environment seems unstable.
3. **Benchmark Link**: This skill requires integration with `benchmark-expert` for quantitative data.

---

## 1. System Doctor (`hb doctor`)
The "Doctor" protocol is the auto-diagnostic engine of the Hub. Use it to verify if the workspace is ready for an agent.

### 🩺 Diagnostic Checklist
When performing a diagnosis, verify:
- **Environment**: Path validity, shell permissions, and key environment variables (e.g., `CLAUDE_CODE_OUTPUT_TOKENS`).
- **Skill Integrity**: Ensure all `.agents/skills` have valid `SKILL.md` files and no broken references.
- **SDD Compliance**: Check for orphaned `spec.md` or `plan.md` files without corresponding tasks.
- **Registry**: Validate that `skills-catalog.md` and `onboarding-navigator` are up to date.

---

## 2. Session VCR (Recording & Replay)
The VCR protocol ensures **determinism**. Use it to record "Golden Paths" and debug divergent agent behavior.

### 📼 Recording Mode
- **Action**: Use a specific flag or logging level to capture all tool inputs/outputs and reasoning chains.
- **Storage**: Save the session trace in `.specs/history/vcr/<session-id>.json`.
- **Purpose**: Create a reproducible baseline for complex tasks.

### ⏪ Replay & Debug
- **Rewind**: If an agent fails a task that previously succeeded, load the VCR trace.
- **A/B Comparison**: Identify the exact turn where the agent's logic diverged from the golden path.

---

## 3. Ablation Testing (Scientific Core)
Ablation is the systematic removal of components to observe the impact on performance.

### 🚩 Standard Ablation Flags
- **`CLAUDE_CODE_DISABLE_THINKING`**: Measures if "thinking" improves the outcome.
- **`DISABLE_MEMORY_RETRIEVAL`**: Measures the importance of `MEMORY.md`.
- **`DISABLE_BACKGROUND_TASKS`**: Measures the benefit of parallel execution.

---

## 4. The Scientific Cycle
1. **Baseline**: Performance in the simplest configuration.
2. **Feature Injection**: Add the new capability.
3. **Full Run**: Performance with all features active.
4. **Counter-factual Ablation**: Disable only the new feature and re-measure.

---

## 5. Prohibited
- **NEVER** ignore a failing `hb doctor` report.
- **NEVER** promote a feature to "Production" without an ablation report.
- **NEVER** skip the baseline measurement.
- **NEVER** assume behavior is deterministic without VCR verification in complex scenarios.

---

## 6. References
- [Ablation Studies in Deep Learning](https://towardsdatascience.com/ablation-studies-in-deep-learning-70337c76899b)
- [Doctor Checklist](resource:resources/DOCTOR_CHECKLIST.md)
- [VCR Protocol](resource:resources/VCR_PROTOCOL.md)
