# Project State

## Current Phase: [PHASE 4: VERIFY] — Features Concluídas ✅

## Features Completas
- [x] Feature: `harness-solids-refactor` — Refatoração SOLID completa.
- [x] Feature: `harness-web-tools` — web_search, code_search, fetch_content.
- [x] Feature: `harness-skill-loading` — Skill Loading Ativo.
- [x] Feature: `harness-budget-management` — Context Budget Management + Makefile.
- [x] Feature: `harness-multi-provider` — Multi-Provider Fallback.
- [x] Feature: `harness-internal-tests` — 119 testes nos pacotes internos.
- [x] Feature: `harness-tui-components` — TUI Components + Error Recovery.
- [x] Feature: `harness-sdd-proactive` — SDD Framework Proativo.
- [x] Feature: `harness-voice-enhancements` — WAV header wrapping & prioritized STT/TTS fallback routing.
- [x] Feature: `harness-cli-bugfixes` — Fixed silent tool exits on successful empty-output commands and duplicate nested user nodes in history.

## SDD Framework Proativo

### Sub-Skills Implementadas
| Sub-Skill | Função | Artefatos |
|---|---|---|
| **sdd-planner** | Visão do projeto e memória | STATE.md, MEMORY.md, LEARNINGS.md |
| **sdd-orchestrator** | Especificação e design | spec.md, plan.md, tasks.md, contract.md |

### Fluxo Automático
```
DetectDevTask() → detecta tarefa de desenvolvimento
         │
         ▼
Phase 0: ALIGN — Grilling Session (sabatina)
  ├── 5 perguntas de alinhamento (objetivo, termos, alternativas, riscos, ADRs)
  ├── Atualiza CONTEXT.md (glossário de domínio)
  ├── Cria ADRs em .specs/architecture/XXXX-nome.md
  └── Registra decisões para a spec
         │
         ▼
sdd-orchestrator: Cria artefatos em .specs/features/[feature]/
  ├── spec.md (BDD scenarios com Given/When/Then)
  ├── plan.md (Mermaid diagram + architecture overview)
  ├── tasks.md (tabela com coluna Evidence para commits)
  └── contract.md (sensors de validação: build, test, lint)
         │
         ▼
sdd-planner: Atualiza memória do projeto
  ├── STATE.md (fase atual + progresso da feature)
  ├── MEMORY.md (fatos persistentes do projeto)
  └── LEARNINGS.md (padrões descobertos)
```

### Templates Seguidos
- spec.md → `sdd/resources/spec-template.md` (Context & Goals, BDD Scenarios, Constraints)
- plan.md → `sdd/resources/plan-template.md` (Mermaid diagram, Architecture, Strategy)
- tasks.md → `sdd/resources/tasks-template.md` (tabela com Evidence column obrigatório)
- contract.md → SDC com sensores de build, test, lint

## Status Final vs pi.dev

| Funcionalidade | Status |
|---|---|
| **Arquitetura SOLID** (SRP, OCP, LSP, ISP, DIP) | ✅ |
| **Agent Loop** completo (system → LLM → parse → tools → iterate) | ✅ |
| **LLM Providers** (Gemini, DeepSeek, Ollama, genérico) | ✅ |
| **Multi-Provider Fallback** (fallback automático entre providers) | ✅ |
| **Ferramentas locais** (read, write, patch, search, exec) | ✅ |
| **Ferramentas Web** (web_search, code_search, fetch_content) | ✅ |
| **Skill Loading Ativo** (detecta menção + carrega SKILL.md) | ✅ |
| **Context Budget Management** (40k warning, 60k critical, 80k max) | ✅ |
| **Token Distiller** (Caveman Lite/Full/Ultra, Compact, Safety) | ✅ |
| **MCP Server/Client** (OCP registry, LSP fix, SSE funcional) | ✅ |
| **Swarm Multi-Agente** (centralized, sequential, GAN) | ✅ |
| **Session Tree DAG** com persistência JSON | ✅ |
| **TUI Components** (Text, Box, Container, Selector, Border, etc.) | ✅ |
| **Error Recovery** (classificação, retry com backoff, fallback, display) | ✅ |
| **SDD Proativo** (Grilling Session → ADRs → spec/plan/tasks/contract) | ✅ |
| **.env auto-load** na inicialização | ✅ |
| **Makefile completo** (build, test, vet, clean, start, stop) | ✅ |
| **119 testes** (27 package main + 92 internal packages) | ✅ |

## Next Steps
- Testes para internal/agent, internal/mcp, internal/ui.
- Qualidade de código e refinamentos.

---

<!-- @sdd-state -->
```yaml
version: "2.3.0"
feature_id: "HARNESS-CLI-BUGFIXES"
phase: "VERIFY"
status: "COMPLETED"
last_update: "2026-05-23T02:10:00Z"
evidence_checksum: "go-build-and-test-pass"
```
