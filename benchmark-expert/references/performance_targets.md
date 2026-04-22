# Performance Targets & SLAs

Este guia define os limites aceitáveis para métricas de performance no nosso ecossistema.

## Core Web Vitals (Navegador)

| Métrica | Ótimo (Verde) | Precisa Melhorar (Amarelo) | Ruim (Vermelho) |
|---------|---------------|---------------------------|-----------------|
| **LCP** | < 2.5s        | 2.5s - 4.0s               | > 4.0s          |
| **FID** | < 100ms       | 100ms - 300ms             | > 300ms         |
| **CLS** | < 0.1         | 0.1 - 0.25                | > 0.25          |
| **FCP** | < 1.8s        | 1.8s - 3.0s               | > 3.0s          |
| **INP** | < 200ms       | 200ms - 500ms             | > 500ms         |

## API Latency (Servidor)

| Escopo | p50 (Mediana) | p95 (95%) | p99 (Pior caso) |
|--------|---------------|-----------|-----------------|
| Interna (Microserviço) | < 20ms | < 50ms | < 100ms |
| Externa (Edge) | < 100ms | < 250ms | < 500ms |
| Processamento Pesado | < 500ms | < 1.5s | < 3s |

## Build & Tooling

- **HMR (Hot Module Replacement)**: Deve ser < 300ms para uma experiência de dev fluida.
- **Cold Install**: `npm install` deve ser < 60s em ambiente limpo.
- **Lint + Type Check**: Deve rodar em < 10s para não quebrar o fluxo mental.

## Boas Práticas (Good Patterns)
- **Code Splitting**: Mantenha bundles iniciais pequenos.
- **Image Optimization**: Use WebP/Avif e `loading="lazy"`.
- **Caching**: Implemente estratégias robustas de Cache-Control e ETag.

## Anti-Padrões (Bad Patterns)
- **Bloat de Dependências**: Adicionar libs gigantes para funções simples.
- **Requisições em Cascata**: Waterfall de rede que bloqueia a renderização.
- **JS Síncrono no Topo**: Script tags sem `defer` ou `async`.
