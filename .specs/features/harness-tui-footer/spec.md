# SDD Specification: TUI Fixed Status Bar Footer (`HARNESS-TUI-FOOTER`)

This is the specification and source of truth for the functional requirements of the TUI Fixed Footer feature.

---

## 1. Context & Goals
- **Problem Statement**: The current interactive loop prints raw text messages, but lacks a persistent, unified view of session metrics (session ID, active provider, model name, accumulated tokens, and USD cost). Users want a real-time fixed status bar at the bottom of their terminal window, mirroring the premium aesthetics of `pi.dev`.
- **Scope**: 
  - **Included**: A fixed footer at the very bottom line of the terminal screen, written to a restricted scroll region. The footer displays Logo, Session ID, Active Provider, Active Model, total tokens, and estimated session cost.
  - **Not Included**: Web-based TUI frames, full-screen graphical applications. Must remain zero-dependency and terminal buffer-friendly.
- **Success Criteria**: 
  - The status bar spans the full width of the terminal.
  - Scrolling of command outputs and assistant bubbles is restricted to the upper area of the screen, leaving the last line untouched.
  - resizes of the terminal window are checked dynamically to adapt the height and scroll margins.
  - On CLI exit, standard terminal scroll regions are clean and reset.

## 2. Requirements (BDD Scenarios)

### Feature: Persistent Status Footer

**Scenario 1: Interactive Loop Startup**
- **Given** The user launches the Harness interactive CLI via `make harness-start`
- **When** The TUI starts up
- **Then** The terminal scroll region is restricted to `1` to `height-1`, and a beautiful styled footer is drawn at line `height` containing the initial session metrics.

**Scenario 2: Real-time Metric Updates**
- **Given** The user completes an LLM call or tool execution which updates total tokens and session cost
- **When** The agent prints the response or transitions back to the prompt
- **Then** The footer is dynamically redrawn on the last line with the updated token count and USD cost.

**Scenario 3: Snappy and Safe Exit**
- **Given** The user types `/exit` or `Ctrl+C` to quit the session
- **When** The CLI interactive loop terminates
- **Then** The terminal scroll region is reset to standard full screen (`\033[r`) and cursor positioning is restored safely.

## 3. Constraints & Risks
- **Zero-Dependency Rule**: Must be implemented purely using standard Go `os/exec` command queries (`stty size`) and ANSI escape sequences.
- **TUI State Consistency**: Must ensure that whenever text is printed (e.g. log files, warnings), the cursor returns safely to its active typing position.

---
<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HARNESS-TUI-FOOTER"
phase: "SPECIFY"
status: "COMPLETED"
last_update: "2026-05-23T03:56:00Z"
evidence_checksum: "NONE"
```
