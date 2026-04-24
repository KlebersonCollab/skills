# Performance Regression Analysis

This guide explains how to interpret performance changes between different code states.

## Why analyze regressions?
New code can introduce imperceptible slowness in development environments, but critical slowness in production. Regression analysis quantifies this impact.

## Comparison Methodology
1. **Confidence Factor**: Always run the benchmark 3 times to eliminate statistical noise.
2. **Change Delta**: Percentage difference between `After` and `Before`.
3. **Alert Thresholds**:
   - < 5%: Normal variation.
   - 5% - 15%: WARNING. Investigate the cause.
   - > 15%: Critical Regression (FAIL). Revert or optimize.

## What to Observe
- **JS Bundle Size**: Increased bundle size directly impacts TTI (Time to Interactive).
- **TTFB (Time to First Byte)**: If it increases, it indicates a regression in the backend or database.
- **LCP**: If it increases, it indicates resource prioritization issues or heavy rendering.

## How to Report
Use clear tables that highlight negative changes in red and positive changes in green.
See [examples/comparison_report.md](../examples/comparison_report.md) for a template.
