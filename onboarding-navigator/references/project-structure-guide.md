# Project Structure Guide: The Source of Truth

Este guia explica a organização física e lógica do projeto para garantir uma navegação eficiente e o cumprimento do workflow SDD.

---

## 1. Hierarquia Global

```text
skills/
├── .specs/                      # O Coração do SDD
│   ├── project/                 # Visão, Roadmap e Memória Triade
│   ├── codebase/                # Stack, Convenções e Arquitetura Global
│   └── features/                # Especificações detalhadas por funcionalidade
├── .gemini/ | .claude/ | .agent/ # Configurações e Mandatos dos Agentes
├── <skill-name>/                # Diretórios de Habilidades (ex: python-uv)
│   ├── SKILL.md                 # Definição técnica
│   ├── references/              # Guias de suporte
│   └── examples/                # Amostras reais
├── README.md                    # Registro central de skills
└── .gitignore                   # Regras de versionamento (inclui Mandatos de Agentes)
```

## 2. O Papel da pasta `.specs/`

Aqui reside toda a inteligência estratégica do projeto. Nunca inicie uma mudança sem consultar este diretório.
- **`STATE.md`**: O que está acontecendo agora?
- **`MEMORY.md`**: O que aprendemos e não devemos esquecer?
- **`LEARNINGS.md`**: Soluções para bugs e padrões técnicos descobertos.

## 3. Diretórios de Agentes

Estes diretórios contêm os arquivos **`GEMINI.md`**, **`CLAUDE.md`** e **`AGENT.md`**.
- Eles são mandatórios e definem como o agente de IA deve se comportar.
- **Importante**: Estes arquivos `.md` são versionados (não estão no ignore), garantindo que todos os agentes sigam as mesmas regras de qualidade.

## 4. Onde encontrar cada coisa?

| Necessidade (O que você quer fazer?) | Localização / Skill Recomendada |
|-------------------------------------|---------------------------------|
| **Criar uma nova skill no Hub** | `skill-factory/` |
| **Orquestrar um Swarm de Agentes** | `multi-agent-orchestrator/` |
| **Consultar/Criar ADRs e Arquitetura Macro** | `architecture/` ou `.specs/codebase/` |
| **Desenhar Contratos e APIs (REST/GraphQL)** | `api-architect/` |
| **Garantir/Verificar regras de Clean Code** | `clean-code-mentor/` |
| **Salvar estado e Gerenciar Memória da Sessão**| `harness-expert/` ou `.specs/project/STATE.md` |
| **Mapear e Visualizar grafos de dependência** | `knowledge-architect/` ou `KNOWLEDGE-MAP.mermaid` |
| **Gerar projetos base / Boilerplates do zero** | `scaffolding-expert/` |
| **Configurar CI/CD, Boards ou Azure Pipelines**| `azure-devops/` |
| **Planejar próxima sprint ou Roadmap** | `.specs/project/ROADMAP.md` |

## 5. Dica de Navegação
Sempre comece sua sessão pedindo ao agente para "Reidratar o contexto lendo os arquivos de memória na pasta `.specs/`". Isso garante que ele saiba exatamente onde paramos.
