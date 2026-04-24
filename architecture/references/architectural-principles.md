# Architectural Principles — Complete Reference

Fundamental principles for creating sustainable and robust systems.

---

## 1. Simplicity (The Core Principle)
> "Simplicity is the ultimate sophistication." (Leonardo da Vinci)

- Start with the minimum necessary to solve the problem.
- Add complexity only when there is evidence of need (scale, performance, security).
- Remember: It is much cheaper to add complexity later than to remove unnecessary complexity now.

## 2. DRY (Don't Repeat Yourself)
- Avoid duplicating business logic or knowledge in the system.
- Every piece of knowledge must have a single, unambiguous representation within a system.

## 3. KISS (Keep It Simple, Stupid)
- Most systems work better if they are kept simple rather than complicated.
- Avoid premature abstractions.

## 4. YAGNI (You Ain't Gonna Need It)
- Do not implement features or infrastructure until they are truly necessary.
- Applying YAGNI reduces architectural "noise" and maintenance costs.

## 5. SOLID
- **S**: Single Responsibility Principle (A class should have a single reason to change).
- **O**: Open/Closed Principle (Entities open for extension, closed for modification).
- **L**: Liskov Substitution Principle (Subtypes must be replaceable by their base types).
- **I**: Interface Segregation Principle (Many specific interfaces are better than one general one).
- **D**: Dependency Inversion Principle (Depend on abstractions, not implementations).
