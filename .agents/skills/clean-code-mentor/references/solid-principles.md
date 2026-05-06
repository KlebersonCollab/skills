# SOLID Principles: Deep Dive

The five fundamental principles of object-oriented design that ensure software maintainability and scalability.

---

## 1. S: Single Responsibility Principle (SRP)
**"A class should have only one reason to change."**

- **Violation**: "God" classes that handle UI logic, data access, and business rules.
- **Solution**: Decompose classes into smaller, specialized components.
- **Check**: If you have difficulty naming the class concisely or if it has multiple "Ands" in its description, it probably breaks SRP.

## 2. O: Open/Closed Principle (OCP)
**"Software entities should be open for extension, but closed for modification."**

- **Violation**: Excessive use of `switch` or `if-else` based on types, forcing the alteration of the original logic for each new type added.
- **Solution**: Use interfaces, polymorphism, and patterns like *Strategy* or *Decorator*.
- **Check**: When adding a new feature, do I need to touch code that is already working? If so, you broke OCP.

## 3. L: Liskov Substitution Principle (LSP)
**"Subclasses should be substitutable for their superclasses without breaking the program's behavior."**

- **Violation**: Overriding a method to throw a "Not implemented" exception or forcing type checks (`is instance of`).
- **Solution**: Ensure that subclasses respect the contract (pre-conditions and post-conditions) defined by the superclass.
- **Check**: If I replace the base class with a child class anywhere, does the system remain stable?

## 4. I: Interface Segregation Principle (ISP)
**"Many specific interfaces are better than one single, general interface."**

- **Violation**: A "Fat" interface that forces a class to implement methods it will not use.
- **Solution**: Segregate interfaces into smaller, cohesive functional units.
- **Check**: Is my class implementing empty methods just to satisfy the interface?

## 5. D: Dependency Inversion Principle (DIP)
**"Depend on abstractions, not on concrete implementations."**

- **Violation**: Instantiating dependencies directly inside a class (e.g., `var repo = new SqlDatabase()`).
- **Solution**: Dependency injection and interface-oriented programming.
- **Check**: Can I swap my database or external service without needing to recompile business classes?

---

## Best Practices

- **Don't overdo it**: Blindly applying all principles to small projects can lead to over-engineering.
- **Context**: SOLID is a tool to reduce the cost of change; use it where change is likely.
- **Simplicity**: If SOLID makes the code impossible to read, review your KISS implementation.


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.380877Z"
evidence_checksum: "NONE"
```
