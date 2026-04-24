# React Best Practices

## Component Design

### Atomic Design Pattern
Organize components in a clear hierarchy:
- **Atoms**: Basic elements (Button, Input, Label)
- **Molecules**: Compositions of atoms (SearchBar, Card)
- **Organisms**: Complex structures (Header, Form)
- **Templates**: Page layouts
- **Pages**: Complete instances

### Component Composition
```tsx
// ✅ Good - Composable
<Button variant="primary" size="large">
  <Icon icon="send" />
  Submit
</Button>

// ❌ Bad - Props drilling
<Button icon="send" text="Submit" variant="primary" size="large" />
```

## State Management

### When to Use What
| Scenario | Solution |
|----------|----------|
| Local UI (modals, toggles) | React Context / useState |
| Server state (API data) | TanStack Query |
| Form state | React Hook Form |
| Global app state | Zustand / Redux |

### Server Components vs Client Components
```tsx
// ✅ Server Component (default)
export default async function ProductList() {
  const products = await db.products.findMany()
  return <div>{products.map(p => <ProductCard key={p.id} product={p} />)}</div>
}

// ✅ Client Component (when needed)
'use client'
export function AddToCartButton({ productId }: { productId: string }) {
  const [loading, setLoading] = useState(false)
  // ... interactive logic
}
```

## Performance

### Code Splitting
```tsx
import dynamic from 'next/dynamic'

const HeavyChart = dynamic(() => import('./HeavyChart'), {
  loading: () => <ChartSkeleton />,
  ssr: false
})
```

### Image Optimization
```tsx
import Image from 'next/image'

<Image
  src="/hero.jpg"
  alt="Hero section"
  width={1200}
  height={600}
  priority
  placeholder="blur"
/>
```

## Accessibility (A11y)

### Keyboard Navigation
```tsx
// ✅ Accessible button
<button
  onClick={handleAction}
  onKeyDown={(e) => e.key === 'Enter' && handleAction()}
  aria-label="Close dialog"
>
  <XIcon />
</button>
```

### Semantic HTML
```tsx
// ✅ Semantic
<nav aria-label="Main navigation">
  <ul>
    <li><Link href="/">Home</Link></li>
  </ul>
</nav>

// ❌ Non-semantic
<div className="nav">
  <div onClick={() => router.push('/')}>Home</div>
</div>
```

## TypeScript Best Practices

### Type Safety
```tsx
// ✅ Explicit types
interface ButtonProps {
  variant: 'primary' | 'secondary' | 'ghost'
  size?: 'sm' | 'md' | 'lg'
  children: React.ReactNode
  onClick?: () => void
}

// ❌ Avoid any
function Button({ variant, ...props }: any) {
  // ...
}
```

### Generic Components
```tsx
interface ListProps<T> {
  items: T[]
  renderItem: (item: T) => React.ReactNode
  keyExtractor: (item: T) => string
}

function List<T>({ items, renderItem, keyExtractor }: ListProps<T>) {
  return (
    <ul>
      {items.map(item => (
        <li key={keyExtractor(item)}>{renderItem(item)}</li>
      ))}
    </ul>
  )
}
```

## Testing

### Component Testing
```tsx
import { render, screen } from '@testing-library/react'
import userEvent from '@testing-library/user-event'

describe('Button', () => {
  it('calls onClick when clicked', async () => {
    const handleClick = vi.fn()
    render(<Button onClick={handleClick}>Click me</Button>)

    await userEvent.click(screen.getByRole('button', { name: 'Click me' }))
    expect(handleClick).toHaveBeenCalledTimes(1)
  })
})
```

## Resources
- [React Documentation](https://react.dev)
- [Next.js Documentation](https://nextjs.org/docs)
- [Tailwind CSS v4](https://tailwindcss.com/docs)
- [Radix UI](https://www.radix-ui.com)
- [Shadcn/UI](https://ui.shadcn.com)
