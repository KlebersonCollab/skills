# Project Incremental Wisdom (Learnings)

## Automação & Scripting
- **Zipping Hidden Folders**: Ao criar ZIPs de pastas ocultas (como `.claude`), é fundamental entrar no diretório pai (`cd dist_staging`) antes de executar o comando `zip` para garantir que o ZIP contenha a pasta como sua raiz, facilitando o uso pelo usuário final.
- **Dynamic Skill Discovery**: O uso de `find . -maxdepth 2 -name "SKILL.md"` é uma estratégia robusta para listar skills sem precisar de um arquivo de manifesto manual, tornando o projeto auto-expansível.

## CI/CD (GitHub Actions)
- A separação de artefatos por nome no `upload-artifact` v4 melhora muito a experiência do usuário.
- **GitHub Releases vs Artifacts**: Para distribuição de software final (como os ZIPs de skills), Releases são preferíveis por serem permanentes e permitirem links estáticos (`/releases/latest/download/`), o que simplifica a integração e documentação.
- **Tag-Driven Releases**: Automatizar a criação de releases a partir do push de tags (`v*`) garante que apenas versões validadas sejam distribuídas aos usuários.

## Skill Management & Documentation
- **Centralized Version Tracking**: Manter um catálogo centralizado com versões atualizadas de todas as 16 skills é essencial para consistência do ecossistema.
- **Mermaid Diagrams for Navigation**: Diagramas visuais no catálogo de skills melhoram significativamente a compreensão do ecossistema e fluxos de trabalho.
- **Automated Version Collection**: Scripts para extrair versões e descrições automaticamente dos arquivos `SKILL.md` previnem desatualização do catálogo.
- **Onboarding Navigator as Living Documentation**: A skill de onboarding deve evoluir junto com o ecossistema, refletindo mudanças em tempo real.
## 2026-04-16 - Knowledge Graphing Strategy
- **Learning**: Implementing a modular 'knowledge-architect' is superior to bloat the 'sdd' or 'harness' skills.
- **Pattern**: Using Mermaid for relational mapping allows agents to visualize impact before coding.

## 2026-04-18 - Expert Skills & Global Sync Audit
- **Sync Blind Spots**: Foi descoberto que sincronizar apenas arquivos de texto (como mandatos globais) não é suficiente para o ecossistema. Novas skills criadas na raiz ficam "invisíveis" aos agentes se o script de sync não realizar o espelhamento completo de diretórios para as pastas ocultas (`.agent/skills`, etc).
- **Automation of Consistency**: Ferramentas simples de "cross-check" entre o README e o frontmatter das skills (como o `verify_versions.py`) são fundamentais para evitar que o repositório entre em um estado de entropia documental.
- **Skill Consolidation**: Integrar conhecimento de múltiplas fontes externas (GitHub) e modularizar em referências específicas (como `orm_performance.md` e `security.md`) torna a skill muito mais acionável do que um arquivo único gigante.
- **Mandatory Skill Governance**: Definir explicitamente em mandatos globais que o uso de skills é mandatório reduz a execução ad-hoc e garante que os agentes operem dentro dos trilhos de qualidade estabelecidos pelo Hub.
- **Workflow Automation with Makefile**: A criação de um Makefile centralizado reduz a carga cognitiva do desenvolvedor e garante que processos complexos (como o sync de mandatos e geração de dist) sejam executados de forma idêntica por qualquer agente ou humano.
- **Automated Knowledge Mapping**: O script `generate_knowledge_map.py` permite visualizar a topologia do projeto e dependências entre skills e features em tempo real, facilitando a análise de impacto em refatorações.

### [2026-04-19] Validação Rigorosa de Changelog
- **Ocorrência**: Falha no pipeline devido ao formato de data no CHANGELOG.md.
- **Solução**: O validador local exige estritamente o formato '## [X.Y.Z] - YYYY-MM-DD' (com hífen simples e espaços). O uso de travessão (—) ou formatos alternativos resulta em FAIL.

### [2026-04-20] Deterministic Enforcement & Context Inertia
- **Context Inertia**: LLMs tendem a ignorar instruções de "limpeza" (sync/memory) ao atingir o objetivo direto do usuário. A solução é transformar essas ações em **Exit Gates** imperativos e visuais no contexto.
- **Enforcement Layering**: O reforço funciona melhor em camadas:
    1. **Mandato Global** (Bootstrap/Exit Gate no topo do contexto).
    2. **Hook Local** (Prerequisites dentro da skill específica).
    3. **Visibilidade Total** (Compilação automática de todas as skills no .cursorrules).
- **Redundancy is Reliability**: Ter o SDD Hook tanto no mandato quanto dentro da skill de stack garante que o agente não pule etapas, mesmo se a tarefa for curta.
- **The "Brain Shell" Pattern**: Agrupar mandatos operacionais sob um subgraph de "Enforcement Shell" no Knowledge Map ajuda o agente a visualizar que a governança não é opcional, mas o esqueleto do sistema.

