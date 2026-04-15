# Architecture — Skill

Arquiteto de Sistemas — guia o agente a projetar sistemas escaláveis, simples e bem documentados, utilizando ADRs e análise rigorosa de trade-offs.

---

## 🚀 Visão Geral

Esta skill transforma o agente em um arquiteto de software pragmático. O foco não é apenas usar tecnologias de ponta, mas garantir que a solução seja a mais simples e mantível possível para o contexto dado. Ela formaliza o processo de tomada de decisão técnica, garantindo que o "porquê" por trás de cada escolha arquitetural seja registrado e validado.

### Capacidades Principais:
- **Descoberta Contextual**: Mapeamento de restrições técnicas, financeiras e de equipe.
- **Análise de Trade-offs**: Comparação objetiva entre abordagens concorrentes.
- **Seleção de Padrões**: Escolha fundamentada de estilos (Hexagonal, Microserviços, Monolito, etc.).
- **Governança via ADR**: Documentação durável de decisões arquiteturais.

---

## 📂 Estrutura da Skill

```text
architecture/
├── SKILL.md                 # Definição e workflow (v1.0.0)
├── README.md                # Este guia
├── CHANGELOG.md             # Histórico de versões
├── references/              # Conhecimento Teórico
│   ├── architectural-principles.md
│   ├── pattern-selection.md
│   └── adr-template.md
└── resources/               # ADRs gerados (sugerido)
```

---

## 🛠️ Como Usar

1. **Ative a skill**: `activate_skill architecture`
2. **Defina o Contexto**: "Quero desenhar a arquitetura para o novo sistema de processamento de pagamentos."
3. **Analise Alternativas**: O agente proporá diferentes caminhos (ex: Sync vs Async). Avalie os trade-offs listados.
4. **Gere o ADR**: Após escolher o caminho, o agente documentará a decisão no formato ADR para consulta futura.

---

## 🎯 Exemplos de Comandos

- *"Analise os trade-offs entre usar um banco de dados relacional e um NoSQL para este catálogo de produtos."*
- *"Como estruturar este projeto seguindo a arquitetura hexagonal para facilitar testes?"*
- *"Crie um ADR justificando a migração para uma arquitetura baseada em eventos."*
- *"Quais são as restrições arquiteturais que devemos considerar para um sistema que precisa de baixíssima latência?"*
