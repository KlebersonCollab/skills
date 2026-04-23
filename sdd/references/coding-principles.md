# Princípios de Codificação

Viés comportamental, não apenas um checklist. Leia antes de cada implementação.

---

## Antes de Codificar
- **Explicit assumições:** Declare suas premissas. Se incerto, pergunte.
- **Interpretações múltiplas:** Se houver várias formas de interpretar a tarefa, apresente-as. Não escolha silenciosamente.
- **Simplicidade:** Existe uma abordagem mais simples? Diga. Questione quando necessário.
- **Incerteza:** Se algo está confuso, PARE. Nomeie o que está confuso e peça esclarecimentos.

---

## Durante a Implementação

### Simplicidade
- Sem features além do que foi solicitado (YAGNI).
- Sem abstrações para códigos de uso único.
- Sem "flexibilidade" ou "configuração" que não foi pedida.
- 200 linhas de código que poderiam ser 50? Refatore para simplificar.

### Alterações Cirúrgicas
- Não "melhore" código adjacente, comentários ou formatações não relacionados à tarefa.
- Não refatore o que não está quebrado.
- Siga o estilo de código existente no repositório.
- Remova APENAS imports/variáveis/funções que as SUAS mudanças deixaram órfãos. Não apague código morto pré-existente, a menos que solicitado.

### Integridade dos Testes
- **NUNCA** enfraqueça uma asserção de teste existente para fazê-lo passar.
- **NUNCA** delete um teste para reduzir a contagem de falhas.
- **NUNCA** use mecanismos de skip/disable para ignorar testes falhando.
- Os testes são a especificação — a implementação se conforma aos testes, não o contrário.

### Orientação a Objetivos
- Transforme tarefas vagas em objetivos verificáveis.
- Em tarefas multi-etapas, declare um plano breve com pontos de verificação.
- Cada linha alterada deve ter rastreabilidade direta com a solicitação do usuário.

---

## Após Cada Alteração
Pergunte a si mesmo: *"Um engenheiro sênior chamaria isso de sobre-engenharia?"*
Se a resposta for sim → **simplifique antes de prosseguir**.
