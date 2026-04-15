# Decision Frameworks — Referência Completa

Padrões para tomada de decisão e documentação de escolhas técnicas.

---

## 1. Matriz de Trade-offs
Toda escolha técnica possui vantagens e sacrifícios. Esta matriz ajuda a tornar esses sacrifícios explícitos.

| Atributo | Decisão A | Decisão B | Trade-off |
|----------|-----------|-----------|-----------|
| Velocidade de Dev | Alta | Baixa | A é mais rápido agora, mas B é melhor a longo prazo. |
| Escalabilidade | Baixa | Alta | B suporta 10x mais carga que A. |
| Custo Operacional | Baixo | Médio | B exige manutenção de infra extra. |

## 2. Decision Log (Registro de Decisão)
Documente o raciocínio para que outros desenvolvedores (ou você no futuro) entendam o porquê de uma escolha.

### Template:
- **Data**: 2026-04-14
- **Status**: Decidido / Em Discussão
- **Contexto**: Qual era o problema?
- **Opções Consideradas**: Opção 1, Opção 2.
- **Opção Escolhida**: Opção 1.
- **Racional**: Por que a Opção 1 venceu? (ex: Menor risco técnico).

## 3. Understanding Lock (Checklist)
Antes de passar do brainstorming para a especificação, garanta que estes pontos foram validados pelo usuário:

- [ ] **Objetivos**: O que queremos alcançar?
- [ ] **Non-goals**: O que **não** faremos?
- [ ] **Restrições**: Tempo, orçamento, stack tecnológica.
- [ ] **Premissas**: O que estamos assumindo como verdade?
- [ ] **Riscos**: O que pode dar errado?
