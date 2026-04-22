---
name: benchmark-expert
version: 1.0.0
description: "Expert Skill para medição de baselines de performance, detecção de regressões e comparação de stacks."
category: performance
---

# Benchmark Expert — Performance Baseline & Regression Detection

## 🔒 Prerequisites (Mandatory)
Esta skill opera DENTRO do framework **SDD**. Antes de iniciar qualquer benchmark:
1. **Context Check**: Você reidratou o contexto lendo `STATE.md`, `MEMORY.md` e `LEARNINGS.md`?
2. **Phase 4 Integration**: benchmarks devem ser executados preferencialmente na **Fase 4 (Review & Persistence)** do SDD para validar o impacto da implementação.
3. **Clean Environment**: Garanta que o ambiente de teste esteja isolado e sem processos de fundo pesados.

---

## Quando Usar

- Antes e depois de um PR para medir o impacto na performance.
- Configuração de baselines iniciais para novos projetos.
- Quando usuários reportam que o sistema "está lento".
- Antes de um lançamento oficial para garantir que os targets foram atingidos.
- Comparação de diferentes bibliotecas ou stacks (ex: Vite vs Webpack).

---

## Modos de Operação

### Modo 1: Performance de Página (Browser)
Mede métricas reais do navegador via ferramentas de automação.

**Fluxo:**
1. Navegar para URLs alvo.
2. Medir **Core Web Vitals**:
   - **LCP (Largest Contentful Paint)**: Alvo < 2.5s.
   - **CLS (Cumulative Layout Shift)**: Alvo < 0.1.
   - **INP (Interaction to Next Paint)**: Alvo < 200ms.
   - **FCP (First Contentful Paint)**: Alvo < 1.8s.
   - **TTFB (Time to First Byte)**: Alvo < 800ms.
3. Medir tamanhos de recursos:
   - Peso total da página (Alvo < 1MB).
   - JS Bundle (Alvo < 200KB gzipped).
   - CSS, Imagens e Scripts de terceiros.

### Modo 2: Performance de API
Benchmarks de endpoints de API para latência e robustez.

**Fluxo:**
1. Executar 100+ requisições no endpoint alvo.
2. Medir percentis: **p50, p95, p99**.
3. Monitorar códigos de status e tamanho das respostas.
4. Testar sob carga concorrente (ex: 10-50 conexões simultâneas).

### Modo 3: Performance de Build & Dev Loop
Mede a agilidade do ecossistema de desenvolvimento.

**Métricas:**
1. **Cold Build**: Tempo de build do zero.
2. **Hot Reload (HMR)**: Tempo entre salvar arquivo e refletir na UI.
3. **Test Duration**: Tempo total da suíte de testes.
4. **Tooling Latency**: Tempo de Lint e Type Check.

### Modo 4: Comparação Antes/Depois (Regressão)
O coração da skill para validação de PRs.

```bash
# Salva o estado atual como baseline
/benchmark-expert baseline

# ... aplica mudanças de código ...

# Compara o estado atual contra o baseline salvo
/benchmark-expert compare
```

---

## Output e Armazenamento

Os baselines são salvos em `.benchmarks/` na raiz do projeto em formato JSON. Devem ser commitados no Git para que a equipe compartilhe as mesmas referências.

---

## Quality Rules

- **Deterministic Measurements**: Sempre execute o benchmark 3 vezes e use a média.
- **Context Awareness**: Documente as especificações da máquina onde o benchmark foi rodado.
- **Fail Fast**: Se o delta de uma métrica crítica (LCP) subir > 15%, o veredito deve ser **FAIL** mandatório.

---

## Referências
- [Performance Targets](./references/performance_targets.md)
- [Análise de Regressão](./references/regression_analysis.md)
