# Next.js App Router Guide

## App Router Fundamentals

### File-Based Routing
```
app/
├── layout.tsx          # Root layout
├── page.tsx            # Home page (/)
├── about/
│   └── page.tsx        # /about
├── blog/
│   ├── page.tsx        # /blog
│   └── [slug]/
│       └── page.tsx    # /blog/:slug
└── api/
    └── route.ts        # /api
```

### Layouts
```tsx
// app/layout.tsx
export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>
        <Header />
        {children}
        <Footer />
      </body>
    </html>
  )
}
```

### Nested Layouts
```tsx
// app/dashboard/layout.tsx
export default function DashboardLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <div className="flex">
      <Sidebar />
      <main>{children}</main>
    </div>
  )
}
```

## Server Components

### Default Behavior
All components in the App Router are Server Components by default.

```tsx
// ✅ Server Component (default)
export default async function ProductList() {
  const products = await db.products.findMany()
  return (
    <div>
      {products.map(product => (
        <ProductCard key={product.id} product={product} />
      ))}
    </div>
  )
}
```

### Client Components
Add `'use client'` directive for interactive components.

```tsx
'use client'

import { useState } from 'react'

export function Counter() {
  const [count, setCount] = useState(0)
  return <button onClick={() => setCount(c => c + 1)}>{count}</button>
}
```

### When to Use Client Components
- Event handlers (`onClick`, `onChange`)
- State management (`useState`, `useReducer`)
- Browser APIs (`window`, `document`)
- Custom hooks that use browser APIs

## Data Fetching

### Server-Side Fetching
```tsx
export default async function Page() {
  const data = await fetch('https://api.example.com/data', {
    cache: 'force-cache', // or 'no-store', 'no-cache'
  }).then(res => res.json())

  return <div>{data.title}</div>
}
```

### Revalidation
```tsx
export const revalidate = 3600 // Revalidate every hour

export default async function Page() {
  const data = await fetch('https://api.example.com/data')
  return <div>{data.title}</div>
}
```

### On-Demand Revalidation
```tsx
import { revalidatePath } from 'next/cache'

export async function POST() {
  // ... update data
  revalidatePath('/blog')
  return Response.json({ success: true })
}
```

## Server Actions

### Form Handling
```tsx
'use server'

import { revalidatePath } from 'next/cache'

export async function createTodo(formData: FormData) {
  const title = formData.get('title') as string
  await db.todos.create({ data: { title } })
  revalidatePath('/todos')
}
```

### Using Server Actions
```tsx
import { createTodo } from './actions'

export default function TodoForm() {
  return (
    <form action={createTodo}>
      <input name="title" type="text" />
      <button type="submit">Add Todo</button>
    </form>
  )
}
```

## Dynamic Routes

### Dynamic Segments
```tsx
// app/blog/[slug]/page.tsx
export default async function BlogPost({
  params,
}: {
  params: { slug: string }
}) {
  const post = await db.posts.findUnique({
    where: { slug: params.slug }
  })
  return <article>{post?.content}</article>
}
```

### Catch-All Segments
```tsx
// app/docs/[...slug]/page.tsx
export default async function DocsPage({
  params,
}: {
  params: { slug: string[] }
}) {
  const path = params.slug.join('/')
  // ...
}
```

## Metadata

### Static Metadata
```tsx
import { Metadata } from 'next'

export const metadata: Metadata = {
  title: 'My App',
  description: 'Description of my app',
}

export default function Page() {
  return <div>...</div>
}
```

### Dynamic Metadata
```tsx
import { Metadata } from 'next'

export async function generateMetadata({
  params,
}: {
  params: { slug: string }
}): Promise<Metadata> {
  const post = await db.posts.findUnique({
    where: { slug: params.slug }
  })
  return {
    title: post?.title,
    description: post?.excerpt,
  }
}
```

## Error Handling

### Error Boundaries
```tsx
// app/error.tsx
'use client'

export default function Error({
  error,
  reset,
}: {
  error: Error
  reset: () => void
}) {
  return (
    <div>
      <h2>Something went wrong!</h2>
      <button onClick={() => reset()}>Try again</button>
    </div>
  )
}
```

### Not Found Pages
```tsx
// app/not-found.tsx
export default function NotFound() {
  return <div>Page not found</div>
}
```

## Loading States

### Loading UI
```tsx
// app/loading.tsx
export default function Loading() {
  return <div>Loading...</div>
}
```

### Suspense Boundaries
```tsx
import { Suspense } from 'react'

export default function Page() {
  return (
    <div>
      <Suspense fallback={<Loading />}>
        <SlowComponent />
      </Suspense>
    </div>
  )
}
```

## Best Practices

### 1. Use Server Components by Default
Only use Client Components when necessary for interactivity.

### 2. Co-locate Data Fetching
Fetch data as close to where it's used as possible.

### 3. Use Streaming for Slow Components
Wrap slow components in Suspense boundaries.

### 4. Optimize Images
Use the `next/image` component for all images.

### 5. Use Server Actions for Mutations
Server Actions are the recommended way to handle form submissions.

## Resources
- [Next.js App Router Documentation](https://nextjs.org/docs/app)
- [Next.js Server Components](https://nextjs.org/docs/app/building-your-application/rendering/server-components)
- [Next.js Server Actions](https://nextjs.org/docs/app/building-your-application/data-fetching/server-actions-and-mutations)
