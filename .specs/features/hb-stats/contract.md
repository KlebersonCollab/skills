# Validation Contract: HB Stats

## Validation Sensors
- [x] **Sensor: Metric Calculation**: Must correctly aggregate task counts from all feature directories.
- [x] **Sensor: Terminal Rendering**: Progress bars must render correctly with colors and percentages.
- [x] **Sensor: Multi-Feature Support**: Must handle repositories with > 20 features without performance degradation.

## Evidence
- Manual audit of output against `tasks.md` files.
- Stress test with artificial feature creation.
