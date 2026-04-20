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
