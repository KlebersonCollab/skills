# Contract: HB-CLI Premium UI

## Agreement
The Implementer (Agent) will deliver a functional and highly aesthetic UI for the `hb` CLI. The Reviewer (User/CI) will validate the delivery based on the following sensors.

## Sensors (Validation)
1.  **Sensor-UI-1 (Visual)**: Running `hb status --ui` must show a fixed status bar at the bottom.
2.  **Sensor-UI-2 (Integrity)**: Executing `go test ./hb/internal/ui/... -cover` must return `100.0% coverage`.
3.  **Sensor-UI-3 (UX)**: Resizing the terminal during a long-running task must not break the UI layout.
4.  **Sensor-UI-4 (Cleanliness)**: After the command finishes, the terminal prompt must return to the top of the line without leaving "ghost" status bars.

## KPIs
- **Max Refresh Latency**: < 150ms.
- **CPU Overhead**: < 2% of a single core.
- **Binary Size Increase**: < 500KB.
