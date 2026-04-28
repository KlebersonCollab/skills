# Spec: CLI Flexbox Layout (Yoga-inspired)

## 1. Vision & Goals
Implement a high-performance, responsive layout engine for the `hb` CLI based on Flexbox principles (inspired by the Yoga Layout used in the legacy `src` project). This will enable complex, adaptive terminal UIs with minimal manual coordinate math.

## 2. Functional Requirements
- **FR-1: Flexbox Container**: Ability to define containers with direction (Row/Column).
- **FR-2: Box Model**: Support for margin, padding, and borders.
- **FR-3: Alignment**: Support for `justify-content` and `align-items`.
- **FR-4: Responsiveness**: Automatic recalculation of layout when terminal size changes.
- **FR-5: Integration**: Seamless integration with the existing `hb/internal/ui/manager.go`.

## 3. Acceptance Criteria (BDD)

### Scenario: Creating a simple responsive header
**Given** a terminal width of 100 characters
**When** I define a Row container with two children (Title and Version)
**And** I set `justify-content` to `space-between`
**Then** the Title should be at the far left and the Version at the far right.

### Scenario: Vertical centering in a HUD
**Given** a fixed-height HUD at the bottom of the terminal
**When** I add a text element with `align-items: center`
**Then** the text should be vertically centered within the HUD.

### Scenario: Resizing the terminal
**Given** an active layout
**When** the terminal window is resized from 100 to 80 columns
**Then** the layout engine should automatically recalculate positions without visual glitches.

## 4. Constraints
- Must be implemented in Go.
- Should avoid heavy external dependencies if possible (or use a lightweight Go port of Yoga).
- Must maintain performance to avoid flickering in the terminal.
