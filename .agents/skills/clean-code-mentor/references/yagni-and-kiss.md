# YAGNI & KISS: Design Simplicity

Simplicity is the ultimate sophistication in software development. These principles help avoid the trap of over-engineering.

---

## 1. YAGNI (You Ain't Gonna Need It)
**"Do not implement something until it is truly necessary."**

Most development time is wasted creating abstractions for future scenarios that will never occur.

- **Violation Signals**:
    - "Just in case" abstractions (e.g., interfaces with only one concrete implementation).
    - Commented-out or inactive code kept for "future use."
    - Plugins and libraries added "just to be sure."
- **Strategy**: Solve today's problem today. If tomorrow's problem arrives, today's simple code will be easy to refactor.
- **Check**: If I remove this code now, does the system still meet the current requirements? If so, it's YAGNI.

## 2. KISS (Keep It Simple, Stupid)
**"Keep things simple and direct."**

Complexity is a cost. The simpler the code, the easier it is to test, maintain, and understand.

- **Violation Signals**:
    - Use of complex design patterns for trivial problems.
    - Deep nested logic (multiple `if`, `for`, `while`).
    - Enigmatic variable names or excessive abbreviations.
- **Strategy**: If you can explain the solution to a beginner, it's probably simple. Avoid "clever code."
- **Check**: Six months from now, will I or my colleague understand what this code block does in 30 seconds?

---

## The Simplicity Balance

1.  **Simple Code != Easy Code**: Simple code requires deep thought and effort to build.
2.  **Abstractions vs. Simplicity**: An abstraction should simplify the component's use, not complicate its understanding.
3.  **Continuous Refactoring**: The path to simplicity often involves multiple iterations of cleaning and code reduction.
4.  **Less is More**: The fastest and safest code is the code that was never written.


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.380987Z"
evidence_checksum: "NONE"
```
