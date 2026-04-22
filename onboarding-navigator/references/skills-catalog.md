# Skills Catalog: AI Agent Hub

Este guia fornece um overview detalhado de todas as 21 habilidades disponíveis neste repositório, servindo como bússola para o Onboarding Navigator.

## 🗺️ Mapa Visual do Ecossistema de Skills

```mermaid
flowchart TD
    Start[Chegada ao Hub] --> Onboarding{onboarding-navigator}
    
    Onboarding --> Categoria{Categoria do Problema}
    
    Categoria --> Metodo[Metodologia & Criação]
    Categoria --> Arquitetura[Arquitetura & Design]
    Categoria --> DevOps[Ecosystems & DevOps]
    Categoria --> Automation[Automação & Utils]
    Categoria --> Navegacao[Navegação & Orchestration]
    
    Metodo --> SDD{sdd}
    Metodo --> Factory{skill-factory}
    Metodo --> Brainstorm{brainstorming}
    Metodo --> Harness{harness-expert}
    Metodo --> Knowledge{knowledge-architect}
    
    Arquitetura --> Arch{architecture}
    Arquitetura --> CleanCode{clean-code-mentor}
    Arquitetura --> API{api-architect}
    
    DevOps --> Python{python-uv}
    DevOps --> Flutter{flutter-fvm}
    DevOps --> AzDO{azure-devops}
    DevOps --> Observability{observability-expert}
    DevOps --> FastAPI{fastapi-expert}
    DevOps --> Django{django-expert}
    DevOps --> Golang{golang-expert}
    DevOps --> Frontend{frontend-expert}
    DevOps --> Git{git-workflow}

    Automation --> YouTube{youtube-transcript}
    Automation --> Scaffolding{scaffolding-expert}
    
    Navegacao --> Self{onboarding-navigator}
    Navegacao --> Facilitator{swarm-facilitator}
    
    style Start fill:#e1f5fe
    style Onboarding fill:#81c784
```

---

## 📚 Catálogo Completo de Skills (21 Total)

### 1. 🏗️ Core Frameworks (Metodologia e Criação)

| Skill | Versão | Propósito | Quando Invocá-la |
|-------|--------|-----------|------------------|
| **[SDD](sdd/)** | `1.4.0` | Spec-Driven Development. Modular workflow with PRD/RFC, BDD, and Mermaid Diagrams mandate. | **Sempre** que for iniciar uma implementação. |
| **[Skill Factory](skill-factory/)** | `1.1.0` | Core Framework para criação padronizada de novas skills com scaffolding, validação e registro automatizados. | Ao criar ou auditar uma habilidade no hub. |
| **[Brainstorming](brainstorming/)** | `1.1.0` | Facilitador de Brainstorming e Design — guia o agente a explorar problemas complexos. | Antes de qualquer especificação técnica. |
| **[Harness Expert](harness-expert/)** | `2.0.0` | Motor técnico para Harness Engineering (Sync, Rehydrate, Automação). | Quando precisar de suporte técnico do motor agêntico. |
| **[Knowledge Architect](knowledge-architect/)** | `1.0.0` | Arquitetura de conhecimento local via grafos relacionais (Local GraphRAG). | Para mapear relações complexas. |

### 2. 🎨 Architecture & Design (Qualidade e Estrutura)

| Skill | Versão | Propósito | Quando Invocá-la |
|-------|--------|-----------|------------------|
| **[Architecture](architecture/)** | `2.0.1` | Arquiteto de Sistemas — projeta sistemas escaláveis, resilientes e distribuídos via ADRs. | Ao desenhar a estrutura macro de um sistema. |
| **[Clean Code Mentor](clean-code-mentor/)** | `1.0.0` | Mentoria técnica e revisão de código com foco em SOLID, YAGNI, DRY e KISS. | Durante revisões de código ou refatorações. |
| **[API Architect](api-architect/)** | `1.3.0` | Arquiteto de APIs — projeta sistemas interoperáveis e seguros. | Ao projetar endpoints e integrações. |

### 3. ⚙️ Ecosystems & DevOps (Ambientes e Automação)

