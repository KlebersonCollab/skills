# Brainstorming — Skill

Facilitador de Brainstorming e Design — guia o agente a explorar problemas complexos, gerar ideias divergentes e convergir para especificações sólidas antes de qualquer implementação.

---

## 🚀 Visão Geral

Esta skill transforma o agente em um facilitador de design thinking. O objetivo não é codificar rápido, mas **pensar devagar** para garantir que a solução construída seja a correta. Ela utiliza técnicas consagradas para quebrar bloqueios criativos e validar premissas antes de gastar recursos em implementação.

### Capacidades Principais:
- **Discovery Disciplinado**: Identificação da causa raiz via "5 Whys".
- **Geração Divergente**: Exploração de múltiplas arquiteturas/abordagens (SCAMPER).
- **Convergência Estratégica**: Avaliação de trade-offs e tomada de decisão fundamentada.
- **Understanding Lock**: Portão de segurança para garantir alinhamento total com o usuário.

---

## 📂 Estrutura da Skill

```text
brainstorming/
├── SKILL.md                 # Definição e workflow (v1.0.0)
├── README.md                # Este guia
├── CHANGELOG.md             # Histórico de versões
├── references/              # Conhecimento Teórico
│   ├── brainstorming-techniques.md
│   └── decision-frameworks.md
└── resources/               # Templates (opcional)
```

---

## 🛠️ Como Usar

1. **Ative a skill**: `activate_skill brainstorming`
2. **Inicie o Discovery**: "Gostaria de fazer um brainstorming sobre a nova arquitetura de notificações."
3. **Siga o Facilitador**: O agente fará perguntas uma por vez para mapear o problema e as premissas.
4. **Valide o Understanding Lock**: Antes de prosseguir para a solução, confirme o resumo de entendimento enviado pelo agente.

---

## 🎯 Exemplos de Comandos

- *"Vamos fazer um brainstorming sobre como reduzir a latência da nossa API de busca."*
- *"Use a técnica SCAMPER para propor novas funcionalidades para o nosso dashboard."*
- *"Quais são os trade-offs entre usar um banco relacional ou NoSQL para este caso de uso?"*
