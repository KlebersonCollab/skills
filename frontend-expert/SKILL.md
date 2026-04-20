---
name: frontend-expert
version: 1.1.0
description: Expert em desenvolvimento de interfaces modernas com React, Next.js e TailwindCSS v4.
---

# Frontend Expert

> "Interfaces não são apenas pixels; são conversas entre o usuário e a lógica do sistema."

---

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar a implementação:
1. **Spec Check**: O arquivo `spec.md` existe e contém os critérios de aceitação (ACs)?
2. **Plan Check**: O arquivo `plan.md` define a arquitetura e schemas?
3. **Task Check**: A lista de tarefas em `tasks.md` está detalhada?

---

## Goal

Prover expertise avançada para a construção de interfaces web de alta performance, acessíveis (A11y), e escaláveis. Foca no ecossistema moderno de React, priorizando Server Components, design system-first com Tailwind v4 e Shadcn/UI.

## Output Structure

A execução desta skill resulta nos seguintes artefatos:

| Artefato | Formato | Descrição |
|----------|---------|-----------|
| **Component Spec** | `.md` | Especificação detalhada de componentes e estados. |
| **Page Layout** | `.tsx` | Estrutura de página usando App Router do Next.js. |
| **Theme Config** | `.css` | Configuração de variáveis de tema via Tailwind v4. |
| **Accessibility Audit** | `.md` | Relatório de conformidade WCAG. |

## Quality Rules

- **Mobile-First**: Design responsivo é a base, não um ajuste posterior.
- **Server-First**: Priorizar React Server Components (RSC) para reduzir o bundle size.
- **A11y Mandate**: Todo componente interativo deve ser navegável por teclado e amigável a leitores de tela.
- **Tailwind v4 CSS-First**: Utilizar o novo motor do Tailwind baseado em variáveis CSS nativas.
- **Type Safety**: 100% TypeScript; evitar `any` a todo custo.

## Prohibited

- **NUNCA** utilizar inline styles (`style={{...}}`) quando utilitários Tailwind estiverem disponíveis.
- **NUNCA** ignorar estados de loading e error em fetches de dados.
- **NUNCA** utilizar bibliotecas de componentes pesadas se o Shadcn/UI (Copy-Paste) for suficiente.
- **NUNCA** misturar Client Components em Server Components sem necessidade técnica real.

---

## 🛠️ Stack Recomendada

1. **Framework**: Next.js (App Router).
2. **Styling**: TailwindCSS v4 (CSS variables engine).
3. **Components**: Shadcn/UI + Radix UI.
4. **State**: React Context para UI local, TanStack Query para Server State.
5. **Forms**: React Hook Form + Zod.

---

## 🔄 Workflow

### Fase 1: Design Review
Validar a hierarquia de componentes e o fluxo de dados (props vs state).

### Fase 2: Component Scaffolding
Criar componentes atômicos (Button, Input, Card) antes de montar as páginas.

### Fase 3: Integration & Performance
Implementar fetch de dados (Server Actions/Fetch API) e otimizar imagens/fonts.

### Fase 4: UX & A11y Validation
Testar em diferentes viewports e validar acessibilidade.
