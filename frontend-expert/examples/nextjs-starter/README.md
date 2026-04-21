# Next.js Starter Example

This is a minimal Next.js 14+ starter with App Router, Tailwind CSS v4, and TypeScript.

## Features

- ✅ Next.js 14+ with App Router
- ✅ Tailwind CSS v4
- ✅ TypeScript
- ✅ Shadcn/UI components
- ✅ Server Components by default
- ✅ Responsive design

## Getting Started

```bash
npm install
npm run dev
```

## Project Structure

```
app/
├── layout.tsx          # Root layout
├── page.tsx            # Home page
├── globals.css         # Global styles
├── components/         # Reusable components
│   ├── ui/            # Shadcn/UI components
│   └── features/      # Feature-specific components
└── lib/               # Utilities and helpers
```

## Key Patterns

### Server Component Example

```tsx
// app/page.tsx
export default async function HomePage() {
  const data = await fetch('https://api.example.com/data')
    .then(res => res.json())

  return (
    <main className="container mx-auto p-4">
      <h1 className="text-3xl font-bold">Welcome</h1>
      <DataDisplay data={data} />
    </main>
  )
}
```

### Client Component Example

```tsx
// components/features/counter.tsx
'use client'

import { useState } from 'react'

export function Counter() {
  const [count, setCount] = useState(0)
  return (
    <button
      onClick={() => setCount(c => c + 1)}
      className="px-4 py-2 bg-blue-600 text-white rounded"
    >
      Count: {count}
    </button>
  )
}
```

### Server Action Example

```tsx
// app/actions.ts
'use server'

export async function createTodo(formData: FormData) {
  const title = formData.get('title') as string
  // Save to database
  revalidatePath('/todos')
}
```

## Styling with Tailwind v4

```css
/* app/globals.css */
@import "tailwindcss";

@theme {
  --color-primary: #3b82f6;
  --color-secondary: #64748b;
}
```

## Accessibility

All components follow WCAG 2.1 AA guidelines:
- Keyboard navigation
- Screen reader support
- Proper semantic HTML
- Focus management

## Performance

- Code splitting by default
- Image optimization with next/image
- Font optimization with next/font
- Automatic static generation where possible