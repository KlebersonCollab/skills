# Stack Tecnológica

## Linguagem & Formato
- **Markdown** (`.md`): Formato exclusivo para definição de skills, specs e documentação.

## Metodologia
- **SDD v1.3.1** (Spec-Driven Development): Workflow modular e adaptativo com integração de **PRD/RFC/BDD**, suporte a **Diagram-as-Code** e gerenciamento de estado persistente.

## Ferramentas de Visualização
- **Mermaid**: Padrão mandatório para diagramas arquiteturais e fluxos de processo integrados ao Markdown.

## Observabilidade & Telemetria
- **OpenTelemetry (OTel)**: Padrão para design de rastreamento distribuído e métricas.
- **Structured Logging (JSON)**: Padrão para emissão de logs legíveis por máquina.

## Ecossistemas de Automação
- **UV**: Gestão de projetos e pacotes Python.
- **FVM**: Gestão de versões do SDK Flutter.
- **Azure DevOps CLI**: Automação de infraestrutura e governança de plataforma.

## Distribuição & CI/CD
- **GitHub Actions**: Automação de pipelines para validação, geração de artefatos e distribuição.
- **Artefatos (.zip)**: Pacotes pré-configurados para múltiplos agentes (.claude, .gemini, .agent) gerados automaticamente.

## Controle de Versão
- **Git**: Versionamento do repositório.
- **GitHub**: Hospedagem e colaboração.

## Ferramentas de Desenvolvimento
- **Antigravity (AI Agent)**: Agente de IA utilizado para pair programming e execução das skills.

## Scaffolding
- **Skill Factory**: Core Framework (`skill-factory/`) para criação padronizada de novas skills.
- **Templates Embedded**: Os templates de scaffolding ficam embedded dentro das instruções das sub-skills (Markdown puro, sem ferramentas externas).
- **Validação Obrigatória**: Toda skill criada passa por auditoria automatizada antes do registro.

## Convenções
- Cada skill reside em seu próprio diretório (ex: `sdd/`).
- O arquivo principal de uma skill é nomeado `SKILL.md`.
- Sub-skills seguem o padrão `<nome-da-skill>-<sub-skill>.skill.md`.
- O `README.md` raiz deve ser atualizado a cada nova skill adicionada.
- Toda nova skill deve iniciar na versão `1.0.0`.
