# 🧠 AI Agent Skills Hub

> Hub centralizado para desenvolvimento, armazenamento e evolução de **Skills** modulares para agentes de IA.

[![Skills](https://img.shields.io/badge/Skills-12-brightgreen)](#-skills-disponíveis)
[![Licença](https://img.shields.io/badge/Licença-MIT-blue)](LICENSE)
[![GitHub Actions](https://img.shields.io/github/actions/workflow/status/KlebersonCollab/skills/generate-skills-artifacts.yml?branch=main&label=Build%20Artifacts)](https://github.com/KlebersonCollab/skills/actions)
[![Status](https://img.shields.io/badge/Status-Ativo-brightgreen)](#-roadmap)

---

## 📖 Sobre o Projeto

Este repositório é a **fonte da verdade** para todas as habilidades (skills) utilizadas por agentes de IA. Cada skill é um módulo independente, documentado e versionado, que pode ser integrado a qualquer agente compatível.

Utilizamos a metodologia **[SDD v2](sdd/)** (Spec-Driven Development) para garantir que cada nova funcionalidade seja rigorosamente especificada, planejada e verificada antes da implementação.

---

## 📦 Download & Uso Rápido

Geramos automaticamente pacotes pré-configurados para os principais agentes de IA. Basta baixar, descompactar na raiz do seu projeto e começar a usar:

| Agente | Artefato | Conteúdo |
|--------|----------|----------|
| **Claude AI** | [📥 claude-skills.zip](https://github.com/KlebersonCollab/skills/releases/latest/download/claude-skills.zip) | Pasta `.claude/` com `CLAUDE.md` e todas as skills. |
| **Gemini CLI** | [📥 gemini-skills.zip](https://github.com/KlebersonCollab/skills/releases/latest/download/gemini-skills.zip) | Pasta `.gemini/` com `GEMINI.md` e todas as skills. |
| **AI Agent (Generic)** | [📥 agent-skills.zip](https://github.com/KlebersonCollab/skills/releases/latest/download/agent-skills.zip) | Pasta `.agent/` com `AGENT.md` e todas as skills. |

> 💡 *Nota: Os links acima baixam automaticamente a versão mais recente da aba **Releases**.*

### 🚀 Como usar as Skills
1. Baixe o ZIP correspondente ao seu agente.
2. Extraia o conteúdo na raiz do diretório do seu projeto.
3. Certifique-se de que a pasta oculta (ex: `.claude/`) foi criada corretamente.
4. Seu agente de IA detectará automaticamente as novas capacidades.

---

## 🛠️ Ferramentas & Automação

O Hub inclui o **SDD CLI**, uma ferramenta Python que automatiza o ciclo de vida das skills e a gestão de conhecimento:

| Comando | Descrição |
|---------|-----------|
| `uv run sdd init <name>` | Inicializa uma nova feature com estrutura SDD completa. |
| `uv run sdd task <feat> <id>` | Marca o progresso de tarefas e sincroniza o estado global. |
| `uv run sdd graph` | Gera automaticamente o **Knowledge Map** (Mermaid) do projeto. |
| `uv run sdd sync` | Sincroniza mandatos globais em todos os agentes (.gemini, .claude, .agent). |

> 💡 *Nota: Requer [Python UV](https://docs.astral.sh/uv/) instalado.*

---

## 🛠️ Skills Disponíveis

| # | Skill | Descrição | Versão |
|---|-------|-----------|--------|
| 1 | **[SDD](sdd/)** | Spec-Driven Development — Workflow modular e adaptativo com **Persistent Memory Protocol**. | `1.3.1` |
| 2 | **[Skill Factory](skill-factory/)** | Core Framework para criação padronizada de novas skills com scaffolding e validação. | `1.1.0` |
| 3 | **[Python com UV](python-uv/)** | Desenvolvimento Python profissional com UV — gerenciador 10-100x mais rápido. | `2.5.0` |
| 4 | **[API Architect](api-architect/)** | Arquiteto de APIs — design interoperável e seguro (OpenAPI, GraphQL, tRPC). | `1.3.0` |
| 5 | **[Brainstorming](brainstorming/)** | Facilitador de design e resolução de problemas — exploração profunda e validação. | `1.1.0` |
| 6 | **[Architecture](architecture/)** | Arquiteto de Sistemas — design pragmático, trade-offs e registros de decisão (ADR). | `2.0.1` |
| 7 | **[Flutter com FVM](flutter-fvm/)** | Desenvolvimento Flutter profissional com FVM — gerenciador de versões SDK por projeto, testes avançados por camada arquitetural e segurança OWASP Mobile Top 10. | `1.1.0` |
| 8 | **[Azure DevOps](azure-devops/)** | Gerenciamento profissional de Boards, Repos, Pipelines e Artifacts no AzDO. | `1.1.0` |
| 9 | **[Clean Code Mentor](clean-code-mentor/)** | Mentor técnico e revisor de código focado em SOLID, YAGNI, DRY e KISS. | `1.0.0` |
| 10 | **[Observability Expert](observability-expert/)** | Especialista em SRE — Logs Estruturados, OpenTelemetry e Confiabilidade (SLI/SLO). | `1.0.0` |
| 11 | **[Onboarding Navigator](onboarding-navigator/)** | Guia interativo de cultura e engenharia para novos membros e projetos. | `1.1.0` |
| 12 | **[YouTube Transcript](youtube-transcript/)** | Extração de transcrições de vídeos com fallback para IA (Whisper) e limpeza automática. | `1.0.0` |
| 13 | **[Hardness Expert](hardness-expert/)** | Infraestrutura de suporte para estado, memória de longo prazo e orquestração de agentes via SDD. | `1.0.0` |
| 14 | **[Knowledge Architect](knowledge-architect/)** | Arquitetura de conhecimento local via grafos relacionais (Local GraphRAG). | `1.0.0` |

---

## 🏗️ Estrutura do Repositório

```
skills/
├── .github/workflows/           # Automação de CI/CD (GitHub Actions)
├── .specs/                      # Especificações do projeto (SDD)
│   ├── project/                 # Visão, Roadmap e Estado
│   └── codebase/                # Stack e Convenções
├── scripts/                     # Scripts de automação e distribuição
├── <skill>/                     # 🧩 Cada skill em seu diretório
│   ├── README.md                # Documentação detalhada
│   ├── SKILL.md                 # Definição técnica principal
│   └── CHANGELOG.md             # Histórico de versões
└── README.md                    # ← Você está aqui
```

---

## 📐 Como Criar uma Nova Skill

Utilize a skill **[Skill Factory](skill-factory/)** para criar novas skills de forma padronizada. O processo garante que cada skill seja gerada com a estrutura correta e passe por uma auditoria de qualidade antes de ser integrada ao hub.

> 📖 Consulte a [documentação do Skill Factory](skill-factory/) para detalhes.

---

## 🗺️ Roadmap & Estado do Projeto

Acompanhe a evolução do hub através dos documentos de planejamento:
- 🛤️ **[ROADMAP.md](.specs/project/ROADMAP.md)**: Visão de longo prazo e metas.
- 📊 **[STATE.md](.specs/project/STATE.md)**: Status atual e tarefas em andamento.

---

## 📄 Licença

Este projeto está licenciado sob a [Licença MIT](LICENSE).

---

<div align="center">

**Construído com 🧠 por [Kleberson Romero](https://github.com/KlebersonCollab)**

*Precision at scale. Rigor when needed, speed when possible.*

</div>
