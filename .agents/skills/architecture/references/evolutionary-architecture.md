# Evolutionary Architecture & Fitness Functions

This guide describes how to design systems that support guided and incremental change across multiple dimensions.

---

## 1. What is Evolutionary Architecture?

It is an architecture that supports guided and incremental change as a first-class principle. Instead of trying to predict the entire future (YAGNI), it focuses on protecting the system's vital characteristics as it evolves.

## 2. Fitness Functions

A **Fitness Function** is any mechanism that provides an objective evaluation of an architectural integrity characteristic. They can be automated (unit tests, scripts) or manual (code reviews).

### Examples of Protected Dimensions:
- **Coupling**: Ensuring that module A does not depend on module B.
- **Security**: Verifying that no endpoint is exposed without authentication.
- **Performance**: Validating that the average response time is below 200ms.
- **Quality**: Preventing cyclomatic complexity above a limit X.

## 3. Types of Fitness Functions

1.  **Atomic**: Executed in a single context (e.g., a unit test verifying package dependencies).
2.  **Holistic**: Evaluates the system as a whole (e.g., load test to validate scalability).
3.  **Continuous**: Executed in the CI/CD pipeline on every commit.
4.  **Temporal**: Triggered by time events (e.g., checking certificate expiration).

## 4. Practical Implementation (Architecture as Code)

Use **Architectural Linter** tools or simple scripts in CI:

### Example: Validating Coupling (Python/Import Linter)
```toml
# setup.cfg or import-linter.config
[importlinter]
root_package = my_project

[[importlinter:contracts:forbidden]]
name = Domain Layer should not depend on Infra Layer
source_modules = my_project.domain
forbidden_modules = my_project.infra
```

### Example: Validating Cycles (Dart/ArchTest)
```dart
void main() {
  test('No circular dependencies between features', () {
    final arch = ArchTest.fromDirectory('lib/features');
    expect(arch, hasNoCircularDependencies());
  });
}
```

## 5. Best Practices

- **Fast Feedback**: Prefer atomic and continuous functions that run in seconds.
- **Fail Fast**: Fail the build if a Fitness Function is violated.
- **Test Evolution**: As the architecture changes, Fitness Functions must also be updated or removed.
