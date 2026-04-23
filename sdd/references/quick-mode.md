# Quick Mode (Modo Rápido)

**Gatilhos:** "Correção rápida", "Ajuste rápido", "Bugfix trivial"

Para tarefas que se enquadram na categoria **Quick** (correção de bugs simples, configuração, alterações em até 3 arquivos e que podem ser descritas em uma frase).

## Fluxo
1. **Pule as fases de Specify, Design e Tasks.** A burocracia é desnecessária.
2. **Execute (Implement & Verify):** Vá direto para a correção do código.
3. Garanta que os testes continuam passando.
4. Confirme que não houve regressões visuais ou lógicas.
5. **Comite diretamente** (ou utilize uma branch de curta vida sem criar uma especificação inteira de feature).

## Válvula de Segurança
Se, ao iniciar a tarefa em *Quick Mode*, você perceber que a implementação exigirá alterações estruturais, adição de novas bibliotecas ou modificar mais de 3 arquivos substancialmente, **PARE**.

Neste caso, declare ao usuário:
> *"Esta tarefa parecia rápida, mas envolve mudanças arquiteturais profundas. Abortando o Quick Mode e elevando para o escopo Small/Medium."*
Em seguida, inicie as etapas convencionais do framework SDD (gerar spec, tasks, etc).
