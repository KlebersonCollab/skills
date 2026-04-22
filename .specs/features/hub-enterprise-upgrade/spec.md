# 📄 Especifição: Hub Enterprise Upgrade

**Feature**: "Enterprise Grade" Skills Hub (Test Suite, Distiller, DevSecOps, UI e CLI)
**Responsável**: Agent Swarm
**Status**: Specification

## 1. Visão Geral (PRD)
O Hub cresceu na direção de abranger 21 skills técnicas especializadas. O foco agora é elevar a "Robustez Interna" e "Usabilidade" do ecossistema, movendo as capacidades atuais de simples scripts e pastas para um sistema autogerenciado, auto-documentado e com um portal visual acessível para a comunidade/outros agentes.

## 2. Objetivos Principais
- **A. Hub Test Suite**: Blindar o código da pasta `/scripts/` garantindo estabilidade nas automações usando `pytest`.
- **B. Automated Knowledge Distiller**: Automatizar o grafo Mermaid para não necessitar mais de intervenção humana quando criar novas skills.
- **C. DevSecOps Skill**: Garantir auditorias de dependências e código.
- **D. Marketplace UI Dashboard**: Um portal web estático gerado a partir do repositório.
- **E. Skill Installer CLI**: Desacoplar o download do repositório inteiro e permitir instalar `skills` de forma isolada.

## 3. Escopo e Limites
- A implementação não deve alterar as skills já validadas na V1 (apenas atualização de versão).
- Deve manter total compatibilidade retroativa com `.claude`, `.gemini` e `.agent`.

## 4. Requisitos Não Funcionais (NFRs)
- Testes: O coverage dos scripts locais deve ser > 80%.
- Performance: O Distiller deve gerar o mapa Mermaid em menos de 2s.

## 5. Estratégia BDD (Comportamento)
- **Dado** que há novas skills criadas, **Quando** for commitar, **Então** o Hook ou Action roda o Distiller e atualiza o KNOWLEDGE-MAP.mermaid automaticamente.
