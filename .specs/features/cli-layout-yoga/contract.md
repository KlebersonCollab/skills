# Contract: CLI Flexbox Layout (Yoga-inspired)

## Parties
- **Implementer**: Antigravity (AI)
- **Reviewer**: User / Audit System

## Delivery
A working Go package `hb/internal/ui/layout` that correctly solves a Flexbox tree for terminal character-based coordinates.

## Sensors (Validation)
1. **Unit Tests**: A suite of tests in `hb/internal/ui/layout/layout_test.go` covering Row, Column, and SpaceBetween scenarios.
2. **Visual Audit**: A demonstration command `hb ui test-layout` that shows a responsive header and footer.
3. **Performance**: Recalculation of a 20-node tree must take less than 1ms on a standard laptop.

## Definition of Done
- Code is 100% Go (no CGO unless strictly necessary for performance).
- Follows the SDD framework (Phases 1-4).
- No flickering during terminal resizing.
