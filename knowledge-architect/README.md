# Knowledge Architect

> Arquitetura de conhecimento local via grafos relacionais (Local GraphRAG). Transforma documentação estática em um mapa de conexões para navegação inteligente de contexto.

[![Versão](https://img.shields.io/badge/Versão-1.0.0-blue)](#changelog)
[![Sub-skills](https://img.shields.io/badge/Sub--skills-0-brightgreen)](#)

---

## 📖 Visão Geral

O **Knowledge Architect** é a inteligência por trás da organização semântica do Hub. Enquanto outras skills focam em *fazer* (implementar, testar, orquestrar), esta skill foca em *entender* as relações estruturais do projeto.

Inspirado em conceitos de **GraphRAG**, ela permite que o agente construa um "segundo cérebro" para o repositório, mapeando como requisitos de negócio se transformam em decisões técnicas e, finalmente, em código e testes. Isso resolve o problema de "visão de túnel" onde o agente entende um arquivo isolado, mas esquece o impacto sistêmico de suas alterações.

---

## ⚙️ Como Usar

Esta skill deve ser invocada em cenários de:
1.  **Onboarding de Codebase**: Para mapear rapidamente as dependências de um projeto legado.
2.  **Análise de Impacto**: Antes de grandes refatorações, para ver quem será afetado.
3.  **Discovery de Feature (SDD)**: Para conectar novos requisitos ao grafo de conhecimento existente.
4.  **Rehydration Relacional (Harness)**: Para carregar o contexto "vizinho" de forma inteligente.

**Exemplo de Invocação:**
*"Use a skill knowledge-architect para mapear as relações entre a nova feature de autenticação e os módulos de banco de dados existentes, gerando um KNOWLEDGE-MAP.mermaid."*

---

## 📝 Changelog

Consulte o [CHANGELOG.md](CHANGELOG.md) para o histórico completo de versões.
