# Pattern Selection — Decision Guide

How to choose the correct architectural style for your context.

---

## 1. Modular Monolith vs Microservices

| Characteristic | Modular Monolith | Microservices |
|----------------|------------------|---------------|
| **Infra Cost** | Low | High |
| **Complexity** | Medium | Very High |
| **Deployment** | Atomic (All or nothing) | Independent per service |
| **Recommendation** | Start here. Ideal for 90% of cases. | Use only if you need extreme independent scaling. |

## 2. Hexagonal Architecture (Ports & Adapters)
Focus on decoupling business logic (Core) from the external world (Infra, DB, UI).
- **Advantage**: Extreme ease for swapping DB or running fast unit tests.
- **When to use**: In systems with complex business logic that needs protection.

## 3. Clean Architecture
Evolution of Hexagonal, organizing code into concentric layers.
- **Advantage**: Independence from frameworks and UI.
- **Trade-off**: Can generate excessive files and "boilerplate" for simple systems.

## 4. Event-Driven Architecture (EDA)
Asynchronous communication based on events (Pub/Sub).
- **Advantage**: Temporal decoupling and resilience.
- **Trade-off**: Difficulty in tracking data flows and debugging.

---

## Quick Decision Tree

1. **Is it a new system (Greenfield)?**
   - Yes → Start with a Modular Monolith.
2. **Is the team small (< 10 people)?**
   - Yes → Avoid Microservices.
3. **Is the logic complex?**
   - Yes → Use Hexagonal Architecture for the core.
4. **Needs high write performance?**
   - Yes → Consider CQRS or Event Sourcing.
