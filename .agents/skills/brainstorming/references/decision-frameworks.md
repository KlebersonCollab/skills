# Decision Frameworks — Complete Reference

Patterns for decision-making and documenting technical choices.

---

## 1. Trade-off Matrix
Every technical choice has advantages and sacrifices. This matrix helps make these sacrifices explicit.

| Attribute | Decision A | Decision B | Trade-off |
|----------|-----------|-----------|-----------|
| Dev Speed | High | Low | A is faster now, but B is better in the long run. |
| Scalability | Low | High | B supports 10x more load than A. |
| Operational Cost | Low | Medium | B requires extra infra maintenance. |

## 2. Decision Log
Document the reasoning so that other developers (or you in the future) understand why a choice was made.

### Template:
- **Date**: 2026-04-14
- **Status**: Decided / In Discussion
- **Context**: What was the problem?
- **Options Considered**: Option 1, Option 2.
- **Option Chosen**: Option 1.
- **Rationale**: Why did Option 1 win? (e.g., Lower technical risk).

## 3. Understanding Lock (Checklist)
Before moving from brainstorming to specification, ensure these points have been validated by the user:

- [ ] **Goals**: What do we want to achieve?
- [ ] **Non-goals**: What will we **not** do?
- [ ] **Constraints**: Time, budget, technology stack.
- [ ] **Assumptions**: What are we assuming to be true?
- [ ] **Risks**: What could go wrong?
