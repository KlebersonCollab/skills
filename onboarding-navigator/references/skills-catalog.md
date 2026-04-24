# Skills Catalog: AI Agent Hub

This guide provides a detailed overview of all 25 abilities available in this repository, serving as a compass for the Onboarding Navigator.

## 🗺️ Visual Map of the Skills Ecosystem

```mermaid
flowchart TD
    Start[Arrival at the Hub] --> Onboarding{onboarding-navigator}
    
    Onboarding --> Category{Problem Category}
    
    Category --> Method[Methodology & Creation]
    Category --> Architecture[Architecture & Design]
    Category --> DevOps[Ecosystems & DevOps]
    Category --> Automation[Automation & Utils]
    Category --> Navigation[Navigation & Orchestration]
    
    Method --> SDD{sdd}
    Method --> Factory{skill-factory}
    Method --> Brainstorm{brainstorming}
    Method --> Harness{harness-expert}
    Method --> Knowledge{knowledge-architect}
    Method --> Distiller{token-distiller}
    
    Architecture --> Arch{architecture}
    Architecture --> CleanCode{clean-code-mentor}
    Architecture --> API{api-architect}
    
    DevOps --> Python{python-uv}
    DevOps --> Flutter{flutter-fvm}
    DevOps --> AzDO{azure-devops}
    DevOps --> Observability{observability-expert}
    DevOps --> FastAPI{fastapi-expert}
    DevOps --> Django{django-expert}
    DevOps --> Golang{golang-expert}
    DevOps --> Frontend{frontend-expert}
    DevOps --> Git{git-workflow}
    DevOps --> GoTest{golang-testing-expert}
    DevOps --> Benchmark{benchmark-expert}
    DevOps --> DevSecOps{devsecops-expert}

    Automation --> YouTube{youtube-transcript}
    Automation --> Scaffolding{scaffolding-expert}
    
    Navigation --> Self{onboarding-navigator}
    Navigation --> Facilitator{swarm-facilitator}
    
    style Start fill:#e1f5fe
    style Onboarding fill:#81c784
```

---

## 📚 Complete Skills Catalog (25 Total)

### 1. 🏗️ Core Frameworks (Methodology and Creation)

| Skill | Version | Purpose | When to Invoke |
|-------|--------|-----------|------------------|
| **[SDD](../../sdd/)** | `1.5.0` | Spec-Driven Development. Modular workflow with PRD/RFC, BDD, and Mermaid Diagrams mandate. | **Always** when starting an implementation. |
| **[Skill Factory](../../skill-factory/)** | `1.1.0` | Core Framework for standardized creation of new skills with automated scaffolding, validation, and registration. | When creating or auditing an ability in the hub. |
| **[Brainstorming](../../brainstorming/)** | `1.1.0` | Brainstorming and Design Facilitator — guides the agent to explore complex problems. | Before any technical specification. |
| **[Harness Expert](../../harness-expert/)** | `2.0.0` | Technical engine for Harness Engineering (Sync, Rehydrate, Automation). | When technical support from the agentic engine is needed. |
| **[Knowledge Architect](../../knowledge-architect/)** | `1.0.0` | Local knowledge architecture via relational graphs (Local GraphRAG). | To map complex relationships. |
| **[Token Distiller](../../token-distiller/)** | `1.0.0` | Token density manager and compression modes (Caveman/Premium). | To optimize token consumption in long tasks. |

### 2. 🎨 Architecture & Design (Quality and Structure)

| Skill | Version | Purpose | When to Invoke |
|-------|--------|-----------|------------------|
| **[Architecture](../../architecture/)** | `2.0.1` | Systems Architect — designs scalable, resilient, and distributed systems via ADRs. | When designing a system's macro structure. |
| **[Clean Code Mentor](../../clean-code-mentor/)** | `1.0.0` | Technical mentoring and code review focused on SOLID, YAGNI, DRY, and KISS. | During code reviews or refactorings. |
| **[API Architect](../../api-architect/)** | `1.3.0` | API Architect — designs interoperable and secure systems. | When designing endpoints and integrations. |

### 3. ⚙️ Ecosystems & DevOps (Environments and Automation)

