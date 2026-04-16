# Harness Expert

> "Se você não é o modelo, então você é o harness." — Infraestrutura de suporte para transformar LLMs em agentes operacionais.

[![Versão](https://img.shields.io/badge/Versão-1.0.0-blue)](#changelog)
[![Sub-skills](https://img.shields.io/badge/Sub--skills-2-brightgreen)](#-sub-skills)

---

## 📖 Visão Geral

O **Harness Expert** é a skill que orquestra a infraestrutura necessária para que um modelo de linguagem (LLM) se comporte como um agente confiável, seguro e persistente. Inspirado nos conceitos de "Harness Engineering", ele foca em garantir que o agente tenha acesso às ferramentas certas, mantenha seu estado entre sessões e valide rigorosamente suas próprias saídas.

No ecossistema deste Hub, o Harness Expert atua como o "zelador" do workflow SDD (Spec-Driven Development), garantindo que a memória de longo prazo do projeto (specs, logs e estados) esteja sempre sincronizada com as ações do agente.

---

## ⚙️ Como Usar

Esta skill é ativada automaticamente sempre que uma tarefa de implementação é iniciada, ou pode ser invocada manualmente para realizar auditorias de estado ou sincronização de progresso.

**Cenários de Aplicação:**
1.  **Início de Sessão**: Rehidratar o contexto lendo os arquivos de especificação.
2.  **Durante a Implementação**: Sincronizar sub-tarefas concluídas no `tasks.md`.
3.  **Finalização**: Gerar relatórios de validação e persistir lições aprendidas.

---

## 🧩 Sub-skills

| Sub-skill | Arquivo | Responsabilidade |
|-----------|---------|------------------|
| **Harness Sync** | [harness-expert-sync.skill.md](harness-expert-sync.skill.md) | Sincroniza o progresso das tarefas e registra aprendizados no File System. |
| **Harness Rehydrate** | [harness-expert-rehydrate.skill.md](harness-expert-rehydrate.skill.md) | Reconstrói o contexto operacional lendo specs e estados persistidos. |
| **Harness Compress** | [harness-expert-compress.skill.md](harness-expert-compress.skill.md) | Compacta o histórico de tarefas concluídas no STATE.md para eficiência de tokens. |

---

## 📝 Changelog

Consulte o [CHANGELOG.md](CHANGELOG.md) para o histórico completo de versões.
