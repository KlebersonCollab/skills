# Specification: HB-CLI Premium UI

## Goal
Implement a high-performance, ANSI-based terminal UI for the `hb` CLI. The UI should provide a persistent status bar at the bottom of the terminal to display real-time SDD progress, task status, and system metrics without interfering with standard command logs.

## Background
The current `hb` CLI is text-only, which makes it difficult for users to track long-running agent tasks or overall project health at a glance. By cannibalizing patterns from the `src/bridge/bridgeUI.ts` (the "other project"), we will implement a "zero-flicker" status bar.

## Requirements
- **FR-1: Persistent Status Bar**: Implement a footer that stays at the bottom of the terminal window.
- **FR-2: SDD Progress Tracking**: Display current feature, active task, and % of completion from `tasks.md`.
- **FR-3: Metric Visualization**: Display token usage and session cost.
- **FR-4: ANSI Control**: Use raw ANSI escape codes for cursor management to ensure "auto-injectable" behavior across different terminals (iTerm2, Tmux, standard bash).
- **FR-5: Terminal Resize Handling**: UI must adapt or re-render correctly when the terminal window is resized.
- **FR-6: No-Pollute Policy**: The status bar must not pollute the standard scrollback buffer (erased on exit).

## Acceptance Criteria (BDD)
### Scenario 1: Displaying the Status Bar
**Given** the user runs an `hb` command that supports UI (e.g., `hb status --watch`)
**When** the command starts
**Then** a colored status bar appears at the bottom of the terminal
**And** it displays the current project name and SDD status.

### Scenario 2: Updating Progress
**Given** a task is marked as completed in `tasks.md`
**When** the UI ticker updates
**Then** the progress bar and percentage should update instantly without scrolling the screen.

### Scenario 3: Clean Exit
**Given** the `hb` command finishes execution
**When** the process exits
**Then** the status bar should be erased from the terminal, leaving only the standard log output.

## Constraints
- **Tech Stack**: Go (Golang).
- **Library**: Use `pterm` or raw ANSI for performance; avoid heavy frameworks if they interfere with "auto-injection".
- **Coverage**: 100% unit test coverage for the UI logic.