| Skill | Versão | Propósito | Quando Invocá-la |
|-------|--------|-----------|------------------|
| **[Python com UV](python-uv/)** | `2.6.0` | Desenvolvimento Python profissional com UV. | Em tarefas macro envolvendo Python. |
| **[FastAPI Expert](fastapi-expert/)** | `1.1.0` | Desenvolvimento avançado de APIs assíncronas com FastAPI e Pydantic. | Ao codificar rotas e schemas FastAPI. |
| **[Django Expert](django-expert/)** | `1.5.0` | Desenvolvimento robusto com Django, focado em arquitetura, segurança e TDD. | Ao trabalhar com aplicações Django. |
| **[Flutter com FVM](flutter-fvm/)** | `1.3.0` | Desenvolvimento Flutter profissional com FVM (Dart 3+, A11y, Performance). | Em qualquer tarefa envolvendo Flutter/Dart. |
| **[Azure DevOps](azure-devops/)** | `1.1.0` | Gerenciamento profissional do Azure DevOps (AzDO). | Para gerenciar tarefas e CI/CD no AzDO. |
| **[Observability Expert](observability-expert/)** | `1.0.0` | Especialista em SRE e Observabilidade (OTel, Logs Estruturados). | Ao garantir que um sistema é monitorável. |
| **[Golang Expert](golang-expert/)** | `1.1.0` | Excelência em Go — Performance, concorrência idiomática e ecossistema Samber. | Ao desenvolver sistemas de alta performance em Go. |
| **[Frontend Expert](frontend-expert/)** | `1.1.0` | Expert em interfaces modernas com React, Next.js e TailwindCSS v4. | Ao projetar ou implementar UIs modernas. |
| **[Git Workflow](git-workflow/)** | `1.0.0` | Padrões de fluxo de trabalho Git, estratégias de branching e integração com SDD. | Ao realizar commits, abrir PRs ou gerenciar branches. |

### 4. 🚀 Automation & Utils (Produtividade)

| Skill | Versão | Propósito | Quando Invocá-la |
|-------|--------|-----------|------------------|
| **[YouTube Transcript](youtube-transcript/)** | `1.0.0` | Automatizar a extração de transcrições de vídeos do YouTube. | Quando precisar de conteúdo de um vídeo. |
| **[Scaffolding Expert](scaffolding-expert/)** | `1.0.0` | Geração dinâmica de templates via CLI (copier/cookiecutter). | Para gerar um novo projeto do zero. |

### 5. 🧭 Navigation & Orchestration

| Skill | Versão | Propósito | Quando Invocá-la |
|-------|--------|-----------|------------------|
| **[Onboarding Navigator](onboarding-navigator/)** | `1.2.1` | Guia mestre do Hub de Skills. Fornece overview e mentoria. | **No início da sessão** para entender o hub. |
| **[Swarm Facilitator](swarm-facilitator/)** | `1.1.0` | Coreografia de equipes de agentes (Arquiteto, Dev, QA) e protocolos de handoff. | Em grandes épicos que requerem vários agentes coordenados. |

---

## 🧠 Matriz de Decisão: Qual Skill usar?

```mermaid
flowchart LR
    Problem[Problema Identificado] --> Decision{Tomada de Decisão}
    
    Decision --> NewApp["🚀 Criar app novo"]
    Decision --> CodeSmell["🤔 Código confuso"]
    Decision --> DeployFailed["🔴 Deploy falhou"]
    Decision --> NoIdea["❓ Não sei como começar"]
    Decision --> PythonTask["🐍 Trabalhar com Python"]
    Decision --> FlutterTask["📱 Trabalho com Flutter"]
    Decision --> Monitor["📊 Monitoramento/Logs"]
    Decision --> NewSkill["✨ Criar nova skill"]
    Decision --> YoutubeVid["🎥 Transcrever vídeo"]
    Decision --> Scaffold["📦 Gerar Template Base"]
    Decision --> Orchestrate["🤖 Projeto Gigante"]
    Decision --> MapKnowledge["🧠 Mapear Relações"]
    Decision --> SaveState["💾 Salvar Progresso"]
    Decision --> NewAPI["🔌 Desenhar Nova API"]
    Decision --> GoTask["🐹 Trabalhar com Go"]
    Decision --> FrontendTask["💻 Trabalho com Frontend"]
    
    NewApp --> Path1[onboarding-navigator → architecture → sdd]
    CodeSmell --> Path2[clean-code-mentor]
    DeployFailed --> Path3[azure-devops → observability-expert]
    NoIdea --> Path4[brainstorming]
    PythonTask --> Path5[python-uv → fastapi-expert / django-expert]
    FlutterTask --> Path6[flutter-fvm]
    Monitor --> Path7[observability-expert]
    NewSkill --> Path8[skill-factory]
    YoutubeVid --> Path9[youtube-transcript]
    Scaffold --> Path10[scaffolding-expert]
    Orchestrate --> Path11[swarm-facilitator]
    MapKnowledge --> Path12[knowledge-architect]
    SaveState --> Path13[harness-expert]
    NewAPI --> Path14[api-architect]
    GoTask --> Path15[golang-expert]
    FrontendTask --> Path16[frontend-expert]
    GitTask --> Path17[git-workflow]
    
    style Problem fill:#bbdefb
    style Decision fill:#ffcc80
```

---

## 📈 Estatísticas do Hub

- **Total de Skills**: 21
- **Skills de Metodologia**: 5
- **Skills de Arquitetura**: 3  
- **Skills de DevOps/Frameworks**: 9
- **Skills de Automação**: 2
- **Skills de Navegação/Orquestração**: 2
- **Última Atualização**: 22 de Abril de 2026
