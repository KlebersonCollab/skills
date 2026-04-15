# Pattern Selection — Guia de Decisão

Como escolher o estilo arquitetural correto para o seu contexto.

---

## 1. Monolito Modular vs Microserviços

| Característica | Monolito Modular | Microserviços |
|----------------|------------------|---------------|
| **Custo de Infra** | Baixo | Alto |
| **Complexidade** | Média | Muito Alta |
| **Deployment** | Atômico (Tudo ou nada) | Independente por serviço |
| **Recomendação** | Comece aqui. Ideal para 90% dos casos. | Use apenas se precisar de escala independente extrema. |

## 2. Arquitetura Hexagonal (Ports & Adapters)
Foco em desacoplar a lógica de negócio (Core) do mundo externo (Infra, DB, UI).
- **Vantagem**: Facilidade extrema para trocar DB ou rodar testes unitários rápidos.
- **Quando usar**: Em sistemas com lógica de negócio complexa que precisa ser protegida.

## 3. Clean Architecture
Evolução da Hexagonal, organizando o código em camadas concêntricas.
- **Vantagem**: Independência de frameworks e UI.
- **Trade-off**: Pode gerar excesso de arquivos e "boilerplate" para sistemas simples.

## 4. Event-Driven Architecture (EDA)
Comunicação assíncrona baseada em eventos (Pub/Sub).
- **Vantagem**: Desacoplamento temporal e resiliência.
- **Trade-off**: Dificuldade em rastrear fluxos de dados e depuração.

---

## Árvore de Decisão Rápida

1. **É um sistema novo (Greenfield)?**
   - Sim → Comece com Monolito Modular.
2. **A equipe é pequena (< 10 pessoas)?**
   - Sim → Fuja de Microserviços.
3. **A lógica é complexa?**
   - Sim → Use Arquitetura Hexagonal para o core.
4. **Precisa de alta performance em escrita?**
   - Sim → Considere CQRS ou Event Sourcing.
