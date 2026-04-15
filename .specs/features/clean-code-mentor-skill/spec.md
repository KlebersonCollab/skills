# Spec: Clean Code Mentor Skill

## 1. Goal
Create a specialized skill that acts as a technical mentor and code reviewer, focusing 100% on **SOLID**, **YAGNI**, **DRY**, and **KISS** principles. This skill will help maintain the highest code quality across all projects.

## 2. Target Audience
Developers using Gemini CLI who want rigorous feedback on their code design and architecture, aiming for maintainability and simplicity.

## 3. Core Requirements (Functional)
- **SOLID Audit**: Analyze classes and modules for violations of the Single Responsibility, Open/Closed, Liskov Substitution, Interface Segregation, and Dependency Inversion principles.
- **YAGNI & KISS Guard**: Identify "over-engineering" and unnecessary complexity, suggesting simpler alternatives.
- **DRY Verification**: Detect logic duplication across the codebase (when multiple files are provided).
- **Refactoring Suggestions**: Provide actionable code snippets to improve design without changing behavior.
- **Technical Mentorship**: Explain *why* a certain pattern is being suggested, citing classic principles.

## 4. Acceptance Criteria (BDD Format)

### Scenario 1: Identify Single Responsibility Violation
**Given** a class with multiple unrelated responsibilities (e.g., UI logic + Database access)
**When** the `clean-code-mentor` analyzes the code
**Then** it must flag the violation and suggest splitting the class into specialized modules.

### Scenario 2: Prevent Over-engineering (YAGNI)
**Given** a feature implementation with complex abstractions for future "possible" use cases
**When** the `clean-code-mentor` reviews the PR
**Then** it must identify the YAGNI violation and recommend a simpler, direct implementation.

### Scenario 3: DRY Audit across files
**Given** two or more files with identical logic blocks
**When** the `clean-code-mentor` is invoked on the directory
**Then** it must suggest a shared abstraction or utility to centralize the logic.

## 5. Technical Requirements (Non-Functional)
- **Language Agnostic**: Must work with Dart (Flutter), Python, TypeScript, and Go.
- **Tone**: Professional, encouraging, and strictly technical.
- **Output Format**: Structured as a "Review Report" with Severity levels (Critical, Warning, Info).

## 6. Architecture & Design
- `clean-code-mentor/SKILL.md`: Entry point with the review workflow.
- `clean-code-mentor/references/`:
    - `solid-principles.md`: Deep dive into each SOLID letter with examples.
    - `yagni-and-kiss.md`: Simplicity as the ultimate sophistication.
    - `dry-and-clean-tests.md`: Principles for logic and test quality.
- `clean-code-mentor/examples/`:
    - `bad-vs-good-code.md`: Comparison samples for training and reference.
