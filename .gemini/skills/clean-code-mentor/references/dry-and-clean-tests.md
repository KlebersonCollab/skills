# DRY & Clean Tests: Quality Assurance

Principles to ensure the cohesion of the codebase and the reliability of the test suite.

---

## 1. DRY (Don't Repeat Yourself)
**"Every piece of knowledge must have a single, unambiguous, authoritative representation within a system."**

- **What is NOT DRY**: Logic duplication (e.g., calculating the same tax in two different places).
- **What IS DRY**: Centralizing the calculation in a single function or service that is reused by all.
- **Violation Signals**:
    - Changes in one place require changes in multiple other places ("Shotgun Surgery").
    - Identical code patterns in unrelated classes.
- **Strategy**: Abstract duplicated logic into utilities, mixins, or specialized services.
- **Note**: Do not confuse *code* duplication (e.g., two similar `for` loops) with *knowledge* duplication (e.g., the same business rule in two places). Forcing DRY on logics that evolve differently can lead to unwanted coupling.

## 2. Clean Tests (FIRST)
Tests should be treated with the same quality rigor as production code.

- **F: Fast**: Tests should run in seconds. If they are slow, nobody will run them.
- **I: Independent**: One test should not depend on the result of another test or the state left by it.
- **R: Repeatable**: They must pass in any environment (Local, CI, Docker), at any time.
- **S: Self-validating**: The test must have a clear boolean result (Pass/Fail). No logs to check manually.
- **T: Timely**: They should preferably be written in the TDD cycle (or along with the production code).

---

## Clean Test Structure (AAA)

Always follow the **Arrange, Act, Assert** pattern for readability:

1.  **Arrange**: Set up the scenario, create objects and mocks.
2.  **Act**: Execute the method or action being tested.
3.  **Assert**: Verify that the result was as expected.

**Tip**: A good test should tell a story. Use descriptive names (e.g., `test_should_throw_exception_when_user_is_null`).


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.380746Z"
evidence_checksum: "NONE"
```
