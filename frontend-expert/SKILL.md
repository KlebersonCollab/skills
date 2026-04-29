---
name: frontend-expert
version: 1.1.0
description: Expert in modern interface development with React, Next.js, and TailwindCSS v4.
---

# Frontend Expert

> "Interfaces are not just pixels; they are conversations between the user and the system logic."

---

## 🔒 Prerequisites (Mandatory)
This skill operates WITHIN the **SDD** framework. Before starting any technical execution:
0. **Mode Check**: Verify the current operational mode (`.hub-mode`) and apply the `token-distiller` skill guidelines.
1. **Context Check**: Did you rehydrate the context by reading `STATE.md`, `MEMORY.md`, and `LEARNINGS.md`?
2. **Spec Check**: Does the `spec.md` file exist with clear requirements and Acceptance Criteria (ACs)? (BDD mandatory for Medium+).
3. **Plan Check**: Does the `plan.md` file define the architecture and schemas, and include **Mermaid** diagrams?
4. **Contract Check**: Was the `contract.md` file established with validation sensors?
5. **Task Check**: Is the task list in `tasks.md` detailed and atomized?

---
## Goal

Provide advanced expertise for building high-performance, accessible (A11y), and scalable web interfaces. Focuses on the modern React ecosystem, prioritizing Server Components, design system-first with Tailwind v4, and Shadcn/UI.

## Output Structure

The execution of this skill results in the following artifacts:

| Artifact | Format | Description |
|----------|---------|-----------|
| **Component Spec** | `.md` | Detailed specification of components and states. |
| **Page Layout** | `.tsx` | Page structure using Next.js App Router. |
| **Theme Config** | `.css` | Theme variable configuration via Tailwind v4. |
| **Accessibility Audit** | `.md` | WCAG compliance report. |

## Quality Rules

- **Mobile-First**: Responsive design is the foundation, not an afterthought.
- **Server-First**: Prioritize React Server Components (RSC) to reduce bundle size.
- **A11y Mandate**: Every interactive component must be keyboard-navigable and screen-reader friendly.
- **Tailwind v4 CSS-First**: Use the new Tailwind engine based on native CSS variables.
- **Type Safety**: 100% TypeScript; avoid `any` at all costs.

## Prohibited

- **NEVER** use inline styles (`style={{...}}`) when Tailwind utilities are available.
- **NEVER** ignore loading and error states in data fetches.
- **NEVER** use heavy component libraries if Shadcn/UI (Copy-Paste) is sufficient.
- **NEVER** mix Client Components into Server Components without a real technical need.

---

## 🛠️ Recommended Stack

1. **Framework**: Next.js (App Router).
2. **Styling**: TailwindCSS v4 (CSS variables engine).
3. **Components**: Shadcn/UI + Radix UI.
4. **State**: React Context for local UI, TanStack Query for Server State.
5. **Forms**: React Hook Form + Zod.

---

## 🧩 Shadcn/UI Integration Rules

- **Direct Installation**: Use `npx shadcn@latest add [component]` to install components directly into the codebase.
- **Customization via `cn()`**: Always use the `cn()` utility (clsx + tailwind-merge) for composing and overriding classes cleanly.
- **Component Variants**: Rely on `class-variance-authority` (`cva`) for deterministic component variations (e.g., button sizes, alert colors).
- **Extensibility**: Do not modify the primitive inside `components/ui` directly unless necessary; instead, compose them in `components/[custom]`.

## 🎨 Stitch HTML-to-React Conversion

- **Modularization**: Break downloaded Stitch HTML into modular, independent files. Do not output massive single-file components.
- **Logic Isolation**: Move all state, event handlers, and business logic into custom hooks in `src/hooks/`.
- **Data Decoupling**: Move static text, image URLs, and lists out of the view layer into `src/data/mockData.ts`.
- **Style Mapping**: Extract Tailwind classes from the Stitch HTML, sync them with the project's style guide (or Shadcn theme), and use theme-mapped utilities over raw hex codes.
- **Type Safety**: Enforce `Readonly<[ComponentName]Props>` on every converted component.

---

## 🔄 Workflow

### Phase 1: Design Review
Validate component hierarchy and data flow (props vs. state).

### Phase 2: Component Scaffolding
Create atomic components (Button, Input, Card) before assembling pages.

### Phase 3: Integration & Performance
Implement data fetching (Server Actions/Fetch API) and optimize images/fonts.

### Phase 4: UX & A11y Validation
Test across different viewports and validate accessibility.
