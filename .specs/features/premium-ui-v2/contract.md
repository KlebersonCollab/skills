# Contract: Premium UI v2

## Delivery Agreement
The delivery will be considered successful when the `ui/` directory contains a functional Vue 3 application that meets all Acceptance Criteria defined in `spec.md`.

## Validation Sensors
- **S1 (Structure)**: `ui/package.json` must exist and include `vue` and `vuetify`.
- **S2 (Visual)**: Dashboard must load without console errors and show at least one stat card.
- **S3 (Functionality)**: Navigation between "Skills" and "Features" views must work.
- **S4 (Data)**: The UI must display real data from the Hub (total skills count or feature names).

## Success Metrics
- 100% ACs met.
- No linting errors in the `ui/` project.
- Build command `npm run build` succeeds.
