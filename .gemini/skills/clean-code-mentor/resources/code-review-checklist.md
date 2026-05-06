# Code Review Checklist

This checklist serves as a guide for code reviews focused on quality and maintainability.

## 1. SOLID Principles
- [ ] **S (Single Responsibility)**: Does each class/function have only one reason to change?
- [ ] **O (Open/Closed)**: Is the code open for extension but closed for modification?
- [ ] **L (Liskov Substitution)**: Can subclasses be substituted for their base classes without breaking the system?
- [ ] **I (Interface Segregation)**: Are interfaces specific and don't force unnecessary implementations?
- [ ] **D (Dependency Inversion)**: Does the code depend on abstractions rather than concrete implementations?

## 2. Engineering Mandates
- [ ] **YAGNI (You Ain't Gonna Need It)**: Is there code that tries to predict the future or solves problems that don't exist yet? (Remove if so)
- [ ] **DRY (Don't Repeat Yourself)**: Is there duplicated logic that should be abstracted in a single place?
- [ ] **KISS (Keep It Simple, Stupid)**: Is the solution the simplest possible for the problem? Is there any unnecessary complexity?

## 3. Readability and Style
- [ ] **Names**: Do classes, methods, and variables have clear names that reveal their intent?
- [ ] **Size**: Are methods and classes small enough to be understood quickly?
- [ ] **Comments**: Is the code self-explanatory? Are there comments that explain "why" instead of "what"?
- [ ] **Error Handling**: Are errors handled appropriately? Is there correct use of logs/exceptions?
- [ ] **Tests**: Are changes covered by automated tests (unit/integration)?

## 4. Security and Performance
- [ ] **Security**: Is there any risk of injection (SQL, XSS, etc)? Is any sensitive data being improperly exposed?
- [ ] **Performance**: Is there any obvious bottleneck (N+1 queries, heavy nested loops, etc)?
