# Memória de Longo Prazo (LTM) no Harness

O `harness-expert` é responsável por garantir que o agente não sofra de "amnésia" entre sessões, mantendo o estado operacional sincronizado.

## 🧠 Conceitos de Memória

### 1. Estado Operacional (`STATE.md`)
É o "ponteiro" atual. Contém o progresso das tarefas, o que está pendente e o foco da sessão atual. É a memória de curto prazo estruturada.

### 2. Memória Persistente (`MEMORY.md`)
Armazena fatos, decisões tomadas e o contexto do projeto que não muda com frequência. Ex: "O projeto usa Postgres 15", "O padrão de logs é JSON".

### 3. Lições Aprendidas (`LEARNINGS.md`)
Registra erros cometidos, correções e "eureka moments". Serve para evitar a repetição de falhas passadas e melhorar a eficiência do agente.

## 🔄 Fluxo de Sincronização

1. **Reidratação**: No início da sessão, o agente lê esses arquivos via `harness-expert-rehydrate`.
2. **Atualização**: Durante a execução, o agente atualiza o `STATE.md` conforme as tasks são concluídas.
3. **Persistência Final**: Ao final, o `harness-expert-sync` garante que as mudanças sejam salvas no disco.

## 💡 Por que isso importa?
Sem LTM, o agente perderia o progresso a cada nova interação, forçando o usuário a repetir instruções e aumentando o custo de tokens e tempo.
