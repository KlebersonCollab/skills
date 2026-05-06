# Performance Targets & SLAs

This guide defines the acceptable limits for performance metrics in our ecosystem.

## Core Web Vitals (Browser)

| Metric | Good (Green) | Needs Improvement (Yellow) | Poor (Red) |
|---------|---------------|---------------------------|-----------------|
| **LCP** | < 2.5s        | 2.5s - 4.0s               | > 4.0s          |
| **FID** | < 100ms       | 100ms - 300ms             | > 300ms         |
| **CLS** | < 0.1         | 0.1 - 0.25                | > 0.25          |
| **FCP** | < 1.8s        | 1.8s - 3.0s               | > 3.0s          |
| **INP** | < 200ms       | 200ms - 500ms             | > 500ms         |

## API Latency (Server)

| Scope | p50 (Median) | p95 (95%) | p99 (Worst case) |
|--------|---------------|-----------|-----------------|
| Internal (Microservice) | < 20ms | < 50ms | < 100ms |
| External (Edge) | < 100ms | < 250ms | < 500ms |
| Heavy Processing | < 500ms | < 1.5s | < 3s |

## Build & Tooling

- **HMR (Hot Module Replacement)**: Must be < 300ms for a fluid dev experience.
- **Cold Install**: `npm install` must be < 60s in a clean environment.
- **Lint + Type Check**: Must run in < 10s to not break mental flow.

## Good Patterns
- **Code Splitting**: Keep initial bundles small.
- **Image Optimization**: Use WebP/Avif and `loading="lazy"`.
- **Caching**: Implement robust Cache-Control and ETag strategies.

## Anti-Patterns
- **Dependency Bloat**: Adding giant libs for simple functions.
- **Waterfall Requests**: Network waterfall that blocks rendering.
- **Synchronous JS at Top**: Script tags without `defer` or `async`.


---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HUB-ALIGNMENT"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-06T13:16:19.359088Z"
evidence_checksum: "NONE"
```
