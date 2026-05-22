
# Linear Design Redesign — Specification

## Objective
Apply the Linear design system (defined in `/home/kleberson/Documentos/skills/DESIGN.md`) to the hub-ui-skills React application, ensuring visual consistency with the Linear brand: near-black canvas, lavender-blue accent, specific typography, spacing, border radius, and component styles.

## Functional Requirements

| ID | Description |
|---|---|
| FR-1 | Implement all design tokens (colors, typography, spacing, border radius) as CSS custom properties. |
| FR-2 | Update `index.html` to load Inter and JetBrains Mono fonts. |
| FR-3 | Update `index.css` to include full typography scale, border radius scale, spacing scale, and component-specific styles. |
| FR-4 | Re-style Header to match Linear top-nav: height 56px, canvas background, tabs as pricing-tab style. |
| FR-5 | Re-style sidebar search input to match `text-input` component (8px radius, 8px 12px padding, focus ring). |
| FR-6 | Re-style skill-list cards to match `feature-card` component (surface-1, 12px radius, 24px padding, hairline border). |
| FR-7 | Re-style harness sidebar to match Linear component cards (info-card, session-card). |
| FR-8 | Re-style detail panel and session detail panel to match Linear panels (surface-1, hairline border, proper spacing). |
| FR-9 | Re-style ecosystem graph nodes to match Linear cards (surface-2, 8px radius, hairline border). |
| FR-10 | Apply button styles where applicable (primary, secondary). |

## Acceptance Criteria

- AC-1: Application background is #010102, accent color is #5e6ad2.
- AC-2: Typography uses Inter at weights 400, 500, 600 with correct letter-spacing (display: -0.5px, body: -0.05px).
- AC-3: Border radii follow Linear scale: 8px buttons/inputs, 12px cards, 16px large panels, pill tabs.
- AC-4: Components match spec: buttons 8px radius, inputs focus ring (2px #5e6ad2 at 50% opacity), cards 1px hairline borders.
- AC-5: Overall visual resembles Linear's marketing aesthetic — dark, dense, product screenshots framed in charcoal panels.

## Out of Scope

- Light mode.
- Additional pages beyond existing Skills / Harness views.
- Real product screenshots (placeholder icon remains).
