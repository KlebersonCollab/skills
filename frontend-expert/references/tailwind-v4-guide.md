# Tailwind CSS v4 Guide

## What's New in v4

### CSS-First Configuration
Tailwind v4 uses native CSS variables instead of JavaScript configuration:

```css
@theme {
  --color-primary: #3b82f6;
  --color-secondary: #64748b;
  --font-sans: 'Inter', system-ui;
}

@utility focus-ring {
  @layer reset {
    outline: none;
  }
  outline: 2px solid var(--color-primary);
  outline-offset: 2px;
}
```

### No More tailwind.config.js
Configuration is now done in CSS:

```css
/* app/globals.css */
@import "tailwindcss";

@theme {
  --color-brand-50: #eff6ff;
  --color-brand-500: #3b82f6;
  --color-brand-900: #1e3a8a;

  --font-display: 'Inter', sans-serif;
  --font-body: 'Inter', sans-serif;

  --animate-fade-in: fade-in 0.3s ease-out;
}

@keyframes fade-in {
  from { opacity: 0; }
  to { opacity: 1; }
}
```

## Custom Utilities

### Creating Reusable Utilities
```css
@utility card {
  @apply rounded-lg border bg-white p-6 shadow-sm;
}

@utility btn {
  @apply inline-flex items-center justify-center rounded-md px-4 py-2 font-medium transition-colors;
}

@utility btn-primary {
  @apply btn bg-blue-600 text-white hover:bg-blue-700;
}
```

### Using Custom Utilities
```tsx
<div className="card">
  <button className="btn-primary">Click me</button>
</div>
```

## Theme Variables

### Color System
```css
@theme {
  --color-*: initial; /* All default colors */
  --color-primary-50: #eff6ff;
  --color-primary-100: #dbeafe;
  --color-primary-500: #3b82f6;
  --color-primary-600: #2563eb;
  --color-primary-700: #1d4ed8;
}
```

### Spacing Scale
```css
@theme {
  --spacing-0: 0;
  --spacing-1: 0.25rem;
  --spacing-2: 0.5rem;
  --spacing-3: 0.75rem;
  --spacing-4: 1rem;
  --spacing-6: 1.5rem;
  --spacing-8: 2rem;
  --spacing-12: 3rem;
  --spacing-16: 4rem;
}
```

### Typography Scale
```css
@theme {
  --font-sans: 'Inter', system-ui;
  --font-mono: 'Fira Code', monospace;

  --text-xs: 0.75rem;
  --text-sm: 0.875rem;
  --text-base: 1rem;
  --text-lg: 1.125rem;
  --text-xl: 1.25rem;
  --text-2xl: 1.5rem;
  --text-3xl: 1.875rem;
}
```

## Responsive Design

### Mobile-First Approach
```tsx
<div className="p-4 md:p-6 lg:p-8">
  <h1 className="text-xl md:text-2xl lg:text-3xl">Responsive Heading</h1>
</div>
```

### Container Queries (New in v4)
```tsx
<div className="@container">
  <div className="@md:text-lg @lg:text-xl">
    Text scales with container, not viewport
  </div>
</div>
```

## Dark Mode

### CSS-Based Dark Mode
```css
@media (prefers-color-scheme: dark) {
  @theme {
    --color-background: #0f172a;
    --color-foreground: #f8fafc;
  }
}
```

### Class-Based Dark Mode
```tsx
<html className="dark">
```

```css
.dark {
  @theme {
    --color-background: #0f172a;
    --color-foreground: #f8fafc;
  }
}
```

## Performance Tips

### Purge Unused Styles
Tailwind v4 automatically purges unused styles in production.

### Optimize Bundle Size
```css
/* Only import what you need */
@import "tailwindcss/theme" layer(theme);
@import "tailwindcss/preflight" layer(base);
@import "tailwindcss/utilities" layer(utilities);
```

## Migration from v3

### Key Changes
1. **No more `tailwind.config.js`** - Use CSS `@theme`
2. **No more `@apply` for everything** - Use custom utilities
3. **Native CSS variables** - Better performance and theming
4. **Container queries** - Responsive components

### Example Migration
```css
/* v3 */
@tailwind base;
@tailwind components;
@tailwind utilities;

@layer components {
  .btn {
    @apply px-4 py-2 rounded-md font-medium;
  }
}

/* v4 */
@import "tailwindcss";

@utility btn {
  padding-inline: 1rem;
  padding-block: 0.5rem;
  border-radius: 0.375rem;
  font-weight: 500;
}
```

## Resources
- [Tailwind CSS v4 Documentation](https://tailwindcss.com/docs)
- [Tailwind CSS v4 Upgrade Guide](https://tailwindcss.com/upgrade-guide)
- [Shadcn/UI with Tailwind v4](https://ui.shadcn.com)
