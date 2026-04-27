import { Suspense } from 'react'
import { Button } from '@/components/ui/button'
import { Card } from '@/components/ui/card'

async function getData() {
  const res = await fetch('https://api.example.com/data', {
    cache: 'no-store',
  })
  return res.json()
}

function DataDisplay() {
  return (
    <Suspense fallback={<div>Loading data...</div>}>
      <DataContent />
    </Suspense>
  )
}

async function DataContent() {
  const data = await getData()

  return (
    <Card className="p-6">
      <h2 className="text-xl font-semibold mb-4">Data</h2>
      <pre className="text-sm">{JSON.stringify(data, null, 2)}</pre>
    </Card>
  )
}

export default function HomePage() {
  return (
    <main className="container mx-auto p-4">
      <header className="mb-8">
        <h1 className="text-4xl font-bold">Next.js Starter</h1>
        <p className="text-muted-foreground mt-2">
          A minimal starter with App Router and Tailwind v4
        </p>
      </header>

      <section className="grid gap-6 md:grid-cols-2">
        <Card className="p-6">
          <h2 className="text-xl font-semibold mb-4">Features</h2>
          <ul className="space-y-2">
            <li>✅ Next.js 14+ with App Router</li>
            <li>✅ Tailwind CSS v4</li>
            <li>✅ TypeScript</li>
            <li>✅ Shadcn/UI components</li>
          </ul>
        </Card>

        <DataDisplay />
      </section>

      <section className="mt-8">
        <Button>Get Started</Button>
      </section>
    </main>
  )
}