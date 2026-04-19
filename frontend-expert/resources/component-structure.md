# Estrutura de Componentes Modernos

Este guia define o padrão para criação de componentes React utilizando TailwindCSS v4 e princípios de Design System.

## 📁 Organização de Pastas
```text
components/
├── ui/                 # Componentes base (Botões, Inputs, Modais)
├── layout/             # Componentes de estrutura (Header, Sidebar)
└── features/           # Componentes específicos de funcionalidades
```

## 🏗️ Estrutura de Arquivo
Cada componente deve ser contido em sua própria pasta se possuir lógica complexa ou estilos específicos:
```text
MyComponent/
├── index.ts            # Exportação limpa
├── MyComponent.tsx     # Lógica e Estrutura
├── MyComponent.test.ts # Testes Unitários
└── types.ts            # Interfaces e Tipos
```

## 🎨 TailwindCSS v4 + CSS Moderno
- **Abstração**: Use `@apply` apenas para padrões globais. Para componentes específicos, prefira classes utilitárias no JSX.
- **Variáveis**: Utilize variáveis CSS para cores e espaçamentos definidos no tema.
- **Interatividade**: Aproveite as novas capacidades do Tailwind v4 para estados complexos (ex: `group-has-*`).

## 📏 Regras de Ouro
1. **Single Responsibility**: Um componente deve fazer apenas uma coisa bem feita.
2. **Composition over Inheritance**: Use o padrão `children` para criar componentes flexíveis.
3. **Props Tipadas**: Sempre use TypeScript para definir as props, garantindo segurança em tempo de desenvolvimento.
