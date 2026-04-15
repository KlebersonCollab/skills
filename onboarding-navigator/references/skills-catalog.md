# Skills Catalog: AI Agent Hub

Este guia fornece um overview detalhado de todas as habilidades disponíveis neste repositório, servindo como bússola para o Onboarding Navigator.

---

## 1. Core Frameworks (Metodologia e Criação)

| Skill | Versão | Propósito | Quando Invocá-la |
|-------|--------|-----------|------------------|
| **[SDD](sdd/)** | `1.3.1` | Workflow de engenharia rigoroso (Spec-Driven). | **Sempre** que for iniciar uma implementação. |
| **[Skill Factory](skill-factory/)** | `1.1.0` | Automação de scaffolding e validação de novas skills. | Ao criar ou auditar uma habilidade no hub. |
| **[Brainstorming](brainstorming/)** | `1.1.0` | Design Thinking e exploração de problemas complexos. | Antes de qualquer especificação técnica. |

## 2. Architecture & Design (Qualidade e Estrutura)

| Skill | Versão | Propósito | Quando Invocá-la |
|-------|--------|-----------|------------------|
| **[Architecture](architecture/)** | `2.0.1` | Design de sistemas (CQRS, Events, ADRs, Mermaid). | Ao desenhar a estrutura macro de um sistema. |
| **[Clean Code Mentor](clean-code-mentor/)** | `1.0.0` | Guardião de SOLID, YAGNI, DRY e KISS. | Durante revisões de código ou refatorações. |
| **[API Architect](api-architect/)** | `1.3.0` | Design de APIs resilientes e contratos (OpenAPI). | Ao projetar endpoints e integrações. |

## 3. Ecosystems & DevOps (Ambientes e Automação)

| Skill | Versão | Propósito | Quando Invocá-la |
|-------|--------|-----------|------------------|
| **[Python com UV](python-uv/)** | `2.5.0` | Gestão ultra-rápida de Python e pacotes. | Em qualquer tarefa envolvendo Python. |
| **[Flutter com FVM](flutter-fvm/)** | `1.0.0` | Gestão de múltiplas versões do SDK Flutter. | Em qualquer tarefa envolvendo Flutter/Dart. |
| **[Azure DevOps](azure-devops/)** | `1.0.0` | Automação de Boards, Repos e Pipelines. | Para gerenciar tarefas e CI/CD no AzDO. |
| **[Observability Expert](observability-expert/)** | `1.0.0` | Especialista em SRE (Logs, Metrics, Traces). | Ao garantir que um sistema é monitorável. |

## 4. Navigation & Mentorship

| Skill | Versão | Propósito | Quando Invocá-la |
|-------|--------|-----------|------------------|
| **[Onboarding Navigator](onboarding-navigator/)** | `1.0.0` | Guia do hub e mentoria de cultura. | **No início da sessão** para entender o hub. |

---

## Matriz de Decisão: Qual Skill usar?

1.  **"Preciso criar um app novo"**: Invoque `onboarding-navigator` -> `architecture` -> `sdd`.
2.  **"O código está confuso"**: Invoque `clean-code-mentor`.
3.  **"O deploy falhou"**: Invoque `azure-devops` -> `observability-expert`.
4.  **"Não sei como começar"**: Invoque `brainstorming`.
