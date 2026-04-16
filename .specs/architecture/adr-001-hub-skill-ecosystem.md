# ADR-001: Arquitetura do Hub de Skills como Ecossistema Modular

## Status

Accepted

## Context

O projeto evoluiu de um conjunto de scripts para um ecossistema organizado de 11 skills especializadas. Com o crescimento, surgiram desafios de:
1. **Descoberta de funcionalidades**: Difícil saber qual skill usar para cada problema
2. **Consistência de versões**: Catálogo desatualizado vs. realidade do código
3. **Onboarding complexo**: Novos usuários não entendem a hierarquia e relacionamentos
4. **Manutenção descentralizada**: Cada skill evolui independentemente

A instrução #3 do CLAUDE.md exige ADRs para decisões arquiteturais significativas.

## Decision

Adotar uma arquitetura de **ecossistema modular** com os seguintes componentes:

1. **Onboarding Navigator como Ponto Único de Entrada**: Toda sessão deve iniciar com esta skill
2. **Catálogo Centralizado Dinâmico**: Única fonte de verdade para todas as 11 skills
3. **Diagramação Mermaid Obrigatória**: Visualização do ecossistema e fluxos de decisão
4. **Versionamento Automatizado**: Scripts para extrair versões dos arquivos SKILL.md
5. **Categorização Hierárquica**: Organização em 4 categorias funcionais:
   - Core Frameworks (Metodologia)
   - Architecture & Design (Qualidade)
   - Ecosystems & DevOps (Ambientes)
   - Navigation & Mentorship (Orientação)

## Consequences

### Positivas
- **Redução do Time-to-Value**: Novos usuários entendem o ecossistema em minutos
- **Consistência Garantida**: Versões sempre atualizadas via automação
- **Descoberta Intuitiva**: Diagramas visuais facilitam navegação
- **Escalabilidade**: Nova skill = atualização automática do catálogo
- **Alinhamento Cultural**: Onboarding Navigator reforça padrões KISS/SDD

### Negativas (e Mitigações)
- **Complexidade Inicial**: Mitigado por diagramas e exemplos visuais
- **Manutenção do Catálogo**: Mitigado por scripts de atualização automática
- **Dependência do Onboarding Navigator**: Mitigado por documentação redundante em READMEs

### Riscos Técnicos
1. **Desatualização do Catálogo**: Mitigado por script de coleta de versões
2. **Diagramas Complexos**: Mitigado por templates Mermaid reutilizáveis
3. **Carga Cognitiva**: Mitigado por categorização hierárquica clara

## Compliance & Verification

### Métricas de Sucesso
- [x] Catálogo com todas as 11 skills documentadas
- [x] Diagramas Mermaid para ecossistema e fluxos de decisão
- [x] Versões extraídas automaticamente dos arquivos SKILL.md
- [x] Referência cruzada com STATE.md, MEMORY.md, LEARNINGS.md
- [ ] Testes automatizados de consistência do catálogo (futuro)

### Verificação Contínua
1. **Pré-commit Hook**: Validar que novas skills são adicionadas ao catálogo
2. **CI/CD Check**: Validar consistência entre versões declaradas e reais
3. **Review Obrigatório**: Mudanças no Onboarding Navigator exigem review de 2 pessoas

## Related Decisions

- **CLAUDE.md Instrução #3**: Exigência de ADRs para decisões arquiteturais
- **ADR-000**: Template de ADR para decisões futuras
- **Skill Factory v1.1.0**: Framework para criação padronizada de novas skills

---

*Created: 2026-04-15*  
*Last Updated: 2026-04-15*  
*Deciders: Sistema de Skills, Onboarding Navigator, Architecture Skill*