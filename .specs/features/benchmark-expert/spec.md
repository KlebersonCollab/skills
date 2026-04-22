# Spec: Benchmark Expert Skill

## 1. Overview
A skill `benchmark-expert` é projetada para medir baselines de performance, detectar regressões antes/depois de PRs e comparar alternativas de stack dentro do ecossistema do Hub. Ela adapta a skill original da ECC para o "Standard de Ouro" do nosso projeto.

## 2. Requisitos Funcionais
- **RF01: Medição de Performance de Página (Mode 1)**: Deve medir métricas reais do navegador (LCP, CLS, INP, FCP, TTFB) e tamanhos de recursos (JS, CSS, Imagens).
- **RF02: Medição de Performance de API (Mode 2)**: Deve realizar benchmarks de endpoints de API medindo latência (p50, p95, p99), tamanho de resposta e códigos de status sob carga.
- **RF03: Medição de Performance de Build (Mode 3)**: Deve medir o ciclo de feedback de desenvolvimento (Cold Build, Hot Reload, Test Suite, Lint, Type Check).
- **RF04: Comparação Antes/Depois (Mode 4)**: Deve permitir salvar um baseline e compará-lo após alterações, emitindo um veredito (SUCCESS, WARNING, FAIL).
- **RF05: Armazenamento de Baselines**: Deve salvar resultados em um diretório padronizado para rastreamento via Git.

## 3. Requisitos Não-Funcionais
- **RNF01: Idioma**: Toda a documentação (`SKILL.md`, `README.md`, `CHANGELOG.md`) deve estar em **Português Brasileiro (PT-BR)**.
- **RNF02: Conformidade SDD**: Deve incluir o hook `🔒 Prerequisites (Mandatory)` vinculado ao processo SDD.
- **RNF03: Padrão de Ouro**: Deve incluir diretórios `references/` e `examples/` com guias de "Bom vs Ruim".
- **RNF04: Estrutura de Arquivos**: Deve seguir rigorosamente o scaffolding da `skill-factory`.

## 4. Critérios de Aceitação (AC)
- [ ] AC1: A skill deve ser criada no diretório `benchmark-expert/` via `skill-factory`.
- [ ] AC2: O arquivo `SKILL.md` deve descrever os 4 modos de operação detalhadamente em PT-BR.
- [ ] AC3: Deve existir um guia de referência em `references/performance_targets.md` definindo os SLAs aceitáveis.
- [ ] AC4: A skill deve passar na validação do `skill-factory-validator`.
- [ ] AC5: A skill deve ser registrada no `README.md` raiz.

## 5. BDD (User Stories)
**Cenário: Comparação de Performance após Refatoração**
- **Dado** que eu executei `/benchmark baseline` na branch `main`.
- **E** que eu apliquei uma refatoração de performance na minha feature branch.
- **Quando** eu executo `/benchmark compare`.
- **Então** eu devo ver uma tabela comparativa com as deltas de métricas.
- **E** o veredito deve ser `✓ BETTER` se as métricas melhorarem além da margem de erro.
