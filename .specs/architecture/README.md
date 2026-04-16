# Architecture Decision Records (ADRs)

Este diretório contém Architecture Decision Records (ADRs) para o projeto, conforme exigido pela instrução #3 do CLAUDE.md.

## O que são ADRs?

ADRs são documentos que capturam decisões arquiteturais significativas, incluindo o contexto, alternativas consideradas, e consequências da decisão. Eles servem como:

- **Memória institucional**: Por que certas decisões foram tomadas
- **Guia para novos membros**: Contexto histórico das escolhas arquiteturais  
- **Base para decisões futuras**: Evita re-discutir decisões já tomadas

## Quando criar um ADR?

Crie um ADR quando a decisão:

1. **Altera a estrutura do sistema** (novos padrões, tecnologias críticas)
2. **Introduz novos padrões arquiteturais** (CQRS, Event-Driven, etc.)
3. **Afeta múltiplas skills ou componentes**
4. **Tem consequências de longo prazo** para manutenibilidade ou escalabilidade
5. **É requerido pela instrução #3 do CLAUDE.md**

## Estrutura dos ADRs

Cada ADR deve seguir o template `adr-template.md` e incluir:

1. **Título**: Frase imperativa curta (< 50 caracteres)
2. **Status**: Proposed, Accepted, Deprecated, ou Superseded
3. **Contexto**: Problema motivador e constraints
4. **Decisão**: O que foi decidido
5. **Consequências**: Impactos positivos e negativos
6. **Compliance & Verification**: Como validar a adoção
7. **Decisões Relacionadas**: Links para ADRs conectados

## ADRs Atuais

### [ADR-001: Arquitetura do Hub de Skills como Ecossistema Modular](adr-001-hub-skill-ecosystem.md)
**Status**: Accepted  
**Resumo**: Adoção de arquitetura modular com Onboarding Navigator como ponto único de entrada, catálogo centralizado dinâmico, e diagramação Mermaid obrigatória para as 11 skills do ecossistema.

## Processo de Aprovação

1. **Proposta**: Criar ADR com status "Proposed"
2. **Discussão**: Review por pelo menos 2 membros (skills `architecture` e `clean-code-mentor`)
3. **Revisão**: Atualizar ADR com feedback incorporado
4. **Aceitação**: Mudar status para "Accepted" após consenso
5. **Comunicação**: Notificar stakeholders e atualizar documentação relacionada

## Manutenção

- **Status Updates**: Atualizar status quando decisões mudarem
- **Superseding**: Quando um ADR é substituído, criar novo ADR referenciando o anterior
- **Arquivamento**: ADRs deprecated devem permanecer para histórico

## Integração com Skills

- **Architecture Skill**: Principal responsável pela criação e manutenção de ADRs
- **Onboarding Navigator**: Referencia ADRs relevantes no processo de onboarding
- **SDD**: ADRs são parte do processo de Spec-Driven Development
- **Clean Code Mentor**: Revisa ADRs por clareza e aderência a princípios

---

*Última atualização: 2026-04-15*