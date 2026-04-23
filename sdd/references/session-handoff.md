# Transição de Sessão (Handoff)

**Gatilhos:** "Pausar trabalho", "Encerrar sessão", "Continuar trabalho", "Retomar"

## Pausando o Trabalho
Antes de encerrar o contexto de uma sessão, garanta que todo o estado atual seja salvo com precisão.
1. **Atualize o STATE.md:** Anote a tarefa ativa atual, os últimos blockers enfrentados e os próximos passos imediatos.
2. **Commit Atômico:** Certifique-se de que o código em andamento esteja commitado na branch de feature (use `[WIP]` se necessário).
3. **Check de Tarefas:** Marque em `tasks.md` exatamente onde a execução parou.

## Retomando o Trabalho
Ao iniciar ou retomar uma sessão:
1. Re-hidrate o contexto lendo primeiramente o `STATE.md`.
2. Identifique a `Active Task`.
3. Valide o estado atual da feature branch no Git.
4. Revise eventuais arquivos em `LEARNINGS.md` ou regras novas estabelecidas recentemente.

## Regras de Continuidade
A premissa principal do Handoff é que a próxima sessão ou o próximo agente deve ter *exatamente* o mesmo entendimento e direcionamento de como a tarefa foi deixada. A memória de curto prazo se perde, então o arquivo `STATE.md` se torna o cérebro persistente da operação.