| Skill | Version | Purpose | When to Invoke |
|-------|--------|-----------|------------------|
| **[Python with UV](../../python-uv/)** | `3.0.0` | Professional Python development with UV. | In macro tasks involving Python. |
| **[FastAPI Expert](../../fastapi-expert/)** | `1.1.0` | Advanced asynchronous API development with FastAPI and Pydantic. | When coding FastAPI routes and schemas. |
| **[Django Expert](../../django-expert/)** | `1.5.0` | Robust development with Django, focused on architecture, security, and TDD. | When working with Django applications. |
| **[Flutter with FVM](../../flutter-fvm/)** | `1.3.0` | Professional Flutter development with FVM (Dart 3+, A11y, Performance). | In any task involving Flutter/Dart. |
| **[Azure DevOps](../../azure-devops/)** | `1.1.0` | Professional Azure DevOps (AzDO) management. | To manage tasks and CI/CD in AzDO. |
| **[Observability Expert](../../observability-expert/)** | `1.0.0` | SRE and Observability specialist (OTel, Structured Logs). | When ensuring a system is monitorable. |
| **[Golang Expert](../../golang-expert/)** | `1.1.0` | Go Excellence — Performance, idiomatic concurrency, and Samber ecosystem. | When developing high-performance systems in Go. |
| **[Frontend Expert](../../frontend-expert/)** | `1.1.0` | Expert in modern interfaces with React, Next.js, and TailwindCSS v4. | When designing or implementing modern UIs. |
| **[Git Workflow](../../git-workflow/)** | `1.0.0` | Git workflow patterns, branching strategies, and SDD integration. | When performing commits, opening PRs, or managing branches. |
| **[Golang Testing Expert](../../golang-testing-expert/)** | `1.0.0` | QA expert for Go — TDD, Table-Driven Tests, Benchmarks, and Fuzzing. | When implementing Go test suites. |
| **[Benchmark Expert](../../benchmark-expert/)** | `1.0.0` | Expert Skill for performance baseline measurement, regression detection, and stack comparison. | When validating performance impact of changes or PRs. |
| **[DevSecOps Expert](../../devsecops-expert/)** | `1.0.0` | Static quality auditing, SAST/DAST security, and Hardening. | When performing security and quality audits on code. |

### 4. 🚀 Automation & Utils (Productivity)

| Skill | Version | Purpose | When to Invoke |
|-------|--------|-----------|------------------|
| **[YouTube Transcript](../../youtube-transcript/)** | `1.0.0` | Automate YouTube video transcript extraction. | When video content is needed. |
| **[Scaffolding Expert](../../scaffolding-expert/)** | `1.0.0` | Dynamic template generation via CLI (copier/cookiecutter). | To generate a new project from scratch. |

### 5. 🧭 Navigation & Orchestration

| Skill | Version | Purpose | When to Invoke |
|-------|--------|-----------|------------------|
| **[Onboarding Navigator](../../onboarding-navigator/)** | `1.2.1` | Skills Hub master guide. Provides overview and mentorship. | **At session start** to understand the hub. |
| **[Swarm Facilitator](../../swarm-facilitator/)** | `1.1.0` | Agent team choreography (Architect, Dev, QA) and handoff protocols. | In large epics requiring multiple coordinated agents. |

---

## 🧠 Decision Matrix: Which Skill to use?

```mermaid
flowchart LR
    Problem[Problem Identified] --> Decision{Decision Making}
    
    Decision --> NewApp["🚀 Create new app"]
    Decision --> CodeSmell["🤔 Confusing code"]
    Decision --> DeployFailed["🔴 Deploy failed"]
    Decision --> NoIdea["❓ Don't know how to start"]
    Decision --> PythonTask["🐍 Work with Python"]
    Decision --> FlutterTask["📱 Work with Flutter"]
    Decision --> Monitor["📊 Monitoring/Logs"]
    Decision --> NewSkill["✨ Create new skill"]
    Decision --> YoutubeVid["🎥 Transcribe video"]
    Decision --> Scaffold["📦 Generate Base Template"]
    Decision --> Orchestrate["🤖 Giant Project"]
    Decision --> MapKnowledge["🧠 Map Relationships"]
    Decision --> SaveState["💾 Save Progress"]
    Decision --> NewAPI["🔌 Design New API"]
    Decision --> GoTask["🐹 Work with Go"]
    Decision --> FrontendTask["💻 Work with Frontend"]
    
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
    GoTestTask --> Path18[golang-testing-expert]
    PerformanceTask --> Path19[benchmark-expert]
    
    style Problem fill:#bbdefb
    style Decision fill:#ffcc80
```

---

## 📈 Hub Statistics

- **Total Skills**: 25
- **Methodology Skills**: 6
- **Architecture Skills**: 3  
- **DevOps/Framework Skills**: 12
- **Automation Skills**: 2
- **Navigation/Orchestration Skills**: 2
- **Last Update**: April 23, 2026