### [2026-04-20] Automated Distilling & Memory Resilience
- **Automated Knowledge Distiller**: Extrair metadados diretamente do frontmatter YAML das skills para gerar o mapa Mermaid reduz o erro humano e permite visualização dinâmica de versões e categorias. O uso de badges visuais (🛡️) para conformidade SDD cria uma camada de pressão social/técnica positiva para a manutenção da qualidade.
- **Auto-Fix as Resilience**: Em ecossistemas complexos, ferramentas que apenas detectam erros (Sensores) não são suficientes. Ter scripts que "reparam" o estado (como o `auto_fix_memory.py`) permite que o sistema recupere a integridade operacional sem intervenção humana, o que é crucial para agents autônomos.
- **The v2.0 Milestone**: A transição para um ecossistema determinístico (Bootstrap/Exit Gates) elimina a necessidade de "torcer" para que o agente siga os padrões. A governança agora está codificada no ciclo de vida da sessão.

### [2026-04-20] Code Audit Findings (Full Hub)
- **Duplication Crisis**: Skills estão duplicadas em ~5 localizações (`.agent/skills/`, `.gemini/skills/`, `.claude/skills/`, `dist_staging/`, `.claude/worktrees/`). Isso viola DRY e cria risco de inconsistência de versionamento.
- **Safe File Operations**: O script `sync_mandates.py` usa `shutil.rmtree` antes de `copytree`, criando risco de perda de dados se houver interrupção. Solução: usar TemporaryDirectory com transação segura.
- **Code Sharing Opportunity**: A lógica de walking de diretórios e listas de exclude está replicada em 4+ scripts. Criar `scripts/utils.py` compartilhado resolveria.

### [2026-04-21] Skill Enrichment & Content Density
- **Learning**: Uma skill sem exemplos práticos e referências externas é 50% menos eficaz para um agente de IA. A densidade de contexto (pasta `examples/` e `references/`) permite que o agente entenda não apenas *o que* fazer, mas *como* fazer seguindo padrões de mercado.
- **Pattern**: O uso de diretórios padronizados (`references/`, `resources/`, `examples/`) em todas as skills cria uma interface de navegação previsível para o agente, reduzindo a latência de busca por informação.

### [2026-04-22] Consolidação de Skills "Expert"
- **Learning**: Integrar conhecimentos de múltiplas fontes externas (Patterns, Security, TDD, Verification) em uma única skill "Expert" local reduz drasticamente a alternância de contexto e cria um "Standard de Ouro" robusto.
- **Pattern**: O workflow em 12 fases (do Setup ao Verification) garante que nenhum pilar técnico seja negligenciado durante o ciclo de vida do desenvolvimento.

### [2026-04-22] Modernização e Resiliência Mobile
- **Learning**: A integração de padrões modernos (Dart 3 Sealed Classes) em skills de Flutter não apenas melhora o código, mas força o agente a pensar em termos de **Exaustividade de Estado**, eliminando classes inteiras de bugs de UI.
- **Pattern**: A inclusão de **Acessibilidade Mandatória** e **Global Error Handling** no workflow core da skill garante que esses pilares "não-funcionais" sejam tratados como requisitos de primeira classe em todas as tarefas.

### [2026-04-22] Cohesion Audit & Validator Evolution
- **Learning (Validator Lag)**: Scripts de validação tendem a ficar defasados em relação à governança documental (ex: `MEMORY.md`). É vital que mudanças na política de governança sejam refletidas imediatamente no código do auditor (`validate_skills.py`).
- **Pattern (Mass Fixing)**: O uso de scripts de correção em massa (`fix_missing_hooks.py`) é a única forma escalável de manter a coesão em um Hub com 20+ skills, garantindo que o "Standard de Ouro" seja aplicado uniformemente sem erro humano.
- **Enforcement Success**: A auditoria de coesão manual seguida de automação elevou o Hub de 40% para 100% de conformidade com o mandado de `🔒 Prerequisites (Mandatory)`.
### [2026-04-22] GAN-style Harness Integration & Quality Enforcement
- **Learning**: Agentes de IA tendem a ser "otimistas patológicos" ao avaliar seu próprio trabalho. Separar o papel de **Generator** (quem implementa) e **Evaluator** (quem critica implacavelmente) via skill `harness-expert` cria uma tensão adversarial que elimina o "AI slop" e garante polimento de nível de produção.
- **Pattern (The Quality Loop)**: Introduzir uma rubrica de 4 pilares (**Design Quality, Craft, Originality, Functionality**) com uma "nota de corte" (Pass Threshold) transforma a Fase 4 do SDD (Review) em um ciclo iterativo, garantindo que o resultado final seja não apenas funcional, mas premium.
- **Implementation (Cross-Skill Fusion)**: Integrar essa lógica diretamente no framework `sdd` e no `onboarding-navigator` institucionaliza a qualidade como parte da infraestrutura, e não apenas como um desejo opcional.

### [2026-04-22] Idiomatic Go Patterns & Composition
- **Learning**: No Go, a simplicidade e a composição são os pilares da manutenibilidade. Adotar padrões como "Accept Interfaces, Return Structs" e "Zero Value Useful" reduz drasticamente a necessidade de mocks complexos e inicializações verbosas, tornando o sistema mais resiliente e fácil de entender.
- **Pattern (Interface Decoupling)**: Definir interfaces no lado do consumidor (Consumer-Defined Interfaces) é a chave para o desacoplamento real em Go, permitindo que diferentes implementações satisfaçam os requisitos sem que o provedor precise conhecer o contrato.
- **Enforcement (Concurrency Safety)**: A adoção do `errgroup` simplifica a orquestração de goroutines e garante que erros em fluxos paralelos não sejam perdidos, elevando o padrão de segurança de concorrência do Hub.
