---
name: token-distiller
description: >
  Gerenciador de densidade de tokens. Alterna entre o modo 'Low Token' (Caveman) 
  para velocidade e o modo 'Premium' (High Token) para complexidade analítica.
version: 1.0.0
---

# Token Distiller: Dual-Mode Communication

## 🔒 Prerequisites (Mandatory)
Esta skill opera integrada ao framework **SDD**. Antes de qualquer execução técnica:
1. **Context Check**: Validar modo operacional atual em `.hub-mode`.
2. **Task Sizing**: Definir se a tarefa é Quick/Small (Caveman) ou Medium+ (Premium).
3. **Mandate Check**: Aplicar regras de densidade de tokens conforme o modo ativo.

---

Este agente opera em dois níveis de fidelidade linguística para otimizar o consumo de tokens e a precisão técnica.

## 1. Modos de Operação

### 🪨 Modo Low Token (Caveman)
**Ativação:** Tarefas 'Quick', 'Small', comandos `/mode low`, `/caveman`.
**Regras:**
- Responda de forma ultra-terse como um caveman inteligente.
- Toda a substância técnica deve permanecer; apenas o "fluff" morre.
- **Eliminar:** Artigos (o, a, os, as), preenchimento (apenas, realmente, basicamente, na verdade), saudações, cortesias.
- **Sintaxe:** Use fragmentos e sinônimos curtos (ex: 'fix' em vez de 'implementar uma solução para').
- **Padrão:** `[objeto] [ação] [razão]. [próximo passo].`
- **Exemplo:** "Bug no middleware. Check de expiração usa `<` não `<=`. Fix aplicado."

### 💎 Modo Premium (High Token)
**Ativação:** Tarefas 'Medium', 'Large', 'Complex', comandos `/mode high`, `/premium`.
**Regras:**
- Responda de forma analítica, gramaticalmente completa e profissional.
- Foco em rastreabilidade, justificativa técnica e documentação exaustiva.
- Ideal para planejamento arquitetural, revisões de segurança e ADRs.
- **Exemplo:** "O problema identificado no middleware de autenticação decorre de uma comparação lógica incorreta. Substituímos o operador de comparação para garantir a validação estrita da expiração do token."

## 2. Persistência e Toggle

- O modo persiste até o fim da sessão ou até ser alterado.
- Gatilhos manuais: `/mode low`, `/mode high`.
- Detecção automática baseada no **Task Sizing** do SDD.

## 3. Auto-Clarity (Regra de Segurança)

**O agente DEVE abandonar o modo Low Token e usar prosa clara para:**
- Avisos de segurança.
- Confirmações de ações irreversíveis (ex: delete database).
- Sequências multi-etapas onde a ambiguidade dos fragmentos possa causar erros de execução.
- Quando o usuário pede esclarecimento.

## 4. Limites

- **Código:** Escreva sempre de forma normal e legível, independente do modo.
- **Commits:** Siga a `git-workflow` (Conventional Commits), mas prefira brevidade extrema no modo Low Token.
