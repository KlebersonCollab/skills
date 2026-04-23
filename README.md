# 🧠 AI Agent Skills Hub

> Hub centralizado para desenvolvimento, armazenamento e evolução **Skills** modulares para agentes IA.

[![Skills](https://img.shields.io/badge/Skills-25-brightgreen)](#-skills-disponíveis)
[![Licença](https://img.shields.io/badge/Licença-MIT-blue)](LICENSE)
[![GitHub Actions](https://img.shields.io/github/actions/workflow/status/KlebersonCollab/skills/generate-skills-artifacts.yml?branch=main&label=Build%20Artifacts)](https://github.com/KlebersonCollab/skills/actions)
[![Status](https://img.shields.io/badge/Status-Ativo-brightgreen)](#-roadmap)

---

## 📖 Sobre o Projeto

Este repositório é **fonte verdade** para todas habilidades (skills) utilizadas por agentes IA. Cada skill é módulo independente, documentado e versionado, que pode ser integrado qualquer agente compatível.

Utilizamos metodologia **[SDD](sdd/)** (Spec-Driven Development) para garantir que cada nova funcionalidade seja rigorosamente especificada, planejada e verificada antes implementação.

---

## 📦 Download & Uso Rápido

Geramos automaticamente pacotes pré-configurados para principais agentes IA. Basta baixar, descompactar na raiz seu projeto e começar usar:

| Agente | Artefato | Conteúdo |
|--------|----------|----------|
| **Claude AI** | [📥 claude-skills.zip](https://github.com/KlebersonCollab/skills/releases/latest/download/claude-skills.zip) | Pasta `.claude/` com `CLAUDE.md` e todas as skills. |
| **Gemini CLI** | [📥 gemini-skills.zip](https://github.com/KlebersonCollab/skills/releases/latest/download/gemini-skills.zip) | Pasta `.gemini/` com `GEMINI.md` e todas as skills. |
| **AI Agent (Generic)** | [📥 agent-skills.zip](https://github.com/KlebersonCollab/skills/releases/latest/download/agent-skills.zip) | Pasta `.agent/` com `AGENT.md` e todas as skills. |

> 💡 *Nota: links acima baixam automaticamente versão mais recente aba **Releases**.*

### 🚀 Como usar as Skills
1. Baixe ZIP correspondente ao seu agente.
2. Extraia conteúdo na raiz diretório seu projeto.
3. Certifique-se que pasta oculta (ex: `.claude/`) foi criada corretamente.
4. Seu agente IA detectará automaticamente novas capacidades.

---

## 🛠️ Ferramentas & Automação

 Hub inclui **SDD CLI**, ferramenta Python que automatiza ciclo vida skills e gestão conhecimento:

| Comando | Descrição |
|---------|-----------|
| `uv run sdd init <name>` | Inicializa nova feature com estrutura SDD completa. |
| `uv run sdd task <feat> <id>` | Marca progresso tarefas e sincroniza estado global. |
| `uv run sdd graph` | Gera automaticamente **Knowledge Map** (Mermaid) projeto. |
| `uv run sdd sync` | Sincroniza mandatos globais em todos agentes (.gemini, .claude, .agent). |

> 💡 *Nota: Requer [Python UV](https://docs.astral.sh/uv/) instalado.*

---

## 🛠️ Skills Disponíveis

| # | Skill | Descrição | Versão |
|---|-------|-----------|--------|
| 1 | **[SDD](sdd/)** | Spec-Driven Development — Workflow modular e adaptativo com **Persistent Memory Protocol**. | `1.5.0` |
| 2 | **[Skill Factory](skill-factory/)** | Core Framework para criação padronizada de novas skills com scaffolding e validação. | `1.1.0` |
| 3 | **[Python com UV](python-uv/)** | Desenvolvimento Python profissional com UV — gerenciador 10-100x mais rápido. | `2.6.0` |
| 4 | **[API Architect](api-architect/)** | Arquiteto de APIs — design interoperável e seguro (OpenAPI, GraphQL, tRPC). | `1.3.0` |
| 5 | **[Brainstorming](brainstorming/)** | Facilitador de design e resolução de problemas — exploração profunda e validação. | `1.1.0` |
| 6 | **[Architecture](architecture/)** | Arquiteto de Sistemas — design pragmático, trade-offs e registros de decisão (ADR). | `2.0.1` |
| 7 | **[Flutter com FVM](flutter-fvm/)** | Desenvolvimento Flutter profissional com FVM. Inclui Dart 3+, performance otimizada, acessibilidade mandatória e segurança OWASP. | `1.3.0` |
| 8 | **[Azure DevOps](azure-devops/)** | Gerenciamento profissional de Boards, Repos, Pipelines e Artifacts no AzDO. | `1.1.0` |
| 9 | **[Clean Code Mentor](clean-code-mentor/)** | Mentor técnico e revisor de código focado em SOLID, YAGNI, DRY e KISS. | `1.0.0` |
| 10 | **[Observability Expert](observability-expert/)** | Especialista em SRE — Logs Estruturados, OpenTelemetry e Confiabilidade (SLI/SLO). | `1.0.0` |
| 11 | **[Onboarding Navigator](onboarding-navigator/)** | Guia interativo de cultura e engenharia para novos membros e projetos. | `1.5.0` |
| 12 | **[YouTube Transcript](youtube-transcript/)** | Extração de transcrições de vídeos com fallback para IA (Whisper) e limpeza automática. | `1.0.0` |
| 13 | **[Harness Expert](harness-expert/)** | Motor técnico para Harness Engineering (Sync, Rehydrate, Automação). | `2.0.0` |
| 14 | **[Knowledge Architect](knowledge-architect/)** | Arquitetura de conhecimento local via grafos relacionais (Local GraphRAG). | `1.0.0` |
| 15 | **[FastAPI Expert](fastapi-expert/)** | Excelência em APIs com FastAPI — performance, Annotated e Pydantic V2. | `1.1.0` |
| 16 | **[Django Expert](django-expert/)** | Desenvolvimento profissional com Django — Arquitetura, Segurança e TDD. | `1.5.0` |
| 17 | **[Swarm Facilitator](swarm-facilitator/)** | Orquestrador de workflows Multi-Agente (Swarm). Define protocolos de handoff. | `1.1.0` |
| 18 | **[Scaffolding Expert](scaffolding-expert/)** | Gerador dinâmico de projetos. Utiliza CLI tools (uvx copier) para inicializar boilerplates. | `1.0.0` |
| 19 | **[Golang Expert](golang-expert/)** | Excelência em Go — Performance, concorrência idiomática e ecossistema Samber. | `1.2.0` |
| 20 | **[Frontend Expert](frontend-expert/)** | Expert em interfaces modernas com React, Next.js e TailwindCSS v4. | `1.1.0` |
| 21 | **[Git Workflow](git-workflow/)** | Padrões de fluxo de trabalho Git, estratégias de branching e integração com SDD. | `1.0.0` |
| 22 | **[Golang Testing Expert](golang-testing-expert/)** | Especialista em QA para Go — TDD, Table-Driven Tests, Benchmarks e Fuzzing. | `1.0.0` |
| 23 | **[Benchmark Expert](benchmark-expert/)** | Expert Skill para medição de baselines de performance, detecção de regressões e comparação de stacks. | `1.0.0` |
| 24 | **[DevSecOps Expert](devsecops-expert/)** | Auditoria estática de qualidade, segurança SAST/DAST e Hardening. | `1.0.0` |
| 25 | **[Token Distiller](token-distiller/)** | Gerenciador de densidade de tokens e modos de compressão (Caveman/Premium). | `1.1.0` |

---

## 🏗️ Estrutura do Repositório

```
skills/
85: ├── .github/workflows/           # Automação de CI/CD (GitHub Actions)
86: ├── .specs/                      # Especificações do projeto (SDD)
87: │   ├── project/                 # Visão, Roadmap e Estado
88: │   └── codebase/                # Stack e Convenções
89: ├── scripts/                     # Scripts de automação e distribuição
90: ├── <skill>/                     # 🧩 Cada skill em seu diretório
91: │   ├── README.md                # Documentação detalhada
92: │   ├── SKILL.md                 # Definição técnica principal
93: │   └── CHANGELOG.md             # Histórico de versões
94: └── README.md                    # ← Você está aqui
95: ```

---

## 📐 Como Criar uma Nova Skill

Utilize skill **[Skill Factory](skill-factory/)** para criar novas skills forma padronizada. processo garante que cada skill seja gerada com estrutura correta e passe por auditoria qualidade antes ser integrada ao hub.

> 📖 Consulte [documentação do Skill Factory](skill-factory/) para detalhes.

---

## 🗺️ Roadmap & Estado do Projeto

Acompanhe evolução hub através documentos planejamento:
- 🛤️ **[ROADMAP.md](.specs/project/ROADMAP.md)**: Visão longo prazo e metas.
- 📊 **[STATE.md](.specs/project/STATE.md)**: Status atual e tarefas em andamento.

---

## 📄 Licença

Este projeto está licenciado sob [Licença MIT](LICENSE).

---

<div align="center">

**Construído com 🧠 por [Kleberson Romero](https://github.com/KlebersonCollab)**

*Precision at scale. Rigor when needed, speed when possible.*

</div>