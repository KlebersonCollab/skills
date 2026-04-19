# Checklist de Revisão de Código

Este checklist serve como guia para revisões de código focadas em qualidade e manutenibilidade.

## 1. Princípios SOLID
- [ ] **S (Single Responsibility)**: Cada classe/função tem apenas uma razão para mudar?
- [ ] **O (Open/Closed)**: O código está aberto para extensão, mas fechado para modificação?
- [ ] **L (Liskov Substitution)**: Subclasses podem ser usadas no lugar de suas classes base sem quebrar o sistema?
- [ ] **I (Interface Segregation)**: Interfaces são específicas e não forçam implementações desnecessárias?
- [ ] **D (Dependency Inversion)**: O código depende de abstrações e não de implementações concretas?

## 2. YAGNI, DRY e KISS
- [ ] **YAGNI (You Ain't Gonna Need It)**: Existe código que tenta prever o futuro ou resolve problemas que ainda não existem? (Remova se sim)
- [ ] **DRY (Don't Repeat Yourself)**: Existe lógica duplicada que deveria ser abstraída em um único lugar?
- [ ] **KISS (Keep It Simple, Stupid)**: A solução é a mais simples possível para o problema? Existe alguma complexidade desnecessária?

## 3. Qualidade Geral e Legibilidade
- [ ] **Nomes**: Classes, métodos e variáveis têm nomes claros e revelam sua intenção?
- [ ] **Tamanho**: Métodos e classes são pequenos o suficiente para serem compreendidos rapidamente?
- [ ] **Comentários**: O código é autoexplicativo? Existem comentários que explicam "por que" em vez de "o que"?
- [ ] **Tratamento de Erros**: Erros são tratados adequadamente? Existe uso correto de logs/exceções?
- [ ] **Testes**: As alterações estão cobertas por testes automatizados (unitários/integração)?

## 4. Segurança e Performance
- [ ] **Segurança**: Existe algum risco de injeção (SQL, XSS, etc)? Algum dado sensível está sendo exposto indevidamente?
- [ ] **Performance**: Existe algum gargalo óbvio (N+1 queries, loops aninhados pesados, etc)?
