# Modern Component Structure

This guide defines the standard for creating React components using TailwindCSS v4 and Design System principles.

## 📁 Folder Organization
```text
components/
├── ui/                 # Base components (Buttons, Inputs, Modals)
├── layout/             # Structural components (Header, Sidebar)
└── features/           # Feature-specific components
```

## 🏗️ File Structure
Each component should be contained in its own folder if it has complex logic or specific styles:
```text
MyComponent/
├── index.ts            # Clean export
├── MyComponent.tsx     # Logic and Structure
├── MyComponent.test.ts # Unit Tests
└── types.ts            # Interfaces and Types
```

## 🎨 TailwindCSS v4 + Modern CSS
- **Abstraction**: Use `@apply` only for global patterns. For specific components, prefer utility classes in the JSX.
- **Variables**: Use CSS variables for colors and spacing defined in the theme.
- **Interactivity**: Leverage Tailwind v4's new capabilities for complex states (e.g., `group-has-*`).

## 📏 Golden Rules
1. **Single Responsibility**: A component should do only one thing well.
2. **Composition over Inheritance**: Use the `children` pattern to create flexible components.
3. **Typed Props**: Always use TypeScript to define props, ensuring safety at development time.
